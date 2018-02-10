package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strconv"
	"strings"
	"time"
)

type ReportBean struct{}

//数据统计
//func (*ReportBean) GetCenter(this *input.BetReportAccount, listParams *global.ListParams, times *global.Times) (
//	back.BetReportAccountAllBack, int64, error) {
//	sess := global.GetXorm().NewSession()
//	defer sess.Close()
//	var info back.BetReportAccountAllBack
//	var data []back.BetReportAccount
//	report_account := new(schema.BetReportAccount)
//	//联表
//	agency := new(schema.Agency)
//	sql := fmt.Sprintf("%s.agency_id = %s.id", report_account.TableName(), agency.TableName())
//	sess.Join("LEFT", agency.TableName(), sql)
//
//	//where条件
//	if len(this.SiteId) != 0 { //站点
//		sess.Where(report_account.TableName()+".site_id = ?", this.SiteId)
//	}
//	if len(this.SiteIndexId) != 0 { //站点前台
//		sess.Where(report_account.TableName()+".site_index_id = ?", this.SiteIndexId)
//	}
//	if len(this.VType) != 0 { //游戏标识 如bbin,bbin_dz,bbin_fc
//		sess.Where(report_account.TableName()+".v_type = ?", this.VType)
//	}
//	if len(this.UserAccount) != 0 { //账号
//		sess.Where(report_account.TableName()+".account like ?", "%"+this.UserAccount+"%")
//	}
//	if len(this.Account) != 0 { //代理id
//		sess.Where(agency.TableName()+".account like ?", "%"+this.Account+"%")
//	}
//	//开始结束时间
//	sess.Where(report_account.TableName()+".create_time>=?", times.StartTime)
//	sess.Where(report_account.TableName()+".create_time<=?", times.EndTime)
//
//	//排序
//	switch this.SortBy { //	排序类型
//	case 1:
//		listParams.OrderBy = report_account.TableName() + ".account"
//	case 2:
//		listParams.OrderBy = report_account.TableName() + ".num"
//	case 3:
//		listParams.OrderBy = report_account.TableName() + ".win_num"
//	case 4:
//		listParams.OrderBy = report_account.TableName() + ".bet_all"
//	case 5:
//		listParams.OrderBy = report_account.TableName() + ".bet_valid"
//	case 6:
//		listParams.OrderBy = report_account.TableName() + ".create_time"
//	}
//	switch this.Sort { //排序方式	1,账号2,总笔数3,赢笔数4,投注额度5,有效投注6,统计时间
//	case 1:
//		listParams.Desc = true
//	default:
//		listParams.Desc = false
//	}
//
//	//分页
//	if this.PageSize == 50 || this.PageSize == 100 || this.PageSize == 200 {
//		listParams.PageSize = this.PageSize
//	}
//	conds := sess.Conds()
//	listParams.Make(sess)
//
//	//查询字段
//	sess.Select(report_account.TableName() + ".*")
//
//	//查询
//	err := sess.Table(report_account.TableName()).Find(&data)
//	if err != nil {
//		global.GlobalLogger.Error("error:%s", err.Error())
//		return info, 0, err
//	}
//	//获得符合条件的记录数
//	count, err := sess.Table(report_account.TableName()).Where(conds).Count()
//	err = sess.Table(report_account.TableName()).Find(&data)
//	if err != nil {
//		global.GlobalLogger.Error("error:%s", err.Error())
//		return info, count, err
//	}
//	//获得符合条件的记录数
//	count, err = sess.Table(report_account.TableName()).Where(conds).Count()
//	if err != nil {
//		global.GlobalLogger.Error("error:%s", err.Error())
//		return info, count, err
//	}
//	//查询总计
//	sess.Select("SUM(sales_bet_report_account.win),SUM(sales_bet_report_account.num)," +
//		"SUM(sales_bet_report_account.win_num),SUM(sales_bet_report_account.bet_all)," +
//		"SUM(sales_bet_report_account.bet_valid),SUM(sales_bet_report_account.jack)")
//	var rt back.BetReportAccountTotal
//	sq2 := fmt.Sprintf("%s.agency_id = %s.id", report_account.TableName(), agency.TableName())
//	sess.Join("LEFT", agency.TableName(), sq2)
//	_, err = sess.Table(report_account.TableName()).Where(conds).Get(&rt)
//	if err != nil {
//		global.GlobalLogger.Error("error:%s", err.Error())
//		return info, count, err
//	}
//	rt.PcTotal = 0
//	if len(data) > 0 {
//		for _, v := range data {
//			rt.SmallWin = rt.SmallWin + v.Win
//			rt.SmallBetAll = rt.SmallBetAll + v.BetAll
//			rt.SmallBetValid = rt.SmallBetValid + v.BetValid
//			rt.SmallJack = rt.SmallJack + v.Jack
//			rt.SmallNum = rt.SmallNum + v.Num
//			rt.SmallWinNum = rt.SmallWinNum + v.WinNum
//			rt.SmallPc = rt.SmallPc + v.Pc
//		}
//	}
//	info.BetReportAccount = data
//	info.BetReportAccountTotal = rt
//	return info, count, err
//}

