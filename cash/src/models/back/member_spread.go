package back

//会员返佣信息
type SpreadInfo struct {
	Id          int64   `xorm:"id" json:"id"`                    //推广Id
	Name        string  `xorm:"realname" json:"name"`            //姓名
	Account     string  `xorm:"account" json:"account"`          //账号
	CreateTime  string  `xorm:"create_time" json:"createTime"`   //创建时间
	RegisterIp  string  `xorm:"register_ip" json:"registerIp"`   //注册Ip
	Status      int     `xorm:"status" json:"status"`            //状态
	Number      int64   `xorm:"-" json:"number"`                 //推荐会员数
	SpreadMoney float64 `xorm:"spread_money" json:"spreadMoney"` //推广获利
}
