//[控制器] [平台]现金系统管理
package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//现金系统管理
type CashRecordController struct {
	controllers.BaseController
}

//现金记录查询
func (c *CashRecordController) GetCashRecordList(ctx echo.Context) error {
	memberCashRecord := new(input.MemberCashRecord)
	code := global.ValidRequest(memberCashRecord, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)
	if len(memberCashRecord.SiteId) == 0 {
		return ctx.JSON(200, global.ReplyPagination(listParams, nil, 0, 0, ctx))

	}
	data, count, err := memberCashRecordBean.GetCashRecordList(memberCashRecord, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	return ctx.JSON(200, global.ReplyCollections("data", data, "count", count))
}

//批量取消现金报表(硬删除)
func (c *CashRecordController) DelCashRecord(ctx echo.Context) error {
	memberCashRecord := new(input.PutMemberCashRecord)
	code := global.ValidRequest(memberCashRecord, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//只能是超管才能操作
	admin := ctx.Get("admin").(*global.AdminRedisStruct)
	if admin.RoleId != 5 {
		return ctx.JSON(200, global.ReplyError(30239, ctx))
	}
	count, err := memberCashRecordBean.DelMemberCashRecord(memberCashRecord.Ids)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30262, ctx))
	}
	return ctx.NoContent(204)
}

//批量删除现金报表(软删除)
func (c *CashRecordController) PutCashRecord(ctx echo.Context) error {
	memberCashRecord := new(input.PutMemberCashRecord)
	code := global.ValidRequest(memberCashRecord, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//只能是超管才能操作
	admin := ctx.Get("admin").(*global.AdminRedisStruct)
	if admin.RoleId != 5 {
		return ctx.JSON(200, global.ReplyError(30239, ctx))
	}
	count, err := memberCashRecordBean.PutMemberCashRecord(memberCashRecord.Ids)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30262, ctx))
	}
	return ctx.NoContent(204)
}
