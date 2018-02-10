package function

import (
	"global"
	"models/back"
	"models/schema"
)

type OtherBean struct{}

//导航栏根据商品类型或许商品
func (*OtherBean) GetProductByType(typeId, productIds []int64) (data []back.OtherProduct, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	products := make([]back.Product, 0)
	var op back.OtherProduct
	err = sess.Table(product.TableName()).In("type_id", typeId).NotIn("id", productIds).Find(&products)
	for k := range typeId {
		op.TypeId = typeId[k]
		data = append(data, op)
	}
	for i := range data {
		for k := range products {
			if data[i].TypeId == products[k].TypeId {
				data[i].ProductName = append(data[i].ProductName, products[k].ProductName)
			}
		}
	}
	return
}

//获取商品剔除
func (*OtherBean) GetProductDel(siteId, siteIndexId string) (ids []int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteProductDel := new(schema.SiteProductDel)
	spd := make([]schema.SiteProductDel, 0)
	sess.Table(siteProductDel.TableName()).Select("product_id").Where("site_id = ?", siteId).Where("site_index_id = ?", siteIndexId).Find(&spd)
	for k := range spd {
		ids = append(ids, spd[k].ProductId)
	}
	return
}
