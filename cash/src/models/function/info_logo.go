package function

import (
	"errors"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type InfoLogoBean struct{}

//首页logo
func (*InfoLogoBean) HomePageOneLogo(site_id, site_index_id string) (*back.HomePageAllInfo, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	data := new(back.HomePageAllInfo)
	//首页logo
	il := new(schema.InfoLogo)
	var logo back.HomePageLogoBack
	_, err := sess.Table(il.TableName()).
		Where("site_id=?", site_id).
		Where("site_index_id=?", site_index_id).
		Where("state=?", 1).Get(&logo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data.HomePageLogoBack = logo
	//首页轮播图
	s_p := new(schema.SiteFlash)
	var lun []back.HomePagePictureBack
	sess.Select("id,img_url,img_link")
	err = sess.Table(s_p.TableName()).
		Where("state=?", 1).
		Where("site_id=?", site_id).
		Where("site_index_id=?", site_index_id).
		Where("ftype=?", 2).Find(&lun)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data.HomePagePictureBack = lun
	//首页公告
	s_n := new(schema.SiteNotice)
	var notice []back.HomePageNoticeBack
	err = sess.Table(s_n.TableName()).
		Where("notice_assign like ? OR notice_assign=?", "%"+site_id+"%", 1).
		Where("delete_time=?", 0).
		Where("notice_cate=?", 1).Find(&notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data.HomePageNoticeBack = notice
	return data, err
}

//得到站点模块信息
func (*InfoLogoBean) GetOrderNumberModuleBySite(this *input.Home) (*schema.SiteOrderModule, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteOrderModuleSchema := new(schema.SiteOrderModule)
	b, err := sess.Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Get(siteOrderModuleSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return siteOrderModuleSchema, err
	}
	if !(b) {
		return siteOrderModuleSchema, errors.New("get 0 row")
	}
	return siteOrderModuleSchema, err
}

//查询与vType对应的Product,Name
func (*InfoLogoBean) ProductList(vTypes ...[]string) ([]back.ProductlistRep, error) {
	var vType []string
	if len(vTypes) > 0 {
		vType = vTypes[0]
	}
	var data []back.ProductlistRep
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	if len(vType) > 0 {
		sess.In("v_type", vTypes)
	}
	err := sess.Table(product.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
