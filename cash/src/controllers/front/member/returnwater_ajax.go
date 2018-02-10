package member

import (
	"encoding/json"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"strconv"
	"time"
)

type ReturnWaterAjax struct {
}

type Today struct {
	NowDate string //今天的年月日
}

type ReturnBetData struct {
	ProductId   int64
	ProductName string
	VType       string
	Rate        float64
	BetValid    float64
	RateMoney   float64
}

//将反水信息存入redis
type WaterData struct {
	Data      []ReturnBetData
	Count     float64
	HasReturn float64
}

type ReturnData struct {
	ProductName string
	Rate        float64
	BetValid    float64
	ReturnMoney float64
}

type WaterRData struct {
	IsSelf int
	RData  []ReturnData
}

//自助反水是否开启
func (m *ReturnWaterAjax) GetMemberIsSelf(ctx echo.Context) error {
	var result WaterRData
	member := ctx.Get("member").(*global.MemberRedisToken)
	//获取会员层级信息
	//获取该会员所处层级详情
	levelInfo, flag, err := member_level_bean.GetLevelInfo(member.LevelId)
	if err != nil {
		result.IsSelf = 2 //1开启 2未开启
	}
	//找不到该层级
	if !flag || levelInfo.LevelId == "" {
		result.IsSelf = 2 //1开启 2未开启
	}
	result.IsSelf = int(levelInfo.IsSelfRebate)
	idDel, _ := siteProductBean.GetProductDel(member.Site, member.SiteIndex)
	//获取站点商品
	productall, _ := siteProductBean.GetProductAll(idDel)
	var data ReturnData
	for _, v := range productall {
		data.ProductName = v.ProductName
		data.Rate = 0
		data.BetValid = 0
		data.ReturnMoney = 0
		result.RData = append(result.RData, data)
	}
	return ctx.JSON(200, global.ReplyItem(result))
}

//获取各个游戏可反水额度
func (m *ReturnWaterAjax) GetMemberGameBet(ctx echo.Context) error {
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	var memberRetreatWaterTotal float64
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
	//会员今日已反水额度
	waterSelf, flag, err := member_retreat_water_selfbean.GetMemberRetreatToday(memberRedis, startTime, endTime)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//会员自助反水总记录一天一条的记录不存在,那么就去反水记录表查询今天的数据,得出总的已经反水的额度
	if !flag || waterSelf.Id == 0 {
		total, err := member_retreat_water_selfbean.GetMemberRetreatNowDayTotal(memberRedis, startTime, endTime)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if total == 0 {
			memberRetreatWaterTotal = 0
		}
	} else {
		memberRetreatWaterTotal = waterSelf.Money
	}
	//拿出每个游戏的有效投注额度
	data, err := member_retreat_water_selfbean.GetProductReWaterRate(memberRedis, startTime, endTime)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//从此处开始数据整合
	if data != nil {
		var betall float64
		for _, v := range data {
			betall += v.BetValid
		}
		rate, err := member_retreat_water_selfbean.GetSiteWaterSetByValidBetAll(memberRedis, betall)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		idDel, _ := siteProductBean.GetProductDel(memberRedis.Site, memberRedis.SiteIndex)
		//获取站点商品
		productall, _ := siteProductBean.GetProductAll(idDel)
		var backdata ReturnBetData
		var backWater []ReturnBetData
		for _, value := range productall {
			backdata.ProductName = value.ProductName
			backdata.ProductId = value.Id
			backdata.VType = value.VType
			backdata.BetValid = 0
			for _, val := range data {
				if value.Id == val.ProductId {
					backdata.BetValid = val.BetValid
				}
			}
			backWater = append(backWater, backdata)
		}
		var waterall, returnwater float64 //本次总计反水  实际反水
		for key, val := range backWater {
			for _, v := range rate {
				if val.ProductId == int64(v.ProductId) {
					backWater[key].Rate = v.Rate
					backWater[key].RateMoney = val.BetValid * v.Rate * 0.01
				}
			}
			waterall += backWater[key].RateMoney
		}
		returnwater = waterall - memberRetreatWaterTotal
		//如果反水总额大于0，将数据存入redis
		if returnwater > 0.01 {
			var water WaterData
			water.Data = backWater
			water.HasReturn = memberRetreatWaterTotal //今日已反水额度
			water.Count = returnwater                 //今日可反水额度
			day := time.Now().Format("20060102")
			key := memberRedis.Site + "_" + memberRedis.SiteIndex + "_" + strconv.Itoa(int(memberRedis.Id)) + "_" + day
			jsonWater, _ := json.Marshal(water)
			err = global.GetRedis().Set(key, string(jsonWater), 0).Err()
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		return ctx.JSON(200, global.ReplyCollections("data", backWater, "count", returnwater))
	} else {
		return ctx.JSON(200, global.ReplyItem(""))
	}
}

//一键存入会员反水，给会员加款
func (m *ReturnWaterAjax) PostReturnWaterSelf(ctx echo.Context) error {
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	day := time.Now().Format("20060102")
	key := memberRedis.Site + "_" + memberRedis.SiteIndex + "_" + strconv.Itoa(int(memberRedis.Id)) + "_" + day
	jsonStr, err := global.GetRedis().Get(key).Result()
	if err != nil || jsonStr == "" {
		return ctx.JSON(500, global.ReplyError(92000, ctx))
	} else {
		var Water back.WaterData
		json.Unmarshal([]byte(jsonStr), &Water)
		//开始处理反水数据
		err := member_retreat_water_selfbean.PostMemberReWater(memberRedis, Water)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(92000, ctx))
		}
	}
	global.GetRedis().Del(key).Result()
	return ctx.NoContent(204)
}
