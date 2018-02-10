package input

//添加红包设置
type RedPacketSetAdd struct {
	Title          string  `json:"title" valid:"Required;MinSize(1);ErrorCode(71008)"`                    //活动标题
	SiteId         string  `json:"siteId" valid:"MaxSize(4);ErrorCode(10050)"`                            //操作站点id
	SiteIndexId    string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`                       //前台id
	Description    string  `json:"description" valid:"ErrorCode(71009)"`                                  //活动简介
	MaxCount       int64   `json:"maxCount"`                                                              //每人最多获奖次数
	StartTime      string  `json:"startTime"  valid:"Required;MinSize(19);MaxSize(19);ErrorCode(71010)"`  //活动开始时间
	EndTime        string  `json:"endTime" valid:"Required;MinSize(19);MaxSize(19);ErrorCode(71011)"`     //活动结束时间
	InStartTime    string  `json:"inStartTime" valid:"Required;MinSize(19);MaxSize(19);ErrorCode(71012)"` //存款起始时间
	InEndTime      string  `json:"inEndTime" valid:"Required;MinSize(19);MaxSize(19);ErrorCode(71013)"`   //存款结束时间
	InSum          float64 `json:"inSum"`                                                                 //存款额度
	AuditStartTime string  `json:"auditStartTime" valid:"Required;MaxSize(20);ErrorCode(71014)"`          //有效打码起始时间
	AuditEndTime   string  `json:"auditEndTime"  valid:"Required;MaxSize(20);ErrorCode(71015)"`           //有效打码结束时间
	BetSum         float64 `json:"betSum"`                                                                //有效打码量
	EndTitle       string  `json:"endTitle"`                                                              //结束标题
	EndDescription string  `json:"endDescription"`                                                        //活动结束简介
	LevelId        string  `json:"levelId"`                                                               //可参加活动会员的分组，0为无限制
	TotalMoney     float64 `json:"totalMoney"  valid:"Required;ErrorCode(71018)"`                         //红包总额
	MinMoney       float64 `json:"minMoney" `                                                             //红包最小额度
	RedNum         int64   `json:"redNum"  valid:"Required;ErrorCode(71019)"`                             //红包数量
	IsIp           int64   `json:"isIp"`                                                                  //1为限制ip；2为不限制
	StyleId        int64   `json:"styleId"  valid:"Required;Min(1);ErrorCode(71020)"`                     //红包皮肤,关联红包皮肤表red_packet_set
	IsShow         int64   `json:"isShow"`                                                                //未到活动开始时间红包是否展示,1为展示，2为不展示
	AppointMoney   int64   `json:"appointMoney"`                                                          //额外领取存款门槛
	RedType        int64   `json:"redType"`                                                               //红包类型，1拼手气，2普通
	Status         int64   `json:"status"`                                                                //活动状态 ，1未开始，2进行中，3已经结束，4已删除
	DepositAchieve float64 `json:"depositAchieve"`                                                        //存款达到
	ReceiveAgain   int8    `json:"receiveAgain"`                                                          //是否能重复领取红包1是2否

	StartTimestamp      int64  `json:"-"` //活动开始时间戳
	EndTimestamp        int64  `json:"-"` //活动结束时间戳
	InStartTimestamp    int64  `json:"-"` //存款起始时间戳
	InEndTimestamp      int64  `json:"-"` //存款结束时间戳
	AuditStartTimestamp int64  `json:"-"` //有效打码起始时间戳
	AuditEndTimestamp   int64  `json:"-"` //有效打码结束时间戳
	CreateIp            string `json:"-"` //创建ip
	CreateUid           int64  `json:"-"` //创建人的uid
	//CreateTime          int64  `json:"-"` //活动创建时间
}

//修改红包设置
type RedPacketSetChange struct {
	Id             int64   `json:"id" valid:"Min(1);ErrorCode(30041)"`                                    //id
	Title          string  `json:"title" valid:"Required;MinSize(1);ErrorCode(71008)"`                    //活动标题
	SiteId         string  `json:"siteId" valid:"MaxSize(4);ErrorCode(10050)"`                            //操作站点id
	SiteIndexId    string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`                       //前台id
	Description    string  `json:"description" valid:"ErrorCode(71009)"`                                  //活动简介
	MaxCount       int64   `json:"maxCount"`                                                              //每人最多获奖次数
	StartTime      string  `json:"startTime"  valid:"Required;MinSize(19);MaxSize(19);ErrorCode(71010)"`  //活动开始时间
	EndTime        string  `json:"endTime" valid:"Required;MinSize(19);MaxSize(19);ErrorCode(71011)"`     //活动结束时间
	InStartTime    string  `json:"inStartTime" valid:"Required;MinSize(19);MaxSize(19);ErrorCode(71012)"` //存款起始时间
	InEndTime      string  `json:"inEndTime" valid:"Required;MinSize(19);MaxSize(19);ErrorCode(71013)"`   //存款结束时间
	InSum          float64 `json:"inSum"`                                                                 //存款额度
	AuditStartTime string  `json:"auditStartTime" valid:"Required;MaxSize(20);ErrorCode(71014)"`          //有效打码起始时间
	AuditEndTime   string  `json:"auditEndTime"  valid:"Required;MaxSize(20);ErrorCode(71015)"`           //有效打码结束时间
	BetSum         float64 `json:"betSum"`                                                                //有效打码量
	EndTitle       string  `json:"endTitle"`                                                              //结束标题
	EndDescription string  `json:"endDescription"`                                                        //活动结束简介
	LevelId        string  `json:"levelId"`                                                               //可参加活动会员的分组，0为无限制
	TotalMoney     float64 `json:"totalMoney"  valid:"Required;ErrorCode(71018)"`                         //红包总额
	MinMoney       float64 `json:"minMoney" `                                                             //红包最小额度
	RedNum         int64   `json:"redNum"  valid:"Required;ErrorCode(71019)"`                             //红包数量
	IsIp           int64   `json:"isIp"`                                                                  //1为限制ip；2为不限制
	StyleId        int64   `json:"styleId"  valid:"Required;Min(1);ErrorCode(71020)"`                     //红包皮肤,关联红包皮肤表red_packet_set
	IsShow         int64   `json:"isShow"`                                                                //未到活动开始时间红包是否展示,1为展示，2为不展示
	AppointMoney   int64   `json:"appointMoney"`                                                          //额外领取存款门槛
	RedType        int64   `json:"redType"`                                                               //红包类型，1拼手气，2普通
	Status         int64   `json:"status"`                                                                //活动状态 ，1未开始，2进行中，3已经结束，4已删除
	DepositAchieve float64 `json:"depositAchieve"`                                                        //存款达到
	ReceiveAgain   int8    `json:"receiveAgain"`                                                          //是否能重复领取红包1是2否

	StartTimestamp      int64  `json:"-"` //活动开始时间戳
	EndTimestamp        int64  `json:"-"` //活动结束时间戳
	InStartTimestamp    int64  `json:"-"` //存款起始时间戳
	InEndTimestamp      int64  `json:"-"` //存款结束时间戳
	AuditStartTimestamp int64  `json:"-"` //有效打码起始时间戳
	AuditEndTimestamp   int64  `json:"-"` //有效打码结束时间戳
	CreateIp            string `json:"-"` //创建ip
	CreateUid           int64  `json:"-"` //创建人的uid
	//CreateTime          int64  `json:"-"` //活动创建时间
}

