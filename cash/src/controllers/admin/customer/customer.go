//[控制器] [平台] 客户模块
package customer

import (
	"models/function"
)

var siteOperateBean = new(function.SiteOperateBean)
var auditsBean = new(function.AuditsBean)                     //稽核日志
var siteLogBean = new(function.SiteLogBean)                   //管理员日志
var memberBean = new(function.MemberBean)                     //会员
var memberLevelBean = new(function.MemberLevelBean)           //会员层级
var bankInOutBean = new(function.BankInOutBean)               //银行出入款
var childAccountBean = new(function.ChildAccountBean)         //子账号
var selfHelpApplyforBean = new(function.SelfHelpApplyforBean) //自助优惠申请
var discountSearchBean = new(function.DiscountSearchBean)     //优惠查询
var noticePopupBean = new(function.NoticePopupBean)           //站点公告
var videoMemberBean = new(function.VideoMemberBean)           //视讯类型
var agencyBean = new(function.AgencyBean)                     //代理管理
var abnormalMemberBean = new(function.AbnormalMemberBean)     //异常会员
var redBagBean = new(function.RedBagBean)                     //红包补数据
var drawMoney = new(function.DrawMoneyBean)                   //取款管理
