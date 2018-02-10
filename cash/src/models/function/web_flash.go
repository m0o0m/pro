package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type WebFlashBean struct{}

//站点轮播图片管理 查询
func (*WebFlashBean) GetWebFlashList(this *input.FlashList) (back.FlashAllList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteFlash)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	var data back.FlashAllList
	sess.Where("ftype=?", 1)
	var pcF []back.FlashList
	err := sess.Table(spv.TableName()).Find(&pcF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	var wapF []back.FlashList
	sess.Where("ftype=?", 2)
	err = sess.Table(spv.TableName()).Where(conds).Find(&wapF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data.PcFlashList = pcF
	data.WapFlashList = wapF
	return data, err
}

//站点轮播图片管理 修改
func (*WebFlashBean) PutWebFlash(this *input.FlashUpdate) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sF := new(schema.SiteFlash)
	sF.ImgLink = this.ImgLink
	sF.ImgTitle = this.ImgTitle
	sF.ImgUrl = this.ImgUrl
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("ftype=?", this.Ftype)
	sess.Cols("img_title,img_url,img_link")
	data, err = sess.Update(sF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//站点轮播图片管理 添加
func (*WebFlashBean) PostWebFlash(this *input.FlashAdd) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteFlash)
	data, err = sess.Table(spv.TableName()).Insert(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//站点轮播图片管理 添加
func (*WebFlashBean) PutWebFlashStatus(this *input.FlashStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteFlash)
	if this.Status == 1 {
		spv.State = 2
	} else {
		spv.State = 1
	}
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("ftype=?", this.Ftype)
	sess.Cols("state")
	data, err := sess.Update(spv)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
