package back

//站点列表/单个站点/返回的struct
type Site struct {
	SiteId      string  `json:"site_id"`
	SiteIndexId string  `json:"site_index_id"`
	SiteName    string  `json:"site_name"`  //站点名称
	ComboId     int64   `json:"combo_id"`   //套餐id
	COmboName   string  `json:"combo_name"` //套餐名称
	DomUp       int     `json:"dom_up"`     //域名上限
	UpCharge    float64 `json:"up_charge"`  //超过上限每一个的收费金额 	//是否默认站点
	Qq          string  `json:"qq"`         //qq
	Wechat      string  `json:"wechat"`     //微信
	Phone       string  `json:"phone"`      //手机
	Email       string  `json:"email"`      //邮箱
	Remark      string  `json:"remark"`     //备注
}

//站点列表 site_id列表
type InSiteList struct {
	SiteId   string `xorm:"id" json:"siteId"`
	SiteName string `xorm:"site_name" json:"siteName"` //站点名称
}

//获取某个开户人下面的所有的站点的返回struct
type OpenUserAllSite struct {
	SiteId      string  `xorm:"id" json:"site_id"`             //主键id
	SiteIndexId string  `xorm:"index_id" json:"site_index_id"` //开户人的site头
	SiteName    string  `xorm:"site_name" json:"site_name"`    //站点名称
	ComboId     int64   `xorm:"combo_id" json:"combo_id"`      //套餐id
	DomUp       int     `xorm:"domain_up" json:"dom_up"`       //域名上限
	UpCharge    float64 `xorm:"up_cose" json:"up_cose"`        //域名超过收费
	ExistDomain int     `xorm:"num" json:"exist_domain"`       //已经有的域名个数
	ExtraCharge float64 `json:"extra_charges"`                 //额外费用
	ComboName   string  `xorm:"combo_name" json:"combo_name"`  //套餐名称
	QQ          string  `xorm:"qq" json:"qq"`                  //qq
	Wechat      string  `xorm:"wechat" json:"wechat"`          //微信
	Phone       string  `xorm:"phone" json:"phone"`            //电话
	Email       string  `xorm:"email" json:"email"`            //邮箱
	Status      int8    `xorm:"status" json:"status"`          //状态
	Remark      string  `xorm:"remark" json:"remark"`          //备注
}

//获取某个开户人下面所有的站点index_id
type IndexBackStruct struct {
	SiteId      string `xorm:"id" json:"siteId"`
	SiteIndexId string `xorm:"index_id" json:"siteIndexId"`
	SiteName    string `xorm:"site_name" json:"siteName"`
	IsDefault   int8   `xorm:"is_default" json:"isDefault"`
}

//SiteUpList 返回站点域名配置列表
type SiteUpList struct {
	Id          int64  `xorm:"id" json:"id"`
	IsMain      int    `xorm:"is_default" json:"is_default"`
	IsUsed      int    `xorm:"is_used" json:"is_used"`
	SiteId      string `xorm:"site_id" json:"site_id"`
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"`
	Domain      string `xorm:"domain" json:"domain"`
	CreatTime   int64  `xorm:"create_time" json:"create_time"`
}

//站点详情返回
type InfoBack struct {
	Id          int64             `xorm:"id" json:"id"`
	Domain      string            `xorm:"domain" json:"domain"`
	SiteId      string            `xorm:"site_id" json:"site_id"`
	SiteIndexId string            `xorm:"site_index_id" json:"site_index_id"`
	FileName    map[string]string `xorm:"file_name" json:"file_name"`
}

//站点下面的不同的域名返回集合
type DomainBack struct {
	//Domain []string `json:"wap_domain"`
	Domain []string `json:"domain"`
}

//平台管理员站点下拉框的返回
type BackSiteDrop struct {
	SiteId      string `xorm:"id" json:"siteId"`
	SiteIndexId string `xorm:"index_id" json:"siteIndexId"` //站点id
	SiteName    string `xorm:"site_name" json:"siteName"`   //站点名称
}

