package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
	"time"
)

type BaseInfoBean struct{}

//会员个人资料
func (*BaseInfoBean) MemberSelfInfo(this *input.MemberInfoSelf) (data *back.MemberInfoSelfBack, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.Member)
	bank := new(schema.MemberBank)
	minfo := new(schema.MemberInfo)
	if this.Id != 0 {
		sess.Where(mb.TableName()+".id=?", this.Id)
	}
	if this.SiteId != "" {
		sess.Where(mb.TableName()+".site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where(mb.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	sess.Where(mb.TableName()+".delete_time=?", 0)
	data = new(back.MemberInfoSelfBack)
	where1 := fmt.Sprintf("%s.id = %s.member_id", mb.TableName(), minfo.TableName())
	where2 := fmt.Sprintf("%s.id=%s.member_id", mb.TableName(), bank.TableName())
	has, err = sess.Table(mb.TableName()).Join("LEFT", minfo.TableName(), where1).
		Join("LEFT", bank.TableName(), where2).
		Get(data)
	return
}

//
////查询会员的余额
//func (*BaseInfoBean) MemberBalanceAll(this *input.MemberInfoSelf) (data []back.MemberBalanceTotalBack, err error) {
//	sess := global.GetXorm().NewSession()
//	defer sess.Close()
//	mb := new(schema.Member)
//	if this.Id != 0 {
//		sess.Where(mb.TableName()+".id=?", this.Id)
//	}
//	if this.SiteId != "" {
//		sess.Where(mb.TableName()+".site_id=?", this.SiteId)
//	}
//	if this.SiteIndexId != "" {
//		sess.Where(mb.TableName()+".site_index_id=?", this.SiteIndexId)
//	}
//	mpcb := new(schema.MemberProductClassifyBalance)
//	pl := new(schema.Platform)
//	var mbalance []back.MemberBalanceBack
//	where1 := fmt.Sprintf("%s.id = %s.member_id", mb.TableName(), mpcb.TableName())
//	where2 := fmt.Sprintf("%s.platform_id = %s.id", mpcb.TableName(), pl.TableName())
//	err = sess.Table(mb.TableName()).Join("LEFT", mpcb.TableName(), where1).
//		Join("LEFT", pl.TableName(), where2).Find(&mbalance)
//	var da back.MemberBalanceTotalBack
//	var dad float64
//	if len(mbalance) > 0 {
//		for _, v := range mbalance {
//			da.Type = 1
//			da.Balance = v.OthersBalance
//			da.Name = v.Platform
//			data = append(data, da)
//			dad = dad + v.OthersBalance
//		}
//		da.Balance = dad + mbalance[0].Balance
//		da.Type = 3
//		da.Name = "账户总余额"
//		data = append(data, da)
//	}
//	da.Name = "账户余额"
//	da.Balance = mbalance[0].Balance
//	da.Type = 2
//	data = append(data, da)
//	return
//}
//查询会员的余额
//func (*BaseInfoBean) MemberBalanceAll(this *input.MemberInfoSelf) (data []back.MemberBalanceTotalBack, err error) {
//	sess := global.GetXorm().NewSession()
//	defer sess.Close()
//	//获取站点可用平台
//	siteOrderModuleSchema := new(schema.SiteOrderModule)
//	b, err := sess.Where("site_id = ?", this.SiteId).
//		Where("site_index_id = ?", this.SiteIndexId).
//		Get(siteOrderModuleSchema)
//	if err != nil {
//		return data, err
//	}
//	if !(b) {
//		return data, errors.New("get 0 row")
//	}
//	moduleArr := strings.Split(siteOrderModuleSchema.Module, ",")
//	moduleInfoArr := []string{}
//	item := make(map[string]bool)
//	for _, v := range moduleArr {
//		//str := siteModelArr[]
//		info := []string{}
//		if v == "video_module" {
//			info = strings.Split(siteOrderModuleSchema.VideoModule, ",")
//		}
//		if v == "fc_module" {
//			info = strings.Split(siteOrderModuleSchema.FcModule, ",")
//			info = strings.Split(siteOrderModuleSchema.VideoModule, ",")
//			for k1, v1 := range info {
//				info[k1] = strings.Split(v1, "_")[0]
//			}
//		}
//		if v == "dz_module" {
//			info = strings.Split(siteOrderModuleSchema.DzModule, ",")
//			info = strings.Split(siteOrderModuleSchema.VideoModule, ",")
//			for k1, v1 := range info {
//				info[k1] = strings.Split(v1, "_")[0]
//			}
//		}
//		if v == "sp_module" {
//			info = strings.Split(siteOrderModuleSchema.SpModule, ",")
//			info = strings.Split(siteOrderModuleSchema.VideoModule, ",")
//			for k1, v1 := range info {
//				info[k1] = strings.Split(v1, "_")[0]
//			}
//		}
//		for _, val := range info {
//			if item[val] == false {
//				moduleInfoArr = append(moduleInfoArr, val)
//				item[val] = true
//			}
//		}
//	}
//
//	//获取相关会员余额
//	mb := new(schema.Member)
//	b, err = sess.Where("site_id = ?", this.SiteId).
//		Where("site_index_id = ?", this.SiteIndexId).
//		Where("id = ?", this.Id).
//		Get(mb)
//	if err != nil {
//		return data, err
//	}
//	if !(b) {
//		return data, errors.New("get 0 row")
//	}
//	//获取余额信息列表
//	mpcb := new(schema.MemberProductClassifyBalance)
//	pl := new(schema.Platform)
//	if this.Id != 0 {
//		sess.Where(mpcb.TableName()+".member_id=?", this.Id)
//	}
//	if len(moduleInfoArr) > 0 {
//		sess.In(pl.TableName()+".platform", moduleInfoArr)
//	}
//	sess.Where(pl.TableName()+".status=?", 1)
//	sess.Where(pl.TableName()+".delete_time=?", 0)
//	var mbalance []back.MemberBalanceBack
//	where2 := fmt.Sprintf("%s.platform_id = %s.id", mpcb.TableName(), pl.TableName())
//	err = sess.Table(mpcb.TableName()).
//		Join("LEFT", pl.TableName(), where2).Find(&mbalance)
//	var da back.MemberBalanceTotalBack
//	var dad float64
//	for _, v := range moduleInfoArr {
//		da.Type = 1
//		if len(mbalance) > 0 {
//			for _, val := range mbalance {
//				if val.Platform == v {
//					da.Balance = val.Balance
//				}
//			}
//		}
//		da.Name = v
//		data = append(data, da)
//		dad = dad + da.Balance
//	}
//	da.Balance = dad + mb.Balance
//	da.Type = 3
//	da.Name = "账户总余额"
//	data = append(data, da)
//
//	da.Name = "账户余额"
//	da.Balance = mb.Balance
//	da.Type = 2
//	data = append(data, da)
//	return
//}

//获取交易记录
func (*BaseInfoBean) DealRecordMemberSelf(this *input.PayRecordToday, listparam *global.ListParams, times *global.Times) (data []back.MemberDealRecord, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id != 0 {
		sess.Where("member_id=?", this.Id)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	//根据时间段查询
	times.Make("create_time", sess)
	conds := sess.Conds()
	listparam.Make(sess)
	mcr := new(schema.MemberCashRecord)
	err = sess.Table(mcr.TableName()).Find(&data)
	count, err = sess.Table(mcr.TableName()).Where(conds).Count()
	return
}

//获取交易记录
func (*BaseInfoBean) GetDealRecordMemberSelf(this *input.PayRecordToday, times *global.Times) (data []back.MemberDealRecord, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id != 0 {
		sess.Where("member_id=?", this.Id)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Desc("id")
	sess.Limit(10)
	//根据时间段查询
	times.Make("create_time", sess)
	mcr := new(schema.MemberCashRecord)
	err = sess.Table(mcr.TableName()).Find(&data)
	return
}

//查询会员
func (*BaseInfoBean) OneMemberInfoForPassword(this *input.MemberPassword) (data *schema.Member, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	data = new(schema.Member)
	has, err = sess.Table(data.TableName()).Where("delete_time=?", 0).Get(data)
	return
}

//更新密码
func (*BaseInfoBean) UpMemberPassword(this *input.MemberPassword) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.Member)
	if this.Type == 1 {
		mb.Password = this.Password
		count, err = sess.Table(mb.TableName()).Where("delete_time=?", 0).Where("id=?", this.Id).Cols("password").Update(mb)
	}
	if this.Type == 2 {
		mb.DrawPassword = this.Password
		count, err = sess.Table(mb.TableName()).Where("delete_time=?", 0).Where("id=?", this.Id).Cols("draw_password").Update(mb)
	}
	return
}

//添加会员出款银行
func (*BaseInfoBean) MemberOutBankAdd(this *input.MemberBankAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mbb := new(schema.MemberBank)
	mbb.Card = this.Card
	mbb.BankId = this.BankId
	mbb.CardAddress = this.CardAddress
	mbb.CardName = this.CardName
	mbb.MemberId = this.MemberId
	count, err = sess.Table(mbb.TableName()).InsertOne(mbb)
	if err != nil {
		return
	}
	member := new(schema.Member)
	_, err = sess.Table(member.TableName()).Where("id=? AND delete_time=? AND status=?", this.MemberId, 0, 1).Get(member)
	if err != nil {
		return
	}
	if len(member.Realname) == 0 {
		member.Realname = this.CardName
		count, err = sess.Table(member.TableName()).Where("id=?", this.MemberId).
			Cols("realname").Update(member)
		if err != nil {
			return
		}
	}
	return
}

//修改会员出款银行
func (*BaseInfoBean) MemberOutBankUpdata(this *input.MemberBankChange) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mbb := new(schema.MemberBank)
	mbb.Card = this.Card
	mbb.CardAddress = this.CardAddress
	mbb.CardName = this.CardName
	mbb.MemberId = this.MemberId
	mbb.BankId = this.BankId
	count, err = sess.Table(mbb.TableName()).Where("delete_time=?", 0).Where("member_id=?", this.MemberId).Where("id=?", this.Id).Cols("bank_id,card,card_name,card_address").Update(mbb)
	return
}

