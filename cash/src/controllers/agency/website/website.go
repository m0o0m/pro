//[控制器] [代理] 咨询管理模块
package website

import "models/function"

type WebSiteController struct{}

var sitePromotionConfigBean = new(function.SitePromotionConfigBean) //自助优惠申请配置
var siteThumbBean = new(function.SiteThumbBean)                     //附件管理
var webInfoBean = new(function.WebInfoBean)                         //网站资讯系统
var webLogoBean = new(function.WebLogoBean)                         //站点logo图片管理
var webFloatBean = new(function.WebFloatBean)                       //站点左右浮动管理
var webFlashBean = new(function.WebFlashBean)                       //站点轮播图管理
var webAdvBean = new(function.WebAdvBean)                           //站点公告弹窗管理
var webPopBean = new(function.WebPopBean)                           //站点左右浮动管理
var siteIwordBean = new(function.SiteIwordBean)                     //站点文案
