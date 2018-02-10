package global

import "errors"

var (
	errInvalidMysqlNode = errors.New("config mysql node is nil")
	Version             = "1.000"
)

const (
	TablePrefix = "sales_" //数据表前缀
	DEFAULT     = "default"
	VIDEO       = "video"

	EncryptSalt = "ocpB8nZG5yBWrMfJDsM2fRB5L5LERMF47A6PWAC4wpM="
	DIR         = "dir"
	FILE        = "file"
)

//全局状态值
const (
	STATUS_ENABLE  = 1 //启用
	STATUS_DISABLE = 2 //停用
	//在线状态
	STATUS_ONLINE  = 1 //在线
	STATUS_OFFLINE = 2 //离线
	//登录状态
	LOGIN_SUCCESS = 1 //登录成功
	LOGIN_FAILD   = 2 //登录失败
)

//全局布尔值
const (
	BOOL_TRUE  = 1 //是
	BOOL_FALSE = 2 //否
)

//操作日志表http请求类型值
const (
	REQUEST_POST   = 1 //增加
	REQUEST_DELETE = 2 //删除
	REQUEST_GET    = 3 //查看
	REQUEST_PUT    = 4 //修改
)

//site_role表固定值
const (
	ROLE_ACCOUNT_OPENING = 1 //开户人及开户人子账号角色id
	ROLE_SHAREHOLDER     = 2 //股东角色id
	ROLE_GENERAL_AGENT   = 3 //总代理角色id
	ROLE_AGENT           = 4 //代理及代理子账号角色id
	ROLE_ADMIN           = 5 //总后台超级管理员角色id
)

//客户端类型
const (
	CLIENT_PC      = 1 //pc
	CLIENT_WAP     = 2 //wap
	CLIENT_ANDROID = 3 //android
	CLIENT_IOS     = 4 //ios
)

//sales_ban_ip表ip段黑名单控制类型
const (
	IP_CONTROL_ACCOUNT_OPENING = 1 //客户后台
	IP_CONTROL_AGENT           = 2 //代理后台
	IP_CONTROL_FRONT           = 3 //前台
	IP_CONTROL_WAP             = 4 //wap
)

//site_domain表type类型
const (
	DOMAIN_FRONT  = 1 //前台域名
	DOMAIN_MANAGE = 2 //客户后台域名
	DOMAIN_AGENCY = 3 //代理后台域名
)

//游戏类型
const (
	GAME_VIDEO       = 1 //视讯
	GAME_ELECTRONICS = 2 //电子
	GAME_FISHING     = 3 //捕鱼
	GAME_LOTTERY     = 4 //彩票
	GAME_SPORT       = 5 //体育
)

//入款方式
const (
	PAYMENT_ARTIFICIAL = 1 //人工存入
	PAYMENT_COMPANY    = 2 //公司入款
	PAYMENT_ONLINE     = 3 //线上入款
)

//注册文案类型
const (
	REGISTER_MEMBER = 1 //会员注册
	REGISTER_AGENT  = 2 //代理注册
	REGISTER_TRIAL  = 3 //试玩注册
)

//出款状态
const (
	DISPENSING_ALREADY = 1 //已出款
	DISPENSING_READY   = 2 //预备出款
	DISPENSING_CANCEL  = 3 //取消出款
	DISPENSING_REFUSE  = 4 //拒绝出款
	DISPENSING_PENDING = 5 //待审核
)

//出入款记录存取款类型
const (
	DEPOSIT_ARTIFICIAL          = 1 //人工存入
	DEPOSIT_DISCOUNT            = 2 //存款优惠
	DEPOSIT_OF_ZERO             = 3 //负数额度归零
	DEPOSIT_CANCEL              = 4 //取消出款
	DEPOSIT_REBATE_DISCOUNT     = 5 //返点优惠
	DEPOSIT_ACTIVITIES_DISCOUNT = 6 //活动优惠
)

//审核状态
const (
	AUDIT_READY      = 1 //待审核
	AUDIT_PASSED     = 2 //审核已通过
	AUDIT_NOT_PASSED = 3 //审核未通过
)

//公司入款记录状态
const (
	COMPANY_DEPOSIT_CONFIRM    = 1 //已确认
	COMPANY_DEPOSIT_CANCEL     = 2 //已取消
	COMPANY_DEPOSIT_NOT_REMIND = 3 //不再提醒
	COMPANY_DEPOSIT_PENDING    = 4 //未处理
)

//会员额度转换操作人类型
const (
	CONVERSION_MANAGE = 1 //平台管理员
	CONVERSION_MEMBER = 2 //会员
)

