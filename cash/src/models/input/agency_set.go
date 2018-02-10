package input

//代理注册设定(查看)
type AgencyRegisterSet struct {
	SiteId      string `query:"site_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       //站点id
	SiteIndexId string `query:"site_index_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//代理注册设定
type AgencyRegister struct {
	SiteId          string `json:"site_id"`                                                                 //站点id
	SiteIndexId     string `json:"site_index_id"`                                                           //站点前台id
	ChineseNickname string `json:"chinese_nickname"`                                                        //中文昵称
	EnglishNickname string `json:"english_nickname"`                                                        //英文昵称
	Card            string `json:"card"`                                                                    //证件号
	Email           string `json:"email"`                                                                   //邮箱
	Qq              string `json:"qq"`                                                                      //qq
	Phone           string `json:"phone"`                                                                   //手机号
	PromoteWebsite  string `json:"promote_website"`                                                         //推广网址
	OtherMethod     string `json:"other_method"`                                                            //其他方式
	AgencyAccount   string `json:"agency_account" valid:"Required;MinSize(4);MaxSize(12);ErrorCode(30009)"` //代理账号
	Password        string `json:"password" valid:"Required;MinSize(6);MaxSize(12);ErrorCode(30010)"`       //密码
	ConfirmPassword string `json:"confirm_password" valid:"Required;ErrorCode(30011)"`                      //确认密码
	Code            string `json:"code" valid:"Required;ErrorCode(30168)"`                                  //验证码
	UserName        string `json:"user_name" valid:"Required;ErrorCode(20011)"`                             //真实姓名
	BankId          string `json:"bank_id" valid:"Required;ErrorCode(30062)"`                               //开户银行
	BackAccount     string `json:"back_account" valid:"Required;ErrorCode(30060)"`                          //银行卡号
	Province        string `json:"province" valid:"Required;ErrorCode(30169)"`                              //银行省份
	Zone            string `json:"zone" valid:"Required;ErrorCode(30170)"`                                  //银行区域

}
