package function

import (
	"fmt"
	"global"
	"models/back"
	"models/schema"
)

type RoleMenuBean struct{}

//获取某个角色菜单导航(role_id:角色id,t:菜单类型)
func (*RoleMenuBean) GetMenuByRoleId(role_id int64, t string) (
	[]back.MenuListBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	roleMenu := new(schema.RoleMenu)
	menu := new(schema.Menu)
	var data []back.MenuListBack
	//join表条件(菜单表id = 角色菜单表menu_id)
	sess.Where(roleMenu.TableName()+".role_id=?", role_id)
	sess.Where(menu.TableName()+".type=?", t)
	sess.Where(menu.TableName() + ".delete_time=0")
	conds := sess.Conds()
	where := fmt.Sprintf("%s.id = %s.menu_id", menu.TableName(), roleMenu.TableName())
	err := sess.Table(roleMenu.TableName()).Join("LEFT", menu.TableName(), where).
		Asc("sort").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(roleMenu.TableName()).Where(conds).
		Join("LEFT", menu.TableName(), where).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//根据角色查菜单id(开户人及下属)
func (*RoleMenuBean) GetMenuByRoleIdAgency(this *schema.RoleMenu) (data []back.MenuListBack, err error) {
	sess := global.GetXorm().NewSession().Table(this.TableName())
	defer sess.Close()
	menu_schema := new(schema.Menu)
	var menu_list []back.MenuBack
	err = sess.Where("role_id=?", this.RoleId).Find(&menu_list)
	if err != nil {
		return
	}
	if menu_list == nil {
		return nil, err
	}
	var menuIds []int64
	for _, v := range menu_list {
		menuIds = append(menuIds, v.MenuId)
	}
	sess.Where("type=?", "agency")
	err = sess.Table(menu_schema.TableName()).In("id", menuIds).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		return
	}
	return
}

//根据角色查菜单id(平台)
func (*RoleMenuBean) GetMenuByRoleIdAdmin(this *schema.RoleMenu) (data []back.MenuListBack, err error) {
	sess := global.GetXorm().NewSession().Table(this.TableName())
	defer sess.Close()
	menu_schema := new(schema.Menu)
	var menu_list []back.MenuBack
	err = sess.Where("role_id=?", this.RoleId).Find(&menu_list)
	if err != nil {
		return
	}
	if menu_list == nil {
		return nil, err
	}
	var menuIds []int64
	for _, v := range menu_list {
		menuIds = append(menuIds, v.MenuId)
	}
	sess.Where("type=?", "admin")
	err = sess.Table(menu_schema.TableName()).In("id", menuIds).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		return
	}
	return
}

//根据ids查询
func (*RoleMenuBean) GetMenuByIds(data []int64) (bdata []back.MenuBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	role_menu := new(schema.RoleMenu)
	err = sess.Table(role_menu.TableName()).In("menu_id", data).Find(&bdata)
	return
}
