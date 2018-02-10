package function

import (
	"global"
	"models/back"
	"models/schema"
)

type RegisterStatusBean struct {
}

//验证是否开启注册

func (m *RegisterStatusBean) GetIsReg(siteId string, siteIndexId string) (*schema.SiteMemberRegisterSet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	detect := new(schema.SiteMemberRegisterSet)
	data := new(schema.SiteMemberRegisterSet)
	has, err := sess.Table(detect.TableName()).
		Select("is_reg").
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Get(data)
	return data, has, err
}

//
func (m *RegisterStatusBean) GetRegSet(siteId string, siteIndexId string) (data []back.MemberRegisterSetting, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	detect := new(schema.SiteMemberRegisterSet)
	err = sess.Table(detect.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Find(&data)
	return
}
