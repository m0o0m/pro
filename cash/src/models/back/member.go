package back

//返回会员列表
type MemberIndex struct {
	Id                 int64                `xorm:"id" json:"id"`                                   //会员id
	SiteId             string               `xorm:"site_id" json:"siteId"`                          //站点id
	SiteIndexId        string               `xorm:"site_index_id" json:"siteIndexId"`               //前台id
	LevelId            string               `xorm:"level_id" json:"levelId"`                        //层级id
	Account            string               `xorm:"account" json:"account"`                         //账号
	Realname           string               `xorm:"realname" json:"realname"`                       //姓名
	Balance            float64              `xorm:"balance" json:"balance"`                         //余额
	AgencyId           int64                `xorm:"third_agency_id" json:"agencyId"`                //所属代理id
	AgencyAccount      string               `xorm:"account" json:"agencyAccount"`                   //所属代理账号
	AgencyName         string               `xorm:"username" json:"agencyName"`                     //所属代理名称
	CreateTime         string               `xorm:"create_time" json:"createTime"`                  //会员注册时间
	Status             int                  `xorm:"status" json:"status"`                           //账号状态
	PcStatus           int8                 `xorm:"pc_status" json:"pcStatus"`                      //pc端在线状态
	WapStatus          int8                 `xorm:"wap_status" json:"wapStatus"`                    //wap端在线状态
	IosStatus          int8                 `xorm:"ios_status" json:"iosStatus"`                    //ios端在线状态
	AndroidStatus      int8                 `xorm:"android_status" json:"androidStatus"`            //安卓端在线状态
	MemberVideoBalance []MemberVideoBalance `xorm:"member_video_balance" json:"memberVideoBalance"` //会员视讯余额
}

//会员视讯余额
type MemberVideoBalance struct {
	MemberId   int64   `json:"memberId"`
	Platform   string  `json:"platform"`   //视讯平台名称
	PlatformId int64   `json:"platformId"` //视讯平台id
	Balance    float64 `json:"balance"`    //额度
}

//会员总数/今日注册人数
type MemberNumberBySite struct {
	TotalNum int64 `xorm:"-" json:"totalNum"` //总人数
	RegNum   int64 `xorm:"-" json:"regNum"`   //注册人数
}

//返回会员基本资料
type MemberInfo struct {
	Id             int64  `xorm:"id PK" json:"id"`                        //会员id
	Account        string `xorm:"account" json:"account"`                 //会员账号
	IsEditPassword int8   `xorm:"is_edit_password" json:"isEditPassword"` //是否可以修改密码
	Realname       string `xorm:"realname" json:"realname"`               //会员姓名
	//DrawPassword   string `xorm:"draw_password"json:"drawpassword"`      //取款密码
}

//返回会员详细资料
type MemberDetail struct {
	Id           int64        `xorm:"id PK" json:"id"`                   //会员id
	Account      string       `xorm:"account" json:"account"`            //会员账号
	Realname     string       `xorm:"realname" json:"realname"`          //会员姓名
	Birthday     int64        `xorm:"birthday" json:"birthday"`          //会员生日
	Card         string       `xorm:"card" json:"card"`                  //身份证号
	Mobile       string       `xorm:"mobile" json:"mobile"`              //手机
	Email        string       `xorm:"email" json:"email"`                //邮箱
	QQ           string       `xorm:"qq" json:"qq"`                      //qq
	Wechat       string       `xorm:"wechat" json:"wechat"`              //微信
	RemarkM      string       `xorm:"remark" json:"remarkM"`             //备注
	Remark       string       `xorm:"remark" json:"remark"`              //备注
	CreateTime   string       `xorm:"create_time" json:"createTime"`     //注册时间
	RegisterIp   string       `xorm:"register_ip" json:"registerIp"`     //注册ip
	LoginTime    string       `xorm:"login_time" json:"loginTime"`       //最后一次登录时间
	LoginIp      string       `xorm:"login_ip" json:"loginIp"`           //最后一次登录ip
	DrawPassword string       `xorm:"draw_password" json:"drawPassword"` //取款密码
	MemberBank   []MemberBank `xorm:"-" json:"memberBank"`               //会员银行卡信息
}

//返回会员出款银行卡集合
type MemberBank struct {
	Id          int64  `xorm:"id PK" json:"id"`                 //银行卡列表主键id
	BankId      int64  `xorm:"bank_id" json:"bankId"`           //银行类型id
	Card        string `xorm:"card" json:"card"`                //卡号
	CardName    string `xorm:"card_name" json:"cardName"`       //账号
	CardAddress string `xorm:"card_address" json:"cardAddress"` //开户行
	CreateTime  string `xorm:"create_time" json:"createTime"`   //创建时间
	Title       string `xorm:"title" json:"title"`              //银行名称
}

//根据会员账号查看会员账号,余额,姓名(资金管理)
type MemberInfos struct {
	Account  string  `json:"account"`  //会员账号
	Balance  float64 `json:"balance"`  //余额
	Realname string  `json:"realname"` //姓名
}

type Members struct {
	SiteId        string  `json:"site_id"`        //站点id
	SiteIndexId   string  `json:"site_index_id"`  //站点前台id
	MemberId      int64   `json:"member_id"`      //会员id
	MemberAccount string  `json:"member_account"` //会员账号
	AgencyId      int64   `json:"agency_id"`      //代理id
	Balance       float64 `json:"balance"`        //余额
	AgencyAccount string  `json:"agency_account"` //代理账号
}

//返回支付密码
type DrawPassData struct {
	DrawPassword string `json:"draw_password" xorm:"draw_password"` //支付密码
}

//出款页面返回会员数据
type MemberBankSiteSet struct {
	BankList []MemberBanksList `json:"bank_list"` //会员出款银行列表
	SiteSet  *SitePaySet       `json:"site_set"`  //会员所在站点出款设置
	Poundage *Poundage         `json:"poundage"`  //出款手续费
	RealName *MemberInfo       `json:"real_name"` //真实姓名
}

//会员排序下拉
type MemberSortDrop struct {
	Id       int64  `json:"id"`
	Platform string `json:"platform"` //视讯名称
}
