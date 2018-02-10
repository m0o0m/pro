package back

type VideoMemberBack struct {
	Id          int64  `xorm:"id" json:"id"`
	ProductName string `xorm:"product_name" json:"productName"`
}

//视讯账号查询
type VideoMemberSearch struct {
	Account    string  `json:"account"`    //登录账号
	GUsername  string  `json:"gUsername"`  //视讯账号
	Balance    float64 `json:"balance"`    //视讯余额
	SiteId     string  `json:"siteId"`     //站点id
	IndexId    string  `json:"indexId"`    //站点前台id
	CreateTime int64   `json:"createTime"` //创建时间
	Status     int8    `json:"status"`     //状态 1正常 2停用
}
