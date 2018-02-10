package schema

import "global"

//电子主题语优惠宽度设置表
type InfoActivityPromotionSet struct {
	Id          int64  `xorm:"'id' PK autoincr" json:"id"`         //文案id
	SiteId      string `xorm:"site_id" json:"site_id"`             //站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	MaxWidth    int64  `xorm:"max_width" json:"max_width"`         //优惠活动宽度
	Bcolor      string `xorm:"bcolor" json:"bcolor"`               //电子内页主题设置
}

func (*InfoActivityPromotionSet) TableName() string {
	return global.TablePrefix + "info_activity_promotion_set"
}
