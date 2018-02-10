package account

import (
	"controllers"
	"global"
	"models/input"

	"github.com/labstack/echo"
)

//体系查询
type SearchController struct {
	controllers.BaseController
}

//体系查询列表
func (sc *SearchController) Index(ctx echo.Context) error {
	listparam := new(global.ListParams)
	//获取listparam的数据
	sc.GetParam(listparam, ctx)
	SearchState := new(input.Search)
	code := global.ValidRequest(SearchState, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if SearchState.Isvague == 0 {
		if SearchState.Account == "" {
			return ctx.JSON(200, global.ReplyPagination(listparam, nil, int64(0), 0, ctx))
		}
		data, count, err := agencyCountBean.SearchSystem(SearchState)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
	}
	if SearchState.Isvague == 1 {
		data, count, err := agencyCountBean.SearchSystemBlur(SearchState, listparam)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, nil, 0, 0, ctx))
}
