package schema

import "global"

//站点对应各平台占成比
type SiteLevelPlatform struct {
	LevelId    int64   `xorm:"level_id"`    //站点层级id
	PlatformId int64   `xorm:"platform_id"` //平台id
	Proportion float64 `xorm:"proportion"`  //占成比
}

func (*SiteLevelPlatform) TableName() string {
	return global.TablePrefix + "site_level_platform"
}
