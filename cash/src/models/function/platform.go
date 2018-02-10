package function

import (
	//"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type PlatformBean struct{}

// 获取交易平台列表
func (*PlatformBean) GetPlatformList(plat *input.PlatformList, listparam *global.ListParams) (data []back.Platform, count int64, err error) {
	session := global.GetXorm().NewSession()
	defer session.Close()
	platform := new(schema.Platform)
	if plat.Status != 0 {
		session.Where("sales_platform.status = ?", plat.Status)
	}
	if plat.Platform != "" {
		session.Where("sales_platform.platform = ?", plat.Platform)
	}
	conds := session.Conds()
	//获得分页记录
	listparam.Make(session)

	//获得符合条件的记录
	err = session.Table(platform.TableName()).Find(&data)
	//获得符合条件的记录数
	count, err = session.Table(platform.TableName()).Where(conds).Count()
	return
}

//获取全部交易平台列表
func (m *PlatformBean) GetPlatformFull() (data []back.Platform, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platformSchema := new(schema.Platform)
	err = sess.Table(platformSchema.TableName()).
		Where("status = 1").
		Where("delete_time = 0").
		Find(&data)
	return
}

// 判断交易平台是否存在
func (*PlatformBean) PlatformIsExist(plat string) (has bool, err error) {
	session := global.GetXorm().NewSession()
	defer session.Close()
	platform := new(schema.Platform)
	//获得符合条件的记录
	has, err = session.Table(platform.TableName()).Where("platform = ?", plat).Get(platform)
	return
}

// 获取交易平台详情
func (*PlatformBean) GetPlatformOne(id int64) (data back.Platform, err error) {
	session := global.GetXorm().NewSession()
	defer session.Close()
	platform := new(schema.Platform)
	//获得符合条件的记录
	_, err = session.Table(platform.TableName()).Where("id = ?", id).Get(&data)
	return
}

// 添加交易平台
func (*PlatformBean) AddPlatform(plat *input.AddPlatform) (count int64, err error) {
	session := global.GetXorm().NewSession()
	defer session.Close()
	platform := new(schema.Platform)
	platform.Platform = plat.Platform
	platform.Status = plat.Status
	//获得符合条件的记录
	count, err = session.Table(platform.TableName()).InsertOne(platform)
	return
}

// 修改交易平台
func (*PlatformBean) UpdatePlatform(plat *input.UpdatePlatform) (count int64, err error) {
	session := global.GetXorm().NewSession()
	defer session.Close()
	platform := new(schema.Platform)
	platform.Platform = plat.Platform
	platform.Status = plat.Status
	//获得符合条件的记录
	count, err = session.Table(platform.TableName()).Where("id = ?", plat.Id).Cols("platform, status").Update(platform)
	return
}

// 修改交易平台状态
func (*PlatformBean) UpdatePlatformStatus(plat *input.UpdatePlatformStatus) (count int64, err error) {
	session := global.GetXorm().NewSession()
	defer session.Close()
	platform := new(schema.Platform)
	platform.Status = plat.Status
	//获得符合条件的记录
	count, err = session.Table(platform.TableName()).Where("id = ?", plat.Id).Cols("status").Update(platform)
	return
}

// 删除交易平台
func (*PlatformBean) DeletePlatform(plat *input.DeletePlatform) (count int64, err error) {
	session := global.GetXorm().NewSession()
	defer session.Close()
	platform := new(schema.Platform)
	platform.Status = 3
	//获得符合条件的记录
	count, err = session.Table(platform.TableName()).Where("id = ?", plat.Id).Cols("status").Update(platform)
	return
}

//视讯项目项目是否存在
func (m *PlatformBean) IsExist(platformId int64) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Platform)
	b, err := sess.Table(platform.TableName()).
		Select("count(*) as count").
		Where("status = 1").
		Where("delete_time = 0").
		Where("id = ?", platformId).
		Get(&count)
	if !(b) {
		count = 0
	}
	return
}
