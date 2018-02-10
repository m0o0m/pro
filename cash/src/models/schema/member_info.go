package schema

import "global"

//会员资料表
type MemberInfo struct {
	MemberId  int64  `xorm:"member_id PK"` //会员id
	Card      string `xorm:"card"`         //身份证号
	LocalCode string `xorm:"local_code"`   //区号
	Mobile    string `xorm:"mobile"`       //手机号码
	Email     string `xorm:"email"`        //邮箱
	Qq        string `xorm:"qq"`           //qq
	Wechat    string `xorm:"wechat"`       //微信
	Birthday  int64  `xorm:"birthday"`     //生日
	Remark    string `xorm:"remark"`       //备注
}

func (*MemberInfo) TableName() string {
	return global.TablePrefix + "member_info"
}
