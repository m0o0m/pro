package function

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"models/thirdParty"
	"time"
)

type MemberBalanceConversionBean struct{}

//获取会员指定视讯的余额
func (*MemberBalanceConversionBean) GetMoneyByVideo(memberId, platformId int64, sessArgs ...*xorm.Session) (back.MemberProductClassifyBalance, bool, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
		defer sess.Close()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	var mp back.MemberProductClassifyBalance
	mpcb := new(schema.MemberProductClassifyBalance)
	has, err := sess.Table(mpcb.TableName()).
		Where("member_id = ?", memberId).
		Where("platform_id = ?", platformId).
		Select("balance").
		Get(&mp)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mp, has, err
	}
	return mp, has, err
}

//根据转出项目获取余额(系统余额)
func (*MemberBalanceConversionBean) GetMoneyByOutTypes(memberId int64) (back.MemberProductClassifyBalance, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mp back.MemberProductClassifyBalance
	member := new(schema.Member)
	has, err := sess.Table(member.TableName()).
		Where("id = ?", memberId).
		Select("balance").Get(&mp)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mp, has, err
	}
	return mp, has, err
}

//根据转出项目和会员id获取余额
func (*MemberBalanceConversionBean) GetMoneys(memberId, forType int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mpcb := new(schema.MemberProductClassifyBalance)
	has, err = sess.Table(mpcb.TableName()).
		Where("platform_id = ?", forType).
		Where("member_id = ?", memberId).
		Select("balance").
		Get(mpcb)
	return
}

