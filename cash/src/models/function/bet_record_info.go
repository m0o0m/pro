package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type BetRecordInfoBean struct{}

func (BetRecordInfoBean) GetBetRecordList(this *input.BetRecordList, listParams *global.ListParams, times *global.Times) (data []back.BetRecordList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bet_record := new(schema.BetRecordInfo)
	if len(this.SiteIndexId) != 0 {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if len(this.Platform) != 0 { //视讯内部的平台
		sess.Where("platform = ?", this.Platform)
	}
	if this.GameType != 0 { //游戏类型，1视讯  2电子 3捕鱼 4彩票 5体育
		sess.Where("game_type = ?", this.GameType)
	}
	if len(this.Account) != 0 {
		sess.Where("account like ?", "%"+this.Account+"%")
	}
	if len(this.OrderId) != 0 {
		sess.Where("order_id = ?", this.OrderId)
	}
	if len(this.GameName) != 0 {
		sess.Where("game_name like ?", "%"+this.GameName+"%")
	}
	sess.Where("settle_timeline>=?", times.StartTime)
	sess.Where("settle_timeline<=?", times.EndTime)
	if len(this.SiteId) != 0 {
		sess.Where("site_id=?", this.SiteId)
	}
	//1,账号2,注单号3,注单时间4,结算时间5,下注金额6,有效投注
	switch this.SortType {
	case 1:
		listParams.OrderBy = "account"
		break
	case 2:
		listParams.OrderBy = "order_id"
		break
	case 3:
		listParams.OrderBy = "bet_time"
		break
	case 4:
		listParams.OrderBy = "settle_time"
		break
	case 5:
		listParams.OrderBy = "bet_all"
		break
	case 6:
		listParams.OrderBy = "bet_yx"
		break
	}
	switch this.Sort { //排序
	case 1:
		listParams.Desc = true
	default:
		listParams.Desc = false
	}
	if this.PageSize == 50 || this.PageSize == 100 || this.PageSize == 200 {
		listParams.PageSize = this.PageSize
	} else {
		listParams.PageSize = 50
	}
	conds := sess.Conds()
	count, err = sess.Table(bet_record.TableName()).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	if count == 0 {
		return
	}
	listParams.Make(sess)
	err = sess.Table(bet_record.TableName()).Where(conds).Select("*").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取会员交易记录
func (*BetRecordInfoBean) GetTransactionRecord(this *input.TransactionRecord, times *global.Times, listParams *global.ListParams) (data []back.TransactionRecord, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	betRecordInfo := new(schema.BetRecordInfo)
	//判断并组合where条件
	if this.OrderId != "" {
		sess.Where("order_id = ?", this.OrderId)
	}
	if this.GameType != 0 {
		sess.Where("game_type = ?", this.GameType)
	}
	//if this.PlatformId != 0 {
	//	sess.Where("platform_id = ?", this.PlatformId)
	//}
	if this.MemberId != 0 {
		sess.Where("member_id = ?", this.MemberId)
	}
	//分页查询
	listParams.Make(sess)
	//根据时间段查询
	times.Make("bet_timeline", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err = sess.Table(betRecordInfo.TableName()).Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(betRecordInfo.TableName()).Where(conds).Count()
	return

}

//获取会员现金流水
func (*BetRecordInfoBean) GetMemberCashRecord(this *input.MemberCashRecords, times *global.Times, listParams *global.ListParams) (data []back.MemberCashRecords, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberCashRecord := new(schema.MemberCashRecord)
	//判断并组合where条件
	switch this.SourceType {
	case 1: //存款：线上存款、公司存款、人工存入
		sess.In("source_type", []int8{1, 2, 11})
	case 2: //取款：线上取款
		sess.Where("source_type=?", 4)
	case 3: //额度转换：系统平台与其他游戏平台之间转换流水。
		sess.Where("source_type=?", 8)
	case 4: //其他：人工取出（备注人工取出原因）、优惠活动（备注具体活动信息）
		sess.In("source_type", []int8{5, 6})
	}
	if this.MemberId != 0 {
		sess.Where("member_id = ?", this.MemberId)
	}
	//分页查询
	listParams.Make(sess)
	//根据时间段查询
	times.Make("create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err = sess.Table(memberCashRecord.TableName()).Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(memberCashRecord.TableName()).Where(conds).Count()
	return
}

//获取商品类型
func (*BetRecordInfoBean) GetProductType() (pt []back.ProductTypeList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	productType := new(schema.ProductType)
	err = sess.Table(productType.TableName()).Where("status=?", 1).Where("delete_time=?", 0).Find(&pt)
	return
}

//获取商品类型下的商品
func (*BetRecordInfoBean) GetProduct(this *input.ProductTypeId) (pn []back.ProductName, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	sess.Where("status=?", 1).Where("delete_time=?", 0)
	err = sess.Table(product.TableName()).Where("type_id=?", this.TypeId).Find(&pn)
	return
}

//wap 投注记录列表
func (*BetRecordInfoBean) WapBetRecordList(this *input.WapBetRecord, listParams *global.ListParams, times *global.Times) (
	[]back.WapBetRecord, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var betRecordList []back.WapBetRecord
	betRecordInfo := new(schema.BetRecordInfo)
	sess.Where("game_type=?", this.GameType).Where("member_id=?", this.MemberId)
	listParams.Make(sess)
	times.Make("bet_timeline", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(betRecordInfo.TableName()).Find(&betRecordList)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return betRecordList, 0, err
	}
	count, err := sess.Table(betRecordInfo.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return betRecordList, count, err
	}
	return betRecordList, count, err
}

//wap 现金流水列表
func (*BetRecordInfoBean) WapMemberCashRecord(this *input.WapMemberCashRecords, listParams *global.ListParams, times *global.Times) (
	[]back.WapMemberCashRecord, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberCashRecord := new(schema.MemberCashRecord)
	var memberCashRecordList []back.WapMemberCashRecord
	//判断并组合where条件
	switch this.SourceType {
	case 1: //存款：线上存款、公司存款、人工存入
		sess.In("source_type", []int8{1, 2, 11})
	case 2: //取款：线上取款
		sess.Where("source_type=?", 4)
	case 3: //额度转换：系统平台与其他游戏平台之间转换流水。
		sess.Where("source_type=?", 8)
	case 4: //其他：人工取出（备注人工取出原因）、优惠活动（备注具体活动信息）
		sess.In("source_type", []int8{5, 6})
	}
	sess.Where("member_id = ?", this.MemberId)
	//分页查询
	listParams.Make(sess)
	times.Make("create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(memberCashRecord.TableName()).Find(&memberCashRecordList)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberCashRecordList, 0, err
	}
	count, err := sess.Table(memberCashRecord.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberCashRecordList, count, err
	}
	return memberCashRecordList, count, err
}

//报表统计
func (*BetRecordInfoBean) WapReportStatistics(this *input.WapReport, times *global.Times) (lists []back.WapReportStatisticsBacks, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	betRecordInfo := new(schema.BetRecordInfo)
	//时间段查询语句
	times.Make("bet_timeline", sess)
	sess.Where("member_id=?", this.MemberId)
	conds := sess.Conds()
	//查询每天的总计
	reportStatistics := make([]back.WapReportStatistics, 0)
	err = sess.Table(betRecordInfo.TableName()).Select("LEFT (`bet_time`, 10) AS date_time," +
		"count(*) AS bet_count,SUM(bet_all) AS bet_all,SUM(bet_yx) AS bet_yx,SUM(win) AS win").
		GroupBy("date_time").Find(&reportStatistics)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return lists, err
	}
	count, err := sess.Table(betRecordInfo.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return lists, err
	}
	//查询每天的内容详情
	reportStatisticsInfo := make([]back.WapReportStatisticsInfo, 0)
	err = sess.Table(betRecordInfo.TableName()).Select("LEFT (`bet_time`, 10) AS date_time,bet_all,bet_yx," +
		"win,platform,v_type").
		Where(conds).Find(&reportStatisticsInfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return lists, err
	}
	var data back.WapReportStatistics
	var datas []back.WapReportStatistics
	var list back.WapReportStatisticsBacks
	list.WeekBetCount = count
	//计算出时间
	days := global.GetEveryDay(time.Unix(times.StartTime, 0), time.Unix(times.EndTime, 0))
	for k := range reportStatistics {
		for i := range days {
			if reportStatistics[k].DateTime == days[i].Format("2006-01-02") {
				data.BetAll = reportStatistics[k].BetAll
				data.BetYx = reportStatistics[k].BetYx
				data.Win = reportStatistics[k].Win
				data.DateTime = reportStatistics[k].DateTime
				data.BetCount = reportStatistics[k].BetCount
				datas = append(datas, data)
			}
		}
	}
	//给返回的数据赋值
	for k := range datas {
		list.WapReportStatisticsInfo = reportStatisticsInfo
		list.WapReportStatisticsBack = datas
		list.WeekWin += datas[k].Win
		list.WeekBetAll += datas[k].BetAll
		list.WeekBetYx += datas[k].BetYx
	}
	lists = append(lists, list)
	return lists, err
}

//wap获取前台会员交易记录
func (c *BetRecordInfoBean) GetMemberRecord(siteId, siteIndexId, Account string, this *input.RecordInfoList, listparams *global.ListParams) (data []back.WapRecordList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	betRecordInfo := new(schema.BetRecordInfo)
	if this.StartTime != 0 {
		sess.Where("bet_timeline>=?", this.StartTime)
	}

	if this.EndTime != 0 {
		sess.Where("bet_timeline<=?", this.EndTime)
	}
	if this.VType != 0 {
		sess.Where("game_type = ?", this.VType)
	}
	if this.OrderNum != 0 {
		sess.Where("order_id =?", this.OrderNum)
	}
	if this.GameResult != "" {
		sess.Where("game_result =?", this.GameResult)
	}
	if this.GameOneType != "" {
		sess.Where("game_name=?", this.GameOneType)
	} else {
		vtypeStr := make([]string, 0)
		if this.GameName != "" {
			switch this.GameName {
			case "pk_fc":
				data, _ := c.GetPkGameList()
				for _, v := range data {
					vtypeStr = append(vtypeStr, v.Type)
				}
			case "cs_fc":
				data, _ := c.GetCsGameList()
				for _, v := range data {
					vtypeStr = append(vtypeStr, v.CsType)
				}
			case "eg_fc":
				data, _ := c.GetEgGameList()
				for _, v := range data {
					vtypeStr = append(vtypeStr, v.EgType)
				}
			default:
				vtypeStr = nil
			}
			if vtypeStr != nil {
				sess.In("game_name", vtypeStr)
			} else {
				sess.Where("game_name=?", this.GameName)
			}
		}
	}
	sess.Where("username = ?", Account)
	sess.Where("status=?", 1)
	sess.OrderBy("id desc")
	conds := sess.Conds()
	listparams.Make(sess)
	err = sess.Table(betRecordInfo.TableName()).Find(&data)
	count, _ = sess.Table(betRecordInfo.TableName()).Where(conds).Count()
	return
}

//获得cs彩票游戏具体分类
func (c *BetRecordInfoBean) GetCsGameList() (data []schema.CsGames, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	betRecordInfo := new(schema.CsGames)
	err = sess.Table(betRecordInfo.TableName()).
		Where("id>?", 0).
		Where("cs_state=?", 1).
		Find(&data)
	return
}

//获得eg彩票游戏具体分类
func (c *BetRecordInfoBean) GetEgGameList() (data []schema.EgGames, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	betRecordInfo := new(schema.EgGames)
	err = sess.Table(betRecordInfo.TableName()).
		Where("id>?", 0).
		Where("eg_state=?", 1).
		Find(&data)
	return
}

//获得pk彩票游戏具体分类
func (c *BetRecordInfoBean) GetPkGameList() (data []schema.PkGames, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	betRecordInfo := new(schema.PkGames)
	err = sess.Table(betRecordInfo.TableName()).
		Where("id>?", 0).
		Where("state=?", 1).
		Find(&data)
	return
}
