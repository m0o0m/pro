//[控制器] [平台]注册模板管理
package templet

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//注册模板管理
type LoginTempController struct {
	controllers.BaseController
}

//注册模板查询
func (c *LoginTempController) GetLoginTempList(ctx echo.Context) error {
	data, err := loginTempBean.GetLoginTempList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//注册模板详情
func (c *LoginTempController) GetLoginTempListDetail(ctx echo.Context) error {
	Reg := new(input.LoginRegListDetailIn)
	code := global.ValidRequest(Reg, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := loginTempBean.GetOneLoginRegOneDetail(Reg)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50153, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//添加注册模板
func (c *LoginTempController) PostLoginTempAdd(ctx echo.Context) error {
	add_data := new(input.LoginAdd)
	code := global.ValidRequest(add_data, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := loginTempBean.AddLogin(add_data)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90101, ctx))
	}
	return ctx.NoContent(204)
}

//修改注册模板
func (c *LoginTempController) PutLoginTempUpdate(ctx echo.Context) error {
	updateData := new(input.LoginUpdate)
	code := global.ValidRequest(updateData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	Reg := new(input.LoginRegListDetailIn)
	Reg.Id = updateData.Id
	_, has, err := loginTempBean.GetOneLoginRegOneDetail(Reg)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50153, ctx))
	}
	count, err := loginTempBean.UpdateLogin(updateData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30287, ctx))
	}
	return ctx.NoContent(204)
}

//状态修改
func (c *LoginTempController) PutLoginTempStatusUpdate(ctx echo.Context) error {
	updateData := new(input.LoginStatus)
	code := global.ValidRequest(updateData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	Reg := new(input.LoginRegListDetailIn)
	Reg.Id = updateData.Id
	_, has, err := loginTempBean.GetOneLoginRegOneDetail(Reg)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50153, ctx))
	}
	count, err := loginTempBean.UpdateStatus(updateData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30288, ctx))
	}
	return ctx.NoContent(204)
}