//转出项目/转入项目是否存在
func (*MemberBalanceConversionBean) IsExistFtype(ftype int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	has, err := sess.Table(product.TableName()).Where("platform_id= ?", ftype).Get(product)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//添加额度转换
func (mb *MemberBalanceConversionBean) BalanceConversionDo(this *input.MemberBalanceConversionAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//根据会员账号获取会员id,所属代理id
	member, _, err := GetMemberInfo(this.Account, this.SiteId, this.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	//获取站点下开户人视讯余额
	videoBalance, err := mb.GetAgency(this.SiteId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	mbc := new(schema.MemberBalanceConversion)
	mpcb := new(schema.MemberProductClassifyBalance)
	m := new(schema.Member)
	//给会员额度转换赋值
	mbc.Money = this.Money
	mbc.ForType = this.ForType
	mbc.FromType = this.FromType
	mbc.DoUserId = this.DoUserId
	mbc.Account = this.Account
	mbc.DoUserType = this.DoUserType
	mbc.Remark = this.Remark
	mbc.SiteId = member.SiteId
	mbc.SiteIndexId = member.SiteIndexId
	mbc.MemberId = member.Id
	mbc.AgencyId = member.ThirdAgencyId
	//给会员现金流水赋值
	mcr := new(schema.MemberCashRecord)
	mcr.SiteId = member.SiteId
	mcr.SiteIndexId = member.SiteIndexId
	mcr.MemberId = member.Id
	mcr.UserName = this.Account
	mcr.AgencyId = member.ThirdAgencyId
	mcr.SourceType = 8
	mcr.Type = 1
	mcr.Balance = this.Money
	mcr.Remark = this.Remark
	sess.Begin()
	//给会员额度转换表添加数据
	count, err := sess.Table(mbc.TableName()).InsertOne(mbc)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if count == 0 {
		return count, err
	}
	//给会员现金流水表添加数据
	count, err = sess.Table(mcr.TableName()).InsertOne(mcr)
	if err != nil || count == 0 {
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			sess.Rollback()
		}
		return count, err
	}
	videoBean := thirdParty.NewThirdParty()
	transferData := &thirdParty.TransferData{
		Credit: this.Money,
	}

	if mbc.ForType == 0 { //系统余额转换到其他余额
		//调用视讯
		transferData.TransferType = "out"
		_, err := videoBean.TransferCredit(transferData)
		if err != nil {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return 0, err
		}
		mbcn := new(schema.MemberBalanceConversion)
		mbcn.Status = 1 //1成功,2失败
		mbcn.UpdateTime = global.GetCurrentTime()
		count, err = sess.Table(mbcn.
			TableName()).
			Cols("status,update_time").
			Where("id = ?", mbc.Id).
			Update(mbcn)
		if err != nil {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return count, err
		}
		//获取会员对应各平台下余额
		has, err := mb.GetMoneys(member.Id, this.FromType)
		mpcb.MemberId = member.Id
		mpcb.PlatformId = this.FromType
		//平台余额=视讯返回的余额
		mpcb.Balance = this.Money
		if !has {
			//给会员对应各平台下余额表添加数据
			count, err = sess.Table(mpcb.TableName()).InsertOne(mpcb)
			if err != nil || count == 0 {
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					sess.Rollback()
				}
				return count, err
			}
		} else {
			//给会员对应各分类下余额表修改数据
			count, err = sess.Table(mpcb.TableName()).
				Where("member_id = ?", mpcb.MemberId).
				Where("platform_id = ?", mpcb.PlatformId).
				Cols("balance").
				Update(mpcb)
			if err != nil || count == 0 {
				if err != nil {
					global.GlobalLogger.Error("err:%s", err.Error())
					sess.Rollback()
				}
				return count, err
			}
		}
		//获取会员余额
		balance, err := mb.GetBalance(this)
		m.Id = member.Id
		m.Balance = balance - this.Money
		//给会员表修改数据
		count, err = sess.Where("id = ?", m.Id).Cols("balance").Update(m)
		if err != nil || count == 0 {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return count, err
		}
		//给代理表中的视讯扣除手续费
		agency := new(schema.Agency)
		agency.VideoBalance = videoBalance.VideoBalance - this.Margin
		sess.Where("site_id = ?", this.SiteId).
			Where("is_sub=2").
			Cols("video_balance")
		if this.SiteIndexId != "" {
			sess.Where("site_index_id = ?", this.SiteIndexId)
		}
		agency.VideoBalance = global.FloatReserve2(agency.VideoBalance)
		count, err = sess.Update(agency)

		if err != nil || count == 0 {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return count, err
		}
		//给站点额度记录表添加数据
		scr := new(schema.SiteCashRecord)
		scr.SiteId = this.SiteId
		scr.SiteIndexId = this.SiteIndexId
		scr.Balance = agency.VideoBalance
		scr.Remark = this.Remark
		scr.Money = this.Money
		scr.AdminName = this.Account
		scr.CashType = 1
		scr.DoType = 2
		scr.State = mbcn.Status
		scr.VdType = this.ForType
		count, err = sess.Table(scr.TableName()).InsertOne(scr)
		if err != nil || count == 0 {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return count, err
		}
	} else { //其他余额转换到系统余额
		//调用视讯
		transferData.TransferType = "in"
		_, err := videoBean.TransferCredit(transferData)
		if err != nil {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return 0, err
		}
		mbcn := new(schema.MemberBalanceConversion)
		mbcn.Status = 1
		mbcn.UpdateTime = global.GetCurrentTime()
		count, err = sess.Table(mbcn.TableName()).Cols("status,update_time").Where("id = ?", mbc.Id).Update(mbcn)
		if err != nil || count == 0 {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return count, err
		}
		//获取会员对应各交易平台下余额
		_, _, err = mb.GetMoneyByVideo(member.Id, this.ForType)
		mpcb.MemberId = member.Id
		//平台余额=视讯返回的余额
		mpcb.Balance = this.Money
		mpcb.PlatformId = this.ForType
		//给会员对应各交易平台下余额表修改数据
		count, err = sess.Table(mpcb.TableName()).
			Where("member_id = ?", mpcb.MemberId).
			Where("platform_id = ?", mpcb.PlatformId).
			Cols("balance").
			Update(mpcb)
		if err != nil || count == 0 {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return count, err
		}
		//获取会员余额
		balance, err := mb.GetBalance(this)
		m.Id = member.Id
		m.Balance = balance + this.Money
		//给会员表修改数据
		count, err = sess.Where("id = ?", m.Id).Cols("balance").Update(m)
		if err != nil || count == 0 {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return count, err
		}
		//给站点表中的视讯余额加上手续费
		agency := new(schema.Agency)
		agency.VideoBalance = videoBalance.VideoBalance + this.Margin
		sess.Where("site_id = ?", this.SiteId).
			Where("is_sub=2").
			Cols("video_balance")
		if this.SiteIndexId != "" {
			sess.Where("site_index_id = ?", this.SiteIndexId)
		}
		agency.VideoBalance = global.FloatReserve2(agency.VideoBalance)
		count, err = sess.Update(agency)
		if err != nil || count == 0 {

			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			} else {
				global.GlobalLogger.Error("err:281 : update agency 0 row")
			}
			return count, err
		}
		//给站点额度记录表添加数据
		scr := new(schema.SiteCashRecord)
		scr.SiteId = this.SiteId
		scr.SiteIndexId = this.SiteIndexId
		scr.Balance = agency.VideoBalance
		scr.Remark = this.Remark
		scr.Money = this.Money
		scr.AdminName = this.Account
		scr.CashType = 1
		scr.DoType = 1
		scr.State = mbcn.Status
		scr.VdType = this.ForType
		count, err = sess.Insert(scr)
		if err != nil || count == 0 {
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
			}
			return count, err
		}
	}
	sess.Commit()
	return count, err
}

