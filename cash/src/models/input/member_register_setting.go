package input

//修改或者添加站点会员注册设定详情
type MemberRegisterSetting struct {
	SiteId      string  `json:"siteId"`                                                          //站点Id
	SiteIndexId string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`                 //站点前台Id
	IsReg       int8    ` json:"isReg" valid:"Range(1,2);ErrorCode(10038)"`                      //是否启用会员注册
	Email       int8    ` json:"email" valid:"Range(1,2);ErrorCode(10032)"`                      //是否需要邮箱
	Wechat      int8    ` json:"wechat" valid:"Range(1,2);ErrorCode(10034)"`                     //是否需要微信
	Passport    int8    ` json:"passport" valid:"Range(1,2);ErrorCode(10035)"`                   //是否需要身份证
	Qq          int8    ` json:"qq" valid:"Range(1,2);ErrorCode(10033)"`                         //是否需要qq
	Mobile      int8    ` json:"mobile" valid:"Range(1,2);ErrorCode(10036)"`                     //是否需要电话
	Birthday    int8    ` json:"birthday" valid:"Range(1,2);ErrorCode(10037)"`                   //是否需要出年生日
	IsName      int8    ` json:"isName" valid:"Range(1,2);ErrorCode(10039)"`                     //姓名是否重复
	IsShowName  int8    ` json:"isShowName" valid:"Range(1,2);ErrorCode(10040)"`                 //是否显示推广人姓名
	IsCardReply int8    ` json:"isCardReply" valid:"Range(1,2);ErrorCode(10041)"`                //银行卡号是否可以重复
	IsTel       int8    ` json:"isTel" valid:"Range(1,2);ErrorCode(10042)"`                      //电话是否可以重复
	IsEmail     int8    ` json:"isEmail" valid:"Range(1,2);ErrorCode(10043)"`                    //邮箱是否可以重复
	IsQq        int8    ` json:"isQq" valid:"Range(1,2);ErrorCode(10044)"`                       //qq是否可以重复
	IsWechat    int8    ` json:"isWechat" valid:"Range(1,2);ErrorCode(10045)"`                   //微信是否可以重复
	IsWapSingle int8    `json:"isWapSingle" valid:"Range(1,2);ErrorCode(60074)"`                 //wap注册是否单页面
	TryPlay     int8    `json:"tryPlay" valid:"Range(1,2);ErrorCode(60073)"`                     //试玩注册是否开启
	Quota       float64 `json:"quota" valid:"Required;ErrorCode(60072)"`                         //试玩赠送额度
	IsCode      int8    ` json:"isCode" valid:"Range(1,2);ErrorCode(10046)"`                     //是否可以需要验证码
	Offer       float64 ` json:"offer" valid:"Match(/(?!0\.00)(\d+\.\d{2}$)/);ErrorCode(10047)"` //会员注册优惠金额
	AddMosaic   int64   ` json:"addMosaic" valid:"Range(1,10000);ErrorCode(10048)"`              //打码(倍数)
	IsIp        int8    ` json:"isIp" valid:"Range(1,2);ErrorCode(10049)"`                       //是否开启限制Ip
}

//获取站点下某个注册设定
type MemberRegisterSettingGet struct {
	SiteId      string ` query:"siteId"`
	SiteIndexId string ` query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
}
