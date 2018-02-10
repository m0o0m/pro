package function

import (
	"global"
	"models/back"
	"models/schema"
)

type Maintain struct{}

// 获取维护信息
func (*Maintain) GetMaintainData() (data []back.MaintainData, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	err = sess.Table(new(schema.MaintainData).TableName()).Find(&data)

	return
}

func (*Maintain) MaintainHas(item schema.MaintainData) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	obj := sess
	obj = obj.Where("mtype=?", item.MType)
	obj = obj.Where("ctype=?", item.CType)
	if item.LindId != "" {
		obj = obj.Where("line_id=?", item.LindId)
	}
	has, err = obj.Get(&item)

	return
}

func (*Maintain) InsertMaintain(item schema.MaintainData) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	count, err = sess.Insert(&item)
	return
}

func (*Maintain) UpdateMaintain(item schema.MaintainData) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	obj := sess
	obj = obj.Where("mtype=?", item.MType)
	obj = obj.Where("ctype=?", item.CType)
	if item.LindId != "" {
		obj = obj.Where("line_id=?", item.LindId)
	}
	if item.SiteId != "" {
		obj = obj.Where("site_id=?", item.SiteId)
	}
	count, err = obj.Update(&item)
	return
}

func (*Maintain) DelMaintain(item schema.MaintainData) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	obj := sess
	obj = obj.Where("mtype=?", item.MType)
	obj = obj.Where("ctype=?", item.CType)
	if item.LindId != "" {
		obj = obj.Where("line_id=?", item.LindId)
	}
	if item.SiteId != "" {
		obj = obj.Where("site_id=?", item.SiteId)
	}
	count, err = obj.Delete(&item)
	return
}
