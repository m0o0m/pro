package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
)

//代理详细资料
type AgencyThirdInfoBean struct{}

//查询某个代理的详细资料
func (*AgencyThirdInfoBean) ThirdAgencyInfo(this *input.ThirdInformation) (*back.ThirdInformationBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info := new(schema.AgencyThirdInfo)
	data := new(back.ThirdInformationBack)
	has, err := sess.Table(info.TableName()).
		Where("agency_id=?", this.Id).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//查询代理的银行卡信息
func (*AgencyThirdInfoBean) ThirdAgencyBankInfo(id int64) ([]back.ThirdAgencyInfoByBank, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	tBank := new(schema.AgencyThirdBank)
	var data []back.ThirdAgencyInfoByBank
	sess.Where("agency_id=?", id)
	err := sess.Table(tBank.TableName()).Find(&data)
	return data, err
}

//查询代理的代理域名
func (*AgencyThirdInfoBean) ThirdAgencyDomainInfo(id int64) ([]back.ThirdAgencyInfoByDomain, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	dM := new(schema.AgencyThirdDomain)
	var data []back.ThirdAgencyInfoByDomain
	sess.Where("delete_time=?", 0)
	sess.Where("agency_id=?", id)
	err := sess.Table(dM.TableName()).Find(&data)
	return data, err
}

//修改代理的详细资料
func (*AgencyThirdInfoBean) ThirdAgencyInfoUpdata(this *input.ThirdInformationUpdata) (int64, int8, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var count int64
	var num int8
	var err error
	info := new(schema.AgencyThirdInfo)
	//查询表
	has, err := sess.Where("agency_id=?", this.AgencyId).Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, num, err
	}
	sess.Begin()
	info.AgencyId = this.AgencyId
	info.AreaId = this.AreaId
	info.Card = this.Card
	info.ChName = this.ChName
	info.CityId = this.CityId
	info.Email = this.Email
	info.Phone = this.Phone
	info.ProvinceId = this.ProvinceId
	info.QQ = this.QQ
	info.Remark = this.Remark
	info.UsName = this.UsName
	info.SpreadId = this.SpreadId
	if has {
		//如果存在就修改
		count, err = sess.Where("agency_id=?", info.AgencyId).
			Cols("ch_name,us_name,card,phone,qq,email,province_id," +
				"city_id,area_id,spread_id").
			Update(info)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, num, err
		}
	} else {
		//不存在就添加
		count, err = sess.InsertOne(info)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, num, err
		}
	}
	//修改银行卡信息
	if len(this.Ids) > 0 {
		ids := strings.Split(this.Ids, ",")
		bankId := strings.Split(this.BankIds, ",")
		tB := new(schema.AgencyThirdBank)
		sql := "UPDATE " + tB.TableName() + " SET" + " bank_id=CASE id"
		for k, v := range ids {
			sql = sql + " WHEN " + v + " THEN " + bankId[k]
		}
		sql = sql + " END," + tB.TableName() + ".card_address = CASE id"
		for q, w := range ids {
			sql = sql + " WHEN " + w + " THEN " + "'" + this.CardAddress[q] + "'"
		}
		sql = sql + " END," + tB.TableName() + ".card = CASE id"
		for t, y := range ids {
			sql = sql + " WHEN " + y + " THEN " + "'" + this.Cards[t] + "'"
		}
		sql = sql + " END"
		sql = sql + " WHERE id IN (" + this.Ids + ")"
		_, err = sess.Query(sql)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, num, err
		}

	}
	if len(this.DoIds) > 0 {
		//查询修改的域名是否被使用
		dM := new(schema.AgencyThirdDomain)
		has, err := sess.Where("agency_id!=?", this.AgencyId).
			Where("delete_time=?", 0).In("domain", this.Domain).Get(dM)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, num, err
		}
		if has {
			num = 1
			sess.Rollback()
			return count, num, err
		}
		doIds := strings.Split(this.DoIds, ",")
		sql := "UPDATE " + dM.TableName() + " SET" + " domain=CASE id"
		for k, v := range doIds {
			sql = sql + " WHEN " + v + " THEN " + "'" + this.Domain[k] + "'"
		}
		sql = sql + " END"
		sql = sql + " WHERE id IN (" + this.DoIds + ")"
		_, err = sess.Query(sql)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, num, err
		}
	}
	sess.Commit()
	return count, num, err
}

//根据代理账号查询代理账号等级level
func (*AgencyThirdInfoBean) ThirdAgencyInfoLevel(Account string) (bool, error, int8) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.Account = Account
	ok, err := sess.Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err, 0
	}
	level := agency.Level
	return ok, err, level
}
