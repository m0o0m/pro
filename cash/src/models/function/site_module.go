package function

import (
	"global"
	"models/back"
	"models/schema"
)

type SiteModuleBean struct {
}

//查询所有维护信息配置
func (m *SiteModuleBean) GetModuleAll() (data []*back.SiteModule, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteModuleSchema := new(schema.SiteModule)
	err = sess.Table(siteModuleSchema.TableName()).
		Where("state = ?", 1).
		Find(&data)
	return
}
