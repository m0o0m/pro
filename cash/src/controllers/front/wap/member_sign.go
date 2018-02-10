package wap

import (
	"encoding/json"
	"fmt"
	"global"
	"models/back"
	"models/function"
	"models/input"
	"models/schema"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"strconv"
)

type SignController struct {
	WapBaseController
}

//登录
func (*SignController) Login(ctx echo.Context) error {
	member := new(input.MemberSign)
	code := global.ValidRequestMember(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//验证验证码
	codes := ctx.Request().Header.Get("code")
	key, err := getMemberRedis(codes)
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
	infos, flag, err := GetSiteInfo(ctx)

	if err != nil || !flag {
		global.GlobalLogger.Error("error%s", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//todo 是否需要判断站点状态
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
	//fmt.Println("token参数", result)
	//redis存储
	if system == "pc" {
		err = memberRedisSet(result, b, info.PcLoginKey)
	} else if system == "wap" {
		err = memberRedisSet(result, b, info.WapLoginKey)
	} else if system == "ios" {
		err = memberRedisSet(result, b, info.IosLoginKey)
	} else {
		err = memberRedisSet(result, b, info.AndroidLoginKey)
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//刚刚是否存储成功
	s, err = getMemberRedis(result)
	//fmt.Println("取出来的redis：", s)
	if err != nil || s == "" {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//刷新数据库
	memberLogin := new(schema.Member)
	memberLogin.Id = info.Id
	count, err := memberBean.RefreshMember(memberLogin, result, system, loginIp)
	if err != nil || count != 1 {
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
	return ctx.JSON(200, global.ReplyItem(backData))
}

//注册
func (*SignController) Register(ctx echo.Context) error {
	register := new(input.MemberRegister)
	code := global.ValidRequest(register, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	domain := ctx.Request().Host
	damainurl := strings.Split(domain, ":")
	siteinfo, flag, err := siteDomainBean.GetSiteByDomain(damainurl[0])
	//根据站点信息获取会员注册设定
	regis, flag, err := memberRegisterSettingBean.GetOneSet(siteinfo.SiteId, siteinfo.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60058, ctx))
	}
	//是否开启会员注册
	if regis.IsReg == 2 {
		return ctx.JSON(200, global.ReplyError(60060, ctx))
	}
	if register.Account == "" {
		return ctx.JSON(200, global.ReplyError(30009, ctx))
	}
	if register.Password == "" {
		return ctx.JSON(200, global.ReplyError(30010, ctx))
	}
	if register.ConfirmPassword == "" {
		return ctx.JSON(200, global.ReplyError(30011, ctx))
	}
	//验证验证码
	if regis.IsCode == 1 {
		codes := ctx.Request().Header.Get("code")
		key, err := getMemberRedis(codes)
		if err == redis.Nil {
			return ctx.JSON(200, global.ReplyError(20021, ctx))
		}
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if key == "" || strings.ToLower(key) != strings.ToLower(register.Code) {
			return ctx.JSON(200, global.ReplyError(20021, ctx))
		}
		//删除验证码
		err = global.GetRedis().Del(codes).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//是否同意协议
	IsAgreeDeal, _ := strconv.ParseInt(register.IsAgreeDeal, 10, 64)
	if IsAgreeDeal == 2 {
		return ctx.JSON(200, global.ReplyError(60070, ctx))
	}

	//判断账号是否重复
	memberBean := new(function.MemberBean)
	flag, err = memberBean.CheckIsExist(register.Account, regis.SiteId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag {
		return ctx.JSON(200, global.ReplyError(60057, ctx))
	}
	//两次密码是否一致
	if register.Password != register.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}
	//登录密码和取款密码加密
	pass, err := global.MD5ByStr(register.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	creatTime := global.GetCurrentTime()
	member := new(schema.Member)
	member.SiteIndexId = siteinfo.SiteIndexId //站点前台id
	member.SiteId = siteinfo.SiteId           //站点id
	member.Account = register.Account         //账号
	member.Password = pass                    //登录密码
	member.IsLockedLevel = 2                  //是否锁定会员层级
	member.Status = 1                         //状态
	member.RegisterIp = ctx.RealIP()          //注册ip
	member.IsEditPassword = 1                 //是否可以修改密码
	member.CreateTime = creatTime             //创建时间
	member.IsAgreeDeal = IsAgreeDeal          //是否同意本平台协议
	member.WapStatus = 2
	member.PcStatus = 2
	member.IosStatus = 2
	member.AndroidStatus = 2

	system := ctx.Request().Header.Get("platform") //注册平台

	if system == "pc" {
		member.RegisterClientType = 1
	} else if system == "wap" {
		member.RegisterClientType = 2
	} else if system == "ios" {
		member.RegisterClientType = 3
	} else {
		member.RegisterClientType = 4
	}
	//默认分层取得
	info, _, err := memberLevelBean.SiteDefault(siteinfo.SiteId, siteinfo.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	member.LevelId = info.LevelId
	var isDiscount int
	//会员注册优惠计算
	if regis.IsIp == 1 { //开启ip限制,那么同一个站点下面的同一个ip只能获取一次优惠，同ip再次注册无法获取优惠
		ip := ctx.RealIP()
		flag, err := memberBean.GetRecordByIp(siteinfo.SiteId, siteinfo.SiteIndexId, ip)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if flag == false {
			member.Balance = regis.Offer
			isDiscount = 1
		} else {
			isDiscount = 2
		}
	} else { //没有开启,直接给优惠
		member.Balance = regis.Offer
		isDiscount = 1
	}
	//所属代理商设置
	err = agencyDef(member, register, ctx, regis)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	_, err = memberBean.MemberRegister(member, register, regis, isDiscount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	loginIp := ctx.RealIP() //获取登录ip

	s := fmt.Sprintf("member %d", time.Now().UnixNano()+member.Id)
	//生成tokenkey
	result, err := global.MD5ByBytes([]byte(s), []byte(global.EncryptSalt))
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	data := new(global.MemberRedisToken)
	data.Id = member.Id
	data.Status = member.Status
	data.Account = member.Account
	data.ExpirTime = time.Now().Add(global.AgencyRedisExp).Unix()
	data.Site = member.SiteId
	data.SiteIndex = member.SiteIndexId
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
	//fmt.Println("token参数", result)
	//redis存储
	if system == "pc" {
		err = memberRedisSet(result, b, member.PcLoginKey)
	} else if system == "wap" {
		err = memberRedisSet(result, b, member.WapLoginKey)
	} else if system == "ios" {
		err = memberRedisSet(result, b, member.IosLoginKey)
	} else {
		err = memberRedisSet(result, b, member.AndroidLoginKey)
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//刚刚是否存储成功
	s, err = getMemberRedis(result)
	//fmt.Println("取出来的redis：", s)
	if err != nil || s == "" {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//刷新数据库
	memberLogin := new(schema.Member)
	memberLogin.Id = member.Id
	memberLogin.LastLoginTime = creatTime
	count, err := memberBean.RefreshMember(memberLogin, result, system, loginIp)
	if err != nil || count != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//返回数据
	backData := new(back.MemberSignBack)
	backData.Id = member.Id
	backData.Type = system
	backData.Account = member.Account
	backData.SiteIndexId = member.SiteIndexId
	backData.SiteId = member.SiteId
	backData.Status = member.Status
	backData.Token = result
	return ctx.JSON(200, global.ReplyItem(backData))
}

//退出登录
func (*SignController) Logout(ctx echo.Context) error {
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

//修改密码
func (msc *SignController) EditPassword(ctx echo.Context) error {
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

////根据请求域名取得站点信息
//func GetSiteInfo(ctx echo.Context) (info schema.SiteDomain, flag bool, err error) {
//	requesturi := ctx.Request().Header.Get("domain")//获取不到数据
//	system := ctx.Request().Header.Get("Platform")//获取不到数据
//	//根据host去获取站点信息
//	requesturi="localhost"
//	system="pc"
//	fmt.Println()
//	if system == "pc" {
//		info, flag, err = siteDomainBean.GetSiteByDomain(requesturi, 1)
//	} else if system == "wap" {
//		info, flag, err = siteDomainBean.GetSiteByDomain(requesturi, 2)
//	}
//	return
//}

//根据请求域名取得站点信息
func GetSiteInfo(ctx echo.Context) (info schema.SiteDomain, flag bool, err error) {
	arr := strings.Split(ctx.Request().Header.Get("Origin")[7:], ":")
	requestDomain := arr[0]
	info, flag, err = siteDomainBean.GetSiteByDomain(requestDomain)
	return
}

//redis 存储
func memberRedisSet(result string, b []byte, beforeKey string) (err error) {
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

//获取登录的时候存储的redis值
func getMemberRedis(token string) (string, error) {
	key, err := global.GetRedis().Get(token).Result()
	return key, err
}

//点击注册的时候根据域名获取会员注册信息和站点
func (*SignController) GetMemberRegister(ctx echo.Context) error {
	info, flag, err := GetSiteInfo(ctx)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60059, ctx))
	}
	//根据站点去取的会员注册设定
	backdata, flag, err := memberRegisterSettingBean.GetOneSet(info.SiteId, info.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60058, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(backdata))
}

//校验注册的时候某项开启的项是否重复
func (*SignController) CheckBankCardExist(ctx echo.Context) error {
	//1.银行卡号是否可以重复2.电话号码是否可以重复3.邮箱是否可以重复4.qq是否可以重复
	//5.微信号是否重复
	checkRegister := new(input.CheckRegister)
	code := global.ValidRequestMember(checkRegister, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	memberBean := new(function.MemberBean)
	flag, err := memberBean.CheckIsExistValue(checkRegister.Types, checkRegister.Conditions)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyItem(1))
	}
	return ctx.NoContent(204)
}

//默认站点线处理
func defaultAgency(ctx echo.Context, regis schema.SiteMemberRegisterSet, member *schema.Member) error {
	thirdAgency, _, err := agencyBean.GetDefault(regis.SiteId, regis.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	member.ThirdAgencyId = thirdAgency.Id
	member.SecondAgencyId = thirdAgency.ParentId //二级
	agency, _, err := agencyBean.GetAgency(thirdAgency.ParentId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	member.FirstAgencyId = agency.ParentId //一级
	return nil
}

//所属代理商判定
func agencyDef(member *schema.Member, register *input.MemberRegister, ctx echo.Context, regis schema.SiteMemberRegisterSet) error {
	Introducer, _ := strconv.ParseInt(register.Introducer, 10, 64)
	IntroducerMember, _ := strconv.ParseInt(register.IntroducerMember, 10, 64)
	if Introducer != 0 { //代理链接直接注册
		info, flag, err := memberBean.GetOneMemberByThird(Introducer) //取得三级经销商名下一个会员，根据该会员经销商获取信息
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if flag { //有属于这个代理的会员
			member.ThirdAgencyId = Introducer           //代理商
			member.SecondAgencyId = info.SecondAgencyId //二级
			member.FirstAgencyId = info.FirstAgencyId   //一级
		} else { //没有属于这个代理的会员,从代理表查找
			info, flag, err := agencyBean.GetAgency(Introducer)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return err
			}
			if !flag { //该代理已经不存在或者状态异常，归入默认代理线
				err := defaultAgency(ctx, regis, member)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					return err
				}
			} else { //正常
				member.ThirdAgencyId = Introducer     //三级
				member.SecondAgencyId = info.ParentId //二级经销商
				agency, _, err := agencyBean.GetAgency(info.ParentId)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					return err
				}
				member.FirstAgencyId = agency.ParentId //一级
			}
		}
	} else if IntroducerMember != 0 { //会员推广链接注册
		info, flag, err := memberBean.GetInfoById(IntroducerMember)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		//没有找到这个推荐人，找到该站点默认代理线
		if !flag {
			err := defaultAgency(ctx, regis, member)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return err
			}
		} else {
			member.SpreadId = IntroducerMember          //推广人id
			member.SpreadMoney = 0                      //推广获利
			member.ThirdAgencyId = info.ThirdAgencyId   //代理商Id
			member.SecondAgencyId = info.SecondAgencyId //二级代理商id
			member.FirstAgencyId = info.FirstAgencyId   //一级代理商Id
		}
	} else if IntroducerMember == 0 && Introducer == 0 { //自己直接过来注册，归属于站点默认的层级
		err := defaultAgency(ctx, regis, member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
	}
	return nil
}
