package function

import (
	"errors"
	"global"
	"models/schema"
)

type SiteInfoVideoUser struct{}

//获取视讯选择模版ID
func (*SiteInfoVideoUser) LiveStyle(siteId, siteIndexId string) (LiveId int, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.InfoVideoUse)
	b, err := sess.Table(infoLogoSchema.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Where("state = ?", 1).
		Select("style").
		Get(&LiveId)
	if err != nil {
		return
	}
	if !b {
		err = errors.New("not found style")
	}
	return
}
