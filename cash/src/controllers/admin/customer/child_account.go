//[控制器] [平台]子账号管理
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"time"
)

//代理管理
type ChildAccountController struct {
	controllers.BaseController
}

//子账号查询
func (c *ChildAccountController) GetChildAccountList(ctx echo.Context) error {
	//获取用户参数
	ca := new(input.ChildAccountList)
	code := global.ValidRequestAdmin(ca, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if ca.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", ca.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if ca.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", ca.EndTime, loc)
		times.EndTime = et.Unix()
	}
	if ca.CType == 1 {
		ca.Account = ca.Value
	} else if ca.CType == 2 {
		ca.Name = ca.Value
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := childAccountBean.AcccountChildList(ca, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var total back.OnlineNumberAndTotal
	if len(data) > 0 {
		var i int64
		for _, v := range data {
			if v.IsLogin == 1 {
				i = i + 1
			}
		}
		total.OnlineNumber = i

	} else {
		total.OnlineNumber = 0
	}
	total.TotalNumber = count
	var list = make(map[string]interface{})
	list["data"] = data
	list["total"] = total
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(data)), count, ctx))
}

//获取一条子帐号
func (c *ChildAccountController) GetOneChildAccount(ctx echo.Context) error {
	//获取用户参数
	ca := new(input.OneChildAccount)
	code := global.ValidRequestAdmin(ca, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := childAccountBean.OneChildInfo(ca)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//子账号修改
func (c *ChildAccountController) PutChildAccountUpdate(ctx echo.Context) error {
	//获取用户参数
	ca := new(input.ChildAccountChange)
	code := global.ValidRequestAdmin(ca, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断子帐号是否存在
	has, err := childAccountBean.BeChildAccount(ca.Id, ca.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50132, ctx))
	}
	if ca.Password != ca.RePassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}
	//密码加密
	if ca.Password != "" {
		//密码不能少于6位
		if len(ca.Password) < 6 {
			return ctx.JSON(200, global.ReplyError(20005, ctx))
		}
		md5Password, err := global.MD5ByStr(ca.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		ca.Password = md5Password
	}
	count, err := childAccountBean.ChildInfoChange(ca)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30271, ctx))
	}
	return ctx.NoContent(204)
}

//子账号状态修改
func (c *ChildAccountController) PutChildAccountStatusUpdate(ctx echo.Context) error {
	//获取用户参数
	ca := new(input.ChildAccountStatus)
	code := global.ValidRequestAdmin(ca, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断子帐号是否存在
	has, err := childAccountBean.BeChildAccount(ca.Id, ca.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50132, ctx))
	}
	if ca.Status == 1 {
		//查询该id下面是否有下级
		data, err := childAccountBean.WhetherSubordinate(ca)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if len(data) > 0 {
			return ctx.JSON(200, global.ReplyError(50121, ctx))
		}
	}
	count, err := childAccountBean.ChildStatusChange(ca)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30272, ctx))
	}
	return ctx.NoContent(204)
}

//站点下拉框
func (c *ChildAccountController) SiteSiteIndexIdByAll(ctx echo.Context) error {
	data, err := childAccountBean.SiteSiteIndexIdBy()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//子站点下拉框
func (c *ChildAccountController) SiteIndexIdList(ctx echo.Context) error {
	siil := new(input.SiteIndexIdList)
	code := global.ValidRequestAdmin(siil, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := childAccountBean.SiteIndexIdList(siil.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
