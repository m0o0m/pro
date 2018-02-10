/*
input包存放所有请求上来的数据的对应的struct
*/

package input

/*
Required: 此参数必须填写；
Min(1): 最小值；
MaxSize(50): 值最大长度；
Match(pattern): 根据正则pattern匹配校验
*/
//AddSite 增加站点的请求数据的struct
type AddSite struct {
	Site            string  `json:"site" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`
	SiteIndex       string  `json:"siteIndex" valid:"Required;MaxSize(4);ErrorCode(10050)"`            //前台id
	SiteName        string  `json:"siteName" valid:"Required;MinSize(1);MaxSize(50);ErrorCode(60104)"` //站点名称
	Domain          string  `json:"domain" valid:"Required;MaxSize(50);ErrorCode(60013)"`              //域名
	BackstageDomain string  `json:"backstageDomain" valid:"Required;MaxSize(50);ErrorCode(60098)"`     //后台域名
	DomUp           int     `json:"domUp" valid:"Required;Min(1);ErrorCode(60106)"`                    //域名上限
	UpCharge        float64 `json:"upCharge" valid:"Required;ErrorCode(60107)"`                        //超过上限每一个的收费金额
	ComboId         int64   `json:"comboId" valid:"Required;Min(1);ErrorCode(60103)"`                  //套餐id
	IsDown          int8    `json:"isDown"`                                                            //是否能下载  1是2否
}

//AddSite 增加站点的请求数据的struct
type AddSiteIn struct {
	Site      string  `json:"site" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`
	SiteIndex string  `json:"siteIndex" valid:"Required;MaxSize(4);ErrorCode(10050)"`            //前台id
	SiteName  string  `json:"siteName" valid:"Required;MinSize(1);MaxSize(50);ErrorCode(60104)"` //站点名称
	Domain    string  `json:"domain" valid:"Required;MaxSize(50);ErrorCode(60013)"`              //域名
	DomUp     int     `json:"domUp" valid:"Required;Min(1);ErrorCode(60106)"`                    //域名上限
	UpCharge  float64 `json:"upCharge" valid:"Required;ErrorCode(60107)"`                        //超过上限每一个的收费金额
	IsDown    int8    `json:"isDown"`                                                            //是否能下载  1是2否
}

//修改站点
type EditSite struct {
	Site      string  `json:"site" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`
	SiteIndex string  `json:"siteIndex" valid:"Required;MaxSize(4);ErrorCode(10050)"`            //前台id
	SiteName  string  `json:"siteName" valid:"Required;MinSize(1);MaxSize(50);ErrorCode(60104)"` //站点名称
	Domain    string  `json:"domain" valid:"Required;MaxSize(50);ErrorCode(60013)"`              //pc域名
	UpCharge  float64 `json:"upCharge" valid:"Required;ErrorCode(60107)"`                        //超过上限每一个的收费金额
	ComboId   int64   `json:"comboId" valid:"Required;Min(1);ErrorCode(60103)"`                  //套餐id
	IsDown    int8    `json:"isDown"`                                                            //是否能下载  1是2否
}

//DelSite 删除站点/取得站点详情请求参数struct
type DelSite struct {
	Site      string `query:"site" json:"site" valid:"Required;MaxSize(4);ErrorCode(60105)"`             //站点id
	SiteIndex string `query:"site_index" json:"site_index" valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
}

//GetSiteDomain 获取该站点各种不同的域名集合
type GetSiteDomain struct {
	Site      string `query:"site" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`
	SiteIndex string `query:"site_index" valid:"Required;MaxSize(4);ErrorCode(10050)"`
}

//GetAllSite  获取某个开户人名下的所有站点请求参数struct
type GetAllSite struct {
	ComboId  int64  `query:"combo_id" valid:"ErrorCode(60103)"`              //套餐id
	Status   int    `query:"status" valid:"Range(0,2);ErrorCode(60108)"`     //站点状态
	SiteName string `query:"site_name" valid:"MaxSize(50);ErrorCode(60009)"` //站点名称
	OperUser int64  `query:"open_user"  valid:"Required;ErrorCode(60110)"`   //开户人id
}

//SiteStatus 更改站点状态请求参数struct
type SiteStatus struct {
	Site      string `json:"site" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`
	SiteIndex string `json:"site_index" valid:"Required;MaxSize(4);ErrorCode(10050)"`
}

//DelDomainSet  删除站点/获取站点域名配置的请求参数struct
type DelDomainSet struct {
	Id int64 `query:"id" valid:"Required;Min(1);ErrorCode(60109)"` //域名配置id
}

//DomainSiteList 获取域名配置列表的请求参数列表
type DomainSiteList struct {
	Site      string `query:"site" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`
	SiteIndex string `query:"site_index" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"`
	Domain    string `query:"domain"` //搜索域名pc或者wap
}

//平台管理员获取站点下拉框的
type SelectSiteDrop struct {
	SiteName string `query:"site_name"`
}

