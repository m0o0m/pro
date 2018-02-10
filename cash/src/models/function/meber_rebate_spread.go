package function

import (
	"errors"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

//会员推广
type MemberSpreadBean struct {
}

//添加会员推广设定
func (*MemberSpreadBean) AddSpreadSet(this *input.SpreadSet) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spreadSet := new(schema.MemberSpreadSet)
	spreadSet.SiteId = this.SiteId
	spreadSet.SiteIndexId = this.SiteIndexId
	spreadSet.IsOpen = this.IsOpen
	spreadSet.IsIp = this.IsIp
	spreadSet.IsMateAgency = this.IsMateAgency
	spreadSet.IsCode = this.IsCode
	spreadSet.RankingMoney = this.RankingMoney
	spreadSet.RankingNum = this.RankingNum
	sess.Begin()
	count, err := sess.Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Delete(spreadSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	count, err = sess.Table(spreadSet.TableName()).InsertOne(spreadSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	sess.Commit()
	return count, err
}

//修改会员推广设定
func (*MemberSpreadBean) UpdateSpreadSet(this *input.SpreadEdit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spreadSet := new(schema.MemberSpreadSet)
	spreadSet.SiteId = this.SiteId
	spreadSet.SiteIndexId = this.SiteIndexId
	spreadSet.IsOpen = this.IsOpen
	spreadSet.IsIp = this.IsIp
	spreadSet.IsMateAgency = this.IsMateAgency
	spreadSet.IsCode = this.IsCode
	spreadSet.RankingMoney = this.RankingMoney
	spreadSet.RankingNum = this.RankingNum
	sess.Table(spreadSet.TableName()).Where("site_id=?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	count, err := sess.Update(spreadSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询会员推广设定
func (m *MemberSpreadBean) FindSpreadSet(siteId, siteIndexId string) (*back.SpreadSet, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	data := new(back.SpreadSet)
	spreadSet := new(schema.MemberSpreadSet)
	_, err := sess.Table(spreadSet.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询会员推广设定
func (m *MemberSpreadBean) GetSpreadSetBySite(siteId, siteIndexId string) (data back.SpreadSet, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spreadSet := new(schema.MemberSpreadSet)
	b, err := sess.Table(spreadSet.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Get(&data)
	if err != nil {
		return
	}
	if !(b) {
		err = errors.New("not found spread set")
		return
	}
	return
}
