package function

import (
	"errors"
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
	"time"
)

type WebInfoBean struct{}

//网站基本信息-查询
func (*WebInfoBean) GetSiteInfo(this *input.OrderModuleList) (
	back.GetSiteInfoPcAndWap, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info := new(back.GetSiteInfoPc)
	infoWap := new(back.GetSiteInfoWap)
	site := new(schema.Site)
	sd := new(schema.SiteDomain)
	si := new(schema.SiteInfo)
	sw := new(schema.SiteWapInfo)
	var data back.GetSiteInfoPcAndWap
	//pc
	sess.Select(
		site.TableName() + ".site_name," +
			site.TableName() + ".id," +
			site.TableName() + ".index_id," +
			si.TableName() + ".remark," +
			si.TableName() + ".qq," +
			si.TableName() + ".wechat," +
			si.TableName() + ".phone," +
			si.TableName() + ".email," +
			si.TableName() + ".url_link," +
			sd.TableName() + ".domain")
	has, err := sess.Table(site.TableName()).
		Join("LEFT", si.TableName(), site.TableName()+
			".id="+si.TableName()+".site_id AND "+site.TableName()+".index_id="+
			si.TableName()+".site_index_id").
		Join("LEFT", sd.TableName(), site.TableName()+
			".id="+sd.TableName()+".site_id AND "+site.TableName()+".index_id="+
			sd.TableName()+".site_index_id").
		Where(si.TableName()+".site_id=?", this.SiteId).
		Where(si.TableName()+".site_index_id=?", this.SiteIndexId).Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	//wap
	sess.Select(
		site.TableName() + ".site_name," +
			site.TableName() + ".id," +
			site.TableName() + ".index_id," +
			sw.TableName() + ".app_url," +
			sw.TableName() + ".wap_color," +
			sw.TableName() + ".wap_bottom," +
			sw.TableName() + ".auto_link_name," +
			sw.TableName() + ".auto_link_url," +
			sw.TableName() + ".wap_quick," +
			sw.TableName() + ".is_download," +
			sw.TableName() + ".website_app_url," +
			sd.TableName() + ".domain")

	has, err = sess.Table(site.TableName()).
		Join("LEFT", sw.TableName(), site.TableName()+
			".id="+sw.TableName()+".site_id AND "+site.TableName()+".index_id="+
			sw.TableName()+".site_index_id").
		Join("LEFT", sd.TableName(), site.TableName()+
			".id="+sd.TableName()+".site_id AND "+site.TableName()+".index_id="+
			sd.TableName()+".site_index_id").
		Where(sw.TableName()+".site_id=?", this.SiteId).
		Where(sw.TableName()+".site_index_id=?", this.SiteIndexId).
		Get(infoWap)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	data.Pc = *info
	data.Wap = *infoWap
	return data, has, err
}

