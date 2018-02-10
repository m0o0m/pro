//[控制器] [平台]稽核日志管理
package customer

import (
	"controllers"
	"encoding/json"
	"framework/uuid"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strconv"
	"time"
)

//稽核日志管理
type AuditLogController struct {
	controllers.BaseController
}

//稽核日志查询
func (c *AuditLogController) GetAuditLogList(ctx echo.Context) error {
	auditLogAdmin := new(input.AuditLogAdmin)
	code := global.ValidRequestAdmin(auditLogAdmin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	c.GetParam(listParam, ctx)
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if auditLogAdmin.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", auditLogAdmin.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if auditLogAdmin.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", auditLogAdmin.EndTime, loc)
		times.EndTime = et.Unix()
	}
	list, count, err := auditsBean.AuditLogList(auditLogAdmin, listParam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, list, int64(len(list)), count, ctx))
}

//稽核列表查询
func (c *AuditLogController) GetAuditRecordList(ctx echo.Context) error {
	auditAdmin := new(input.AuditLogAdmin)
	code := global.ValidRequestAdmin(auditAdmin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	c.GetParam(listParam, ctx)
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if auditAdmin.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", auditAdmin.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if auditAdmin.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", auditAdmin.EndTime, loc)
		times.EndTime = et.Unix()
	}
	list, count, err := auditsBean.AuditList(auditAdmin, listParam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, list, int64(len(list)), count, ctx))
}

//即时稽核查询
func (c *AuditLogController) GetAuditNowList(ctx echo.Context) error {
	audit := new(input.AuditNow)
	code := global.ValidRequest(audit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := new(global.MemberRedisToken)
	member.Site = audit.SiteId
	member.Account = audit.Account
	//查询稽核表
	data, err := drawMoney.GetDrawData(member)
	//切割时间
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	//出款展示数据 最终成型的稽核数据结构体
	sdata := new(back.MemberAuditOutBack)
	var etime = time.Now().Unix() //当前时间
	if data != nil {
		var ftime = data[0].BeginTime       //第一笔稽核开始时间
		today0, today1 := global.GetToday() //获取当天0点和23点59分59秒的时间戳
		var reportValid []back.AuditBet
		var sbettime, ebettime int64
		mp := make(map[int]*back.AuditBetRecord)
		if ftime > today0 && ftime < today1 {
			//稽核计算日期在当天，则只需要查询注单表即可
			sbettime = today0 //注单表起始时间
			ebettime = etime  //注单表起始时间
			//查询注单表
			betValid, err := drawMoney.GetBetValid(member.Account, member.Site, sbettime, ebettime)
			for _, v := range betValid {
				mp[v.GameType] = &v
			}
			if err != nil {
				global.GlobalLogger.Error("err:s%", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		} else {
			//切割时间
			sreporttime := global.GetTimeStart(data[0].BeginTime) //统计表起始时间
			ereporttime := today0                                 //统计表结束时间
			sbettime = today0                                     //注单表起始时间
			ebettime = etime                                      //注单表起始时间
			//查询统计表
			reportValid, err = drawMoney.GetReportValid(member.Account, member.Site, sreporttime, ereporttime)
			if err != nil {
				global.GlobalLogger.Error("err:s%", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			//查询注单表
			betValid, err := drawMoney.GetBetValid(member.Account, member.Site, sbettime, ebettime)
			if err != nil {
				global.GlobalLogger.Error("err:s%", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			for k, v := range betValid {
				mp[v.GameType] = &betValid[k]
			}
			for k, v := range reportValid {
				if m, ok := mp[v.GameType]; ok {
					m.BetValid += reportValid[k].BetValid
				} else {
					temp := back.AuditBetRecord{
						SiteId:      reportValid[k].SiteId,
						SiteIndexId: reportValid[k].SiteIndexId,
						Username:    reportValid[k].Account,
						BetValid:    reportValid[k].BetValid,
						GameType:    reportValid[k].GameType,
					}
					mp[v.GameType] = &temp
				}
			}
			mergeValid, _ := json.Marshal(&mp)
			global.GlobalLogger.Error("%s", string(mergeValid))
		}
		sdata.BeginTime = ftime                     //稽核开始时间
		sdata.EndTime = etime                       //稽核结束时间
		sdata.Money = data[0].Money                 //存款金额
		sdata.NormalMoney = data[0].NormalMoney     //常态稽核金额
		sdata.MultipleMoney = data[0].MultipleMoney //综合稽核金额
		sdata.AdminMoney = data[0].AdminMoney       //扣除的行政费用
		sdata.DepositMoney = data[0].DepositMoney   //优惠金额
		sdata.RelaxMoney = data[0].RelaxMoney       //放宽额度
		//初始化打码数据
		sdata.Vdbet = 0
		sdata.Dzbet = 0
		sdata.Spbet = 0
		sdata.Fcbet = 0
		if mp != nil { //打码数据不为空，整合有效打码
			for _, v := range mp {
				if v.GameType == 1 {
					sdata.Vdbet = v.BetValid
				}
				if v.GameType == 2 || v.GameType == 3 {
					sdata.Dzbet += v.BetValid
				}
				if v.GameType == 4 {
					sdata.Fcbet = v.BetValid
				}
				if v.GameType == 5 {
					sdata.Spbet = v.BetValid
				}
			}
		}
		sdata.Allbet = sdata.Vdbet + sdata.Dzbet + sdata.Fcbet + sdata.Spbet //所有打码合计
		if sdata.Allbet+float64(sdata.RelaxMoney) >= sdata.NormalMoney {     //判断常态稽核是否通过
			sdata.IsNormal = 1 //通过
			sdata.AdminMoney = 0.00
		} else {
			sdata.IsNormal = 2 //未通过       扣除行政费
		}
		if sdata.MultipleMoney > 0 { //开启了综合稽核
			if sdata.Allbet >= sdata.MultipleMoney { //综合稽核是否通过
				sdata.IsMultiple = 1 //通过
				sdata.DepositMoney = 0.00
			} else {
				sdata.IsMultiple = 2 //未通过，扣除优惠
			}
		} else {
			sdata.DepositMoney = 0.00
		}
	} else {
		sdata.BeginTime = etime
		sdata.EndTime = etime
		sdata.Money = 0         //存款金额
		sdata.AdminMoney = 0    //扣除的行政费用
		sdata.DepositMoney = 0  // 优惠金额
		sdata.MultipleMoney = 0 //综合稽核
		sdata.NormalMoney = 0   //常态稽核
		sdata.RelaxMoney = 0    //放宽额度
		sdata.Vdbet = 0         // 视讯打码
		sdata.Dzbet = 0         //电子打码
		sdata.Spbet = 0         // 体育打码
		sdata.Fcbet = 0         // 彩票打码
		sdata.Allbet = 0        //打码总计
		sdata.IsNormal = 1      //通过   //通过常态稽核 1为通过，2为未通过，未通过扣取行政费用
		sdata.IsMultiple = 1    //通过  //是否通过综合稽核 1为通过，2为未通过，未通过扣取优惠金额
		sdata.Charge = 0        //出款手续费
		sdata.OutStatus = 1     //可出款状态 1未可出款 2未不可出款（提出金额减去费用小于0）
	}
	key := uuid.NewV4().String()
	sdata.OrderNum = key
	redisKey := member.Site + "_" + member.SiteIndex + "_" + strconv.FormatInt(member.Id, 10) + "outCharge"
	b, err := json.Marshal(sdata)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	err = global.GetRedis().Set(redisKey, b, 0).Err()
	return ctx.JSON(200, global.ReplyItem(sdata))
}
