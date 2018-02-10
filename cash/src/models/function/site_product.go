package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SiteProductBean struct{}

//该站点下套餐中的商品列表
func (*SiteProductBean) ProductList(comboId int64) (siteMaintenance []back.SiteProduct, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	productType := new(schema.ProductType)
	comboProduct := new(schema.ComboProduct)
	sql2 := fmt.Sprintf("%s.product_id=%s.id", comboProduct.TableName(), product.TableName())
	sql := fmt.Sprintf("%s.type_id = %s.id", product.TableName(), productType.TableName())
	err = sess.Table(product.TableName()).Join("LEFT", productType.TableName(), sql).
		Join("LEFT", comboProduct.TableName(), sql2).Where(comboProduct.TableName()+
		".combo_id=?", comboId).Where(product.TableName() + ".delete_time=0").Find(&siteMaintenance)
	return
}

//站点模块管理（站点商品剔除）
func (sp *SiteProductBean) SiteProductUpdate(this *input.SiteProductEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteProductDel := new(schema.SiteProductDel)
	sess.Begin()
	//删除剔除表中该站点下数据
	count, err = sess.Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Delete(siteProductDel)
	if err != nil {
		sess.Rollback()
		return
	}
	if len(this.ProductId) != 0 {
		var spdl schema.SiteProductDel
		spd := make([]schema.SiteProductDel, 0)
		for _, k := range this.ProductId {
			spdl.SiteId = this.SiteId
			spdl.SiteIndexId = this.SiteIndexId
			spdl.ProductId = k
			spd = append(spd, spdl)
		}
		//批量添加数据
		count, err = sess.Table(siteProductDel.TableName()).Insert(&spd)
		if err != nil || count == 0 {
			sess.Rollback()
			return
		}
	}
	sess.Commit()
	return
}

//查看站点是否存在
func (*SiteProductBean) IsExistSite(siteId, siteIndexId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	has, err = sess.Where("id=?", siteId).Where("index_id=?", siteIndexId).
		Where("delete_time=0").Get(site)
	return
}

//查看站点下的套餐id
func (*SiteProductBean) GetSiteCombo(siteId, siteIndexId string) (comboId int64, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	has, err = sess.Where("id=?", siteId).Where("index_id=?", siteIndexId).
		Where("delete_time=0").Select("combo_id").Get(site)
	comboId = site.ComboId
	return
}

//站点下的商品剔除
func (*SiteProductBean) SiteProductDel(siteId, siteIndexId string) (productIds []int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spd := new(schema.SiteProductDel)
	var siteProduct []back.SiteProductDelBack
	sess.Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId)
	err = sess.Table(spd.TableName()).Find(&siteProduct)
	for k := range siteProduct {
		productIds = append(productIds, siteProduct[k].ProductId)
	}
	return
}

//查看商品是否存在
func (*SiteProductBean) IsExistProduct(productId []int64, comboId int64) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	comboProduct := new(schema.ComboProduct)
	if comboId != 0 {
		sess.Where("combo_id=?", comboId)
	}
	count, err = sess.Table(comboProduct.TableName()).In("product_id", productId).Count()
	return
}

//套餐商品中查出productIds以外的商品id
func (*SiteProductBean) GetProductId(productId []int64, comboId int64) (ids []int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var comboProduct []schema.ComboProduct
	if comboId != 0 {
		sess.Where("combo_id=?", comboId)
	}
	err = sess.NotIn("product_id", productId).Select("product_id").Find(&comboProduct)
	for k := range comboProduct {
		ids = append(ids, comboProduct[k].ProductId)
	}
	return
}

//获取商品表
func (*SiteProductBean) GetProductList() (Products []back.ProductVideo, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.Product)
	err = sess.Table(infoLogoSchema.TableName()).
		Where("delete_time = ?", 0).
		Where("status = ?", 1).
		Where("type_id = ?", 1). //视讯类型
		Select("id,product_name,v_type").
		Find(&Products)
	if err != nil {
		return
	}
	return
}

//获取所有站点商品
func (*SiteProductBean) GetProductAll(DelId []int) (Products []back.ProductlistRep, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.Product)
	err = sess.Table(infoLogoSchema.TableName()).
		Where("delete_time = ?", 0).
		Where("status = ?", 1).
		NotIn("id", DelId).
		Select("id,product_name,v_type").
		Find(&Products)
	if err != nil {
		return
	}
	return
}

//获取剔除表数据
func (*SiteProductBean) GetProductDel(siteId, siteIndexId string) (ids []int, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.SiteProductDel)
	err = sess.Table(infoLogoSchema.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Select("product_id").
		Find(&ids)
	if err != nil {
		return
	}
	return
}
