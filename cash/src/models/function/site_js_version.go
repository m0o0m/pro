package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SiteJsVersionBean struct{}

//js版本列表
func (*SiteJsVersionBean) GetSiteJsVersionList() (data []back.SiteJsVersion, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	press := new(schema.SiteJsVersion)
	err = sess.Table(press.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//修改下载地址详情
func (*SiteJsVersionBean) UpdateInfo(this *input.SiteJsVersion) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteJsVersion)
	admin.Vers = this.Vers
	count, err = sess.Where("id = ?", this.Id).Cols("vers").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