//查询设置
type RedPacketSet struct {
	Id         int64 `json:"id"  valid:"Required;ErrorCode(71022)"`           //设置id
	ClientType int64 `json:"client_type" valid:"Range(0,3);ErrorCode(70023)"` //客户端类型0pc 1wap 2android 3ios
}

//查询列表
type RedPacketSetList struct {
	SiteId          string `query:"siteId" valid:"MaxSize(4);ErrorCode(10050)"`               //操作站点id
	SiteIndexId     string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台id
	Status          int64  `query:"status"`                                                   //活动状态 ，1未开始，2进行中，3已经结束，4已删除
	CreateTimeStart string `query:"createTimeStart"`                                          //活动创建时间
	CreateTimeEnd   string `query:"createTimeEnd"`                                            //活动创建时间
}

//查询列表详情
type RedPacketSetListInfo struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(10050)"`      //操作站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台id
	Id          int64  `query:"id" valid:"Min(1);ErrorCode(30041)"`              //红包id
}

//红包查看
type RedBagSee struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(10050)"`      //操作站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台id
	Id          int64  `query:"id" valid:"Min(1);ErrorCode(30041)"`              //红包id
}

//终止红包设置
type RedPacketSetDelete struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(30041)"` //id
	SiteId      string `json:"siteId"`                             //操作站点id
	SiteIndexId string `json:"siteIndexId"`                        //前台id
}
