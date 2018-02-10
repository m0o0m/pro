package input

//会员列表筛选条件
type MemberIndex struct {
	SiteId      string `query:"siteId"`
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	FirstId     int64  `query:"firstId"`                                      //股东id
	SecondId    int64  `query:"secondId"`                                     //总代id
	AgencyId    int64  `query:"agencyId"`                                     //所属代理id
	Status      int    `query:"status" valid:"Range(0, 2);ErrorCode(30050)"`  //状态状态 0全部,1正常,2禁用
	Online      int    `query:"online" valid:"Range(0, 10);ErrorCode(50083)"` //是否在线 0全部,1pc在线,2wap在线,3android在线,4ios在线,5pc离线,6wap离线,7android离线,8ios离线,9全部在线10全部离线
	Source      int    `query:"source" valid:"Range(0, 4);ErrorCode(60033)"`  //注册来源 0全部,1pc,2wap,3android,4ios
	StartTime   string `query:"startTime"`                                    //注册时间,开始
	EndTime     string `query:"endTime"`                                      //注册时间,结束
	Type        int    `query:"type"`                                         //类型 1账号,2姓名,3注册ip,4登录ip,5手机,6银行卡,7邮箱,8qq,9微信号
	TypeValue   string `query:"typeValue"`                                    //类型值
	IsVague     int    `query:"isVague"`                                      //是否模糊查询 1是,2否
	Sort        int    `query:"sort"`                                         //排序方式 1desc从大到小,2asc从小到大
	SortBy      string `query:"sortBy"`                                       //排序类型 1注册时间,2账号,3登录时间,4系统余额，其他的为视讯余额
	PageSize    int    `query:"pageSize"`                                     //每页显示
	Page        int    `query:"page"`                                         //页码
	IsHide      int8   `query:"isHide" valid:"Range(0,2);ErrorCode(50159)"`   //是否隐藏  1是2否
}

//站点会员总数和今日注册人数
type MemberNumber struct {
	SiteId      string `query:"siteId"`
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	FirstId     int64  `query:"firstId"`                    //股东id
	SecondId    int64  `query:"secondId"`                   //总代id
	AgencyId    int64  `query:"agencyId"`                   //所属代理id
	Status      int    `query:"status" valid:"Range(0, 2)"` //状态状态 0全部,1正常,2禁用
	Online      int    `query:"online" valid:"Range(0, 2)"` //是否在线 0全部,1在线,2离线
	Source      int    `query:"source" valid:"Range(0, 4)"` //注册来源 0全部,1pc,2wap,3android,4ios
	StartTime   string `query:"startTime"`                  //注册时间,开始
	EndTime     string `query:"endTime"`                    //注册时间,结束
	Type        int    `query:"type"`                       //类型 1账号,2姓名,3注册ip,4登录ip,5手机,6银行卡,7邮箱,8qq,9微信号
	TypeValue   string `query:"typeValue"`                  //类型值
	IsVague     int    `query:"isVague"`                    //是否模糊查询 1是,2否
	Sort        int    `query:"sort"`                       //排序方式 1desc从大到小,2asc从小到大
	SortBy      int    `query:"sortBy"`                     //排序类型 1注册时间,2账号,3登录时间
	PageSize    int    `query:"pageSize"`                   //每页显示
	Page        int    `query:"page"`                       //页码
}

//修改会员状态过滤条件/获取会员基本资料/获取会员详细资料
type MemberStatus struct {
	SiteId      string
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	Id          int64  `query:"id" json:"id" valid:"Required;ErrorCode(10020)"` //会员id
}

type MemberBankInfo struct {
	SiteId      string
	SiteIndexId string `query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"`
	Card        string `query:"card" valid:"Required;ErrorCode(30060)"` //卡号
}

