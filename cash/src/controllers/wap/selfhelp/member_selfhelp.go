//会员自助反水
package selfhelp

import (
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"time"
)

type SelfHelpController struct{}

//自助反水列表
func (*SelfHelpController) GetSelfHelpList(ctx echo.Context) error {
	//获取当天时间
	t := time.Now().Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02", t, loc)
	startTime := t1.Unix()
	dd, _ := time.ParseDuration("24h")
	end := t1.Add(dd)
	endTime := end.Unix()
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//获取会员所属层级详情，判断是否开启自助反水
	memberInfo, flag, err := memberBean.GetInfoById(memberRedis.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag || memberInfo.Id == 0 {
		return ctx.JSON(200, global.ReplyError(60051, ctx))
	}
	//获取该会员所处层级详情
	levelInfo, flag, err := memberLevelBean.GetLevelInfo(memberInfo.LevelId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//找不到该层级
	if !flag || levelInfo.LevelId == "" {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//未开启自助反水
	if levelInfo.IsSelfRebate == 2 {
		return ctx.JSON(200, global.ReplyError(60213, ctx))
	}
	var memberRetreatWaterTotal float64
	//首先获取会员今日已经反水的额度
	waterSelf, flag, err := memberRetreatWaterSelfBean.GetMemberRetreatToday(memberRedis, startTime, endTime)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//会员自助反水总记录一天一条的记录不存在,那么就去反水记录表查询今天的数据,得出总的已经反水的额度
	if !flag || waterSelf.Id == 0 {
		total, err := memberRetreatWaterSelfBean.GetMemberRetreatNowDayTotal(memberRedis, startTime, endTime)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if total == 0 {
			memberRetreatWaterTotal = 0
		}
	} else {
		memberRetreatWaterTotal = waterSelf.Money
	}
	//拿出每个游戏的有效投注额度
	data, err := memberRetreatWaterSelfBean.GetProductReWaterRate(memberRedis, startTime, endTime)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取不同投注额的反水设定,根据有效投注额度取得反水比例
	lenGth := len(data)
	for i := 0; i < lenGth; i++ {
		//比例,最高优惠,错误信息
		rate, disocuntup, err := memberRetreatWaterSelfBean.GetInfoBySiteAndProSetId(memberRedis, data[i].ProductId, data[i].BetValid)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		data[i].Rate = rate
		//如果刚开始进来可以返回有效投注总额，将这里删除
		//data[i].BetValid = 0
		//data[i].ReWaterQutoa = 0
		//如果列表需要金额，将这里开启
		if data[i].BetValid*(rate/100) >= float64(disocuntup) {
			data[i].ReWaterQutoa = int(disocuntup)
		} else {
			data[i].ReWaterQutoa = int(data[i].BetValid * (rate / 100))
		}
	}
	backWater := new(back.ReWater)
	backWater.AlreadyReNowDate = memberRetreatWaterTotal
	backWater.BackList = data
	return ctx.JSON(200, global.ReplyItem(backWater))
}

//一键查看反水额度
func (*SelfHelpController) OneClickSee(ctx echo.Context) error {
	//获取当天时间
	t := time.Now().Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02", t, loc)
	//当天零点的时间戳
	startTime := t1.Unix()
	dd, _ := time.ParseDuration("24h")
	end := t1.Add(dd)
	//另天零点的时间戳
	endTime := end.Unix()
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//获取会员所属层级详情，判断是否开启自助反水(避免直接请求依然加上是否开启自助反水的校验)
	memberInfo, flag, err := memberBean.GetInfoById(memberRedis.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag || memberInfo.Id == 0 {
		return ctx.JSON(200, global.ReplyError(60051, ctx))
	}
	//获取该会员所处层级详情
	levelInfo, flag, err := memberLevelBean.GetLevelInfo(memberInfo.LevelId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//找不到该层级
	if !flag || levelInfo.LevelId == "" {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//未开启自助反水
	if levelInfo.IsSelfRebate == 2 {
		return ctx.JSON(500, global.ReplyError(60213, ctx))
	}
	//拿出每个游戏的有效投注额度
	data, err := memberRetreatWaterSelfBean.GetProductReWaterRate(memberRedis, startTime, endTime)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取不同投注额的反水设定,根据有效投注额度取得反水比例
	lenGth := len(data)
	for i := 0; i < lenGth; i++ {
		//比例,最高优惠,错误信息
		rate, discountup, err := memberRetreatWaterSelfBean.GetInfoBySiteAndProSetId(memberRedis, data[i].ProductId, data[i].BetValid)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		data[i].Rate = rate
		//反水金额计算
		if data[i].BetValid*(rate/100) >= float64(discountup) {
			data[i].ReWaterQutoa = int(discountup)
		} else {
			data[i].ReWaterQutoa = int(data[i].BetValid * (rate / 100))
		}
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//全部领取 todo 采用的是前端上传有效总额度和总反水，需要修改为后台查询这个总额或者在下发所有的额度的时候缓存，然后在这里从缓存取出来，但是可能这个中间会有变化？？
func (*SelfHelpController) AllReWaterPickUp(ctx echo.Context) error {
	getAllRewater := new(input.OneClickGetAllReWater)
	code := global.ValidRequest(getAllRewater, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当天时间
	t := time.Now().Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02", t, loc)
	//当天零点的时间戳
	startTime := t1.Unix()
	dd, _ := time.ParseDuration("24h")
	end := t1.Add(dd)
	//另天零点的时间戳
	endTime := end.Unix()
	memberInfo := ctx.Get("member").(*global.MemberRedisToken)
	//查看记录表当天的退水金额
	rebateWater, err := memberRetreatWaterSelfBean.GetMemberRetreatNowDayTotal(memberInfo, startTime, endTime)
	//判断退水金额是否已被领取
	if rebateWater >= getAllRewater.RewaterTotal {
		return ctx.JSON(200, global.ReplyError(30245, ctx))
	}
	getAllRewater.RewaterTotal = getAllRewater.RewaterTotal - rebateWater
	count, err := memberRetreatWaterSelfBean.OneClickGetAllReWater(memberInfo, getAllRewater)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30294, ctx))
	}
	return ctx.NoContent(204)
}

//推广返佣
func (*SelfHelpController) WapMemberRebate(ctx echo.Context) error {
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//根据站点id和站点前台id判断是否开启会员推广
	has, err := memberRebate.WapMemberSpread(member.Site, member.SiteIndex)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//未开启
	if !has {
		return ctx.JSON(200, global.ReplyError(30228, ctx))
	}
	list, err := memberRebate.WapMemberRebate(member.Site, member.SiteIndex)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//排行榜
func (*SelfHelpController) Leaderboards(ctx echo.Context) error {
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//根据站点id和站点前台id获取排行榜人数系数
	rebateRanking, has, err := memberRebate.WapMemberRebateRanking(member.Site, member.SiteIndex)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30229, ctx))
	}
	var ranking back.WapRanking
	var rankings []back.WapRanking
	//获取2017-12-12时间戳
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02", "2017-12-12", loc)
	t12 := t1.Unix()
	//获取当前日期的时间戳
	t2, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), loc)
	tn := t2.Unix()
	//一天的秒数
	var d int64
	d = 24 * 60 * 60
	//计算距离2017-12-12有多少天
	day := (tn - t12) / d
	//每天上涨的人数（天数*系数）
	var num int64
	if day > 0 {
		num = day * int64(rebateRanking)
	}
	//获取生成的十个账号账号
	accounts := []string{"uzii", "mata", "mlxg", "xiaohu", "mystic", "bengi", "deft", "keaidian", "marin", "faker"}
	for k := range accounts {
		ranking.Account = accounts[k]
		ranking.PromotionNumber = int64(k+50) + num
		rankings = append(rankings, ranking)
	}
	return ctx.JSON(200, global.ReplyItem(rankings))

}

//查看单个的反水额度
func (*SelfHelpController) GetSingleReWater(ctx echo.Context) error {
	singleWater := new(input.SingleReWater)
	code := global.ValidRequest(singleWater, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//根据productid，日期获取有效投注总额
	info, flag, err := memberRetreatWaterSelfBean.GetBetValidById(memberRedis, singleWater)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if info.Id == 0 || !flag {
		return ctx.JSON(200, global.ReplyError(60215, ctx))
	}
	backReWater := new(back.ReWaterSingle)
	//获取不同投注额的反水设定,根据有效投注额度取得反水比例
	rate, disocuntup, err := memberRetreatWaterSelfBean.GetInfoBySiteAndProSetId(memberRedis, singleWater.ProductId, info.BetValid)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if info.BetValid*(rate/100) >= float64(disocuntup) {
		backReWater.ReWaterQutoa = int(disocuntup)
	} else {
		backReWater.ReWaterQutoa = int(info.BetValid * (rate / 100))
	}
	return ctx.JSON(200, global.ReplyItem(backReWater))
}
