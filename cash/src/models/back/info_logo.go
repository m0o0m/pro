package back

//获取站点logo图片列表
type InfoLogoList struct {
	Id          int64  `json:"id"`          //
	SiteId      string `json:"siteId"`      //站点ID
	SiteIndexId string `json:"siteIndexId"` //前台ID
	Title       string `json:"title"`       //logo名称
	LogoUrl     string `json:"logoUrl"`     //logo地址
	Type        int8   `json:"type"`        //文案类型
	State       int8   `json:"state"`       //状态1启用 2停用
	Form        int64  `json:"form"`        //1pc     2wap
}
type GetLogoInfo struct {
	Id    int64  `json:"id"`    //
	Title string `json:"title"` //logo名称
}

//站点浮动管理
type FloatList struct {
	Id          int64  `xorm:"id"  json:"id"`                     //轮播id
	SiteId      string `xorm:"site_id"  json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id"  json:"siteIndexId"` //站点前台id
	ImgA        string `xorm:"img_a"  json:"imgA"`                //常规显示图片
	ImgB        string `xorm:"img_b"  json:"imgB"`                //鼠标覆盖事件显示图片
	Url         string `xorm:"url"  json:"url"`                   //链接
	UrlInter    string `xorm:"url_inter"  json:"urlInter"`        //内链
	IsBlank     int64  `xorm:"is_blank"  json:"isBlank"`          //新开窗口 1-是 2-否
	IsSlide     int64  `xorm:"is_slide"  json:"isSlide"`          //滑动效果 1-是 2-否
	IsClose     int64  `xorm:"is_close"  json:"isClose"`          //关闭按钮
	State       int64  `xorm:"state"  json:"state"`               //状态 1-启用 2-关闭
	Sort        int64  `xorm:"sort"  json:"sort"`                 //排序
	Ftype       int64  `xorm:"ftype"  json:"ftype"`               //类型 1-左 2-右
}

//站点浮动图片状态
type FloatAllStatus struct {
	LeftFloat  *FloatStatus `json:"leftFloat"`  //左浮动图片状态
	RightFloat *FloatStatus `json:"rightFloat"` //右浮动图片状态
}

//站点浮动图片状态
type FloatStatus struct {
	State int64 `xorm:"state"  json:"state"` //状态 1-启用 2-关闭
	Ftype int64 `xorm:"ftype"  json:"ftype"` //类型 1-左 2-右
}

//logo
type HomePageLogoBack struct {
	Id      int64  `xorm:"id" json:"id"`            //文案id
	Title   string `xorm:"title" json:"title"`      //logo名称
	LogoUrl string `xorm:"logo_url" json:"logoUrl"` //logo地址
}

//轮播图
type HomePagePictureBack struct {
	Id      int64  `xorm:"id" json:"id"`            //轮播id
	ImgUrl  string `xorm:"img_url" json:"imgUrl"`   //图片路径
	ImgLink string `xorm:"img_link" json:"imgLink"` //链接地址
}

//公告
type HomePageNoticeBack struct {
	ID            int64  `xorm:"id" json:"id"`
	NoticeTitle   string `xorm:"notice_title" json:"noticeTitle"`       //公告标题
	NoticeContent string `xorm:"notice_content" json:"noticeContent"`   //公告内容
	NoticeDate    int64  `xorm:"notice_date created" json:"noticeDate"` //公告时间
}

//首页总
type HomePageAllInfo struct {
	HomePageLogoBack           HomePageLogoBack      `json:"homePageLogoBack"`           //logo
	HomePagePictureBack        []HomePagePictureBack `json:"homePagePictureBack"`        //轮播图
	HomePageNoticeBack         []HomePageNoticeBack  `json:"homePageNoticeBack"`         //公告
	HomePageProductAndTypeBack []OrderModuleBySite   `json:"homePageProductAndTypeBack"` //商品分类、商品
}

//template 下拉框
type OrderModuleBySite struct {
	Module    ProductMapping
	SubModule []*ProductMapping
}

//首页商品分类总
type HomePageProductAndTypeBack struct {
	Id                  int64                 `json:"id"`    //商品分类id
	Title               string                `json:"title"` //商品分类名称
	HomePageProductBack []HomePageProductBack `json:"homePageProductBack"`
}

//商品
type HomePageProductBack struct {
	Id          int64  `xorm:"id" json:"id"`                    //主键id
	ProductName string `xorm:"product_name" json:"productName"` //商品名 30
	TypeId      int64  `xorm:"type_id" json:"typeId"`           //商品类型id
	PlatformId  int64  `xorm:"platform_id" json:"platformId"`   //平台id
	VType       string `xorm:"v_type" json:"vType"`             //游戏类型
	Api         string `xorm:"api" json:"api"`                  //商品域名
}
