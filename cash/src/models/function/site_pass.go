package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type SitePassBean struct{}

//站点口令列表
func (*SitePassBean) SitePassList(this *input.SitePassList, listParams *global.ListParams) (sitePass []back.SitePass, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SitePasswordVerification)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	conds := sess.Conds()
	//获得符合条件的记录数
	count, err = sess.Table(spv.TableName()).Count()
	if count > 0 {
		listParams.Make(sess)
		err = sess.Table(spv.TableName()).Where(conds).Find(&sitePass)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return
		}
	}
	return
}

//站点口令修改 有则修改  无则添加
func (sp *SitePassBean) SitePassUpdate(this *input.SitePassUpdate) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spvfn := new(schema.SitePasswordVerification)
	has, err := sess.Where("site_id=?", this.SiteId).Get(spvfn)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	spv := new(schema.SitePasswordVerification)
	spv.SiteId = this.SiteId
	spv.Status = this.Status
	spv.Account = this.Account
	spv.PassKey = this.PassKey
	spv.UpdateTime = time.Now().Unix()
	if !has {
		count, err = sess.InsertOne(spv)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else {
		if spvfn.SiteId != spv.SiteId {
			sess.Cols("site_id")
		} else {
			sess.Where("site_id=?", spv.SiteId)
		}
		count, err = sess.Cols("status,account,pass_key,update_time").Update(spv)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	}
	return count, err
}

//查看站点是否存在
func (*SitePassBean) IsExistSite(siteId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	has, err = sess.Where("id=?", siteId).Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return
}

func (*SitePassBean) BatchDelChanges(this *input.BatchDelChanges) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	//判断所选站点是否存在
	site := new(schema.Site)
	siteList := []schema.Site{}
	err = sess.Table(site.TableName()).In("id", this.SiteId).Find(&siteList)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if len(siteList) > 0 {
		siteArr := []string{}
		for _, value := range siteList {
			siteArr = append(siteArr, value.Id)
		}
		sitePasswordVerification := new(schema.SitePasswordVerification)
		sitePasswordVerification.Account = this.Account
		sitePasswordVerification.UpdateTime = global.GetCurrentTime()
		sitePasswordVerification.Status = int(this.Status)
		count, err = sess.In("site_id", this.SiteId).Cols("status,account,update_time").Update(sitePasswordVerification)
		return
	}
	return
}
