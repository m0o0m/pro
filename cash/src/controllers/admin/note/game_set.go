//[控制器] [平台]电子管理
package note

import (
	"controllers"
	"encoding/json"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"strconv"
)

//电子管理
type GameSetController struct {
	controllers.BaseController
}

//查询电子游戏列表
func (c *GameSetController) GetVdGameList(ctx echo.Context) error {
	vdGameList := new(input.VdGameList)
	code := global.ValidRequestAdmin(vdGameList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)
	//查询电子游戏列表
	list, count, err := noteGameBean.GetVdGameList(vdGameList, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//添加电子游戏
func (c *GameSetController) PostVdGameAdd(ctx echo.Context) error {
	vdGameAdd := new(input.VdGameAdd)
	code := global.ValidRequestAdmin(vdGameAdd, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	code, count, err := noteGameBean.PostVdGameAdd(vdGameAdd)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(60312, ctx))
	}
	return ctx.NoContent(204)
}

//游戏视讯类型列表
func (c *GameSetController) GetVdGameTypeList(ctx echo.Context) error {
	list, err := gameSetBean.GameTypeList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//添加游戏视讯类型
func (c *GameSetController) PostVdGameTypeAdd(ctx echo.Context) error {
	game := new(input.VdGameType)
	code := global.ValidRequestAdmin(game, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := gameSetBean.GameTypeAdd(game)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10146, ctx))
	}
	return ctx.NoContent(204)
}

//修改游戏视讯类型
func (c *GameSetController) PutVdGameTypeUpdate(ctx echo.Context) error {
	game := new(input.VdGameTypeUpdate)
	code := global.ValidRequestAdmin(game, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := gameSetBean.GameTypeUpdate(game)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10144, ctx))
	}
	return ctx.NoContent(204)
}

//删除游戏视讯类型
func (c *GameSetController) DelVdGameType(ctx echo.Context) error {
	game := new(input.VdGameType)
	code := global.ValidRequestAdmin(game, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := gameSetBean.DelGameType(game)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10145, ctx))
	}
	return ctx.NoContent(204)
}

//电子游戏修改
func (c *GameSetController) PutVdGameUpdate(ctx echo.Context) error {
	vdGameUpdate := new(input.VdGameUpdate)
	code := global.ValidRequestAdmin(vdGameUpdate, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	if vdGameUpdate.Status == 0 || vdGameUpdate.Status > 3 {
		return ctx.JSON(200, global.ReplyError(60307, ctx))
	}
	code, count, err := noteGameBean.PutVdGameUpdate(vdGameUpdate)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(60311, ctx))
	}
	return ctx.NoContent(204)
}

//电子游戏修改（修改内容）
func (c *GameSetController) PutVdGameContentUpdate(ctx echo.Context) error {
	vdGameContentUpdate := new(input.VdGameContentUpdate)
	code := global.ValidRequestAdmin(vdGameContentUpdate, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	code, count, err := noteGameBean.PutVdGameContentUpdate(vdGameContentUpdate)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(60311, ctx))
	}
	return nil
}

//排序推荐度修改
func (c *GameSetController) PutVdGameOrderUpdate(ctx echo.Context) error {
	vdGameOrderUpdate := new(input.VdGameUpdate)
	code := global.ValidRequestAdmin(vdGameOrderUpdate, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	code, count, err := noteGameBean.PutVdGameOrderUpdate(vdGameOrderUpdate)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(60311, ctx))
	}
	return ctx.NoContent(204)
}

//状态修改（剔除）
func (c *GameSetController) PutVdGameStatusUpdate(ctx echo.Context) error {
	vdGameStatusUpdate := new(input.VdGameStatusUpdate)
	code := global.ValidRequestAdmin(vdGameStatusUpdate, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	code, count, err := noteGameBean.PutVdGameStatusUpdate(vdGameStatusUpdate)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30096, ctx))
	}
	return ctx.NoContent(204)
}

//正式试玩线路的开关
func (c *GameSetController) PutVdGameDemoSwitch(ctx echo.Context) error {
	vdGameDemoSwitch := new(input.VdGameUpdate)
	code := global.ValidRequestAdmin(vdGameDemoSwitch, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if vdGameDemoSwitch.IsSw == 0 && vdGameDemoSwitch.IsZs == 0 {
		return ctx.JSON(200, global.ReplyError(60309, ctx))
	}
	code, count, err := noteGameBean.PutVdGameOrderUpdate(vdGameDemoSwitch)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(60311, ctx))
	}
	return ctx.NoContent(204)
}

//清楚缓存
func (c *GameSetController) DelVdGameRedis(ctx echo.Context) error {
	data, err := noteGameBean.GameListRedis()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	sdata := make(map[string]interface{})
	for _, v := range data {
		num := strconv.FormatInt(v.Id, 10)
		sdata[num], _ = json.Marshal(v)
	}
	err = global.GetRedis().HMSet("game_list", sdata).Err()
	return ctx.NoContent(204)
}
