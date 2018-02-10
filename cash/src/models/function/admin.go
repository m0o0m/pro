package function

import (
	"fmt"
	"framework/logger"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type AdminBean struct{}

//平台账号列表
func (*AdminBean) GetList(this *input.AdminList, listparam *global.ListParams) (data []back.Admin, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.Admin)
	if this.Status != 0 {
		sess.Where("sales_admin.status = ?", this.Status)
	}
	if this.Account != "" {
		sess.Where("sales_admin.account = ?", this.Account)
	}
	if this.RoleId != 0 {
		sess.Where("sales_admin.role_id = ?", this.RoleId)
	}
	if this.OnlineStatus != 0 {
		sess.Where("sales_admin.online_status = ?", this.OnlineStatus)
	}
	sess.Where("sales_admin.delete_time = ?", 0)
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	//重新传入表名和where条件查询记录
	data = make([]back.Admin, 0)
	role := new(schema.Role)
	sql := fmt.Sprintf("%s.role_id = %s.id", admin.TableName(), role.TableName())
	err = sess.Table(admin.TableName()).Join("INNER", role.TableName(), sql).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	//获得符合条件的记录数
	count, err = sess.Table(admin.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//添加平台账号
func (*AdminBean) Add(this *input.AdminAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.Admin)
	admin.Account = this.Account
	admin.Password = this.Password
	admin.RoleId = this.RoleId
	admin.Status = this.Status
	admin.LoginIp = this.LoginIp
	admin.CreateTime = global.GetCurrentTime()
	count, err = sess.Table(admin.TableName()).InsertOne(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//平台账号详情
func (*AdminBean) GetInfo(this *input.AdminId) (a back.AdminInfo, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.Admin)
	role := new(schema.Role)
	sql := fmt.Sprintf("%s.role_id = %s.id", admin.TableName(), role.TableName())
	has, err = sess.Table(admin.TableName()).Join("INNER", role.TableName(), sql).
		Where("sales_admin.id = ?", this.Id).Where("sales_admin.delete_time = ?", 0).Get(&a)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return a, has, err
	}
	return a, has, err
}

//修改平台账号
func (*AdminBean) UpdateInfoNew(this *input.AdminEditNew) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.Admin)
	admin.Password = this.Password
	admin.RoleId = this.RoleId
	admin.Status = this.Status
	if admin.Status == 2 {
		admin.OnlineStatus = 2
		sess.Cols("login_key")
	}
	admin.LoginIp = this.LoginIp
	if this.Password != "" {
		sess.Cols("password")
	}
	count, err = sess.Where("id = ?", this.Id).Cols("status,login_ip,role_id").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改平台账号状态(新)
func (*AdminBean) UpdateStatusNew(this *input.AdminStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ad := new(schema.Admin)
	if this.Status == 1 {
		ad.Status = 2
		sess.Cols("login_key")
		ad.OnlineStatus = 2
	} else {
		ad.Status = 1
	}
	count, err := sess.Where("id = ?", this.Id).Cols("status", "online_status").Update(ad)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除平台账号
func (*AdminBean) Delete(this *input.AdminId) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.Admin)
	admin.Id = this.Id
	admin.DeleteTime = time.Now().Unix()
	count, err = sess.Table(admin.TableName()).Where("id = ?", this.Id).Cols("delete_time").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看账号是否存在
func (*AdminBean) GetAccount(account string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.Admin)
	has, err = sess.Table(admin.TableName()).Where("account= ?", account).Get(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看账号是否存在
func (*AdminBean) GetAccountById(id int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.Admin)
	has, err = sess.Where("id= ?", id).Where("delete_time = 0").Get(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//根据账号查询信息
func (*AdminBean) GetInfoByAcPa(account string) (schema.Admin, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.Admin
	flag, err := sess.Where("account=?", account).Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, err
}

//登录的时候刷新登录凭证，增加登录记录
func (*AdminBean) RefreshLoginKey(info schema.Admin, log schema.AdminLoginLog) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info.OnlineStatus = 1
	sess.Begin()
	count, err := sess.Where("id=?", info.Id).Cols("login_key", "online_status").Update(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	count, err = sess.InsertOne(&log)
	if err != nil || count == 0 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//根据id清除login_key
func (*AdminBean) UpAdminLoginStatus(adminId int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info := new(schema.Admin)
	info.LoginKey = ""
	count, err := sess.Where("id=?", adminId).Cols("login_key").Update(info)
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return count, err
	}
	return count, err
}

//根据token获取info
func (*AdminBean) GetInfoByToken(token string) (schema.Admin, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.Admin
	sess.Where("login_key=?", token)
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	flag, err := sess.Get(&info)
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return info, flag, err
	}
	return info, flag, err
}

//根据id获取平台管理员账号
func (*AdminBean) GetInfomationById(id int64) (schema.Admin, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.Admin
	sess.Where("id=?", id)
	flag, err := sess.Get(&info)
	return info, flag, err
}

//修改平台管理员密码
func (*AdminBean) ChangePassword(newAdmin *schema.Admin) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err := sess.Cols("password").Update(newAdmin)
	return count, err
}

//查询角色是否存在
func (*AdminBean) BeRoleId(role_id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ad := new(schema.Admin)
	has, err := sess.Where("id=?", role_id).Where("delete_time=?", 0).Get(ad)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//初始化管理员密码
func (*AdminBean) InitPassword(id int64, password string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.Admin)
	admin.Password = password
	count, err := sess.Where("id = ?", id).Cols("password").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
