package member

import (
	"controllers"
	"framework/validation"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"strings"
)

type AgencyRegisterController struct {
	controllers.BaseController
}

//获取代理注册设定
func (arc *AgencyRegisterController) Set(ctx echo.Context) error {
	//获取站点信息
	agencyRegisterSet, flag, err := GetSiteInfo(ctx)
	if err != nil {
		global.GlobalLogger.Error("error%s", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	list, has, err := secondDistributionRegisterSetupBeen.SiteIdExist(agencyRegisterSet.SiteId, agencyRegisterSet.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	if list.RegisterProxy != 1 {
		return ctx.JSON(500, global.ReplyError(30167, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//提交代理注册申请
func (arc *AgencyRegisterController) Register(ctx echo.Context) error {
	agencyRegister := new(input.AgencyRegister)
	//请求参数绑定
	if err := ctx.Bind(agencyRegister); err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(10000, ctx))
	}
	//请求参数验证
	valid := validation.Validation{}
	ok, err := valid.Valid(agencyRegister)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !ok {
		for _, e := range valid.Errors {
			return ctx.JSON(200, global.ReplyError(e.Code(), ctx))
		}
	}
	//获取站点信息
	site, flag, err := GetSiteInfo(ctx)
	if err != nil {
		global.GlobalLogger.Error("error%s", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	agencyRegister.SiteId = site.SiteId
	agencyRegister.SiteIndexId = site.SiteIndexId
	list, _, err := secondDistributionRegisterSetupBeen.SiteIdExist(agencyRegister.SiteId, agencyRegister.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if list.RegisterProxy != 1 { //不启用代理注册
		return ctx.JSON(200, global.ReplyError(30167, ctx))
	}
	if list.ChineseNickname == 1 && list.IsMustChineseNickname == 1 { //启用中文昵称且必填
		if agencyRegister.ChineseNickname == "" {
			return ctx.JSON(200, global.ReplyError(30171, ctx))
		}
	} else if list.ChineseNickname == 2 { //禁用中文昵称
		agencyRegister.ChineseNickname = ""
	}
	if list.EnglishNickname == 1 && list.IsMustEnglishNickname == 1 { //启用英文昵称且必填
		if agencyRegister.EnglishNickname == "" {
			return ctx.JSON(200, global.ReplyError(30172, ctx))
		}
	} else if list.EnglishNickname == 2 { //禁用英文昵称
		agencyRegister.EnglishNickname = ""
	}
	if list.NeedCard == 1 && list.IsMustIdentity == 1 { //启用证件号且必填
		if agencyRegister.Card == "" {
			return ctx.JSON(200, global.ReplyError(30173, ctx))
		}
	} else if list.NeedCard == 2 { //禁用证件号
		agencyRegister.Card = ""
	}
	if list.NeedEmail == 1 && list.IsMustEmail == 1 { //启用邮箱且必填
		if agencyRegister.Email == "" {
			return ctx.JSON(200, global.ReplyError(30174, ctx))
		}
	} else if list.NeedEmail == 2 { //禁用邮箱
		agencyRegister.Email = ""
	}
	if list.NeedQq == 1 && list.IsMustQq == 1 { //启用QQ且必填
		if agencyRegister.Qq == "" {
			return ctx.JSON(200, global.ReplyError(30175, ctx))
		}
	} else if list.NeedQq == 2 { //禁用qq
		agencyRegister.Qq = ""
	}
	if list.NeedPhone == 1 && list.IsMustPhone == 1 { //启用手机号且必填
		if agencyRegister.Phone == "" {
			return ctx.JSON(200, global.ReplyError(30176, ctx))
		}
	} else if list.NeedPhone == 2 { //禁用手机号
		agencyRegister.Phone = ""
	}
	if list.PromoteWebsite == 1 && list.IsMustPromoteWebsite == 1 { //启用推广网址且必填
		if agencyRegister.PromoteWebsite == "" {
			return ctx.JSON(200, global.ReplyError(30177, ctx))
		}
	} else if list.PromoteWebsite == 2 { //禁用推广网址
		agencyRegister.PromoteWebsite = ""
	}
	if list.OtherMethod == 1 && list.IsMustMethod == 1 { //启用其他方式且必填
		if agencyRegister.OtherMethod == "" {
			return ctx.JSON(200, global.ReplyError(30178, ctx))
		}
	} else if list.OtherMethod == 2 { //禁用其他方法
		agencyRegister.OtherMethod = ""
	}
	//代理账号是否已存在（查看代理表和代理注册表）
	has, have, err := distributionApplyBeen.IsExistAccount(agencyRegister.AgencyAccount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has || have {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//两次密码是否一致
	if agencyRegister.Password != agencyRegister.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30012, ctx))
	}
	//验证验证码
	codes := ctx.Request().Header.Get("code")
	key, err := getMemberRedis(codes)
	if err == redis.Nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if key == "" || strings.ToLower(key) != strings.ToLower(agencyRegister.Code) {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	//删除验证码
	err = global.GetRedis().Del(codes).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	count, err := distributionApplyBeen.Add(agencyRegister)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30179, ctx))
	}
	return ctx.JSON(200, global.ReplyError(204, ctx))
}
