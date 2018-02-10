package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type ProductTypeBean struct{}

//添加商品分类
func (*ProductTypeBean) Add(this *input.ProductType) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProductType := &schema.ProductType{
		Title:  this.Title,
		Status: this.Status}
	count, err = sess.Table(schemaProductType.TableName()).Insert(schemaProductType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//通过商品名称获取商品分类详情
func (*ProductTypeBean) GetInfo(title string) (*back.ProductType, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProductType := new(schema.ProductType)
	sess.Where("title=?", title)
	sess.Where("delete_time=?", 0)
	backProductType := new(back.ProductType)
	have, err := sess.Table(schemaProductType.TableName()).Get(backProductType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backProductType, have, err
	}
	return backProductType, have, err
}

//通过商品Id获取获取商品分类详情
func (*ProductTypeBean) GetInfoById(id int64) (*back.ProductType, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProductType := new(schema.ProductType)
	sess.Where("id=?", id)
	sess.Where("delete_time=?", 0)
	backProductType := new(back.ProductType)
	have, err := sess.Table(schemaProductType.TableName()).Get(backProductType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backProductType, have, err
	}
	return backProductType, have, err
}

//更新商品分类信息
func (*ProductTypeBean) Update(this *input.ProductTypeEdit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", this.Id)
	sess.Where("delete_time=?", 0)
	schemaProductType := new(schema.ProductType)
	schemaProductType.Status = this.Status
	schemaProductType.Title = this.Title
	count, err := sess.Table(schemaProductType.TableName()).Cols("title", "status").Update(schemaProductType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//更新商品分类的状态
func (*ProductTypeBean) Status(this *input.ProductTypeStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", this.Id)
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	schemaProductType := new(schema.ProductType)
	schemaProductTypeInfo := new(schema.ProductType)
	_, err := sess.Table(schemaProductType.TableName()).Get(schemaProductTypeInfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if schemaProductTypeInfo.Status == 1 {
		schemaProductType.Status = 2
	}
	if schemaProductTypeInfo.Status == 2 {
		schemaProductType.Status = 1
	}
	count, err := sess.Table(schemaProductType.TableName()).Where(conds).Cols("status").Update(schemaProductType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除商品分类
func (*ProductTypeBean) Delete(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", id)
	sess.Where("delete_time=?", 0)
	schemaProductType := new(schema.ProductType)
	schemaProductType.DeleteTime = time.Now().Unix()
	count, err := sess.Table(schemaProductType.TableName()).Cols("delete_time").Update(schemaProductType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//商品分类下拉框
func (*ProductTypeBean) List() ([]*back.ProductTypeList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	schemaProductType := new(schema.ProductType)
	list := make([]*back.ProductTypeList, 0)
	err := sess.Table(schemaProductType.TableName()).Find(&list)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	return list, err
}

//商品分类列表
func (*ProductTypeBean) ProductTypeList(this *input.ProductTypeList, listParams *global.ListParams) (data []back.ProductTypes, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	productType := new(schema.ProductType)
	sess.Where(productType.TableName()+".delete_time=?", 0)
	//根据商品分类名称查找
	if this.Title != "" {
		sess.Where(productType.TableName()+".title=?", this.Title)
	}
	conds := sess.Conds()
	listParams.Make(sess)
	product := new(schema.Product)
	sql := fmt.Sprintf("%s.id = %s.type_id", productType.TableName(), product.TableName())
	sess.Select(productType.TableName() + ".id," + productType.TableName() + ".title," + productType.TableName() +
		".status," + "count(sales_product.type_id) as product_count").GroupBy(productType.TableName() + ".id")
	err = sess.Table(productType.TableName()).Join("LEFT", product.TableName(), sql).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(productType.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//商品分类下是否有商品
func (*ProductTypeBean) ExistProduct(typeId int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProduct := new(schema.Product)
	sess.Where("type_id=?", typeId)
	sess.Where("delete_time=?", 0)
	have, err := sess.Table(schemaProduct.TableName()).Get(schemaProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return have, err
	}
	return have, err
}

//查询商品分类和商品
func (*ProductTypeBean) GetAllProductByType(this *schema.Product) (data []back.AllProductClassifyListBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaProductType := new(schema.ProductType)
	var proType []schema.ProductType
	err = sess.Table(schemaProductType.TableName()).Where("delete_time=?", 0).Where("status=?", 1).Find(&proType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if proType == nil {
		return nil, err
	}
	var ids []int64
	for k := range proType {
		ids = append(ids, proType[k].Id)
	}
	var pro_duct []back.PlatformlistBack
	platform := new(schema.Platform)
	Product := new(schema.Product)
	sql := fmt.Sprintf("%s.platform_id = %s.id", Product.TableName(), platform.TableName())
	err = sess.Table(Product.TableName()).Join("LEFT", platform.TableName(), sql).In("sales_product.type_id", ids).Where("sales_product.delete_time=?", 0).Where("sales_product.status=?", 1).Find(&pro_duct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	for i := range proType {
		var da back.AllProductClassifyListBack
		for k := range pro_duct {
			da.Id = proType[i].Id
			da.Title = proType[i].Title
			var p back.PlatformlistBack
			if proType[i].Id == pro_duct[k].TypeId {
				p.Id = pro_duct[k].Id
				p.TypeId = pro_duct[k].TypeId
				p.ProductId = pro_duct[k].ProductId
				p.Platform = pro_duct[k].Platform
				p.Proportion = pro_duct[k].Proportion
				p.IsProduct = pro_duct[k].IsProduct
				da.Children = append(da.Children, p)
			}
		}
		data = append(data, da)
	}
	return data, err
}

//商品下拉框
func (*ProductTypeBean) ProductListBeanDrop() ([]back.ProductListDropBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	pro := new(schema.Product)
	var data []back.ProductListDropBack
	err := sess.Table(pro.TableName()).Where("status=?", 1).
		Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
