package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type MemberRetreatWaterSetBean struct {
}

//优惠查询列表
func (*MemberRetreatWaterSetBean) SearchMemberRetreatWaterSet(this *input.ListRetreatWater, listparam *global.ListParams, sTime, eTime int64) (
	[]back.ListRetreatWater, int64, error) {
	water := new(schema.MemberRetreatWater)
	sess := global.GetXorm().Table(water.TableName())
	defer sess.Close()
	var data []back.ListRetreatWater
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if sTime != 0 && eTime != 0 {
		sess.Where("start_time >= ?", sTime).Where("start_time <= ?", eTime)
	}
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	err := sess.Table(water.TableName()).
		Select("id,admin_user,start_time," +
			"end_time,create_time,event,no_people_num,people_num,money,bet").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(water.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取单个列表详情
func (*MemberRetreatWaterSetBean) DetailMemberRetreatWaterSet(this *input.GetOneRetreatWaterDetails) (
	back.RetreatWaterSetList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var has bool
	var data back.RetreatWaterSetList
	//查询退水设定表(主表)==================================================
	//---------------------查询条件-------------------
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("id=?", this.Id)
	sess.Where("delete_time=?", 0)
	//---------------------查询执行-------------------
	water := new(schema.MemberRetreatWaterSet)
	has, err := sess.Table(water.TableName()).Get(water)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	if !has {
		return data, 0, err
	}
	//-------------------data数据组装-----------------
	count := int64(1)                  //主表有多少条记录，GET最多只有一条
	data.DiscountUp = water.DiscountUp //第一层的参数加进返回值里来
	data.ValidMoney = water.ValidMoney //第一层的参数加进返回值里来
	//查询商品表(副表)======================================================
	//---------------------查询条件-------------------
	sess.Where("status=?", 1)
	sess.Where("delete_time=?", 0)
	//---------------------查询执行-------------------
	products := make([]*schema.Product, 0)
	err = sess.Table(new(schema.Product).TableName()).Cols("id,product_name").Find(&products)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	//------------根据中间表外键转成map+外键组-----------
	ids := make([]int64, 0)
	productMap := make(map[int64]string)
	for _, v := range products {
		ids = append(ids, v.Id)          //外键组，用于中间表条件
		productMap[v.Id] = v.ProductName //map,用于中间表转换数据
	}
	//查询退水商品表(中间表)================================================
	//---------------------查询条件-------------------
	sess.Where("set_id=?", this.Id)
	sess.In("product_id", ids)
	//---------------------查询执行-------------------
	waterPros := make([]*schema.MemberRetreatWaterProduct, 0)
	err = sess.Table(new(schema.MemberRetreatWaterProduct).TableName()).Cols("set_id,product_id,rate").Find(&waterPros)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	//-------------------data数据组装-----------------
	if len(waterPros) > 0 {
		for _, v := range waterPros {
			data.Params = append(data.Params, back.RetreatWaterProductList{
				v.ProductId,
				productMap[v.ProductId],
				v.Rate})
		}
	}
	return data, count, err
}

//返点优惠设定列表
func (*MemberRetreatWaterSetBean) ListMemberRetreatWaterSet(this *input.RetreatWaterSetList) (
	[]back.RetreatWaterSetList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.RetreatWaterSetList
	//查询退水设定表(主表)==================================================
	//---------------------查询条件-------------------
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("delete_time=?", 0)
	//---------------------查询执行-------------------
	waters := make([]*schema.MemberRetreatWaterSet, 0)
	err := sess.Table(new(schema.MemberRetreatWaterSet).TableName()).Find(&waters)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count := int64(len(waters))
	if count == 0 {
		return data, 0, err
	}
	//-----------根据中间表外键转成map+外键组-----------
	ids := make([]int64, 0)
	waterMap := make(map[int64]*back.RetreatWaterSetList)
	for _, v := range waters {
		ids = append(ids, v.Id)
		waterMap[v.Id] = &back.RetreatWaterSetList{
			v.ValidMoney,
			v.DiscountUp,
			make([]back.RetreatWaterProductList, 0)}
	}
	//查询商品表(副表)======================================================
	//---------------------查询条件-------------------
	sess.Where("status=?", 1)
	sess.Where("delete_time=?", 0)
	//---------------------查询执行-------------------
	products := make([]*schema.Product, 0)
	err = sess.Table(new(schema.Product).TableName()).Cols("id,product_name").Find(&products)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	//------------根据中间表外键转成map+外键组-----------
	ids2 := make([]int64, 0)
	productMap := make(map[int64]string)
	for _, v := range products {
		ids2 = append(ids2, v.Id)
		productMap[v.Id] = v.ProductName
	}
	//查询退水商品表(中间表)==================================================
	//---------------------查询条件-------------------
	sess.In("set_id", ids)      //外键组1
	sess.In("product_id", ids2) //外键组2
	waterPros := make([]*schema.MemberRetreatWaterProduct, 0)
	//---------------------查询执行-------------------
	err = sess.Table(new(schema.MemberRetreatWaterProduct).TableName()).
		Cols("set_id,product_id,rate").Find(&waterPros)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	//三个表的三组数据统一组装=================================================
	for _, v := range waterPros {
		waterMap[v.SetId].Params = append(waterMap[v.SetId].Params, back.RetreatWaterProductList{
			v.ProductId,
			productMap[v.ProductId],
			v.Rate})
	}
	for _, v := range waterMap {
		data = append(data, *v)
	}
	return data, count, err
}

//新增返点优惠设定
func (*MemberRetreatWaterSetBean) AddMemberRetreatWaterSet(this *input.AddRetreatWaterSet) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	water := new(schema.MemberRetreatWaterSet)
	var waterPro schema.MemberRetreatWaterProduct
	var waterPros []schema.MemberRetreatWaterProduct
	water.SiteIndexId = this.SiteIndexId
	water.SiteId = this.SiteId
	water.ValidMoney = this.ValidMoney
	water.DiscountUp = this.DiscountUp
	sess.Begin()
	//新增返点设定数据
	count, err := sess.Table(water.TableName()).InsertOne(water)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	for k := range this.Params {
		waterPro.SetId = water.Id
		waterPro.ProductId = this.Params[k].ProductId
		waterPro.Rate = this.Params[k].Rate
		waterPros = append(waterPros, waterPro)
	}
	//新增返点设定之商品百分比
	count, err = sess.Table(waterPro.TableName()).Insert(waterPros)
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

//修改返点优惠设定
func (*MemberRetreatWaterSetBean) EditMemberRetreatWaterSet(this *input.EditRetreatWaterSet) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	water := new(schema.MemberRetreatWaterSet)
	var waterPro schema.MemberRetreatWaterProduct
	var waterPros []schema.MemberRetreatWaterProduct
	water.Id = this.Id
	water.SiteIndexId = this.SiteIndexId
	water.SiteId = this.SiteId
	water.ValidMoney = this.ValidMoney
	water.DiscountUp = this.DiscountUp
	sess.Begin()
	count, err := sess.Table(water.TableName()).Where("id = ?", water.Id).Cols("site_id,site_index_id,valid_money,discount_up").Update(water)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//删除设置id下的所有数据
	count, err = sess.Table(waterPro.TableName()).Where("set_id = ?", water.Id).Delete(waterPro)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	for k := range this.Params {
		waterPro.SetId = water.Id
		waterPro.ProductId = this.Params[k].ProductId
		waterPro.Rate = this.Params[k].Rate
		waterPros = append(waterPros, waterPro)
	}
	//新增返点设定之商品百分比
	count, err = sess.Table(waterPro.TableName()).Insert(waterPros)
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

//删除返点优惠设定
func (*MemberRetreatWaterSetBean) DelMemberRetreatWaterSet(l *input.DelRetreatWaterSet) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.MemberRetreatWaterSet)
	member.Id = int64(l.Id)
	member.DeleteTime = time.Now().Unix()
	count, err := sess.Table(member.TableName()).
		Where("id = ?", member.Id).
		Cols("delete_time").Update(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
