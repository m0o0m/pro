package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type SubAccountBeen struct{}

//子账号列表
func (*SubAccountBeen) GetList(this *input.SubAccountList, listParams *global.ListParams) ([]back.SubAccount, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	var data []back.SubAccount
	//判断并组合where条件
	sess.Where("site_id = ?", this.SiteId)
	sess.Where("delete_time = ?", 0).Where("is_sub = ?", 1)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.RoleId != 0 {
		sess.Where("role_id = ?", this.RoleId)
	}
	if this.Status == 1 || this.Status == 2 {
		sess.Where("status = ?", this.Status)
	}
	if this.ParentId != 0 {
		sess.Where("parent_id = ?", this.ParentId)
	}
	if this.Key == "username" {
		sess.Where("username = ?", this.Value)
	} else if this.Key == "account" {
		sess.Where("account = ?", this.Value)
	}
	if this.IsLogin == 1 {
		//查询在线
		sess.Where("is_login = ?", this.IsLogin)
	} else if this.IsLogin == 2 {
		//查询离线
		sess.Where("is_login = ?", this.IsLogin)
	}
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	//获得分页记录
	listParams.Make(sess)
	//重新传入表名和where条件查询记录
	err := sess.Table(agency.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	//获得符合条件的记录数
	count, err := sess.Table(agency.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//添加子账号
func (*SubAccountBeen) Add(this *input.SubAccountAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	//给账号表赋值
	agency.SiteId = this.SiteId
	agency.SiteIndexId = this.SiteIndexId
	agency.Account = this.Account
	agency.Username = this.Username
	agency.Password = this.Password
	agency.RoleId = this.RoleId
	agency.ParentId = this.ParentId
	agency.Level = this.Level
	agency.OperatePassword = this.OperatePassword
	agency.Status = this.Status
	agency.IsLogin = 2
	agency.IsSub = 1
	agency.IsDefault = 2
	count, err := sess.Table(agency.TableName()).Insert(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//子账号基本资料详情
func (*SubAccountBeen) GetInfo(id int64) (*back.SubAccountInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	subAccount := new(back.SubAccountInfo)
	has, err := sess.Table(agency.TableName()).
		Where("id = ?", id).Where("is_sub = ?", 1).
		Where("delete_time = ?", 0).Get(subAccount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return subAccount, has, err
	}
	return subAccount, has, err
}

//子账号基本资料详情
func (*SubAccountBeen) GetInfos(id int64, siteId, siteIndexId string) (*back.SubAccountInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	subAccount := new(back.SubAccountInfo)
	has, err := sess.Table(agency.TableName()).Where("id=?", id).Where("is_sub=?", 1).
		Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).
		Where("delete_time = ?", 0).Get(subAccount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return subAccount, has, err
	}
	return subAccount, has, err
}

//修改子账号基本资料(修改内容:密码，用户名，操作密码,状态)
func (*SubAccountBeen) UpdateInfo(this *input.SubAccountEdit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.Password = this.Password
	agency.OperatePassword = this.OperatePassword
	agency.Username = this.Username
	if agency.Password != "" {
		sess.Cols("password")
	}
	if agency.OperatePassword != "" {
		sess.Cols("operate_password")
	}
	count, err := sess.Table(agency.TableName()).
		Where("id = ?", this.Id).
		Cols("username").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//子账号权限列表
func (*SubAccountBeen) GetPermission(this *input.SubAccountPermission) ([]back.SubAgencyPermission, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	data := make([]back.SubAgencyPermission, 0)
	permission := new(schema.Permission)
	subAgency := new(schema.Agency)
	agencyPermission := new(schema.AgencyPermission)
	sql1 := fmt.Sprintf("%s.permission_id = %s.id", agencyPermission.TableName(), permission.TableName())
	sql2 := fmt.Sprintf("%s.agency_id = %s.id", agencyPermission.TableName(), subAgency.TableName())
	sess.Where("sales_agency_permission.agency_id = ?", this.Id)
	sess.Where("sales_agency.is_sub = ?", 1)
	sess.Where("sales_agency.parent_id = ?", this.ParentId)
	err := sess.Table(agencyPermission.TableName()).Join("LEFT", subAgency.TableName(), sql2).
		Join("LEFT", permission.TableName(), sql1).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//设置子账号权限
func (*SubAccountBeen) UpdatePermission(this *input.PermissionEdit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	subAgencyPermission := new(schema.AgencyPermission)
	var sap schema.AgencyPermission
	var saps []schema.AgencyPermission
	//删除子账号权限中间表数据
	count, err := sess.Table(subAgencyPermission.TableName()).Where("agency_id = ?", this.Id).
		Delete(subAgencyPermission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	for k := range this.PermissionId {
		sap.AgencyId = this.Id
		sap.PermissionId = this.PermissionId[k]
		saps = append(saps, sap)
	}
	count, err = sess.Table(subAgencyPermission.TableName()).Insert(saps)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//判断表中是否存在
	dm := new(schema.DetailsMember)
	has, err := sess.Table(dm.TableName()).Where("child_id=?", this.Id).Get(dm)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	dm.ChildPower = this.ChildPower
	dm.ChildSite = this.ChildSite
	if !has {
		//添加   判断count
		dm.ChildId = this.Id
		count, err = sess.Table(dm.TableName()).InsertOne(dm)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if count == 0 {
			return count, err
		}
	} else {
		//可能不会更改，因此不判断count
		count, err = sess.Table(dm.TableName()).Where("child_id=?", this.Id).
			Cols("child_power,child_site").
			Update(dm)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if count == 0 {
			return count, err
		}
	}
	sess.Commit()
	return count, err
}

//获取子账号细分权限
func (*SubAccountBeen) GetDetailPermission(id int64) (bool, schema.DetailsMember, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data schema.DetailsMember
	ok, err := sess.Where("child_id = ?", id).Get(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return ok, data, err
}

//设置子账号状态[启用/禁用]
func (*SubAccountBeen) UpdateStatus(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	this := new(schema.Agency)
	_, err := sess.Table(this.TableName()).Where("id = ?", id).Get(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	status := this.Status
	if status == 1 {
		this.Status = 2
	} else if status == 2 {
		this.Status = 1
	}
	count, err := sess.Table(this.TableName()).
		Where("id = ?", id).
		Cols("status").Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除子账号
func (*SubAccountBeen) Delete(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.DeleteTime = time.Now().Unix()
	count, err := sess.Table(agency.TableName()).
		Where("id = ?", id).Cols("delete_time").
		Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看账号是否存在
func (*SubAccountBeen) GetAccount(account string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	subAccount := new(schema.Agency)
	has, err := sess.Where("account= ?", account).Get(subAccount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看子账号是否存在(子账号权限)
func (*SubAccountBeen) SubAccountIsExist(this *input.PermissionEdit) (*schema.Agency, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	subAccount := new(schema.Agency)
	sess.Where("id = ?", this.Id)
	sess.Where("is_sub = ?", 1)
	sess.Where("site_id=?", this.SiteId)
	if this.ParentId != 0 {
		sess.Where("parent_id = ?", this.ParentId)
	}
	has, err := sess.Where("delete_time = ?", 0).Get(subAccount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return subAccount, has, err
	}
	return subAccount, has, err
}

//修改口令验证
func (sub *SubAccountBeen) SubAccessToken(this *input.SubAccountToken, account string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.SitePasswordVerification)

	//查询当前登录开户人所属site_id是否设置过口令验证
	has, err := sess.Table(agency.TableName()).Where("site_id=?", this.SiteId).Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	agency.Account = account
	agency.PassKey = this.PassKey
	agency.UpdateTime = time.Now().Unix()
	agency.Status = this.Status
	//如果存在则更新
	if has {
		count, err := sess.Table(agency.TableName()).Where("site_id=?", this.SiteId).
			Update(agency)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else {
		count, err := sess.Table(agency.TableName()).InsertOne(agency)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	}
	return 1, err
}

//根据站点信息查询口令验证信息
func (*SubAccountBeen) SubAccessTokenInfo(SiteId string) (
	*schema.SitePasswordVerification, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ag := new(schema.SitePasswordVerification)
	has, err := sess.Table(ag.TableName()).
		Where("site_id = ?", SiteId).Get(ag)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ag, has, err
	}
	return ag, has, err
}

//查看账号在代理申请表中是否存在
func (*SubAccountBeen) GetAccountByReg(account string) (*schema.SiteAgencyRegister, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	data := new(schema.SiteAgencyRegister)
	has, err := sess.Table(data.TableName()).Where("account= ?", account).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//查看id是否为股东
func (*SubAccountBeen) GetFirstAgent(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err := sess.Where("id=?", id).
		Where("role_id=2").
		Select("account").Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}
