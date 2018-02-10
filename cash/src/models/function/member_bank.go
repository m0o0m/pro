package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type MemberBankBean struct{}

//会员银行卡列表
func (*MemberBankBean) MemberBankListById(id int64) ([]back.MemberBanksListBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberBank)
	var data []back.MemberBanksListBack
	err := sess.Table(mb.TableName()).
		Where(mb.TableName()+".member_id=?", id).
		Where(mb.TableName()+".delete_time=?", 0).Find(&data)

	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	bank := new(schema.Bank)
	var allBank []schema.Bank
	//获取所有银行
	err = sess.Table(bank.TableName()).
		Where(bank.TableName()+".status=?", 1).
		Select("id").Find(allBank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	ids := make([]int64, 0)
	for _, v := range allBank {
		ids = append(ids, v.Id)
	}
	err = sess.Table(mb.TableName()).
		In("bank_id", ids).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return data, err
}

//添加会员银行卡
func (*MemberBankBean) MemberAddBankCard(this *input.MemberBankAddIn) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberBank)
	mb.MemberId = this.MemberId
	mb.Card = this.CardNumber
	mb.CardAddress = this.CardAddress
	mb.BankId = this.BankId
	mb.IsDefaultBank = 2
	mb.CardName = this.CardMan
	count, err := sess.InsertOne(mb)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询银行是否存在
func (*MemberBankBean) BeMemberBankById(id, member_id int64) (*schema.MemberBank, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberBank)
	has, err := sess.Where("id=?", id).
		Where("member_id=?", member_id).
		Where("delete_time=?", 0).Get(mb)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mb, has, err
	}
	return mb, has, err
}

//解绑
func (*MemberBankBean) MemberBankUnBundling(id, member_id int64, status int8) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberBank)
	if status == 1 {
		mb.IsDefaultBank = 2
	} else if status == 2 {
		mb.IsDefaultBank = 1
	}
	count, err := sess.Where("id=?", id).
		Where("member_id=?", member_id).
		Where("delete_time=?", 0).Cols("is_default_bank").Update(mb)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//会员银行卡删除
func (*MemberBankBean) MemberBankDelete(this *input.MemberBankDeleteIn) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberBank)
	mb.DeleteTime = time.Now().Unix()
	count, err := sess.Where("id=?", this.Id).Where("member_id=?", this.MemberId).
		Cols("delete_time").Update(mb)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//银行下拉框
func (*MemberBankBean) BankDropBySiteAndSiteIndexId(site_id, site_index_id string) ([]back.BankDropBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	out_del := new(schema.BankOutDel)
	//查询所有的允许出款的银行
	var a_b []back.BankDropBack
	err := sess.Table(bank.TableName()).Where("is_out=?", 1).
		Where("delete_time=?", 0).Find(&a_b)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return a_b, err
	}
	//根据会员的站点查询出款银行剔除表中的数据
	var b_d []schema.BankOutDel
	err = sess.Table(out_del.TableName()).Where("site_id=?", site_id).
		Where("site_index_id=?", site_index_id).Find(&b_d)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return a_b, err
	}
	//两个表相比较，将该站点允许使用的银行整理出来
	var data []back.BankDropBack
	var da back.BankDropBack
	if len(b_d) > 0 {
		for k, v := range a_b {
			for _, j := range b_d {
				if v.Id == j.BankId {
					a_b[k].Status = 2
				}
			}
		}
	} else {
		data = a_b
		return data, err
	}
	for _, h := range a_b {
		if h.Status == 1 {
			da.Id = h.Id
			da.Title = h.Title
			da.Status = h.Status
			data = append(data, da)
		}
	}
	return data, err
}

//会员银行详情
func (*MemberBankBean) MemberBankCardDetails(this *input.MemberBankCardDetailsIn) (*back.MemberBankCardDetailsBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	mb := new(schema.MemberBank)
	info := new(back.MemberBankCardDetailsBack)
	has, err := sess.Table(mb.TableName()).
		Join("LEFT", bank.TableName(), mb.TableName()+".bank_id="+
			bank.TableName()+".id").Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err
	}
	return info, has, err
}

//获取单条出款银行卡号
func (*MemberBankBean) GetOneBank(this *input.OneMemberBankInfo) (*back.WapOutBank, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberBank)
	data := new(back.WapOutBank)
	has, err := sess.Table(mb.TableName()).
		Select("`id`,`card`,`card_name`").
		Where("id=?", this.Id).
		Where("member_id=?", this.MemberId).
		Get(data)
	return data, has, err
}

//会员银行卡列表
func (*MemberBankBean) MemberBankList(id int64) (data []back.MemberBanksList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberBank)
	var mData []schema.MemberBank
	bank := new(schema.Bank)
	err = sess.Table(mb.TableName()).Where("member_id=?", id).Find(&mData)
	if err != nil {
		return
	}
	if len(mData) == 0 {
		return
	}
	var notStr []int64
	for _, v := range mData {
		notStr = append(notStr, v.BankId)
	}
	err = sess.Table(bank.TableName()).
		Select("id,title").
		In("id", notStr).
		Where("delete_time=?", 0).
		Where("status=?", 1).
		Find(&data)
	if err != nil {
		return
	}
	for k := range data {
		data[k].Card = mData[k].Card
		data[k].CardName = mData[k].CardName
		data[k].Id = mData[k].Id
	}
	return
}

//查询会员银行是否存在
func (*MemberBankBean) GetMemberBankOne(id int64) (*schema.MemberBank, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberBank)
	has, err := sess.
		Where("member_id=?", id).
		Where("delete_time=?", 0).Get(mb)
	return mb, has, err
}
