//[控制器] [平台]异常会员管理
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//异常会员管理
type AbnormalMemberController struct {
	controllers.BaseController
}

//异常会员查询
func (c *AbnormalMemberController) GetAbnormalMemberList(ctx echo.Context) error {
	abnormalMember := new(input.AbnormalMemberList)
	code := global.ValidRequestAdmin(abnormalMember, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	c.GetParam(listParam, ctx)
	list, count, err := abnormalMemberBean.GetAbnormalMemberList(abnormalMember, listParam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, list, int64(len(list)), count, ctx))
}

//异常会员处理
func (c *AbnormalMemberController) PutAbnormalMemberSet(ctx echo.Context) error {
	abnormalMember := new(input.AbnormalMemberSet)
	code := global.ValidRequest(abnormalMember, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := abnormalMemberBean.AbnormalMemberSet(abnormalMember)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10140, ctx))
	}
	return ctx.NoContent(204)
}
