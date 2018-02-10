package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type RebateCountBean struct{}

//期数列表
func (*RebateCountBean) RetirementList(this *input.PeriodsGet) (data []back.RetirementList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	retirement := new(schema.SiteRebatePeriods)
	err = sess.Table(retirement.TableName()).Where("site_id=?", this.SiteId).Find(&data)
	return
}

//获取期数对应时间
func (*RebateCountBean) GetTime(this *input.RebateInput) (data []back.Periods, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SiteRebatePeriods)
	err = sess.Table(periods.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("id=?", this.PeriodsId).
		Find(&data)
	return
}

//获取有效会员代理id
func (*RebateCountBean) GetAgencyId(this *input.RebateInput) ([]back.RebateCountList1, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//获取有会员的代理id
	var data []back.RebateCountList1
	rebate := new(schema.AgencyCount)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("member_count>?", 0)
	err := sess.Table(rebate.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取代理信息
func (*RebateCountBean) GetAgencyInfo(agencylist []int64, this *input.RebateInput) ([]back.AgencyInfo, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//获取代理信息
	var data []back.AgencyInfo
	agency := new(schema.Agency)
	sess.Where("site_id=?", this.SiteId)
	sess.In("id", agencylist)
	err := sess.Table(agency.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取有效会员门槛
func (*RebateCountBean) GetRebate(this *input.RebateInput) ([]back.Rebatec, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//获取有效会员门槛
	var data []back.Rebatec
	rebateset := new(schema.SiteRebateSet)
	err := sess.Table(rebateset.TableName()).
		Where("site_id=?", this.SiteId).
		Where("delete_time=?", 0).
		Where("site_index_id=?", this.SiteIndexId).
		OrderBy("self_profit desc").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取退佣退水比例
func (*RebateCountBean) GetPoundageSet(id int64) (data []back.RebateProduct) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SiteRebateProduct)
	sess.Table(periods.TableName()).Where("set_id=?", id).Find(&data)
	return
}

//获取退佣手续费设定
func (*RebateCountBean) GetCharge(this *input.RebateInput) ([]back.Charge, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//获取退佣手续费设定
	sitepoun := new(schema.SitePoundage)
	var data []back.Charge
	err := sess.Table(sitepoun.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取其他优惠数据
//返水费用
//入款费用 公司线上入款//后台人工入款//出款费用
func (*RebateCountBean) ReturnCompany(time1 int64, time2 int64, agencylist []int64, soucetype []int, this *input.RebateInput) (
	[]back.CashRecord, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.CashRecord
	cash := new(schema.MemberCashRecord)
	err := sess.Table(cash.TableName()).Select("id,site_id,source_type,site_index_id,"+
		"sum(balance) as balance,sum(dis_balance) as dis_balance,"+
		"member_id,user_name,agency_id").
		In("source_type", soucetype).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		In("agency_id", agencylist).
		Where("create_time>=?", time1).
		Where("create_time<=?", time2).
		GroupBy("source_type,agency_id").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取所有报表
func (*RebateCountBean) GetReport(time1 int64, time2 int64, agencylist []int64, gameType []int64, this *input.RebateInput) (data []back.ReportAll, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	cash := new(schema.BetReportAccount) //每日打码统计
	err = sess.Table(cash.TableName()).
		Select("sum(num) as num,sum(bet_all) as bet_all,sum(bet_valid) as bet_valid,sum(win) as win,sum(win_num) as win_num,sum(jack) as jack,day_time,game_type,platform,v_type,agency_id").
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		In("agency_id", agencylist).
		Where("day_time>=?", time1).
		Where("day_time<=?", time2).
		In("game_type", gameType).
		GroupBy("member_id,game_type,agency_id").
		Find(&data)
	return
}

//获取上一期期数id
func (*RebateCountBean) GetLastId(this *input.RebateInput) (data []back.RebateCountList) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SiteRebatePeriods)
	sess.Table(periods.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("status=?", 1).
		Where("delete_time=?", 0).
		OrderBy("id desc").
		Find(&data)
	return
}

//获得上一期数据 agency_rebate_record
func (*RebateCountBean) GetLastData(pid int64, this *input.RebateInput) (data []back.LastRebate) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebate := new(schema.AgencyRebateRecord)
	//获取前期数据
	sess.Table(rebate.TableName()).
		Select("sum(before_jack+now_jack) as before_jack,sum(before_cost+now_cost) as before_cost,sum(before_betting+now_betting) as before_betting,sum(now_profit+before_profit) as before_profit,agency_id,site_id,site_index_id,statue,effective_member").
		Where("periods_id=?", pid).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		GroupBy("agency_id").
		Find(&data)
	return
}

//代理退佣存档
func (*RebateCountBean) RebateFile(data []back.RebateListIn) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebate := new(schema.AgencyRebateRecord)
	num, err := sess.Table(rebate.TableName()).Insert(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return num, err
	}
	return num, err
}
