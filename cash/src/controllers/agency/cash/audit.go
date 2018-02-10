package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"time"
)

//稽核管理
type AuditController struct {
	controllers.BaseController
}

//获取稽核日志记录(get列表)
func (pc *AuditController) AuditLogList(ctx echo.Context) error {
	combo := new(input.AuditLogGet)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	if combo.SiteId == "" {
		combo.SiteId = user.SiteId //将缓存信息传给条件
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	pc.GetParam(listParam, ctx)
	data, count, err := auditsBean.GetAuditLogList(combo, listParam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, data, int64(len(data)), count, ctx))
}

//获取即时稽核
func (pc *AuditController) GetNowAudit(ctx echo.Context) error {
	combo := new(input.AuditNow)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if combo.Account != "" {
		//验证会员账号是否存在
		_, hs, err := auditsBean.IsMemberAccount(combo.Account, combo.SiteId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !hs {
			return ctx.JSON(200, global.ReplyError(30138, ctx))
		}
	}
	//查询未处理的稽核记录
	data, err := auditsBean.GetMenmberNowAudit(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	sdata := new(back.MemberAuditNowBack)
	if data != nil {
		var ftime = data[0].BeginTime //第一笔稽核开始时间
		tm := time.Unix(ftime, 0)
		ttime := tm.Format("2006-01-02") + " 00:00:00"
		timeLayout := "2006-01-02 15:04:05"                        //转化所需模板
		loc, _ := time.LoadLocation("Local")                       //重要：获取时区
		theTime, _ := time.ParseInLocation(timeLayout, ttime, loc) //使用模板在对应时区转化为time.time类型
		stime := theTime.Unix()                                    //转化为时间戳 类型是int64
		var etime = time.Now().Unix()

		//查询改时间段内的会员打码
		betdata, err := auditsBean.GetMemberValidBet(combo.Account, combo.SiteId, stime, etime)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		sdata.BeginTime = data[0].BeginTime
		sdata.EndTime = etime
		for _, v := range data { //整合稽核为一条
			sdata.Money += v.Money
			sdata.NormalMoney += v.NormalMoney
			sdata.MultipleMoney += v.MultipleMoney
			sdata.AdminMoney += v.AdminMoney
			sdata.DepositMoney += v.DepositMoney
			sdata.RelaxMoney += v.RelaxMoney
		}
		sdata.Vdbet = 0 //初始化数据
		sdata.Dzbet = 0
		sdata.Spbet = 0
		sdata.Fcbet = 0
		if betdata != nil { //打码数据不为空，整合有效打码
			for _, v2 := range betdata {
				if v2.GameType == 1 {
					sdata.Vdbet = v2.BetValid
				}
				if v2.GameType == 2 || v2.GameType == 3 {
					sdata.Dzbet += v2.BetValid
				}
				if v2.GameType == 4 {
					sdata.Fcbet = v2.BetValid
				}
				if v2.GameType == 5 {
					sdata.Spbet = v2.BetValid
				}
			}
		}
		sdata.Allbet = sdata.Vdbet + sdata.Dzbet + sdata.Fcbet + sdata.Spbet
		if sdata.Allbet+float64(sdata.RelaxMoney) >= sdata.NormalMoney { //判断常态稽核是否通过
			sdata.IsNormal = 1 //通过
		} else {
			sdata.IsNormal = 2 //未通过	   扣除行政费
		}
		if sdata.Allbet >= sdata.MultipleMoney { //综合稽核是否通过
			sdata.IsMultiple = 1 //通过
		} else {
			sdata.IsMultiple = 2 //未通过，扣除优惠
		}
	} else {
		sdata.BeginTime = time.Now().Unix()
		sdata.BeginTime = time.Now().Unix()
		sdata.Money = 0
		sdata.NormalMoney = 0
		sdata.MultipleMoney = 0
		sdata.AdminMoney = 0
		sdata.DepositMoney = 0
		sdata.RelaxMoney = 0
		sdata.Vdbet = 0
		sdata.Dzbet = 0
		sdata.Spbet = 0
		sdata.Fcbet = 0
		sdata.Allbet = 0
		sdata.IsNormal = 1   //通过
		sdata.IsMultiple = 1 //通过
	}
	return ctx.JSON(200, global.ReplyItem(sdata))
}