//删除会员出款银行
func (*BaseInfoBean) MemberOutBankDelete(this *input.MemberBankDelete) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	if this.MemberId != 0 {
		sess.Where("member_id=?", this.MemberId)
	}
	mbb := new(schema.MemberBank)
	mbb.DeleteTime = time.Now().Unix()
	count, err = sess.Table(mbb.TableName()).Cols("delete_time").Update(mbb)
	return
}

//会员出款列表
//func (*BaseInfoBean) MemberBankList(this *input.MemberBankList) (data []back.MemberBankListBack, err error) {
//	sess := global.GetXorm().NewSession()
//	defer sess.Close()
//	mb := new(schema.MemberBank)
//	bank := new(schema.Bank)
//	if this.MemberId != 0 {
//		sess.Where(mb.TableName()+".member_id=?", this.MemberId)
//	}
//	sess.Where(mb.TableName()+".delete_time=?", 0)
//	sess.Where(bank.TableName()+".delete_time=?", 0)
//	where1 := fmt.Sprintf("%s.bank_id = %s.id", mb.TableName(), bank.TableName())
//	err = sess.Table(mb.TableName()).Join("LEFT", bank.TableName(), where1).Find(&data)
//	fmt.Println("出款银行列表参数**********", data)
//	return
//}
//会员出款列表
func (*BaseInfoBean) MemberBankList(this *input.MemberBankList) (data []back.MemberBankListBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//可用出款银行信息
	bankList, err := BankList("out")
	if err != nil {
		return
	}
	//获取站点剔除的出款银行
	BankOutDel := new(schema.BankOutDel)
	delBankOunt := []schema.BankOutDel{}
	sess.Where("site_id = ?", this.SiteId)
	sess.Where("site_index_id = ?", this.SiteIndexId)
	err = sess.Table(BankOutDel.TableName()).Find(&delBankOunt)
	outIds := []int64{}
	for _, value := range delBankOunt {
		outIds = append(outIds, value.BankId)
	}
	//获取会员出款银行
	memberBank := new(schema.MemberBank)
	sess.Where("delete_time = ?", 0)
	sess.Where("member_id = ?", this.MemberId)
	if len(outIds) > 0 {
		sess.NotIn("bank_id", outIds)
	}
	sess.Select("id, bank_id, card, card_name, card_address")
	err = sess.Table(memberBank.TableName()).Find(&data)
	for key, value := range data {
		for _, v := range bankList {
			if value.BankId == v.Id {
				data[key].Title = v.Title
			}
		}
	}
	return
}

