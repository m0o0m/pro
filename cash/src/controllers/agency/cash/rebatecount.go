package cash

import (
	"controllers"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strconv"
	"time"
)

//退佣查询
type RebateCountController struct {
	controllers.BaseController
}

//退佣统计条件
func (pc *RebateCountController) GetList(ctx echo.Context) error {
	combo := new(input.PeriodsGet)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获得站点列表
	data1 := pc.getSiteList(combo.SiteId, combo.SiteIndexId)
	if combo.SiteIndexId == "" {
		combo.SiteIndexId = data1[0].SiteIndexId
	}
	//获取期数列表
	list, err := retirementBean.RetirementList(combo)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10120, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("list", list, "data", data1))
}

//退佣统计
func (pc *RebateCountController) CheckList(ctx echo.Context) error {
	//1.前台传入参数
	combo := new(input.RebateInput)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//fmt.Println(combo)
	// 2.获取期数对应时间区间
	gettime, err := rebateCountBean.GetTime(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var time1 int64
	var time2 int64
	time1 = gettime[0].StartTime
	time2 = gettime[0].EndTime
	//3.获取有会员的代理id
	agencylist, err := rebateCountBean.GetAgencyId(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//4.获取代理信息
	agencylist1 := make([]int64, len(agencylist))
	for k, v := range agencylist {
		agencylist1[k] = v.AgencyId
	}
	agencyinfo, err := rebateCountBean.GetAgencyInfo(agencylist1, combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	fmt.Println("agencyinfo---", agencyinfo)
	//数据来源类型
	sourcetype := make([]int, 6)
	sourcetype[0] = 0
	sourcetype[1] = 1
	sourcetype[2] = 5
	sourcetype[3] = 6
	sourcetype[4] = 9
	sourcetype[5] = 10

	// 5.获取优惠数据。数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换9优惠返水10自助返水11会员返佣
	returncompany, err := rebateCountBean.ReturnCompany(time1, time2, agencylist1, sourcetype, combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	allfee := make(map[int64]map[int]back.AllFee)
	item := make(map[int64]bool)
	for _, v := range returncompany {
		if item[v.AgencyId] == false {
			allfee[v.AgencyId] = make(map[int]back.AllFee)
			allfee[v.AgencyId][v.SourceType] = back.AllFee{}
			item[v.AgencyId] = true
		}
		allFee := back.AllFee{}
		allFee.SiteId = v.SiteId
		allFee.SiteIndexId = v.SiteIndexId
		allFee.Balance = v.Balance
		allFee.DisBalance = v.DisBalance
		allFee.AgencyId = v.AgencyId
		allFee.SourceType = v.SourceType
		allfee[v.AgencyId][v.SourceType] = allFee
	}
	//游戏类型 1视讯  2电子 3捕鱼 4彩票 5体育
	game_type := make([]int64, 5)
	game_type[0] = 1
	game_type[1] = 2
	game_type[2] = 3
	game_type[3] = 4
	game_type[4] = 5
	//6.获取所有报表
	all_report, err := rebateCountBean.GetReport(time1, time2, agencylist1, game_type, combo)
	report := make(map[int64]map[int64][]back.ReportAll)
	item1 := make(map[int64]bool)
	for _, v := range all_report {
		if item1[v.AgencyId] == false {
			report[v.AgencyId] = make(map[int64][]back.ReportAll)
			report[v.AgencyId][v.GameType] = []back.ReportAll{}
			item1[v.AgencyId] = true
		}
		allFee := back.ReportAll{}
		allFee.GameType = v.GameType
		allFee.BetValid = v.BetValid
		allFee.BetAll = v.BetAll
		allFee.Jack = v.Jack
		allFee.Num = v.Num
		allFee.WinNum = v.WinNum
		allFee.Win = v.Win
		allFee.VType = v.VType
		allFee.DayTime = v.DayTime
		allFee.AgencyId = v.AgencyId
		report[v.AgencyId][v.GameType] = append(report[v.AgencyId][v.GameType], allFee)
	}
	//7.获取有效会员门槛
	rebateset, err := rebateCountBean.GetRebate(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取代理退佣退水设定
	poundageset := rebateCountBean.GetPoundageSet(rebateset[0].Id)
	PoundageList := back.PoundageRateList{}
	WaterList := back.PoundageWaterList{}
	for _, v := range poundageset {
		switch v.ProductId {
		case 1:
			PoundageList.VideoRate = v.RebateRadio
			WaterList.VideoWater = v.WaterRadio
		case 2:
			PoundageList.ElectRate = v.RebateRadio
			WaterList.ElectWater = v.WaterRadio
		case 3:
			PoundageList.FishRate = v.RebateRadio
			WaterList.FishWater = v.WaterRadio
		case 4:
			PoundageList.LotteryRate = v.RebateRadio
			WaterList.LotteryWater = v.WaterRadio
		case 5:
			PoundageList.SportRate = v.RebateRadio
			WaterList.SportWater = v.WaterRadio
		}
	}
	//8.有效会员数量
	ReportData := make(map[int64]back.ReportData)
	reportDataList := make(map[int64]map[int64]back.ReportData)
	for k, v := range report {

		report_data1 := back.ReportData{}
		report_data2 := make(map[int64]back.ReportData)
		for k1, v1 := range v {
			report_data := back.ReportData{}
			for _, v2 := range v1 {
				if v2.BetValid >= float64(rebateset[0].ValidMoney) {
					//分类总计
					report_data.ValidMember += 1
					report_data.BetValid += v2.BetValid
					report_data.Jack += v2.Jack
					report_data.WinNum += v2.WinNum
					report_data.Win += v2.Win
					report_data.Num += v2.Num
					report_data.BetAll += v2.BetAll
					//代理总计
					report_data1.ValidMember += 1
					report_data1.BetValid += v2.BetValid
					report_data1.Jack += v2.Jack
					report_data1.WinNum += v2.WinNum
					report_data1.Win += v2.Win
					report_data1.Num += v2.Num
					report_data1.BetAll += v2.BetAll
				}
			}
			report_data2[k1] = report_data
		}
		ReportData[k] = report_data1
		reportDataList[k] = report_data2
	}
	//9.获取上一期期数id
	pirods := rebateCountBean.GetLastId(combo)
	var pid int64
	pid = pirods[0].Id
	//10,获取上一期数据
	data := rebateCountBean.GetLastData(pid, combo)
	//处理上期期数数据
	last_data := make(map[int64]back.LastRebate)
	for _, v := range data {
		allFee := back.LastRebate{}
		allFee.BeforeJack = v.BeforeJack
		allFee.BeforeProfit = v.BeforeProfit
		allFee.BeforeBetting = v.BeforeBetting
		allFee.EffectiveMember = v.EffectiveMember
		allFee.BeforeCost = v.BeforeCost
		allFee.NowPayoffElect = v.NowPayoffElect
		allFee.NowPayoffLottery = v.NowPayoffLottery
		allFee.NowPayoffSport = v.NowPayoffSport
		allFee.NowPayoffVideo = v.NowPayoffVideo
		allFee.NowPayoffFish = v.NowPayoffFish
		last_data[v.AgencyId] = allFee
	}
	//11.获取退佣手续费设定
	poundage_list, err := rebateCountBean.GetCharge(combo)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10119, ctx))
	}
	//12.循环代理进行数据处理
	for _, v := range agencylist1 {
		//1.入款手续费
		bank_in := pc.getCash(allfee[v][1].DisBalance, poundage_list[0].IncomPoundageRatio, poundage_list[0].IncomPoundageUp)
		//2.出款手续费
		bank_out := pc.getCash(allfee[v][5].DisBalance, poundage_list[0].OutPoundageRatio, poundage_list[0].OutPoundageUp)
		//3.人工入款手续费
		catm_in_fee := pc.getCash(allfee[v][0].DisBalance, poundage_list[0].IncomPoundageRatio, poundage_list[0].IncomPoundageUp)
		//7.有效会员个数 是否累计
		report_data1 := back.ReportData{}
		if combo.IsCumulative == 1 {
			report_data1.ValidMember = ReportData[v].ValidMember + data[v].EffectiveMember
		} else if combo.IsCumulative == 2 {
			report_data1.ValidMember = ReportData[v].ValidMember
		}
		//赋值report_data1
		report_data1.BeforeBetting = last_data[v].BeforeBetting
		report_data1.BeforeCost = last_data[v].BeforeCost
		report_data1.BeforeJack = last_data[v].BeforeJack
		report_data1.BeforeProfit = last_data[v].BeforeProfit
		report_data1.BankInFee = bank_in
		report_data1.BankOutFee = bank_out
		report_data1.CatmInFee = catm_in_fee
		report_data1.BetAll = ReportData[v].BetAll
		report_data1.Win = ReportData[v].Win
		report_data1.WinNum = ReportData[v].WinNum
		report_data1.Jack = ReportData[v].Jack
		report_data1.BetValid = ReportData[v].BetValid
		report_data1.Num = ReportData[v].Num
		report_data1.EffectiveMember = report_data1.ValidMember
		report_data1.AgencyAccount = ReportData[v].AgencyAccount
		//判断是否退佣条件是否满足
		if (ReportData[v].Win+last_data[v].BeforeProfit) >= float64(rebateset[0].ValidMoney) && ReportData[v].ValidMember >= rebateset[0].EffectiveUser {
			report_data1.IsGrand = 0
		} else {
			report_data1.IsGrand = 1
		}
		//4.jack判断
		if report_data1.Jack < 0 {
			report_data1.Jack = 0
		}
		//5.盈利减去彩金
		report_data1.Win -= report_data1.Jack + last_data[v].BeforeProfit
		//6.总计费用 = 入款手续费+出款手续费+人工入款手续费+返水费用+入款费用+人工入款+其他费用
		total_fee := bank_in + bank_out + catm_in_fee + allfee[v][9].DisBalance + allfee[v][10].DisBalance + allfee[v][0].DisBalance + allfee[v][6].DisBalance
		//其他优惠
		report_data1.OtherFee = allfee[v][6].DisBalance
		//当期费用
		report_data1.NowCost = total_fee
		all_sport_fee := last_data[v].NowPayoffSport + reportDataList[v][5].BetValid     //体育总费用
		all_video_fee := last_data[v].NowPayoffVideo + reportDataList[v][1].BetValid     //视讯总费用
		all_elect_fee := last_data[v].NowPayoffElect + reportDataList[v][2].BetValid     //电子总费用
		all_lottery_fee := last_data[v].NowPayoffLottery + reportDataList[v][4].BetValid //彩票总费用
		all_fish_fee := last_data[v].NowPayoffFish + reportDataList[v][3].BetValid       //捕鱼总费用
		if poundage_list[0].IsDeliveryModel == 1 {
			if all_sport_fee < 0 {
				all_sport_fee = 0
			}
			if all_video_fee < 0 {
				all_video_fee = 0
			}
			if all_elect_fee < 0 {
				all_elect_fee = 0
			}
			if all_lottery_fee < 0 {
				all_lottery_fee = 0
			}
			if all_fish_fee < 0 {
				all_fish_fee = 0
			}
		}
		//退水计算
		report_data1.TotalWater = 0
		/**
		* 1	视讯
		* 2	电子
		* 3	捕鱼
		* 4	彩票
		* 5	体育
		 */
		for k, v := range reportDataList[v] {
			switch k {
			case 1:
				report_data1.TotalWater += WaterList.VideoWater * (v.BetValid + last_data[k].BeforeBetting) * 0.01
			case 2:
				report_data1.TotalWater += WaterList.ElectWater * (v.BetValid + last_data[k].BeforeBetting) * 0.01
			case 3:
				report_data1.TotalWater += WaterList.FishWater * (v.BetValid + last_data[k].BeforeBetting) * 0.01
			case 4:
				report_data1.TotalWater += WaterList.LotteryWater * (v.BetValid + last_data[k].BeforeBetting) * 0.01
			case 5:
				report_data1.TotalWater += WaterList.SportWater * (v.BetValid + last_data[k].BeforeBetting) * 0.01
			}
		}
		//返水费用
		report_data1.DiscountFee = report_data1.TotalWater
		//总费用 上期费用+当期费用+上期奖池+当期奖池
		all_fee := total_fee + last_data[v].BeforeCost + last_data[v].BeforeJack + ReportData[v].Jack
		//视讯退佣处理
		Video_back := pc.Commission(all_video_fee, all_fee, PoundageList.VideoRate, 0, 0)
		//电子退佣处理
		elect_back := pc.Commission(all_elect_fee, Video_back.TotalMoney, PoundageList.ElectRate, Video_back.ReturnCash, Video_back.AddStatus)
		//捕鱼退佣处理
		fish_back := pc.Commission(all_fish_fee, elect_back.TotalMoney, PoundageList.FishRate, elect_back.ReturnCash, elect_back.AddStatus)
		//体育退佣处理
		sports_back := pc.Commission(all_sport_fee, fish_back.TotalMoney, PoundageList.SportRate, fish_back.ReturnCash, fish_back.AddStatus)
		//彩票退佣处理
		lottery_back := pc.Commission(all_lottery_fee, sports_back.TotalMoney, PoundageList.LotteryRate, sports_back.ReturnCash, sports_back.AddStatus)
		//fmt.Println(lottery_back)
		if lottery_back.AddStatus < 0 {
			lottery_back.TotalMoney = 0 - lottery_back.TotalMoney
		}
		//视讯 电子 体育 彩票 捕鱼盈利
		report_data1.NowPayoffElect = elect_back.ReturnCash
		report_data1.NowPayoffVideo = Video_back.ReturnCash
		report_data1.NowPayoffSport = sports_back.ReturnCash
		report_data1.NowPayoffFish = fish_back.ReturnCash
		report_data1.NowPayoffLottery = lottery_back.ReturnCash
		report_data1.TotalBack = Video_back.ReturnCash + elect_back.ReturnCash + fish_back.ReturnCash + sports_back.ReturnCash + lottery_back.ReturnCash - lottery_back.TotalMoney
		report_data1.PeriodsId = combo.PeriodsId
		report_data1.AgencyId = v
		report_data1.SiteId = ReportData[v].SiteId
		report_data1.SiteIndexId = ReportData[v].SiteIndexId
		report_data1.NowJack = ReportData[v].Jack

		//如果退佣为负数 或者未达门槛 则累计下期
		if report_data1.IsGrand == 1 || report_data1.TotalBack < 0 {
			report_data1.TotalBack = 0
			report_data1.Status = 0
		} else {
			report_data1.Status = 1
		}
		report_data1.Rebate = report_data1.TotalBack
		report_data1.RebateWater = report_data1.TotalWater
		//总可退金额
		report_data1.TotalCost = report_data1.TotalBack + report_data1.TotalWater

		if report_data1.TotalBack <= 0 {
			report_data1.TotalCost = 0
		}
		ReportData[v] = report_data1
	}
	//转换json格式
	newdata, err := json.Marshal(ReportData)
	//存入redis
	redis_key := "rebate_" + combo.SiteId + combo.SiteIndexId + strconv.Itoa(combo.PeriodsId)
	err = keySet(redis_key, newdata, "")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(ReportData))
}

//退佣存档
func (pc *RebateCountController) RebateFile(ctx echo.Context) error {
	combo := new(input.RebateFile)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//去除redis中的数据
	redis_key := "rebate_" + combo.SiteId + combo.SiteIndexId + strconv.Itoa(combo.PeriodsId)
	data, err := GetTokenS(redis_key)
	if err == nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	report := make(map[int64]back.RebateListIn)
	json.Unmarshal([]byte(data), &report)
	reportdata := make([]back.RebateListIn, len(report))
	for _, v := range report {
		reportdatalist := back.RebateListIn{}
		reportdatalist.DiscountFee = v.DiscountFee           //返水费用
		reportdatalist.OtherFee = v.OtherFee                 //其它优惠
		reportdatalist.BeforeProfit = v.BeforeProfit         //前期赢利
		reportdatalist.BeforeJack = v.BeforeJack             //前期彩金
		reportdatalist.BeforeCost = v.BeforeCost             //前期费用
		reportdatalist.BeforeBetting = v.BeforeBetting       //前期有效投注
		reportdatalist.NowPayoffElect = v.NowPayoffElect     //当期电子盈利
		reportdatalist.NowCost = v.NowCost                   //当期费用
		reportdatalist.CatmInFee = v.CatmInFee               //手动入款费用
		reportdatalist.BankOutFee = v.BankOutFee             //出款费用
		reportdatalist.BankInFee = v.BankInFee               //入款费用
		reportdatalist.NowPayoffLottery = v.NowPayoffLottery //当期彩票盈利
		reportdatalist.NowPayoffFish = v.NowPayoffFish       //当期捕鱼费用
		reportdatalist.NowPayoffSport = v.NowPayoffSport     //当期体育盈利
		reportdatalist.NowPayoffVideo = v.NowPayoffVideo     //当期视讯盈利
		reportdatalist.SiteIndexId = v.SiteIndexId           //站点前台id
		reportdatalist.SiteId = v.SiteId                     //站点id
		reportdatalist.PeriodsId = v.PeriodsId               //期数id
		reportdatalist.EffectiveMember = v.EffectiveMember   //有效会员
		reportdatalist.Status = v.Status                     //退佣状态
		reportdatalist.AgencyId = v.AgencyId                 //代理id
		reportdatalist.Balance = v.Balance                   //代理余额
		reportdatalist.AgencyAccount = v.AgencyAccount       //代理账号
		reportdatalist.Rebate = v.Rebate                     //本次退佣金额
		reportdatalist.CreateTime = time.Now().Unix()        //创建时间
		reportdatalist.NowBetting = v.NowBetting             //当期有效投注
		reportdatalist.NowJack = v.NowJack                   //本期彩金jack
		reportdatalist.NowProfit = v.NowProfit               //当期赢利
		reportdatalist.RebateWater = v.RebateWater           //本次退水金额
		reportdatalist.Remark = v.Remark                     //备注
		reportdata[v.AgencyId] = reportdatalist
	}
	//fmt.Printf("%+v\n",reportdata)
	num, err := rebateCountBean.RebateFile(reportdata)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(num))
}

