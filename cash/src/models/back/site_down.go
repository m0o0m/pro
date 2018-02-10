package back

//视讯下载链接表
type SiteDown struct {
	Id         int64  `json:"id"`          //主键id
	IosUrl     string `json:"ios_url"`     //ios下载地址
	AndroidUrl string `json:"android_url"` //安卓下载地址
	State      int64  `json:"state"`       //状态 1启用 2停用
	Platform   string `json:"platform"`    //平台
	Vers       string `json:"vers"`        //版本号
	PcUrl      string `json:"pc_url"`      //pc下载地址
}