//会员一条出款列表
func (*BaseInfoBean) OneMemberBankInfo(this *input.OneMemberBankInfo) (data *back.OneMemberBankInfoBack, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	if this.MemberId != 0 {
		sess.Where("member_id=?", this.MemberId)
	}
	sess.Where("delete_time=?", 0)
	mb := new(schema.MemberBank)
	data = new(back.OneMemberBankInfoBack)
	has, err = sess.Table(mb.TableName()).Get(data)
	return
}

//检验会员出款银行卡是否已添加
func (*BaseInfoBean) CheckOutBank(this *input.MemberBankAdd) (data *schema.MemberBank, has bool, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.MemberId != 0 {
		sess.Where("member_id=?", this.MemberId)
	}
	if this.BankId != 0 {
		sess.Where("bank_id=?", this.BankId)
	}
	if this.Card != "" {
		sess.Where("card=?", this.Card)
	}
	mb := new(schema.MemberBank)
	data = new(schema.MemberBank)
	has, err = sess.Table(mb.TableName()).Where("delete_time=?", 0).Get(data)
	count, err = sess.Table(mb.TableName()).Where("member_id=?", this.MemberId).
		Where("delete_time=?", 0).Count()
	return
}

//修改手机号
func (*BaseInfoBean) ChangePhoneNum(this *input.PhoneNum) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberInfo)
	mb.Mobile = this.PhoneNum
	mb.LocalCode = this.LocalCode
	count, err = sess.Table(mb.TableName()).Where("member_id=?", this.MemberId).Cols("mobile,local_code").Update(mb)
	return
}

