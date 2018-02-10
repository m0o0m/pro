package back

//返回站点会员注册设定详情
type MemberRegisterSetting struct {
	IsReg       int8    ` xorm:"is_reg" json:"isReg"  `             //是否启用会员注册
	Email       int8    ` xorm:"email" json:"email" `               //是否需要邮箱
	Wechat      int8    ` xorm:"wechat" json:"wechat" `             //是否需要微信
	Passport    int8    ` xorm:"passport" json:"passport" `         //是否需要身份证
	Qq          int8    ` xorm:"qq"  json:"qq" `                    //是否需要qq
	Mobile      int8    ` xorm:"mobile" json:"mobile" `             //是否需要电话
	Birthday    int8    ` xorm:"birthday" json:"birthday" `         //是否需要出年生日
	IsName      int8    ` xorm:"is_name" json:"isName" `            //是否需要真实姓名
	IsShowName  int8    ` xorm:"is_show_name" json:"isShowName" `   //是否显示推广人姓名
	IsCardReply int8    ` xorm:"is_card_reply" json:"isCardReply" ` //银行卡号是否可以重复
	IsTel       int8    ` xorm:"is_tel"  json:"isTel" `             //电话是否可以重复
	IsEmail     int8    ` xorm:"is_email" json:"isEmail" `          //邮箱是否可以重复
	IsQq        int8    ` xorm:"is_qq" json:"isQq" `                //qq是否可以重复
	IsWechat    int8    ` xorm:"is_wechat" json:"isWechat" `        //微信是否可以重复
	IsWapSingle int8    ` xorm:"is_wap_single" json:"isWapSingle"`  //wap注册是否单页面
	TryPlay     int8    ` xorm:"try_play" json:"tryPlay"`           //试玩注册是否开启
	Quota       float64 ` xorm:"quota" json:"quota"`                //试玩赠送额度
	IsCode      int8    ` xorm:"is_code" json:"isCode" `            //是否需要验证码
	Offer       float64 ` xorm:"offer" json:"offer"`                //会员注册优惠金额
	AddMosaic   int64   ` xorm:"add_mosaic" json:"addMosaic"`       //打码
	IsIp        int8    ` xorm:"is_ip" json:"isIp"`                 //是否开启限制Ip
}

//返回注册设定 wap
type WapRegSet struct {
	RegSet    []MemberRegisterSetting `json:"regSet"`    //注册设定
	AgencyId  int64                   `json:"agencyId"`  //推荐id
	Agreement string                  `json:"agreement"` //协议
}
