package account

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//站点会员注册设定
type MemberRegController struct {
	controllers.BaseController
}

//获取会员注册设定
func (mrc *MemberRegController) RegisterSet(ctx echo.Context) error {
	memberRegSettingGet := new(input.MemberRegisterSettingGet)
	code := global.ValidRequest(memberRegSettingGet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, have, err := memberRegisterSettingBean.Get(memberRegSettingGet)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改会员注册设定
func (mrc *MemberRegController) RegisterSetEdit(ctx echo.Context) error {
	//TODO:金额验证未完成
	memberRegSetting := new(input.MemberRegisterSetting)
	code := global.ValidRequest(memberRegSetting, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断该站点是否进行过会员注册设定
	validInfo := new(input.MemberRegisterSettingGet)
	validInfo.SiteIndexId = memberRegSetting.SiteIndexId
	validInfo.SiteId = memberRegSetting.SiteId
	_, have, err := memberRegisterSettingBean.Get(validInfo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//如果没有,则进行添加操作,否则进行更新操作.
	if !have {
		count, err := memberRegisterSettingBean.Add(memberRegSetting)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count != 1 {
			return ctx.JSON(200, global.ReplyError(10002, ctx))
		}
	} else {
		count, err := memberRegisterSettingBean.Update(memberRegSetting)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count != 1 {
			return ctx.JSON(200, global.ReplyError(50173, ctx))
		}
	}
	return ctx.NoContent(204)
}
