package function

import (
	"errors"
	"github.com/go-xorm/xorm"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type RedPacketSetBean struct {
}

//添加设置
func (*RedPacketSetBean) Add(add *input.RedPacketSetAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketSetSchema := new(schema.RedPacketSet)
	redPacketSetSchema.Title = add.Title
	redPacketSetSchema.SiteId = add.SiteId
	redPacketSetSchema.SiteIndexId = add.SiteIndexId
	redPacketSetSchema.Description = add.Description
	redPacketSetSchema.MaxCount = add.MaxCount
	redPacketSetSchema.StartTime = add.StartTimestamp
	redPacketSetSchema.EndTime = add.EndTimestamp
	redPacketSetSchema.InStartTime = add.InStartTimestamp
	redPacketSetSchema.InEndTime = add.InEndTimestamp
	redPacketSetSchema.InSum = add.InSum
	redPacketSetSchema.AuditStartTime = add.AuditStartTimestamp
	redPacketSetSchema.AuditEndTime = add.AuditEndTimestamp
	redPacketSetSchema.BetSum = add.BetSum
	redPacketSetSchema.EndTitle = add.EndTitle
	redPacketSetSchema.EndDescription = add.EndDescription
	redPacketSetSchema.LevelId = add.LevelId
	redPacketSetSchema.TotalMoney = add.TotalMoney
	redPacketSetSchema.MinMoney = add.MinMoney
	redPacketSetSchema.RedNum = add.RedNum
	redPacketSetSchema.CreateIp = add.CreateIp
	redPacketSetSchema.CreateUid = add.CreateUid
	redPacketSetSchema.CreateTime = global.GetCurrentTime()
	redPacketSetSchema.IsIp = add.IsIp
	redPacketSetSchema.StyleId = add.StyleId
	redPacketSetSchema.IsShow = add.IsShow
	redPacketSetSchema.AppointMoney = add.AppointMoney
	redPacketSetSchema.RedType = add.RedType
	redPacketSetSchema.Status = add.Status
	redPacketSetSchema.DepositAchieve = add.DepositAchieve
	redPacketSetSchema.ReceiveAgain = add.ReceiveAgain
	redPacketSetSchema.Status = 1
	redPacketSetSchema.IsGenerate = 1
	num, err := sess.InsertOne(redPacketSetSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return num, err
	}
	return num, err
}

//添加设置
func (*RedPacketSetBean) Change(this *input.RedPacketSetChange) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketSetSchema := new(schema.RedPacketSet)
	redPacketSetSchema.Title = this.Title
	redPacketSetSchema.SiteId = this.SiteId
	redPacketSetSchema.SiteIndexId = this.SiteIndexId
	redPacketSetSchema.Description = this.Description
	redPacketSetSchema.MaxCount = this.MaxCount
	redPacketSetSchema.StartTime = this.StartTimestamp
	redPacketSetSchema.EndTime = this.EndTimestamp
	redPacketSetSchema.InStartTime = this.InStartTimestamp
	redPacketSetSchema.InEndTime = this.InEndTimestamp
	redPacketSetSchema.InSum = this.InSum
	redPacketSetSchema.AuditStartTime = this.AuditStartTimestamp
	redPacketSetSchema.AuditEndTime = this.AuditEndTimestamp
	redPacketSetSchema.BetSum = this.BetSum
	redPacketSetSchema.EndTitle = this.EndTitle
	redPacketSetSchema.EndDescription = this.EndDescription
	redPacketSetSchema.LevelId = this.LevelId
	redPacketSetSchema.TotalMoney = this.TotalMoney
	redPacketSetSchema.MinMoney = this.MinMoney
	redPacketSetSchema.RedNum = this.RedNum
	redPacketSetSchema.CreateIp = this.CreateIp
	redPacketSetSchema.CreateUid = this.CreateUid
	redPacketSetSchema.CreateTime = global.GetCurrentTime()
	redPacketSetSchema.IsIp = this.IsIp
	redPacketSetSchema.StyleId = this.StyleId
	redPacketSetSchema.IsShow = this.IsShow
	redPacketSetSchema.AppointMoney = this.AppointMoney
	redPacketSetSchema.RedType = this.RedType
	redPacketSetSchema.Status = this.Status
	redPacketSetSchema.DepositAchieve = this.DepositAchieve
	redPacketSetSchema.ReceiveAgain = this.ReceiveAgain
	redPacketSetSchema.Status = 1
	redPacketSetSchema.IsGenerate = 1
	num, err := sess.
		Where("id=?", this.Id).
		Cols("title,site_id,site_index_id,description,max_count,start_time,end_time,in_start_time," +
			"in_end_time,audit_start_time,audit_end_time,bet_sum,end_title,end_description," +
			"level_id,total_money,min_money,red_num,create_ip,create_uid,create_time,is_ip,style_id," +
			"is_show,appoint_money,red_type,status,is_generate,deposit_achieve,receive_again").
		Update(redPacketSetSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return num, err
	}
	return num, err
}

//红包终止
func (*RedPacketSetBean) Delete(this *input.RedPacketSetDelete) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rD := new(schema.RedPacketSet)
	rD.Status = 4
	count, err := sess.Where("id=?", this.Id).Cols("status").Update(rD)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//得到一个
func (*RedPacketSetBean) GetOne(id int64) (schema.RedPacketSet, error) {
	sess := global.GetXorm().NewSession()
	var redPacketSet schema.RedPacketSet
	b, err := sess.ID(id).Get(&redPacketSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return redPacketSet, err
	}
	if !b {
		err = errors.New("read null")
	}
	return redPacketSet, err
}

