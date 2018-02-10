package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type NoteModuleBean struct{}

//添加商品分类
func (*NoteModuleBean) Add(this *input.ProductType) (count int64, err error) {
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
func (*NoteModuleBean) GetInfo(title string) (*back.ProductType, bool, error) {
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
func (*NoteModuleBean) GetInfoById(id int64) (*back.ProductType, bool, error) {
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

//通过商品Id获取商品信息
func (*NoteModuleBean) GetProductInfoById(productId int64) (*schema.Product, bool, error) {
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

//更新商品分类信息
func (*NoteModuleBean) Update(this *input.ProductTypeEdit) (int64, error) {
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
func (*NoteModuleBean) Status(this *input.ProductTypeStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", this.Id)
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	schemaProductType := new(schema.ProductType)
	schemaProductTypeInfo := new(schema.ProductType)
	_, err := sess.Table(schemaProductType.TableName()).Get(schemaProductTypeInfo)
	if err != nil {
		return 0, err
	}
	if schemaProductTypeInfo.Status == 1 {
		schemaProductType.Status = 2
	}
	if schemaProductTypeInfo.Status == 2 {
		schemaProductType.Status = 1
	}
	count, err := sess.Table(schemaProductType.TableName()).Where(conds).Cols("status").Update(schemaProductType)
	return count, err
}

//删除商品分类
func (*NoteModuleBean) Delete(id int64) (int64, error) {
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
func (*NoteModuleBean) List() ([]*back.ProductTypeList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	schemaProductType := new(schema.ProductType)
	list := make([]*back.ProductTypeList, 0)
	err := sess.Table(schemaProductType.TableName()).Find(&list)
	return list, err
}

//商品分类列表
func (*NoteModuleBean) ProductTypeList(this *input.ProductTypeList) (data []back.ProductTypes, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("delete_time=?", 0)
	//根据商品分类名称查找
	if this.Title != "" {
		sess.Where("title=?", this.Title)
	}
	sess.Select("id,title,status")
	productType := new(schema.ProductType)
	err = sess.Table(productType.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	var data1 []back.ProductTypes
	product := new(schema.Product)
	err = sess.Table(product.TableName()).Select("count(*) as product_count,type_id").Where("delete_time=?", 0).GroupBy("type_id").Find(&data1)
	for k, v := range data {
		for k1, v1 := range data1 {
			if v.Id == v1.TypeId {
				data[k].ProductCount = data1[k1].ProductCount
				data[k].TypeId = data1[k1].TypeId
			}
		}
	}
	return
}

//商品分类下是否有商品
func (*NoteModuleBean) ExistProduct(typeId int64) (bool, error) {
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
func (*NoteModuleBean) GetAllProductByType(this *schema.Product) (data []back.AllProductClassifyListBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	schemaProductType := new(schema.ProductType)
	var proType []schema.ProductType
	err = sess.Table(schemaProductType.TableName()).Where("delete_time=?", 0).Where("status=?", 1).Find(&proType)
	if err != nil {
		sess.Rollback()
		return
	}
	if proType == nil {
		return nil, err
	}

	var proDuct []back.PlatformlistBack
	Product := new(schema.Product)
	err = sess.Table(Product.TableName()).
		Where("delete_time=?", 0).
		Find(&proDuct)
	if err != nil {
		return
	}
	platform := new(schema.Platform)
	var platList []schema.Platform
	err = sess.Table(platform.TableName()).
		Where("delete_time=?", 0).
		Where("status=?", 1).
		Find(&platList)
	for k, v := range proDuct {
		for k1, v1 := range platList {
			if v.PlatformId == v1.Id {
				proDuct[k].Platform = platList[k1].Platform
			}
		}
	}
	for i := range proType {
		var da back.AllProductClassifyListBack
		for k := range proDuct {
			da.Id = proType[i].Id
			da.Title = proType[i].Title
			var p back.PlatformlistBack
			if proType[i].Id == proDuct[k].TypeId {
				p.Id = proDuct[k].Id
				p.TypeId = proDuct[k].TypeId
				p.ProductId = proDuct[k].ProductId
				p.Platform = proDuct[k].Platform
				p.Proportion = proDuct[k].Proportion
				p.IsProduct = proDuct[k].IsProduct
				da.Children = append(da.Children, p)
			}
		}
		data = append(data, da)
	}
	sess.Commit()
	return
}

//游戏平台查询
func (*NoteModuleBean) GetGamePlatformList(this *input.GamePlatform) (data []back.Platform, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Platform)
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	if this.Platform != "" {
		sess.Where("platform=?", this.Platform)
	}
	err = sess.Table(platform.TableName()).
		Where("delete_time=?", 0).
		Find(&data)
	if err != nil {
		return data, err
	}
	return data, err
}

//查询平台是否存在
func (*NoteModuleBean) GetPlatformInfo(this *input.PlatformAdd) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Platform)
	if this.Platform != "" {
		sess.Where("platform=?", this.Platform)
	}
	has, err = sess.Table(platform.TableName()).
		Where("delete_time=?", 0).
		Get(platform)
	return
}

//添加游戏平台
func (*NoteModuleBean) AddPlatform(this *input.PlatformAdd) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Platform)
	data, err = sess.Table(platform.TableName()).Insert(this)
	return
}

//修改游戏平台
func (*NoteModuleBean) UpdatePlatform(this *input.PlatformUpdate) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Platform)
	data, err = sess.Table(platform.TableName()).Cols("platform,status").Where("id=?", this.Id).Update(this)

	return
}

//查询平台下是否有正常运行的游戏
func (*NoteModuleBean) GetPlatformGame(this *input.PlatformDelete) (have bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Product)
	have, err = sess.Table(platform.TableName()).
		Where("platform_id=?", this.Id).
		Where("status=?", 1).
		Where("delete_time=?", 0).
		Get(platform)
	return have, err
}

