package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type ReportFormBean struct{}

//会员报表统计
//func (*ReportFormBean) MemberReport(this *input.ReportForm, ids []int64, times *global.Times) (list []back.ReportFormBack, err error) {
//	sess := global.GetXorm().NewSession()
//	betReportAccount := new(schema.BetReportAccount)
//	times.Make(betReportAccount.TableName()+".create_time", sess)
//	sess.Where(betReportAccount.TableName()+".site_id = ?", this.SiteId)
//	sess.Where(betReportAccount.TableName()+".site_index_id = ?", this.SiteIndexId)
//	sess.Where(betReportAccount.TableName()+".member_id = ?", this.MemberId)
//	product := new(schema.Product)
//	productType := new(schema.ProductType)
//	sess.NotIn(product.TableName()+".id", ids)
//	sql1 := fmt.Sprintf("%s.v_type = %s.v_type", product.TableName(), betReportAccount.TableName())
//	sql2 := fmt.Sprintf("%s.game_type = %s.id", betReportAccount.TableName(), productType.TableName())
//	sess.Select(betReportAccount.TableName() + ".v_type," + "SUM(" + betReportAccount.TableName() + ".num) as num," +
//		"SUM(" + betReportAccount.TableName() + ".bet_all) as bet_all," + "SUM(" + betReportAccount.TableName() +
//		".bet_valid) as bet_valid," + "SUM(" + betReportAccount.TableName() + ".win) as win," + productType.TableName() + ".title")
//	data := make([]back.ReportForm, 0)
//	err = sess.Table(betReportAccount.TableName()).Join("LEFT", product.TableName(), sql1).
//		Join("LEFT", productType.TableName(), sql2).GroupBy(betReportAccount.TableName() + ".v_type").Find(&data)
//	if err != nil {
//		return
//	}
//	var re back.ReportFormBack
//	for k := range data {
//		re.Project = data[k].VType + data[k].Title
//		re.BetAll = data[k].BetAll
//		re.BetValid = data[k].BetValid
//		re.Num = data[k].Num
//		re.Win = data[k].Win
//		list = append(list, re)
//	}
//	return
//}

//会员报表统计
func (*ReportFormBean) MemberReport(this *input.ReportForm, ids []int64, times *global.Times) (list []back.ReportFormBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	productType := new(schema.ProductType)
	productTypeData := []schema.ProductType{}
	err = sess.Table(productType.TableName()).
		Where("status = ? AND delete_time = ?", 1, 0).Find(&productTypeData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}

	product := new(schema.Product)
	productData := []schema.Product{}
	err = sess.Table(product.TableName()).
		In("id", ids).Find(&productData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}

	betReportAccount := new(schema.BetReportAccount)

	if len(productData) > 0 {
		vTypeList := []string{}
		for _, value := range productData {
			vTypeList = append(vTypeList, value.VType)
		}
		sess.NotIn("v_type", vTypeList)
	}
	times.Make("create_time", sess)
	sess.Where("site_id = ?", this.SiteId)
	sess.Where("site_index_id = ?", this.SiteIndexId)
	sess.Where("member_id = ?", this.MemberId)

	sess.Select("v_type,SUM(num) as num,SUM(bet_all) as bet_all,SUM(bet_valid) as bet_valid,SUM(win) as win, game_type")
	data := []schema.BetReportAccount{}
	err = sess.Table(betReportAccount.TableName()).
		GroupBy("v_type").
		Find(&data)
	if err != nil {
		return
	}

	for _, v := range data {
		var re back.ReportFormBack
		for _, value := range productTypeData {
			if value.Id == int64(v.GameType) {
				re.Project = v.VType + value.Title
			}
		}
		re.BetAll = v.BetAll
		re.BetValid = v.BetValid
		re.Num = int8(v.Num)
		re.Win = v.Win
		list = append(list, re)
	}
	return
}

//站点报表统计
func (*ReportFormBean) GetSiteReportList(this *input.GetDataCenterList, listparam *global.ListParams, times *global.Times) (sdata map[string]interface{}, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteReport := new(schema.SiteReport)

	sdata = make(map[string]interface{})

	if len(this.SiteId) != 0 {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.Account != "" {
		sess.Where("account LIKE ?", "%"+this.Account+"%")
	}
	if len(this.VType) != 0 {
		sess.Where("v_type = ?", this.VType)
	}
	if len(this.AgencyAccount) != 0 {
		sess.Where("agency_account =  LIKE ?", "%"+this.AgencyAccount+"%")
	}

	sess.Where("create_time >= ?", times.StartTime)
	sess.Where("create_time <= ?", times.EndTime)

	//排序
	switch this.SortBy { //	排序类型
	case 1:
		listparam.OrderBy = "num"
	case 2:
		listparam.OrderBy = "win_num"
	case 3:
		listparam.OrderBy = "bet_all"
	case 4:
		listparam.OrderBy = "bet_valid"
	case 5:
		listparam.OrderBy = "jack"
	}
	listparam.Desc = true

	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)

	data := []back.ReportFormDetailBack{}
	err = sess.Table(siteReport.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return sdata, count, err
	}
	sdata["data"] = []back.ReportFormDetailBack{}
	sdata["data"] = data

	//获得符合条件的记录数
	count, err = sess.Table(siteReport.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return sdata, count, err
	}
	//获取总计小计
	total := back.SiteReportTotal{}
	_, err = sess.Table(siteReport.TableName()).
		Where(conds).
		Select("sum(num) num,sum(win_num) win_num,sum(bet_all) bet_all,sum(bet_valid) bet_valid,sum(jack) jack").
		Get(&total)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return sdata, count, err
	}
	sdata["total"] = back.SiteReportTotal{}
	sdata["total"] = total
	//清空total 计算小计
	total = back.SiteReportTotal{}
	for _, v := range data {
		total.Num = total.Num + v.Num
		total.WinNum = total.WinNum + v.WinNum
		total.BetAll = total.BetAll + v.BetAll
		total.BetValid = total.BetValid + v.BetValid
		total.Jack = total.Jack + v.Jack
	}
	sdata["Subtotal"] = back.SiteReportTotal{}
	sdata["Subtotal"] = total
	return
}