//获取会员的系统余额
func (*MemberBalanceConversionBean) GetBalance(memberBalanceConversion *input.MemberBalanceConversionAdd) (float64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	var balance float64
	sess.Table(member.TableName()).
		Where("account = ?", memberBalanceConversion.Account).
		Where("site_id = ?", memberBalanceConversion.SiteId).
		Select("balance")
	if memberBalanceConversion.SiteIndexId != "" {
		sess.Where("site_index_id = ?", memberBalanceConversion.SiteIndexId)
	}
	b, err := sess.Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return balance, err
	}
	if !(b) {
		err = errors.New("not found member money")
		return balance, err
	}
	balance = member.Balance
	return balance, err
}

//获取会员余额统计
func (mb *MemberBalanceConversionBean) ClassifyBalance(this *input.MemberClassifyBalance) ([]back.MemberClassifyBalance, error) {
	sess := global.GetXorm().NewSession()
	mpcb := new(schema.MemberProductClassifyBalance)
	member := new(schema.Member)
	product := new(schema.Product)
	agency := new(schema.Agency)
	var data []back.MemberClassifyBalance
	defer sess.Close()

	if this.DataType == 0 {
		this.DataType = 1
	}
	switch this.DataType { //关联条件
	case 1: //商品	传入参数 站点site_index_id 代理 agency_id
		if this.AgencyId != 0 { //代理id为0 查询全部
			sess.Where(fmt.Sprintf("%s.third_agency_id = ?", member.TableName()), this.AgencyId)
		}
		sess.Join("LEFT", member.TableName(), fmt.Sprintf("%s.member_id = %s.id", mpcb.TableName(), member.TableName()))
		sess.Join("LEFT", product.TableName(), fmt.Sprintf("%s.product_id = %s.id", mpcb.TableName(), product.TableName()))

		var selectstr string
		selectstr = fmt.Sprintf("%s.account,%s.balance mbalance,%s.status,%s.balance,%s.product_name", member.TableName(), member.TableName(), member.TableName(), mpcb.TableName(), product.TableName())
		sess.Select(selectstr)
		break
	case 2: //代理	代理账号agency_name  站点site_index_id  商品id product_id
		if len(this.AgencyName) != 0 { //代理账号为空 查全部
			sess.Where(fmt.Sprintf("%s.account = ?", agency.TableName()), this.AgencyName)
		}
		sess.Join("LEFT", agency.TableName(), fmt.Sprintf("%s.third_agency_id = %s.id", member.TableName(), agency.TableName()))

		var selectstr string
		if this.ProductId != 0 { //ProductId为0时 为系统余额
			sess.Join("LEFT", mpcb.TableName(), fmt.Sprintf("%s.member_id = %s.id", mpcb.TableName(), member.TableName()))
			sess.Where(fmt.Sprintf("%s.product_id = ?", mpcb.TableName()), this.ProductId)
			selectstr = fmt.Sprintf("%s.account,%s.account product_name, %s.status,%s.balance", member.TableName(), agency.TableName(), member.TableName(), mpcb.TableName())
		} else {
			selectstr = fmt.Sprintf("%s.account,%s.account product_name, %s.status,%s.balance", member.TableName(), agency.TableName(), member.TableName(), member.TableName())
		}
		sess.Select(selectstr)
		break
	case 3: //会员 代理id agency_id
		sess.Where(fmt.Sprintf("%s.third_agency_id = ?", member.TableName()), this.AgencyId)
		if len(this.Account) != 0 { //账号为空 查全部
			sess.Where(fmt.Sprintf("%s.account = ?", member.TableName()), this.Account)
		}
		var selectstr string
		if this.ProductId != 0 { //ProductId为0时 为系统余额
			sess.Join("LEFT", mpcb.TableName(), fmt.Sprintf("%s.member_id = %s.id", mpcb.TableName(), member.TableName()))
			sess.Where(fmt.Sprintf("%s.product_id = ?", mpcb.TableName()), this.ProductId)
			selectstr = fmt.Sprintf("%s.account,%s.status,%s.balance", member.TableName(), member.TableName(), mpcb.TableName())
		} else {
			selectstr = fmt.Sprintf("%s.account,%s.status,%s.balance", member.TableName(), member.TableName(), member.TableName())
		}
		sess.Select(selectstr)
		break
	}

	sess.Where(fmt.Sprintf("%s.site_id = ?", member.TableName()), this.SiteId)
	if len(this.SiteIndexId) != 0 {
		sess.Where(fmt.Sprintf("%s.site_index_id = ?", member.TableName()), this.SiteIndexId)
	}

	//查询 站点下的所有会员
	var sdata []back.MemberClassifyArr
	if this.DataType == 1 {
		err := sess.Table(mpcb.TableName()).Find(&sdata)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	} else {
		err := sess.Table(member.TableName()).Find(&sdata)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
	}
	data = mb.classifyarr(sdata, this.DataType)
	return data, nil
}

