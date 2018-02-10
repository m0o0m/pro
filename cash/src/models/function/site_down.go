package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SiteDownBean struct{}

//下载列表
func (*SiteDownBean) GetSiteDownList() (data []back.SiteDown, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	press := new(schema.SiteDown)
	err = sess.Table(press.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//添加下载地址
func (*SiteDownBean) SiteDownAdd(this schema.SiteDown) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteDown)
	count, err = sess.Table(admin.TableName()).InsertOne(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//单条下载地址详情
func (*SiteDownBean) GetInfo(this *input.SiteDown) (data back.SiteDown, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteDown)
	has, err = sess.Table(admin.TableName()).Where("id = ?", this.Id).Get(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//修改下载地址详情
func (*SiteDownBean) UpdateInfo(this *input.SiteDownEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteDown)
	admin.IosUrl = this.IosUrl
	admin.AndroidUrl = this.AndroidUrl
	admin.Vers = this.Vers
	admin.Platform = this.Platform
	count, err = sess.Where("id = ?", this.Id).Cols("ios_url,android_url,version,platform").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改下载地址状态
func (*SiteDownBean) UpdateStatus(this *input.SiteDownState) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteDown)
	admin.State = this.State
	count, err = sess.Where("id = ?", this.Id).Cols("state").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
