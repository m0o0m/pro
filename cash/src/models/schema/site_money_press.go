package schema

import "global"

//催款账单
type SiteMoneyPress struct {
	Id         int64   `xorm:"id"`
	SiteId     string  `xorm:"site_id"`     //站点
	AddDate    int64   `xorm:"add_date"`    //添加时间
	UpdateDate int64   `xorm:"update_date"` //更新时间
	Qishu      int64   `xorm:"qishu"`       //新增时间
	Money      float64 `xorm:"money"`       //应缴金额
	PayName    string  `xorm:"pay_name"`    //收款人姓名
	PayAddress string  `xorm:"pay_address"` //收款地址
	Bank       string  `xorm:"bank"`        //收款银行
	PayCard    string  `xorm:"pay_card"`    //银行账号
	Remark     string  `xorm:"remark"`      //备注
	Status     int8    `xorm:"status"`      //1业主未提交,2业主已提交
	State      int8    `xorm:"state"`       //1已催款,2已完成催款
}

func (*SiteMoneyPress) TableName() string {
	return global.TablePrefix + "site_money_press"
}
