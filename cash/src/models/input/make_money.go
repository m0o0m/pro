package input

//ConfirmOutMoney 确认出款请求struct
type ConfirmOutMoney struct {
	SiteId   string
	Id       int64 `json:"id" valid:"Required;Min(1);ErrorCode(60056)"`
	AgencyId int64 `json:"agencyId" valid:"Min(0);ErrorCode(60043)"` //操作者Id
}

//CancleOutMoney 取消出款请求struct
type CancleOutMoney struct {
	SiteId   string
	Id       int64  `json:"id" valid:"Required;Min(1);ErrorCode(60056)"`
	Reason   string `json:"reason" valid:"Required"` //取消原因
	AgencyId int64  `json:"agencyId"`                //操作者Id
}

//RefuseOutMoney 拒绝出款请求struct
type RefuseOutMoney struct {
	SiteId   string
	Id       int64  `json:"id" valid:"Required;Min(1);ErrorCode(60056)"`
	Reason   string `json:"reason" valid:"Required;MinSize(1);ErrorCode()"` //拒绝出款原因
	AgencyId int64  `json:"agencyId"`                                       //操作者Id
}

//PrepareOutMoney 预备出款请求struct
type PrepareOutMoney struct {
	SiteId   string
	Id       int64 `json:"id" valid:"Required;Min(1);ErrorCode(60056)"` //出款请求Id
	AgencyId int64 `json:"agencyId"`                                    //操作者Id
}

//OutMoneyList 出款列表请求struct
type OutMoneyList struct {
	SiteId        string
	SiteIndexId   string  `query:"siteIndexId"`    //站点
	Level         string  `query:"level_id"`       //层级
	OutStatus     int8    `query:"outStatus"`      //状态 5待审核1已出款2预备出款3取消出款4拒绝出款
	StartTime     string  `query:"startTime"`      //起始时间
	AgencyId      int64   `query:"agency_id"`      //代理id
	AgencyAccount string  `query:"agency_account"` //代理账号
	EndTime       string  `query:"endTime"`        //结束时间
	UpperLimit    float64 `query:"upperLimit"`     //金额上限
	LowerLimit    float64 `query:"lowerLimit"`     //金额下限
	ClientType    int     `query:"clientType"`     //出款来源(1.pc 2.wap)
	Automatic     int     `query:"automatic"`      //自动下发(1.自动2.手动0.默认全部)
	SelectBy      int     `query:"selectBy"`       //搜索条件(1.账号2.操作者)
	Conditions    string  `query:"conditions"`     //条件
}

//OutMoneyRequest 出款请求struct
type OutMoneyRequest struct {
	SiteId          string  `json:"site_id"`                                            //站点Id
	SiteIndexId     string  `json:"site_index_id" valid:" MaxSize(4);ErrorCode(10050)"` //站点前台id
	MemberId        int64   `json:"member_id"`                                          //会员Id
	OutMoney        float64 `json:"out_money"`                                          //请求出款金额
	RequestUrl      string  `json:"request_url"`                                        //提出请求网址
	Charge          float64 `json:"charge"`                                             //手续费
	FavourableMoney float64 `json:"favourable_money"`                                   //优惠金额
	ExpeneseMoney   float64 `json:"expenese_money"`                                     //行政费
	HasDiscount     int     `json:"has_discount"`                                       //是否有优惠金额
	ClientType      int     `json:"client_type"`                                        //类型
}