//数据统计
func (*ReportBean) GetCenter(this *input.BetReportAccount, listParams *global.ListParams, times *global.Times) (
	info back.BetReportAccountAllBack, count int64, err error) {
	fmt.Println(1111)
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	var agencyId int64
	if len(this.Account) != 0 {

		fmt.Println(222)
		//代理账号不为空
		agency := new(schema.Agency)
		_, err = sess.Table(agency.TableName()).
			Where("account like ? AND site_id = ? AND site_index_id = ?", "%"+this.VType+"%", this.SiteId, this.SiteIndexId).Get(agency)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return
		}
		agencyId = agency.Id
	}

	betReportAccount := new(schema.BetReportAccount)
	if len(this.SiteId) != 0 {
		sess.Where("site_id = ?", this.SiteId)
	}
	if len(this.SiteIndexId) != 0 {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if len(this.VType) != 0 {
		sess.Where("v_type = ?", this.VType)
	}
	if len(this.UserAccount) != 0 {
		sess.Where("account like ?", "%"+this.VType+"%")
	}
	if agencyId != 0 {
		sess.Where("agency_id = ?", agencyId)
	}
	//开始结束时间
	sess.Where("create_time>=?", times.StartTime)
	sess.Where("create_time<=?", times.EndTime)

	//排序
	switch this.SortBy { //	排序类型
	case 1:
		listParams.OrderBy = "account"
	case 2:
		listParams.OrderBy = "num"
	case 3:
		listParams.OrderBy = "win_num"
	case 4:
		listParams.OrderBy = "bet_all"
	case 5:
		listParams.OrderBy = "bet_valid"
	case 6:
		listParams.OrderBy = "create_time"
	}
	switch this.Sort { //排序方式	1,账号2,总笔数3,赢笔数4,投注额度5,有效投注6,统计时间
	case 1:
		listParams.Desc = true
	default:
		listParams.Desc = false
	}

	//分页
	if this.PageSize == 50 || this.PageSize == 100 || this.PageSize == 200 {
		listParams.PageSize = this.PageSize
	}
	conds := sess.Conds()
	listParams.Make(sess)
	//查询
	sess.Select("id,site_id,site_index_id,agency_id,member_id,account,win,num,win_num,bet_all,bet_valid,day_time,jack")
	err = sess.Table(betReportAccount.TableName()).Find(&info.BetReportAccount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	//获得符合条件的记录数
	count, err = sess.Table(betReportAccount.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, count, err
	}
	//查询总计
	sess.Select("SUM(win),SUM(num),SUM(win_num),SUM(bet_all),SUM(bet_valid),SUM(jack)")
	_, err = sess.Table(betReportAccount.TableName()).Where(conds).Get(&info.BetReportAccountTotal)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, count, err
	}
	//总派彩
	info.BetReportAccountTotal.PcTotal = 0

	if len(info.BetReportAccount) > 0 {
		for _, v := range info.BetReportAccount {
			info.BetReportAccountTotal.SmallWin = info.BetReportAccountTotal.SmallWin + v.Win
			info.BetReportAccountTotal.SmallBetAll = info.BetReportAccountTotal.SmallBetAll + v.BetAll
			info.BetReportAccountTotal.SmallBetValid = info.BetReportAccountTotal.SmallBetValid + v.BetValid
			info.BetReportAccountTotal.SmallJack = info.BetReportAccountTotal.SmallJack + v.Jack
			info.BetReportAccountTotal.SmallNum = info.BetReportAccountTotal.SmallNum + v.Num
			info.BetReportAccountTotal.SmallWinNum = info.BetReportAccountTotal.SmallWinNum + v.WinNum
			info.BetReportAccountTotal.SmallPc = info.BetReportAccountTotal.SmallPc + v.Pc
		}
	}
	return
}

