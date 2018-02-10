package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type MenuBean struct{}

//查询菜单
func (*MenuBean) GetOneMenuByName(this *input.MenuInfo) (data *back.MenuListBack, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu := new(schema.Menu)
	if this.MenuName != "" {
		sess.Where("menu_name=?", this.MenuName)
	}
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	if this.Type != "" {
		sess.Where("type=?", this.Type)
	}
	data = new(back.MenuListBack)
	has, err = sess.Table(menu.TableName()).Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//添加菜单
func (*MenuBean) AddMenu(this *input.AddMenu) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu_info := new(schema.Menu)
	menu_info.MenuName = this.MenuName
	menu_info.Icon = this.Icon
	menu_info.Route = this.Route
	menu_info.ParentId = this.ParentId
	menu_info.Sort = this.Sort
	menu_info.Status = 1
	menu_info.Level = this.Level
	menu_info.LanguageKey = this.LanguageKey
	menu_info.Type = this.Type
	count, err = sess.InsertOne(menu_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//更新菜单
func (*MenuBean) UpdataMenu(this *input.UpdataMenu) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu_info := new(schema.Menu)
	menu_info.Id = this.Id
	menu_info.MenuName = this.MenuName
	menu_info.Icon = this.Icon
	menu_info.Route = this.Route
	menu_info.ParentId = this.ParentId
	menu_info.Sort = this.Sort
	menu_info.LanguageKey = this.LanguageKey
	count, err = sess.Where("id=?", this.Id).
		Cols("language_key,menu_name,route,sort,icon,parent_id").
		Update(menu_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//菜单列表(平台)
func (*MenuBean) FindAllMenu(this *schema.Menu) ([]back.MenuListBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.MenuListBack
	err := sess.Table(this.TableName()).Where("delete_time=?", 0).
		Asc("sort").Where("type=?", this.Type).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取一级、二级菜单的id和名称
func (*MenuBean) GetIdMenuName(this *input.MenuInfo) (data []back.MenuIdNameBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Type != "" {
		sess.Where("type=?", this.Type)
	}
	if this.Level != 0 {
		sess.Where("level=?", this.Level)
	}
	menu := new(schema.Menu)
	err = sess.Table(menu.TableName()).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//根据id查看菜单是否存在
func IsMenuById(ids []int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu := new(schema.Menu)
	has, err = sess.Table(menu.TableName()).In("id", ids).Where("status = ?", 1).Where("delete_time = ?", 0).Get(menu)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查询一个id的下级id
func (*MenuBean) GetNextIdById(this *input.MenuOpenClose) (data []int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var me []schema.Menu
	menu := new(schema.Menu)
	err = sess.Table(menu.TableName()).Where("parent_id=?", this.Id).
		Where("status=?", 1).
		Where("delete_time=?", 0).
		Find(&me)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if me != nil {
		for _, v := range me {
			data = append(data, v.Id)
		}
	}
	return
}

//菜单禁用
func (*MenuBean) CloseMenu(this *input.MenuOpenClose) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu := new(schema.Menu)
	menu.Status = 2
	count, err = sess.Table(menu.TableName()).Where("id=?", this.Id).Cols("status").Update(menu)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//菜单开启
func (*MenuBean) OpenMenu(this *input.MenuOpenClose) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu := new(schema.Menu)
	menu.Status = 1
	count, err = sess.Where("id=?", this.Id).Cols("status").Update(menu)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//菜单删除
func (*MenuBean) MenuDeleteTime(this *input.MenuOpenClose) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu := new(schema.Menu)
	menu.DeleteTime = time.Now().Unix()
	count, err = sess.Table(menu.TableName()).Where("id=?", this.Id).Cols("delete_time").Update(menu)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//菜单列表
func (*MenuBean) FindMenu(types string) (data []back.MenuListBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu := new(schema.Menu)
	err = sess.Table(menu.TableName()).Where("type = ?", types).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询一条菜单
func (*MenuBean) BeOneMenu(id int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	menu := new(schema.Menu)
	has, err = sess.Where("id=?", id).Get(menu)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}
