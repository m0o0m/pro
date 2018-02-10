//[控制器] [平台]银行列表管理
package site

import (
	"controllers"
	"encoding/json"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"strconv"
)

//银行列表管理
type BankCardController struct {
	controllers.BaseController
}

//银行列表查询
func (c *BankCardController) GetBankCardList(ctx echo.Context) error {
	//获取用户参数
	bc := new(input.BankCardList)
	code := global.ValidRequestAdmin(bc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取分页数据
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := bankCardBean.BankCardList(bc, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//银行卡添加
func (c *BankCardController) PostBankCardAdd(ctx echo.Context) error {
	//获取用户参数
	bc := new(input.BankCardAdd)
	code := global.ValidRequestAdmin(bc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//银行官网格式校验
	flag := global.DomainCheck(bc.BankWebsiteUrl)
	if !flag {
		return ctx.JSON(200, global.ReplyError(50127, ctx))
	}
	_, has, err := bankCardBean.BeOneBankCardByName(bc)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(50128, ctx))
	}
	count, err := bankCardBean.BankCardAdd(bc)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//银行卡修改
func (c *BankCardController) PutBankCardUpdate(ctx echo.Context) error {
	//获取用户参数
	bc := new(input.BankCardChange)
	code := global.ValidRequestAdmin(bc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//银行官网格式校验
	flag := global.DomainCheck(bc.BankWebsiteUrl)
	if !flag {
		return ctx.JSON(200, global.ReplyError(50127, ctx))
	}
	//查询银行是否存在
	has, err := bankCardBean.BeBankCard(bc.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50133, ctx))
	}
	count, err := bankCardBean.BankCardChange(bc)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//银行状态修改
func (c *BankCardController) PutBankCardStatusUpdate(ctx echo.Context) error {
	//获取用户参数
	bc := new(input.BankCardStatus)
	code := global.ValidRequestAdmin(bc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询银行是否存在
	has, err := bankCardBean.BeBankCard(bc.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50133, ctx))
	}
	count, err := bankCardBean.BankCardStatus(bc)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//银行卡删除
func (c *BankCardController) PutBankCardDel(ctx echo.Context) error {
	//获取用户参数
	bc := new(input.BankCardDelete)
	code := global.ValidRequestAdmin(bc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询银行是否存在
	has, err := bankCardBean.BeBankCard(bc.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50133, ctx))
	}
	count, err := bankCardBean.BankCardDelete(bc)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//同步银行列表同步到缓存redis
func (c *BankCardController) PostBankCardRedis(ctx echo.Context) error {
	//获取用户参数
	bc := new(input.BankCardList)
	code := global.ValidRequestAdmin(bc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := bankCardBean.BankCardRedis()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	sdata := make(map[string]interface{})
	for _, v := range data {
		num := strconv.FormatInt(v.Id, 10)
		sdata[num], _ = json.Marshal(v)
	}
	err = global.GetRedis().HMSet("bank_type", sdata).Err()
	return ctx.NoContent(204)
}
