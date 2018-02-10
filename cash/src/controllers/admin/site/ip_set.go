//[控制器] [平台]ip开关控制
package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"strings"
)

//ip开关管理
type IpSetController struct {
	controllers.BaseController
}

//ip控制列表查询
func (c *IpSetController) GetIpSetList(ctx echo.Context) error {
	//获取用户参数
	ip_set := new(input.IpSetList)
	code := global.ValidRequestAdmin(ip_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取分页数据
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := ipSetBean.IpSetList(ip_set, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//ip区间添加
func (c *IpSetController) PostIpSetAdd(ctx echo.Context) error {
	//获取用户参数
	ip_set := new(input.IpSetAdd)
	code := global.ValidRequestAdmin(ip_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询添加的ip起始是否在同一ip段
	start := strings.Split(ip_set.IpStart, ".")
	end := strings.Split(ip_set.IpEnd, ".")
	if start[0] != end[0] || start[1] != end[1] || start[2] != end[2] {
		return ctx.JSON(200, global.ReplyError(50190, ctx))
	}
	if start[3] > end[3] {
		return ctx.JSON(200, global.ReplyError(50191, ctx))
	}
	//查询ip开关是否存在
	has, err := ipSetBean.BeOneIpSet(ip_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(50131, ctx))
	}

	count, err := ipSetBean.IpSetAdd(ip_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//Ip区间段修改
func (c *IpSetController) PutIpSetUpdate(ctx echo.Context) error {
	//获取用户参数
	ip_set := new(input.IpSetChange)
	code := global.ValidRequestAdmin(ip_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询添加的ip起始是否在同一ip段
	start := strings.Split(ip_set.IpStart, ".")
	end := strings.Split(ip_set.IpEnd, ".")
	if start[0] != end[0] || start[1] != end[1] || start[2] != end[2] {
		return ctx.JSON(200, global.ReplyError(50190, ctx))
	}
	if start[3] > end[3] {
		return ctx.JSON(200, global.ReplyError(50191, ctx))
	}
	//判断是否存在
	has, err := ipSetBean.BeIpSet(ip_set.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50134, ctx))
	}
	//修改
	count, err := ipSetBean.IpSetChange(ip_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//Ip白名单列表查询
func (c *IpSetController) GetIpWhiteList(ctx echo.Context) error {
	//获取用户参数
	ip_set := new(input.WhiteList)
	code := global.ValidRequestAdmin(ip_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取分页数据
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := ipSetBean.WhiteList(ip_set, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//Ip白名单添加
func (c *IpSetController) PostIpwhiteAdd(ctx echo.Context) error {
	//获取用户参数
	ip_set := new(input.WhiteListAdd)
	code := global.ValidRequestAdmin(ip_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询白名单中是否存在该站点的信息
	ip_set.Id = 0
	data, has, err := ipSetBean.BeWhiteList(ip_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//如果不存在，直接添加
	if !has {
		_, err = ipSetBean.WhiteListAdd(ip_set)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.NoContent(204)
	}
	//将字符串转换成数组
	ips := strings.Split(data.Ip, ",")
	//如果取出的ip中存在该ip
	for _, v := range ips {
		if ip_set.Ip == v {
			return ctx.JSON(200, global.ReplyError(50136, ctx))
		}
	}
	//如果取出的ip中不存在该ip
	ips = append(ips, ip_set.Ip)
	//将数组转换成用逗号分隔的字符串
	ip_set.Ip = strings.Join(ips, ",")
	//修改
	ip_change := new(input.WhiteListChange)
	ip_change.SiteId = ip_set.SiteId
	ip_change.Ip = ip_set.Ip
	ip_change.Remark = ip_set.Remark
	ip_change.Id = data.Id
	ip_change.Status = ip_set.Status
	count, err := ipSetBean.WhiteListChange(ip_change)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//Ip白名单修改
func (c *IpSetController) PutIpwhiteUpdate(ctx echo.Context) error {
	//获取用户参数
	ip_set := new(input.WhiteListChange)
	code := global.ValidRequestAdmin(ip_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//new input.WhiteListAdd  从而直接直接查询
	ip_add := new(input.WhiteListAdd)
	ip_add.Id = ip_set.Id
	ip_add.SiteId = ip_set.SiteId
	//查询白名单中是否存在该站点的信息
	_, has, err := ipSetBean.BeWhiteList(ip_add)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//如果不存在，返回提示
	if !has {
		return ctx.JSON(200, global.ReplyError(50137, ctx))
	}
	//修改
	count, err := ipSetBean.WhiteListChange(ip_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}
