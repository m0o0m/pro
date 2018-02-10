package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type PermissionBean struct{}

//获取权限中的module
func (*PermissionBean) GetModules(this *input.PromissRoleId) ([]back.Module, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Type == 1 {
		sess.Where("type=?", "agency")
	} else if this.Type == 2 {
		sess.Where("type=?", "admin")
	}
	if this.Typed != "" {
		sess.Where("type=?", this.Typed)
	}
	permssion := new(schema.Permission)
	var m []back.Module
	err := sess.Table(permssion.TableName()).
		Select("module").GroupBy("module").
		Where("delete_time = ?", 0).Find(&m)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return m, err
	}
	return m, err
}

//权限列表
func (*PermissionBean) GetListByModule(module []string) ([]back.Pmn, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permssion := new(schema.Permission)
	var data []back.Pmn
	err := sess.Table(permssion.TableName()).
		In("module", module).
		Where("delete_time = ?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//子帐号查看权限细项
func (*PermissionBean) GetDetailsMemberByChild(id int64) (*schema.DetailsMember, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	dm := new(schema.DetailsMember)
	has, err := sess.Where("child_id=?", id).Get(dm)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return dm, has, err
	}
	return dm, has, err
}

//查询站点
func (*PermissionBean) GetSiteIndexBySite(siteId, siteIndexId string) ([]back.SiteIndexBySiteBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	var data []back.SiteIndexBySiteBack
	if siteId != "" {
		sess.Where("id=?", siteId)
	}
	if siteIndexId != "" {
		sess.Where("index_id=?", siteIndexId)
	}
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	err := sess.Table(site.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//权限列表
func (*PermissionBean) GetLists(this *input.PermissionList, listparam *global.ListParams) (data []back.PermissionList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Type == 1 {
		sess.Where("type=?", "agency")
	} else if this.Type == 2 {
		sess.Where("type=?", "admin")
	}
	sess.Where("delete_time = ?", 0)
	condes := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	permssion := new(schema.Permission)
	err = sess.Table(permssion.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	//获得符合条件的记录数
	count, err = sess.Table(permssion.TableName()).Where(condes).Count()
	if err != nil {
		return data, count, err
	}
	return data, count, err
}

//添加权限
func (*PermissionBean) Add(this *input.PermissionAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permission := new(schema.Permission)
	permission.PermissionName = this.PermissionName
	permission.Method = this.Method
	permission.Route = this.Route
	permission.Module = this.Module
	permission.Status = this.Status
	if this.Type == 1 {
		permission.Type = "agency"
	} else if this.Type == 2 {
		permission.Type = "admin"
	}
	count, err = sess.Table(permission.TableName()).InsertOne(permission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//权限详情
func (*PermissionBean) GetInfo(this *input.PermissionId) (p back.PermissionList, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permission := new(schema.Permission)
	has, err = sess.Table(permission.TableName()).Where("id = ?", this.Id).Where("delete_time = ?", 0).Get(&p)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return p, has, err
	}
	return p, has, err
}

//修改权限(新接口)
func (*PermissionBean) ColumnUpdate(this *input.PermissionUpdate) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permission := new(schema.Permission)
	permission.Id = this.Id
	permission.PermissionName = this.PermissionName
	permission.Method = this.Method
	permission.Module = this.Module
	permission.Route = this.Route
	count, err := sess.Table(permission.TableName()).Where("id = ?", this.Id).
		Cols("module,permission_name,method,route").Update(permission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改权限状态(新)
func (*PermissionBean) ColumnStatus(this *input.ColumnStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	var count int64
	permission := new(schema.Permission)
	if this.Status == 1 {
		permission.Status = 2
	} else {
		permission.Status = 1
	}
	//修改权限状态
	count, err := sess.Table(permission.TableName()).
		Where("id = ?", this.Id).Cols("status").Update(permission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//如果是禁用权限,需要把中间表的数据同时删除
	if permission.Status == 2 {
		//删除角色权限中间表
		rolePermission := new(schema.RolePermission)
		rolePermission.PermissionId = this.Id
		count2, err := sess.Table(rolePermission.TableName()).Delete(rolePermission)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count2, err
		}
		//删除子账号权限中间表
		agencyPermission := new(schema.AgencyPermission)
		agencyPermission.PermissionId = this.Id
		count3, err := sess.Table(agencyPermission.TableName()).Delete(agencyPermission)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count3, err
		}
	}
	sess.Commit()
	return count, err
}

//删除权限(新）
func (*PermissionBean) ColumnDelete(this *input.ColumnDelete) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	permission := new(schema.Permission)
	permission.DeleteTime = time.Now().Unix()
	//删除权限
	count, err := sess.Table(permission.TableName()).
		Where("id = ?", this.Id).Cols("delete_time").Update(permission)
	if err != nil || count == 0 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//删除角色权限中间表
	rolePermission := new(schema.RolePermission)
	rolePermission.PermissionId = this.Id
	count, err = sess.Table(rolePermission.TableName()).Delete(rolePermission)
	if err != nil || count == 0 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//删除子账号权限中间表
	agencyPermission := new(schema.AgencyPermission)
	agencyPermission.PermissionId = this.Id
	count, err = sess.Table(agencyPermission.TableName()).Delete(agencyPermission)
	if err != nil || count == 0 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//查看权限名称(添加权限)
func (*PermissionBean) GetPermissionName(module, permissionName string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permission := new(schema.Permission)
	has, err = sess.Table(permission.TableName()).Where("module = ?", module).Where("permission_name = ?", permissionName).Where("delete_time = ?", 0).Get(permission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看权限名称（修改权限）
func (*PermissionBean) GetPermissionNames(id int64, module, permissionName string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permission := new(schema.Permission)
	has, err = sess.Table(permission.TableName()).Where("id != ?", id).Where("permission_name = ?", permissionName).Where("delete_time = ?", 0).Get(permission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//根据id查看权限是否存在
func IsPermissionById(ids []int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permission := new(schema.Permission)
	has, err := sess.Table(permission.TableName()).
		In("id", ids).Where("status = ?", 1).
		Where("delete_time = ?", 0).Get(permission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看路由和方法是否存在(添加)
func (*PermissionBean) GetRouteAndMethod(module, route, method string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permission := new(schema.Permission)
	has, err = sess.Table(permission.TableName()).Where("module = ?", module).Where("route = ?", route).Where("method = ?", method).Where("delete_time = ?", 0).Get(permission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看路由和方法是否存在（修改）
func (*PermissionBean) GetRouteAndMethods(id int64, module, route, method string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permission := new(schema.Permission)
	has, err = sess.Table(permission.TableName()).Where("id != ?", id).Where("module = ?", module).
		Where("route = ?", route).Where("method = ?", method).
		Where("delete_time = ?", 0).Get(permission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查询权限是否存在
func (*PermissionBean) BeOnePermission(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	perm := new(schema.Permission)
	has, err := sess.Where("id=?", id).Where("delete_time=?", 0).Get(perm)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查询权限是否被使用
func (*PermissionBean) BeOneColumn(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	perm := new(schema.RolePermission)
	has, err := sess.Where("permission_id=?", id).Get(perm)
	return has, err
}
