package input

//视讯下载链接表
type SiteDown struct {
	Id int64 `query:"id" valid:"Required;ErrorCode(30041)"` //主键id
}

//视讯下载链接表添加
type SiteDownAdd struct {
	IosUrl     string `json:"ios_url"`                                                          //ios下载地址
	AndroidUrl string `json:"android_url"`                                                      //安卓下载地址
	Platform   string `json:"platform" valid:"Required;MinSize(2);MaxSize(4);ErrorCode(50113)"` //平台
	Vers       string `json:"vers"`                                                             //版本号
}

//视讯下载链接表修改
type SiteDownEdit struct {
	Id         int64  `json:"id" valid:"Required;ErrorCode(30041)"`                             //主键id
	IosUrl     string `json:"ios_url"`                                                          //ios下载地址
	AndroidUrl string `json:"android_url"`                                                      //安卓下载地址
	Platform   string `json:"platform" valid:"Required;MinSize(2);MaxSize(4);ErrorCode(50113)"` //平台
	Vers       string `json:"vers"`                                                             //版本号
}

//视讯下载链接表状态修改
type SiteDownState struct {
	Id    int64 `json:"id" valid:"Required;ErrorCode(30041)"`               //主键id
	State int64 `json:"state" valid:"Required;Range(1,2);ErrorCode(30050)"` //状态
}
