package input

//站点浮动管理
type FloatUpdate struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(90850);"`                       //id
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Ftype       int64  `json:"ftype" valid:"Min(1);Max(2);ErrorCode(90851);"`             //浮动类型 1 左 2右
	Url         string `json:"url" valid:"MaxSize(200);ErrorCode(90852)"`                 //链接
	UrlInter    int64  `json:"urlInter" valid:"Range(0,99);ErrorCode(90853)"`             //内链
	IsBlank     int64  `json:"isBlank" valid:"Max(2);ErrorCode(90854)"`                   //新开窗口 1-是 2-否
	IsSlide     int64  `json:"isSlide" valid:"Max(2);ErrorCode(90855)"`                   //滑动效果 1-是 2-否
	Sort        int64  `json:"sort" valid:"Max(99);ErrorCode(90858)"`                     //排序
}

//站点浮动图片地址修改
type FloatUpdatePicture struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(90850);"`                       //id
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Ftype       int64  `json:"ftype" valid:"Min(1);Max(2);ErrorCode(90851);"`             //浮动类型 1 左 2右
	CType       int8   `json:"cType" valid:"Range(1,2);ErrorCode(50114)"`                 //修改地址类型（1鼠标滑出2鼠标滑入）
	Address     string `json:"address" valid:"MaxSize(200);ErrorCode(50158)"`             //图片地址
	Url         string `json:"url" valid:"MaxSize(200);ErrorCode(90852)"`                 //链接
}

type FloatList struct {
	SiteId      string `query:"siteId"`                                                    //站点id
	SiteIndexId string `query:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Ftype       int64  `query:"ftype" valid:"Range(1,2);ErrorCode(90851);"`                //浮动类型 1 左 2右
}

//浮动图片禁用启用
type FloatListStatus struct {
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Ftype       int64  `json:"ftype" valid:"Range(1,2);ErrorCode(90851);"`                //浮动类型 1 左 2右
	Status      int8   `json:"status" valid:"Range(1,2);ErrorCode(50085)"`                //当前状态
}

//查看站点图片状态
type GetFloatListStatus struct {
	SiteId      string `query:"siteId"`                                                    //站点id
	SiteIndexId string `query:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
}

//删除站点浮动图片
type DeleteFloatPicture struct {
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(30041)"`                        //id
}

type FloatAdd struct {
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Ftype       int64  `json:"ftype" valid:"Min(1);Max(2);ErrorCode();"`                  //浮动类型 1 左 2右
	Url         string `json:"url" valid:"MaxSize(200);ErrorCode(90852)"`                 //链接
	UrlInter    int64  `json:"urlInter" valid:"Max(99);ErrorCode(90853)"`                 //内链
	IsBlank     int64  `json:"isBlank" valid:"Max(2);ErrorCode(90854)"`                   //新开窗口 1-是 2-否
	IsSlide     int64  `json:"isSlide" valid:"Max(2);ErrorCode(90855)"`                   //滑动效果 1-是 2-否
	Sort        int64  `json:"sort" valid:"Max(99);ErrorCode(90858)"`                     //排序
}
