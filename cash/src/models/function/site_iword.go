package function

import (
	"global"
	"time"

	"html/template"
	"models/back"
	"models/input"
	"models/schema"
)

type SiteIwordBean struct{}

//查询文案列表
func (*SiteIwordBean) IWordList(this *input.SiteIWodList) ([]back.IndexWord, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	iword := new(schema.Iword)
	var data []back.IndexWord
	sess.Where("site_id = ?", this.SiteId)
	sess.Where("site_index_id = ?", this.SiteIndexId)
	if this.State != 0 {
		sess.Where("state = ?", this.State)
	}
	if this.Itype != 0 {
		sess.Where("itype = ?", this.Itype)
	}
	err := sess.Table(iword.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询 存款文案内容 列表
func (*SiteIwordBean) IWordContentList(siteId, siteIndexId string) (data []back.IndexContentWord, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	iword := new(schema.Iword)
	sess.Where("site_id = ?", siteId)
	sess.Where("site_index_id = ?", siteIndexId)
	sess.Where("state = ?", 1)
	sess.Where("itype >= ?", 7) //存款文案类型
	err = sess.Table(iword.TableName()).Find(&data)
	return
}

//查询文案单条

func (*SiteIwordBean) IwordInfo(this *input.SiteCopyInfo, siteId string) (bool, back.IndexWordInfo, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	iword := new(schema.Iword)
	var data back.IndexWordInfo
	sess.Where("id = ?", this.Id)
	sess.Where("itype = ?", this.Itype)
	sess.Where("site_id = ?", siteId)
	ok, err := sess.Table(iword.TableName()).Get(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, data, err
	}
	return ok, data, err
}

////添加文案
//func (*SiteIwordBean)IwordAdd(this *input.SiteIwordAdd)(count int64, err error){
//	sess := global.GetXorm().NewSession()
//	defer sess.Close()
//	iword := new(schema.Iword)
//	iword.AddTime = time.Now().Unix()
//	iword.TopId = 0
//	iword.SiteId = this.SiteId
//	iword.SiteIndexId = this.SiteIndexId
//	iword.State = 1
//	iword.Sort = this.Sort
//	iword.Title = this.Title
//	iword.Content = this.Content
//	iword.Itype = this.Itype
//	iword.TypeName = this.TypeName
//	iword.From = 1
//
//	count, err = sess.Table(iword.TableName()).InsertOne(iword)
//	return
//}

//修改文案

func (*SiteIwordBean) IwordEidt(this *input.IwordUpdate, SiteId string) (int64, int64, error) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	iword := new(schema.Iword)
	var code, count int64
	has, err := sess.Table(iword.TableName()).
		Where("id = ? and site_id = ? and itype=?", this.Id, SiteId, this.Itype).Get(iword)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if !has {
		code = 90716
		return code, count, err
	}

	iword.Content = this.Content
	//if len(this.Img) != 0 {
	//	iword.Img = this.Img
	//}

	count, err = sess.Table(iword.TableName()).
		Where("id = ?", this.Id).
		Cols("content, add_time").Update(iword)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//查询优惠文案列表

func (*SiteIwordBean) IwordActivityList(this *input.SiteActivityCopyList, SiteId string) (data []back.IndexActivityWord, err error) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()

	iword := new(schema.SiteActivity)
	sess.Where("site_id = ?", SiteId)
	sess.Where("site_index_id = ?", this.SiteIndexId)
	if this.From != 0 {
		sess.Where("`from`= ?", this.From)
	}
	sess.Where("top_id = ?", 0)
	sess.Where("delete_time = ?", 0)
	err = sess.Table(iword.TableName()).Find(&data)
	return
}

//查询优惠分类详情列表

func (*SiteIwordBean) IwordActivityType(this *input.SiteActivityCopyInfo, SiteId string) (
	[]back.IndexActivityWord, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.IndexActivityWord
	iword := new(schema.SiteActivity)
	sess.Where("site_id = ?", SiteId)
	sess.Where("site_index_id = ?", this.SiteIndexId)
	sess.Where("`from` = ?", this.From)
	sess.Where("top_id = ?", this.Id)
	sess.Where("type_name = ?", this.TypeName)
	sess.Where("delete_time = ?", 0)
	err := sess.Table(iword.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询优惠活动列表

func (*SiteIwordBean) IndexActivityList(siteId, siteIndexId string, from int8) (data []schema.SiteActivity, err error) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()

	iword := new(schema.SiteActivity)
	sess.Where("site_id = ?", siteId)
	sess.Where("site_index_id = ?", siteIndexId)
	sess.Where("`from` = ?", from)
	sess.Where("delete_time = ?", 0)
	sess.Asc("sort")
	err = sess.Table(iword.TableName()).Find(&data)
	return
}

//查询优惠详情

func (*SiteIwordBean) IwordActivityInfo(this *input.SiteActivityCopyInfo, SiteId string) (
	bool, back.IndexWordActivityInfo, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data back.IndexWordActivityInfo
	siteActivity := new(schema.SiteActivity)
	if this.From != 0 {
		sess.Where(siteActivity.TableName()+".from=?", this.From)
	}
	sess.Where("delete_time = ?", 0)
	sess.Where("id = ?", this.Id)
	sess.Where("site_id = ?", SiteId)
	ok, err := sess.Table(siteActivity.TableName()).Get(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, data, err
	}
	return ok, data, err
}

//添加优惠文案

func (*SiteIwordBean) IwordActivityAdd(this *input.ActivityEditeTitle, siteId string) (int64, error) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteActivity := new(schema.SiteActivity)
	siteActivity.Title = this.Title
	if this.Ctype == 1 {
		siteActivity.TypeName = this.Title
	}
	siteActivity.AddTime = time.Now().Unix()
	siteActivity.TopId = this.TopId
	siteActivity.SiteId = siteId
	siteActivity.SiteIndexId = this.SiteIndexId
	siteActivity.State = this.State
	siteActivity.Sort = this.Sort
	siteActivity.From = this.From
	siteActivity.DeleteTime = 0
	count, err := sess.Table(siteActivity.TableName()).InsertOne(siteActivity)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//优惠标题修改

func (*SiteIwordBean) ActivityEditeTitle(this *input.ActivityEditeTitle, siteId string) (int64, int64, error) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteActivity := new(schema.SiteActivity)
	var code, count int64
	has, err := sess.Table(siteActivity.TableName()).
		Where("id=? and delete_time = ? and site_id = ?", this.Id, 0, siteId).Get(siteActivity)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if !has {
		code = 90716
		return code, count, err
	}
	siteActivity.State = this.State
	siteActivity.Title = this.Title
	siteActivity.Sort = this.Sort
	siteActivity.From = this.From
	if this.Ctype == 1 && this.TopId == 0 {
		siteActivity.TypeName = this.Title

	} else {

		siteActivity.TypeName = ""
	}
	sess.Cols("title, state, type_name, sort")
	count, err = sess.Table(siteActivity.TableName()).
		Where("id=? and delete_time = ? and site_id = ?", this.Id, 0, siteId).
		Cols("title, state, type_name, sort").
		Update(siteActivity)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}

	return code, count, err
}

