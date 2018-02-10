package back

//设置列表
type RedPacketSetList struct {
	Id         int64   `json:"id" xorm:"'id' PK autoincr"`
	CreateTime int64   `json:"createTime" xorm:"'create_time' created"` //活动创建时间
	StartTime  int64   `json:"startTime" xorm:"start_time"`             //活动开始时间
	EndTime    int64   `json:"endTime" xorm:"end_time"`                 //活动结束时间
	BetSum     float64 `json:"betSum" xorm:"bet_sum"`                   //有效打码量
	InSum      float64 `json:"inSum" xorm:"in_sum"`                     //存款额度
	LevelId    string  `json:"levelId" xorm:"level_id"`                 //可参加活动会员的分组，0为无限制
	IsIp       int64   `json:"isIp" xorm:"is_ip"`                       //1为限制ip；2为不限制
	MinMoney   float64 `json:"minMoney" xorm:"min_money"`               //红包最小额度
	TotalMoney float64 `json:"totalMoney" xorm:"total_money"`           //红包总额
	Title      string  `json:"title" xorm:"title"`                      //活动标题
	Status     int64   `json:"status" xorm:"status"`                    //活动状态 ，1未开始，2进行中，3已经结束，4已删除
	IsGenerate int64   `json:"isGenerate" xorm:"is_generate"`           //是否生成 1,未生成,2,已生成',
}

//红包活动
type RedPacket struct {
	Id             int64   `xorm:"'id' PK autoincr" json:"id"`
	Title          string  `xorm:"title" json:"title"`                      //活动标题
	MaxCount       int64   `xorm:"max_count" json:"maxCount"`               //每人最多获奖次数
	StartTime      int64   `xorm:"start_time" json:"startTime"`             //活动开始时间
	InStartTime    int64   `xorm:"in_start_time" json:"inStartTime"`        //存款起始时间
	InEndTime      int64   `xorm:"in_end_time" json:"inEndTime"`            //存款结束时间
	InSum          float64 `xorm:"in_sum" json:"inSum"`                     //存款额度
	AuditStartTime int64   `xorm:"audit_start_time" json:"auditStartTime"`  //有效打码起始时间
	AuditEndTime   int64   `xorm:"audit_end_time" json:"auditEndTime"`      //有效打码结束时间
	BetSum         float64 `xorm:"bet_sum" json:"betSum"`                   //有效打码量
	LevelId        string  `xorm:"level_id" json:"levelId"`                 //可参加活动会员的分组，0为无限制
	TotalMoney     float64 `xorm:"total_money" json:"totalMoney"`           //红包总额
	MinMoney       float64 `xorm:"min_money" json:"minMoney"`               //红包最小额度
	RedNum         int64   `xorm:"red_num" json:"redNum"`                   //红包数量
	CreateTime     int64   `xorm:"'create_time' created" json:"createTime"` //活动创建时间
	IsIp           int64   `xorm:"is_ip" json:"isIp"`                       //1为限制ip；2为不限制
	StyleId        int64   `xorm:"style_id" json:"styleId"`                 //红包皮肤,关联红包皮肤表red_packet_set
	IsShow         int64   `xorm:"is_show" json:"isShow"`                   //1为展示，2为不展示
	AppointMoney   int64   `xorm:"appoint_money" json:"appointMoney"`       //额外领取存款门槛
	RedType        int64   `xorm:"red_type" json:"redType"`                 //红包类型，1拼手气，2普通
	Status         int64   `xorm:"status" json:"status"`                    //活动状态 ，1未开始，2进行中，3已经结束，4已删除
	IsGenerate     int64   `xorm:"is_generate" json:"isGenerate"`           //是否生成 1,未生成,2,已生成',
	Opencount      int64   `json:"opencount"`
	Closecount     int64   `json:"closecount"`
	Pic            int64   `json:"pic"`
}