//根据代理账号查询代理ID
func (*MemberBalanceConversionBean) AgencyByid(Account string) (Id int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if len(Account) != 0 {
		sess.Where("account=?", Account)
	}
	_, err = sess.Table(agency.TableName()).Select("id").Get(agency)
	Id = agency.Id
	return
}

//会员余额统计数据处理
func (*MemberBalanceConversionBean) classifyarr(sdata []back.MemberClassifyArr, DataType int) (data []back.MemberClassifyBalance) {
	item := make(map[string]bool)
	mitem := make(map[string]bool)
	var fdata back.MemberClassifyBalance
	StrMoney := make(map[string]float64)
	EndMoney := make(map[string]float64)

	for _, v := range sdata {
		if DataType == 1 {
			if mitem[v.Account] != true {
				mitem[v.Account] = true
				if v.Status == 1 {
					StrMoney["系统额度"] = StrMoney["系统额度"] + v.Mbalance
				} else {
					EndMoney["系统额度"] = StrMoney["系统额度"] + v.Mbalance
				}
			}
		} else if DataType == 3 {
			v.ProductName = v.Account
		}

		if item[v.ProductName] == true {
			if v.Status == 1 {
				StrMoney[v.ProductName] = StrMoney[v.ProductName] + v.Balance
			} else {
				EndMoney[v.ProductName] = StrMoney[v.ProductName] + v.Balance
			}
		} else {
			item[v.ProductName] = true
			if v.Status == 1 {
				StrMoney[v.ProductName] = v.Balance
			} else {
				EndMoney[v.ProductName] = v.Balance
			}
		}
	}

	for k, v := range StrMoney {
		fdata.ProductName = k
		fdata.StrMoney = v
		fdata.EndMoney = EndMoney[k]
		fdata.UpdateTime = time.Now().Unix()
		data = append(data, fdata)
	}
	return
}

//获取会员id
func (*MemberBalanceConversionBean) GetMemberId(account string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	_, err := sess.Table(member.TableName()).Where("account = ?", account).
		Select("id").Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	memberId := member.Id
	return memberId, err
}

//获取总站点套餐id
func (*MemberBalanceConversionBean) GetSiteCombo(siteId string) (*schema.Site, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	_, err := sess.Table(site.TableName()).
		Where("id = ?", siteId).
		Where("index_id = ?", "a").
		Select("combo_id").
		Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return site, err
	}
	return site, err
}

