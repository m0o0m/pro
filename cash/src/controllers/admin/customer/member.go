//[控制器] [平台]会员管理
package customer

import (
	"controllers"
	"global"
	"models/input"

	"github.com/labstack/echo"
	"models/back"
)

//代理管理
type MemberController struct {
	controllers.BaseController
}

//站点会员查询
func (c *MemberController) GetMemberList(ctx echo.Context) error {
	siteMemberInfo := new(input.SiteMemberInfo)
	code := global.ValidRequestAdmin(siteMemberInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if siteMemberInfo.StartTime != "" {
		_, _, code = global.FormatTime2Timestamp(siteMemberInfo.StartTime, siteMemberInfo.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	if siteMemberInfo.Ip != "" {
		flag := global.CheckIp(siteMemberInfo.Ip)
		if !flag {
			return ctx.JSON(200, global.ReplyError(10205, ctx))
		}
	}
	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)
	data, count, err := siteOperateBean.Member(siteMemberInfo, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	statusCount, err := siteOperateBean.StatusCount()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	var total back.OnlineNumberAndTotalAndStatus
	if len(data) > 0 {
		var i int64
		for _, v := range data {
			if v.PcStatus == 1 {
				i = i + 1
			} else if v.IosStatus == 1 {
				i = i + 1
			} else if v.WapStatus == 1 {
				i = i + 1
			} else if v.AndroidStatus == 1 {
				i = i + 1
			}
		}
		total.OnlineNumber = i

	} else {
		total.OnlineNumber = 0
	}
	total.TotalNumber = count
	total.StatusNumber = statusCount
	var list = make(map[string]interface{})
	list["data"] = data
	list["total"] = total
	return ctx.JSON(200, global.ReplyPagination(listParams, list, int64(len(data)), count, ctx))
}

//查询会员视讯余额
func (c *MemberController) GetMemberVdMoney(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//会员资料修改
func (c *MemberController) PutMemberInfoUpdate(ctx echo.Context) error {
	member := new(input.MemberDetail)
	code := global.ValidRequestAdmin(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	//判断当前账号角色是否为超级管理员.
	user := ctx.Get("admin").(global.AdminRedisStruct)
	if user.RoleId != 5 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	if member.DrawPassword != "" {
		pwd, err := global.MD5ByStr(member.DrawPassword, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		member.DrawPassword = pwd
	}

	if member.Birthday != "" {
		_, code = global.FormatDay2Timestamp2(member.Birthday)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}

	ok, err := memberBean.UpdateDetail(member, user.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(60086, ctx))
	}
	return ctx.NoContent(204)
}

//会员状态修改（停用，暂停时将会员踢线）
func (c *MemberController) PutMemberStatusUpdate(ctx echo.Context) error {
	member := new(input.MemberStatus)
	code := global.ValidRequestAdmin(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	ok, err := memberBean.Status(member)

	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(10019, ctx))
	}
	return ctx.NoContent(204)
}

//会员资料详情
func (c *MemberController) MemberInfo(ctx echo.Context) error {
	member := new(input.MemberStatus)
	code := global.ValidRequestAdmin(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := siteOperateBean.MemberInfo(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	bankInfo, err := siteOperateBean.MemberBankInfoById(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data.MemberBank = bankInfo
	return ctx.JSON(200, global.ReplyItem(data))
}
