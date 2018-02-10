//[控制器] [平台]代理管理
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//代理管理
type AgencyController struct {
	controllers.BaseController
}

//站点代理查询
func (c *AgencyController) GetAgencyList(ctx echo.Context) error {
	//获取用户参数
	ag := new(input.AgencyManageList)
	code := global.ValidRequestAdmin(ag, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := agencyBean.AgencyManageList(ag, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))

}

//代理占成
func (c *AgencyController) AgencyOccupationRatio(ctx echo.Context) error {
	//获取用户参数
	ag := new(input.AgencyOccupationRatio)
	code := global.ValidRequestAdmin(ag, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := agencyBean.OccupationRatio(ag)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
