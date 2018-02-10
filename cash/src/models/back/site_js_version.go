package back

//视讯下载链接表
type SiteJsVersion struct {
	Id    int64  `json:"id"`    //主键id
	State int64  `json:"state"` //状态  1启用 2停用
	Type  string `json:"type"`  //类别  1pc  2wap
	Vers  string `json:"vers"`  //版本号
}