//网站基本信息
func (*WebInfoBean) GetSiteInfos(this *input.OrderModuleList) (data back.GetSiteInfoPcAndWap, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	info := new(back.GetSiteInfoPc)
	infoWap := new(back.GetSiteInfoWap)

	site := new(schema.Site)
	has, err = sess.Table(site.TableName()).
		Where("id=? AND index_id = ?", this.SiteId, this.SiteIndexId).
		Select("site_name, id, index_id").
		Get(site)
	if err != nil {
		return
	}
	if !has {
		return
	}
	infoWap.SiteId, info.SiteId = site.Id, site.Id
	infoWap.SiteIndexId, info.SiteIndexId = site.IndexId, site.IndexId
	infoWap.SiteName, info.SiteName = site.SiteName, site.SiteName
	fmt.Println(infoWap, info)

	si := new(schema.SiteInfo)
	has, err = sess.Table(si.TableName()).
		Where("site_id=? AND site_index_id = ?", this.SiteId, this.SiteIndexId).
		Select("remark, qq, wechat, phone, email, url_link").
		Get(si)
	if err != nil {
		return
	}
	if !has {
		return
	}
	info.Remark, info.Qq, info.Wechat, info.Phone, info.Email, info.UrlLink = si.Remark, si.Qq, si.Wechat, si.Phone, si.Email, si.UrlLink

	sw := new(schema.SiteWapInfo)
	has, err = sess.Table(sw.TableName()).
		Where("site_id=? AND site_index_id = ?", this.SiteId, this.SiteIndexId).
		Select("app_url, wap_color, wap_bottom, auto_link_name, auto_link_url, wap_quick, is_download, website_app_url").
		Get(sw)
	if err != nil {
		return
	}
	if !has {
		return
	}
	infoWap.AppUrl = sw.AppUrl
	infoWap.WapColor = sw.WapColor
	infoWap.WapBottom = sw.WapBottom
	infoWap.AutoLinkName = sw.AutoLinkName
	infoWap.AutoLinkUrl = sw.AutoLinkUrl
	infoWap.WapQuick = sw.WapQuick
	infoWap.IsDownload = sw.IsDownload
	infoWap.WebsiteAppUrl = sw.WebsiteAppUrl

	sd := new(schema.SiteDomain)
	has, err = sess.Table(sd.TableName()).
		Where("site_id=? AND site_index_id = ?", this.SiteId, this.SiteIndexId).
		Select("domain").
		Get(sd)
	if err != nil {
		return
	}
	if !has {
		return
	}
	infoWap.WapDomain, info.PcDomain = sd.Domain, sd.Domain
	data.Pc = *info
	data.Wap = *infoWap
	fmt.Println("网站基本信息-查询", data)
	return
}

//网站基本信息-添加或修改pc
func (*WebInfoBean) PostSiteInfo(this *input.PostSiteInfo) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//pc
	info := new(schema.SiteInfo)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Begin()
	conds := sess.Conds()
	has, err := sess.Table(info.TableName()).Get(info)
	if err != nil {
		return 0, err
	}
	info.SiteId = this.SiteId
	info.SiteIndexId = this.SiteIndexId
	info.Email = this.Email
	info.Phone = this.Phone
	info.Qq = this.Qq
	info.Remark = this.Remark
	info.UrlLink = this.UrlLink
	info.Wechat = this.Wechat
	if has {
		//更新
		count, err := sess.Cols("remark,qq,wechat,phone,email,url_link").Where(conds).Update(info)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	} else {
		//新增
		count, err := sess.InsertOne(info)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if count == 0 {
			return count, err
		}
	}
	//wap
	infoWap := new(schema.SiteWapInfo)
	infoWapp := new(schema.SiteWapInfo)
	has, err = sess.Table(infoWap.TableName()).Where(conds).Get(infoWap)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	infoWapp.AppUrl = this.AppUrl
	infoWapp.AutoLinkName = this.AutoLinkName
	infoWapp.IsDownload = this.IsDownload
	infoWapp.WapBottom = this.WapBottom
	infoWapp.WapColor = this.WapColor
	infoWapp.WapQuick = this.WapQuick
	infoWapp.WebsiteAppUrl = this.WebsiteAppUrl
	infoWapp.AutoLinkUrl = this.AutoLinkUrl
	if has {
		id := infoWap.Id
		//更新
		count, err := sess.Cols("app_url,wap_color,wap_bottom,auto_link_name,auto_link_url,"+
			"wap_quick,is_download,website_app_url").Where("id=?", id).Update(infoWapp)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	} else {
		infoWapp.SiteIndexId = this.SiteIndexId
		infoWapp.SiteId = this.SiteId
		//新增
		count, err := sess.InsertOne(infoWapp)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if count == 0 {
			return count, err
		}
	}
	sess.Commit()
	return 1, err
}

//网站基本信息-添加或修改wap
func (*WebInfoBean) PostSiteInfoWap(this *input.PostSiteInfoWap) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoWap := new(schema.SiteWapInfo)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err = sess.Table(infoWap.TableName()).Get(infoWap)
	if err != nil {
		return
	}
	if has {
		//更新
		count, err = sess.Table(infoWap.TableName()).Where(conds).Update(this)
	} else {
		//新增
		count, err = sess.Table(infoWap.TableName()).InsertOne(this)
	}
	return
}

