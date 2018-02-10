package member

import (
	"framework/render"
	"global"
	"models/input"
	"time"

	"controllers/front/page"
	"fmt"
	"github.com/labstack/echo"
	"models/back"
	"strings"
)

type RedController struct {
	page.PageBaseController
}

//红包
func (c *RedController) RedPacketLog(ctx echo.Context) error {
	//获取站点信息
	infos, flag, err := GetSiteInfo(ctx)

	if err != nil || !flag {
		global.GlobalLogger.Error("error%s", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	siteId := infos.SiteId
	siteIndexId := infos.SiteIndexId
	//红包返回数据
	sdata := []back.RedPacket{}

	//获取红包数据
	data, err := redPacketSetBean.SiteFind(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	for _, v := range data {
		info := back.RedPacket{}
		info.Id = v.Id
		info.Title = v.Title
		info.MaxCount = v.MaxCount
		info.StartTime = v.StartTime
		info.InStartTime = v.InStartTime
		info.InEndTime = v.InEndTime
		info.InSum = v.InSum
		info.AuditStartTime = v.AuditStartTime
		info.AuditEndTime = v.AuditEndTime
		info.BetSum = v.BetSum
		info.LevelId = v.LevelId
		info.TotalMoney = v.TotalMoney
		info.MinMoney = v.MinMoney
		info.RedNum = v.RedNum
		info.CreateTime = v.CreateTime
		info.IsIp = v.IsIp
		info.StyleId = v.StyleId
		info.IsShow = v.IsShow
		info.AppointMoney = v.AppointMoney
		info.RedType = v.RedType
		info.Status = v.Status
		info.IsGenerate = v.IsGenerate
		timeNum := time.Now().Unix()
		info.Opencount = timeNum - v.StartTime
		info.Closecount = v.EndTime - timeNum
		info.Pic = 1
		sdata = append(sdata, info)
	}
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//抢红包
func (*RedController) GetSnatch(ctx echo.Context) error {
	snatch := new(input.GetSnatch)
	code := global.ValidRequest(snatch, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	member := ctx.Get("member").(*global.MemberRedisToken)

	data, err := redPacketSetBean.GetOne(snatch.SetId) //查询红包活动
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	if data.Id == 0 { //没有该红包活动
		return ctx.JSON(200, global.ReplyError(71106, ctx))
	}
	if data.StartTime > time.Now().Unix() {
		return ctx.JSON(200, global.ReplyError(91201, ctx))
	}
	var Ip string
	if data.IsIp == 2 {
		Ip = ctx.RealIP()
		count, err := redPacketSetBean.IpRed(snatch.SetId, Ip) //查询红包活动

		fmt.Println("红包数据：", data)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return render.PageErr(60000, ctx)
		}
		if count > 0 { //没有该红包活动
			return ctx.JSON(200, global.ReplyError(71108, ctx))
		}
	}
	if data.LevelId != "" {
		strArr := strings.Split(data.LevelId, ",")
		for _, v := range strArr {
			if member.LevelId == v {
				return ctx.JSON(200, global.ReplyError(71105, ctx))
			}
		}
		count, err := redPacketSetBean.UserRed(snatch.SetId, member.Id) //查询红包活动
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return render.PageErr(60000, ctx)
		}
		if count >= data.MaxCount { //已经抢过了
			return ctx.JSON(200, global.ReplyError(71109, ctx))
		}
	}

	//查询可分配的红包
	b, redinfo, err := redPacketSetBean.GetRebInfo(snatch.SetId) //查询红包活动
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	if !b { //没有查到可分配的红包
		return ctx.JSON(200, global.ReplyError(71107, ctx))
	}

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	//修改红包状态
	redinfo.MemberId = member.Id
	redinfo.Account = member.Account
	redinfo.PType = 1
	num, err := redPacketSetBean.SetRebMakeSure(redinfo, sess) //修改红包状态
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//增加会员金额
	code, num, newCashRecord, err := redPacketSetBean.SetRebMemBalance(member.Id, redinfo.Money, member.Site, sess)

	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if num == 0 {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	if data.IsGenerate != 2 {
		return ctx.JSON(200, global.ReplyError(71104, ctx))
	}

	agencyInfo, _, err := thirdAgencyBean.BaseInfo(&input.ThirdAgencyInfo{member.Site, member.SiteIndex, newCashRecord.AgencyId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	newCashRecord.AgencyAccount = agencyInfo.Account

	//会员现金流水
	num, err = memberCashBean.AddNewRecordSess(&newCashRecord, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if num == 0 {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	err = sess.Commit()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(redinfo))
}