//报表列表数据输出
func (*ReportBean) ReportList(this *input.ReportList) ([]back.ReportList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	reportAccount := new(schema.BetReportAccount)
	var data []back.ReportList
	var numsql string //select 查询字段
	//where条件
	if len(this.SiteId) != 0 { //站点
		sess.Where("site_id = ?", this.SiteId)
	}

	switch this.Rtype { //报表类型 1.代理报表 2.会员报表 0.总报表
	case 2:
		//报表类型 代理 根据代理id查询
		sess.Where("agency_id = ?", this.AgencyId) //代理id需要在controllers中获取
	case 3:
		//根据会员查询
		sess.Where("account link ?", "%"+this.Username)
	}

	num := len(this.VType)
	if num > 0 {
		sql := ""
		for i := 0; i < num; i++ { //商品信息
			arr := strings.Split(this.VType[i], "_")
			var gameTypeId int64
			if len(arr) > 1 {
				gameTypeId = gameTypeInt(arr[1])
			} else {
				gameTypeId = 1
			}
			if i == 0 {
				sql = sql + "'" + arr[0] + strconv.FormatInt(gameTypeId, 10) + "'"
			}
			sql = sql + ", '" + arr[0] + strconv.FormatInt(gameTypeId, 10) + "'"
		}
		sess.Where("(CONCAT(`platform`,`game_type`)) IN (" + sql + ")")
		//sess.In("v_type",this.VType)
	}

	if len(this.SiteIndexId) != 0 { //前台
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}

	locStr := ""
	if this.TimeZone == 2 {
		locStr = "UTC" //世间时间
	} else {
		locStr = "Local" //北京时间
	}
	loc, _ := time.LoadLocation(locStr) //时区
	if this.StartTime != "" {           //开始时间
		t, err := time.ParseInLocation("2006-01-02", this.StartTime, loc)
		if err == nil {
			sess.Where("create_time >= ?", t.Unix())
		}
	} else { //起始时间为空的时候 默认查询当天
		StartTime := time.Now().Format("2006-01-02")
		t, err := time.ParseInLocation("2006-01-02", StartTime, loc)
		if err == nil {
			sess.Where("create_time >= ?", t.Unix())
		}
	}
	this.EndTime = this.EndTime + " 23:59:59" //获取日期的结束时间
	if this.EndTime != "" {                   //结束时间
		t, err := time.ParseInLocation("2006-01-02 15:04:05", this.EndTime, loc)
		if err == nil {
			sess.Where("create_time<=?", t.Unix())
		}
	}

	sess.GroupBy("platform,game_type,site_id")
	numsql = numsql + " sum(num) num,sum(win_num) win_num,sum(bet_all) bet_all,sum(bet_valid) bet_valid,sum(win) win, sum(jack) jack,"
	numsql = numsql + "site_id, site_index_id, agency_id, ua_id, sh_id, account, platform, game_type"
	err := sess.Table(reportAccount.TableName()).Select(numsql).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	var totalReport = back.ReportList{}
	for k, v := range data { //拼接商品v_type 拼接总报表数据
		data[k].VType = v.Platform + gameTypeStr(v.GameType)
		totalReport.SiteId = v.SiteId
		totalReport.SiteIndexId = v.SiteIndexId
		totalReport.ShId = v.ShId
		totalReport.UaId = v.UaId
		totalReport.AgencyId = v.AgencyId
		totalReport.Account = v.Account

		totalReport.Num = totalReport.Num + v.Num
		totalReport.WinNum = totalReport.WinNum + v.WinNum
		totalReport.BetAll = totalReport.BetAll + v.BetAll
		totalReport.BetValid = totalReport.BetValid + v.BetValid
		totalReport.Win = totalReport.Win + v.Win
		totalReport.Jack = totalReport.Jack + v.Jack
	}
	data = append(data, totalReport)
	return data, err
}