//获取电子管理排序列表
func (*WebInfoBean) GetDzOrderList(this *input.OrderModuleList) ([]back.OrderModuleList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.OrderModuleList
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data1 := new(back.OrderModuleList)
	//------------------------取本站点的商品连成字符串------------------------

	product := new(schema.Product)
	productDel := new(schema.SiteProductDel)
	var moduleStr string
	var ids []int64
	VTypes := make([]string, 0)
	VTypesMap := make(map[string]string)
	sess.Where("type_id=?", 2)
	sql1 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), productDel.TableName())
	sess.Table(product.TableName()).
		Select("product_id").Where(conds).
		Join("LEFT", productDel.TableName(), sql1).
		Find(&ids)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	sess.Where("type_id=?", 2)
	if len(ids) > 0 {
		sess.NotIn("id", ids)
	}
	sess.Table(product.TableName()).
		Select("v_type").Where("delete_time=?", 0).
		Find(&VTypes)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if order.DzModule != "" {
		sP := strings.Split(order.DzModule, ",")
		for i, v := range sP {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, *data1)
			VTypesMap[v] = v
		}
	} else {
		for i, v := range VTypes {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, *data1)
			VTypesMap[v] = v
		}
	}
	moduleStr = strings.Join(VTypes, ",")
	//------------------------判断操作数据库-----------------------

	order.SiteId = this.SiteId
	order.SiteIndexId = this.SiteIndexId
	arr := strings.Split(order.DzModule, ",")
	checkStr := 0
	if has && len(VTypes) == len(arr) {
		for _, v := range arr {
			_, ok := VTypesMap[v]
			if !ok {
				checkStr = 1
				break
			}
		}
	}
	//如果有且checkStr=0,说明商品完全相同，不操作
	if (has && len(VTypes) != len(arr)) || checkStr == 1 {
		//如果有且含有商品数目不等，或者含有商品数目相同但内容不完全相同 更新
		order.DzModule = moduleStr
		_, err = sess.Table(order.TableName()).Where(conds).Cols("dz_module").Update(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	} else if !has {
		//如果没有 新增
		order.DzModule = moduleStr
		_, err = sess.Table(order.TableName()).InsertOne(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	}
	return data, err
}

//修改电子管理排序
func (*WebInfoBean) EditDzOrder(this *input.EditModuleOrder) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if !has {
		return 0, err
	}
	if order.DzModule == "" {
		return 0, err
	}
	data := make([]back.OrderModuleList, 0)
	data1 := new(back.OrderModuleList)
	var moduleStr string
	arr := strings.Split(order.DzModule, ",")
	checkAddProduct := 0
	for i, v := range arr {
		if v == this.VType {
			//商品等于传入商品
			checkAddProduct = 1
			continue
		} else if int64(i+1) == this.IndexNum {
			//序号等于传入序号
			if checkAddProduct == 1 {
				data1.IndexNum = int64(len(data) + 1)
				data1.Product = v
				data = append(data, *data1)
				moduleStr = moduleStr + v + ","
				checkAddProduct = 2
			}
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = this.VType
			data = append(data, *data1)
			moduleStr = moduleStr + this.VType + ","
		}
		if checkAddProduct == 2 {
			checkAddProduct = 0
		} else {
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = v
			data = append(data, *data1)
			moduleStr = moduleStr + v + ","
		}
	}
	if len(arr)-len(data) == 1 {
		//防止传入序号大于总商品数
		data1.IndexNum = int64(len(data) + 1)
		data1.Product = this.VType
		data = append(data, *data1)
		moduleStr = moduleStr + this.VType + ","
	}
	fmt.Print(moduleStr[:len(moduleStr)-1])
	order.DzModule = moduleStr[:len(moduleStr)-1]
	count, err := sess.Table(order.TableName()).Where(conds).Cols("dz_module").Update(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取体育管理排序列表
func (*WebInfoBean) GetSpOrderList(this *input.OrderModuleList) (
	[]back.OrderModuleList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.OrderModuleList
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data1 := new(back.OrderModuleList)
	//------------------------取本站点的商品连成字符串------------------------

	product := new(schema.Product)
	productDel := new(schema.SiteProductDel)
	var moduleStr string
	var ids []int64
	VTypes := make([]string, 0)
	VTypesMap := make(map[string]string)
	sess.Where("type_id=?", 5)
	sql1 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), productDel.TableName())
	sess.Table(product.TableName()).
		Select("product_id").Where(conds).
		Join("LEFT", productDel.TableName(), sql1).
		Find(&ids)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	sess.Where("type_id=?", 5)
	if len(ids) > 0 {
		sess.NotIn("id", ids)
	}
	sess.Table(product.TableName()).
		Select("v_type").Where("delete_time=?", 0).
		Find(&VTypes)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if order.SpModule != "" {
		sP := strings.Split(order.SpModule, ",")
		for i, v := range sP {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, *data1)
			VTypesMap[v] = v
		}
	} else {
		for i, v := range VTypes {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, *data1)
			VTypesMap[v] = v
		}
	}
	moduleStr = strings.Join(VTypes, ",")
	//------------------------判断操作数据库-----------------------

	order.SiteId = this.SiteId
	order.SiteIndexId = this.SiteIndexId
	arr := strings.Split(order.SpModule, ",")
	checkStr := 0
	if has && len(VTypes) == len(arr) {
		for _, v := range arr {
			_, ok := VTypesMap[v]
			if !ok {
				checkStr = 1
				break
			}
		}
	}
	//如果有且checkStr=0,说明商品完全相同，不操作
	if (has && len(VTypes) != len(arr)) || checkStr == 1 {
		//如果有且含有商品数目不等，或者含有商品数目相同但内容不完全相同 更新
		order.SpModule = moduleStr
		_, err = sess.Table(order.TableName()).Where(conds).Cols("sp_module").Update(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	} else if !has {
		//如果没有 新增
		order.SpModule = moduleStr
		_, err = sess.Table(order.TableName()).InsertOne(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	}
	return data, err
}

//修改体育管理排序
func (*WebInfoBean) EditSpOrder(this *input.EditModuleOrder) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if !has {
		return 0, err
	}
	if order.SpModule == "" {
		return 0, err
	}
	data := make([]back.OrderModuleList, 0)
	data1 := new(back.OrderModuleList)
	var moduleStr string
	arr := strings.Split(order.SpModule, ",")
	checkAddProduct := 0
	for i, v := range arr {
		if v == this.VType {
			//商品等于传入商品
			checkAddProduct = 1
			continue
		} else if int64(i+1) == this.IndexNum {
			//序号等于传入序号
			if checkAddProduct == 1 {
				data1.IndexNum = int64(len(data) + 1)
				data1.Product = v
				data = append(data, *data1)
				moduleStr = moduleStr + v + ","
				checkAddProduct = 2
			}
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = this.VType
			data = append(data, *data1)
			moduleStr = moduleStr + this.VType + ","
		}
		if checkAddProduct == 2 {
			checkAddProduct = 0
		} else {
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = v
			data = append(data, *data1)
			moduleStr = moduleStr + v + ","
		}
	}
	if len(arr)-len(data) == 1 {
		//防止传入序号大于总商品数
		data1.IndexNum = int64(len(data) + 1)
		data1.Product = this.VType
		data = append(data, *data1)
		moduleStr = moduleStr + this.VType + ","
	}
	fmt.Print(moduleStr[:len(moduleStr)-1])
	order.SpModule = moduleStr[:len(moduleStr)-1]
	count, err := sess.Table(order.TableName()).Where(conds).Cols("sp_module").Update(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取彩票管理排序列表
func (*WebInfoBean) GetFcOrderList(this *input.OrderModuleList) (
	[]back.OrderModuleList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	order := new(schema.SiteOrderModule)
	var has bool
	var data []back.OrderModuleList
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data1 := new(back.OrderModuleList)
	//------------------------取本站点的商品连成字符串------------------------

	product := new(schema.Product)
	productDel := new(schema.SiteProductDel)
	var moduleStr string
	var ids []int64
	VTypes := make([]string, 0)
	VTypesMap := make(map[string]string)
	sess.Where("type_id=?", 4)
	sql1 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), productDel.TableName())
	sess.Table(product.TableName()).
		Select("product_id").Where(conds).
		Join("LEFT", productDel.TableName(), sql1).
		Find(&ids)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	sess.Where("type_id=?", 4)
	if len(ids) > 0 {
		sess.NotIn("id", ids)
	}
	sess.Table(product.TableName()).
		Select("v_type").Where("delete_time=?", 0).
		Find(&VTypes)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if order.FcModule != "" {
		sP := strings.Split(order.FcModule, ",")
		for i, v := range sP {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, *data1)
			VTypesMap[v] = v
		}
	} else {
		for i, v := range VTypes {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, *data1)
			VTypesMap[v] = v
		}
	}
	moduleStr = strings.Join(VTypes, ",")
	//------------------------判断操作数据库-----------------------

	order.SiteId = this.SiteId
	order.SiteIndexId = this.SiteIndexId
	arr := strings.Split(order.FcModule, ",")
	checkStr := 0
	if has && len(VTypes) == len(arr) {
		for _, v := range arr {
			_, ok := VTypesMap[v]
			if !ok {
				checkStr = 1
				break
			}
		}
	}
	//如果有且checkStr=0,说明商品完全相同，不操作
	if (has && len(VTypes) != len(arr)) || checkStr == 1 {
		//如果有且含有商品数目不等，或者含有商品数目相同但内容不完全相同 更新
		order.FcModule = moduleStr
		_, err = sess.Table(order.TableName()).Where(conds).Cols("fc_module").Update(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	} else if !has {
		//如果没有 新增
		order.FcModule = moduleStr
		_, err = sess.Table(order.TableName()).InsertOne(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	}
	return data, err
}

//彩票大厅重置排序
func (*WebInfoBean) GetFcOrderFcReset(this *input.OrderModuleList) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	//------------------------取本站点的商品连成字符串------------------------

	product := new(schema.Product)
	productDel := new(schema.SiteProductDel)
	var moduleStr string
	var ids []int64
	VTypes := make([]string, 0)
	sess.Where("type_id=?", 4)
	sql1 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), productDel.TableName())
	err = sess.Table(product.TableName()).
		Select("product_id").Where(conds).
		Join("LEFT", productDel.TableName(), sql1).
		Find(&ids)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	sess.Where("type_id=?", 4)
	if len(ids) > 0 {
		sess.NotIn("id", ids)
	}
	err = sess.Table(product.TableName()).
		Select("v_type").Where("delete_time=?", 0).
		Find(&VTypes)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	moduleStr = strings.Join(VTypes, ",")
	//------------------------重置-----------------------

	order.SiteId = this.SiteId
	order.SiteIndexId = this.SiteIndexId
	order.FcModule = moduleStr
	//如果有重置
	if has {
		//如果有且含有商品数目不等，或者含有商品数目相同但内容不完全相同 更新
		count, err := sess.Table(order.TableName()).Where(conds).Cols("fc_module").Update(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
		return count, err
	}
	//如果没有 新增
	order.FcModule = moduleStr
	count, err := sess.Table(order.TableName()).InsertOne(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改彩票管理排序
func (*WebInfoBean) EditFcOrder(this *input.EditModuleOrder) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if !has {
		return 0, err
	}
	if order.FcModule == "" {
		return 0, err
	}
	data := make([]back.OrderModuleList, 0)
	data1 := new(back.OrderModuleList)
	var moduleStr string
	arr := strings.Split(order.FcModule, ",")
	checkAddProduct := 0
	for i, v := range arr {
		if v == this.VType {
			//商品等于传入商品
			checkAddProduct = 1
			continue
		} else if int64(i+1) == this.IndexNum {
			//序号等于传入序号
			if checkAddProduct == 1 {
				data1.IndexNum = int64(len(data) + 1)
				data1.Product = v
				data = append(data, *data1)
				moduleStr = moduleStr + v + ","
				checkAddProduct = 2
			}
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = this.VType
			data = append(data, *data1)
			moduleStr = moduleStr + this.VType + ","
		}
		if checkAddProduct == 2 {
			checkAddProduct = 0
		} else {
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = v
			data = append(data, *data1)
			moduleStr = moduleStr + v + ","
		}
	}
	if len(arr)-len(data) == 1 {
		//防止传入序号大于总商品数
		data1.IndexNum = int64(len(data) + 1)
		data1.Product = this.VType
		data = append(data, *data1)
		moduleStr = moduleStr + this.VType + ","
	}
	fmt.Print(moduleStr[:len(moduleStr)-1])
	order.FcModule = moduleStr[:len(moduleStr)-1]
	count, err := sess.Table(order.TableName()).Where(conds).Cols("fc_module").Update(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取视讯管理排序列表
func (*WebInfoBean) GetVideoOrderList(this *input.OrderModuleList) (
	[]back.OrderModuleList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.OrderModuleList
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data1 := new(back.OrderModuleList)
	//------------------------取本站点的商品连成字符串------------------------

	product := new(schema.Product)
	productDel := new(schema.SiteProductDel)
	var moduleStr string
	var ids []int64
	VTypes := make([]string, 0)
	VTypesMap := make(map[string]string)
	sess.Where("type_id=?", 1)
	sql1 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), productDel.TableName())
	err = sess.Table(product.TableName()).
		Select("product_id").Where(conds).
		Join("LEFT", productDel.TableName(), sql1).
		Find(&ids)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	sess.Where("type_id=?", 1)
	if len(ids) > 0 {
		sess.NotIn("id", ids)
	}
	err = sess.Table(product.TableName()).
		Select("v_type").Where("delete_time=?", 0).
		Find(&VTypes)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if order.VideoModule != "" {
		art := strings.Split(order.VideoModule, ",")
		for i, v := range art {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, *data1)
			VTypesMap[v] = v
		}
	} else {
		for i, v := range VTypes {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, *data1)
			VTypesMap[v] = v
		}
	}
	moduleStr = strings.Join(VTypes, ",")
	//------------------------判断操作数据库-----------------------
	order.SiteId = this.SiteId
	order.SiteIndexId = this.SiteIndexId
	arr := strings.Split(order.VideoModule, ",")
	checkStr := 0
	if has && len(VTypes) == len(arr) {
		for _, v := range arr {
			_, ok := VTypesMap[v]
			if !ok {
				checkStr = 1
				break
			}
		}
	}
	//如果有且checkStr=0,说明商品完全相同，不操作
	if (has && len(VTypes) != len(arr)) || checkStr == 1 {
		//如果有且含有商品数目不等，或者含有商品数目相同但内容不完全相同 更新
		order.VideoModule = moduleStr
		_, err = sess.Table(order.TableName()).Where(conds).Cols("video_module").Update(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	} else if !has {
		//如果没有 新增
		order.VideoModule = moduleStr
		_, err = sess.Table(order.TableName()).InsertOne(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	}
	return data, err
}

//修改视讯管理排序
func (*WebInfoBean) EditVideoOrder(this *input.EditModuleOrder) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if !has {
		return 0, err
	}
	if order.VideoModule == "" {
		return 0, err
	}
	styleUse := new(schema.InfoVideoUse)
	data := make([]back.OrderModuleList, 0)
	data1 := new(back.OrderModuleList)
	var moduleStr string
	arr := strings.Split(order.VideoModule, ",")
	checkAddProduct := 0
	for i, v := range arr {
		if v == this.VType {
			//商品等于传入商品
			checkAddProduct = 1
			continue
		} else if int64(i+1) == this.IndexNum {
			//序号等于传入序号
			if checkAddProduct == 1 {
				data1.IndexNum = int64(len(data) + 1)
				data1.Product = v
				data = append(data, *data1)
				moduleStr = moduleStr + v + ","
				checkAddProduct = 2
			}
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = this.VType
			data = append(data, *data1)
			moduleStr = moduleStr + this.VType + ","
		}
		if checkAddProduct == 2 {
			checkAddProduct = 0
		} else {
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = v
			data = append(data, *data1)
			moduleStr = moduleStr + v + ","
		}
	}
	if len(arr)-len(data) == 1 {
		//防止传入序号大于总商品数
		data1.IndexNum = int64(len(data) + 1)
		data1.Product = this.VType
		data = append(data, *data1)
		moduleStr = moduleStr + this.VType + ","
	}
	fmt.Print(moduleStr[:len(moduleStr)-1])
	order.VideoModule = moduleStr[:len(moduleStr)-1]
	styleUse.Video = order.VideoModule
	var has1 bool
	has1, err = sess.Table(styleUse.TableName()).Where(conds).Get(styleUse)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	sess.Begin()
	count, err := sess.Table(order.TableName()).Where(conds).Cols("video_module").Update(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if has1 {
		count, err = sess.Table(styleUse.TableName()).Where(conds).Cols("video").Update(styleUse) //更新视讯模板使用表
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	}
	sess.Commit()
	return count, err
}

//获取视讯管理类型下拉框
func (*WebInfoBean) GetVideoOrderTypeList() ([]back.TypeOrderModuleList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.TypeOrderModuleList
	vtype := new(schema.SiteInfoVideoStyle)
	sess.Where("pid=?", 0)
	sess.Where("status=?", 1)
	err := sess.Table(vtype.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取视讯管理风格下拉框
func (*WebInfoBean) GetVideoOrderStyleList(this *input.StyleOrderModuleList) (
	[]back.StyleOrderModuleList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.StyleOrderModuleList
	vstyle := new(schema.SiteInfoVideoStyle)
	sess.Where("pid=?", this.Pid)
	sess.Where("status=?", 1)
	err := sess.Table(vstyle.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//视讯管理-风格配置使用
func (*WebInfoBean) PostVideoOrderStyleUse(this *input.PostVideoOrderStyleUse) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	styleUse := new(schema.InfoVideoUse)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(styleUse.TableName()).Get(styleUse)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	styleUse.Style = this.Style
	styleUse.Type = this.Type
	styleUse.State = 1
	styleUse.DoTime = time.Now().Unix()
	if has {
		//更新
		count, err := sess.Table(styleUse.TableName()).Where(conds).Cols("style", "type", "do_time").Update(styleUse)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
		if count == 0 {
			return count, err
		}
	} else {
		//新增
		order := new(schema.SiteOrderModule)
		var has1 bool
		has1, err = sess.Table(order.TableName()).Select("video_module").Where(conds).Get(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return 0, err
		}
		if has1 {
			styleUse.Video = order.VideoModule
		}
		styleUse.SiteIndexId = this.SiteIndexId
		styleUse.SiteId = this.SiteId
		count, err := sess.Table(styleUse.TableName()).InsertOne(styleUse)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
		if count == 0 {
			return count, err
		}
	}
	return 1, err
}

//视讯管理-风格还原默认
func (*WebInfoBean) PutVideoOrderStyleUseUpdate(this *input.PutVideoOrderStyleUseUpdate) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	styleUse := new(schema.InfoVideoUse)
	styleUse.Style = 0
	styleUse.Type = 0
	styleUse.DoTime = time.Now().Unix()
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	count, err := sess.Table(styleUse.TableName()).Cols("style", "type", "do_time").Update(styleUse)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取模块管理排序列表
func (*WebInfoBean) GetModuleOrderList(this *input.OrderModuleList) (
	[]back.OrderModuleList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.OrderModuleList
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	var data1 back.OrderModuleList
	module := "video,dz,sp,fc"
	order.SiteId = this.SiteId
	order.SiteIndexId = this.SiteIndexId
	if !has {
		order.Module = module
		_, err := sess.Table(order.TableName()).InsertOne(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	}
	if order.Module == "0" || order.Module == "" {
		order.Module = module
		_, err := sess.Table(order.TableName()).Where(conds).Cols("module").Update(order)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	}
	if order.Module != "" {
		mo1 := strings.Split(order.Module, ",")
		for i, v := range mo1 {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, data1)
		}
	} else {
		mo2 := strings.Split(module, ",")
		for i, v := range mo2 {
			data1.IndexNum = int64(i + 1)
			data1.Product = v
			data = append(data, data1)
		}
	}
	return data, err
}

//修改模块管理排序
func (*WebInfoBean) EditModuleOrder(this *input.EditModuleOrder) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	order := new(schema.SiteOrderModule)
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(order.TableName()).Get(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if !has {
		return 0, err
	}
	if order.Module == "" {
		return 0, err
	}
	data := make([]back.OrderModuleList, 0)
	data1 := new(back.OrderModuleList)
	var moduleStr string
	arr := strings.Split(order.Module, ",")
	checkAddProduct := 0
	for i, v := range arr {
		if v == this.VType {
			//商品等于传入商品
			checkAddProduct = 1
			continue
		} else if int64(i+1) == this.IndexNum {
			//序号等于传入序号
			if checkAddProduct == 1 {
				data1.IndexNum = int64(len(data) + 1)
				data1.Product = v
				data = append(data, *data1)
				moduleStr = moduleStr + v + ","
				checkAddProduct = 2
			}
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = this.VType
			data = append(data, *data1)
			moduleStr = moduleStr + this.VType + ","
		}
		if checkAddProduct == 2 {
			checkAddProduct = 0
		} else {
			data1.IndexNum = int64(len(data) + 1)
			data1.Product = v
			data = append(data, *data1)
			moduleStr = moduleStr + v + ","
		}
	}
	if len(arr)-len(data) == 1 {
		//防止传入序号大于总商品数
		data1.IndexNum = int64(len(data) + 1)
		data1.Product = this.VType
		data = append(data, *data1)
		moduleStr = moduleStr + this.VType + ","
	}
	fmt.Print(moduleStr[:len(moduleStr)-1])
	order.Module = moduleStr[:len(moduleStr)-1]
	count, err := sess.Table(order.TableName()).Where(conds).Cols("module").Update(order)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//得到站点模块信息
func (*WebInfoBean) GetOrderNumberBySite(siteId, siteIndexId string) (*schema.SiteOrderModule, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteOrderModuleSchema := new(schema.SiteOrderModule)
	b, err := sess.Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Get(siteOrderModuleSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return siteOrderModuleSchema, err
	}
	if !(b) {
		return siteOrderModuleSchema, errors.New("get 0 row")
	}
	return siteOrderModuleSchema, nil
}

//电子内页主题配置修改/新增
func (*WebInfoBean) PutElectConfig(this *input.PromotionSet) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	styleUse := new(schema.InfoActivityPromotionSet)
	colorset := new(input.PromotionColorSet)
	colorset.Bcolor = this.TitleBcolor + "," + this.TitleColor + "," + this.ButtonBcolor + "," + this.ButtonColor + "," + this.BborderColor + "," + this.PopBcolor
	var has bool
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	has, err := sess.Table(styleUse.TableName()).Get(styleUse)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if has {
		//更新
		count, err := sess.Table(styleUse.TableName()).
			Where(conds).Cols("bcolor").
			Update(colorset)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else {
		//新增
		colorset.SiteIndexId = this.SiteIndexId
		colorset.SiteId = this.SiteId
		count, err := sess.Table(styleUse.TableName()).InsertOne(colorset)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	}
	return 1, err
}

//获取电子内页主题配置
func (*WebInfoBean) GetElectConfig(this *input.Site) (*back.BcolorList, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	electconfig := new(schema.InfoActivityPromotionSet)
	data := new(back.BcolorList)
	has, err := sess.Table(electconfig.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//电子内页主题配置初始化
func (*WebInfoBean) InitializationDianZi(this *input.DianZiInitialization) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	el := new(schema.InfoActivityPromotionSet)
	//查询该站点是否存在电子主题配置
	has, err := sess.Get(el)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if !has {
		el.SiteId = this.SiteId
		el.SiteIndexId = this.SiteIndexId
		count, err := sess.InsertOne(el)
		return count, err
	}
	id := el.Id
	el.Bcolor = ""
	count, err := sess.Cols("bcolor").Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("id=?", id).
		Update(el)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
