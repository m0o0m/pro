package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
)

type RedPacketSetController struct {
	controllers.BaseController
}

//添加一个红包设置
func (*RedPacketSetController) Add(ctx echo.Context) error {
	reqData := new(input.RedPacketSetAdd)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//活动开始结束时间
	reqData.StartTimestamp, reqData.EndTimestamp, code = global.FormatTime2Timestamp(reqData.StartTime, reqData.EndTime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//存款开始结束时间
	reqData.InStartTimestamp, reqData.InEndTimestamp, code = global.FormatTime2Timestamp(reqData.InStartTime, reqData.InEndTime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//有效打码开始结束时间
	reqData.AuditStartTimestamp, reqData.AuditEndTimestamp, code = global.FormatTime2Timestamp(reqData.AuditStartTime, reqData.AuditEndTime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	reqData.CreateIp = ctx.RealIP()
	user := ctx.Get("user").(*global.RedisStruct)
	reqData.SiteId = user.SiteId
	reqData.CreateUid = user.Id
	num, err := redPacketSetBean.Add(reqData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		return ctx.JSON(500, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//查询红包设置列表
func (*RedPacketSetController) List(ctx echo.Context) error {
	reqData := new(input.RedPacketSetList)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	reqData.SiteId = user.SiteId
	var sTime, eTime int64
	if reqData.CreateTimeStart != "" {
		sTime, code = global.FormatTime2Timestamp2(reqData.CreateTimeStart)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		sTime = 0
	}
	if reqData.CreateTimeEnd != "" {
		eTime, code = global.FormatTime2Timestamp2(reqData.CreateTimeEnd)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		eTime = 0
	}
	redPackets, err := redPacketSetBean.FindList(reqData.SiteId, reqData.SiteIndexId, sTime, eTime, reqData.Status)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(redPackets, len(redPackets)))
}

//查询红包设置详情
func (*RedPacketSetController) ListInfo(ctx echo.Context) error {
	reqData := new(input.RedPacketSetListInfo)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	reqData.SiteId = user.SiteId
	redPackets, has, err := redPacketSetBean.FindListInfo(reqData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10142, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(redPackets))
}

//查看红包
func (*RedPacketSetController) RedBagInfo(ctx echo.Context) error {
	rI := new(input.RedBagSee)
	code := global.ValidRequest(rI, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := redPacketSetBean.RedBagSeeById(rI)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var num back.RedBagNumTotal
	if len(data) > 0 {
		for _, v := range data {
			num.RedBagMoney = num.RedBagMoney + v.Money
			if v.MakeSure == 2 {
				num.Already = num.Already + 1
				num.AlreadyMoney = num.AlreadyMoney + v.BalanceMoney
			}
		}
		num.RedBag = int64(len(data))
		num.Spare = num.RedBag - num.Already
		num.SpareMoney = num.RedBagMoney - num.AlreadyMoney
	} else {
		num.RedBag = 0
		num.Spare = 0
		num.SpareMoney = 0
		num.RedBagMoney = 0
		num.Already = 0
		num.AlreadyMoney = 0
	}
	var list = make(map[string]interface{})
	list["data"] = data
	list["num"] = num
	return ctx.JSON(200, global.ReplyItem(list))
}

//修改红包设置
func (*RedPacketSetController) RedBagChange(ctx echo.Context) error {
	reqData := new(input.RedPacketSetChange)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//活动开始结束时间
	reqData.StartTimestamp, reqData.EndTimestamp, code = global.FormatTime2Timestamp(reqData.StartTime, reqData.EndTime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//存款开始结束时间
	reqData.InStartTimestamp, reqData.InEndTimestamp, code = global.FormatTime2Timestamp(reqData.InStartTime, reqData.InEndTime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//有效打码开始结束时间
	reqData.AuditStartTimestamp, reqData.AuditEndTimestamp, code = global.FormatTime2Timestamp(reqData.AuditStartTime, reqData.AuditEndTime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	reqData.CreateIp = ctx.RealIP()
	user := ctx.Get("user").(*global.RedisStruct)
	reqData.SiteId = user.SiteId
	reqData.CreateUid = user.Id
	num, err := redPacketSetBean.Change(reqData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		return ctx.JSON(500, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//红包设置终止
func (*RedPacketSetController) RedBagRemove(ctx echo.Context) error {
	rD := new(input.RedPacketSetDelete)
	code := global.ValidRequest(rD, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := redPacketSetBean.Delete(rD)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50188, ctx))
	}
	return ctx.NoContent(204)
}