//修改邮箱
func (*BaseInfoBean) ChangeEmailNum(this *input.EmailNum) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberInfo)
	mb.Email = this.EmailNum
	count, err = sess.Table(mb.TableName()).Where("member_id=?", this.MemberId).Cols("email").Update(mb)
	return
}

//修改生日
func (*BaseInfoBean) ChangeBirthdayNum(this *input.BirthdayNum) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.MemberInfo)
	tineUnix, err := time.Parse("2006-01-02", this.BirthdayNum)
	if err != nil {
		return
	}
	mb.Birthday = tineUnix.Unix()
	count, err = sess.Table(mb.TableName()).Where("member_id=?", this.MemberId).Cols("birthday").Update(mb)
	return
}

//修改资料
func (*BaseInfoBean) ChangeMeans(this *input.EditMeans) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	_, err = sess.Table(member.TableName()).Where("id=? AND delete_time=? AND status=?", this.MemberId, 0, 1).Get(member)
	if err != nil {
		return
	}

	if len(member.Realname) == 0 {
		member.Realname = this.Realname
		count, err = sess.Table(member.TableName()).Where("id=?", this.MemberId).
			Cols("realname").Update(member)
		if err != nil {
			return
		}
	}

	mb := new(schema.MemberInfo)

	var Colss []string
	if len(this.BirthdayNum) != 0 {
		mb.Birthday, _ = global.FormatDay2Timestamp2(this.BirthdayNum)
		Colss = append(Colss, "birthday")
	}

	if len(this.PhoneNum) != 0 {
		mb.Mobile = this.PhoneNum
		Colss = append(Colss, "mobile")
	}

	if len(this.EmailNum) != 0 {
		mb.Email = this.EmailNum
		Colss = append(Colss, "email")
	}

	if len(this.LocalCode) != 0 {
		mb.LocalCode = this.LocalCode
		Colss = append(Colss, "local_code")
	}

	if len(this.Card) != 0 {
		mb.Card = this.Card
		Colss = append(Colss, "card")
	}

	if len(this.QqNum) != 0 {
		mb.Qq = this.QqNum
		Colss = append(Colss, "qq")
	}

	if len(this.Wechat) != 0 {
		mb.Wechat = this.Wechat
		Colss = append(Colss, "wechat")
	}

	if len(this.Remark) != 0 {
		mb.Remark = this.Remark
		Colss = append(Colss, "remark")
	}

	ColsStr := strings.Join(Colss, ",")
	fmt.Println("修改的参数", mb)
	count, err = sess.Table(mb.TableName()).Where("member_id=?", this.MemberId).
		Cols(ColsStr).Update(mb)
	return
}

//获取一条会员详细资料
func (*BaseInfoBean) OneMemberDetail(this *input.MemberInfoSelf) (data *back.MemberDetailOne, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mi := new(schema.MemberInfo)
	data = new(back.MemberDetailOne)
	has, err = sess.Table(mi.TableName()).Where("member_id=?", this.Id).Get(data)
	return
}

/**
 * 获取银行列表
 * @param types string	银行列表类型	income 入款 out 出款 third 三方 为空的时候 查询全部
 * @return data []schema.Bank 银行卡列表
 */
func BankList(types string) (data []schema.Bank, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	bank := new(schema.Bank)
	sess.Where("delete_time = ?", 0)
	sess.Where("status = ?", 1)
	switch types {
	case "income":
		sess.Where("is_income = ?", 1)
	case "out":
		sess.Where("is_out = ?", 1)
	case "third":
		sess.Where("is_third = ?", 1)
	}
	err = sess.Table(bank.TableName()).Find(&data)
	return
}
