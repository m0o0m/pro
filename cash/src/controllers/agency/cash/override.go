package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
)

//代理退佣
type OverrideController struct {
	controllers.BaseController
}

//代理退佣(get列表)
func (pc *OverrideController) OverrideList(ctx echo.Context) error {
	combo := new(input.OverRideGet)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取listparam的数据
	list, count, err := overRideBean.OverRideList(combo)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10114, ctx))
	}
	rewater := new(back.Rewater)
	rebate := new(back.Rebate)
	list1 := new(back.List)
	data3 := new(back.Data)
	var data []float64
	var data1 []float64
	var data2 []string
	for _, v := range list {
		rewater.Percent1 = v.Percent1
		rebate.Percent = v.Percent
		data3.BetMoney = v.BetMoney
		data2 = append(data2, v.Name)
		list1.Id = v.Id
		list1.Amount = v.Amount
		list1.Member = v.Member
		data = append(data, rebate.Percent)
		data1 = append(data1, rewater.Percent1)
		list1.Rebate = data
		list1.Rewater = data1
	}
	data3.Arr = data2
	data3.List = list1
	return ctx.JSON(200, global.ReplyCollection(data3, count))
}

//获取单条代理退佣设定数据
func (pc *OverrideController) OverGetOne(ctx echo.Context) error {
	over := new(input.OverGetOne)
	code := global.ValidRequest(over, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取数据
	list, err := overRideBean.OverGetOne(over)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10114, ctx))
	}
	var data = new(back.List1)
	var data1 []float64
	var data2 []float64
	var data3 []string

	for _, v := range list {
		data.Member = v.Member
		data.Amount = v.Amount
		data1 = append(data1, v.Percent)
		data2 = append(data2, v.Percent1)
		data3 = append(data3, v.Name)
	}
	data.Rewater = data2
	data.Rebate = data1
	data.Name = data3
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改代理退佣设定

func (*OverrideController) UpdateOver(ctx echo.Context) error {
	list := new(input.OverRideUpdate)
	code := global.ValidRequest(list, ctx)

	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	count, err := overRideBean.UpdateOver(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//删除一条代理退佣设定
func (*OverrideController) DeleteOver(ctx echo.Context) error {
	over := new(input.OverRideDelet)
	code := global.ValidRequest(over, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := overRideBean.DeleteOver(over)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10116, ctx))
	}
	return ctx.NoContent(204)
}

//添加一条代理退佣设定
func (*OverrideController) OverRideAdd(ctx echo.Context) error {
	manualAccess := new(input.OverRideAdd)
	code := global.ValidRequest(manualAccess, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := overRideBean.OverRideAdd(manualAccess)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10117, ctx))
	}
	return ctx.NoContent(204)
}

//修改有效会员投注金额
func (*OverrideController) UpdataMoney(ctx echo.Context) error {
	list := new(input.UpdateMoney)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := overRideBean.UpdataMoney(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10118, ctx))
	}
	return ctx.NoContent(204)
}