//根据转入项目和套餐id获取手续费占成比
func (*MemberBalanceConversionBean) GetProductProportion(forType, comboId int64) ([]float64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var proportion []float64
	cp := new(schema.ComboProduct)
	var p []back.Proportion
	err := sess.Table(cp.TableName()).
		Where("platform_id = ?", forType).
		Where("combo_id = ?", comboId).
		Desc("proportion").
		Find(&p)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return proportion, err
	}
	if len(p) != 0 {
		for k := range p {
			proportion = append(proportion, p[k].Proportion)
		}
	}
	return proportion, err
}

//额度转换列表
func (*MemberBalanceConversionBean) GetList(this *input.MemberBalanceConversionList, listParams *global.ListParams, times *global.Times) ([]back.MemberBalanceConversionBackList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mbc := new(schema.MemberBalanceConversion)
	var data []back.MemberBalanceConversionBackList
	//判断并组合where条件
	if this.SiteId != "" {
		sess.Where("sales_member_balance_conversion.site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("sales_member_balance_conversion.site_index_id = ?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("sales_member_balance_conversion.account = ?", this.Account)
	}
	//根据时间段查询
	times.Make("sales_member_balance_conversion.create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	//获得分页记录
	listParams.Make(sess)
	mbcl := make([]back.MemberBalanceConversionList, 0)
	product := new(schema.Product)
	platform := new(schema.Platform)
	sess.Select("sales_member_balance_conversion.for_type,sales_platform.platform," +
		"sales_member_balance_conversion.id,sales_member_balance_conversion.create_time," +
		"sales_member_balance_conversion.account,sales_member_balance_conversion.money," +
		"sales_member_balance_conversion.remark,sales_member_balance_conversion.update_time," +
		"sales_member_balance_conversion.status")
	sql1 := fmt.Sprintf("%s.platform_id = %s.id", product.TableName(), platform.TableName())
	sql2 := fmt.Sprintf("%s.from_type = %s.platform_id OR %s.for_type = %s.platform_id", mbc.TableName(), product.TableName(), mbc.TableName(), product.TableName())
	//重新传入表名和where条件查询记录
	err := sess.Table(mbc.TableName()).Join("LEFT", product.TableName(), sql2).
		Join("LEFT", platform.TableName(), sql1).GroupBy("sales_member_balance_conversion.id," +
		"sales_member_balance_conversion.for_type,sales_platform.platform," +
		"sales_member_balance_conversion.create_time,sales_member_balance_conversion.account," +
		"sales_member_balance_conversion.money,sales_member_balance_conversion.remark," +
		"sales_member_balance_conversion.update_time,sales_member_balance_conversion.status").Find(&mbcl)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	var m back.MemberBalanceConversionBackList
	for k := range mbcl {
		if mbcl[k].ForType == 0 {
			m.Type = "入款"
		} else {
			m.Type = "取款"
		}
		m.Platform = mbcl[k].Platform
		m.Id = mbcl[k].Id
		m.Account = mbcl[k].Account
		m.Money = mbcl[k].Money
		m.CreateTime = mbcl[k].CreateTime
		m.UpdateTime = mbcl[k].UpdateTime
		m.Status = mbcl[k].Status
		m.Remark = mbcl[k].Remark
		data = append(data, m)
	}
	//获得符合条件的记录数
	count, err := sess.Table(mbc.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取站点视讯余额(保证金)
func (*MemberBalanceConversionBean) GetAgency(siteId string, sessArgs ...*xorm.Session) (schema.Agency, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
		defer sess.Close()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	var agency schema.Agency
	b, err := sess.Where("site_id=?", siteId).
		Where("status=1").
		Where("is_sub=2").
		Where("level=1").
		Get(&agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return agency, err
	}
	if !(b) {
		err = errors.New("not found agency")
		return agency, nil
	}
	return agency, err
}

//wap 会员中心--会员余额刷新
func (*MemberBalanceConversionBean) MemberBalance(memberId int64) (
	back.MemberBalance, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var memberBalance back.MemberBalance
	member := new(schema.Member)
	//获取会员账号余额
	has, err := sess.Select("balance").Where("id=?", memberId).Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberBalance, err
	}
	if !has {
		return memberBalance, err
	}
	//余额
	balance := make([]back.Balance, 0)
	//获取会员各视讯余额
	memberProductClassifyBalance := new(schema.MemberProductClassifyBalance)
	err = sess.Table(memberProductClassifyBalance.TableName()).Select("balance").
		Where("member_id=?", memberId).Find(&balance)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberBalance, err
	}
	//赋值
	memberBalance.AccountBalance = member.Balance
	for k := range balance {
		memberBalance.GameBalance += balance[k].Balance
	}
	return memberBalance, err
}

//wap 额度转换-获取各平台余额
func (m *MemberBalanceConversionBean) GetPlatformBalance(memberId int64) (
	back.MemberPlatformBalance, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	var memberBalance back.MemberPlatformBalance
	//获取会员账号余额
	has, err := sess.Select("account,realname,balance").Where("id=?", memberId).Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberBalance, err
	}
	if !has {
		return memberBalance, err
	}
	sumBalance, balances, err := m.GetVideoBalance(memberId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberBalance, err
	}
	//赋值
	memberBalance.AccountBalance = member.Balance
	memberBalance.Account = member.Account
	memberBalance.Realname = member.Realname
	memberBalance.GameBalance = sumBalance
	memberBalance.ProductClassifyBalance = balances
	return memberBalance, err
}

//获取会员各视讯余额
func (*MemberBalanceConversionBean) GetVideoBalance(memberId int64) (sumBalance float64, balances []back.ProductClassifyBalance, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//获取会员各视讯余额
	memberProductClassifyBalance := new(schema.MemberProductClassifyBalance)
	platform := new(schema.Platform)

	// DESCRIPTION:查询所有平台
	err = sess.Table(platform.TableName()).
		Select("id as platform_id,platform").
		Where("status = 1").
		Where("delete_time = 0").
		Find(&balances)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	var moneys []schema.MemberProductClassifyBalance
	// // DESCRIPTION:查询会员所有视讯余额
	err = sess.Table(memberProductClassifyBalance.TableName()).
		Where("member_id=?", memberId).
		Find(&moneys)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	moneyMap := make(map[int64]*schema.MemberProductClassifyBalance)
	for i, _ := range moneys {
		moneyMap[moneys[i].PlatformId] = &moneys[i]
		sumBalance += moneys[i].Balance
	}
	for k, balance := range balances {
		if money, ok := moneyMap[balance.PlatformId]; ok {
			balances[k].Balance = money.Balance
		}
	}
	sumBalance = global.FloatReserve2(sumBalance)
	return
}

//重置会员视讯额度
func (*MemberBalanceConversionBean) ResetVideoBalances(memberId int64, platformIds []int64, sessArgs ...*xorm.Session) (count int64, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
		defer sess.Close()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	memberProductClassifyBalance := new(schema.MemberProductClassifyBalance)
	count, err = sess.Table(memberProductClassifyBalance.TableName()).
		Where("member_id = ?", memberId).
		In("platform", platformIds).
		Update(map[string]interface{}{"balance": 0})
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	return
}

//重置会员视讯额度
func (*MemberBalanceConversionBean) ResetVideoBalance(memberId int64, platformId int64, sessArgs ...*xorm.Session) (err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
		defer sess.Close()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	memberProductClassifyBalance := new(schema.MemberProductClassifyBalance)
	count, err := sess.Table(memberProductClassifyBalance.TableName()).
		Where("member_id = ?", memberId).
		In("platform_id", platformId).
		Update(map[string]interface{}{"balance": 0})
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	if count == 0 {
		err = errors.New("update 0 row")
	}
	return
}

//wap 额度转换-单个平台余额刷新
func (*MemberBalanceConversionBean) PlatformBalanceRefresh(memberId, platformId int64) (
	back.PlatformBalanceRefresh, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var memberBalance back.PlatformBalanceRefresh
	member := new(schema.Member)
	memberProductClassifyBalance := new(schema.MemberProductClassifyBalance)
	//调用视讯接口
	resBalance, err := global.PlatformBalanceRefresh()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberBalance, err
	}
	//获取会员账号余额
	has, err := sess.Select("balance").Where("id=?", memberId).Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberBalance, err
	}
	if !has {
		return memberBalance, err
	}
	//查看会员视讯余额是否存在(这个查看和不存在的判断或许是冗余的，后期如果不需要可以删除)
	has, err = sess.Where("member_id=?", memberId).Where("platform_id=?", platformId).Get(memberProductClassifyBalance)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberBalance, err
	}
	memberProductClassifyBalance.Balance = resBalance.Balance
	if !has {
		//不存在，新增本地会员视讯余额
		memberProductClassifyBalance.MemberId = memberId
		memberProductClassifyBalance.PlatformId = platformId
		count, err := sess.InsertOne(memberProductClassifyBalance)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return memberBalance, err
		}
		if count != 1 {
			return memberBalance, err
		}
	} else {
		//存在，修改本地会员视讯余额
		count, err := sess.Where("member_id=?", memberId).Where("platform_id=?", platformId).
			Cols("balance").Update(memberProductClassifyBalance)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return memberBalance, err
		}
		if count != 1 {
			return memberBalance, err
		}
	}
	//申明余额
	balance := make([]back.ProductClassifyBalance, 0)
	platform := new(schema.Platform)
	sql := fmt.Sprintf("%s.id=%s.platform_id", platform.TableName(), memberProductClassifyBalance.TableName())
	err = sess.Table(memberProductClassifyBalance.TableName()).Join("LEFT", platform.TableName(), sql).
		Where("member_id=?", memberId).Where("platform_id=?", platformId).Find(&balance)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberBalance, err
	}
	//赋值
	memberBalance.AccountBalance = member.Balance //账号总额
	for k := range balance {
		memberBalance.GameBalance += balance[k].Balance //视讯总额
	}
	memberBalance.ProductClassifyBalance = balance //视讯各平台余额
	return memberBalance, err
}

