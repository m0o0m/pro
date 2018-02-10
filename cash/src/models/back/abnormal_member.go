package back

//异常会员查询
type AbnormalMemberList struct {
	SiteId      string `xorm:"site_id"`      //站点id
	Id          int64  `xorm:"id"`           //会员id
	Account     string `xorm:"account"`      //登录账号
	Card        string `xorm:"card"`         //卡号
	CardName    string `xorm:"card_name"`    //开户人姓名
	CardAddress string `xorm:"card_address"` //卡开户行
}
