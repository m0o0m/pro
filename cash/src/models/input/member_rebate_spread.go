package input

//会员推广设定
type SpreadSet struct {
	SiteId       string  `json:"siteId" `                                           //站点id
	SiteIndexId  string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`   //前台站点ID
	IsOpen       int8    `json:"isOpen"  valid:"Range(1, 2);ErrorCode(10000)"`      //是否开启会员推广
	IsIp         int8    `json:"isIp" valid:"Range(1, 2);ErrorCode(10000)"`         //是否过滤ip
	IsMateAgency int8    `json:"isMateAgency" valid:"Range(1, 2);ErrorCode(10000)"` //是否匹配推广会员代理
	IsCode       int8    `json:"isCode" valid:"Range(1, 2);ErrorCode(10000)"`       //返佣会员是否需要打码
	RankingNum   float64 `json:"rankingNum"`                                        //排行榜人数系数
	RankingMoney float64 `json:"rankingMoney" valid:"Required"`                     //排行榜金额系数
}

//会员推广修改
type SpreadEdit struct {
	SiteId       string  `json:"siteId"`                                            //站点id
	SiteIndexId  string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`   //前台站点ID
	IsOpen       int8    `json:"isOpen"  valid:"Range(1, 2);ErrorCode(10000)"`      //是否开启会员推广
	IsIp         int8    `json:"isIp" valid:"Range(1, 2);ErrorCode(10000)"`         //是否过滤ip
	IsMateAgency int8    `json:"isMateAgency" valid:"Range(1, 2);ErrorCode(10000)"` //是否匹配推广会员代理
	IsCode       int8    `json:"isCode" valid:"Range(1, 2);ErrorCode(10000)"`       //返佣会员是否需要打码
	RankingNum   float64 `json:"rankingNum"`                                        //排行榜人数系数
	RankingMoney float64 `json:"rankingMoney" valid:"Required"`                     //排行榜金额系数
}

//会员推广查询
type SpreadInfo struct {
	SiteId      string `query:"site_id"`       //站点id
	SiteIndexId string `query:"site_index_id"` //前台站点ID
	Account     string `query:"account"`       //账号
	RegisterIp  string `query:"register_ip"`   //注册ip
	SpreadId    string `query:"spread_id"`     //推广id
}
