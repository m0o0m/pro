//总后台登录日志
package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type AdminLoginLogBean struct{}

//查询总后台登录日志
func (*AdminLoginLogBean) AdminLoginLogList(this *input.AdminLoginLog,
	listparam *global.ListParams, times *global.Times) ([]back.AdminLoginLogBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	a_l := new(schema.AdminLoginLog)
	role := new(schema.Role)
	if this.Ip != "" {
		sess.Where(a_l.TableName()+".login_ip=?", this.Ip)
	}
	if this.Account != "" {
		sess.Where(a_l.TableName()+".account=?", this.Account)
	}
	if this.RoleId != 0 {
		sess.Where(a_l.TableName()+".login_role=?", this.RoleId)
	}
	if this.Device != 0 {
		sess.Where(a_l.TableName()+".device=?", this.Device)
	}
	times.Make(a_l.TableName()+".login_time", sess)
	conds := sess.Conds()
	listparam.Make(sess)
	var data []back.AdminLoginLogBack
	err := sess.Table(a_l.TableName()).
		Join("LEFT", role.TableName(),
			a_l.TableName()+".login_role="+role.TableName()+".id").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(a_l.TableName()).
		Join("LEFT", role.TableName(),
			a_l.TableName()+".login_role="+role.TableName()+".id").
		Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}