//软删除平台
func (*NoteModuleBean) DeletePlatformGame(this *input.PlatformDelete) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Platform)
	platform.DeleteTime = time.Now().Unix()
	sess.Where("id=?", this.Id)
	sess.Where("platform=?", this.Platform)
	data, err = sess.Table(platform.TableName()).Cols("delete_time").Update(platform)

	return
}

//获取商品列表
func (*NoteModuleBean) ProductList(this *input.ProductList, listparam *global.ListParams) (data []back.ProductList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	sess.Where("delete_time=?", 0)
	//根据商品Id查找
	if this.ProductId != 0 {
		sess.Where("id=?", this.ProductId)
	}
	//根据商品名称查找
	if this.ProductName != "" {
		sess.Where("product_name=?", this.ProductName)
	}
	//根据商品状态进行查找
	if this.Status != 0 {
		sess.Where("status=?", this.Status)
	}
	productType := new(schema.ProductType)
	//根据商品类型进行查找
	if this.Title != "" {
		sess.Where("title=?", this.Title)
	}
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	//where1 := fmt.Sprintf("%s.type_id = %s.id", product.TableName(), productType.TableName())
	//err = sess.Table(product.TableName()).Join("LEFT", productType.TableName(), where1).Find(&data)
	err = sess.Table(product.TableName()).
		Select("id,product_name,api,create_time,status,type_id").
		Find(&data)
	if err != nil {
		return
	}
	var ids []int64
	for k := range data {
		ids = append(ids, data[k].TypeId)
	}
	var productTypeList []schema.ProductType
	err = sess.Table(productType.TableName()).
		Select("id,title").
		In("id", ids).
		Where("delete_time=?", 0).
		Where("status=?", 1).
		Find(&productTypeList)
	if err != nil {
		return
	}
	for k, v := range data {
		for k1, v1 := range productTypeList {
			if v.TypeId == v1.Id {
				data[k].Title = productTypeList[k1].Title
			}
		}
	}
	count, err = sess.Table(product.TableName()).Where(conds).Count()
	return
}

//通过商品名称和商品分类名称判断是否存在该条纪录
func (*NoteModuleBean) Exist(typeId int64, productName string) (*schema.Product, bool, error) {
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

//查看商品id是否被哪个套餐所使用
func (*NoteModuleBean) GetProductId(productId int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.ComboProduct)
	sess.Where("platform_id = ?", productId)
	has, err = sess.Table(combo.TableName()).Get(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//添加商品
func (*NoteModuleBean) AddProduct(this *input.Product) (count int64, err error) {
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

//更新商品信息
func (*NoteModuleBean) UpdateProduct(this *input.ProductEdit) (int64, error) {
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

//删除商品(DELETE)
func (*NoteModuleBean) DeleteProduct(productId int64) (int64, error) {
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
