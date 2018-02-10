package function

import (
	"errors"
	"global"
	"models/schema"
)

type SiteOrderModule struct{}

//获取视讯平台和排名
func (*SiteOrderModule) LiveModule(siteId, siteIndexId string) (LiveModule string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.SiteOrderModule)
	b, err := sess.Table(infoLogoSchema.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Select("video_module").
		Get(&LiveModule)
	if err != nil {
		return
	}
	if !b {
		err = errors.New("not found video_module")
	}
	return
}

//获取彩票平台和排名
func (*SiteOrderModule) LotteryModule(siteId, siteIndexId string) (FcModule string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.SiteOrderModule)
	b, err := sess.Table(infoLogoSchema.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Select("fc_module").
		Get(&FcModule)
	if err != nil {
		return
	}
	if !b {
		err = errors.New("not found fc_module")
	}
	return
}

//获取体育平台和排名
func (*SiteOrderModule) SportModule(siteId, siteIndexId string) (SportModule string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.SiteOrderModule)
	b, err := sess.Table(infoLogoSchema.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Select("sp_module").
		Get(&SportModule)
	if err != nil {
		return
	}
	if !b {
		err = errors.New("not found sport_module")
	}
	return
}
