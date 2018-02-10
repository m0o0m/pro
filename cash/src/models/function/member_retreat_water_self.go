package function

import (
	"errors"
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strconv"
	"time"
)

type MemberRetreatWaterSelfBean struct {
}

//自助返水查询列表
func (*MemberRetreatWaterSelfBean) SearchMemberRetreatWaterSelf(this *input.ListRetreatWaterSelf, listparam *global.ListParams) (
	[]back.ListRetreatWaterSelf, int64, error) {
	water := new(schema.MemberRetreatWaterSelf)
	sess := global.GetXorm().Table(water.TableName())
	defer sess.Close()
	var data []back.ListRetreatWaterSelf
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.StartTime != 0 {
		sess.Where("create_time>=?", this.StartTime)
	}
	if this.EndTime != 0 {
		sess.Where("create_time<?", this.EndTime)
	}
	if this.Account != "" {
		sess.Where("account=?", this.Account)
	}
	if this.OrderNum != 0 {
		sess.Where("order_num=?", this.OrderNum)
	}
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	err := sess.Table(water.TableName()).
		Select("id,order_num,account,betting,money,create_time").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(water.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取单个列表详情
func (*MemberRetreatWaterSelfBean) DetailMemberRetreatWaterSelf(this *input.DetailRetreatWaterSelf) (
	back.RetreatWaterRecordListTotalSelf, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	water := new(schema.MemberRetreatWaterRecordSelf)
	var data back.RetreatWaterRecordListTotalSelf
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("sales_member_retreat_water_record_self.period_id=?", this.Id)
	conds := sess.Conds()
	count, err := sess.Table(water.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	waterPro := new(schema.MemberRetreatWaterRecordProductSelf)
	product := new(schema.Product)
	sql2 := fmt.Sprintf("%s.id = %s.record_id", water.TableName(), waterPro.TableName())
	sql3 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), waterPro.TableName())
	data1 := make([]back.RetreatWaterRecordAndProductSelf, 0)
	err = sess.Table(water.TableName()).
		Select("sales_member_retreat_water_record_self.id,account,member_id,betall,rebate_water,sales_member_retreat_water_record_self.create_time,product_id,product_name,product_bet,money").
		Join("LEFT", waterPro.TableName(), sql2).
		Join("LEFT", product.TableName(), sql3).OrderBy("sales_member_retreat_water_record_self.id,product_id").
		Find(&data1)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	var checkid int64 = 0
	waters := make([]back.RetreatWaterRecordSelf, 0)
	waterPros := make([]back.RetreatWaterRecordProductSelf, 0)
	water2 := new(back.RetreatWaterRecordSelf)
	waterPro2 := new(back.RetreatWaterRecordProductSelf)
	waterTot2 := new(back.RetreatWaterRecordTotalSelf)           //明细列表总计total
	waterProTot := make([]back.RetreatWaterRecordProductSelf, 0) //明细列表商品总计 商品切片
	waterProTot2 := new(back.RetreatWaterRecordProductSelf)      //明细列表商品总计
	var n int64 = -1
	m := 0
	for i, d := range data1 {
		if checkid != d.Id { //id不同时，组装有效打码，上限，商品列表，总组装
			n = n + 1         //明细列表条数
			m = i - m         //用于取余，遍历每条明细的商品
			if checkid != 0 { //给checkid赋初始值
				water2.Params = waterPros        //商品列表 组装到总组装Params参数
				waters = append(waters, *water2) //总组装
			}
			waterPros = nil //id不同时清空
			checkid = d.Id  //因为是以id排序的,所以id都是联系的序列组，置id相同

			//会员账号,所属会员id,有效总投注,本次退水金额,返水时间
			water2.Id = d.Id
			water2.Account = d.Account
			water2.MemberId = d.MemberId
			water2.Betall = d.Betall
			water2.RebateWater = d.RebateWater
			water2.CreateTime = d.CreateTime

			//累加总个数，有效总投注，本次退水金额
			waterTot2.TotalNum = waterTot2.TotalNum + n
			waterTot2.TotalBetall = waterTot2.TotalBetall + d.Betall
			waterTot2.TotalRebateWater = waterTot2.TotalRebateWater + d.RebateWater

		}
		//每次都组装商品列表
		waterPro2.ProductId = d.ProductId     //商品分类id
		waterPro2.ProductName = d.ProductName //商品名
		waterPro2.ProductBet = d.ProductBet   //商品投注额
		waterPro2.Money = d.Money             //金额
		waterPros = append(waterPros, *waterPro2)

		if n == 0 {
			waterProTot2.ProductId = d.ProductId
			waterProTot2.ProductName = d.ProductName
			waterProTot2.ProductBet = d.ProductBet
			waterProTot2.Money = d.Money
			waterProTot = append(waterProTot, *waterProTot2)
		} else { //建立在商品都固定的基础上，这样循环，每次都能对应上相应的商品
			waterProTot[i%m].ProductId = d.ProductId
			waterProTot[i%m].ProductName = d.ProductName
			waterProTot[i%m].ProductBet = d.ProductBet + waterProTot[i%m].ProductBet
			waterProTot[i%m].Money = d.Money + waterProTot[i%m].Money
		}
	}
	water2.Params = waterPros
	if checkid != 0 {
		waters = append(waters, *water2)
	}
	waterTot2.TotalNum = waterTot2.TotalNum + 1
	waterTot2.Params = waterProTot

	data.List = waters
	data.Total = *waterTot2

	return data, count, err
}

//获取各商品反水比例
func (*MemberRetreatWaterSelfBean) GetProductReWaterRate(memberInfo *global.MemberRedisToken, startTime, endTime int64) (data []back.BackRate, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	result, err := sess.Query("SELECT sales_bet_report_account.v_type,"+
		"sales_bet_report_account.id,"+
		"sales_bet_report_account.bet_valid,"+
		"c.product_name,c.id as product_id FROM "+
		"sales_bet_report_account LEFT JOIN "+
		"sales_product c ON sales_bet_report_account.v_type = "+
		"c.v_type WHERE sales_bet_report_account.site_id = ? "+
		"AND sales_bet_report_account.site_index_id = ? "+
		"AND sales_bet_report_account.account = ? "+
		"AND sales_bet_report_account.day_time >= ? "+
		"AND sales_bet_report_account.day_time <= ? "+
		"AND sales_bet_report_account.v_type "+
		"IN (SELECT sales_product.v_type "+
		"FROM sales_product WHERE id NOT "+
		"IN (SELECT product_id FROM sales_site_product_del "+
		"WHERE site_id = ? AND site_index_id = ?)"+
		"AND status = ? AND delete_time = ?)",
		memberInfo.Site,
		memberInfo.SiteIndex,
		memberInfo.Account,
		startTime,
		endTime,
		memberInfo.Site,
		memberInfo.SiteIndex, 1, 0)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	var rate back.BackRate
	for _, k := range result {
		rate.Id, err = strconv.ParseInt(string(k["id"]), 10, 64)
		if err != nil {
			return data, err
		}
		rate.BetValid, err = strconv.ParseFloat(string(k["bet_valid"]), 64)
		if err != nil {
			return data, err
		}
		rate.ProductId, err = strconv.ParseInt(string(k["product_id"]), 10, 64)
		if err != nil {
			return data, err
		}
		rate.ProductName = string(k["product_name"])
		rate.Vtype = string(k["v_type"])
		data = append(data, rate)
	}
	return
}

//获取会员今日已经反水的额度
func (*MemberRetreatWaterSelfBean) GetMemberRetreatToday(memberInfo *global.MemberRedisToken, startTime, endTime int64) (backData schema.MemberRetreatWaterSelf, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Where("site_id=?", memberInfo.Site).
		Where("site_index_id=?", memberInfo.SiteIndex).
		Where("member_id=?", memberInfo.Id).
		Where("create_time>=?", startTime).
		Where("create_time<=?", endTime).
		Get(&backData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, flag, err
	}
	return backData, flag, err
}

//在sales_member_retreat_water_record_self中间获取今日反水总额
func (*MemberRetreatWaterSelfBean) GetMemberRetreatNowDayTotal(memberInfo *global.MemberRedisToken, startTime, endTime int64) (total float64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	retreatWaterRecord := new(schema.MemberRetreatWaterRecordSelf)
	total, err = sess.Table(retreatWaterRecord.TableName()).Where("site_id=?", memberInfo.Site).
		Where("site_index_id=?", memberInfo.SiteIndex).
		Where("member_id=?", memberInfo.Id).
		Where("create_time>=?", startTime).
		Where("create_time<=?", endTime).
		Sum(retreatWaterRecord, "rebate_water")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return total, err
	}
	return total, err
}

