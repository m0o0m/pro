package back

//站点口令列表返回
type SitePass struct {
	SiteId  string `json:"site_id"`  //站点id
	Status  int8   `json:"status" `  //状态 1 未开启 2 开启
	PassKey string `json:"pass_key"` //密钥
}