//wap 添加额度转换
func (mb *MemberBalanceConversionBean) WapBalanceConversionDo(this *input.WapMemberBalanceConversion) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//根据会员账号获取会员id,所属代理id
	member, _, err := WapGetMemberInfo(this.MemberId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	//获取站点下开户人视讯余额
	videoBalance, err := mb.GetAgency(this.SiteId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	mbc := new(schema.MemberBalanceConversion)
	mpcb := new(schema.MemberProductClassifyBalance)
	m := new(schema.Member)
	//给会员额度转换赋值
	mbc.Money = this.Money
	mbc.ForType = this.ForType
	mbc.FromType = this.FromType
	mbc.DoUserId = this.MemberId
	mbc.DoUserType = 2
	mbc.Account = member.Account
	mbc.SiteId = member.SiteId
	mbc.SiteIndexId = member.SiteIndexId
	mbc.MemberId = this.MemberId
	mbc.AgencyId = member.ThirdAgencyId
	//给会员现金流水赋值
	mcr := new(schema.MemberCashRecord)
	mcr.SiteId = member.SiteId
	mcr.SiteIndexId = member.SiteIndexId
	mcr.MemberId = this.MemberId
	mcr.UserName = member.Account
	mcr.AgencyId = member.ThirdAgencyId
	mcr.SourceType = 8
	mcr.Type = 1
	mcr.Balance = this.Money
	sess.Begin()
	//给会员额度转换表添加数据
	count, err := sess.Table(mbc.TableName()).InsertOne(mbc)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//给会员现金流水表添加数据
	count, err = sess.Table(mcr.TableName()).InsertOne(mcr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	if mbc.ForType == 0 { //系统余额转换到其他余额
		//调用视讯
		res, err := global.Video()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		mbcn := new(schema.MemberBalanceConversion)
		mbcn.Status = res.Status
		mbcn.UpdateTime = res.ConfirmTime
		count, err = sess.Table(mbcn.TableName()).Cols("status,update_time").
			Where("id = ?", mbc.Id).Update(mbcn)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if count == 0 {
			return count, err
		}
		//获取会员对应各平台下余额
		has, err := mb.GetMoneys(this.MemberId, this.FromType)
		mpcb.MemberId = this.MemberId
		mpcb.PlatformId = this.FromType
		//平台余额=视讯返回的余额
		mpcb.Balance = res.Balance
		if !has {
			//给会员对应各平台下余额表添加数据
			count, err = sess.Table(mpcb.TableName()).InsertOne(mpcb)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				sess.Rollback()
				return 0, err
			}
			if count == 0 {
				return count, err
			}
		} else {
			//给会员对应各分类下余额表修改数据
			count, err = sess.Table(mpcb.TableName()).Where("member_id = ?", mpcb.MemberId).
				Where("platform_id = ?", mpcb.PlatformId).Cols("balance").Update(mpcb)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				sess.Rollback()
				return 0, err
			}
			if count != 1 {
				return count, err
			}
		}
		//获取会员余额
		balance, err := mb.WapGetBalance(this.MemberId)
		m.Id = this.MemberId
		m.Balance = balance - this.Money
		//给会员表修改数据
		count, err = sess.Where("id = ?", m.Id).Cols("balance").Update(m)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
		//给代理表中的视讯扣除手续费
		agency := new(schema.Agency)
		agency.VideoBalance = videoBalance.VideoBalance - this.Fee
		count, err = sess.Where("site_id = ?", this.SiteId).
			Where("is_sub=2").Cols("video_balance").Update(agency)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
		//给站点额度记录表添加数据
		scr := new(schema.SiteCashRecord)
		scr.SiteId = this.SiteId
		scr.SiteIndexId = member.SiteIndexId
		scr.Balance = agency.VideoBalance
		scr.Money = this.Money
		scr.AdminName = member.Account
		scr.CashType = 1
		scr.DoType = 2
		scr.State = mbcn.Status
		scr.VdType = this.ForType
		count, err = sess.Table(scr.TableName()).InsertOne(scr)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
	} else { //其他余额转换到系统余额
		//调用视讯
		res, err := global.Video()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		mbcn := new(schema.MemberBalanceConversion)
		mbcn.Status = res.Status
		mbcn.UpdateTime = res.ConfirmTime
		count, err = sess.Table(mbcn.TableName()).Cols("status,update_time").
			Where("id = ?", mbc.Id).Update(mbcn)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
		//获取会员对应各交易平台下余额
		_, _, err = mb.GetMoneyByVideo(this.MemberId, this.ForType)
		mpcb.MemberId = this.MemberId
		//平台余额=视讯返回的余额
		mpcb.Balance = res.Balance
		mpcb.PlatformId = this.ForType
		//给会员对应各交易平台下余额表修改数据
		count, err = sess.Table(mpcb.TableName()).Where("member_id = ?", mpcb.MemberId).
			Where("platform_id = ?", mpcb.PlatformId).Cols("balance").Update(mpcb)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
		//获取会员余额
		balance, err := mb.WapGetBalance(this.MemberId)
		m.Id = this.MemberId
		m.Balance = balance + this.Money
		//给会员表修改数据
		count, err = sess.Where("id = ?", m.Id).Cols("balance").Update(m)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
		//给站点表中的视讯余额加上手续费
		agency := new(schema.Agency)
		agency.VideoBalance = videoBalance.VideoBalance + this.Fee
		count, err = sess.Where("site_id = ?", this.SiteId).
			Where("is_sub=2").Cols("video_balance").Update(agency)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
		//给站点额度记录表添加数据
		scr := new(schema.SiteCashRecord)
		scr.SiteId = this.SiteId
		scr.SiteIndexId = member.SiteIndexId
		scr.Balance = agency.VideoBalance
		scr.Money = this.Money
		scr.AdminName = member.Account
		scr.CashType = 1
		scr.DoType = 1
		scr.State = mbcn.Status
		scr.VdType = this.ForType
		count, err = sess.Table(scr.TableName()).InsertOne(scr)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
	}
	sess.Commit()
	return count, err
}

//wap 获取会员的系统余额
func (*MemberBalanceConversionBean) WapGetBalance(memberId int64) (float64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	_, err := sess.Table(member.TableName()).Where("id = ?", memberId).Select("balance").Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	balance := member.Balance
	return balance, err
}

//获取会员额度转换记录
func (*MemberBalanceConversionBean) GetBalanceList(siteId, siteIndexId string, memberId int64, times *global.Times, listparams *global.ListParams) (data []schema.MemberBalanceConversion, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mbc := new(schema.MemberBalanceConversion)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("member_id=?", memberId)
	conds := sess.Conds()
	listparams.Make(sess)
	err = sess.Table(mbc.TableName()).Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(mbc.TableName()).Where(conds).Count()
	return
}
