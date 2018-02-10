package cash

import (
	"controllers"
	"framework/uuid"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/schema"
)

type RedPacketLogController struct {
	controllers.BaseController
}

//生成红包
func (*RedPacketLogController) GenerateRedPacket(ctx echo.Context) error {
	reqData := new(input.RedPacketSet)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询设置
	redPacketSet, err := redPacketSetBean.GetOne(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if redPacketSet.IsGenerate == 2 {
		return ctx.JSON(500, global.ReplyError(71023, ctx))
	}

	//生成多个红包
	redPackets, code := redPacketLogBean.GenerateRedPacket(redPacketSet.TotalMoney, redPacketSet.RedNum, redPacketSet.MinMoney)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	redPacketLogTemplate := new(schema.RedPacketLog)
	redPacketLogTemplate.SetId = redPacketSet.Id                      //红包设置id
	redPacketLogTemplate.SiteId = redPacketSet.SiteId                 //站点id
	redPacketLogTemplate.SiteIndexId = redPacketSet.SiteIndexId       //前台id
	redPacketLogTemplate.CreateIp = ctx.RealIP()                      //创建ip
	redPacketLogTemplate.StartTime = redPacketSet.StartTime           //开始时间
	redPacketLogTemplate.EndTime = redPacketSet.EndTime               //结束时间
	redPacketLogTemplate.InStartTime = redPacketSet.InStartTime       //存款起始时间
	redPacketLogTemplate.InEndTime = redPacketSet.InEndTime           //存款结束时间
	redPacketLogTemplate.InSum = redPacketSet.InSum                   //存款额度
	redPacketLogTemplate.AuditStartTime = redPacketSet.AuditStartTime //有效打码开始时间
	redPacketLogTemplate.AuditEndTime = redPacketSet.AuditEndTime     //有效打码结束时间
	redPacketLogTemplate.BetSum = redPacketSet.BetSum                 //有效打码量
	redPacketLogTemplate.MinMoney = redPacketSet.MinMoney             //红包最小额度
	redPacketLogTemplate.LevelId = redPacketSet.LevelId               //可参加活动会员的分组，0为无限制
	redPacketLogTemplate.Title = redPacketSet.Title                   //红包名
	redPacketLogTemplate.MakeSure = 1                                 //是否被抢1,未被抢,2已抢
	redPacketLogTemplate.PType = reqData.ClientType                   //客户端类型0pc 1wap 2android 3ios
	redPacketLogTemplate.IsIp = 2                                     //1为限制ip,2为不限制
	//redPacketLogTemplate.Uuid           string  `xorm:"uuid"`             //红包的uuid
	//redPacketLogTemplate.Money  =          float64 `xorm:"money"`            //红包金额
	//redPacketLogTemplate.MemberId       = redPacketSet int     `xorm:"member_id"`        //用户id
	//redPacketLogTemplate.Account        = redPacketSet  //用户名
	//redPacketLogTemplate.CreateTime     = redPacketSet  //创建时间
	//redPacketLogTemplate.BalanceMoney   = redPacketSet   //以下三个参数,不知道是什么用的
	//redPacketLogTemplate.LevelEs        = redPacketSet
	//redPacketLogTemplate.Finish         = redPacketSet

	var redPacketLogs = make([]*schema.RedPacketLog, len(redPackets))
	for i, value := range redPackets {
		temp := *redPacketLogTemplate
		temp.Money = value                //红包里的钱
		temp.Uuid = uuid.NewV4().String() //uuid
		redPacketLogs[i] = &temp
	}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	num, err := redPacketLogBean.AddRedPacket(redPacketLogs, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != int64(len(redPacketLogs)) {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//修改设置为已生成
	num, err = redPacketSetBean.SetGenerate(reqData.Id)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
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
	return ctx.NoContent(204)
}
