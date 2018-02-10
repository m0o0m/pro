package input

//获取电子管理列表、体育管理列表、彩票管理列表、网站基本信息-查询、模板管理列表
type OrderModuleList struct {
	SiteId      string `query:"siteId"`                                                   //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//修改电子管理顺序、体育管理顺序、彩票管理列表、模板管理列表
type EditModuleOrder struct {
	SiteId      string `json:"siteId"`                                                   //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	VType       string `json:"vType"`                                                    //商品类型v_type
	IndexNum    int64  `json:"indexNum" valid:"Required;Min(1);ErrorCode(10149)"`        //序号
}

//获取视讯风格下拉框
type StyleOrderModuleList struct {
	Pid int64 `query:"pid"` //父id,即上级类型的aid是本级的pid
}

//视讯管理-风格配置使用
type PostVideoOrderStyleUse struct {
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台
	Style       int64  `json:"style"`       //0默认风格
	Type        int64  `json:"type"`        //模板
}

//视讯管理-风格还原默认
type PutVideoOrderStyleUseUpdate struct {
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台
}

//网站基本信息-添加或修改pc
type PostSiteInfo struct {
	SiteId        string `json:"siteId"`
	SiteIndexId   string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"`
	Remark        string `json:"remark"`
	Qq            string `json:"qq"`
	Wechat        string `json:"wechat"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	UrlLink       string `json:"urlLink"`       //客服连接地址
	AppUrl        string `json:"appUrl"`        //app下载地址
	WapColor      string `json:"wapColor"`      //wap头部颜色
	WapBottom     string `json:"wapBottom"`     //wap端底部文案
	AutoLinkName  string `json:"autoLinkName"`  //自定义链接名称
	AutoLinkUrl   string `json:"autoLinkUrl"`   //自定义链接url
	WapQuick      int8   `json:"wapQuick"`      //1默认不开启，2为开启
	IsDownload    int8   `json:"isDownload"`    //是否允许下载
	WebsiteAppUrl string `json:"websiteAppUrl"` //官网app下载地址
}

//网站基本信息-添加或修改wap
type PostSiteInfoWap struct {
	SiteId        string `json:"site_id"`         //站点id
	SiteIndexId   string `json:"site_index_id"`   //站点前台id
	AppUrl        string `json:"app_url"`         //app下载地址
	WapColor      string `json:"wap_color"`       //wap头部颜色
	WapBottom     string `json:"wap_bottom"`      //wap端底部文案
	AutoLinkName  string `json:"auto_link_name"`  //自定义链接名称
	AutoLinkUrl   string `json:"auto_link_url"`   //自定义链接url
	WapQuick      int8   `json:"wap_quick"`       //1默认不开启，2为开启
	IsDownload    int8   `json:"is_download"`     //是否允许下载
	WebsiteAppUrl string `json:"website_app_url"` //官网app下载地址
}
