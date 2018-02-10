//子帐号管理[admin]
package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type ChildAccountBean struct{}

//子帐号列表
func (*ChildAccountBean) AcccountChildList(this *input.ChildAccountList, listparms *global.ListParams,
	times *global.Times) (data []back.ChildAccountListBack, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.Account != "" {
		sess.Where("account=?", this.Account)
	}
	if this.IsLogin != 0 {
		sess.Where("is_login=?", this.IsLogin)
	}
	if this.Name != "" {
		sess.Where("username=?", this.Name)
	}
	//开户人和代理
	if this.Type == 1 {
		sess.In("level", 1, 4).Where("is_sub=?", 2)
	} else if this.Type == 2 { //子账号
		sess.Where("is_sub=?", 1)
	} else if this.Type == 0 { //所有
		sess.In("level", 1, 4)
	}
	sess.Where("delete_time=?", 0)
	agency := new(schema.Agency)
	times.Make("create_time", sess)
	conds := sess.Conds()
	listparms.Make(sess)
	err = sess.Table(agency.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	//总数
	count, err = sess.Table(agency.TableName()).Where(conds).Count()
	if err != nil {
		return data, count, err
	}
	return
}

//获取一条子帐号信息
func (*ChildAccountBean) OneChildInfo(this *input.OneChildAccount) (data *back.OneChildInfo, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	data = new(back.OneChildInfo)
	has, err = sess.Table(agency.TableName()).Where("delete_time=?", 0).
		Where("id=?", this.Id).Where("site_id=?", this.SiteId).
		Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return
}

//修改子帐号
func (*ChildAccountBean) ChildInfoChange(this *input.ChildAccountChange) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if this.Password != "" {
		agency.Password = this.Password
		sess.Cols("password,username")
	}
	if this.Name != "" {
		agency.Username = this.Name
		sess.Cols("username")
	}
	count, err = sess.Where("id=?", this.Id).
		Where("site_id=?", this.SiteId).Update(agency)
	return
}

//修改子帐号的状态
func (*ChildAccountBean) ChildStatusChange(this *input.ChildAccountStatus) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if this.Status == 1 {
		agency.Status = 2
	} else if this.Status == 2 {
		agency.Status = 1
	}
	count, err = sess.Where("id=?", this.Id).
		Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询帐号
func (*ChildAccountBean) WhetherSubordinate(this *input.ChildAccountStatus) (data []schema.Agency, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	sess.Where("delete_time=?", 0)
	sess.Where("parent_id=?", this.Id)
	err = sess.Table(agency.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询子帐号是否存在
func (*ChildAccountBean) BeChildAccount(id int64, site string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if site != "" {
		sess.Where("site_id=?", site)
	}
	has, err = sess.Where("delete_time=?", 0).
		Where("is_sub=?", 1).
		Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//站点下拉框
func (*ChildAccountBean) SiteSiteIndexIdBy() ([]back.SiteSiteIndexBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	var data []back.SiteSiteIndexBack
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	err := sess.Table(site.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//子站点下拉框
func (*ChildAccountBean) SiteIndexIdList(siteId string) ([]back.SiteIndexId, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	var data []back.SiteIndexId
	err := sess.Table(site.TableName()).Where("id=?", siteId).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