//修改会员基本资料
type MemberBaseInfo struct {
	SiteId         string
	SiteIndexId    string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	Id             int    `json:"id" valid:"Required;Min(1);ErrorCode(10020)"`            //会员id
	NewPassword    string `json:"newPassword" valid:"MaxSize(12);ErrorCode(20005)"`       //新密码
	ReplyPassword  string `json:"replyPassword" valid:"MaxSize(12);ErrorCode(20007)"`     //重复密码
	IsEditPassword int    `json:"isEditPassword" valid:"Range(1, 2);ErrorCode(10000)"`    //是否可以修改密码 1是,2否
	Realname       string `json:"realname" valid:"Required;MaxSize(20);ErrorCode(20011)"` //会员姓名
}

//修改会员详细资料
type MemberDetail struct {
	SiteId       string
	SiteIndexId  string   `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	Id           int      `json:"id" valid:"Required;Min(1);ErrorCode(10020)"`      //会员id
	Realname     string   `json:"realname" valid:"MaxSize(20);ErrorCode(20011)"`    //会员姓名
	Birthday     string   `json:"birthday" valid:"MaxSize(10);ErrorCode(20012)"`    //出生日期不合法
	Card         string   `json:"card" valid:"MaxSize(20);ErrorCode(20013)"`        //身份证号不合法
	Mobile       string   `json:"mobile" valid:"MaxSize(11);ErrorCode(20014)"`      //手机号不合法
	Email        string   `json:"email" valid:"MaxSize(30);ErrorCode(20015)"`       //邮箱不合法
	QQ           string   `json:"qq" valid:"MaxSize(12);ErrorCode(20016)"`          // qq不合法
	Wechat       string   `json:"wechat" valid:"MaxSize(12);ErrorCode(20017)"`      //微信号不合法
	DrawPassword string   `json:"drawPassword" valid:"MaxSize(4);ErrorCode(20018)"` //取款密码不合法
	Remark       string   `json:"remark" valid:"MaxSize(200);ErrorCode(20019)"`     //备注不合法
	Ids          string   `json:"ids"`                                              //会员银行卡id
	BankIds      string   `json:"bankIds"`                                          //银行id
	BankAccount  []string `json:"bankAccount"`                                      //银行帐号
	CardAddress  []string `json:"cardAddress"`                                      //开户行
}

//修改会员银行
type MemberBankEdit struct {
	SiteId      string
	SiteIndexId string `json:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"`
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(30057)"`
	BankId      int64  `json:"bank_id" valid:"Required;Min(1);ErrorCode(30059)"` //卡类型
	Card        string `json:"card" valid:"Required;ErrorCode(30060)"`           //卡号
	CardName    string `json:"card_name" valid:"Required;ErrorCode(30061)"`      //卡账号
	CardAddress string `json:"card_address" valid:"Required;ErrorCode(30062)"`   //卡地址
}

type MemberInfo struct {
	SiteId      string
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	Account     string `query:"account" valid:"Required;ErrorCode(30124)"` //会员账号
}

//会员登录
type MemberSign struct {
	Account  string `json:"account" valid:"MinSize(4);MaxSize(12);ErrorCode(20000)"`
	Password string `json:"password" valid:"MinSize(4);MaxSize(12);ErrorCode(20000)"`
	Code     string `json:"code" valid:"MinSize(4);MaxSize(4);ErrorCode(20021)"` //验证码
}