func (*ReportBean) ReportClick(this *input.ReportClick) ([]back.ReportClick, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.ReportClick
	reportAccount := new(schema.BetReportAccount)

	if len(this.SiteId) != 0 {
		sess.Where("site_id = ?", this.SiteId)
	}
	if len(this.SiteIndexId) != 0 {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}

	switch this.Select { //报表层级(all:站点 sh:股东 ua:总代理 at:代理)
	case "sh": //点股东查总代
		sess.Where("sh_id = ?", this.ShId)
		sess.GroupBy("ua_id")
	case "ua": //点总代查代理
		sess.Where("sh_id = ? AND ua_id = ?", this.ShId, this.UaId)
		sess.GroupBy("agency_id")
	case "at": //点代理查会员
		sess.Where("sh_id = ? AND ua_id = ? AND agency_id = ?", this.ShId, this.UaId, this.AgencyId)
		sess.GroupBy("account")
	case "all": //默认点站点查股东
		sess.GroupBy("sh_id")
	default:
		sess.GroupBy("sh_id")
	}

	if len(this.VType) != 0 {
		arr := strings.Split(this.VType, "_")
		gameTypeId := gameTypeInt(arr[1])
		sess.Where("platform = ? and game_type = ?", arr[0], gameTypeId)
	}

	loc, _ := time.LoadLocation("Local") //时区
	if this.StartTime != "" {            //开始时间
		t, err := time.ParseInLocation("2006-01-02", this.StartTime, loc)
		if err == nil {
			sess.Where(reportAccount.TableName()+".create_time>=?", t.Unix())
		}
	} else { //起始时间为空的时候 默认查询当天
		StartTime := time.Now().Format("2006-01-02")
		t, err := time.ParseInLocation("2006-01-02", StartTime, loc)
		if err == nil {
			sess.Where("create_time>=?", t.Unix())
		}
	}
	this.EndTime = this.EndTime + " 23:59:59" //获取日期的结束时间
	if this.EndTime != "" {                   //结束时间
		t, err := time.ParseInLocation("2006-01-02 15:04:05", this.EndTime, loc)
		if err == nil {
			sess.Where("create_time<=?", t.Unix())
		}
	}

	selectSql := fmt.Sprintf(" sum(num) num,sum(win_num) win_num,sum(bet_all) bet_all,sum(bet_valid) bet_valid,sum(win) win, sum(jack) jack")
	selectSql = selectSql + fmt.Sprintf("site_id, site_index_id,sh_id, ua_id, agency_id, account")
	//查询
	err := sess.Table(reportAccount.TableName()).Select(selectSql).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取报表查询页面的商品列表
func (*ReportBean) PorductList(this *input.RepSearch) ([]back.ProductlistRep, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	var data []back.ProductlistRep
	ses := global.GetXorm().NewSession()
	productDel := new(schema.SiteProductDel)
	delData := []back.SiteProductDel{}

	productIdDl := []int64{} //剔除的商品id集合
	if len(this.SiteId) != 0 {
		ses.Where("site_id = ?", this.SiteId)
		if len(this.SiteIndexId) == 0 { //根据site_index_id来查询剔除的商品
			indexArr, _ := OneSiteId(this.SiteId)
			ses.GroupBy("product_id, site_id")
			err := ses.Table(productDel.TableName()).
				Select("product_id,COUNT(site_index_id) count, site_id").Find(&delData)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, err
			}
			for _, val := range delData {
				if val.Count == indexArr {
					productIdDl = append(productIdDl, val.ProductId)
				}
			}
		} else {
			ses.GroupBy("product_id, site_id, site_index_id")
			ses.Where("site_index_id = ?", this.SiteIndexId)
			err := ses.Table(productDel.TableName()).Select("product_id,COUNT(site_index_id) count, site_id").Find(&delData)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, err
			}
			for _, val := range delData {
				productIdDl = append(productIdDl, val.ProductId)
			}
		}

		if len(productIdDl) > 0 { //剔除
			sess.NotIn("id", productIdDl)
		}
	}

	//查询商品列表
	err := sess.Table(product.TableName()).Where("status = 1").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//根据site_id查询前台数量
func OneSiteId(SiteId string) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	count, err = sess.Table(site.TableName()).Where("id = ?", SiteId).Count()
	return
}

