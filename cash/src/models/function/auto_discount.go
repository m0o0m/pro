package function

import (
	"time"

	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SelfHelpApplyforBean struct{}

//自助优惠申请列表
func (*SelfHelpApplyforBean) SelfApplyList(inputRequest *input.DiscountList, listparam *global.ListParams) (backData []back.AutoDiscountBack, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	maa := new(schema.MemberAutoApplypro)
	//站点
	if inputRequest.SiteId != "" {
		sess.Where("site_id=?", inputRequest.SiteId)
	}
	//前台id
	if inputRequest.SiteIndexId != "" {
		sess.Where("site_index_id=?", inputRequest.SiteIndexId)
	}
	//状态
	if inputRequest.Status != 0 {
		sess.Where("status=?", inputRequest.Status)
	}
	//账号
	if inputRequest.ApplyAccount != "" {
		sess.Where("account=?", inputRequest.ApplyAccount)
	}
	//开始时间
	if inputRequest.StartDate != "" {
		startDate, err := time.Parse("2006-01-02 15:04:05", inputRequest.StartDate)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return backData, 0, err
		}
		sess.Where("apply_time>?", startDate.Unix())
	}
	//结束时间
	if inputRequest.EndDate != "" {
		endDate, err := time.Parse("2006-01-02 15:04:05", inputRequest.EndDate)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return backData, 0, err
		}
		sess.Where("apply_time<?", endDate.Unix())
	}
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	//获取数据
	err = sess.Table(maa.TableName()).Find(&backData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, 0, err
	}
	//获得符合条件的记录数
	count, err = sess.Table(maa.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, count, err
	}
	return backData, count, err
}

//自主优惠状态修改
func (*SelfHelpApplyforBean) ChangeStatusDiscount(this *input.AutoDiscountStatus) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	site.SelfHelpSwitch = this.Status
	count, err = sess.Table(site.TableName()).Where("id=?", this.SiteId).
		Where("index_id=?", this.SiteIndexId).
		Where("delete_time=?", 0).
		Cols("self_help_switch").
		Update(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取自助优惠详情
func (*SelfHelpApplyforBean) GetAutoDiscountInfo(id int64) (back.MemberAutoApplypro, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	maa := new(schema.MemberAutoApplypro)
	var backData back.MemberAutoApplypro
	flag, err := sess.Table(maa.TableName()).Where("id=?", id).Get(&backData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, flag, err
	}
	return backData, flag, err
}

//拒绝一条优惠申请
func (*SelfHelpApplyforBean) RefusedFor(inputData *input.RefuseApplyFor, id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	selfApply := new(schema.SelfHelpApplyfor)
	selfApply.RejectReason = inputData.Reason
	selfApply.Status = 2
	selfApply.AuditTime = time.Now().Unix()
	selfApply.OperateId = id
	count, err := sess.Where("id=?", inputData.Id).
		Cols("status,reject_reason,audit_time,operate_id").
		Update(selfApply)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//自主优惠状态添加
func (*SelfHelpApplyforBean) AutoDiscountAdd(this *input.SelfHelpApllyAdd, memberInfo *global.MemberRedisToken) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//selfHelpApplyfor := new(schema.SelfHelpApplyfor)
	selfHelpApplyfor := new(schema.MemberAutoApplypro)
	selfHelpApplyfor.Createtime = time.Now().Unix()
	selfHelpApplyfor.Status = 3
	selfHelpApplyfor.Account = this.Account
	selfHelpApplyfor.SiteIndexId = memberInfo.SiteIndex
	selfHelpApplyfor.SiteId = memberInfo.Site
	selfHelpApplyfor.ApplyMoney = this.ApplyMoney
	selfHelpApplyfor.Applyreason = this.Applyreason
	selfHelpApplyfor.PromotionId = this.ProId
	selfHelpApplyfor.PromotionContent = this.ProContent
	selfHelpApplyfor.PromotionTitle = this.ProTitle
	selfHelpApplyfor.MemberId = memberInfo.Id
	count, err = sess.Table(selfHelpApplyfor.TableName()).InsertOne(selfHelpApplyfor)
	return
}

//获取会员优惠申请列表
func (*SelfHelpApplyforBean) GetMemberApplyList(memberId int64, siteId, siteIndexId string, orderNum int64, times *global.Times, listparams *global.ListParams) (data []back.MemberAutoApplypro, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	applyData := new(schema.MemberAutoApplypro)
	sess.Where("member_id=?", memberId)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	if orderNum != 0 {
		sess.Where("promotion_id=?", orderNum)
	}
	if times != nil {
		times.Make("createtime", sess)
	}
	sess.OrderBy("id desc")
	conds := sess.Conds()
	listparams.Make(sess)
	err = sess.Table(applyData.TableName()).Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(applyData.TableName()).Where(conds).Count()
	return
}