//会员注册
type MemberRegister struct {
	Introducer       string `json:"introducer"`         //代理注册
	IntroducerMember string `json:"introducer_member"`  //会员推广注册
	Account          string `json:"account" `           //账号
	Password         string `json:"password" `          //密码
	ConfirmPassword  string `json:"confirm_password" `  //确认密码
	Site             string `json:"site_id" `           //站点id
	SiteIndexId      string `json:"site_index_id" `     //站点前台id
	RealName         string `json:"real_name"`          //真实姓名
	Code             string `json:"code" `              //验证码
	WithdrawPassword string `json:"withdraw_password" ` //取款密码
	IsAgreeDeal      string `json:"isAgreeDeal" `       //是否同意协议(1.同意2.不同意)
	Email            string `json:"email"`              //邮箱
	PassPort         string `json:"pass_port" `         //身份证 //valid:"Match(^(\d{6})(\d{4})(\d{2})(\d{2})(\d{3})([0-9]|X)$);ErrorCode(50091)"
	Qq               string `json:"qq" `                //qq //valid:"Match([1-9][0-9]{4,});ErrorCode(20016)"
	Phone            string `json:"phone" `             //电话
	Birthday         string `json:"birthday"`           //生日
	Wechat           string `json:"wechat"`             //微信
	BankCard         string `json:"bank_card" `         //银行卡号 // valid:"Match(^([0-9]){16,19}$);ErrorCode(60200)"
	LocalCode        string `json:"local_code"`         //注册地区号
}

//会员注册校验是否重复的struct
type CheckRegister struct {
	Types      int    `json:"types" valid:"Required;Range(1,4);ErrorCode(60079)"`      //注册终端类型(1.pc2.wap3.ios,4android)
	Conditions string `json:"conditions" valid:"Required;MixSize(1);ErrorCode(60080)"` //填写的内容
}

//退出登录
type Logout struct {
	Id int64 `json:"id" valid:"Required;Min(1);ErrorCode(10020)"` //会员id
}

//修改密码
type ChangePassword struct {
	Types           int    `json:"types" valid:"Required;Range(1,2);ErrorCode(60075)"`           //密码修改类型(1.登录密码2.取款密码)
	BeforePassword  string `json:"beforePassword" valid:"Required;MinSize(1);ErrorCode(60076)"`  //原密码
	Password        string `json:"password" valid:"Required;MinSize(1);ErrorCode(60077)"`        //新密码
	ConfirmPassword string `json:"confirmPassword" valid:"Required;MinSize(1);ErrorCode(60078)"` //重复的新密码
}

//Wap端会员线上存款-微信
type WapDeposit struct {
	Account string  `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"` //会员账号
	Money   float64 `json:"money"`                                            //存款金额
}

//会员线上存款-微信、支付宝、百度钱包、京东钱包、visa支付、qq钱包等
type WapOnlineDeposit struct {
	Account  string  `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"` //会员账号
	Money    float64 `json:"money"`                                            //存款金额
	PaidType int     `json:"paidType"`                                         //支付类型id
	IsFast   int8    `json:"isFast"`                                           //是否快速充值 1是2否
	Id       int64   `json:"id"`                                               //设定的第三方id
}

//快捷支付传入会员
type GetFastIncomeData struct {
	Account   string `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"`   //会员账号
	Reaccount string `json:"reaccount" valid:"RangeSize(4,12);ErrorCode(20000)"` //确认会员账号
}

//网银
type OnlineBankDeposit struct {
	Account  string  `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"` //会员账号
	Money    float64 `json:"money"`                                            //存款金额
	PaidType int     `json:"paidType"`                                         //支付类型id
	Bank     string  `json:"bank"`                                             //收款银行
	IsFast   int8    `json:"isFast"`                                           //是否快速充值 1是2否
	Id       int64   `json:"id"`                                               //设定的第三方id
}

//会员线上存款-微信、支付宝、百度钱包、京东钱包、visa支付、qq钱包等回调处理
type WapOnlineDepositCallback struct {
	Status  bool     `json:"status"`  //状态 true
	Code    int64    `json:"code"`    //状态码 200
	Message string   `json:"message"` //信息 成功
	Data    []string `json:"data"`    //[] or 入款失败
}

//Wap端会员线上存款-网银
type WapDepositBank struct {
	Account string  `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"` //会员账号
	Money   float64 `json:"money"`                                            //存款金额
	Bank    int64   `json:"id" valid:"Min(1);ErrorCode(10220)"`               //卡类型
}

