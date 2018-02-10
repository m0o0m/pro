package schema

import "global"

//站点退佣退水设定
type SiteRebateSet struct {
	Id            int64  `xorm:"'id' PK autoincr"`
	SiteId        string `xorm:"site_id"`        //操作站点id
	SiteIndexId   string `xorm:"site_index_id"`  //站点前台id
	SelfProfit    int64  `xorm:"self_profit"`    //自身盈利金额
	EffectiveUser int64  `xorm:"effective_user"` //有效会员数
	ValidMoney    int64  `xorm:"valid_money"`    //有效会员投注金额
	DeleteTime    int64  `xorm:"delete_time"`    //删除时间
}

func (*SiteRebateSet) TableName() string {
	return global.TablePrefix + "site_rebate_set"
}
