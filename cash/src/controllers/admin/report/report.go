//[控制器] [平台] 财务报表模块
package report

import (
	"models/function"
)

//财务报表 - 额度计算
var (
	siteLevelBean          = new(function.SiteLevelBean)          //站点层级管理
	pressMoneyBean         = new(function.PressMoneyBean)         //催款账单管理
	quotaBean              = new(function.QuotaCountBean)         //额度统计
	retreatWaterRecordBean = new(function.RetreatWaterRecordBean) //优惠统计
	discountCountBean      = new(function.DiscountCountBean)      //优惠统计
	cashCountBean          = new(function.CashCountBean)          //入款统计
	reportBean             = new(function.ReportBean)             //账单查询
	manualAccessBean       = new(function.ManualAccessBean)       //人工出入款记录
	cashRecordBean         = new(function.MemberCashRecordBean)   //现金记录
	makeMoneyBean          = new(function.MakeMoneyBean)          //会员出款
	sitePackageBean        = new(function.SitePackageBean)        //套餐管理
	siteOperateBean        = new(function.SiteOperateBean)        //站点管理
	productBean            = new(function.ProductBean)            //站点管理
	comboBeen              = new(function.ComboBeen)              //套餐列表
	memberCashRecordBean   = new(function.MemberCashRecordBean)   //现金记录
	reportFormBean         = new(function.ReportFormBean)         //报表统计
	agencyBean             = new(function.AgencyBean)             //代理
	memberBean             = new(function.MemberBean)             //会员
)
