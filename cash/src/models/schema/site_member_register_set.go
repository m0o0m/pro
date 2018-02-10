package schema

import (
	"global"
)

//SiteMemberRegisterSet  站点会员注册设定
type SiteMemberRegisterSet struct {
	SiteId      string  `xorm:"'site_id' PK" json:"siteId"`            //站点id
	SiteIndexId string  `xorm:"'site_index_id' PK" json:"siteIndexId"` //站点前台id
	IsReg       int8    `xorm:"is_reg" json:"isReg"`                   //是否启用会员注册 1:是2:否
	Email       int8    `xorm:"email" json:"email"`                    //注册是否需要邮箱 1:是2:否
	Passport    int8    `xorm:"passport" json:"passport"`              //是否需要身份证号 1:是2:否
	Wechat      int8    `xorm:"wechat" json:"wechat"`                  //是否需要微信
	Qq          int8    `xorm:"qq" json:"qq"`                          //是否需要qq 1:是2:否
	Mobile      int8    `xorm:"mobile" json:"mobile"`                  //是否需要手机号 1:是2:否
	Birthday    int8    `xorm:"birthday" json:"birthday"`              //是否需要出生日期 1:是2:否
	IsShowName  int8    `xorm:"is_show_name" json:"isShowName"`        //推广人姓名是否显示 1:是2:否
	IsName      int8    `xorm:"is_name" json:"isName"`                 //姓名是否可以重复 1:是2:否
	IsTel       int8    `xorm:"is_tel" json:"isTel"`                   //电话是否重复 1:是2:否
	IsEmail     int8    `xorm:"is_email" json:"isEmail"`               //邮箱是否可以重复 1:是2:否
	IsWechat    int8    `xorm:"is_wechat" json:"isWechat"`             //微信号是否可以重复
	IsQq        int8    `xorm:"is_qq" json:"isQq"`                     //qq是否可以重复 1:是2:否
	IsCardReply int8    `xorm:"is_card_reply" json:"isCardReply"`      //银行卡号是否可以重复 1:是2:否
	IsCode      int8    `xorm:"is_code" json:"isCode"`                 //验证码 1:是2:否
	IsWapSingle int8    `xorm:"is_wap_single" json:"isWapSingle"`      //wap注册是否单页面
	TryPlay     int8    `xorm:"try_play" json:"tryPlay"`               //试玩注册是否开启1.是2.不是
	Quota       float64 `xorm:"quota" json:"quota"`                    //试玩赠送额度【默认2000,最大100000】
	Offer       float64 `xorm:"offer" json:"offer"`                    //申请会员赠送优惠
	AddMosaic   int64   `xorm:"add_mosaic" json:"addMosaic"`           //优惠打码倍数
	IsIp        int8    `xorm:"is_ip" json:"isIp"`                     //限制ip
}

func (*SiteMemberRegisterSet) TableName() string {
	return global.TablePrefix + "site_member_register_set"
}
