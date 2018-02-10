package login

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"controllers"
	"global"
	"models/back"
	"models/input"
	"models/schema"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

type WapLoginController struct {
	controllers.BaseController
}

//MemberLogin wap会员登录
func (*WapLoginController) MemberLogin(ctx echo.Context) error {
	memberLogin := new(input.MemberLogin)
	code := global.ValidRequestMember(memberLogin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//其他需要的参数获取
	system := ctx.Request().Header.Get("platform") //客户端类型
	//如果不是安卓或者ios则返回404
	if system != "android" && system != "ios" {
		return echo.ErrNotFound
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
	if key == "" || strings.ToLower(key) != strings.ToLower(memberLogin.Code) {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	//删除验证码
	err = global.GetRedis().Del(codes).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//会员是否存在,根据账号站点及站点前台id查询会员
	memberInfo, flag, err := memberBean.GetInfoBySite(memberLogin.SiteId, memberLogin.SiteIndexId, memberLogin.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//账号不存在,已经被删除在获取的时候已经被排除出去，或者是该站点下面不存在该账号
	if !flag {
		return ctx.JSON(200, global.ReplyError(60051, ctx))
	}
	//账号被禁用
	if memberInfo.Status != 1 {
		return ctx.JSON(200, global.ReplyError(20002, ctx))
	}
	//加密密码
	password, err := global.MD5ByStr(memberLogin.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//记录日志
	log := new(schema.LoginLog)
	log.Domains = "" //app没有域名
	var device int8
	if system == "android" {
		device = 3
	} else {
		device = 4
	}
	log.Device = device
	log.Account = memberInfo.Account
	log.LoginIp = ctx.RealIP()
	log.SiteId = memberInfo.SiteId
	log.SiteIndexId = memberInfo.SiteIndexId
	log.LoginRole = 1

	//密码错误
	if memberInfo.Password != password {
		//插入一条登录日志
		log.LoginResult = 2
		log.LoginTime = time.Now().Unix()
		r, err := memberBean.AddLoginLog(log)
		if err != nil || r != 1 {
			global.GlobalLogger.Error("error:%s", err.Error())
		}
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}

	s := fmt.Sprintf(system+"member%s", time.Now().UnixNano()+memberInfo.Id)
	//生成token
	result, err := global.MD5ByBytes([]byte(s), []byte(global.EncryptSalt))
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	loginIp := ctx.RealIP() //获取登录ip

	//存储数据
	memberRes := new(global.MemberRedisToken)
	memberRes.Id = memberInfo.Id
	memberRes.Account = memberInfo.Account
	memberRes.Status = memberInfo.Status
	memberRes.Site = memberInfo.SiteId
	memberRes.SiteIndex = memberInfo.SiteIndexId
	memberRes.LevelId = memberInfo.LevelId
	memberRes.Type = "member" + system
	if system == "android" {
		memberRes.ExpirTime = time.Now().Add(global.DefaultRedisExp.Android).Unix()
	} else {
		memberRes.ExpirTime = time.Now().Add(global.DefaultRedisExp.Ios).Unix()
	}
	//序列化
	b, err := json.Marshal(memberRes)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//redis存储
	if system == "android" {
		err = memberRedisSet(result, b, memberInfo.AndroidLoginKey)
	} else {
		err = memberRedisSet(result, b, memberInfo.IosLoginKey)
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取登录的时候存储的redis值
	redsiData, err := getMemberRedis(result)
	if err != nil || redsiData == "" {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//刷新数据库
	members := new(schema.Member)
	members.Id = memberInfo.Id
	count, err := memberBean.RefreshMember(members, result, system, loginIp)
	if err != nil || count != 1 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//插入一条登录日志
	log.LoginResult = 1
	log.LoginTime = time.Now().Unix()
	r, err := memberBean.AddLoginLog(log)
	if err != nil || r != 1 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//返回数据
	backData := new(back.MemberSignBack)
	backData.Id = memberInfo.Id
	backData.Type = system
	backData.Account = memberInfo.Account
	backData.SiteIndexId = memberInfo.SiteIndexId
	backData.SiteId = memberInfo.SiteId
	backData.Status = memberInfo.Status
	backData.Token = result
	return ctx.JSON(200, global.ReplyItem(backData))
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
	err = global.GetRedis().RPush("app_member", result).Err()
	return err
}

func getMemberRedis(token string) (string, error) {
	key, err := global.GetRedis().Get(token).Result()
	return key, err
}

//MemberLogout 退出登录
func (*WapLoginController) MemberLogout(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	token := ctx.Get("token").(string)
	//更改redis
	err := global.GetRedis().Del(token).Err()
	if err != nil {
		global.GlobalLogger.Error("Error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//弹出list
	err = global.GetRedis().LPop(token).Err()
	if err != nil {
		global.GlobalLogger.Error("Error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//修改数据库
	count, err := memberBean.ChangeLoginkey(member.Id, "wap")
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(20004, ctx))
	}
	return ctx.NoContent(204)
}

//MemberRegister 会员注册 todo 会员推广的话是否需要更新会员推广获利金额，从哪里获取推广一个返利多少
func (*WapLoginController) MemberRegister(ctx echo.Context) error {
	member_register := new(input.WapMemberRegister)
	code := global.ValidRequestMember(member_register, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//邮箱校验
	//if member_register.Email != "" {
	//	has := global.CheckEmail(member_register.Email)
	//	if !has {
	//		return ctx.JSON(200, global.ReplyError(20015, ctx))
	//	}
	//
	//}
	//qq验证
	//if member_register.Qq != "" {
	//	has := global.Checkqq(member_register.Qq)
	//	if !has {
	//		return ctx.JSON(200, global.ReplyError(20016, ctx))
	//	}
	//}
	//银行卡验证
	//if member_register.BankCard != "" {
	//	has := global.CheckCardNumber(member_register.BankCard)
	//	if !has {
	//		return ctx.JSON(200, global.ReplyError(60200, ctx))
	//	}
	//}
	//手机验证
	//if member_register.Phone != "" {
	//	has := global.CheckPhoneNumber(member_register.Phone)
	//	if !has {
	//		return ctx.JSON(200, global.ReplyError(20014, ctx))
	//	}
	//}
	//根据站点信息获取会员注册设定
	regis, flag, err := memberRegisterSetup.GetOneSet(member_register.Site, member_register.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60058, ctx))
	}
	//是否开启会员注册
	//if regis.IsReg == 2 {
	//	return ctx.JSON(200, global.ReplyError(60060, ctx))
	//}
	//验证验证码
	//if regis.IsCode == 1 {
	//	codes := ctx.Request().Header.Get("code")
	//	key, err := getMemberRedis(codes)
	//	if err == redis.Nil {
	//		return ctx.JSON(200, global.ReplyError(20021, ctx))
	//	}
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	if key == "" || strings.ToLower(key) != strings.ToLower(member_register.Code) {
	//		return ctx.JSON(200, global.ReplyError(20021, ctx))
	//	}
	//	//删除验证码
	//	err = global.GetRedis().Del(codes).Err()
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//}
	//是否同意协议
	if member_register.IsAcceptAgreement == 2 {
		return ctx.JSON(200, global.ReplyError(60070, ctx))
	}
	//判断账号是否重复
	//flag, err = memberBean.CheckIsExist(member_register.Account, regis.SiteId)
	//if err != nil {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}
	//if flag {
	//	return ctx.JSON(200, global.ReplyError(60057, ctx))
	//}
	//两次密码是否一致
	if member_register.Password != member_register.RepeatPassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}
	//邮箱不能重复
	//if regis.Email == 1 && regis.IsEmail == 1 {
	//	if member_register.Email == "" {
	//		return ctx.JSON(200, global.ReplyError(60063, ctx))
	//	}
	//	flag, err := memberBean.CheckIsExistValue(3, member_register.Email)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	if flag {
	//		return ctx.JSON(200, global.ReplyError(60064, ctx))
	//	}
	//}
	//电话号码不能重复
	//if regis.Mobile == 1 && regis.IsTel == 1 {
	//	if member_register.Phone == "" {
	//		return ctx.JSON(200, global.ReplyError(60065, ctx))
	//	}
	//	flag, err := memberBean.CheckIsExistValue(2, member_register.Phone)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	if flag {
	//		return ctx.JSON(200, global.ReplyError(60066, ctx))
	//	}
	//}
	//qq不能重复
	//if regis.Qq == 1 && regis.IsQq == 1 {
	//	if member_register.Qq == "" {
	//		return ctx.JSON(200, global.ReplyError(60067, ctx))
	//	}
	//	flag, err := memberBean.CheckIsExistValue(4, member_register.Qq)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	if flag {
	//		return ctx.JSON(200, global.ReplyError(60068, ctx))
	//	}
	//}
	//银行卡号是否可以重复
	//if regis.IsCardReply == 1 {
	//	flag, err := memberBean.CheckIsExistValue(1, member_register.BankCard)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	if flag {
	//		return ctx.JSON(200, global.ReplyError(60069, ctx))
	//	}
	//}
	//需要身份号
	//if regis.Passport == 1 {
	//	if member_register.PassPort == "" {
	//		return ctx.JSON(200, global.ReplyError(60061, ctx))
	//	}
	//}
	//需要生日
	//if regis.Birthday == 1 {
	//	if member_register.Birthday == "" {
	//		return ctx.JSON(200, global.ReplyError(60062, ctx))
	//	}
	//}
	//姓名是否重复
	//if regis.IsName == 1 {
	//	flag, err := memberBean.CheckRealNameExist(member_register.RealName)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	if flag {
	//		return ctx.JSON(200, global.ReplyError(60071, ctx))
	//	}
	//}
	//登录密码加密
	pass, err := global.MD5ByStr(member_register.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//操作密码加密
	//withdraw, err := global.MD5ByStr(member_register.OperatePassword, global.EncryptSalt)
	//if err != nil {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}

	member := new(schema.Member)
	member.SiteIndexId = member_register.SiteIndexId //站点前台id
	member.SiteId = member_register.Site             //站点id
	member.Account = member_register.Account         //账号
	member.Password = pass                           //登录密码
	//member.DrawPassword = member_register.OperatePassword  //取款密码
	//member.Realname = member_register.RealName             //真实姓名
	member.IsLockedLevel = 2                               //是否锁定会员层级
	member.Status = 1                                      //状态
	member.RegisterIp = ctx.RealIP()                       //注册ip
	member.IsEditPassword = 1                              //是否可以修改密码
	member.CreateTime = time.Now().Unix()                  //创建时间
	member.IsAgreeDeal = member_register.IsAcceptAgreement //是否同意本平台协议
	member.RegisterClientType = 2                          //注册平台
	member.WapStatus = 2
	member.PcStatus = 2
	member.IosStatus = 2
	member.AndroidStatus = 2
	//默认分层取得
	info, _, err := memberLevelBean.SiteDefault(member_register.Site, member_register.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	member.LevelId = info.LevelId
	var isDiscount int
	//会员注册优惠计算
	if regis.IsIp == 1 { //开启ip限制,那么同一个站点下面的同一个ip只能获取一次优惠，同ip再次注册无法获取优惠
		ip := ctx.RealIP()
		flag, err := memberBean.GetRecordByIp(member_register.Site, member_register.SiteIndexId, ip)
		if err != nil {
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
	err = agencyDef(member, member_register, ctx, regis)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	count, err := memberBean.WapMemberRegister(member, member_register, regis, isDiscount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(60118, ctx))
	}
	return ctx.NoContent(204)
}

//点击注册的时候根据域名获取会员注册信息和站点
func (*WapLoginController) GetMemberRegister(ctx echo.Context) error {
	memberRegisterSet := new(input.MemberRegisterSet)
	code := global.ValidRequestMember(memberRegisterSet, ctx)
	if code != 0 {
		return echo.ErrNotFound
	}
	//根据站点去取的会员注册设定
	backData, flag, err := memberRegisterSetup.GetOneSet(memberRegisterSet.SiteId, memberRegisterSet.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60058, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(backData))
}

//所属代理商判定
func agencyDef(member *schema.Member, register *input.WapMemberRegister, ctx echo.Context, regis schema.SiteMemberRegisterSet) error {
	if register.Introducer != 0 { //代理链接直接注册
		info, flag, err := memberBean.GetOneMemberByThird(register.Introducer) //取得三级经销商名下一个会员，根据该会员经销商获取信息
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if flag { //有属于这个代理的会员
			member.ThirdAgencyId = register.Introducer  //代理商
			member.SecondAgencyId = info.SecondAgencyId //二级
			member.FirstAgencyId = info.FirstAgencyId   //一级
		} else { //没有属于这个代理的会员,从代理表查找
			info, flag, err := agnecyBean.GetAgency(register.Introducer)
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
				member.ThirdAgencyId = register.Introducer //三级
				member.SecondAgencyId = info.ParentId      //二级经销商
				agency, _, err := agnecyBean.GetAgency(info.ParentId)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					return err
				}
				member.FirstAgencyId = agency.ParentId //一级
			}
		}
	} else if register.IntroducerMember != 0 { //会员推广链接注册
		info, flag, err := memberBean.GetInfoById(register.IntroducerMember)
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
			member.SpreadId = register.IntroducerMember //推广人id
			member.SpreadMoney = 0                      //推广获利
			member.ThirdAgencyId = info.ThirdAgencyId   //代理商Id
			member.SecondAgencyId = info.SecondAgencyId //二级代理商id
			member.FirstAgencyId = info.FirstAgencyId   //一级代理商Id
		}
	} else if register.IntroducerMember == 0 && register.Introducer == 0 { //自己直接过来注册，归属于站点默认的层级
		err := defaultAgency(ctx, regis, member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
	}
	return nil
}

//默认站点线处理
func defaultAgency(ctx echo.Context, regis schema.SiteMemberRegisterSet, member *schema.Member) error {
	thirdAgency, _, err := agnecyBean.GetDefault(regis.SiteId, regis.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	member.ThirdAgencyId = thirdAgency.Id
	member.SecondAgencyId = thirdAgency.ParentId //二级
	agency, _, err := agnecyBean.GetAgency(thirdAgency.ParentId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	member.FirstAgencyId = agency.ParentId //一级
	return nil
}
