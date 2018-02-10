package rebate

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

type RebateSetController struct {
	controllers.BaseController
}

//添加或者更新会员返佣优惠设定
func (c *RebateSetController) AddOrUpdate(ctx echo.Context) error {
	requestData := &input.MemberRebateSetAdd{}
	code := global.ValidRequest(requestData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var count int64
	var err error
	if requestData.Id > 0 {
		count, err = rebateSetBean.Update(requestData)
		if count != 1 {
			return ctx.JSON(200, global.ReplyError(70005, ctx))
		}
	} else {
		count, err = rebateSetBean.Add(requestData)
		if count != 1 {
			return ctx.JSON(200, global.ReplyError(70004, ctx))
		}
	}
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//删除会员优惠设定
func (c *RebateSetController) Del(ctx echo.Context) error {
	requestData := &input.MemberRebateSetDel{}
	code := global.ValidRequest(requestData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := rebateSetBean.Del(requestData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(70006, ctx))
	}
	return ctx.NoContent(204)
}

//查询优惠设定详情
func (c *RebateSetController) GetOne(ctx echo.Context) error {
	requestData := &input.MemberRebateSetDel{}
	code := global.ValidRequest(requestData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	result, b, err := rebateSetBean.Select(requestData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(6000, ctx))
	}
	if !b {
		return ctx.JSON(200, global.ReplyError(70007, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(result))
}

//查询优惠设定列表
func (c *RebateSetController) GetAll(ctx echo.Context) error {
	reqData := new(input.MemberRebateSetList)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询商品表信息
	products, err := productBean.GetList(reqData.SiteId, reqData.SiteIndexId, &input.ProductList{Status: 1})
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(products) == 0 {
		return ctx.JSON(500, global.ReplyError(70013, ctx))
	}
	//查询设置列表
	rebateSets, err := rebateSetBean.GetAll(reqData.SiteId, reqData.SiteIndexId, products)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(rebateSets, int64(len(rebateSets))))
}
