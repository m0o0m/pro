package input

//站点轮播图片管理
type FlashUpdate struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(30041)"`                       //轮播id
	SiteId      string `json:"siteId"`                                                   //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	ImgTitle    string `json:"imgTitle" valid:"Required;MaxSize(50);ErrorCode(90801)"`   //标题
	ImgUrl      string `json:"imgUrl" valid:"Required;MaxSize(200);ErrorCode(90802)"`    //图片路径
	ImgLink     string `json:"imgLink" valid:"Required;MaxSize(200);ErrorCode(90803)"`   //链接
	Ftype       int8   `json:"ftype" valid:"Range(1,2);ErrorCode(90806)"`                //类型 1-PC端 2-WAP端
}