//优惠内容修改
func (*SiteIwordBean) ActivityEditeContent(this *input.ActivityEditeContent, siteId string) (int64, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var code, count int64
	siteActivity := new(schema.SiteActivity)
	has, err := sess.Table(siteActivity.TableName()).
		Where("id=? and delete_time = ? and site_id = ?", this.Id, 0, siteId).Get(siteActivity)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if !has {
		code = 90716
		return code, count, err
	}
	if this.Content != "" {
		siteActivity.Content = this.Content
		sess.Cols("content")
	}
	if this.Img != "" {
		siteActivity.Img = this.Img
		sess.Cols("img")
	}
	count, err = sess.Table(siteActivity.TableName()).
		Where("id=?", this.Id).
		Update(siteActivity)

	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}

	return code, count, err
}

//优惠删除
func (*SiteIwordBean) ActivityDel(Id int64, siteId string) (int64, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteActivity := new(schema.SiteActivity)
	var code, count int64
	has, err := sess.Table(siteActivity.TableName()).
		Where("id=? and delete_time = ? and site_id = ?", Id, 0, siteId).Get(siteActivity)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if !has {
		code = 90716
		return code, count, err
	}
	siteActivity.DeleteTime = time.Now().Unix()
	sess.Where("id=? and site_id = ?", Id, siteId)
	sess.Cols("delete_time")
	count, err = sess.Table(siteActivity.TableName()).Update(siteActivity)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//站点线路检测查询

func (*SiteIwordBean) SiteDetectList(siteId, siteIndexId string) ([]schema.InfoDomain, error) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []schema.InfoDomain
	infoDomain := new(schema.InfoDomain)
	sess.Where("site_id = ?", siteId)
	sess.Where("site_index_id = ?", siteIndexId)
	sess.Where("status=?", 1)
	err := sess.Table(infoDomain.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//站点线路检测修改

func (*SiteIwordBean) SiteDetectEdit(Id int64, SiteId, SiteIndexId, Domain string) (int64, int64, error) {
	var code, count int64
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoDomain := new(schema.InfoDomain)
	//检查数据是否存在
	has, err := sess.Table(infoDomain.TableName()).
		Where("id=? and site_id = ? and site_index_id = ? and status = ?", Id, SiteId, SiteIndexId, 1).
		Get(infoDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if !has {
		code = 90716
		return code, count, err
	}

	infoDomain.Domain = Domain
	count, err = sess.Table(infoDomain.TableName()).
		Where("id=? and site_id = ? and site_index_id = ? and status = ?", Id, SiteId, SiteIndexId, 1).
		Cols("domain").Update(infoDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}

	return code, count, err
}

//站点线路检测删除

func (*SiteIwordBean) SiteDetectDel(this *input.SiteDetectDel, SiteId string) (int64, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var code, count int64
	infoDomain := new(schema.InfoDomain)
	//检查数据是否存在
	has, err := sess.Table(infoDomain.TableName()).
		Where("id=? and site_id = ? and site_index_id = ? and status = ?", this.Id, SiteId, this.SiteIndexId, 1).
		Get(infoDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if !has {
		code = 90716
		return code, count, err
	}
	infoDomain.Status = 2
	count, err = sess.Table(infoDomain.TableName()).
		Where("id=? and site_id = ? and site_index_id = ? and status = ?", this.Id, SiteId, this.SiteIndexId, 1).
		Cols("status").Update(infoDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//添加站点线路检测

func (*SiteIwordBean) SiteDetectAdd(SiteId, SiteIndexId, Domain string) (int64, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var code, count int64
	infoDomain := new(schema.InfoDomain)
	has, err := sess.Table(infoDomain.TableName()).
		Where("site_id = ? and site_index_id = ? and domain = ?", SiteId, SiteIndexId, Domain).
		Get(infoDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if has {
		code = 90716
		return code, count, err
	}
	infoDomain.Domain = Domain
	infoDomain.SiteId = SiteId
	infoDomain.SiteIndexId = SiteIndexId
	infoDomain.Status = 1
	count, err = sess.Table(infoDomain.TableName()).InsertOne(infoDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//查询站点下文案列表
func (*SiteIwordBean) GetIWordList(siteId, siteIndexId string, Id int64) (data []back.IwordList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	detect := new(schema.Iword)
	var datalist []back.IwordList
	err = sess.Table(detect.TableName()).
		Select("id,top_id,site_id,site_index_id,title,title_color,content,url,img,state,sort,`from`,itype,type_name,add_time").
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Where("state=?", 1).
		Where("id=?", Id).
		Find(&data)
	if len(data) != 0 {
		err = sess.Table(detect.TableName()).
			Select("id,top_id,site_id,site_index_id,title,title_color,content,url,img,state,sort,`from`,`itype`,type_name,add_time").
			Where("site_id=?", siteId).
			Where("site_index_id=?", siteIndexId).
			Where("top_id=?", Id).
			Where("state=?", 1).
			OrderBy("sort").
			Find(&datalist)
		if err != nil {
			global.GlobalLogger.Error("err:s%", err.Error())
			return nil, err
		}
		if datalist != nil {
			data[0].IList = datalist
		}
	}
	for k, v := range data {
		data[k].ContentHtml = template.HTML(v.Content)
	}
	for k, v := range datalist {
		datalist[k].ContentHtml = template.HTML(v.Content)
	}
	return
}

//查询站点开会协议
func (*SiteIwordBean) GetAgreeMent(siteId, siteIndexId string) (data back.AgreeContent, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	detect := new(schema.Iword)
	has, err := sess.Table(detect.TableName()).
		Select("id,title,content").
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Where("state=?", 1).
		Where("itype=?", 6).
		Where("delete_time=?", 0).
		Get(&data)
	if has {
		return
	}
	if err != nil {
		return
	}
	return
}