//查看红包
type RedPacketLogInfoBack struct {
	Id           int64   `xorm:"id" json:"id"`
	Money        float64 `xorm:"money" json:"money"`            //红包金额
	Account      string  `xorm:"account" json:"account"`        //用户名
	LevelId      string  `json:"level_id" json:"levelId"`       //可参加活动会员的分组，0为无限制
	CreateTime   int64   `xorm:"create_time" json:"createTime"` //创建时间
	BalanceMoney float64 `xorm:"balance_money" json:"balanceMoney"`
	SiteId       string  `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId  string  `xorm:"site_index_id" json:"siteIndexId"` //前台id
	MakeSure     int64   `xorm:"make_sure"`                        //是否被抢1,未被抢,2已抢
}

//红包总计
type RedBagNumTotal struct {
	RedBag       int64   `json:"redBag"`       //红包总个数
	Already      int64   `json:"already"`      //已经领取的个数
	Spare        int64   `json:"spare"`        //剩余红包
	RedBagMoney  float64 `json:"redBagMoney"`  //红包金额
	AlreadyMoney float64 `json:"alreadyMoney"` //已经领取金额
	SpareMoney   float64 `json:"spareMoney"`   //剩余金额
}

//红包设置详情
type RedBagInfo struct {
	Id             int64   `xorm:"id" json:"id"`
	Title          string  `xorm:"title" json:"title"`                     //活动标题
	SiteId         string  `xorm:"site_id" json:"siteId"`                  //站点id
	SiteIndexId    string  `xorm:"site_index_id" json:"siteIndexId"`       //前台id
	MaxCount       int64   `xorm:"max_count" json:"maxCount"`              //每人最多获奖次数
	StartTime      int64   `xorm:"start_time" json:"startTime"`            //活动开始时间
	EndTime        int64   `xorm:"end_time" json:"endTime"`                //活动结束时间
	InStartTime    int64   `xorm:"in_start_time" json:"inStartTime"`       //存款起始时间
	InEndTime      int64   `xorm:"in_end_time" json:"inEndTime"`           //存款结束时间
	InSum          float64 `xorm:"in_sum" json:"inSum"`                    //存款额度
	AuditStartTime int64   `xorm:"audit_start_time" json:"auditStartTime"` //有效打码起始时间
	AuditEndTime   int64   `xorm:"audit_end_time" json:"auditEndTime"`     //有效打码结束时间
	BetSum         float64 `xorm:"bet_sum" json:"betSum"`                  //有效打码量
	EndTitle       string  `xorm:"end_title" json:"endTitle"`              //结束标题
	LevelId        string  `xorm:"level_id" json:"leveId"`                 //可参加活动会员的分组，0为无限制
	TotalMoney     float64 `xorm:"total_money" json:"totalMoney"`          //红包总额
	MinMoney       float64 `xorm:"min_money" json:"minMoney"`              //红包最小额度
	RedNum         int64   `xorm:"red_num" json:"redNum"`                  //红包数量
	CreateIp       string  `xorm:"create_ip" json:"createIp"`              //创建ip
	CreateUid      int64   `xorm:"create_uid" json:"createUid"`            //创建管理员的uid
	CreateTime     int64   `xorm:"create_time" json:"createTime"`          //活动创建时间
	IsIp           int64   `xorm:"is_ip" json:"isIp"`                      //1为限制ip；2为不限制
	StyleId        int64   `xorm:"style_id" json:"styleId"`                //红包皮肤,关联红包皮肤表red_packet_set
	IsShow         int64   `xorm:"is_show" json:"isShow"`                  //1为展示，2为不展示
	AppointMoney   int64   `xorm:"appoint_money" json:"appointMoney"`      //额外领取存款门槛
	RedType        int64   `xorm:"red_type" json:"redType"`                //红包类型，1拼手气，2普通
	Status         int64   `xorm:"status" json:"status"`                   //活动状态 ，1未开始，2进行中，3已经结束，4已删除
	IsGenerate     int64   `xorm:"is_generate" json:"isGenerate"`          //是否生成 1,未生成,2,已生成',
	DepositAchieve float64 `xorm:"deposit_achieve" json:"depositAchieve"`  //存款达到
	ReceiveAgain   int8    `xorm:"receive_again" json:"receiveAgain"`      //是否能重复领取红包1是2否
	BgPic          string  `xorm:"bg_pic" json:"bgPic"`                    //背景图片
	ClickPic       string  `xorm:"click_pic" json:"clickPic"`              //点击图片
}
