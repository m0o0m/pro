package input

//异常会员查询
type AbnormalMemberList struct {
	//SiteId string `query:"siteId"` //站点id
	Type int8   `query:"type"` //1已处理 2未处理
	Key  string `query:"key"`  //关键字
}

//异常会员处理
type AbnormalMemberSet struct {
	AbnormalMembers []AbnormalMemberFeild `json:"abnormal_members"` //异常会员组
}

//异常会员处理唯一字段
type AbnormalMemberFeild struct {
	SiteId  string `json:"site_id"` //站点id
	Account string `json:"account"` //会员账号
}
