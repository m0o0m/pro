//[控制器] [平台]视讯账号管理
package customer

import (
	"global"

	"controllers"
	"github.com/labstack/echo"
	"models/input"
)

//视讯账号管理
type VideoMemberController struct {
	controllers.BaseController
}

//视讯账号查询
func (c *VideoMemberController) GetVideoMemberList(ctx echo.Context) error {
	siteVideoMember := new(input.SiteVideoMemberSearch)
	code := global.ValidRequestAdmin(siteVideoMember, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	c.GetParam(listParam, ctx)
	list, count, err := videoMemberBean.SiteVideoMemberSearch(siteVideoMember, listParam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, list, int64(len(list)), count, ctx))
}

//视讯类型下拉框
func (v *VideoMemberController) GetVideoTypeList(ctx echo.Context) error {
	infolist, err := videoMemberBean.GetVideoList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(infolist))
}