//获取该站点，该产品，该设定的反水比例
func (*MemberRetreatWaterSelfBean) GetInfoBySiteAndProSetId(memberInfo *global.MemberRedisToken, productId int64, validMoney float64) (rate float64, discountup int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	result, err := sess.Query("SELECT rate,sales_member_retreat_water_set.discount_up FROM sales_member_retreat_water_product LEFT JOIN sales_member_retreat_water_set ON sales_member_retreat_water_product.set_id = sales_member_retreat_water_set.id WHERE set_id = (SELECT id FROM sales_member_retreat_water_set WHERE valid_money = (SELECT MAX(valid_money) FROM sales_member_retreat_water_set WHERE valid_money <= ?) AND site_id = ? AND site_index_id = ?)AND product_id = ?",
		validMoney, memberInfo.Site, memberInfo.SiteIndex, productId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return rate, discountup, err
	}
	for _, k := range result {
		if string(k["rate"]) != "" {
			rate, err = strconv.ParseFloat(string(k["rate"]), 64)
			if err != nil {
				return rate, discountup, err
			}
			discountup, err = strconv.ParseInt(string(k["discount_up"]), 10, 64)
			if err != nil {
				return rate, discountup, err
			}
		} else {
			rate = 0.00
			discountup, err = strconv.ParseInt(string(k["discount_up"]), 10, 64)
			if err != nil {
				return rate, discountup, err
			}
		}
	}
	return
}

