//[控制器] [平台]催款管理
package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/schema"
	"time"
)

//催款管理
type PressMoneyController struct {
	controllers.BaseController
}

//催款账单查询
func (c *PressMoneyController) GetPressMoneyList(ctx echo.Context) error {
	pressmoney := new(input.PressMoneyList)
	code := global.ValidRequest(pressmoney, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := pressMoneyBean.GetPressMoneyList(pressmoney)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//催款账单添加
func (c *PressMoneyController) PostPressMoneyAdd(ctx echo.Context) error {
	pressmoney := new(input.PressMoneyAdd)
	code := global.ValidRequest(pressmoney, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var data schema.SiteMoneyPress
	var data2 []schema.SiteMoneyPress
	for _, v := range pressmoney.SiteId {
		data.SiteId = v
		data.Bank = pressmoney.Bank
		data.PayAddress = pressmoney.PayAddress
		data.PayCard = pressmoney.PayCard
		data.PayName = pressmoney.PayName
		data.Remark = pressmoney.Remark
		data.AddDate = time.Now().Unix()
		data.Status = 1
		data.State = 1
		data.Qishu = 0 //期数，留作备用
		data.Money = 0 //应交金额，留作备用
		data2 = append(data2, data)
	}
	count, err := pressMoneyBean.Add(data2)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//获取单条催款账单详情
func (c *PressMoneyController) GetPressMoneyOne(ctx echo.Context) error {
	pressmoney := new(input.PressMoneyOne)
	code := global.ValidRequest(pressmoney, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := pressMoneyBean.GetInfo(pressmoney)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has == false {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//催款账单修改
func (c *PressMoneyController) PutPressMoneyUpdate(ctx echo.Context) error {
	pressmoney := new(input.PressMoneyEdit)
	code := global.ValidRequest(pressmoney, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := pressMoneyBean.UpdateInfo(pressmoney)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//催款账单确认
func (c *PressMoneyController) PutPressMoneyStatus(ctx echo.Context) error {
	pressmoney := new(input.PressMoneyStatus)
	code := global.ValidRequest(pressmoney, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if pressmoney.Status != 1 {
		return ctx.JSON(200, global.ReplyError(50154, ctx))
	}
	count, err := pressMoneyBean.UpdateStatus(pressmoney)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//siteId
func (c *PressMoneyController) PressMoneySiteDrop(ctx echo.Context) error {
	data, err := pressMoneyBean.PressMoneySiteByDrop()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
