package back

//会员个人资料
type MemberInfoSelfBack struct {
	Id            int64   `xorm:"'id' PK autoincr" json:"id"`
	SiteId        string  `xorm:"site_id" json:"site_id"`                   //站点Id
	SiteIndexId   string  `xorm:"site_index_id" json:"site_index_id"`       //站点前台Id
	Account       string  `xorm:"account" json:"account"`                   //登陆账号
	Realname      string  `xorm:"realname" json:"realname"`                 //真实姓名
	Balance       float64 `xorm:"balance" json:"balance"`                   //账号余额
	LastLoginTime int64   `xorm:"last_login_time" json:"last_login_time"`   //上次登录时间
	CreateTime    int64   `xorm:"'create_time' created" json:"create_time"` //创建时间
	Mobile        string  `xorm:"mobile" json:"mobile"`                     //手机号码
	Email         string  `xorm:"email" json:"email"`                       //邮箱
	Birthday      int64   `xorm:"birthday" json:"birthday"`                 //生日
	Card          string  `xorm:"card" json:"card"`                         //卡号
	BankCard      string  `xorm:"card" json:"bank_card"`
}

//连表查询的系统余额和各种账户余额
type MemberBalanceBack struct {
	Id            int64   `xorm:"id" json:"id"`
	Platform      string  `xorm:"platform" json:"platform"`
	Balance       float64 `xorm:"balance" json:"balance"`
	OthersBalance float64 `xorm:"balance" json:"others_balance"`
}

//系统余额以及各种账户余额，总余额
type MemberBalanceTotalBack struct {
	Name    string  `json:"name" json:"name"`
	Balance float64 `json:"balance" json:"balance"`
	Type    int8    `json:"type" json:"type"`
}

//会员交易记录
type MemberDealRecord struct {
	Id           int64   `xorm:"id" json:"id"`
	SiteId       string  `xorm:"site_id" json:"site_id"`             //操作站点id
	SiteIndexId  string  `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	SourceType   int     `xorm:"source_type" json:"source_type"`     //数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8取消出款
	Balance      float64 `xorm:"balance" json:"balance"`             //金额
	Types        int     `xorm:"type" json:"types"`                  //1.存入2.取出
	DisBalance   float64 `xorm:"dis_balance" json:"dis_balance"`     //优惠金额
	AfterBalance float64 `xorm:"after_balance" json:"after_balance"` //操作后余额
	Remark       string  `xorm:"remark" json:"remark"`               //备注
	CreateTime   int64   `xorm:"create_time" json:"create_time"`     //添加时间
}

//会员出款列表
type MemberBankListBack struct {
	Id          int64  `xorm:"id" json:"id"`
	BankId      int64  `xorm:"bank_id" json:"bank_id"` //卡类型
	Title       string `xorm:"title" json:"title" json:"title"`
	Card        string `xorm:"card" json:"card"`                 //卡号
	CardName    string `xorm:"card_name" json:"card_name"`       //卡账号
	CardAddress string `xorm:"card_address" json:"card_address"` //卡地址
}

//会员出款
type OneMemberBankInfoBack struct {
	Id          int64  `xorm:"id" json:"id"`
	BankId      int64  `xorm:"bank_id" json:"bank_id"`           //卡类型
	Card        string `xorm:"card" json:"card"`                 //卡号
	CardName    string `xorm:"card_name" json:"card_name"`       //卡账号
	CardAddress string `xorm:"card_address" json:"card_address"` //卡地址
}

//会员详细资料
type MemberDetailOne struct {
	MemberId  int64  `xorm:"member_id" json:"member_id"`   //会员id
	Card      string `xorm:"card" json:"card"`             //身份证号
	LocalCode string `xorm:"local_code" json:"local_code"` //区号
	Mobile    string `xorm:"mobile" json:"mobile"`         //手机号码
	Email     string `xorm:"email" json:"email"`           //邮箱
	Qq        string `xorm:"qq" json:"qq"`                 //qq
	Wechat    string `xorm:"wechat" json:"wechat"`         //微信
	Birthday  int64  `xorm:"birthday" json:"birthday"`     //生日
	Remark    string `xorm:"remark" json:"remark"`         //备注
}

//会员登录验证请求
type AjaxLoginIn struct {
	Id       int64                    `xorm:"'id' PK autoincr" json:"id"`
	Account  string                   `xorm:"account" json:"account"`   //登陆账号
	Realname string                   `xorm:"realname" json:"realname"` //真实姓名
	Balance  float64                  `xorm:"balance" json:"balance"`   //账号余额
	Count    int64                    `xorm:"count" json:"count"`       //未读消息数量
	TBalance []MemberBalanceTotalBack //系统余额以及各种账户余额，总余额
}

//会员返佣情况
type MemberRebateInfo struct {
	Account     string              `xorm:"account" json:"account"`           //账号
	SpreadNum   int64               `xorm:"-" json:"spread_num"`              //推广人数
	SpreadMoney float64             `xorm:"spread_money" json:"spread_money"` //推广获利
	SpreadUrl   string              `xorm:"-" json:"spread_url"`              //推广地址,用Host拼接id
	RebateSets  []*MemberRebateRate `xorm:"-" json:"rebate_sets"`             //返佣设定详情
	RankingList RankingList         `xorm:"-" json:"ranking_list"`            //排行榜
}

type RankingList struct {
	RankingNumList   []*RankingNumList   `json:"ranking_num_list" `   //人数排行榜
	RankingMoneyList []*RankingMoneyList `json:"ranking_money_list" ` //金额排行榜
}

type MemberRebateRate struct {
	Title  string    `json:"title" `  //pc端我要推广页面的展示,最左边的一列,有效打码或商品名
	Values []float64 `json:"values" ` //有效打码或比例
}

//推广人数排行榜
type RankingNumList struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
	Num     int64  `json:"num"`
}

//推广金额排行榜
type RankingMoneyList struct {
	Id      int64   `json:"id"`
	Account string  `json:"account"`
	Money   float64 `json:"money"`
}