func gameTypeInt(str string) (id int64) {
	switch {
	case str == "dz":
		return 2
	case str == "by":
		return 3
	case str == "fc":
		return 4
	case str == "sp":
		return 5
	default:
		return 1
	}
	return
}
func gameTypeStr(id int64) (str string) {
	switch {
	case id == 1:
		return ""
	case id == 2:
		return "dz"
	case id == 3:
		return "by"
	case id == 4:
		return "fc"
	case id == 5:
		return "sp"
	}
	return
}

//查询
func (*ReportBean) ReportExport(this *input.ReportBills) (data []schema.SiteReport, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	siteReport := new(schema.SiteReport)

	if len(this.SiteId) != 0 {
		sess.In("site_id", this.SiteId)
	}
	if len(this.VType) != 0 {
		sess.In("v_type", this.VType)
	}

	loc, _ := time.LoadLocation("Local") //时区
	if this.StartTime != "" {            //开始时间
		t, err := time.ParseInLocation("2006-01-02", this.StartTime, loc)
		if err == nil {
			sess.Where("create_time>=?", t.Unix())
		}
	} else { //起始时间为空的时候 默认查询当天
		StartTime := time.Now().Format("2006-01-02")
		t, err := time.ParseInLocation("2006-01-02", StartTime, loc)
		if err == nil {
			sess.Where("create_time>=?", t.Unix())
		}
	}
	this.EndTime = this.EndTime + " 23:59:59" //获取日期的结束时间
	if this.EndTime != "" {                   //结束时间
		t, err := time.ParseInLocation("2006-01-02 15:04:05", this.EndTime, loc)
		if err == nil {
			sess.Where("create_time<=?", t.Unix())
		}
	}
	sess.GroupBy("site_id, v_type")
	sess.Select("site_id, v_type, sum(win) win,sum(jack)")
	err = sess.Table(siteReport.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//账单添加
func (*ReportBean) BillsAdd(this *input.BillsAdd, startDate, endDate int64) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteReportBill)
	admin.ReportData = this.ReportData
	admin.SiteId = this.SiteId
	admin.Qishu = this.Year + "-" + this.Qishu
	admin.Status = 2
	admin.StartDate = startDate
	admin.EndDate = endDate
	count, err = sess.Table(admin.TableName()).InsertOne(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//账单查询
func (*ReportBean) GetSiteReportBill(this *input.BillList) ([]schema.SiteReportBill, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteReportBill := new(schema.SiteReportBill)
	var data []schema.SiteReportBill
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if len(this.Year) != 0 && len(this.Qishu) != 0 {
		sess.Where("qishu = ?", this.Year+"-"+this.Qishu)
	}
	if this.Status != 0 {
		sess.Where("status = ?", this.Status)
	}
	err := sess.Table(siteReportBill.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//账单批量下发
func (*ReportBean) ReportBillBatch(this *input.BillListBatch) (int64, error) {
	sess := global.GetXorm().NewSession()
	sess.Close()
	rE := new(schema.SiteReportBill)
	rE.Status = 2
	var count int64
	var err error
	if len(this.Id) > 0 {
		ids := strings.Split(this.Id, ",")
		count, err = sess.In("id", ids).Update(rE)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
		return count, err
	}
	return count, err
}

//账单修改
func (*ReportBean) PutSiteBillUpdate(this *input.SiteBillUpdate) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteReportBill := new(schema.SiteReportBill)
	has, err := sess.Table(siteReportBill.TableName()).
		Where("id=?", this.Id).Get(siteReportBill)
	if !has {
		code = 60212
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if this.Status != 0 {
		siteReportBill.Status = this.Status
	} else {
		siteReportBill.ReportData = this.ReportData
	}
	sess.Where("site_id = ? and qishu = ?", this.SiteId, this.Qishu)
	count, err = sess.Where("id = ?", this.Id).Cols("report_data, status").Update(siteReportBill)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//账单删除
func (*ReportBean) DelSiteBill(this *input.BillList) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteReportBill := new(schema.SiteReportBill)
	has, err := sess.Table(siteReportBill.TableName()).
		Where("site_id=? and qishu = ?", this.SiteId, this.Year+"-"+this.Qishu).Get(siteReportBill)
	if !has {
		code = 60212
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	siteReportBill.Status = this.Status
	sess.Where("site_id = ?", this.SiteId)
	sess.Where("qishu = ?", this.Year+"-"+this.Qishu)
	count, err = sess.Cols("status").Update(siteReportBill)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}
