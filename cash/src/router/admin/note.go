//[平台] 注单模块控制管理（合并彩票体育电子视讯）
package admin

import (
	"controllers/admin/note"
	"github.com/labstack/echo"
	"router"
)

//注单管理
func NoteRouter(c *echo.Echo) {
	e := c.Group(apiPath, router.AdminRedisCheck, router.AdminAccessLog) //开启操作日志记录

	//注单管理
	noteBet := new(note.NoteBetController)
	e.GET("/note/get_bet_list", noteBet.GetNoteBetList) //额度统计

	//电子游戏类型管理
	gameSet := new(note.GameSetController)
	e.GET("/note/vdgameType/list", gameSet.GetVdGameTypeList)        //游戏类型列表
	e.POST("/note/vdgame_type/add", gameSet.PostVdGameTypeAdd)       //添加游戏类型
	e.PUT("/note/vdgame_type/update", gameSet.PutVdGameTypeUpdate)   //修改游戏类型
	e.DELETE("/note/vdgame_type/del", gameSet.DelVdGameType)         //删除游戏类型
	e.GET("/note/get_game_list", gameSet.GetVdGameList)              //查询电子游戏列表
	e.POST("/note/game_add", gameSet.PostVdGameAdd)                  //添加电子游戏
	e.PUT("/note/game_edit", gameSet.PutVdGameUpdate)                //电子游戏修改
	e.PUT("/note/game_content_edit", gameSet.PutVdGameContentUpdate) //电子游戏修改(内容)
	e.PUT("/note/game_order_edit", gameSet.PutVdGameOrderUpdate)     //排序推荐度修改
	e.PUT("/note/game_status_edit", gameSet.PutVdGameStatusUpdate)   //状态修改（剔除）
	e.PUT("/note/game_demo_switch", gameSet.PutVdGameDemoSwitch)     //正式试玩线路的开关
	e.GET("/note/game_redis_del", gameSet.DelVdGameRedis)            //清楚同步redis缓存
}
