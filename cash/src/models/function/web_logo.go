package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type WebLogoBean struct{}

//站点logo图片管理列表
func (*WebLogoBean) GetWebLogoList(this *input.LogoInfoList) (infolist []back.InfoLogoList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.InfoLogo)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("type=?", 5)
	err = sess.Table(spv.TableName()).Find(&infolist)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, err
	}
	return infolist, err
}

//站点log修改(基本信息)
func (*WebLogoBean) PutWebLogo(this *input.UpdateLogoInfo) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.InfoLogo)
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	data, err = sess.Table(spv.TableName()).Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//站点log修改(logo路径)
func (*WebLogoBean) PutWebLogoWay(this *input.UpdateLogoInfoWay) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.InfoLogo)
	spv.LogoUrl = this.LogoUrl
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Cols("logo_url")
	data, err = sess.Update(spv)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

func (*WebLogoBean) PostWebLogo(this *input.PostLogoInfo) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.InfoLogo)
	data, err = sess.Table(spv.TableName()).Insert(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//站点logo查询
func (*WebLogoBean) GetWebLogo(Id int64, SiteId string, SiteIndexId string, Type int8) (info []back.GetLogoInfo, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.InfoLogo)
	if Id != 0 {
		sess.Where("id=?", Id)
	}
	if Type != 0 {
		sess.Where("type=?", Type)
	}
	sess.Where("site_id=?", SiteId)
	sess.Where("site_index_id=?", SiteIndexId)
	err = sess.Table(spv.TableName()).Find(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, err
	}
	return info, err
}
