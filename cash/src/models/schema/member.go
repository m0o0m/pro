package schema

import "global"

//会员表
type Member struct {
	Id                 int64   `xorm:"id PK autoincr"`
	SiteId             string  `xorm:"site_id"`               //站点Id
	SiteIndexId        string  `xorm:"site_index_id"`         //站点前台Id
	LevelId            string  `xorm:"level_id"`              //会员所属层级
	Account            string  `xorm:"account"`               //登陆账号
	Password           string  `xorm:"password"`              //登陆密码
	DrawPassword       string  `xorm:"draw_password"`         //取款密码
	Realname           string  `xorm:"realname"`              //真实姓名
	SpreadId           int64   `xorm:"spread_id"`             //这个是该会员的推广人id
	SpreadMoney        float64 `xorm:"spread_money"`          //推广获得佣金
	FirstAgencyId      int64   `xorm:"first_agency_id"`       //一级经销商Id
	SecondAgencyId     int64   `xorm:"second_agency_id"`      //二级经销商Id
	ThirdAgencyId      int64   `xorm:"third_agency_id"`       //所属代理Id
	IsLockedLevel      int8    `xorm:"is_locked_level"`       //是否锁定会员所在层级(1.锁定2.不锁定)
	Status             int8    `xorm:"status"`                //状态(1.正常2.禁用)
	IsEditPassword     int8    `xorm:"is_edit_password"`      //是否可以修改密码
	RegisterClientType int8    `xorm:"register_client_type"`  //注册终端(1.pc2.wap3.ios4.android)
	RegisterIp         string  `xorm:"register_ip"`           //注册ip
	PcLoginKey         string  `xorm:"pc_login_key"`          //pc登陆key
	WapLoginKey        string  `xorm:"wap_login_key"`         //wap登录key
	IosLoginKey        string  `xorm:"ios_login_key"`         //ios端登录key
	AndroidLoginKey    string  `xorm:"android_login_key"`     //安卓端登录key
	LoginCount         int64   `xorm:"login_count"`           //登陆次数
	LoginIp            string  `xorm:"login_ip"`              //登陆Ip
	LoginTime          int64   `xorm:"login_time"`            //登陆时间
	LastLoginIp        string  `xorm:"last_login_ip"`         //上次登陆Ip
	LastLoginTime      int64   `xorm:"last_login_time"`       //上次登录时间
	Balance            float64 `xorm:"balance"`               //账号余额
	FirstDepositTime   int64   `xorm:"first_deposit_time"`    //首次存款时间
	FirstDepositId     int64   `xorm:"first_deposit_id"`      //首次存款订单id
	IsAgreeDeal        int64   `xorm:"is_agree_deal"`         //是否同意注册协议
	CreateTime         int64   `xorm:"'create_time' created"` //创建时间
	DeleteTime         int64   `xorm:"delete_time"`           //软删除时间
	PcStatus           int8    `xorm:"pc_status"`             //pc端登录状态
	WapStatus          int8    `xorm:"wap_status"`            //wap端登录状态
	IosStatus          int8    `xorm:"ios_status"`            //ios端登录状态
	AndroidStatus      int8    `xorm:"android_status"`        //安卓端登录状态
	Remark             string  `xorm:"remark"`                //备注
	IdHide             int8    `xorm:"id_hide"`               //是否隐藏1是2否
	IsImport           int8    `xorm:"is_import"`             //是否导入会员 0.注册  1.导入
}

func (*Member) TableName() string {
	return global.TablePrefix + "member"
}
