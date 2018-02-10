package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type OverRideBean struct{}

//退佣列表
func (*OverRideBean) OverRideList(this *input.OverRideGet) ([]back.OverRide, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	override := new(schema.SiteRebateProduct)
	site := new(schema.SiteRebateSet)
	product := new(schema.ProductType)
	var data []back.OverRide
	sql2 := fmt.Sprintf("%s.set_id = %s.id", override.TableName(), site.TableName())
	sql3 := fmt.Sprintf("%s.product_id = %s.id", override.TableName(), product.TableName())
	//"sales_site_rebate_set.site_id = sales_site_rebate_product.id"
	sess.Where("sales_site_rebate_set.delete_time=?", 0)
	sess.Where("sales_site_rebate_set.site_id=?", this.SiteId)
	sess.Where("sales_site_rebate_set.site_index_id=?", this.SiteIndexId)
	err := sess.Table(override.TableName()).
		Select("sales_site_rebate_set.id,sales_site_rebate_product.product_id,sales_site_rebate_product.`rebate_radio`,sales_site_rebate_product.`water_radio`,sales_site_rebate_set.`effective_user`,sales_site_rebate_set.site_id,sales_site_rebate_set.site_index_id,sales_site_rebate_set.`self_profit`,sales_site_rebate_set.`valid_money`,sales_product_type.`title`").
		Join("LEFT", site.TableName(), sql2).
		Join("LEFT", product.TableName(), sql3).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}

	count, err := sess.Table(site.TableName()).
		Where("sales_site_rebate_set.site_id=?", this.SiteId).
		Where("sales_site_rebate_set.site_index_id=?", this.SiteIndexId).
		Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取单条代理退佣数据
func (*OverRideBean) OverGetOne(this *input.OverGetOne) ([]back.OverRide, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	override := new(schema.SiteRebateProduct)
	site := new(schema.SiteRebateSet)
	product := new(schema.ProductType)
	var data []back.OverRide
	sql2 := fmt.Sprintf("%s.set_id = %s.id", override.TableName(), site.TableName())
	sql3 := fmt.Sprintf("%s.product_id = %s.id", override.TableName(), product.TableName())
	//"sales_site_rebate_set.site_id = sales_site_rebate_product.id"
	sess.Where("sales_site_rebate_set.delete_time=?", 0)
	sess.Where("sales_site_rebate_set.site_id=?", this.SiteId)
	sess.Where("sales_site_rebate_set.site_index_id=?", this.SiteIndexId)
	sess.Where("sales_site_rebate_set.id=?", this.Id)
	err := sess.Table(override.TableName()).
		Select("sales_site_rebate_set.id,sales_site_rebate_product.product_id,sales_site_rebate_product.`rebate_radio`,sales_site_rebate_product.`water_radio`,sales_site_rebate_set.`effective_user`,sales_site_rebate_set.site_id,sales_site_rebate_set.site_index_id,sales_site_rebate_set.`self_profit`,sales_site_rebate_set.`valid_money`,sales_product_type.`title`").
		Join("LEFT", site.TableName(), sql2).
		Join("LEFT", product.TableName(), sql3).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//修改单条代理退佣设定数据
func (*OverRideBean) UpdateOver(this *input.OverRideUpdate) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	override := new(schema.SiteRebateSet)
	site := new(schema.SiteRebateProduct)
	override.SelfProfit = this.Amount
	override.EffectiveUser = this.Member
	count, err := sess.Table(override.TableName()).Where("id = ?", this.Id).Update(override)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	count, err = sess.Table(site.TableName()).Where("set_id = ?", this.Id).Delete(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}

	for key := range this.List {
		this.List[key].SetId = this.Id
		fmt.Printf("%#v\n", this.Id)
	}
	fmt.Printf("%#v\n", this.List)
	count, err = sess.Table(site.TableName()).Insert(this.List)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//删除一条代理退佣设定
func (*OverRideBean) DeleteOver(this *input.OverRideDelet) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.SiteRebateSet)
	role.DeleteTime = global.GetCurrentTime()
	count, err := sess.Table(role.TableName()).Where("id = ?", this.Id).Update(role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//新增一条代理退佣设定
func (*OverRideBean) OverRideAdd(this *input.OverRideAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	var data int64
	defer sess.Close()
	combo := new(schema.SiteRebateSet)
	combo.SiteId = this.SiteId
	combo.SiteIndexId = this.SiteIndexId
	combo.SelfProfit = this.Amount
	combo.EffectiveUser = this.Member
	err := sess.Table(combo.TableName()).Select("valid_money").
		Where("site_index_id=?", this.SiteIndexId).
		Where("site_id=?", this.SiteId).Find(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	combo.ValidMoney = data
	sess.Begin()
	count, err := sess.Table(combo.TableName()).Insert(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	Id := combo.Id
	site := new(schema.SiteRebateProduct)
	for key := range this.List {
		this.List[key].SetId = Id
	}
	count, err = sess.Table(site.TableName()).Insert(this.List)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	sess.Commit()
	return count, err
}

//修改有效会员投注金额
func (*OverRideBean) UpdataMoney(this *input.UpdateMoney) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebate := new(schema.SiteRebateSet)
	rebate.ValidMoney = int64(this.Money)
	sess.Where("site_id=?", this.SiteId)
	count, err := sess.Table(rebate.TableName()).
		Where("site_index_id = ?", this.SiteIndexId).Update(rebate)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
