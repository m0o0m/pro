package input

//登录表单
type AgencySignLogin struct {
	Account    string `json:"account" valid:"MinSize(4);MaxSize(12);ErrorCode(20000)"`
	Password   string `json:"password" valid:"MinSize(4);MaxSize(12);ErrorCode(20000)"`
	Code       string `json:"code" valid:"MinSize(4);MaxSize(4);ErrorCode(20021)"`
	VerifyCode string `json:"verify_code"`
}

//修改密码表单
type AgencySignPassword struct {
	OldPassword   string `json:"old_password" valid:"MinSize(4);MaxSize(12);ErrorCode(20005)"`
	NewPassword   string `json:"new_password" valid:"MinSize(4);MaxSize(12);ErrorCode(20006)"`
	ReplyPassword string `json:"reply_password" valid:"MinSize(4);MaxSize(12);ErrorCode(20007)"`
}
