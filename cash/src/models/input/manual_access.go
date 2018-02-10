package input

//添加一条人工存款
type ManualAccessAdd struct {
	SiteId            string  `json:"siteId"`                                                //用户站点ID
	SiteIndexId       string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`       //用户前台id
	Account           string  `json:"account" valid:"Required;ErrorCode(30124)"`             // 会员账号
	Money             float64 `json:"money" valid:"Required;ErrorCode(30127)"`               // 本次存款金额
	IsDepositDiscount int8    `json:"isDepositDiscount" valid:"Range(1,2);ErrorCode(30137)"` // 是否有存款优惠 1:是   2:否
	DepositDiscount   float64 `json:"depositDiscount" valid:"Required;ErrorCode(30128)"`     //存款优惠
	IsRemitDiscount   int8    `json:"isRemitDiscount" valid:"Range(1,2);ErrorCode(30129)"`   // 是否有汇款优惠 1:是   2:否
	RemitDiscount     float64 `json:"remitDiscount" valid:"Required;ErrorCode(30130)"`       // 汇款优惠
	IsCodeCount       int8    `json:"isCodeCount" valid:"Range(1,2);ErrorCode(30131)"`       // 是否有综合稽核 1:是   2:否
	CodeCount         int64   `json:"codeCount" valid:"ErrorCode(30132)"`                    // 综合稽核打码量
	IsRoutineCheck    int8    `json:"isRoutineCheck" valid:"Range(0,2);ErrorCode(30133)"`    // 是否有常态性稽核 1:是   2:否
	DepositType       int8    `json:"depositType" valid:"Range(1,9);ErrorCode(30134)"`       // 存款项目(类型)1人工存入2存款优惠3负数额度归零4取消出款5返点优惠6活动优惠7其他8体育投注余额9额度掉单
	IsWriteRebate     int8    `json:"isWriteRebate" valid:"Range(1,2);ErrorCode(30135)"`     //是否写入退佣 1:是   2:否
	Remark            string  `json:"remark" valid:"MaxSize(255);ErrorCode(30136)"`          // 备注
	DoAgencyId        int64   `json:"doAgencyId" `                                           //操作人id(agency表主键)
	DoAgencyAccount   string  `json:"doAgencyAccount"`                                       // 操作人账号
}

//存取款记录列表
type ManualAccessList struct {
	SiteId      string `query:"siteId"`                                          //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Account     string `query:"account"`                                         //会员账号
	StartTime   string `query:"startTime"`                                       //开始时间
	EndTime     string `query:"endTime"`                                         //结束时间
}

//添加多条人工存款
type AddManualAccess struct {
	SiteId            string   `json:"siteId"`                                                //用户站点ID
	SiteIndexId       string   `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`       //用户前台id
	Types             int8     `json:"types" valid:"Range(1,2);ErrorCode(30140)"`             //批量方式 1:账号   2:层级
	Account           []string `json:"account"`                                               //会员账号
	LevelId           []string `json:"levelId"`                                               //会员层级
	Money             float64  `json:"money" valid:"Required;ErrorCode(30127)"`               // 本次存款金额
	IsDepositDiscount int8     `json:"isDepositDiscount" valid:"Range(1,2);ErrorCode(30137)"` // 是否有存款优惠 1:是   2:否
	DepositDiscount   float64  `json:"depositDiscount" valid:"Required;ErrorCode(30128)"`     //存款优惠
	IsRemitDiscount   int8     `json:"isRemitDiscount" valid:"Range(1,2);ErrorCode(30129)"`   // 是否有汇款优惠 1:是   2:否
	RemitDiscount     float64  `json:"remitDiscount" valid:"Required;ErrorCode(30130)"`       // 汇款优惠
	IsCodeCount       int8     `json:"isCodeCount" valid:"Range(1,2);ErrorCode(30131)"`       // 是否有综合稽核 1:是   2:否
	CodeCount         int64    `json:"codeCount" valid:"ErrorCode(30132)"`                    // 综合稽核打码量
	IsRoutineCheck    int8     `json:"isRoutineCheck" valid:"Range(1,2);ErrorCode(30133)"`    // 是否有常态性稽核 1:是   2:否
	DepositType       int8     `json:"depositType" valid:"Required;ErrorCode(30134)"`         // 存款项目(类型)1人工存入2存款优惠3负数额度归零4取消出款5返点优惠6活动优惠
	IsWriteRebate     int8     `json:"isWriteRebate" valid:"Range(1,2);ErrorCode(30135)"`     //是否写入退佣 1:是   2:否
	Remark            string   `json:"remark" valid:"MaxSize(255);ErrorCode(30136)"`          // 备注
	DoAgencyId        int64    `json:"doAgencyId" `                                           //操作人id(agency表主键)
	DoAgencyAccount   string   `json:"doAgencyAccount"`                                       // 操作人账号
}

//人工取款
type ManualWithdrawalAdd struct {
	SiteId          string  `json:"siteId"`                                            //用户站点ID
	SiteIndexId     string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`   //用户前台id
	Account         string  `json:"account" valid:"Required;ErrorCode(30124)"`         // 会员账号
	Money           float64 `json:"money" valid:"Required;ErrorCode(30127)"`           // 本次取款金额
	Remark          string  `json:"remark" valid:"MaxSize(255);ErrorCode(30136)"`      // 备注
	DepositType     int8    `json:"depositType" valid:"Range(10,16);ErrorCode(30134)"` //操作类型15重复出款16公司入款误存10公司负数回冲11手动申请出款12扣除非法下注派彩13放弃存款优惠14其他17体育投注余额18额度掉单
	DoAgencyId      int64   `json:"doAgencyId" `                                       //操作人id(agency表主键)
	DoAgencyAccount string  `json:"doAgencyAccount"`                                   // 操作人账号
}

//账目汇总
type ManualAccessLists struct {
	SiteId      string `query:"siteId"`                                          //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	AgencyId    int64  `query:"agencyId"`                                        //代理id
	Account     string `query:"account" valid:"MaxSize(12);ErrorCode(30009)"`    //会员账号
	DateTime    string `query:"dateTime"`                                        //日期
}
