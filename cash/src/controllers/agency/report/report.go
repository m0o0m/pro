//[控制器] [代理] 报表模块
package report

import "models/function"

var (
	reportBean      = new(function.ReportBean)      //数据中心
	thirdAgencyBean = new(function.ThirdAgencyBean) //代理信息
	agencyBean      = new(function.AgencyBean)
	memberBean      = new(function.MemberBean)      //会员信息
	productBean     = new(function.ProductBean)     //商品信息
	siteOperateBean = new(function.SiteOperateBean) //商品信息
	webInfoBean     = new(function.WebInfoBean)     //站点信息
	pressMoneyBean  = new(function.PressMoneyBean)  //催款账单管理
)
