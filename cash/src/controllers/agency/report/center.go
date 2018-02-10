package report

import (
	"controllers"
	"fmt"
	"framework/uuid"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strings"
	"time"
)

type CenterController struct {
	controllers.BaseController
}

//数据中心
func (this *CenterController) GetCenter(ctx echo.Context) error {
	betReportAccount := new(input.BetReportAccount)
	code := global.ValidRequest(betReportAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	listParams := new(global.ListParams)
	this.GetParam(listParams, ctx)

	user := ctx.Get("user").(*global.RedisStruct)
	if len(betReportAccount.SiteId) == 0 {
		betReportAccount.SiteId = user.SiteId
	}
	times := new(global.Times)
	if len(betReportAccount.StartTime) != 0 {
		times.StartTime, code = global.FormatTime2Timestamp2(betReportAccount.StartTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.StartTime = global.GetTimeStart(global.GetCurrentTime())
	}
	if len(betReportAccount.EndTime) != 0 {
		times.EndTime, code = global.FormatTime2Timestamp2(betReportAccount.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.EndTime = global.GetTimeStart(global.GetCurrentTime())
	}
	//获取数据
	data, count, err := reportBean.GetCenter(betReportAccount, listParams, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(data, count))
}

//获取报表查询页面的商品列表
func (this *CenterController) RepSearch(ctx echo.Context) error {
	repSearch := new(input.RepSearch)
	code := global.ValidRequest(repSearch, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	if len(repSearch.SiteId) == 0 {
		repSearch.SiteId = user.SiteId
	}
	// 获取商品数据
	data, err := reportBean.PorductList(repSearch)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	var orderModules []back.OrderModule
	if len(repSearch.SiteIndexId) != 0 {
		productMap := make(map[string]*back.ProductMappingList)
		for k, _ := range data {
			productMap[data[k].VType] = &back.ProductMappingList{
				Name:       data[k].ProductName,
				VType:      data[k].VType,
				PlatformId: data[k].PlatformId,
			}
		}
		f := func(vTypes []string) (mappings []*back.ProductMappingList) {
			for k, _ := range vTypes {
				mapping, ok := productMap[vTypes[k]]
				if ok {
					mappings = append(mappings, mapping)
				} else {
					global.GlobalLogger.Error("err:%s not exist", vTypes[k])
				}
			}
			return
		}
		orderNumbers, err := webInfoBean.GetOrderNumberBySite(repSearch.SiteId, repSearch.SiteIndexId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		modules := strings.Split(orderNumbers.Module, ",")
		for _, v := range modules {
			var orderModule back.OrderModule
			switch v {
			case "video_module":
				orderModule.Module = back.ProductMapping{Name: "视讯直播", VType: "video"}
				orderModule.SubModule = f(strings.Split(orderNumbers.VideoModule, ","))
			case "fc_module":
				orderModule.Module = back.ProductMapping{Name: "彩票游戏", VType: "fc"}
				orderModule.SubModule = f(strings.Split(orderNumbers.FcModule, ","))
			case "dz_module":
				orderModule.Module = back.ProductMapping{Name: "电子游艺", VType: "dz"}
				orderModule.SubModule = f(strings.Split(orderNumbers.DzModule, ","))
			case "sp_module":
				orderModule.Module = back.ProductMapping{Name: "体育赛事", VType: "sp"}
				orderModule.SubModule = f(strings.Split(orderNumbers.SpModule, ","))
			default:
				global.GlobalLogger.Error("%s not in (video_module,fc_module,dz_module,sp_module)", v)
				continue
			}
			orderModules = append(orderModules, orderModule)
		}
	} else {
		for _, v := range data {
			var orderModule back.OrderModule
			arr := strings.Split(v.VType, "_")
			if len(arr) > 1 {
				switch arr[1] {
				case "dz":
					orderModule.Module = back.ProductMapping{"电子游艺", "dz"}
					orderModule.SubModule = append(orderModule.SubModule, &back.ProductMappingList{v.ProductName, v.VType, 0})
				case "fc":
					orderModule.Module = back.ProductMapping{"彩票游戏", "fc"}
					orderModule.SubModule = append(orderModule.SubModule, &back.ProductMappingList{v.ProductName, v.VType, 0})
				case "sp":
					orderModule.Module = back.ProductMapping{"体育赛事", "sp"}
					orderModule.SubModule = append(orderModule.SubModule, &back.ProductMappingList{v.ProductName, v.VType, 0})
				}
			} else {
				orderModule.Module = back.ProductMapping{"视讯直播", "video"}
				orderModule.SubModule = append(orderModule.SubModule, &back.ProductMappingList{v.ProductName, v.VType, 0})
			}
			orderModules = append(orderModules, orderModule)
		}
	}

	return ctx.JSON(200, global.ReplyCollections("data", orderModules))
}

//报表统计数据查询
func (this *CenterController) ReportList(ctx echo.Context) error {
	reportList := new(input.ReportList)
	code := global.ValidRequest(reportList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	if len(reportList.SiteId) == 0 { //代理平台中 当SiteId没有值得时候 给默认值
		reportList.SiteId = user.SiteId
	}

	account, name := "", ""
	switch reportList.Rtype {
	case 2: //查代理报表
		agencyinfo, _, err := thirdAgencyBean.AgencyNameInfo(reportList.Username, reportList.SiteId, reportList.SiteIndexId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if agencyinfo.Id != 0 {
			reportList.AgencyId = agencyinfo.Id
			account = agencyinfo.Account
			name = agencyinfo.Username
		} else {
			return ctx.JSON(500, global.ReplyError(30181, ctx))
		}
	case 3: //查会员报表
		_, err, memberinfo := memberBean.AccountInfo(reportList.Username, reportList.SiteId, reportList.SiteIndexId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if memberinfo.Id != 0 {
			reportList.AgencyId = memberinfo.Id
			account = memberinfo.Account
			name = memberinfo.Realname
		} else {
			return ctx.JSON(500, global.ReplyError(30009, ctx))
		}
	case 1: //默认查站点报表
		account = reportList.SiteId
		name = "站点"
	}
	//获取统计数据
	data, err := reportBean.ReportList(reportList)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	productList, _ := productBean.ProductList()
	for k, v := range data {
		if len(v.VType) != 0 {
			for _, val := range productList {
				if v.VType == val.VType {
					data[k].ProductName = val.ProductName
					break
				}
			}
		} else {
			data[k].ProductName = "总报表"
		}
		data[k].Name = "(" + account + ")" + name
		data[k].Payout = v.BetValid + v.Win
	}
	return ctx.JSON(200, global.ReplyCollections("data", data))
}

//点击金额加载数据
func (this *CenterController) ReportClick(ctx echo.Context) error {
	reportClick := new(input.ReportClick)
	code := global.ValidRequest(reportClick, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	if len(reportClick.SiteId) == 0 {
		reportClick.SiteId = user.SiteId
	}

	//获取数据
	data, err := reportBean.ReportClick(reportClick)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//根据select处理数据
	switch reportClick.Select { //获取数据名称
	case "sh":
	case "ua": // 总代获取代理
		agentIdNameBySite := new(input.SecondIdNameBySite)
		agentIdNameBySite.SiteId = reportClick.SiteId
		agentIdNameBySite.SiteIndexId = reportClick.SiteIndexId
		agentIdNameBySite.FirstId = reportClick.UaId
		agentData, err := agencyBean.GetAllSecondIdName(agentIdNameBySite)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		for k, v := range data {
			for _, val := range agentData {
				if reportClick.Select == "sh" {
					if val.Id == v.UaId {
						data[k].Name = "(" + val.Account + ")" + val.Username
						data[k].Select = "ua"
						break
					}
				} else {
					if val.Id == v.AgencyId {
						data[k].Name = "(" + val.Account + ")" + val.Username
						data[k].Select = "at"
						break
					}
				}
			}
		}
	case "at": // 代理获取会员
		thirdAgencyInfo := new(input.ThirdAgencyInfo)
		thirdAgencyInfo.SiteId = reportClick.SiteId
		thirdAgencyInfo.SiteIndexId = reportClick.SiteIndexId
		thirdAgencyInfo.Id = reportClick.AgencyId
		memberList, err := memberBean.GetAllList(thirdAgencyInfo)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		for k, v := range data {
			for _, val := range memberList {
				if val.Id == v.UaId {
					data[k].Name = "(" + val.Account + ")" + val.Realname
					break
				}
			}
		}
	case "all":
		firstIdNameBySite := new(input.FirstIdNameBySite)
		firstIdNameBySite.SiteId = reportClick.SiteId
		firstIdNameBySite.SiteIndexId = reportClick.SiteIndexId
		//获取站点下的股东列表
		agencyList, err := agencyBean.GetAllFirstIdNameAccount(firstIdNameBySite)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		for k, v := range data {
			for _, val := range agencyList {
				if val.Id == v.UaId {
					data[k].Name = "(" + val.Account + ")" + val.Username
					break
				}
			}
		}
		break
	default:
		firstIdNameBySite := new(input.FirstIdNameBySite)
		firstIdNameBySite.SiteId = reportClick.SiteId
		firstIdNameBySite.SiteIndexId = reportClick.SiteIndexId
		//获取站点下的股东列表
		agencyList, err := agencyBean.GetAllFirstIdNameAccount(firstIdNameBySite)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		for k, v := range data {
			for _, val := range agencyList {
				if val.Id == v.UaId {
					data[k].Name = "(" + val.Account + ")" + val.Username
					break
				}
			}
		}
		break
	}
	for k, v := range data {
		data[k].Payout = v.BetValid + v.Win
	}

	return ctx.JSON(200, global.ReplyCollections("data", data))
}

//账单查询
func (this *CenterController) ReportBills(ctx echo.Context) error {
	billList := new(input.BillList)
	code := global.ValidRequest(billList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	user := ctx.Get("user").(*global.RedisStruct)
	if len(billList.SiteId) == 0 {
		billList.SiteId = user.SiteId
	}

	//查询账单
	data, err := reportBean.GetSiteReportBill(billList)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("data", data))
}

func (*CenterController) Periods(ctx echo.Context) error {
	periods := make(map[string][]string)
	Year := time.Now().Format("2006")
	lastYear := time.Now().AddDate(-1, 0, 0).Format("2006")
	periods[Year] = PeriodsList(Year)
	periods[lastYear] = PeriodsList(lastYear)
	return ctx.JSON(200, global.ReplyCollections("data", periods))
}

//获取当前期数的第一天和最后一天
func getYearAndMonth(dd time.Time) (start time.Time, end time.Time) {
	year, month, _ := dd.Date()
	loc := dd.Location()
	//获取每月的第一天和最后一天
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	//获取每个月的第一个星期一和最后一个星期日
	startData := startOfMonth.AddDate(0, 0, (8-WeekDayNum(startOfMonth.Weekday().String()))%7)
	endData := endOfMonth.AddDate(0, 0, (7-WeekDayNum(endOfMonth.Weekday().String()))%7)
	return startData, endData
}
func WeekDayNum(str string) (num int) {
	switch str {
	case "Monday":
		return 1
	case "Tuesday":
		return 2
	case "Wednesday":
		return 3
	case "Thursday":
		return 4
	case "Friday":
		return 5
	case "Saturday":
		return 6
	case "Sunday":
		return 7
	}
	return
}

//获取该年份的期数
func PeriodsList(year string) (strs []string) {
	for i := 1; i <= 12; i++ {
		var dd time.Time
		if i < 10 {
			dd, _ = time.Parse("2006-01-02", fmt.Sprintf("%s-0%d-01", year, i))
		} else {
			dd, _ = time.Parse("2006-01-02", fmt.Sprintf("%s-%d-01", year, i))
		}
		startData, endData := getYearAndMonth(dd)
		strs = append(strs, fmt.Sprintf("%s~%s", startData, endData))
	}
	return
}

//账单缴款数据查询
func (*CenterController) PressMoneySiteInfo(ctx echo.Context) error {
	user := ctx.Get("user").(*global.RedisStruct)
	data, has, err := pressMoneyBean.PressMoneySiteInfo(user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has == false {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//站点业主提交缴款
func (*CenterController) PutSiteUpdateStatus(ctx echo.Context) error {
	pressMoneyStatus := new(input.PressMoneyStatus)
	code := global.ValidRequest(pressMoneyStatus, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	count, err := pressMoneyBean.PressIdInfo(user.SiteId, pressMoneyStatus.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(70031, ctx))
	}
	//账单确认
	count, err = pressMoneyBean.UpdateStatus(pressMoneyStatus)
	if err != nil || count < 1 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

func (*CenterController) PrePayment(ctx echo.Context) error {
	prePayment := new(input.PrePayment)
	code := global.ValidRequest(prePayment, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//if prePayment.PayId <= 0 {
	//	return ctx.JSON(200, global.ReplyError(60019, ctx))
	//}
	user := ctx.Get("user").(*global.RedisStruct)
	prePayment.OrderNum = uuid.NewV4().String()
	prePayment.SiteId = user.SiteId
	prePayment.SiteIndexId = user.SiteIndexId
	prePayment.AdminUser = user.Account
	count, err := pressMoneyBean.AddPrePayment(prePayment)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(70032, ctx))
	}
	return ctx.NoContent(204)
}
