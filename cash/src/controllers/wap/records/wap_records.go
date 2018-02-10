package records

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

//交易记录管理
type WapRecordController struct {
	controllers.BaseController
}

//投注记录列表
func (wrc *WapRecordController) WapBetRecordList(ctx echo.Context) error {
	betRecordList := new(input.WapBetRecord)
	code := global.ValidRequestMember(betRecordList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if betRecordList.DateTime != "" {
		st, _ := time.ParseInLocation("2006-01-02", betRecordList.DateTime, loc)
		times.StartTime = st.Unix()
		dd, _ := time.ParseDuration("24h")
		t1 := st.Add(dd)
		times.EndTime = t1.Unix()
	}
	//获取登录会员信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	betRecordList.MemberId = member.Id
	//获取listparam的数据
	listParams := new(global.ListParams)
	wrc.GetParam(listParams, ctx)
	//获取投注记录列表
	list, count, err := betRecordInfo.WapBetRecordList(betRecordList, listParams, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, list, int64(len(list)), count, ctx))
}

//现金流水列表
func (wrc *WapRecordController) WapMemberCashRecordList(ctx echo.Context) error {
	memberCashRecords := new(input.WapMemberCashRecords)
	code := global.ValidRequestMember(memberCashRecords, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if memberCashRecords.DateTime != "" {
		st, _ := time.ParseInLocation("2006-01-02", memberCashRecords.DateTime, loc)
		times.StartTime = st.Unix()
		dd, _ := time.ParseDuration("24h")
		t1 := st.Add(dd)
		times.EndTime = t1.Unix()
	}
	//获取登录会员信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	memberCashRecords.MemberId = member.Id
	listparam := new(global.ListParams)
	//获取listparam的数据
	wrc.GetParam(listparam, ctx)
	//获取现金流水列表
	list, count, err := betRecordInfo.WapMemberCashRecord(memberCashRecords, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}