//根据site_id获取站点前台信息
type SiteIndexList struct {
	SiteId      string  `xorm:"id" json:"id"`                  //站点id
	SiteIndexId string  `xorm:"index_id" json:"site_index_id"` //前台
	SiteName    string  `xorm:"site_name" json:"site_name"`    //站点名称
	ComboName   string  `xorm:"combo_name" json:"combo_name"`  //套餐id
	ProductId   int64   `xorm:"product_id" json:"product_id"`  //套餐id
	Proportion  float64 `xorm:"proportion" json:"proportion"`  //占成比
}

//站点会员列表
type SiteMemberInfo struct {
	Id                 int64                    `xorm:"id" json:"id"`
	SiteId             string                   `xorm:"site_id" json:"siteId"`
	SiteIndexId        string                   `xorm:"site_index_id" json:"siteIndexId"`
	Account            string                   `xorm:"account" json:"account"`
	Realname           string                   `xorm:"realname" json:"realname"`
	Balance            float64                  `xorm:"balance" json:"balance"`
	Status             int                      `xorm:"status" json:"status"`
	LoginIp            string                   `xorm:"login_ip" json:"loginIp"`
	CreateTime         string                   `xorm:"create_time" json:"createTime"`
	Device             uint                     `xorm:"device" json:"device"`
	PcStatus           int8                     `xorm:"pc_status" json:"-"`
	WapStatus          int8                     `xorm:"wap_status" json:"-"`
	IosStatus          int8                     `xorm:"ios_status" json:"-"`
	AndroidStatus      int8                     `xorm:"android_status" json:"-"`
	MemberVideoBalance []MemberVideoBalanceBack `xorm:"-" json:"memberVideBalance"`
}

//会员视讯余额
type MemberVideoBalanceBack struct {
	MemberId   int64   `xorm:"member_id" json:"memberId"`     //会员id
	Platform   string  `xorm:"platform" json:"platform"`      //视讯平台名称
	PlatformId int64   `xorm:"platform_id" json:"platformId"` //视讯平台id
	Balance    float64 `xorm:" balance" json:"balance"`       //额度
}

//前台文案列表
type CopyList struct {
	Id          int64  `json:"id" xorm:"id"`                     //文案id
	TopId       int64  `json:"topId" xorm:"top_id"`              //上级栏目
	SiteId      string `json:"siteId" xorm:"site_id"`            //站点id
	SiteIndexId string `json:"siteIndexId" xorm:"site_index_id"` //站点前台id
	Title       string `json:"title" xorm:"title"`               //标题
	TitleColor  string `json:"titleColor" xorm:"title_color"`    //标题颜色
	Content     string `json:"content" xorm:"content"`           //内容
	Url         string `json:"url" xorm:"url"`                   //链接地址
	Img         string `json:"img" xorm:"img"`                   //图片路径
	State       int8   `json:"state" xorm:"state"`               //状态 1-启用 2-关闭
	Sort        int64  `json:"sort" xorm:"sort"`                 //排序
	From        int8   `json:"from" xorm:"from"`                 //'1-PC 2-WAP
	Itype       int64  `json:"itype" xorm:"itype"`               //类型代码
	TypeName    string `json:"typeName" xorm:"type_name"`        //类型名称
	AddTime     int64  `json:"addTime" xorm:"add_time"`          //操作时间
}

type CopyListInfo struct {
	Id      int64  `json:"id" xorm:"id"`           //文案id
	Content string `json:"content" xorm:"content"` //内容
}

//轮播图片（pc、wap）
type FlashAllList struct {
	PcFlashList  []FlashList `json:"pcFlashList"`  //PC端轮播图片
	WapFlashList []FlashList `json:"wapFlashList"` //wap端轮播图片
}

//轮播查询列表
type FlashList struct {
	Id          int64  `json:"id" xorm:"id"`                     //轮播id
	SiteId      string `json:"siteId" xorm:"site_id"`            //站点id
	SiteIndexId string `json:"siteIndexId" xorm:"site_index_id"` //站点前台id
	ImgTitle    string `json:"imgTitle" xorm:"img_title"`        //标题
	ImgUrl      string `json:"imgUrl" xorm:"img_url"`            //图片路径
	ImgLink     string `json:"imgLink" xorm:"img_link"`          //链接地址
	State       int8   `json:"state" xorm:"state"`               //状态 1-启用 2-关闭
	Sort        int64  `json:"sort" xorm:"sort"`                 //排序
	Ftype       int8   `json:"ftype" xorm:"ftype"`               //类型 1-PC端 2-WAP端
	IsLink      int    `json:"isLink" xorm:"is_link"`            //新开页面 1-是 2-否
}

