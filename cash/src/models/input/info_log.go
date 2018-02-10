//站点logo图片管理
package input

//站点logo列表
type LogoInfoList struct {
	SiteId      string //站点id
	SiteIndexId string `query:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
}

//站点logo修改
type UpdateLogoInfo struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(90905);"`                       //id
	Title       string `json:"title" valid:"Required;MaxSize(80);ErrorCode(90901);"`      //logo名称
	State       int8   `json:"state" valid:"Range(1,2);ErrorCode(90903);"`                //状态
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
}

//站点logo修改路径
type UpdateLogoInfoWay struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(90905);"`                       //id
	LogoUrl     string `json:"logoUrl" valid:"Required;MaxSize(200);ErrorCode(90902)"`    //路径
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
}

//查询单条信息
type GetLogoInfo struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(90905);"`                          //id
	Title       string `json:"title" valid:"Min(1);ErrorCode(90901);"`                       //logo名称
	SiteId      string `query:"site_id" valid:"Required;MaxSize(4);ErrorCode(60105)"`        //站点id
	SiteIndexId string `query:"site_index_id"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
}

//新增站点logo
type PostLogoInfo struct {
	Title       string `json:"title" valid:"Required;MaxSize(20);ErrorCode(90901);"`      //logo名称
	State       int8   `json:"state" valid:"Min(1);ErrorCode(90903);"`                    //状态
	SiteId      string `json:"siteId" `                                                   //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	LogoUrl     string `json:"logoUrl" valid:"Required;MaxSize(200);ErrorCode(90902)"`    //
	Type        int8   `json:"type" valid:"Min(1);ErrorCode(90906);"`                     //
	Form        int64  `json:"form" valid:"Min(1);Max(2);ErrorCode(90907);"`              //
}
