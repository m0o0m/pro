package input

//视讯账号查询
type SiteVideoMemberSearch struct {
	SiteId   string `query:"siteId"`   //站点ID
	Platform string `query:"platform"` //视讯类型
	Type     int8   `query:"type"`     //1平台用户名2视讯用户名
	Account  string `query:"account"`  //账号
}
