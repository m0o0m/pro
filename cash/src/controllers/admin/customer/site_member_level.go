//[控制器] [平台]会员层级管理
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//会员层级管理
type MemberLevelController struct {
	controllers.BaseController
}

//站点会员层级查询
func (c *MemberLevelController) GetMemeberLevelList(ctx echo.Context) error {
	levelIndex := new(input.LevelIndex)
	code := global.ValidRequestAdmin(levelIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)
	data, count, err := memberLevelBean.SiteLevelList(levelIndex, listParams)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}
