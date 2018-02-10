package back

//查询一条三方对接验证
type GetOneApiClients struct {
	UserId int64  `json:"user_id"` //客户id
	Name   string `json:"name"`    //名称
	Secret string `json:"secret"`  //证书
	SiteId string `json:"site_id"` //证书
}
