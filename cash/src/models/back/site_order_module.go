package back

//获取电子管理列表、体育管理列表、彩票管理列表
type OrderModuleList struct {
	IndexNum int64  `json:"indexNum"` //序号
	Product  string `json:"product"`  //商品
}

//视讯类型下拉框
type TypeOrderModuleList struct {
	Id   int64  `json:"id"`   //视讯模板id
	Name string `json:"name"` //类型名字
	Aid  int8   `json:"aid"`  //中间关联字段
}

//视讯风格下拉框
type StyleOrderModuleList struct {
	Id   int64  `json:"id"`   //视讯模板id
	Name string `json:"name"` //类型名字
}

//网站基本信息-查询pc&wap
type GetSiteInfoPcAndWap struct {
	Pc  GetSiteInfoPc  `json:"pc"`  //pc信息
	Wap GetSiteInfoWap `json:"wap"` //wap信息
}

//网站基本信息-查询pc
type GetSiteInfoPc struct {
	SiteId      string `xorm:"id" json:"siteId"`
	SiteIndexId string `xorm:"index_id" json:"siteIndexId"`
	Remark      string `xorm:"remark" json:"remark"`
	Qq          string `xorm:"qq" json:"qq"`
	Wechat      string `xorm:"wechat" json:"wechat"`
	Phone       string `xorm:"phone" json:"phone"`
	Email       string `xorm:"email" json:"email"`
	UrlLink     string `xorm:"url_link" json:"urlLink"` //客服连接地址
	PcDomain    string `xorm:"pc_domain" json:"pcDomain"`
	SiteName    string `xorm:"site_name" json:"siteName"`
}

//网站基本信息-查询wap
type GetSiteInfoWap struct {
	SiteId        string `xorm:"id" json:"siteId"`                     //站点id
	SiteIndexId   string `xorm:"index_id" json:"siteIndexId"`          //站点前台id
	AppUrl        string `xorm:"app_url" json:"appUrl"`                //app下载地址
	WapColor      string `xorm:"wap_color" json:"wapColor"`            //wap头部颜色
	WapBottom     string `xorm:"wap_bottom" json:"wapBottom"`          //wap端底部文案
	AutoLinkName  string `xorm:"auto_link_name" json:"autoLinkName"`   //自定义链接名称
	AutoLinkUrl   string `xorm:"auto_link_url" json:"autoLinkUrl"`     //自定义链接url
	WapQuick      int8   `xorm:"wap_quick" json:"wapQuick"`            //1默认不开启，2为开启
	IsDownload    int8   `xorm:"is_download" json:"isDownload"`        //是否允许下载
	WebsiteAppUrl string `xorm:"website_app_url" json:"websiteAppUrl"` //官网app下载地址
	WapDomain     string `xorm:"wap_domain" json:"wapDomain"`
	SiteName      string `xorm:"site_name" json:"siteName"`
}

//template 下拉框
type OrderModule struct {
	Module    ProductMapping
	SubModule []*ProductMappingList
}

type ProductMapping struct {
	Name  string
	VType string
}
type ProductMappingList struct {
	Name       string
	VType      string
	PlatformId int64
}
type BcolorList struct {
	Id           int64  `xorm:"id" json:"id"`                     //
	SiteId       string `xorm:"site_id" json:"siteId"`            //
	SiteIndexId  string `xorm:"site_index_id" json:"siteIndexId"` //
	Bcolor       string `xorm:"bcolor" json:"bcolor"`             //电子内页主题色
	TitleBcolor  string `xorm:"-" json:"titleBcolor"`             //导航背景颜色
	TitleColor   string `xorm:"-" json:"titleColor"`              //导航选中颜色
	ButtonBcolor string `xorm:"-" json:"buttonBcolor"`            //按钮背景颜色
	ButtonColor  string `xorm:"-" json:"buttonColor"`             //按钮选中颜色
	BborderColor string `xorm:"-" json:"bborderColor"`            //按钮边框颜色
	PopBcolor    string `xorm:"-" json:"popBcolor"`               //背景颜色
}
