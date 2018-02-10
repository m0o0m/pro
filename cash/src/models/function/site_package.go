package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SitePackageBean struct {
}

//获取套餐配置列表
func (*SitePackageBean) GetList(this *input.GetList) (data []back.GetList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	comboSchema := new(schema.Combo)
	product := new(schema.Product)
	comboProduct := new(schema.ComboProduct)
	err = sess.Table(comboSchema.TableName()).Alias("c").
		Join("LEFT", []string{comboProduct.TableName(), "cp"}, "c.id=cp.combo_id").
		Join("LEFT", []string{product.TableName(), "p"}, "cp.product_id=p.id").
		Where("c.status=?", 1).Where("c.delete_time=?", 0).
		Select("c.id as combo_id,cp.product_id as product_id,cp.proportion as proportion,p.product_name as product_name,c.combo_name as combo_name").
		Find(&data)
	if err != nil {
		return
	}
	return
}

//新增套餐
func (m *SitePackageBean) PostPackageAdd(this *input.PackAdd) (pack_add back.PackAdd, data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	comboSchema := new(schema.Combo)
	//pack_add := back.PackAdd{}
	pack_add.ComboName = this.ComboName

	data, err = sess.Table(comboSchema.TableName()).InsertOne(&pack_add)

	return
}

//新增套餐下商品比例
func (m *SitePackageBean) AddProduct(list *input.PackAdd) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	comboProduct := new(schema.ComboProduct)
	data, err = sess.Table(comboProduct.TableName()).Insert(list.List)
	return
}

//套餐修改
func (m *SitePackageBean) PutPackageUpdate(list *input.PackUpdate) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	comboPro := new(schema.ComboProduct)
	//删除套餐id下的所有数据
	count, err = sess.Table(comboPro.TableName()).Where("combo_id = ?", list.Id).Delete(comboPro)
	if err != nil {
		sess.Rollback()
		return
	}
	count, err = sess.Table(comboPro.TableName()).Insert(list.List)
	if err != nil {
		sess.Rollback()
		return
	}
	sess.Commit()
	return
}

//套餐详情
func (m *SitePackageBean) GetPackageInfo(this *input.GetPackage) (data []back.GetList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	getpackage := new(schema.Combo)
	getinfo := new(schema.ComboProduct)
	getlist := new(schema.Product)
	err = sess.Table(getpackage.TableName()).Alias("c").
		Join("LEFT", []string{getinfo.TableName(), "p"}, "c.id=p.combo_id").
		Join("LEFT", []string{getlist.TableName(), "pr"}, "pr.id=p.product_id").
		Where("c.id = ?", this.Id).
		Select("c.id as combo_id,p.product_id as product_id,p.proportion as proportion,pr.product_name as product_name,c.combo_name as combo_name").
		Find(&data)
	return
}

//查看套餐id是否存在
func (m *SitePackageBean) GetComboId(comboId int64) (has bool, err error) {
	site := new(schema.Site)
	sess := global.GetXorm().NewSession().Table(site.TableName())
	defer sess.Close()
	has, err = sess.Where("combo_id = ?", comboId).Where("status = ?", 1).Where("delete_time = ?", 0).Get(site)
	return
}

//修改套餐状态
func (m *SitePackageBean) UpdateStatus(this *input.ComboId) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	combo := new(schema.Combo)
	combos := new(schema.Combo)
	sess.Begin()
	_, err = sess.Table(combo.TableName()).Where("id = ?", this.Id).Get(combo)
	if err != nil {
		sess.Rollback()
	}
	//状态取反
	if combo.Status == 1 {
		combos.Status = 2
	} else if combo.Status == 2 {
		combos.Status = 1
	}
	count, err = sess.Table(combo.TableName()).Where("id = ?", this.Id).Cols("status").Update(combos)
	if err != nil {
		sess.Rollback()
	}
	sess.Commit()
	return
}

//套餐下拉框
func (*SitePackageBean) ComboDropAll() ([]back.ComboDropBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	com := new(schema.Combo)
	var data []back.ComboDropBack
	err := sess.Table(com.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