//查询列表
func (*RedPacketSetBean) FindList(siteId, siteIndexId string, sTime, eTime, status int64) (
	[]*back.RedPacketSetList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var redPackets []*back.RedPacketSetList
	redPacketSetSchema := new(schema.RedPacketSet)
	sess.Where("site_id = ?", siteId)
	sess.Where("site_index_id = ?", siteIndexId)
	if sTime != 0 {
		sess.Where("create_time >= ?", sTime)
	}
	if eTime != 0 {
		sess.Where("create_time <= ?", eTime)
	}
	if status != 0 {
		sess.Where("status = ?", status)
	}
	err := sess.Table(redPacketSetSchema.TableName()).Find(&redPackets)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return redPackets, err
	}
	return redPackets, err
}

//查询红包设置详情
func (*RedPacketSetBean) FindListInfo(this *input.RedPacketSetListInfo) (*back.RedBagInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rd := new(schema.RedPacketSet)
	rS := new(schema.RedPacketStyle)
	data := new(back.RedBagInfo)
	has, err := sess.Table(rd.TableName()).
		Join("LEFT", rS.TableName(), rd.TableName()+".style_id="+
			rS.TableName()+".id").
		Where(rd.TableName()+".id=?", this.Id).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//查询当前红包活动
func (*RedPacketSetBean) SiteFind(siteId, siteIndexId string) (redPackets []schema.RedPacketSet, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketSetSchema := new(schema.RedPacketSet)
	sess.Where("site_id = ?", siteId)
	sess.Where("site_index_id = ?", siteIndexId)
	sess.Where("start_time <= ?", time.Now().Unix())
	sess.Where("end_time >= ?", time.Now().Unix())
	sess.Where("status >= ?", 1)
	sess.Where("is_generate = ?", 2)

	err = sess.Table(redPacketSetSchema.TableName()).Find(&redPackets)
	return
}

//修改为已生成
func (*RedPacketSetBean) SetGenerate(id int64, sessArgs ...*xorm.Session) (num int64, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	redPacketSetSchema := new(schema.RedPacketSet)
	redPacketSetSchema.IsGenerate = 2
	return sess.ID(id).Update(redPacketSetSchema)
}

//查询该IP是否已经抢过红包
func (*RedPacketSetBean) IpRed(setId int64, ip string) (b int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketLog := new(schema.RedPacketLog)
	sess.Where("set_id = ?", setId)
	//sess.Where("set_id = ?",setId)
	sess.Where("make_sure = ?", 2)
	b, err = sess.Table(redPacketLog.TableName()).Count()
	return
}

//查询该会员抢的红包数量
func (*RedPacketSetBean) UserRed(setId, memberId int64) (b int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketLog := new(schema.RedPacketLog)
	sess.Where("set_id = ?", setId)
	sess.Where("member_id = ?", memberId)
	sess.Where("make_sure = ?", 2)
	b, err = sess.Table(redPacketLog.TableName()).Count()
	return
}

//查询可抢的红包
func (*RedPacketSetBean) GetRebInfo(setId int64) (b bool, redPacketLog schema.RedPacketLog, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("set_id = ?", setId)
	sess.Where("make_sure = ?", 1)
	sess.Limit(1)
	b, err = sess.Table(redPacketLog.TableName()).Get(&redPacketLog)
	return
}

//修改红包状态
func (*RedPacketSetBean) SetRebMakeSure(data schema.RedPacketLog, sessArgs ...*xorm.Session) (num int64, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	data.MakeSure = 2
	return sess.Table(data.TableName()).ID(data.Id).Update(data)
}

//将红包金额加入会员余额
func (*RedPacketSetBean) SetRebMemBalance(id int64, money float64, siteId string, sessArgs ...*xorm.Session) (code, num int64, newCashRecord schema.MemberCashRecord, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}

	Menber := new(schema.Member)
	has, err := sess.Table(Menber.TableName()).
		Where("id = ? and site_id = ?", id, siteId).Get(Menber)
	if !has {
		code = 90716
		return
	}
	if err != nil {
		return
	}
	Menber.Balance = Menber.Balance + money
	num, err = sess.Table(Menber.TableName()).ID(Menber.Id).Update(Menber)

	//会员现金流水数据
	newCashRecord.SiteIndexId = Menber.SiteIndexId
	newCashRecord.SiteId = Menber.SiteId
	newCashRecord.MemberId = Menber.Id
	newCashRecord.UserName = Menber.Account
	newCashRecord.AgencyId = Menber.ThirdAgencyId
	//newCashRecord.AgencyAccount = info.AgencyAccount
	newCashRecord.SourceType = 12
	newCashRecord.TradeNo = ""
	newCashRecord.Type = 1
	newCashRecord.Balance = money               //操作金额
	newCashRecord.AfterBalance = Menber.Balance //操作后的余额
	newCashRecord.Remark = "红包"
	newCashRecord.ClientType = 2
	newCashRecord.CreateTime = time.Now().Unix()
	return
}

//查看红包
func (*RedPacketSetBean) RedBagSeeById(this *input.RedBagSee) ([]back.RedPacketLogInfoBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bg := new(schema.RedPacketLog)
	var data []back.RedPacketLogInfoBack
	err := sess.Table(bg.TableName()).Where("set_id=?", this.Id).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
