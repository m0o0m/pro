package account

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//代理申请管理
type AgencyRegController struct {
	controllers.BaseController
}

//代理申请列表
func (arc *AgencyRegController) Index(ctx echo.Context) error {
	listparam := new(global.ListParams)
	//获取listparam的数据
	arc.GetParam(listparam, ctx)
	agencyIndex := new(input.AgencyIndex)
	code := global.ValidRequest(agencyIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	//如果当前登录者身份不是开户人
	if user.Level != 1 || user.RoleId != 1 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	//查看站点是否存在
	if agencyIndex.SiteIndexId != "" {
		has, err := distributionApplyBeen.SiteIdExists(agencyIndex.SiteId, agencyIndex.SiteIndexId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyPagination(listparam, nil, int64(0), 0, ctx))
		}
	}
	list, count, err := distributionApplyBeen.Get(agencyIndex, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//审核代理申请（审核通过就添加账号）
func (arc *AgencyRegController) Add(ctx echo.Context) error {
	agentRegEdit := new(input.AgentRegEdit)
	code := global.ValidRequest(agentRegEdit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	//如果当前登录者身份不是开户人
	if user.Level != 1 || user.RoleId != 1 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}
	//判断账号是否存在
	has, err := subAccountBeen.GetAccount(agentRegEdit.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//在代理申请表中获取数据
	data, has, err := subAccountBeen.GetAccountByReg(agentRegEdit.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30243, ctx))
	}
	//判断两次密码是否输入一致
	if agentRegEdit.Password != agentRegEdit.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30012, ctx))
	}
	if agentRegEdit.Password != "" {
		//给密码加密
		md5Password, err := global.MD5ByStr(agentRegEdit.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		agentRegEdit.Password = md5Password
	} else {
		agentRegEdit.Password = data.Password
	}
	//得到站点id
	agentRegEdit.SiteId = user.SiteId
	count, err := distributionApplyBeen.Update(agentRegEdit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30003, ctx))
	}
	return ctx.NoContent(204)
}

//删除代理申请
func (arc *AgencyRegController) AgentRegState(ctx echo.Context) error {
	agentRegState := new(input.AgentRegState)
	code := global.ValidRequest(agentRegState, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看agency_id是否为0，不是就不能删除
	ars, err := distributionApplyBeen.IsAgencyId(agentRegState.RegisterId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if ars.AgencyId != 0 {
		return ctx.JSON(200, global.ReplyError(30099, ctx))
	}
	count, err := distributionApplyBeen.Delete(agentRegState)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30004, ctx))
	}
	return ctx.NoContent(204)
}

//获取代理申请设定
func (arc *AgencyRegController) Set(ctx echo.Context) error {
	agentSet := new(input.AgentSet)
	code := global.ValidRequest(agentSet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	//如果当前登录者身份不是开户人
	if user.Level != 1 || user.RoleId != 1 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	list, has, err := secondDistributionRegisterSetupBeen.SiteIdExist(agentSet.SiteId, agentSet.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	} else if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//修改代理申请设定
func (arc *AgencyRegController) SetEdit(ctx echo.Context) error {
	agentSetDo := new(input.AgentSetDo)
	code := global.ValidRequest(agentSetDo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	//如果当前登录者身份不是开户人
	if user.Level != 1 || user.RoleId != 1 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	//查看注册设定表的站点id是否存在
	_, has, err := secondDistributionRegisterSetupBeen.
		SiteIdExist(agentSetDo.SiteId, agentSetDo.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	} else if !has {
		//不存在站点id就添加数据
		count, err := secondDistributionRegisterSetupBeen.Add(agentSetDo)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50174, ctx))
		}
	} else {
		//存在就修改数据
		count, err := secondDistributionRegisterSetupBeen.Update(agentSetDo)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50173, ctx))
		}
	}
	return ctx.NoContent(204)
}

//获取一条代理注册申请
func (arc *AgencyRegController) OneAgencyRegById(ctx echo.Context) error {
	agentSetDo := new(input.OneAgencyReg)
	code := global.ValidRequest(agentSetDo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	agentSetDo.SiteId = user.SiteId
	data, _, err := secondDistributionRegisterSetupBeen.GetOneAgencyReg(agentSetDo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
