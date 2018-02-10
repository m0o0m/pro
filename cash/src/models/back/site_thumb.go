package back

//自助优惠申请配置
type SiteThumb struct {
	Id          int64  `json:"id"`          //主键id
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台id
	FilePath    string `json:"filePath"`    //文件路径
	FileName    string `json:"fileName"`    //文件名称
	FileType    string `json:"fileType"`    //文件类型
	FileMd5     string `json:"fileMd5"`     //文件加密
	AddTime     int64  `json:"addTime"`     //申请时间
	DeleteTime  int64  `json:"deleteTime"`  //更新时间
	State       int8   `json:"status"`      //1可用 2禁用
}
