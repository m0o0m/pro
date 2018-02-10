package schema

import "global"

//SiteAgencyRegisterSet 站点代理申请设定表
type SiteAgencyRegisterSet struct {
	SiteId                string `xorm:"'site_id' PK"`             //站点id
	SiteIndexId           string `xorm:"'site_index_id' PK"`       //站点前台id
	RegisterProxy         int8   `xorm:"register_proxy"`           //是否启用代理注册(1.启用 2.不启用)
	ChineseNickname       int8   `xorm:"chinese_nickname"`         //是否需要中文昵称(1.需要 2.不需要)
	EnglishNickname       int8   `xorm:"english_nickname"`         //是否需要英文昵称(1.需要 2.不需要)
	NeedCard              int8   `xorm:"need_card"`                //是否需要证件号(1.需要 2.不需要)
	NeedEmail             int8   `xorm:"need_email"`               //是否需要邮箱(1.需要 2.不需要)
	NeedQq                int8   `xorm:"need_qq"`                  //是否需要qq(1.需要 2.不需要)
	NeedPhone             int8   `xorm:"need_phone"`               //是否需要手机号(1.需要 2.不需要)
	PromoteWebsite        int8   `xorm:"promote_website"`          //推广网址(1.需要 2.不需要)
	OtherMethod           int8   `xorm:"other_method"`             //其他方式(1.需要 2.不需要)
	IsMustChineseNickname int8   `xorm:"is_must_chinese_nickname"` //中文昵称是否必填(1.必须填2.非必填)
	IsMustEnglishNickname int8   `xorm:"is_must_english_nickname"` //英文昵称是否必填(1.必须填2.非必填)
	IsMustEmail           int8   `xorm:"is_must_email"`            //邮箱是否必填(1.必须填2.非必填)
	IsMustIdentity        int8   `xorm:"is_must_identity"`         //证件是否必填(1.必须填2.非必填)
	IsMustQq              int8   `xorm:"is_must_qq"`               //qq是否必填(1.必须填2.非必填)
	IsMustPhone           int8   `xorm:"is_must_phone"`            //手机号是否必填(1.必须填2.非必填)
	IsMustPromoteWebsite  int8   `xorm:"is_must_promote_website"`  //推广网址是否必填(1.必须填2.非必填)
	IsMustMethod          int8   `xorm:"is_must_method"`           //其他方式是否必填(1.必须填2.非必填)
}

func (*SiteAgencyRegisterSet) TableName() string {
	return global.TablePrefix + "site_agency_register_set"
}
