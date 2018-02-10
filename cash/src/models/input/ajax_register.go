package input

type AjaxRegister struct {
	Ajax       string `query:"ajax"`     //ajax类型
	Account    string `query:"account"`  //会员账号
	RealName   string `query:"realName"` //
	Code       string `query:"code"`     //
	AgencyUser string `query:"user"`     //
}
type AjaxIsRepeat struct {
	Id      string `query:"id"`      //
	Content string `query:"content"` //
}