//Wap端会员线上存款-点卡
type WapDepositCard struct {
	Account      string  `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"`    //会员账号
	Money        float64 `json:"money"`                                               //存款金额
	CardMoney    float64 `json:"cardMoney" valid:"MinFloat64(0.00);ErrorCode(10222)"` //卡面额
	CardNumber   string  `json:"cardNumber"`                                          //卡序列号
	CardPassword string  `json:"cardPassword"`                                        //卡密码
	Bank         int64   `json:"id" valid:"Min(1);ErrorCode(10220)"`                  //卡类型
}

//Wap端会员线上存款-点卡
type WapDepositCardIn struct {
	Account string  `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"` //会员账号
	Money   float64 `json:"money"`                                            //存款金额
}

//Wap端会员线上存款-点卡
type WapOnlineDepositCard struct {
	Money        float64 `json:"money"`                                               //存款金额
	CardMoney    float64 `json:"cardMoney" valid:"MinFloat64(0.00);ErrorCode(10222)"` //卡面额
	CardNumber   string  `json:"cardNumber"`                                          //卡序列号
	CardPassword string  `json:"cardPassword"`                                        //卡密码
	Bank         int64   `json:"id" valid:"Min(1);ErrorCode(10220)"`                  //卡类型
	PaidType     int     `json:"paidType"`                                            //支付类型id
	Id           int64   `json:"id"`                                                  //设定的第三方id
}

//Wap端会员公司存款 -微信
type WapCompanyDeposit struct {
	Account         string  `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"` //会员账号
	Money           float64 `json:"money"`                                            //存款金额
	DepositDisCount int64   `json:"depositDiscount"`                                  //存款优惠类型
	Time            string  `json:"time" valid:"Date;ErrorCode(10223)"`               //转帐时间
	DepositAccount  string  `json:"depositAccount" valid:"Required;ErrorCode(10224)"` //转账账号
}

//Wap端会员公司存款 -网银
type WapCompanyDepositBank struct {
	Account        string  `json:"account" valid:"RangeSize(4,12);ErrorCode(20000)"` //会员账号
	Money          float64 `json:"money"`                                            //存款金额
	DepositName    string  `json:"depositName" valid:"Chinese;ErrorCode(10225)"`     //存款人姓名
	DepositType    int8    `json:"depositType" valid:"Range(1,2);ErrorCode(10226)"`  //转账方式(网银转账,ATM现金入款)
	Time           string  `json:"time" valid:"Date;ErrorCode(10223)"`               //转帐时间
	DepositAccount string  `json:"depositAccount" valid:"Required;ErrorCode(10224)"` //转账账号
	Bank           int64   `json:"id" valid:"Min(1);ErrorCode(10220)"`               //卡类型
	BranchBank     string  `json:"branchBank"`                                       //分行名称
	City           string  `json:"city"`                                             //城市
	Provice        string  `json:"provice"`                                          //省份
}

//检测会员支付密码
type CheckDrawPass struct {
	MemberId     int64  `query:"id"`            //会员id
	DrawPassword string `query:"draw_password"` //取款密码
}

// 交易记录附带参数获取
type IsParameter struct {
	IsOther int64 `query:"isOther"` //附带参数
}

//会员视讯余额
type MemberVideoBalance struct {
	SiteId      string
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	Id          int64  `query:"id" valid:"Required;ErrorCode(10020)"` //会员id
}

//批量启用禁用会员
type BatchMember struct {
	Ids    []int64 `json:"ids" query:"ids"` //会员id
	Status int8    `json:"status" query:"status" valid:"Range(1,2);ErrorCode(10229)"`
}

//批量踢线会员
type OfflineMember struct {
	Ids []int64 `json:"ids" query:"ids"` //会员id
}

//会员排序
type MemberSortDrop struct {
	SiteId string
}

//获取网银在线的银行
type GetOnlineIncomeBank struct {
	PayId int64 `json:"pay_id" query:"pay_id"` //三方配置id
}
