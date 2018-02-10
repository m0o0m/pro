package input

//添加会员额度转换
type MemberBalanceConversionAdd struct {
	SiteId      string  `json:"-"`                                        //操作站点id
	SiteIndexId string  `json:"-"`                                        //站点前台id
	Account     string  `json:"-"`                                        //会员账号
	Margin      float64 `json:"-"`                                        //手续费
	Money       float64 `json:"money"  valid:"Required;ErrorCode(30150)"` //金额
	FromType    int64   `json:"fromType"`                                 //转入类型,0为系统金额
	ForType     int64   `json:"forType"`                                  //转出类型,0为系统金额
	DoUserId    int64   `json:"doUserId"`                                 //操作人
	DoUserType  int8    `json:"doUserType"`                               //操作人类型1平台管理员2会员
	Media       string  `json:"media" `                                   //设备:pc wap app
	Remark      string  `json:"remark"`                                   //备注
	Platform    string  `json:"platform"`                                 //平台名称
}

//添加会员额度转换
type BalanceConvert struct {
	PlatformId int64  `json:"platformId" valid:"Required;Min(1);ErrorCode(60226)"`   //平台id
	Platform   string `json:"platform" valid:"Required;MinSize(1);ErrorCode(60227)"` //平台名称
	TType      string `json:"tType" `                                                //转换方式 in or out
	Money      int64  `json:"money" valid:"Required;Min(10);ErrorCode(60228)"`       //转换金额,至少10块,整数
	Media      string `json:"media" valid:"Required;MinSize(2);ErrorCode(60229)"`    //设备:pc wap app
	Remark     string `json:"remark" `                                               //备注
}

type OutType struct {
	SiteId      string `query:"site_id"`                                           //操作站点id
	SiteIndexId string `query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Account     string `query:"account" valid:"Required;ErrorCode(30124)"`         //会员账号
	ForType     int64  `query:"for_type"`                                          //转出类型,0为系统金额
}

//额度转换列表
type MemberBalanceConversionList struct {
	SiteId      string `query:"siteId"`                                          //操作站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	StartTime   string `query:"startTime"`                                       //开始时间
	EndTime     string `query:"endTime"`                                         //结束时间
	Account     string `query:"account"`                                         //账号
}

//会员余额统计查询
type MemberClassifyBalance struct {
	SiteId      string `query:"site_id" `      //操作站点id
	SiteIndexId string `query:"site_index_id"` //站点前台id
	Account     string `query:"account"`       //会员账号
	AgencyName  string `query:"agency_name"`   //代理账号
	AgencyId    int64  `query:"agency_id"`     //代理id
	ProductId   int64  `query:"product_id"`    //商品id
	DataType    int    `query:"data_type"`     //查询的数据显示类型  1商品， 2代理， 3会员
}

//wap 额度转换-单个平台余额刷新
type PlatformBalanceRefresh struct {
	PlatformId int64 `json:"platformId"` //平台id
}

//wap 额度转换
type WapMemberBalanceConversion struct {
	Money       float64 `json:"money"  valid:"Required;ErrorCode(30150)"` //金额
	FromType    int64   `json:"fromType"`                                 //类型,0为系统金额
	ForType     int64   `json:"forType"`                                  //类型,0为系统金额
	SiteId      string  `json:"siteId"`                                   //站点id
	SiteIndexId string  `json:"siteIndexId"`                              //站点前台id
	MemberId    int64   `json:"memberId"`                                 //会员id
	Fee         float64 `json:"fee"`                                      //手续费
}
