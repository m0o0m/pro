package data_merge

import (
	"encoding/json"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strconv"
	"time"
)

type ReturnWaterAjax struct {
	IsSelf int
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
	Data  []ReturnBetData
	Count float64
}

//自助反水是否开启
func (m *ReturnWaterAjax) GetMemberIsSelf(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	//获取会员层级信息
	//获取该会员所处层级详情
	levelInfo, flag, err := member_level_bean.GetLevelInfo(member.LevelId)
	if err != nil {
		m.IsSelf = 2 //1开启 2未开启
	}
	//找不到该层级
	if !flag || levelInfo.LevelId == "" {
		m.IsSelf = 2 //1开启 2未开启
	} else {
		m.IsSelf = int(levelInfo.IsSelfRebate)
	}

	return ctx.JSON(200, global.ReplyItem(m.IsSelf))
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
			water.Count = returnwater
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

//获取个人银行信息和出款银行列表
func (m *ReturnWaterAjax) GetMemberBankList(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	mi := new(input.MemberBankList)
	mi.MemberId = member.Id
	mi.SiteId = member.Site
	mi.SiteIndexId = member.SiteIndex
	data, err := MemberBankBean.MemberBankList(member.Id)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for k, v := range data {
		v.Card = global.BankStr(v.Card)
		data[k].Card = v.Card
	}
	has, err, realnameinfo := memberBean.AccountInfo(member.Account, member.Site, member.SiteIndex)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(20011, ctx))
	}
	poundage, has, err := drawMoney.GetPoundageSet(member.Site, member.SiteIndex)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10201, ctx))
	}
	sitepay, has, err := drawMoney.GetSiteSet(member.Site, member.SiteIndex)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50084, ctx))
	}
	memberOutInfo := new(back.MemberBankSiteSet)
	memberOutInfo.Poundage = poundage
	memberOutInfo.BankList = data
	memberOutInfo.SiteSet = sitepay
	memberOutInfo.RealName = realnameinfo
	return ctx.JSON(200, global.ReplyItem(memberOutInfo))
}

//获取单条出款银行卡号
func (m *ReturnWaterAjax) GetOneBank(ctx echo.Context) error {

	bankinfo := new(input.OneMemberBankInfo)
	code := global.ValidRequest(bankinfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	bankinfo.MemberId = member.Id
	data, has, err := MemberBankBean.GetOneBank(bankinfo)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data.Card = global.BankStr(data.Card)
	return ctx.JSON(200, global.ReplyItem(data))
}

//检测支付密码是否正确
func (m *ReturnWaterAjax) CheckDrawPass(ctx echo.Context) error {
	bankinfo := new(input.CheckDrawPass)
	code := global.ValidRequest(bankinfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	bankinfo.MemberId = member.Id
	data, has, err := MemberSelfInfoBean.CheckDrawPass(bankinfo)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(20018, ctx))
	}
	//加密密码
	password, err := global.MD5ByStr(bankinfo.DrawPassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data.DrawPassword != password {
		return ctx.JSON(200, global.ReplyError(30234, ctx))
	}
	return ctx.NoContent(204)
}