//会员现金流水数据来源
const (
	SOURCE_ARTIFICIAL_INCOME     = 0  //人工存入
	SOURCE_COMPANY_INCOME        = 1  //公司入款
	SOURCE_ONLINE_INCOME         = 2  //线上入款
	SOURCE_ARTIFICIAL_DISPENSING = 3  //人工取款
	SOURCE_ONLINE_DISPENSING     = 4  //人工取款
	SOURCE_OUT                   = 5  //出款
	SOURCE_REGISTER_DISCOUNT     = 6  //注册优惠
	SOURCE_ORDER                 = 7  //下单
	SOURCE_QUOTA_CONVERSION      = 8  //额度转换
	SOURCE_DISCOUNT_WATER        = 9  //优惠返水
	SOURCE_SELF_WATER            = 10 //自助返水
	SOURCE_MEMBER_REBATE         = 11 //会员返佣
	SOURCE_RED_ENVELOPE          = 12 //红包
)

//会员消息类型
const (
	MESSAGE_PERSONAL = 1 //个人消息
	MESSAGE_PLATFORM = 2 //平台消息
)

//菜单等级
const (
	MENU_ONE_LEVEL   = 1 //一级菜单
	MENU_TWO_LEVEL   = 2 //二级菜单
	MENU_THREE_LEVEL = 3 //三级菜单
)

//菜单类型
const (
	MENU_AGENCY = 1 //代理菜单
	MENU_ADMIN  = 2 //平台菜单
)

//线上入款记录状态
const (
	ONLINE_ENTRY_NONE   = 1 //未支付
	ONLINE_ENTRY_OK     = 2 //已支付
	ONLINE_ENTRY_CANCEL = 3 //已取消
)

//支付设定表存款优惠次数
const (
	DEPOSIT_DISCOUNT_FIRST = 1 //每次
	DEPOSIT_DISCOUNT_EVERY = 2 //首次
)

//红包活动状态
const (
	RED_PACKET_NOT_START  = 1 //未开始
	RED_PACKET_PROCESSING = 2 //进行中
	RED_PACKET_END        = 3 //已结束
	RED_PACKET_DELETE     = 4 //已删除
)

//站点状态
const (
	SITE_ENABLE   = 1 //正常
	SITE_DISABLE  = 2 //禁用
	SITE_PAUSE    = 3 //暂停
	SITE_MAINTAIN = 4 //维护
)

//弹窗广告位置
const (
	ADV_MIDDLE     = 1 //中间
	ADV_LEFT_DOWN  = 2 //左下
	ADV_RIGHT_DOWN = 3 //右下
)

//站点额度记录类型
const (
	QUOTA_RECORD_CONVERSION = 1 //额度转换
	QUOTA_RECORD_ADD        = 2 //额度加款
	QUOTA_RECORD_REDUCE     = 3 //额度扣款
	QUOTA_RECORD_READY      = 4 //预借
	QUOTA_RECORD_RECHARGE   = 5 //业主充值
)

//站点额度记录状态
const (
	QUOTA_READY  = 1 //待审核
	QUOTA_NORMAL = 2 //正常
	QUOTA_LOST   = 3 //掉单
)

//公告类型
const (
	NOTICE_ORDINARY    = 1 //普通公告
	NOTICE_SYSTEM      = 2 //系统公告
	NOTICE_MAINTAIN    = 3 //维护公告
	NOTICE_VIDEO       = 4 //视讯公告
	NOTICE_ELECTRONICS = 5 //电子公告
	NOTICE_LOTTERY     = 6 //彩票公告
	NOTICE_SPORT       = 7 //体育公告
)

//后台充值第三方配置类型
const (
	THIRD_PARTY      = 1 //第三方
	THIRD_PARTY_BANK = 3 //银行卡
)

//站点额度购买充值记录支付方式
const (
	QUOTA_RECHARGE_THIRD   = 1 //第三方入款
	QUOTA_RECHARGE_COMPANY = 2 //公司入款
)

//账单下发状态
const (
	REPORT_BILL_NONE   = 1 //未下发
	REPORT_BILL_OK     = 2 //已下发
	REPORT_BILL_DELETE = 3 //删除
)

//视讯额度掉单申请类型
const (
	QUOTA_LOST_SYSTEM = 1 //系统转视讯掉单
	QUOTA_LOST_VIDEO  = 2 //视讯转系统掉单
)

//视讯额度掉单申请状态
const (
	QUOTA_LOST_READY   = 1 //审核中
	QUOTA_LOST_PASSED  = 2 //审核通过
	QUOTA_LOST_INVALID = 3 //无效申请
)
