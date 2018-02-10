package schema

import "global"

//视讯额度掉单申请表
type SiteSingleRecord struct {
	Id             int64   `xorm:"id"`
	SiteId         string  `xorm:"site_id"`             //站点ID
	SiteIndexId    string  `xorm:"site_index_id"`       //多站点ID
	Username       string  `xorm:"username"`            //会员账号
	AdminUser      string  `xorm:"admin_user"`          //提交人
	Money          float64 `xorm:"money"`               //交易额度
	Ctype          int64   `xorm:"ctype"`               //转出方
	Vtype          int64   `xorm:"vtype"`               //转入方
	DoType         int8    `xorm:"do_type"`             //1系统转视讯掉单，扣除视讯额度   2视讯转系统掉单，加视讯额度
	DoTime         int64   `xorm:"do_time"`             //掉单时间
	CreateTime     int64   `xorm:"create_time created"` //申请时间
	UpdateTime     int64   `xorm:"update_time"`         //操作时间
	UpdateUsername string  `xorm:"update_username"`     //操作人
	Type           int8    `xorm:"type"`                //1表示掉单审核中，2表示审核通过，3无效申请
	Remark         string  `xorm:"remark"`              //备注
}

func (*SiteSingleRecord) TableName() string {
	return global.TablePrefix + "site_single_record"
}
