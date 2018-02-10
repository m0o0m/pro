//[控制器] [平台]自助优惠申请
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//自助优惠
type AutoDisCountController struct {
	controllers.BaseController
}

//自助优惠申请列表查询
func (c *AutoDisCountController) GetAutoDiscountList(ctx echo.Context) error {
	discount := new(input.DiscountList)
	code := global.ValidRequestAdmin(discount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	infoList, count, err := selfHelpApplyforBean.SelfApplyList(discount, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, infoList, int64(len(infoList)), count, ctx))
}

//自助优惠申请列表修改
func (c *AutoDisCountController) PutAutoDiscountUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//自助优惠设定查询
func (c *AutoDisCountController) GetAutoDiscountSet(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//获取自助优惠详情
func (c *AutoDisCountController) GetInfoMation(ctx echo.Context) error {
	discount := new(input.AutoDiscountInfo)
	code := global.ValidRequestAdmin(discount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	infolist, _, err := selfHelpApplyforBean.GetAutoDiscountInfo(discount.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(infolist))
}

//自助优惠设定修改
func (c *AutoDisCountController) PutAutoDiscountSet(ctx echo.Context) error {
	discount := new(input.AutoDiscountStatus)
	code := global.ValidRequestAdmin(discount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询是否存在
	_, has, err := siteOperateBean.GetSingleSite(discount.SiteIndexId, discount.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	count, err := selfHelpApplyforBean.ChangeStatusDiscount(discount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30273, ctx))
	}
	return ctx.NoContent(204)
}

//自助优惠开关列表
func (c *AutoDisCountController) SelfDiscountSwitch(ctx echo.Context) error {
	discount := new(input.SelfDiscountSwitch)
	code := global.ValidRequestAdmin(discount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	infolist, err := siteOperateBean.GetSiteSelfSwitch(discount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(infolist))
}

//拒绝一条优惠申请[这个后台没有]
func (c *AutoDisCountController) RefuseOneDiscountApply(ctx echo.Context) error {
	refuse := new(input.RefuseApplyFor)
	code := global.ValidRequestAdmin(refuse, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作人
	userinfo := ctx.Get("admin").(global.AdminRedisStruct)
	//判断该条申请是否存在
	info, flag, err := selfHelpApplyforBean.GetAutoDiscountInfo(refuse.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if info.Id == 0 || !flag {
		return ctx.JSON(200, global.ReplyError(60115, ctx))
	}
	//拒绝
	count, err := selfHelpApplyforBean.RefusedFor(refuse, userinfo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}
