package input

//邮箱添加或修改
type EmailAddOrChangeIn struct {
	Id          int64  //会员id
	SiteId      string //站点
	SiteIndexId string //站点前台id
	Email       string `json:"email" valid:"Required;MaxSize(50);ErrorCode(20015)"`
}

//出生日期添加或修改
type BirthAddOrChangeIn struct {
	Id          int64  //会员id
	SiteId      string //站点
	SiteIndexId string //站点前台id
	Birth       string `json:"birth" valid:"Required;MaxSize(10);ErrorCode(20012)"`
}

//手机验证码
type PhoneCode struct {
	Phone string `query:"phone" valid:"Required;MaxSize(11);ErrorCode(20014)"` //手机号
}

//手机号绑定
type PhoneBind struct {
	Id          int64  //会员id
	SiteId      string //站点
	SiteIndexId string //站点前台id
	Phone       string `json:"phone" valid:"Required;MaxSize(11);ErrorCode(20014)"` //手机号码
	//PhoneCode   string `json:"phone_code" valid:"Required;MaxSize(4);ErrorCode(30168)"` //验证码
	LocalCode string `json:"localCode" valid:"Required;MaxSize(4);ErrorCode(50120)"` //区号
}

//修改密码
type PasswordMemberChange struct {
	Type             int    `json:"type"  valid:"Required;Min(1);ErrorCode(50114)"`                 //修改的密码类型
	OriginalPassword string `json:"originalPassword" valid:"Required;MaxSize(18);ErrorCode(50148)"` //原始密码
	Password         string `json:"password" valid:"Required;MaxSize(18);ErrorCode(20006)"`         //新密码
	RePassword       string `json:"rePassword" valid:"Required;MaxSize(18);ErrorCode(20007)"`       //重复密码
}
