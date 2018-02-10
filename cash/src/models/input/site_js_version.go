package input

//视讯下载链接表修改
type SiteJsVersion struct {
	Id   int64  `json:"id" valid:"Required;ErrorCode(30041)"` //主键id
	Vers string `json:"vers"`                                 //版本号
	//State int64  `json:"state" valid:"Required;Range(1,2);ErrorCode(30050)"` //状态
}