//根据产品id和日期,类型,会员每日打码统计表id获取有效投注总额
func (*MemberRetreatWaterSelfBean) GetBetValidById(memberInfo *global.MemberRedisToken, selectDiffer *input.SingleReWater) (memberWater schema.BetReportAccount, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	timeUnix, err := time.Parse("2006-01-02", selectDiffer.DateTimes)
	if err != nil {
		return
	}
	sess.Where("id=?", selectDiffer.Id)
	sess.Where("day_time=?", timeUnix.Unix())
	sess.Where("v_type=?", selectDiffer.Vtype)
	sess.Where("member_id=?", memberInfo.Id)
	sess.Where("site_id=?", memberInfo.Site)
	sess.Where("site_index_id=?", memberInfo.SiteIndex)
	flag, err = sess.Get(&memberWater)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberWater, flag, err
	}
	return memberWater, flag, err
}

//一键领取所有的反水
func (*MemberRetreatWaterSelfBean) OneClickGetAllReWater(memberInfo *global.MemberRedisToken, oneClick *input.OneClickGetAllReWater) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	newTime, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		return
	}
	//获取会员详情
	memberBeans := new(MemberBean)
	info, flag, err := memberBeans.GetInfoById(memberInfo.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	if !flag || info.Id == 0 {
		return count, err
	}

	//获取会员自助反水期数，存在更新会员自助水期数，不存在添加
	memberRetreatWaterSelf := new(schema.MemberRetreatWaterSelf)
	reWaterRecord := new(schema.MemberRetreatWaterRecordSelf)
	memberCashRecord := new(schema.MemberCashRecord)
	member := new(schema.Member)
	flag, err = sess.Where("site_id=?", memberInfo.Site).
		Where("site_index_id=?", memberInfo.SiteIndex).
		Where("member_id=?", memberInfo.Id).
		Where("create_time>=?", newTime.Unix()).
		Get(memberRetreatWaterSelf)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//存在更新
	if flag {
		memberRetreatWaterSelf.Betting = oneClick.BetValid   //有效投注额度
		memberRetreatWaterSelf.Money = oneClick.RewaterTotal //总的反水额度
		count, err := sess.Where("id=?", memberRetreatWaterSelf.Id).
			Cols("betting,money").
			Update(memberRetreatWaterSelf)
		if err != nil || count != 1 {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		reWaterRecord.RebateWater = oneClick.RewaterTotal - memberRetreatWaterSelf.Money                    //本次反水记录反水额度
		member.Balance = info.Balance + oneClick.RewaterTotal                                               //会员余额
		memberCashRecord.DisBalance = oneClick.RewaterTotal - memberRetreatWaterSelf.Money                  //优惠金额
		memberCashRecord.AfterBalance = info.Balance + oneClick.RewaterTotal - memberRetreatWaterSelf.Money //操作后金额
	} else { //不存在添加
		memberRetreatWaterSelf.MemberId = strconv.FormatInt(memberInfo.Id, 10)
		memberRetreatWaterSelf.Money = oneClick.RewaterTotal
		memberRetreatWaterSelf.SiteId = memberInfo.Site
		memberRetreatWaterSelf.SiteIndexId = memberInfo.SiteIndex
		memberRetreatWaterSelf.Account = memberInfo.Account
		memberRetreatWaterSelf.CreateTime = newTime.Unix()
		memberRetreatWaterSelf.Betting = oneClick.BetValid   //有效投注额度
		memberRetreatWaterSelf.Money = oneClick.RewaterTotal //总的反水额度
		count, err := sess.Insert(memberRetreatWaterSelf)
		if err != nil || count != 1 {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		reWaterRecord.RebateWater = oneClick.RewaterTotal
		member.Balance = info.Balance + oneClick.RewaterTotal                //会员余额
		memberCashRecord.DisBalance = oneClick.RewaterTotal                  //优惠金额
		memberCashRecord.AfterBalance = info.Balance + oneClick.RewaterTotal //操作后余额
	}
	//添加会员自助反水记录
	reWaterRecord.SiteId = memberInfo.Site
	reWaterRecord.SiteIndexId = memberInfo.SiteIndex
	reWaterRecord.MemberId = strconv.Itoa(int(memberInfo.Id))
	reWaterRecord.Account = memberInfo.Account
	reWaterRecord.CreateTime = time.Now().Unix()
	reWaterRecord.Betall = oneClick.BetValid
	reWaterRecord.PeriodId = memberRetreatWaterSelf.Id
	count, err = sess.Insert(reWaterRecord)
	if err != nil || count != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//更新会员余额
	count, err = sess.Cols("balance").Where("id=?", memberInfo.Id).Update(member)
	if err != nil || count != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//增加现金流水记录
	memberCashRecord.SiteId = memberInfo.Site
	memberCashRecord.SiteIndexId = memberInfo.SiteIndex
	memberCashRecord.SourceType = 9
	memberCashRecord.Type = 1
	memberCashRecord.TradeNo = ""
	memberCashRecord.ClientType = 2
	memberCashRecord.CreateTime = time.Now().Unix()
	memberCashRecord.Remark = "自助优惠反水"
	memberCashRecord.UserName = memberInfo.Account
	memberCashRecord.MemberId = memberInfo.Id
	memberCashRecord.Balance = 0
	memberCashRecord.AgencyId = info.ThirdAgencyId
	//根据代理id获取代理账号
	agencyBean := new(AgencyBean)
	agencyInfo, flag, err := agencyBean.GetAgency(info.ThirdAgencyId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//todo 这里代理没有找到应该怎么处理,我暂时给空值
	if flag || agencyInfo.Id == 0 {
		memberCashRecord.AgencyAccount = agencyInfo.Account
	} else {
		memberCashRecord.AgencyAccount = ""
	}
	memberCash := new(MemberCashRecordBean)
	count, err = memberCash.AddNewRecord(memberCashRecord)
	if err != nil || count != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	sess.Commit()
	return count, err
}

//获取该站点，符合有效打码设定的反水比例
func (*MemberRetreatWaterSelfBean) GetSiteWaterSetByValidBetAll(memberInfo *global.MemberRedisToken, betValidAll float64) (result []back.ReMembetWaterRate, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err = sess.SQL("SELECT product_id,rate,sales_member_retreat_water_set.discount_up FROM sales_member_retreat_water_product LEFT JOIN sales_member_retreat_water_set ON sales_member_retreat_water_product.set_id = sales_member_retreat_water_set.id WHERE set_id = (SELECT id FROM sales_member_retreat_water_set WHERE valid_money = (SELECT MAX(valid_money) FROM sales_member_retreat_water_set WHERE valid_money <= ? AND site_id = ? AND site_index_id = ?) AND site_id = ? AND site_index_id = ?)",
		betValidAll, memberInfo.Site, memberInfo.SiteIndex, memberInfo.Site, memberInfo.SiteIndex).Find(&result)
	return
}

//自助反水存入
func (*MemberRetreatWaterSelfBean) PostMemberReWater(memberInfo *global.MemberRedisToken, oneClick back.WaterData) (err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err = sess.Begin()
	loc, _ := time.LoadLocation("Local")
	newTime, err := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 00:00:00", loc)
	if err != nil {
		return err
	}
	//获取会员详情
	memberBeans := new(MemberBean)
	info, flag, err := memberBeans.GetInfoById(memberInfo.Id)
	if err != nil {
		return err
	}
	if !flag || info.Id == 0 {
		return errors.New("this member is not exist")
	}

	//获取会员自助反水期数，存在更新会员自助水期数，不存在添加
	memberRetreatWaterSelf := new(schema.MemberRetreatWaterSelf)
	reWaterRecord := new(schema.MemberRetreatWaterRecordSelf)
	memberCashRecord := new(schema.MemberCashRecord)
	member := new(schema.Member)
	flag, err = sess.Table(memberRetreatWaterSelf.TableName()).
		Where("site_id=?", memberInfo.Site).
		Where("site_index_id=?", memberInfo.SiteIndex).
		Where("member_id=?", memberInfo.Id).
		Where("create_time=?", newTime.Unix()).
		Get(memberRetreatWaterSelf)
	if err != nil {
		return err
	}
	//存在更新
	waterData := oneClick.Data
	var betall float64
	for _, v := range waterData {
		betall += v.BetValid
	}
	if flag {
		//在改变总额前赋值
		reWaterRecord.RebateWater = oneClick.Count - memberRetreatWaterSelf.Money                    //本次反水记录反水额度
		member.Balance = info.Balance + oneClick.Count - memberRetreatWaterSelf.Money                //会员余额
		memberCashRecord.DisBalance = oneClick.Count - memberRetreatWaterSelf.Money                  //优惠金额
		memberCashRecord.AfterBalance = info.Balance + oneClick.Count - memberRetreatWaterSelf.Money //操作后金额

		memberRetreatWaterSelf.Betting = betall       //有效投注额度
		memberRetreatWaterSelf.Money = oneClick.Count //总的反水额度
		count, err := sess.Where("site_id=?", memberInfo.Site).
			Where("site_index_id=?", memberInfo.SiteIndex).
			Where("member_id=?", memberInfo.Id).
			Where("create_time=?", newTime.Unix()).
			Cols("betting,money").
			Update(memberRetreatWaterSelf)
		if err != nil || count != 1 {
			sess.Rollback()
			return err
		}
	} else { //不存在添加
		memberRetreatWaterSelf.MemberId = strconv.Itoa(int(memberInfo.Id))
		memberRetreatWaterSelf.Money = oneClick.Count
		memberRetreatWaterSelf.SiteId = memberInfo.Site
		memberRetreatWaterSelf.SiteIndexId = memberInfo.SiteIndex
		memberRetreatWaterSelf.Account = memberInfo.Account
		memberRetreatWaterSelf.CreateTime = newTime.Unix()
		memberRetreatWaterSelf.Betting = betall //有效投注额度
		count, err := sess.Insert(memberRetreatWaterSelf)
		if err != nil || count != 1 {
			sess.Rollback()
			return err
		}
		reWaterRecord.RebateWater = oneClick.Count
		member.Balance = info.Balance + oneClick.Count                //会员余额
		memberCashRecord.DisBalance = oneClick.Count                  //优惠金额
		memberCashRecord.AfterBalance = info.Balance + oneClick.Count //操作后余额
	}
	//添加会员自助反水记录
	reWaterRecord.SiteId = memberInfo.Site
	reWaterRecord.SiteIndexId = memberInfo.SiteIndex
	reWaterRecord.MemberId = strconv.Itoa(int(memberInfo.Id))
	reWaterRecord.Account = memberInfo.Account
	reWaterRecord.CreateTime = time.Now().Unix()
	reWaterRecord.Betall = betall
	reWaterRecord.PeriodId = memberRetreatWaterSelf.Id
	count, err := sess.Insert(reWaterRecord)
	if err != nil || count != 1 {
		sess.Rollback()
		return
	}
	//添加自助反水记录的对应商品记录表
	var SelfData []*schema.MemberRetreatWaterRecordProductSelf
	for _, v := range waterData {
		reWaterRecordProduct := new(schema.MemberRetreatWaterRecordProductSelf)
		reWaterRecordProduct.RecordId = reWaterRecord.Id
		reWaterRecordProduct.ProductId = v.ProductId
		reWaterRecordProduct.ProductBet = v.BetValid
		reWaterRecordProduct.Rate = v.Rate
		reWaterRecordProduct.Money = v.RateMoney
		SelfData = append(SelfData, reWaterRecordProduct)
	}
	count, err = sess.InsertMulti(SelfData)
	if err != nil || count <= 1 {
		sess.Rollback()
		return
	}

	//更新会员余额
	count, err = sess.Cols("balance").Where("site_id=?", memberInfo.Site).
		Where("site_index_id=?", memberInfo.SiteIndex).
		Where("id=?", memberInfo.Id).Update(member)
	if err != nil || count != 1 {
		sess.Rollback()
		return err
	}
	//增加现金流水记录
	memberCashRecord.SiteId = memberInfo.Site
	memberCashRecord.SiteIndexId = memberInfo.SiteIndex
	memberCashRecord.SourceType = 9
	memberCashRecord.Type = 1
	memberCashRecord.TradeNo = ""
	memberCashRecord.ClientType = 2
	memberCashRecord.CreateTime = time.Now().Unix()
	memberCashRecord.Remark = "自助优惠反水"
	memberCashRecord.UserName = memberInfo.Account
	memberCashRecord.MemberId = memberInfo.Id
	memberCashRecord.Balance = 0
	memberCashRecord.AgencyId = info.ThirdAgencyId
	//根据代理id获取代理账号
	agencyBean := new(AgencyBean)
	agencyInfo, flag, err := agencyBean.GetAgency(info.ThirdAgencyId)
	if err != nil {
		sess.Rollback()
		return
	}
	if flag && agencyInfo.Id != 0 {
		memberCashRecord.AgencyAccount = agencyInfo.Account
	} else {
		memberCashRecord.AgencyAccount = ""
	}
	memberCash := new(MemberCashRecordBean)
	count, err = memberCash.AddNewRecord(memberCashRecord)
	if err != nil || count != 1 {
		return err
	}
	err = sess.Commit()
	return
}
