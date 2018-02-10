package schema

import "global"

//站点额度记录表
type SiteCashRecord struct {
	Id          int64   `xorm:"id" xorm:"'id' PK autoincr"`
	SiteId      string  `xorm:"site_id"`       //站点id
	SiteIndexId string  `xorm:"site_index_id"` //站点前台id
	AdminName   string  `xorm:"admin_name"`    //操作账号，会员或登录人
	Money       float64 `xorm:"money"`         //交易额度
	Balance     float64 `xorm:"balance"`       //站点视讯余额
	CashType    int8    `xorm:"cash_type"`     //1额度转换  2额度加款  3额度扣款  4预借   5业主充值
	VdType      int64   `xorm:"vd_type"`       //视讯类型
	DoType      int8    `xorm:"do_type"`       //操作类型 1存入 2 取出
	CreateTime  int64   `xorm:"create_time"`   //操作时间
	State       int8    `xorm:"state"`         //1 正常  2 掉单
	Remark      string  `xorm:"remark"`        //备注
}

func (*SiteCashRecord) TableName() string {
	return global.TablePrefix + "site_cash_record"
}
