package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strconv"
	"time"
)

type ProductBean struct{}

//添加商品
func (*ProductBean) Add(this *input.Product) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProduct := &schema.Product{
		ProductName: this.ProductName,
		TypeId:      this.TypeId,
		Status:      this.Status,
		Api:         this.Api,
		PlatformId:  this.PlatformId,
		VType:       this.VType}
	count, err = sess.Table(schemaProduct.TableName()).Insert(schemaProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//通过商品名称和商品分类名称判断是否存在该条纪录
func (*ProductBean) Exist(typeId int64, productName string) (*schema.Product, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProduct := new(schema.Product)
	sess.Where("product_name=?", productName)
	sess.Where("type_id=?", typeId)
	sess.Where("delete_time=?", 0)
	have, err := sess.Table(schemaProduct.TableName()).Get(schemaProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return schemaProduct, have, err
	}
	return schemaProduct, have, err
}

//获取单个商品信息
func (*ProductBean) GetInfo(this *input.ProductInfo) (*back.ProductInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProduct := new(schema.Product)
	sess.Where(schemaProduct.TableName()+".delete_time=?", 0)
	sess.Where(schemaProduct.TableName()+".id=?", this.ProductId)
	data := new(back.ProductInfo)
	schemaProductType := new(schema.ProductType)
	where1 := fmt.Sprintf("%s.type_id = %s.id", schemaProduct.TableName(), schemaProductType.TableName())
	have, err := sess.Table(schemaProduct.TableName()).Join("LEFT", schemaProductType.TableName(), where1).Get(data)
	if err != nil {
		return data, have, err
	}
	return data, have, err
}

//通过商品Id获取商品信息
func (*ProductBean) GetInfoById(productId int64) (*schema.Product, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProduct := new(schema.Product)
	sess.Where("id=?", productId)
	sess.Where("delete_time=?", 0)
	have, err := sess.Table(schemaProduct.TableName()).Get(schemaProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return schemaProduct, have, err
	}
	return schemaProduct, have, err
}

//更新商品信息
func (*ProductBean) Update(this *input.ProductEdit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProduct := schema.Product{
		ProductName: this.ProductName,
		Status:      this.Status,
		Api:         this.Api,
		TypeId:      this.TypeId}
	sess.Where("delete_time=?", 0)
	sess.Where("id=?", this.ProductId)
	count, err := sess.Table(schemaProduct.TableName()).Cols("product_name", "status", "api", "type_id").Update(schemaProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//启用或禁用商品
func (*ProductBean) Status(this *input.ProductStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProduct := new(schema.Product)
	sess.Where("id=?", this.ProductId)
	sess.Where("delete_time=?", 0)
	schemaProductInfo := new(schema.Product)
	conds := sess.Conds()
	_, err := sess.Table(schemaProduct.TableName()).Get(schemaProductInfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if schemaProductInfo.Status == 1 {
		schemaProduct.Status = 2
	}
	if schemaProductInfo.Status == 2 {
		schemaProduct.Status = 1
	}
	count, err := sess.Table(schemaProduct.TableName()).Where(conds).Cols("status").Update(schemaProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除商品(DELETE)
func (*ProductBean) Delete(productId int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProduct := new(schema.Product)
	schemaProduct.DeleteTime = time.Now().Unix()
	sess.Where("id=?", productId)
	sess.Where("delete_time=?", 0)
	count, err := sess.Table(schemaProduct.TableName()).Cols("delete_time").Update(schemaProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取商品列表
func (*ProductBean) List(this *input.ProductList, listparam *global.ListParams) (data []*back.ProductList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//获得分页记录
	listparam.Make(sess)
	product := new(schema.Product)
	sess.Where(product.TableName()+".delete_time=?", 0)
	//根据商品Id查找
	if this.ProductId != 0 {
		sess.Where(product.TableName()+".id=?", this.ProductId)
	}
	//根据商品名称查找
	if this.ProductName != "" {
		sess.Where(product.TableName()+".product_name=?", this.ProductName)
	}
	//根据商品状态进行查找
	if this.Status != 0 {
		sess.Where(product.TableName()+".status=?", this.Status)
	}
	productType := new(schema.ProductType)
	//根据商品类型进行查找
	if this.Title != "" {
		sess.Where(productType.TableName()+".title=?", this.Title)
	}
	conds := sess.Conds()
	where1 := fmt.Sprintf("%s.type_id = %s.id", product.TableName(), productType.TableName())
	err = sess.Table(product.TableName()).Join("LEFT", productType.TableName(), where1).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(product.TableName()).Join("LEFT", productType.TableName(), where1).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取商品列表(关联剔除表)
func (m *ProductBean) GetList(siteId, siteIndexId string, this *input.ProductList) ([]*schema.Product, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []*schema.Product
	productSchema := new(schema.Product)
	productDelSchema := new(schema.SiteProductDel)
	sql := "SELECT * FROM " + productSchema.TableName() + " WHERE delete_time = 0"
	var params []interface{}
	//根据商品Id查找
	if this.ProductId != 0 {
		sql += " AND id = ?"
		params = append(params, this.ProductId)
	}
	//根据商品名称查找
	if this.ProductName != "" {
		sql += " AND product_name = ?"
		params = append(params, this.ProductName)
	}
	//根据商品状态进行查找
	if this.Status != 0 {
		sql += " AND status = ?"
		params = append(params, this.Status)
	}
	sql += " AND id NOT IN (SELECT product_id FROM " + productDelSchema.TableName() +
		" WHERE site_id = ? AND site_index_id = ?)"
	params = append(params, siteId, siteIndexId)
	err := sess.SQL(sql, params...).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取商品列表(关联剔除表并合并同一平台的商品)
func (m *ProductBean) GetListOnPlatform(siteId, siteIndexId string, this *input.ProductList) (data []*schema.Product, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	productSchema := new(schema.Product)
	productDelSchema := new(schema.SiteProductDel)
	sql := "SELECT * FROM " + productSchema.TableName() + " WHERE delete_time = 0"
	var params []interface{}
	//根据商品Id查找
	if this.ProductId != 0 {
		sql += " AND id = ?"
		params = append(params, this.ProductId)
	}
	//根据商品名称查找
	if this.ProductName != "" {
		sql += " AND product_name = ?"
		params = append(params, this.ProductName)
	}
	//根据商品状态进行查找
	if this.Status != 0 {
		sql += " AND status = ?"
		params = append(params, this.Status)
	}
	sql += " AND id NOT IN (SELECT product_id FROM " + productDelSchema.TableName() +
		" WHERE site_id = ? AND site_index_id = ?) group by platform_id"
	params = append(params, siteId, siteIndexId)
	err = sess.SQL(sql, params...).Find(&data)
	return
}

//获取额度转换平台列表(关联剔除表)
//func (m *ProductBean) SitePlatform(siteId, siteIndexId string) (data []back.TypeList, err error) {
//	sess := global.GetXorm().NewSession()
//	defer sess.Close()
//	productSchema := new(schema.Product)
//	productDelSchema := new(schema.SiteProductDel)
//	sess.Where(productSchema.TableName()+".delete_time = ?", 0)
//	//sess.Where(productSchema.TableName()+".site_id = ?",siteId)
//	//sess.Where(productSchema.TableName()+".site_index_id = ?",siteIndexId)
//	sess.Where(productSchema.TableName()+".status = ?", 1)
//	sess.NotIn(productSchema.TableName()+".id", "SELECT product_id FROM "+productDelSchema.TableName()+
//		" WHERE site_id = "+siteId+" AND site_index_id = "+siteIndexId)
//	sess.Join("LEFT", "sales_platform", productSchema.TableName()+".platform_id = sales_platform.id")
//	sess.GroupBy("platform_id")
//	sess.Select("sales_platform. platform,sales_platform.id")
//	err = sess.Table(productSchema.TableName()).Find(&data)
//	fmt.Println("平台数据******************************",data)
//	return
//}
//获取额度转换平台列表(关联剔除表)
func (m *ProductBean) SitePlatform(siteId, siteIndexId string) (data []back.TypeList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	//查詢剔除的商品
	productDelSchema := new(schema.SiteProductDel)
	delProductId := []int64{}
	productDelList := []schema.SiteProductDel{}
	_ = sess.Table(productDelSchema.TableName()).
		Where("site_id = ? AND site_index_id = ?", siteId, siteIndexId).
		Find(&productDelList)
	if len(productDelList) > 0 {
		for _, value := range productDelList {
			delProductId = append(delProductId, value.ProductId)
		}
	}
	//平台列表數據
	platformList, err := PlatformList()
	if err != nil {
		return
	}
	//查詢可用平台
	productSchema := new(schema.Product)
	sess.Where("delete_time = ?", 0)
	sess.Where("status = ?", 1)
	sess.NotIn("id", delProductId)
	sess.GroupBy("platform_id")
	sess.Select("platform_id id")
	err = sess.Table(productSchema.TableName()).Find(&data)
	if err != nil {
		return
	}
	for key, value := range data {
		for _, v := range platformList {
			if value.Id == v.Id {
				data[key].Platform = v.Platform
			}
		}
	}
	fmt.Println("平台数据******************************", data)
	return
}

//获取交易平台列表（转出/转入项目）
func (*ProductBean) TypeList() ([]*back.TypeList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Platform)
	var data []*back.TypeList
	err := sess.Table(platform.TableName()).Where("status=?", 1).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return data, err
	}
	return data, err
}

//查询与vType对应的Product,Name
func (*ProductBean) ProductList(vTypes ...[]string) (data []back.ProductlistRep, err error) {
	var vType []string
	if len(vTypes) > 0 {
		vType = vTypes[0]
	}

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	if len(vType) > 0 {
		sess.In("v_type", vTypes)
	}
	err = sess.Table(product.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取wap商品列表(关联剔除表)
func (m *ProductBean) GetProductList(siteId, siteIndexId string) (data []schema.Product, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	productSchema := new(schema.Product)
	productDelSchema := new(schema.SiteProductDel)
	var delData []back.SiteProductDelBack
	err = sess.Table(productDelSchema.TableName()).
		Select("product_id").
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Find(&delData)
	var notInStr string
	for k, v := range delData {
		if k == (len(delData) - 1) {
			notInStr += strconv.FormatInt(v.ProductId, 10)
		} else {
			notInStr += strconv.FormatInt(v.ProductId, 10) + ","
		}
	}
	err = sess.Table(productSchema.TableName()).
		Where("status=?", 1).
		Where("delete_time=?", 0).
		NotIn("id", notInStr).
		Find(&data)
	return
}

//商品分类下拉框
func (*ProductBean) GetProductTypeList() (list []back.ProductTypeList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	schemaProductType := new(schema.ProductType)
	err = sess.Table(schemaProductType.TableName()).Find(&list)
	return
}

//根据5大模块id获取游戏列表
func (m *ProductBean) GetGameList(siteId, siteIndexId string, typeId int64) (data []schema.Product, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	productSchema := new(schema.Product)
	productDelSchema := new(schema.SiteProductDel)
	var delData []back.SiteProductDelBack

	err = sess.Table(productDelSchema.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Find(&delData)
	var notInStr string
	for k, v := range delData {
		if k == (len(delData) - 1) {
			notInStr += strconv.FormatInt(v.ProductId, 10)
		} else {
			notInStr += strconv.FormatInt(v.ProductId, 10) + ","
		}
	}
	if typeId != 0 {
		sess.Where("type_id=?", typeId)
	}
	err = sess.Table(productSchema.TableName()).
		Where("status=?", 1).
		Where("delete_time=?", 0).
		NotIn("id", notInStr).
		Find(&data)
	return
}

//获取平台列表
func PlatformList() (data []schema.Platform, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	platform := new(schema.Platform)
	sess.Where("delete_time = ?", 0)
	sess.Where("status = ?", 1)
	err = sess.Table(platform.TableName()).Find(&data)
	return
}
