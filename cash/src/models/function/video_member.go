package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type VideoMemberBean struct{}

//获取视讯类型下拉框
func (*VideoMemberBean) GetVideoList() (backData []back.VideoMemberBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newProduct := new(schema.Product)
	sess.Where("delete_time=?", 0)
	sess.Where("type_id=?", 1)
	err = sess.Table(newProduct.TableName()).Find(&backData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, err
	}
	return backData, err
}

//视讯账号查询
func (*VideoMemberBean) SiteVideoMemberSearch(this *input.SiteVideoMemberSearch, listParam *global.ListParams) (data []back.VideoMemberSearch, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteVideoMember := new(schema.SiteVideoMember)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.Platform != "" {
		sess.Where("platform=?", this.Platform)
	}
	if this.Type == 1 { //平台用户名
		sess.Where("account = ?", this.Account)
	} else if this.Type == 2 { //视讯用户名
		sess.Where("g_username = ?", this.Account)
	}
	listParam.Make(sess)
	conds := sess.Conds()
	err = sess.Table(siteVideoMember.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(siteVideoMember.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}
