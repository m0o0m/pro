//[控制器] [代理] 账号管理模块注释
package account

import "models/function"

var agencyDomainBeen = new(function.AgencyDomainBeen)
var distributionApplyBeen = new(function.DistributionApplyBeen)                             //站点代理申请
var secondDistributionRegisterSetupBeen = new(function.SecondDistributionRegisterSetupBeen) //站点代理申请注册设定
var agencySignBean = new(function.AgencySignBean)                                           //代理登录
var agencyCountBean = new(function.AgencyCountBean)                                         //代理统计
var agencyMemberRegisterDiscountSetBean = new(function.AgencyMemberRegisterDiscountSetBean) //会员注册优惠设定
var memberBean = new(function.MemberBean)                                                   //会员
var memberRegisterSettingBean = new(function.MemberRegisterSettingBean)                     //会员注册设定
var memberLevelBean = new(function.MemberLevelBean)                                         //会员层级
var agencyBean = new(function.AgencyBean)                                                   //股东、总代的下拉框
var subAccountBeen = new(function.SubAccountBeen)                                           //子账号
var agencyThirdInfoBean = new(function.AgencyThirdInfoBean)                                 //代理详细资料
var roleMenuBean = new(function.RoleMenuBean)                                               //角色菜单中间表
var firstAgencyBean = new(function.FirstAgencyBean)                                         //股东
var secondAgencyBean = new(function.SecondAgencyBean)                                       //总代
var thirdAgencyBean = new(function.ThirdAgencyBean)                                         //代理
var permissionBean = new(function.PermissionBean)                                           //权限
var siteOperateBean = new(function.SiteOperateBean)                                         //站点
var bankCardBean = new(function.BankCardBean)                                               //银行
var ipSetBean = new(function.IpSetBean)                                                     //ip控制
