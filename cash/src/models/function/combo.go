package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type ComboBeen struct{}

//套餐商品列表
func (*ComboBeen) GetProductList(this *input.ComboList, listparam *global.ListParams) (data []back.ComboPro, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	if this.Status != 0 {
		sess.Where("sales_combo.status = ?", this.Status)
	}
	if this.ComboName != "" {
		sess.Where("sales_combo.combo_name = ?", this.ComboName)
	}
	sess.Where("sales_combo.delete_time = ?", 0)
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	site := new(schema.Site) //站点表
	sql2 := fmt.Sprintf("%s.id = %s.combo_id", combo.TableName(), site.TableName())
	//重新传入表名和where条件查询记录
	err = sess.Table(combo.TableName()).Select("sales_combo.id,sales_combo.combo_name,sales_combo.status,sales_combo.create_time,count(sales_site.combo_id) as num").Join("LEFT", site.TableName(), sql2).GroupBy("sales_combo.id").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(combo.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//套餐列表
func (*ComboBeen) GetList(this *input.ComboList, listparam *global.ListParams) (data []back.Combo, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	if this.Status != 0 {
		sess.Where("sales_combo.status = ?", this.Status)
	}
	if this.ComboName != "" {
		sess.Where("sales_combo.combo_name = ?", this.ComboName)
	}
	sess.Where("sales_combo.delete_time = ?", 0)
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	site := new(schema.Site) //站点表
	sql2 := fmt.Sprintf("%s.id = %s.combo_id", combo.TableName(), site.TableName())
	//重新传入表名和where条件查询记录
	err = sess.Table(combo.TableName()).Select("sales_combo.id,sales_combo.combo_name,sales_combo.status,sales_combo.create_time,count(sales_site.combo_id) as num").Join("LEFT", site.TableName(), sql2).GroupBy("sales_combo.id").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	//获得符合条件的记录数
	count, err = sess.Table(combo.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//添加套餐名称
func (*ComboBeen) Add(this *input.ComboAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	combo.ComboName = this.ComboName
	combo.Status = 1
	combo.CreateTime = time.Now().Unix()
	count, err = sess.Table(combo.TableName()).InsertOne(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//套餐详情
func (*ComboBeen) GetInfo(this *input.ComboId) (c back.ComboInfo, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	has, err = sess.Table(combo.TableName()).Where("id = ?", this.Id).Where("delete_time = ?", 0).Get(&c)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return c, has, err
	}
	return c, has, err
}

//修改套餐名称
func (*ComboBeen) UpdateInfo(this *input.ComboEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	combo.ComboName = this.ComboName
	count, err = sess.Table(combo.TableName()).Where("id = ?", this.Id).Cols("combo_name").Update(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改套餐状态
func (*ComboBeen) UpdateStatus(this *input.ComboStatus) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	combo.Status = this.Status
	count, err = sess.Where("id = ?", this.Id).Cols("status").Update(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除套餐
func (*ComboBeen) Delete(this *input.ComboId) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	combo.Id = this.Id
	combo.DeleteTime = time.Now().Unix()
	count, err = sess.Table(combo.TableName()).Where("id = ?", this.Id).Cols("delete_time").Update(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看套餐名（添加）
func (*ComboBeen) GetComboName(comboName string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	has, err = sess.Table(combo.TableName()).Where("combo_name = ?", comboName).Where("delete_time = ?", 0).Get(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看套餐名（修改）
func (*ComboBeen) GetComboNames(id int64, comboName string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	has, err = sess.Table(combo.TableName()).Where("id != ?", id).Where("combo_name = ?", comboName).Where("delete_time = ?", 0).Get(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//套餐详情
func (*ComboBeen) GetComboNameById(id int64) (combo *schema.Combo, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo = new(schema.Combo)
	has, err = sess.Table(combo.TableName()).Where("id = ?", id).Select("combo_name").Where("delete_time = ?", 0).Get(combo)
	return
}

//查看套餐商品
func (*ComboBeen) GetProductInfo(this *input.ComboProductId) (cp []back.ComboProduct, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	comboPro := new(schema.ComboProduct)
	cp = make([]back.ComboProduct, 0)
	product := new(schema.Product)
	sql := fmt.Sprintf("%s.product_id = %s.id", comboPro.TableName(), product.TableName())
	err = sess.Table(comboPro.TableName()).Where("combo_id = ?", this.ComboId).Join("LEFT", product.TableName(), sql).Find(&cp)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return cp, err
	}
	return cp, err
}

//根据商品类型名查看平台
func (*ComboBeen) GetProductByType(this *input.ProductTypeName) (cp []back.ComboProduct, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	comboPro := new(schema.ComboProduct)
	cp = make([]back.ComboProduct, 0)
	platform := new(schema.Platform)
	sess.Where(fmt.Sprintf("%s.status = 1", platform.TableName()))
	sql := fmt.Sprintf("%s.platform_id = %s.id", comboPro.TableName(), platform.TableName())
	err = sess.Table(comboPro.TableName()).Where("combo_id = ?", this.ComboId).Join("LEFT", platform.TableName(), sql).Find(&cp)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return cp, err
	}
	return cp, err
}

//添加套餐商品占比
func (*ComboBeen) AddProductProportion(this *input.ComboProducts) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var comboPro schema.ComboProduct
	var comboPros []schema.ComboProduct
	combo := new(schema.ComboProduct)
	comboPro.ComboId = this.ComboId
	sess.Begin()
	//删除套餐id下的所有数据
	count, err = sess.Table(comboPro.TableName()).Where("combo_id = ?", this.ComboId).Delete(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	for k := range this.Params {
		comboPro.PlatformId = this.Params[k].PlatformId
		comboPro.ProductId = this.Params[k].ProductId
		comboPro.Proportion = this.Params[k].Proportion
		comboPros = append(comboPros, comboPro)
	}
	//添加数据
	count, err = sess.Table(combo.TableName()).Insert(comboPros)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//查看商品id是否被哪个套餐所使用
func (*ComboBeen) GetProductId(productId int64) (has bool, err error) {
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

//根据套餐id获取套餐名称
func (*ComboBeen) GetNameByID(id int64) (info schema.Combo, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("delete_time=?", 0)
	sess.Where("id=?", id)
	flag, err = sess.Table(info.TableName()).Get(&info)
	return
}

//套餐下拉框
func (*ComboBeen) GetComboDrop(this *schema.Combo) (data []back.ComboDrop, err error) {
	sess := global.GetXorm().NewSession().Table(this.TableName())
	defer sess.Close()
	err = sess.Where("delete_time=?", 0).Where("status=?", 1).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//套餐类型模糊搜索
func (*ComboBeen) GetProductType(this *input.ProductTypeName) (data []back.AllProductClassifyListBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	productType := new(schema.ProductType)
	Product := new(schema.Product)
	var pro_type []schema.ProductType
	var da back.AllProductClassifyListBack
	err = sess.Table(productType.TableName()).Where("title like ?", "%"+this.Name+"%").Find(&pro_type)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if pro_type == nil {
		return nil, err
	}
	var id []int64
	for _, v := range pro_type {
		id = append(id, v.Id)
		da.Id = v.Id
		da.Title = v.Title
		data = append(data, da)
	}
	var pro_duct []back.PlatformlistBack
	platform := new(schema.Platform)
	sql := fmt.Sprintf("%s.platform_id = %s.id", Product.TableName(), platform.TableName())
	err = sess.Table(Product.TableName()).Join("LEFT", platform.TableName(), sql).In("sales_product.type_id", id).Where("sales_product.delete_time=?", 0).Where("sales_product.status=?", 1).Find(&pro_duct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	for k := range data {
		for l := range pro_duct {
			if data[k].Id == pro_duct[l].TypeId {
				data[k].Children = append(data[k].Children, pro_duct[l])
			}
		}
	}
	return data, err
}

//检验域名是否被使用
func (*ComboBeen) CheckThirdDoMain(domain string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var cou int64
	sD := new(schema.SiteDomain)
	has, err := sess.Where("domain=?", domain).Get(sD)
	if has && err == nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		cou = 60015
		return cou, err
	}
	return cou, err
}

//检验域名是否被使用(修改)
func (*ComboBeen) CheckThirdDoMainChange(siteId, indexId, domain string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var cou int64
	sD := new(schema.SiteDomain)
	has, err := sess.Where("domain=?", domain).
		Where("site_id!=?", siteId).
		Where("site_index_id!=?", indexId).
		Get(sD)
	if has && err == nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		cou = 60015
		return cou, err
	}
	return cou, err
}
