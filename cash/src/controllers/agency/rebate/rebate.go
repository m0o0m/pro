package rebate

import (
	"models/function"
)

var rebateSetBean = &function.MemberRebateSetBean{}                       //返佣设置
var spreadBean = new(function.MemberSpreadBean)                           //会员推广
var rebateBean = new(function.MemberRebateBean)                           //返佣查询
var rebateRecordBean = new(function.MemberRebateRecordBean)               //返佣记录
var rebateRecordProductBean = new(function.MemberRebateRecordProductBean) //返佣记录商品对应价格
var productBean = new(function.ProductBean)                               //商品
var betReportBean = new(function.BetReportBean)                           //会员每日打码统计
var memberBean = new(function.MemberBean)                                 //会员
var auditBean = new(function.MemberAuditBean)                             //会员稽核
var auditLogBean = new(function.MemeberAuditLogBean)                      //会员稽核日志
var cashRecordBean = new(function.MemberCashRecordBean)                   //现金记录
