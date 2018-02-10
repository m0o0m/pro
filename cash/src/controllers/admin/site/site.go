//站点管理 function引入
package site

import "models/function"

var agencyBean = new(function.AgencyBean)                       //开户人及下属表
var siteOperateBean = new(function.SiteOperateBean)             //站点
var siteControllBean = new(function.SiteControllBean)           //站点管理
var memberLevelBean = new(function.MemberLevelBean)             //会员层级
var noticeBean = new(function.NoticeBean)                       //站点公告
var comboBeen = new(function.ComboBeen)                         //站点管理
var bankCardBean = new(function.BankCardBean)                   //银行卡
var menuBean = new(function.MenuBean)                           //菜单
var ipSetBean = new(function.IpSetBean)                         //ip开关
var siteDownBean = new(function.SiteDownBean)                   //下载管理
var sitePassBean = new(function.SitePassBean)                   //站点口令
var siteProductBean = new(function.SiteProductBean)             //站点商品剔除
var onlineIncomeThirdBean = new(function.OnlineIncomeThirdBean) //第三方入款
var siteJsVersionBean = new(function.SiteJsVersionBean)         //js版本控制
var siteDomainBean = new(function.SiteDomainBean)               //站点域名配置
var maintenanceBean = new(function.MaintenanceBean)             //]维护管理
var siteADBean = new(function.SiteADBean)                       //广告管理
