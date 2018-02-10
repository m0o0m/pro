package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SiteLogBean struct{}

//操作日志
func (*SiteLogBean) SiteDoLog(this *input.SiteDoLog, listParam *global.ListParams, times *global.Times) (data []back.SiteDoLog, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	adminLog := new(schema.AgencyLog)
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Key == 1 {
		sess.Where("operate_account = ?", this.Value)
	} else if this.Key == 2 {
		sess.Where("ip = ?", this.Value)
	}
	listParam.Make(sess)
	times.Make("operate_time", sess)
	conds := sess.Conds()
	err = sess.Table(adminLog.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(adminLog.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//登录日志
func (*SiteLogBean) SiteLoginLog(this *input.SiteLoginLog, listParam *global.ListParams, times *global.Times) (data []back.SiteLoginLog, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	loginLog := new(schema.LoginLog)
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Value != "" {
		if this.Key1 == 1 && this.Key2 == 1 { //会员账号
			sess.Where("login_role=1").And("account = ?", this.Value)
		} else if this.Key1 == 1 && this.Key2 == 2 { //会员ip
			sess.Where("login_role=1").And("login_ip = ?", this.Value)
		} else if this.Key1 == 2 && this.Key2 == 1 { //管理员账号
			sess.Where("login_role=2").And("account = ?", this.Value)
		} else if this.Key1 == 2 && this.Key2 == 2 { //管理员ip
			sess.Where("login_role=2").And("login_ip = ?", this.Value)
		}
	}
	listParam.Make(sess)
	times.Make("login_time", sess)
	conds := sess.Conds()
	err = sess.Table(loginLog.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(loginLog.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//自动稽核
func (sl *SiteLogBean) AutoAudit(this *input.AutoAudit, listParam *global.ListParams) (data []back.SiteLoginLog, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	loginLog := new(schema.LoginLog)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	cond := sess.Conds()
	if this.Device != 0 {
		sess.Where("device = ?", this.Device)
	}
	if this.Key == 1 { //会员账号
		sess.Where("account = ?", this.Value)
	} else if this.Key == 2 { //会员ip
		sess.Where("login_ip = ?", this.Value)
	}
	err = sess.Table(loginLog.TableName()).Find(&data) //第一次 ip
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	//sl.ipOrAccount(data,2)//取出结果中所有的ip
	sess.In("login_ip", sl.ipOrAccount(data, 1))
	data = nil
	err = sess.Table(loginLog.TableName()).Where(cond).Find(&data) //第一次 会员
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	sess.In("account", sl.ipOrAccount(data, 2))
	data = nil
	err = sess.Table(loginLog.TableName()).Where(cond).Find(&data) //第二次 ip
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	sess.In("login_ip", sl.ipOrAccount(data, 1))
	data = nil
	err = sess.Table(loginLog.TableName()).Where(cond).Find(&data) //第二次 会员
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	sess.In("account", sl.ipOrAccount(data, 2))
	data = nil
	sess.Where(cond)
	listParam.Make(sess)
	conds := sess.Conds()
	err = sess.Table(loginLog.TableName()).Find(&data) //第三次 ip
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(loginLog.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//ip查会员，会员查ip
func (*SiteLogBean) ipOrAccount(this []back.SiteLoginLog, ipOrAccount int8) (data []string) {
	if ipOrAccount == 1 { //会员查ip 返回ip集合
		for _, v := range this {
			data = append(data, v.LoginIp)
		}
	} else if ipOrAccount == 2 { //ip查会员 返回会员集合
		for _, v := range this {
			data = append(data, v.Account)
		}
	}
	return
}
