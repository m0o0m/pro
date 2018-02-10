package function

import (
	"errors"
	"fmt"
	"framework/logger"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type RoleBean struct{}

//角色列表
func (*RoleBean) GetList(listparam *global.ListParams) (data []back.Role, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	//获得分页记录
	listparam.Make(sess)
	err = sess.Table(role.TableName()).Where("delete_time = ?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err = sess.Table(role.TableName()).Where("delete_time = ?", 0).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//添加角色
func (*RoleBean) Add(this *input.RoleAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	role.RoleName = this.RoleName
	//添加的角色都用admin标识
	role.RoleMark = "admin"
	role.Remark = this.Remark
	role.Status = this.Status
	count, err = sess.Table(role.TableName()).InsertOne(role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//角色详情
func (*RoleBean) GetInfo(this *input.RoleId) (role back.Role, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	r := new(schema.Role)
	has, err = sess.Table(r.TableName()).Where("id = ?", this.Id).Where("delete_time = ?", 0).Get(&role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return role, has, err
	}
	return role, has, err
}

//修改角色
func (*RoleBean) UpdateInfo(this *input.RoleEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	role.Id = this.Id
	role.RoleName = this.RoleName
	role.Remark = this.Remark
	role.Status = this.Status
	count, err = sess.Table(role.TableName()).Where("id = ?", this.Id).Cols("role_name,remark,status").Where("delete_time = ?", 0).Update(role)
	return
}

//修改角色(新)
func (*RoleBean) UpdateInfoNew(this *input.RoleEditNew) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	role.Id = this.Id
	role.RoleName = this.RoleName
	role.Remark = this.Remark
	count, err = sess.Table(role.TableName()).Where("id = ?", this.Id).Cols("role_name,remark").Where("delete_time = ?", 0).Update(role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return
}

//角色权限列表(查询角色和权限对应表)
func (*RoleBean) GetPermission(this *input.PromissRoleId) (data []back.RolePermission, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Type == 1 {
		this.Typed = "agency"
	} else if this.Type == 2 {
		this.Typed = "admin"
	}
	data = make([]back.RolePermission, 0)
	permission := new(schema.Permission)
	role := new(schema.Role)
	rolePermission := new(schema.RolePermission)
	if this.Typed != "" {
		sess.Where(permission.TableName()+".type = ?", this.Typed)
	}
	sql1 := fmt.Sprintf("%s.permission_id = %s.id", rolePermission.TableName(), permission.TableName())
	sql2 := fmt.Sprintf("%s.role_id = %s.id", rolePermission.TableName(), role.TableName())
	err = sess.Table(rolePermission.TableName()).Where(permission.TableName()+".delete_time = ?", 0).Where(rolePermission.TableName()+".role_id = ?", this.Id).Join("INNER", role.TableName(), sql2).Join("INNER", permission.TableName(), sql1).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//设置角色权限
func (*RoleBean) UpdatePermission(this *input.RolePermission) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	rolePermission := new(schema.RolePermission)
	role := new(schema.Role)
	role.Id = this.RoleId
	role.Status = this.Status
	role.RoleName = this.RoleName
	rolePermission.RoleId = this.RoleId
	//修改角色表
	if this.RoleId > 5 {
		count, err = sess.Table(role.TableName()).
			Where("id = ?", role.Id).
			Cols("status", "role_name").
			Update(role)
		if err != nil || count == 0 {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return
		}
	} else {
		count, err = sess.Table(role.TableName()).Where("id = ?", role.Id).Cols("role_name").Update(role)
		if err != nil || count == 0 {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return
		}
	}

	var rp schema.RolePermission
	var rps []schema.RolePermission
	//删除角色权限中间表数据
	count, err = sess.Table(rolePermission.TableName()).Where("role_id = ?", this.RoleId).Delete(rolePermission)
	if err != nil || count == 0 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	if len(this.PermissionId) > 0 {
		for k := range this.PermissionId {
			rp.RoleId = this.RoleId
			rp.PermissionId = this.PermissionId[k]
			rps = append(rps, rp)
		}
		_, err = sess.Table(rolePermission.TableName()).Insert(rps)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return
		}
	}
	sess.Commit()
	return
}

//删除角色
func (*RoleBean) Delete(this *input.RoleId) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	role.Id = this.Id
	role.DeleteTime = time.Now().Unix()
	count, err = sess.Table(role.TableName()).Where("id = ?", this.Id).Cols("delete_time").Update(role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//设置角色菜单
func (*RoleBean) UpdateMenu(this *input.RoleMenu) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	roleMenu := new(schema.RoleMenu)
	var rm schema.RoleMenu
	var rms []schema.RoleMenu
	roleMenu.RoleId = this.RoleId
	sess.Begin()
	//删除角色菜单中间表数据
	count, err = sess.Table(roleMenu.TableName()).Delete(roleMenu)
	if err != nil || count == 0 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	for k := range this.MenuId {
		rm.RoleId = this.RoleId
		rm.MenuId = this.MenuId[k]
		rms = append(rms, rm)
	}
	count, err = sess.Table(roleMenu.TableName()).Insert(rms)
	if err != nil || count == 0 {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	sess.Commit()
	return
}

//根据角色查菜单列表
func (*RoleBean) GetMenuByRoleId(this *input.RoleId) (data []back.MenuId, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	roleMenu := new(schema.RoleMenu)
	err = sess.Table(roleMenu.TableName()).Where("role_id = ?", this.Id).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查看角色是否存在(添加角色)
func IsRoleName(roleName string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	has, err = sess.Table(role.TableName()).Where("role_name= ?", roleName).Get(role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看角色是否存在（修改角色）
func IsRoleNames(id int64, roleName string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	has, err = sess.Table(role.TableName()).Where("id != ?", id).Where("role_name= ?", roleName).Get(role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看角色是否存在(角色权限)
func RoleIsExist(roleId int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	has, err = sess.Table(role.TableName()).Where("id= ?", roleId).Where("delete_time = ?", 0).Get(role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//修改角色状态
func (*RoleBean) UpdateStatus(this *input.RoleId) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	roles := new(schema.Role)
	sess.Begin()
	//查看角色状态
	_, err = sess.Table(role.TableName()).Where("id = ?", this.Id).Select("status").Where("delete_time = ?", 0).Get(role)
	if err != nil {
		sess.Rollback()
		return
	}
	//将查询的状态取反
	if role.Status == 1 {
		roles.Status = 2
	} else if role.Status == 2 {
		roles.Status = 1
	}
	count, err = sess.Table(role.TableName()).Where("id = ?", this.Id).Cols("status").Update(roles)
	if err != nil {
		sess.Rollback()
		return
	}
	sess.Commit()
	return
}

//修改角色状态(新)
func (*RoleBean) UpdateStatusNew(this *input.RoleStatus) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role := new(schema.Role)
	if this.Status == 1 {
		role.Status = 2
	} else {
		role.Status = 1
	}
	count, err = sess.Table(role.TableName()).Where("id = ?", this.Id).
		Cols("status").Update(role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//平台账号中的角色下拉框
func (*RoleBean) GetRoleList() (role []back.RoleList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	r := new(schema.Role)
	err = sess.Table(r.TableName()).Where("id > ?", 4).Where("status = ?", 1).Where("delete_time = ?", 0).Find(&role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return role, err
	}
	return role, err
}

//根据子账号id获取权限id
func (m *RoleBean) GetPermissionIdByAgencyId(agencyId int64) (permissionId int64, err error) {
	agencyPermissionSchema := new(schema.AgencyPermission)
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	has, err := sess.Table(agencyPermissionSchema.TableName()).
		Select("permission_id").
		Where("id = ?", agencyId).
		Get(&permissionId)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	if !(has) {
		err = errors.New(fmt.Sprintf("Not found agency permission by <%d>", agencyId))
	}
	return permissionId, err
}

//根据权限获取权限路由和方法
func (m *RoleBean) GetRoutesByPermissionId(permissionId int64) (list map[string]string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	permissionSchema := new(schema.Permission)
	list = make(map[string]string, 0)
	var routes []back.RouteList

	err = sess.Table(permissionSchema.TableName()).
		Select("route,method").
		Where("delete_time = 0").
		Where("status = 1").
		Where("type = agency").
		Find(&routes)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
	}
	for _, value := range routes {
		if value.Route != "" {
			list[value.Route] = value.Method
		}
	}
	return
}

//根据角色Id获取权限路由和方法
func (*RoleBean) GetRoutesByRole(id int64) (list map[string]string, err error) {
	//list = make(map[string]string, 0)
	//routes := make([]back.RouteList, 0)
	//rolePer := new(schema.RolePermission)
	//per := new(schema.Permission)
	//sess := global.GetXorm().NewSession().Table(per.TableName())
	//defer sess.Close()
	//sess.Where(rolePer.TableName()+".role_id=?", id)
	//sess.Where(per.TableName()+".delete_time=?", 0)
	//sess.Where(per.TableName()+".status=?", 1)
	//where1 := fmt.Sprintf("%s.permission_id = %s.id", rolePer.TableName(), per.TableName())
	//err = sess.Join("LEFT", rolePer.TableName(), where1).Find(&routes)
	//if err != nil {
	//	return
	//}
	//for index := range routes {
	//	list[routes[index].Route] = routes[index].Method
	//}
	return
}

//查看角色是否可以操作
func (*RoleBean) RoleIsOperate(roleId int64) (role schema.Role, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	_, err = sess.Table(role.TableName()).Select("is_operate").Where("id= ?", roleId).Where("delete_time = ?", 0).Get(&role)
	return
}

//查看角色是否可以操作(新)
func (*RoleBean) RoleIsOperateNew(roleId int64) (role schema.Role, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	has, err = sess.Table(role.TableName()).Where("id= ?", roleId).
		Where("delete_time = ?", 0).Get(&role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return role, has, err
	}
	return role, has, err
}

//根据id获取角色名，状态，是否可以禁用和删除
func (*RoleBean) GetRoleInfo(roleId int64) (role schema.Role, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	_, err = sess.Table(role.TableName()).Select("role_name,status,is_operate").Where("id= ?", roleId).Where("delete_time = ?", 0).Get(&role)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return role, err
	}
	return role, err
}

//查询角色是否被使用
func (*RoleBean) BeOneRole(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rl := new(schema.RolePermission)
	has, err := sess.Where("role_id=?", id).Get(rl)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//根据path,method和平台类型查询出权限
func (*RoleBean) GetPermissionByPM(path, method, t string) (schema.Permission, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var p schema.Permission
	ok, err := sess.Where("route=?", path).
		Where("method=?", method).
		Where("type=?", t).
		Where("status=1").
		Where("delete_time=0").Get(&p)
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return p, ok, err
	}
	return p, ok, err
}
