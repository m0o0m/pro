package account

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"time"

	"framework/validation"
	"github.com/go-redis/redis"
	"global"
	"models/back"
	"models/function"
	"models/input"
	"models/schema"
	"strings"
)

type AgencySignController struct {
}

//登录
func (asc *AgencySignController) Login(ctx echo.Context) error {
	login := new(input.AgencySignLogin)
	if err := ctx.Bind(login); err != nil {
		return ctx.JSON(200, global.ReplyError(10000, ctx))
	}
	valid := validation.Validation{}
	ok, err := valid.Valid(login)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		for _, e := range valid.Errors {
			return ctx.JSON(200, global.ReplyError(e.Code(), ctx))
		}
	}
	//验证验证码
	code := ctx.Request().Header.Get("code")
	key, err := GetTokenS(code)
	if err == redis.Nil {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if key == "" || strings.ToLower(key) != strings.ToLower(login.Code) {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	//删除验证码
	err = global.GetRedis().Del(code).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//根据域名查询出站点id
	host := ctx.Request().Host
	SiteInfoBean := new(function.SiteDomainBean)
	siteDomin := strings.Split(host, ":")[0]
	if len(siteDomin) == 0 {
		global.GlobalLogger.Error("Login get host is nil!")
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	siteInfo, flag, err := SiteInfoBean.GetSiteInfoByDomian(siteDomin)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}
	//未查询到域名
	if !flag {
		global.GlobalLogger.Error("Login GetSiteInfoByDomian error: %s", siteDomin)
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}
	//根据登录账号和密码查询出账号
	ok, err, agency := agencySignBean.Login(login, siteInfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//未查询到账号
	if !ok {
		global.GlobalLogger.Error("Login getusername error:%v", ok)
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}
	//加密密码
	login.Password, err = global.MD5ByStr(login.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//账号被删除
	if agency.DeleteTime != 0 {
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}
	//账号被禁用
	if agency.Status != 1 {
		return ctx.JSON(200, global.ReplyError(20002, ctx))
	}
	//密码错误
	if agency.Password != login.Password {
		//更新登录错误次数
		agency.LoginErrCount += 1
		_, err := agencySignBean.LoginErrUpdate(agency)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}

	//获取账号角色来确定在ip开关数据表中取哪种类型数据
	var accountType int
	if agency.RoleId == 1 {
		accountType = 1
	} else {
		accountType = 2
	}
	ip := ctx.RealIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	//验证是否是ip限制中
	ok, err = ipSetBean.IpCheck(accountType, agency.SiteId, ip)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(30250, ctx))
	}

	//获取当前登录代理所属站点是否设置登录口令
	subAccountTokenInfo, has, err := subAccountBeen.SubAccessTokenInfo(agency.SiteId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//如果该站点设置过登录口令
	if has && subAccountTokenInfo.Status == 1 {
		if len(login.VerifyCode) != 6 {
			//请输入口令验证
			return ctx.JSON(200, global.ReplyError(20022, ctx))
		}
		passKey, err := global.ParseBase64Str(subAccountTokenInfo.PassKey, agency.SiteId)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		has := agencySignBean.VerifyCode(passKey, login.VerifyCode)
		if !has {
			return ctx.JSON(200, global.ReplyError(20023, ctx))
		}
	}

	data := new(global.RedisStruct)
	var accessSID []string
	//如果是开户人或者开户人子账号
	if agency.RoleId == 1 {
		//如果是子账号
		if agency.IsSub == 2 {
			//获取子账号细分权限
			ok, dm, err := subAccountBeen.GetDetailPermission(agency.Id)
			if err != nil {
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			//如果查询不到或者设置的站点权限为空,那么可控制的站点是所属开户人下所有子站
			if !ok || dm.ChildSite == "" {
				sob := new(function.SiteOperateBean)
				s := new(schema.Site)
				s.Id = agency.SiteId
				sites, err := sob.GetSiteIndexId(s)
				if err != nil {
					return ctx.JSON(500, global.ReplyError(60000, ctx))
				}
				for _, site := range sites {
					accessSID = append(accessSID, site.SiteIndexId)
				}
			} else {
				sites := strings.Split(dm.ChildSite, ",")
				accessSID = append(accessSID, sites...)
				if len(sites) == 1 {
					//如果子账号站点权限只有一个,那么该子账号的站点前台id就是这一个
					agency.SiteIndexId = sites[0]
				}
			}
		} else {
			//如果不是子账号,那么可控制站点是当前登录开户人所有子站
			sob := new(function.SiteOperateBean)
			s := new(schema.Site)
			s.Id = agency.SiteId
			sites, err := sob.GetSiteIndexId(s)
			if err != nil {
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			for _, site := range sites {
				accessSID = append(accessSID, site.SiteIndexId)
			}
		}
	} else {
		//如果是代理,那么当前登录账号可控制站点只有所属子站
		accessSID = append(accessSID, agency.SiteIndexId)
	}

	data.Id = agency.Id
	data.Account = agency.Account
	data.SiteId = agency.SiteId
	data.SiteIndexId = agency.SiteIndexId
	data.RoleId = agency.RoleId
	data.ExpirTime = time.Now().Add(global.AgencyRedisExp).Unix()
	data.IsSub = agency.IsSub
	data.Username = agency.Username
	data.Level = agency.Level
	data.Type = "agency"
	data.AccessSID = accessSID

	s := fmt.Sprintf("agency %d", time.Now().UnixNano()+agency.Id)
	//生成tokenkey
	result, err := global.MD5ByBytes([]byte(s), []byte(global.EncryptSalt))
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//序列化
	b, err := json.Marshal(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	res := new(back.AgencySign)
	res.Token = result
	res.Username = agency.Username
	res.SiteId = agency.SiteId
	res.SiteIndexId = agency.SiteIndexId

	//redis存储
	err = tokenSetup(result, b, agency.LoginKey)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	s, err = GetTokenS(result)
	if err != nil || s == "" {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//更新登录相关字段
	agency.IsLogin = 1
	agency.LastLoginIp = agency.LoginIp
	agency.LastLoginTime = agency.LoginTime
	agency.LoginIp = ip
	agency.LoginTime = time.Now().Unix()
	agency.LoginCount += 1
	agency.LoginKey = result //存储序列化之后的值

	row, err := agencySignBean.LoginUpdate(agency)
	if err != nil || row != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(res))
}

//退出
func (asc *AgencySignController) Logout(ctx echo.Context) error {
	userinfo := ctx.Get("user").(*global.RedisStruct)
	token := ctx.Get("token").(string)
	//更改redis
	global.GetRedis().Del(token)
	//弹出list
	global.GetRedis().LPop(token).Err()
	//更改数据库登录状态
	err := agencySignBean.Logout(userinfo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//修改密码
func (asc *AgencySignController) SetPwd(ctx echo.Context) error {
	set_pwd := new(input.AgencySignPassword)
	code := global.ValidRequest(set_pwd, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//两次密码不一致
	if set_pwd.NewPassword != set_pwd.ReplyPassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}
	//判断新密码是否和旧密码一致
	if set_pwd.NewPassword == set_pwd.OldPassword {
		return ctx.JSON(200, global.ReplyError(60214, ctx))
	}
	userinfo := ctx.Get("user").(*global.RedisStruct)
	old_pwd, err := global.MD5ByStr(set_pwd.OldPassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//验证原密码
	ok, err := agencySignBean.ValidPassword(old_pwd, userinfo.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//原密码错误
	if !ok {
		return ctx.JSON(200, global.ReplyError(20009, ctx))
	}

	//加密新密码
	pwd, err := global.MD5ByStr(set_pwd.NewPassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	_, err = agencySignBean.UpdatePassword(userinfo.Id, pwd)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//操作redis
//存储redis值，并且设置该token值的过期时间
func tokenSetup(token string, jsonUser []byte, beforeKey string) (err error) {
	if beforeKey != "" {
		//删除旧的key
		err = global.GetRedis().Del(beforeKey).Err()
		if err == redis.Nil {
		} else if err != nil {
			return
		}
		//将旧的删除
		err = global.GetRedis().LPop(token).Err()
		if err == redis.Nil {
		} else if err != nil {
			return
		}
	}
	//存储新token
	err = global.GetRedis().Set(token, jsonUser, 0).Err()
	if err == redis.Nil {
	} else if err != nil {
		return
	}
	//将推进list
	err = global.GetRedis().RPush("is_login", token).Err()
	if err == redis.Nil {
	} else if err != nil {
		return
	}
	return nil
}

//获取登录的时候存储的redis值
func GetTokenS(token string) (string, error) {
	key, err := global.GetRedis().Get(token).Result()
	return key, err
}

//根据登录人角色获取菜单
func (asc *AgencySignController) Menu(ctx echo.Context) error {
	user := ctx.Get("user").(*global.RedisStruct)
	menu, count, err := roleMenuBean.GetMenuByRoleId(user.RoleId, "agency")
	if err != nil {
		return ctx.JSON(500, global.ReplyError(50012, ctx))
	}
	var data []back.Trees
	if count > 1 {
		data = AndLevel(menu, 0)
	}
	return ctx.JSON(200, global.ReplyCollection(data, count))
}

//递归，菜单无限级目录树
func AndLevel(data []back.MenuListBack, parentid int64) []back.Trees {
	//递归调用当所有的循环没有完成的时候是没有进行child的存值操作
	var lend = 0
	var x = 0
	//这里是为了计算我存储数据的slice的长度
	for _, v := range data {
		if v.ParentId == parentid {
			lend = lend + 1
		}
	}
	//这里根据上面取得的长度定义slice
	var tree []back.Trees = make([]back.Trees, lend)
	if lend != 0 {
		for k, v := range data {
			//这里的k是不定的，所以需要定义另外的累加值进行累加计数
			//将计数累加放在这里会导致数组越界，因为没有满足条件，循环次数会超过上面定义的slice的长度
			if v.ParentId == parentid {
				k = x
				x = x + 1
				//满足条件赋值
				tree[k].MenuName = v.MenuName
				tree[k].Icon = v.Icon
				tree[k].Sort = v.Sort
				tree[k].Route = v.Route
				tree[k].Id = v.Id
				tree[k].Status = v.Status
				tree[k].Type = v.Type
				tree[k].Level = v.Level
				tree[k].LanguageKey = v.LanguageKey
				//下级菜单的个数不定所以这里更改id值和层级 循环再次调用自己
				child := AndLevel(data, v.Id)
				//将取出来的值赋值给子项
				tree[k].Children = child
			}
		}
	}
	return tree
}
