package input

//登录请求的struct
type MemberLogin struct {
	Account     string `json:"account" valid:"Required;MinSize(4);MaxSize(11);ErrorCode(20000)"`
	Password    string `json:"password" valid:"Required;MinSize(6);MaxSize(11);ErrorCode(20000)"`
	Code        string `json:"code" valid:"Required;MaxSize(4);ErrorCode(20021)"`        //验证码
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(30016)"` //站点前台id
}

//注册
type WapMemberRegister struct {
	Introducer       int64  `json:"introducer"`                                                              //代理注册
	IntroducerMember int64  `json:"introducerMember"`                                                        //会员推广注册
	Account          string `json:"account" valid:"Required;MaxSize(11);MinSize(4);ErrorCode(30009)"`        //账号
	Password         string `json:"password" valid:"Required;MaxSize(11);MinSize(6);ErrorCode(30010)"`       //密码
	RepeatPassword   string `json:"repeatPassword" valid:"Required;MaxSize(11);MinSize(6);ErrorCode(30010)"` //重复密码
	//RealName          string `json:"realName" valid:"Required;MinSize(1);ErrorCode(20011)"`                                               //真实姓名
	//OperatePassword   string `json:"operatePassword" valid:"MinSize(4);MaxSize(4);MinSize(4);ErrorCode(30015)"`                           //操作密码
	Site        string `json:"site" valid:"MaxSize(4);ErrorCode(60105)"`        //站点id
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(30016)"` //站点前台id
	//Code              string `json:"code"`                                                                                                //验证码
	IsAcceptAgreement int64 `json:"isAcceptAgreement" valid:"Required;Range(0,2);ErrorCode(60119)"` //是否同意协议(1.同意2.不同意)
	//Email             string `json:"email"`                                                                                               //邮箱
	//PassPort          string `json:"passPort" valid:"Match(^(\d{6})(\d{4})(\d{2})(\d{2})(\d{3})([0-9]|X)$);MaxSize(18);ErrorCode(50091)"` //身份证
	//Qq                string `json:"qq"`                                                                                                  //qq
	//Phone             string `json:"phone"`                                                                                               //电话
	//Birthday          string `json:"birthday" valid:"MaxSize(10);ErrorCode(20012)"`                                                       //生日
	//Wechat            string `json:"wechat"`                                                                                              //微信
	//BankCard          string `json:"bankCard"`                                                                                            //银行卡号
	//LocalCode         string `json:"localCode" valid:"MaxSize(4);ErrorCode(50120)"`                                                       //区号
}

//获取站点会员注册设定
type MemberRegisterSet struct {
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(30016)"` //站点前台id
}

type Home struct {
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(30016)"` //站点前台id
}
