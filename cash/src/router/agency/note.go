package agency

import (
	"controllers/agency/note"
	//"controllers/agency/rebate"
	"github.com/labstack/echo"
	//"router"
	//"router"
	"router"
)

func NoteRouter(c *echo.Echo) {
	e := c.Group(apiPath, router.GetRedisToken, router.AgencyAccessLog) //开启代理操作日志
	bet_record_info := new(note.BetRecordInfoController)
	e.GET("/betrecord_list", bet_record_info.GetBeRecordList) //注单列表
}
