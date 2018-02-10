package input

//会员个人资料
type MemberInfoSelf struct {
	Id          int64
	SiteId      string
	SiteIndexId string
}

//今日交易记录
type PayRecordToday struct {
	Id          int64
	SiteId      string
	SiteIndexId string
	StartTime   string `query:"start_time" valid:"Required;MaxSize(10);ErrorCode(50118)"` //开始时间
	EndTime     string `query:"end_time" valid:"Required;MaxSize(10);ErrorCode(50119)"`   //结束时间
}

//修改密码
type MemberPassword struct {
	Id               int64
	SiteId           string
	SiteIndexId      string `json:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"`
	OriginalPassword string `json:"original_password" valid:"Required;MaxSize(12);ErrorCode(20005)"`
	Password         string `json:"password" valid:"Required;MaxSize(12);ErrorCode(20006)"`
	RePassword       string `json:"re_password" valid:"Required;MaxSize(12);ErrorCode(20007)"`
	Type             int    `json:"type" valid:"Required;Min(1);ErrorCode(50114)"`
}

//添加会员出款银行
type MemberBankAdd struct {
	SiteId      string
	SiteIndexId string
	MemberId    int64  //会员Id
	BankId      int64  `json:"bank_id" valid:"Required;Min(1);ErrorCode(30059)"`            //卡类型
	Card        string `json:"card" valid:"Required;MaxSize(19);ErrorCode(30060)"`          //卡号
	CardName    string `json:"card_name" valid:"Required;MaxSize(20);ErrorCode(30061)"`     //收款人姓名
	CardAddress string `json:"card_address" valid:"Required;MaxSize(100);ErrorCode(30062)"` //卡地址
	Password    string `json:"password"`                                                    //取款密码
}

//修改会员出款银行
type MemberBankChange struct {
	SiteId      string
	SiteIndexId string
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(50013)"`
	MemberId    int64  //会员Id
	BankId      int64  `json:"bank_id" valid:"Required;Min(1);ErrorCode(30059)"`            //卡类型
	Card        string `json:"card" valid:"Required;MaxSize(19);ErrorCode(30060)"`          //卡号
	CardName    string `json:"card_name" valid:"Required;MaxSize(20);ErrorCode(30061)"`     //收款人姓名
	CardAddress string `json:"card_address" valid:"Required;MaxSize(100);ErrorCode(30062)"` //卡地址
}

//删除会员出款银行
type MemberBankDelete struct {
	Id          int64 `query:"id" valid:"Required;Min(1);ErrorCode(50013)"`
	MemberId    int64
	SiteId      string
	SiteIndexId string
}

//会员出款银行列表
type MemberBankList struct {
	MemberId    int64
	SiteId      string
	SiteIndexId string
}

//查询会员一条出款银行信息
type OneMemberBankInfo struct {
	SiteId      string
	SiteIndexId string
	Id          int64 `query:"id" valid:"Required;Min(1);ErrorCode(50013)"`
	MemberId    int64
}

//修改手机号
type PhoneNum struct {
	SiteId      string
	SiteIndexId string
	MemberId    int64
	PhoneNum    string `json:"phone_num" valid:"Required;MaxSize(11);ErrorCode(20014)"`
	Code        string `json:"code" valid:"Required;MaxSize(6);ErrorCode(20021)"`
	LocalCode   string `json:"local_code" valid:"Required;MaxSize(4);ErrorCode(50120)"`
}

//修改邮箱
type EmailNum struct {
	SiteId      string
	SiteIndexId string
	MemberId    int64
	EmailNum    string `json:"email_num" valid:"Required;MaxSize(50);ErrorCode(20015)"`
}

//修改生日
type BirthdayNum struct {
	SiteId      string
	SiteIndexId string
	MemberId    int64
	BirthdayNum string `json:"birthday_num" valid:"Required;MaxSize(10);ErrorCode(20012)"`
}

//修改资料
type EditMeans struct {
	SiteId      string
	SiteIndexId string
	MemberId    int64
	Realname    string `json:"realname"`
	BirthdayNum string `json:"birthday_num" valid:"MaxSize(10);ErrorCode(20012)"`
	EmailNum    string `json:"email_num" valid:"MaxSize(50);ErrorCode(20015)"`
	PhoneNum    string `json:"phone_num" valid:"MaxSize(11);ErrorCode(20014)"`
	QqNum       string `json:"qq_num" valid:"MaxSize(12);ErrorCode(20016)"`
	Wechat      string `json:"wechat" valid:"MaxSize(20);ErrorCode(20017)"`
	Card        string `json:"card" valid:"MaxSize(18);ErrorCode(20013)"`
	LocalCode   string `json:"local_code" valid:"MaxSize(4);ErrorCode(50120)"`
	Remark      string `json:"remark"`
}
