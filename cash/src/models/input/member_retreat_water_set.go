package input

//返点优惠设定列表
type RetreatWaterSetList struct {
	SiteId      string `json:"siteId" query:"siteId" `                                              //操作站点id
	SiteIndexId string `json:"siteIndexId" query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//新增返点优惠设定
type AddRetreatWaterSet struct {
	SiteId      string                    //操作站点id
	SiteIndexId string                    `json:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"`    //站点前台id
	ValidMoney  int64                     `json:"valid_money" valid:"Required;Min(1);ErrorCode(10103)"` //有效总投注
	DiscountUp  int64                     `json:"discount_up" valid:"Required;Min(1);ErrorCode(10104)"` //优惠上限
	Params      []ListRetreatWaterProduct `json:"params"`                                               //引用商品列表
}

//获取单个列表详情
type GetOneRetreatWaterDetails struct {
	Id          int64  `json:"id" query:"id"`                   //返点优惠设定id
	SiteId      string `json:"siteId" query:"siteId" `          //操作站点id
	SiteIndexId string `json:"siteIndexId" query:"siteIndexId"` //站点前台id
}

//修改返点优惠设定
type EditRetreatWaterSet struct {
	Id          int64                     `json:"id"`
	SiteId      string                    `json:"site_id" query:"site_id" `                                                //操作站点id
	SiteIndexId string                    `json:"site_index_id" query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	ValidMoney  int64                     `json:"valid_money" valid:"Required;Min(1);ErrorCode(10103)"`                    //有效总投注
	DiscountUp  int64                     `json:"discount_up" valid:"Required;Min(1);ErrorCode(10104)"`                    //优惠上限
	Params      []ListRetreatWaterProduct `json:"params"`                                                                  //引用商品列表
}

//删除返点设定
type DelRetreatWaterSet struct {
	Id          int64  `json:"id"`                                                                      //优惠设定ID
	SiteIndexId string `json:"site_index_id" query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	SiteId      string
}