//站点logo图片管理列表
type LogoList struct {
	Id          int64  `json:"id" xorm:"id"`                     //文案id
	SiteId      string `json:"siteId" xorm:"site_id"`            //站点id
	SiteIndexId string `json:"siteIndexId" xorm:"site_index_id"` //站点前台id
	Title       string `json:"title" xorm:"title"`               //logo名称
	LogoUrl     string `json:"logoUrl" xorm:"logo_url"`          //logo地址
	Type        int8   `json:"type" xorm:"type"`                 //文案类型
	State       int8   `json:"state" xorm:"state"`               //状态1启用 2停用
	Form        int8   `json:"form" xorm:"form"`                 //1PC 2WAP
}

//多站点-列表
type SiteMoreList struct {
	SiteId      string  `xorm:"site_id" json:"siteId"`        //站点id
	Id          int64   `xorm:"id" json:"id"`                 //域名id
	SiteIndexId string  `json:"siteIndexId"`                  //站点前台id
	SiteName    string  `json:"siteName"`                     //站点名称
	Domain      string  `json:"Domain"`                       //pc域名
	IsDefaultD  int8    `xorm:"is_default" json:"isDefaultD"` //是否默认站点
	IsDefault   int8    `json:"isDefault"`                    //是否主域名 1是  2否
	Status      int8    `json:"status"`                       //状态1正常2禁用3.暂停4.维护
	CreateTime  int64   `json:"createTime"`                   //创建时间
	IsDownApp   int8    `json:"isDownApp"`                    //是否可以下载app(1.可以2.不可以)
	UpCose      float64 `json:"upCose"`                       //收费金额
	DomainUp    int     `json:"domainUp"`                     //域名上限
}

//报表负数-负数+商品
type ReportNegativeAndProducts struct {
	Id          int64   `json:"id"`
	SiteId      string  `json:"site_id"`       //站点id
	SiteIndexId string  `json:"site_index_id"` //站点前台id
	Years       string  `json:"years"`         //年月
	ProductId   int64   `json:"product_id"`    //商品分类id
	ProductName string  `json:"product_name"`  //商品名
	ReportWin   float64 `json:"report_win"`    //报表盈利数字
}

//报表负数-列表
type ReportNegativeList struct {
	SiteId      string                  `json:"site_id"`       //站点id
	SiteIndexId string                  `json:"site_index_id"` //站点前台id
	Years       string                  `json:"years"`         //年月
	Products    []ReportNegativeProduct `json:"products"`      //商品列表
}

//报表负数-商品
type ReportNegativeProduct struct {
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品名
	ReportWin   float64 `json:"report_win"`   //报表盈利数字
}

//皮肤数据
type Theme struct {
	SiteId      string `xorm:"site_id"`       // 站点id
	SiteIndexId string `xorm:"site_index_id"` //前台id
	ThemeName   string `xorm:"theme_name"`    //皮肤名称
}

type GetDomainBack struct {
	Domain string `json:"domain"`
}

//站点出款设置
type SitePaySet struct {
	IsFree    int8    `json:"is_free" xorm:"is_free"`       // 出款是否免手续费
	FreeNum   int     `json:"free_num" xorm:"free_num"`     // 出款免费次数
	OutCharge float64 `json:"out_charge" xorm:"out_charge"` // 出款手续费金额
}

//站点详情
type SiteInfoBySiteAndSiteIndexBack struct {
	Id        string  `xorm:"id" json:"id"`                 //主键id
	IndexId   string  `xorm:"index_id" json:"indexId"`      //前台id
	SiteName  string  `xorm:"site_name" json:"siteName"`    //站点名称
	ComboId   int64   `xorm:"combo_id" json:"comboId"`      //套餐id
	DomainUp  int     `xorm:"domain_up" json:"domainUp"`    //域名上限
	UpCose    float64 `xorm:"up_cose" json:"upCose"`        //超过上线收费金额
	IsDownApp int8    `xorm:"is_down_app" json:"isDownApp"` //是否可以下载app
	Domain    string  `xorm:"domain" json:"domain"`         //域名  50
}
