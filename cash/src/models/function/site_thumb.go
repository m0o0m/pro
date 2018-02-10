package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type SiteThumbBean struct{}

//附件列表
func (*SiteThumbBean) List(l *input.GetSiteThumb, listParams *global.ListParams) (
	[]back.SiteThumb, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site_thumb := new(schema.SiteThumb)
	var data []back.SiteThumb
	sess.Where("site_id=?", l.SiteId)
	sess.Where("delete_time=0")
	if l.SiteIndexId != "" {
		sess.Where("site_index_id=?", l.SiteIndexId)
	}
	if l.State != 0 {
		sess.Where("state=?", l.State)
	}

	conds := sess.Conds()
	listParams.Make(sess)

	err := sess.Table(site_thumb.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(site_thumb.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//附件修改
func (*SiteThumbBean) SiteThumbUpdate(l *input.SiteThumbEdit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", l.Id)
	sess.Where("site_id=?", l.SiteId)
	if l.SiteIndexId != "" {
		sess.Where("site_index_id=?", l.SiteIndexId)
	}
	sitethum := new(schema.SiteThumb)
	sitethum.FileName = l.FileName
	count, err := sess.Table(sitethum.TableName()).Cols("file_name").Update(sitethum)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//附件删除
func (*SiteThumbBean) SiteThumbDelete(l *input.SiteThumbDelete) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", l.Id)
	sess.Where("site_id=?", l.SiteId)
	if l.SiteIndexId != "" {
		sess.Where("site_index_id=?", l.SiteIndexId)
	}
	sitethum := new(schema.SiteThumb)
	sitethum.DeleteTime = time.Now().Unix()
	count, err := sess.Table(sitethum.TableName()).Cols("delete_time").Update(sitethum)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
