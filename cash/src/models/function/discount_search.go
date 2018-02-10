//优惠查询
package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type DiscountSearchBean struct{}

//优惠列表
func (*DiscountSearchBean) DiscountSearchLIst(this *input.DiscountSearchList, listparams *global.ListParams) (data []back.DiscountSearchListBack, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	listparams.Make(sess)
	site := new(schema.Site)
	err = sess.Table(site.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(site.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//优惠查询统计
func (*DiscountSearchBean) GetDiscountCount(this *input.DiscountAllList,
	search_time int64, end_time int64) (data []back.DiscountList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	retreat := new(schema.MemberRetreatWater)
	err = sess.Table(retreat.TableName()).
		Where("create_time >?", search_time).
		Where("create_time<?", end_time).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Find(&data)
	if err != nil {
		return data, err
	}
	return data, err
}

//优惠查询明细
func (*DiscountSearchBean) GetDiscountList(this *input.DiscountInfo) (data []back.DiscountGetInfo, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	record := new(schema.MemberRetreatWaterRecord)
	recordProduct := new(schema.MemberRetreatWaterRecordProduct)
	if this.Id != 0 {
		sess.Where("r.periods_id=?", this.Id)
	}
	err = sess.Table(record.TableName()).Alias("r").
		Select("r.id,r.site_id,r.site_index_id,r.account,r.start_time,r.end_time,r.member_id,r.level_id,r.betall,r.all_money,r.self_money,r.rebate_water,r.status,r.create_time,p.product_id,p.product_bet,p.rate,p.money").
		Join("LEFT", []string{recordProduct.TableName(), "p"}, "r.id=p.record_id").
		Where("r.site_id=?", this.SiteId).
		Where("r.site_index_id=?", this.SiteIndexId).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
