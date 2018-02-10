package schema

import "global"

//SiteAgencyRegister 站点代理申请表
type SiteAgencyRegister struct {
	Id             int64  `xorm:"'id' PK autoincr"`      //主键id
	SiteId         string `xorm:"site_id"`               //站点id
	SiteIndexId    string `xorm:"site_index_id"`         //站点前台id
	ZhName         string `xorm:"zh_name"`               //中文昵称
	UsName         string `xorm:"us_name"`               //英文昵称
	Card           string `xorm:"card"`                  //证件
	Email          string `xorm:"email"`                 //邮箱
	Account        string `xorm:"account"`               //账号
	Password       string `xorm:"password"`              //密码
	Qq             string `xorm:"qq"`                    //qq
	Wechat         string `xorm:"wechat"`                //微信号
	Skype          string `xorm:"skype"`                 //网络电话
	Phone          string `xorm:"phone"`                 //手机号
	UserName       string `xorm:"user_name"`             //真实姓名（银行开户人姓名）
	BackAccount    string `xorm:"back_account"`          //银行账号
	BankId         int64  `xorm:"bank_id"`               //开户银行
	OtherMethod    string `xorm:"other_method"`          //其他方法
	PromoteWebsite string `xorm:"promote_website"`       //推广网址
	Province       int64  `xorm:"province"`              //省份
	Zone           int64  `xorm:"zone"`                  //区域
	Remark         string `xorm:"remark"`                //备注
	Status         int8   `xorm:"status"`                //申请状态 1已添加账号2未处理'
	AgencyId       int64  `xorm:"agency_id"`             //添加的账号id 为0表示未添加到账号表
	CreateTime     int64  `xorm:"'create_time' created"` //申请时间
	UpdateTime     int64  `xorm:"update_time"`           //操作时间
	DeleteTime     int64  `xorm:"delete_time"`           //删除时间'
}

func (*SiteAgencyRegister) TableName() string {
	return global.TablePrefix + "site_agency_register"
}