//获取站点列表
func (pc *RebateCountController) getSiteList(SiteId string, SiteIndexId string) []back.SiteList {
	list := new(input.GetSiteList)
	list.SiteId = SiteId
	list.SiteIndexId = SiteIndexId
	data, err := periodsBean.GetSiteList(list.SiteId, list.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil
	}
	return data
}

//根据手续费获取出入款金额 money 金额 fee手续费 maxfee最大上限
func (pc *RebateCountController) getCash(Money float64, Fee float64, MaxFee int) float64 {
	data := Money * Fee * 0.01
	if data >= float64(MaxFee) {
		data = float64(MaxFee)
	}
	return data
}

//处理退佣数据
/**
*Money 盈利
*TotalMoney 总费用
*FeeRate 退佣比例
 */
func (pc *RebateCountController) Commission(Money float64, TotalMoney float64, FeeRate float64, ReturnCash float64, AddStatus int64) back.HelperData {
	//剩余盈利=盈利-费用
	data := Money - TotalMoney
	var negative float64
	returndata := back.HelperData{}
	if Money > 0 {
		if AddStatus == 1 {
			//用户当期可获退佣
			ReturnCash += Money * FeeRate * 0.01
		} else {
			if data >= 0 {
				TotalMoney = 0                      //剩余盈利大于0，则费用清0
				ReturnCash += data * FeeRate * 0.01 //用户当期可获退佣
				AddStatus = 1                       //累计标识
			} else {
				TotalMoney = data //剩余盈利小于0 费用等于剩余盈利
				AddStatus = -1    //累计标识
			}
		}
	} else {
		negative = Money * FeeRate * 0.01 //负数退佣
		AddStatus = -1                    //累计标识
	}
	returndata.AddStatus = AddStatus
	returndata.TotalMoney = TotalMoney
	returndata.ReturnCash = ReturnCash
	returndata.Negative = negative
	return returndata
}

//存入redis
func keySet(result string, b []byte, beforeKey string) (err error) {
	if beforeKey != "" {
		//删除旧的key
		err = global.GetRedis().Del(beforeKey).Err()
		//将旧的删除
		err = global.GetRedis().LPop(result).Err()
	}
	//存储新token
	err = global.GetRedis().Set(result, b, 0).Err()
	//将推进list
	err = global.GetRedis().RPush("rebate_info", result).Err()
	return err
}

//获取存储的redis值
func GetTokenS(token string) (string, error) {
	key, err := global.GetRedis().Get(token).Result()
	return key, err
}
