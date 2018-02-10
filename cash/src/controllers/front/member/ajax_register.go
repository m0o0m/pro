package member

import (
	"framework/render"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"global"
	"models/function"
	"models/input"
	"strings"
)

type AjaxRegister struct {
}

//注册ajax请求
func (c *AjaxRegister) RegisterAjax(ctx echo.Context) error {

	ajaxinfo := new(input.AjaxRegister)
	code := global.ValidRequest(ajaxinfo, ctx)
	if code != 0 {
		return render.PageErr(code, ctx)
	}
	domain := ctx.Request().Host
	damainurl := strings.Split(domain, ":")

	siteinfo, _, _ := siteDomainBean.GetSiteByDomain(damainurl[0])

	member_bean := new(function.MemberBean)
	switch ajaxinfo.Ajax {
	case "CheckUser": //判断账号是否重复
		flag, err := member_bean.CheckIsExist(ajaxinfo.Account, siteinfo.SiteId)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if flag {
			return ctx.JSON(200, global.ReplyError(60057, ctx))
		}

		return ctx.JSON(200, global.ReplyItem(flag))
	case "CheckCode":
		codes := ctx.Request().Header.Get("Code")
		key, err := getMemberRedis(codes)
		if err == redis.Nil {
			global.GlobalLogger.Error("err:s%", err.Error())
			return ctx.JSON(200, global.ReplyError(20021, ctx))
		}
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if key == "" || strings.ToLower(key) != strings.ToLower(ajaxinfo.Code) {
			return ctx.JSON(200, global.ReplyError(20021, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(1))
	case "CheckRealName":
		flag, err := member_bean.CheckRealNameExist(ajaxinfo.RealName)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if flag {
			return ctx.JSON(200, global.ReplyError(60071, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(flag))
	case "CheckAgentUser":
		agencybean := new(function.DistributionApplyBeen)
		has, have, err := agencybean.IsExistAccount(ajaxinfo.AgencyUser)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if has || have {
			return ctx.JSON(200, global.ReplyError(30039, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(1))
	}
	return ctx.JSON(500, global.ReplyError(60000, ctx))

}

//注册状态是否开启检测
func (c *AjaxRegister) GetIsRegStatus(ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	data, has, err := registerStatusBean.GetIsReg(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(200, global.ReplyError(60071, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10051, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//代理注册ajax请求

//获取登录的时候存储的redis值
func getMemberRedis(token string) (string, error) {
	key, err := global.GetRedis().Get(token).Result()
	return key, err
}
