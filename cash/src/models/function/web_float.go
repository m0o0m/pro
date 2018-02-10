package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type WebFloatBean struct{}

//站点左右浮动管理
func (*WebFloatBean) GetWebFloatList(this *input.FloatList) (infolist []back.FloatList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteFloat)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("ftype=?", this.Ftype)
	err = sess.Table(spv.TableName()).Find(&infolist)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, err
	}
	return infolist, err
}

//删除站点左右浮动图片
func (*WebFloatBean) DeleteWebFloatList(this *input.DeleteFloatPicture) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sF := new(schema.SiteFloat)
	sF.DeleteTime = global.GetCurrentTime()
	count, err := sess.Where("id=?", this.Id).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Cols("delete_time").Update(sF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取站点左右浮动图片状态
func (*WebFloatBean) GetWebFloatListStatus(this *input.GetFloatListStatus) (*back.FloatAllStatus, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sF := new(schema.SiteFloat)
	//左浮动图片状态查询
	lF := new(back.FloatStatus)
	_, err := sess.Table(sF.TableName()).Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("ftype=?", 1).Where("delete_time=?", 0).
		Get(lF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil, err
	}
	//右浮动图片状态查询
	rF := new(back.FloatStatus)
	_, err = sess.Table(sF.TableName()).Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("ftype=?", 2).Where("delete_time=?", 0).
		Get(rF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil, err
	}
	//站点左右图片状态
	aF := new(back.FloatAllStatus)
	aF.LeftFloat = lF
	aF.RightFloat = rF
	return aF, err
}

//浮动图片禁用开启
func (*WebFloatBean) PutWebFloatStatus(this *input.FloatListStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sF := new(schema.SiteFloat)
	sF.State = int64(this.Status)
	count, err := sess.Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("ftype=?", this.Ftype).Cols("state").
		Update(sF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点左右浮动修改
func (*WebFloatBean) PutWebFloat(this *input.FloatUpdate) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sF := new(schema.SiteFloat)
	sF.IsBlank = this.IsBlank
	sF.IsSlide = this.IsSlide
	sF.Sort = this.Sort
	sF.Url = this.Url
	sF.UrlInter = this.UrlInter
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("ftype=?", this.Ftype)
	sess.Cols("is_blank,is_slide,sort,url,url_inter")
	count, err := sess.Update(sF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点左右浮动图片修改
func (*WebFloatBean) PutWebFloatPicture(this *input.FloatUpdatePicture) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sF := new(schema.SiteFloat)
	sF.Url = this.Url
	if this.CType == 1 {
		sF.ImgA = this.Address
		sess.Cols("img_a,url")
	} else {
		sF.ImgB = this.Address
		sess.Cols("img_b,url")
	}
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("ftype=?", this.Ftype)
	count, err := sess.Update(sF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点左右浮动添加
func (*WebFloatBean) PostWebFloat(this *input.FloatAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sF := new(schema.SiteFloat)
	sF.Ftype = this.Ftype
	sF.Url = this.Url
	sF.IsBlank = this.IsBlank
	sF.IsSlide = this.IsSlide
	sF.SiteId = this.SiteId
	sF.SiteIndexId = this.SiteIndexId
	sF.Sort = this.Sort
	sF.UrlInter = this.UrlInter
	count, err := sess.Insert(sF)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询浮动图
func (m *WebFloatBean) GetListBySite(siteId, siteIndexId string) (floatList []back.FloatList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	floatSchema := new(schema.SiteFloat)
	err = sess.Table(floatSchema.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		In("ftype", []int{1, 2}).
		Find(&floatList)
	return
}
