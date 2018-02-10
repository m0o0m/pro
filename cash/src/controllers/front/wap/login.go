package wap

import (
	"encoding/json"
	"fmt"
	"framework/render"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
	"time"
)

type LoginController struct {
	WapBaseController
}

func (*LoginController) LoginDo(ctx echo.Context) error {
	member := new(input.MemberSign)
	code := global.ValidRequestMember(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//验证验证码
	codes := ctx.Request().Header.Get("code")
	key, err := getWapMemberRedis(codes)
	if err == redis.Nil {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if key == "" || strings.ToLower(key) != strings.ToLower(member.Code) {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	//删除验证码
	err = global.GetRedis().Del(codes).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取站点信息
	infos, flag, err := GetWapSiteInfo(ctx)

	if err != nil || !flag {
		global.GlobalLogger.Error("error%s", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	system := ctx.Request().Header.Get("platform") //客户端类型
	loginIp := ctx.RealIP()                        //获取登录ip
	//根据站点账号取出对应的会员信息
	info, flag, err := memberBean.GetInfoBySite(infos.SiteId, infos.SiteIndexId, member.Account)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//账号不存在,已经被删除在获取的时候已经被排除出去，或者是该站点下面不存在该账号
	if !flag {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//账号被禁用
	if info.Status == 2 {
		return ctx.JSON(200, global.ReplyError(20002, ctx))
	}
	//加密密码
	password, err := global.MD5ByStr(member.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//密码错误
	if info.Password != password {
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}

	s := fmt.Sprintf("member %d", time.Now().UnixNano()+info.Id)
	//生成tokenkey
	result, err := global.MD5ByBytes([]byte(s), []byte(global.EncryptSalt))
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	data := new(global.MemberRedisToken)
	data.Id = info.Id
	data.Status = info.Status
	data.Account = info.Account
	data.ExpirTime = time.Now().Add(global.AgencyRedisExp).Unix()
	data.Site = infos.SiteId
	data.SiteIndex = infos.SiteIndexId
	data.LevelId = info.LevelId
	data.Type = "member"
	if system == "pc" {
		data.ExpirTime = time.Now().Add(global.DefaultRedisExp.Pc).Unix()
	} else if system == "wap" {
		data.ExpirTime = time.Now().Add(global.DefaultRedisExp.Wap).Unix()
	} else if system == "ios" {
		data.ExpirTime = time.Now().Add(global.DefaultRedisExp.Ios).Unix()
	} else {
		data.ExpirTime = time.Now().Add(global.DefaultRedisExp.Android).Unix()
	}

	//序列化
	b, err := json.Marshal(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	fmt.Println("token参数", result)
	//redis存储
	if system == "pc" {
		err = wapMemberRedisSet(result, b, info.PcLoginKey)
	} else if system == "wap" {
		err = wapMemberRedisSet(result, b, info.WapLoginKey)
	} else if system == "ios" {
		err = wapMemberRedisSet(result, b, info.IosLoginKey)
	} else {
		err = wapMemberRedisSet(result, b, info.AndroidLoginKey)
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//刚刚是否存储成功
	s, err = getMemberRedis(result)
	fmt.Println("取出来的redis：", s)
	if err != nil || s == "" {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//刷新数据库
	memberLogin := new(schema.Member)
	memberLogin.Id = info.Id
	memberLogin.LastLoginTime = info.LoginTime
	count, err := memberBean.RefreshMember(memberLogin, result, system, loginIp)
	if err != nil || count != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//获取会员上级代理基本信息
	agencyInfo, _, err := thirdAgencyBean.BaseInfo(&input.ThirdAgencyInfo{infos.SiteId, infos.SiteIndexId, info.ThirdAgencyId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//返回数据
	backData := new(back.MemberSignBack)
	backData.Id = info.Id
	backData.Type = system
	backData.Account = info.Account
	backData.SiteIndexId = infos.SiteIndexId
	backData.SiteId = infos.SiteId
	backData.Status = info.Status
	backData.LevelId = info.LevelId
	backData.Token = result
	backData.ThirdAgencyId = info.ThirdAgencyId
	backData.ThirdAgencyAccount = agencyInfo.Account
	return ctx.JSON(200, global.ReplyItem(backData))
}

//退出登录
func (*LoginController) Logout(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	token := ctx.Get("token").(string)
	system := ctx.Request().Header.Get("platform")
	//更改redis
	err := global.GetRedis().Del(token).Err()
	if err != nil {
		global.GlobalLogger.Error("Error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	err = global.GetRedis().LPop(token).Err()

	//修改数据库
	count, err := memberBean.ChangeLoginkey(member.Id, system)
	if err != nil || count != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//获取登录的时候存储的redis值
func getWapMemberRedis(token string) (string, error) {
	key, err := global.GetRedis().Get(token).Result()
	return key, err
}

//根据请求域名取得站点信息
func GetWapSiteInfo(ctx echo.Context) (info schema.SiteDomain, flag bool, err error) {
	requesturi := ctx.Request().Header.Get("domain") //获取不到数据
	//system := ctx.Request().Header.Get("Platform")   //获取不到数据
	//根据host去获取站点信息
	requesturi = "localhost"
	info, flag, err = siteDomainBean.GetSiteByDomain(requesturi)
	return
}

//redis 存储
func wapMemberRedisSet(result string, b []byte, beforeKey string) (err error) {
	if beforeKey != "" {
		//删除旧的key
		err = global.GetRedis().Del(beforeKey).Err()
		//将旧的删除
		err = global.GetRedis().LPop(result).Err()
	}
	//存储新token
	err = global.GetRedis().Set(result, b, 0).Err()
	//将推进list
	err = global.GetRedis().RPush("member_login", result).Err()
	return err
}

//登陆验证
func (c *LoginController) GetAjaxLoginVerify(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	Time := time.Now().Unix()
	if member.ExpirTime < Time {
		return ctx.JSON(200, global.ReplyError(91201, ctx))
	}

	mi := new(input.MemberInfoSelf)
	mi.Id = member.Id
	mi.SiteId = member.Site
	mi.SiteIndexId = member.SiteIndex
	data, _, err := baseInfoBean.MemberSelfInfo(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}

	sdata := back.AjaxLoginIn{}
	sdata.Id = data.Id
	sdata.Account = data.Account
	sdata.Realname = data.Realname
	sdata.Balance = data.Balance

	platformBalance, err := memberBalanceConversionBean.GetPlatformBalance(member.Id)
	var da back.MemberBalanceTotalBack
	for _, v := range platformBalance.ProductClassifyBalance {
		da.Type = 1
		da.Balance = v.Balance
		da.Name = v.Platform
		sdata.TBalance = append(sdata.TBalance, da)
	}
	sdata.TBalance = append(sdata.TBalance, back.MemberBalanceTotalBack{"账户余额", platformBalance.AccountBalance, 2})
	dad := platformBalance.AccountBalance + platformBalance.GameBalance
	sdata.TBalance = append(sdata.TBalance, back.MemberBalanceTotalBack{"账户总余额", dad, 3})

	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	times := new(global.Times)
	timeNow := global.GetCurrentTime()
	times.StartTime = timeNow - 7*24*3600
	times.EndTime = timeNow
	sdata.Count, err = memMessageBean.MemMessageCount(member.Site, member.SiteIndex, member.Id, times)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	if sdata.Realname == "" {
		mbank, has, err := MemberBankBean.GetMemberBankOne(member.Id)
		if has && err == nil {
			_, err = memberBean.UpdateMemberReallname(member.Site, member.SiteIndex, member.Account, mbank.CardName)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
			}
			sdata.Realname = mbank.CardName
		}
	}
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//修改密码
func (msc *LoginController) EditPassword(ctx echo.Context) error {
	newPassword := new(input.ChangePassword)
	code := global.ValidRequestMember(newPassword, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	//根据id取出会员信息
	info, flag, err := memberBean.GetInfoById(member.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60051, ctx))
	}
	withdraw, err := global.MD5ByStr(newPassword.BeforePassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//旧密码错误
	if newPassword.Types == 1 { //登录密码
		if withdraw != info.Password {
			return ctx.JSON(200, global.ReplyError(20009, ctx))
		}
	} else if newPassword.Types == 2 { //取款密码
		if withdraw != info.DrawPassword {
			return ctx.JSON(200, global.ReplyError(20009, ctx))
		}
	}

	//新密码两次输入不一致
	if newPassword.Password != newPassword.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}

	//判断新密码是否和旧密码一致
	if newPassword.BeforePassword == newPassword.Password {
		return ctx.JSON(200, global.ReplyError(60214, ctx))
	}
	password, err := global.MD5ByStr(newPassword.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	_, err = memberBean.ChangePassword(member.Id, password, newPassword.Types)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}