//根据站点查询会员
type SiteMemberInfo struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	Account     string `query:"account"`
	StartTime   string `query:"startTime"`
	EndTime     string `query:"endTime"`
	PageSize    int    `query:"pageSize"`
	Page        int    `query:"page"`
	Ip          string `query:"ip"`
	Status      uint   `query:"status"`
	Device      uint   `query:"device"`
}

//站点视讯额度修改
type SiteVideoBalance struct {
	SiteId         string  `json:"siteId" valid:"Required;MaxSize(4);MinSize(1);ErrorCode(60105)"`      //站点id
	SiteIndexId    string  `json:"siteIndexId" valid:"Required;MaxSize(4);MinSize(1);ErrorCode(10050)"` //站点前台id
	NowBalance     float64 `json:"nowBalance" valid:"Required;ErrorCode(60086)"`                        //当前金额
	Operate        int     `json:"operate" valid:"Required;Range(1,3);ErrorCode(60087)"`                //1.存款2.扣款 3.预借存款(就是赊欠)
	OperateBalance float64 `json:"operateBalance" valid:"Required;ErrorCode(60088)"`                    //操作金额
	Remark         string  `json:"remark"`                                                              //备注
}

//站点上线
type SiteOnline struct {
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);MinSize(1);ErrorCode(60105)"`      //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);MinSize(1);ErrorCode(10050)"` //站点前台id
	OnlineTime  string `json:"onlineTime" valid:"Required;ErrorCode(60089)"`                        //上线时间
}

//一键生成三级代理
type GenerationAgency struct {
	SiteId                     string `json:"siteId" valid:"Required;MaxSize(4);MinSize(1);ErrorCode(60105)"`          //站点id
	SiteIndexId                string `json:"siteIndexId" valid:"Required;MaxSize(4);MinSize(1);ErrorCode(10050)"`     //站点前台id
	DefaultShareholdersName    string `json:"defaultShareholdersName" valid:"Required;MaxSize(30);ErrorCode(60090)"`   //默认股东名称
	DefaultShareholdersAccount string `json:"defaultShareholdersAccount" valid:"Required;MaxSize(30)ErrorCode(60091)"` //默认股东账号
	DefaultShareholdersRemark  string `json:"defaultShareholdersRemark"`                                               //股东备注
	DefaultTotalAgencyName     string `json:"defaultTotalAgencyName" valid:"Required;MaxSize(30)ErrorCode(60092)"`     //默认总代名称
	DefaultTotalAgencyAccount  string `json:"defaultTotalAgencyAccount" valid:"Required;MaxSize(30)ErrorCode(60093)"`  //默认总代账号
	DefaultTotalAgencyRemark   string `json:"defaultTotalAgencyRemark"`                                                //总代理备注
	DefaultAgencyName          string `json:"defaultAgencyName" valid:"Required;MaxSize(30)ErrorCode(60094)"`          //默认代理名称
	DefaultAgencyAccount       string `json:"defaultAgencyAccount" valid:"Required;MaxSize(30);ErrorCode(60095)"`      //默认代理账号
	DefaultAgencyRemark        string `json:"defaultAgencyRemark"`                                                     //代理备注
}

//多站点-列表
type SiteMoreList struct {
	SiteId   string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"` //站点id
	Page     int    `query:"page"`                                                //页码
	PageSize int    `query:"pageSize"`                                            //每页条数
}

//多站点-添加代理
type SiteMoreAdd struct {
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	Account     string `json:"account" valid:"Required;MaxSize(30);ErrorCode(50010)"`    //账号
	Username    string `json:"username" valid:"Required;MaxSize(50);ErrorCode(50110)"`   //用户名
	Remark      string `json:"remark" valid:"MaxSize(255);ErrorCode(20019)"`             //备注
}

//报表负数查询
type ReportNegativeList struct {
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//报表负数添加
type ReportNegativeAdd struct {
	SiteId      string                  `json:"site_id" valid:"Required;MaxSize(4);ErrorCode(60105)"`       //站点id
	SiteIndexId string                  `json:"site_index_id" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	Years       string                  `json:"years"`                                                      //年月
	Products    []ReportNegativeProduct `json:"products"`                                                   //商品列表
}

//报表负数-商品
type ReportNegativeProduct struct {
	ProductId int64   `json:"product_id"` //商品分类id
	ReportWin float64 `json:"report_win"` //报表盈利数字
}

//报表负数修改
type ReportNegativeEdit struct {
	Id          int64                   `json:"id"`
	SiteId      string                  `json:"site_id" valid:"Required;MaxSize(4);ErrorCode(60105)"`       //站点id
	SiteIndexId string                  `json:"site_index_id" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	Years       string                  `json:"years"`                                                      //年月
	State       int8                    `json:"state"`                                                      //状态 1累计 2清零
	Products    []ReportNegativeProduct `json:"products"`                                                   //商品列表
}

//站点管理-一键生成三级代理账号获取
type ProductAccount struct {
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
}
