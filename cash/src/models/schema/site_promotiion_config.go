package schema

import "global"

//自助优惠申请配置
type SitePromotionConfig struct {
	Id          int64  `xorm:"id"`            //主键id
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	ProTitle    string `xorm:"pro_title"`     //申请标题
	ProContent  string `xorm:"pro_content"`   //申请内容
	Createtime  int64  `xorm:"createtime"`    //申请时间
	Updatetime  int64  `xorm:"updatetime"`    //更新时间
	Status      int8   `xorm:"status"`        //1有效 2无效

}

func (*SitePromotionConfig) TableName() string {
	return global.TablePrefix + "site_promotion_config"
}
