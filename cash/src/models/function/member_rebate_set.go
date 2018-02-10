package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

//会员返佣设定
type MemberRebateSetBean struct {
}

//删除会员返佣设定
func (*MemberRebateSetBean) Del(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebateSchema := &schema.MemberRebateSet{}
	rebateSchema.DeleteTime = time.Now().Unix()
	count, err := sess.Table(rebateSchema.TableName()).
		Where("id = ?", id).Cols("delete_time").
		Update(rebateSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//添加会员返佣设定
func (*MemberRebateSetBean) Add(rebateSet *input.MemberRebateSetAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebateSchema := &schema.MemberRebateSet{}
	rebateSchema.SiteIndexId = rebateSet.SiteIndexId
	rebateSchema.SiteId = rebateSet.SiteId
	rebateSchema.ValidMoney = rebateSet.ValidMoney
	rebateSchema.DiscountUp = rebateSet.DiscountUp
	count, err = sess.Table(rebateSchema.TableName()).InsertOne(rebateSchema)
	return
}

//更新会员返佣设定
func (*MemberRebateSetBean) Update(rebateSet *input.MemberRebateSetAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebateSchema := &schema.MemberRebateSet{}
	rebateSchema.SiteIndexId = rebateSet.SiteIndexId
	rebateSchema.SiteId = rebateSet.SiteId
	rebateSchema.ValidMoney = rebateSet.ValidMoney
	rebateSchema.DiscountUp = rebateSet.DiscountUp
	count, err := sess.Table(rebateSchema.TableName()).
		Where("delete_time >0").Update(rebateSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询会员返佣设定详情
func (*MemberRebateSetBean) Select(id int64) (back.MemberRebateSet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebateSchema := &schema.MemberRebateSet{}
	var rebateSet back.MemberRebateSet
	b, err := sess.Table(rebateSchema.TableName()).
		Where("id = ?", id).
		Where("delete_time >0").Get(&rebateSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return rebateSet, b, err
	}
	return rebateSet, b, err
}

//查询会员返佣设定列表
func (*MemberRebateSetBean) SelectAll() (rebateSetAll []back.MemberRebateSet, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebateSchema := &schema.MemberRebateSet{}
	err = sess.Table(rebateSchema.TableName()).Where("delete_time >0").Find(&rebateSetAll)
	return
}

//查询会员返佣设定列表 (按照有效投注倒序)
func (*MemberRebateSetBean) GetListBySite(siteId, siteIndexId string) (rebateSets []*back.MemberRebateSet, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebateSetSchema := &schema.MemberRebateSet{}
	sess.Table(rebateSetSchema.TableName())
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	if siteIndexId != "" {
		sess.Where("site_index_id = ?", siteIndexId)
	}
	sess.Where("delete_time = 0")
	sess.Desc("valid_money")
	sess.Find(&rebateSets)
	return
}

//查询会员优惠设定详情,用来补充到[]*back.MemberRebateSet 上
func (m *MemberRebateSetBean) GetRebateProductSetByRebateSetIds(rebateSetIds []int64) (
	[]*back.MemberRebateProduct, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var rebateProducts []*back.MemberRebateProduct
	rebateSetSchema := &schema.MemberRebateProduct{}
	product := new(schema.Product)
	err := sess.Table(rebateSetSchema.TableName()).
		Join("INNER", product.TableName(), rebateSetSchema.TableName()+".product_id = "+product.TableName()+".id").
		Select("set_id,product_id,v_type,product_name,rate").
		In("set_id", rebateSetIds).
		Find(&rebateProducts)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return rebateProducts, err
	}
	return rebateProducts, err
}

//查询会员优惠设定列表以及详情 (按照有效投注倒序)
func (m *MemberRebateSetBean) GetAll(siteId, siteIndexId string, products []*schema.Product) ([]*back.MemberRebateSet, error) {
	//查询设置列表
	var rebateSets []*back.MemberRebateSet
	rebateSets, err := m.GetListBySite(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return rebateSets, err
	}
	if len(rebateSets) == 0 {
		return rebateSets, err
	}
	var rebateSetIds []int64 //收集设置的id
	temp := make(map[int64]*back.MemberRebateSet)
	for i, rebateSet := range rebateSets {
		rebateSetIds = append(rebateSetIds, rebateSet.Id)
		productMap := make(map[string]*back.MemberRebateProduct)
		for _, product := range products {
			productMap[product.VType] = &back.MemberRebateProduct{VType: product.VType, ProductName: product.ProductName}
		}
		rebateSets[i].ProductRates = &productMap
		temp[rebateSet.Id] = rebateSets[i]
	}
	//js, _ := json.Marshal(products)
	//fmt.Println("所有返佣比率:", string(js))

	//查询每个商品对应的优惠比例
	rebateProductSets, err := m.GetRebateProductSetByRebateSetIds(rebateSetIds)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return rebateSets, err
	}
	if len(rebateProductSets) == 0 {
		err = global.GlobalLogger.Error("A single product corresponding to the discount ratio does not exist")
		return rebateSets, err
	}
	//js, _ = json.Marshal(rebateProductSets)
	//fmt.Println("每个商品对应的优惠比例:", string(js))
	for i, rebateProductSet := range rebateProductSets {
		rebateSet, ok := temp[rebateProductSet.SetId]
		if !ok {
			//global.GlobalLogger.Info("不存在优惠设置 %d ", rebateProductSet.SetId)
			continue
		}
		//如果有这个商品(没有被剔除的),才会给赋值返佣比例
		_, ok = (*rebateSet.ProductRates)[rebateProductSet.VType]
		if ok {
			(*rebateSet.ProductRates)[rebateProductSet.VType] = rebateProductSets[i]
		}
	}
	//js, _ = json.Marshal(rebateSets)
	//fmt.Println("每个商品对应的优惠比例:", string(js))
	return rebateSets, err
}
