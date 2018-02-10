package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

//后台人员操作日志
type AdminLogBean struct {
}

//查询客户后台操作日志
func (*AdminLogBean) FindAdminLogList(this *input.AdminLog, listParams *global.ListParams, times *global.Times) (
	[]*back.AdminLog, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	adminLogSchema := new(schema.AdminLog)
	if this.Account != "" {
		sess.Where("operate_account = ?", this.Account)
	}
	if this.Ip != "" {
		sess.Where("ip = ?", this.Ip)
	}
	if this.OperatePath != "" {
		sess.Where("operate_path = ?", this.OperatePath)
	}
	if this.Type != 0 {
		sess.Where("type=?", this.Type)
	}
	times.Make("operate_time", sess)
	conds := sess.Conds()
	listParams.Make(sess)
	var data []*back.AdminLog
	err := sess.Table(adminLogSchema.TableName()).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(adminLogSchema.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}
