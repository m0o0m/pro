package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type RetirementBean struct{}

//期数列表
func (*RetirementBean) RetirementList(this *input.PeriodsGet) ([]back.RetirementList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.RetirementList
	retirement := new(schema.MemberRetreatWaterRecord)
	err := sess.Table(retirement.TableName()).Where("site_id=?", this.SiteId).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
func (*RetirementBean) CheckList(this *input.CheckList) (
	[]back.Retirement, []back.GetProductList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	retirement := new(schema.AgencyRebateRecord)
	var data []back.Retirement
	var list []back.GetProductList
	rebate := new(schema.AgencyRebateRecordProduct)
	sess.Where("c.site_index_id=?", this.SiteIndexId)
	sess.Where("c.periods_id=?", this.PeriodsId)
	sess.Where("c.agency_account=?", this.AgencyAccount)
	if this.Rebate == 1 {
		sess.Where("c.rebate>?", 0)
	} else {
		sess.Where("c.rebate=?", 0)
	}
	sess.Where("effective_member >?", this.StartNum)
	sess.Where("effective_member <?", this.EndNum)
	sql1 := "r.record_id=c.id"
	err := sess.Table(retirement.TableName()).Alias("c").
		Join("LEFT", []string{rebate.TableName(), "r"}, sql1).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, list, err
	}
	product := new(schema.ProductType)
	err = sess.Table(product.TableName()).Select("id,title").Where("delete_time=?", 0).Find(&list)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, list, err
	}
	return data, list, err
}
