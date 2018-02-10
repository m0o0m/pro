package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/function"
	"models/input"
	"models/schema"
	"time"
)

// 维护管理
type MaintainController struct {
	controllers.BaseController
}

// 主页
func (c *MaintainController) Index(ctx echo.Context) error {
	result := back.MaintainIndexRes{}

	c.SetMaintain()
	sites, _ := new(function.SiteOperateBean).SiteSiteIndexIdBy()
	result.Sites = sites

	data := c.GetMaintainGroupByCType()
	result.Data = data

	list := c.GetMaintainGroupByLineSite()
	result.List = list
	return ctx.JSON(200, global.ReplyItem(result))
}

// 添加/修改
func (c *MaintainController) Create(ctx echo.Context) error {
	code := c.SetMaintain()

	var result []string
	if code == 1 {
		c.ResetRedis()
	} else if code == 0 {

	}
	return ctx.JSON(200, global.ReplyItem(result))
}

// 关闭
func (c *MaintainController) Close(ctx echo.Context) error {
	code := c.DelMaintain()

	var result []string
	if code == 1 {
		c.ResetRedis()
	} else if code == 0 {

	}
	return ctx.JSON(200, global.ReplyItem(result))
}

// 重置Redis
func (c *MaintainController) ResetRedis() {
	// del code...

	c.SetRedis()
}

// 写入Redis
func (c *MaintainController) SetRedis() {
	// set code...
}

// 添加/修改 处理
func (c *MaintainController) SetMaintain() (code int64) {
	request := new(input.MaintainData)

	// debug start
	request.MType = 3
	request.CType = "1"
	request.LindId = "aaa"
	request.SiteId = "a"
	request.ProductId = "1"
	request.Remark = "Test"
	// debug end

	item := schema.MaintainData{}
	item.MType = request.MType
	item.CType = request.CType
	item.LindId = request.LindId
	item.SiteId = request.SiteId
	item.ProductId = request.ProductId
	item.StartTime = request.StartTime
	item.EndTime = request.EndTime
	item.Remark = request.Remark

	has, err := new(function.Maintain).MaintainHas(item)
	if err != nil {
		global.GlobalLogger.Error("err:%v", err)
		return 0
	}
	if !has {
		item.AddTime = time.Now().Unix()
		count, err := new(function.Maintain).InsertMaintain(item)
		if err != nil || count != 1 {
			global.GlobalLogger.Error("err:%v", err)
			return 0
		}
	} else {
		item.UpdateTime = time.Now().Unix()
		count, err := new(function.Maintain).UpdateMaintain(item)
		if err != nil || count != 1 {
			global.GlobalLogger.Error("err:%v", err)
			return 0
		}
	}

	return 1
}

// 关闭 处理
func (c *MaintainController) DelMaintain() (code int64) {
	request := new(input.MaintainData)

	// debug start
	request.MType = 3
	request.CType = "1"
	request.LindId = "aaa"
	request.SiteId = "a"
	// debug end

	item := schema.MaintainData{}
	item.MType = request.MType
	item.CType = request.CType
	item.LindId = request.LindId
	item.SiteId = request.SiteId

	count, err := new(function.Maintain).DelMaintain(item)
	if err != nil || count != 1 {
		global.GlobalLogger.Error("err:%v", err)
		return 0
	}

	return 1
}

// 获取维护信息 并组装成需要的格式
func (c *MaintainController) GetMaintainGroupByCType() (data map[int]*map[string]*map[int]*back.MaintainData) {
	maintainData, _ := new(function.Maintain).GetMaintainData()

	data = make(map[int]*map[string]*map[int]*back.MaintainData)
	for k, m := range maintainData {
		oneLevel, ok := data[m.MType]
		if !ok {
			oneMap := make(map[string]*map[int]*back.MaintainData)
			oneLevel = &oneMap
			data[m.MType] = &oneMap
		}
		twoLevel, ok := (*oneLevel)[m.CType]
		if !ok {
			twoMap := make(map[int]*back.MaintainData)
			twoLevel = &twoMap
			(*oneLevel)[m.CType] = &twoMap
		}
		(*twoLevel)[k] = &maintainData[k]
	}

	return
}
func (c *MaintainController) GetMaintainGroupByLineSite() (data map[int]*map[string]*map[string]*map[int]*back.MaintainData) {
	maintainData, _ := new(function.Maintain).GetMaintainData()

	data = make(map[int]*map[string]*map[string]*map[int]*back.MaintainData)
	for k, m := range maintainData {
		oneLevel, ok := data[m.MType]
		if !ok {
			oneMap := make(map[string]*map[string]*map[int]*back.MaintainData)
			oneLevel = &oneMap
			data[m.MType] = &oneMap
		}
		twoLevel, ok := (*oneLevel)[m.LindId]
		if !ok {
			twoMap := make(map[string]*map[int]*back.MaintainData)
			twoLevel = &twoMap
			(*oneLevel)[m.LindId] = &twoMap
		}
		thrLevel, ok := (*twoLevel)[m.SiteId]
		if !ok {
			thrMap := make(map[int]*back.MaintainData)
			thrLevel = &thrMap
			(*twoLevel)[m.SiteId] = &thrMap
		}
		(*thrLevel)[k] = &maintainData[k]
	}

	return
}
