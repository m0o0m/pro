package member

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

type PayRecordController struct {
	controllers.BaseController
}

//投注记录//TODO:GameResult返回部分还没解析出来
func (prc *PayRecordController) Cathectic(ctx echo.Context) error {
	transactionRecord := new(input.TransactionRecord)
	code := global.ValidRequestMember(transactionRecord, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	memberId := member.Id
	transactionRecord.MemberId = memberId
	listparam := new(global.ListParams)
	//获取listparam的数据
	prc.GetParam(listparam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if transactionRecord.BetTime != "" {
		tm, _ := time.ParseInLocation("2006-01-02", transactionRecord.BetTime, loc)
		times.StartTime = tm.Unix()
		times.EndTime = tm.Unix()
	}
	list, count, err := betRecordInfoBean.GetTransactionRecord(transactionRecord, times, listparam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//现金流水
func (prc *PayRecordController) MemberCashRecord(ctx echo.Context) error {
	memberCashRecords := new(input.MemberCashRecords)
	code := global.ValidRequestMember(memberCashRecords, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	memberCashRecords.MemberId = member.Id
	listparam := new(global.ListParams)
	//获取listparam的数据
	prc.GetParam(listparam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if memberCashRecords.CreateTime != "" {
		tm, _ := time.ParseInLocation("2006-01-02", memberCashRecords.CreateTime, loc)
		times.StartTime = tm.Unix()
		times.EndTime = tm.Unix()
	}
	//时间区间为最近三个月(90天)
	diff := times.EndTime - times.StartTime
	day := diff / 3600 / 24
	if day > 90 {
		return ctx.JSON(200, global.ReplyError(30182, ctx))
	}
	list, count, err := betRecordInfoBean.GetMemberCashRecord(memberCashRecords, times, listparam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//获取商品类型(二级菜单栏)
func (*PayRecordController) GetProductType(ctx echo.Context) error {
	data, err := betRecordInfoBean.GetProductType()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取商品类型下的商品
func (*PayRecordController) GetProductName(ctx echo.Context) error {
	typeId := new(input.ProductTypeId)
	code := global.ValidRequestMember(typeId, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := betRecordInfoBean.GetProduct(typeId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
