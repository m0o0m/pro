package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type AgencyMemberRegisterDiscountSetBean struct {
}

//查询会员注册设定
func (*AgencyMemberRegisterDiscountSetBean) GetOneDiscountSet(this *input.FirstDiscountSet) (*back.FirstDiscountSetBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	set := new(schema.AgencyMemberRegisterDiscountSet)
	data := new(back.FirstDiscountSetBack)
	if this.AcountId != 0 {
		sess.Where("agency_id=?", this.AcountId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	ok, err := sess.Table(set.TableName()).
		Select("site_index_id,agency_id,offer,add_mosaic,is_ip").
		Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, ok, err
	}
	return data, ok, err
}

//修改、增加会员注册设定
func (*AgencyMemberRegisterDiscountSetBean) UpdataSet(this *input.FirstDiscountUpdata) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var count int64
	set := new(schema.AgencyMemberRegisterDiscountSet)
	var set_info schema.AgencyMemberRegisterDiscountSet
	has, err := sess.Table(set.TableName()).
		Where("delete_time=?", 0).
		Where("agency_id=?", this.AgencyId).
		Where("site_index_id=?", this.SiteIndexId).
		Get(&set_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	if !has {
		set_info.SiteId = this.SiteId
		set_info.SiteIndexId = this.SiteIndexId
		set_info.AgencyId = this.AgencyId
		set_info.AddMosaic = this.AddMosaic
		set_info.IsIp = this.IsIp
		set_info.Offer = this.Offer
		count, err = sess.Table(set.TableName()).InsertOne(set_info)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else {
		set_info.AddMosaic = this.AddMosaic
		set_info.IsIp = this.IsIp
		set_info.Offer = this.Offer
		count, err = sess.Table(set.TableName()).
			Where("site_id=?", this.SiteId).
			Where("agency_id=?", this.AgencyId).
			Where("site_index_id=?", this.SiteIndexId).
			Where("delete_time=?", 0).
			Cols("site_id,site_index_id,agency_id,offer,add_mosaic,is_ip").
			Update(set_info)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	}
	return count, err
}

//查询总代对会员注册设定
func (*AgencyMemberRegisterDiscountSetBean) GetOneSecondDiscountSet(this *input.SecondDiscountSet) (
	*back.SecondDiscountSetBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	set := new(schema.AgencyMemberRegisterDiscountSet)
	data := new(back.SecondDiscountSetBack)
	if this.AcountId != 0 {
		sess.Where("agency_id=?", this.AcountId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	ok, err := sess.Table(set.TableName()).
		Select("site_index_id,agency_id,offer,add_mosaic,is_ip").
		Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, ok, err
	}
	return data, ok, err
}

//修改、增加会员注册设定(总代)
func (*AgencyMemberRegisterDiscountSetBean) UpdataSecondSet(this *input.SecondDiscountUpdata) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	set := new(schema.AgencyMemberRegisterDiscountSet)
	var set_info schema.AgencyMemberRegisterDiscountSet
	var count int64
	has, err := sess.Table(set.TableName()).Where("delete_time=?", 0).Where("agency_id=?", this.AgencyId).Where("site_index_id=?", this.SiteIndexId).Get(&set_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	if !has {
		set_info.SiteId = this.SiteId
		set_info.SiteIndexId = this.SiteIndexId
		set_info.AgencyId = this.AgencyId
		set_info.AddMosaic = this.AddMosaic
		set_info.IsIp = this.IsIp
		set_info.Offer = this.Offer
		count, err = sess.Table(set.TableName()).InsertOne(set_info)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	} else {
		set_info.AddMosaic = this.AddMosaic
		set_info.IsIp = this.IsIp
		set_info.Offer = this.Offer
		count, err = sess.Table(set.TableName()).
			Where("site_id=?", this.SiteId).
			Where("agency_id=?", this.AgencyId).
			Where("site_index_id=?", this.SiteIndexId).
			Where("delete_time=?", 0).
			Cols("site_id,site_index_id,agency_id,offer,add_mosaic,is_ip").
			Update(set_info)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	}
	return count, err
}

//查询总代对会员注册设定
func (*AgencyMemberRegisterDiscountSetBean) GetOneThirdDiscountSet(this *input.ThirdDiscountSet) (*back.ThirdDiscountSetBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	set := new(schema.AgencyMemberRegisterDiscountSet)
	data := new(back.ThirdDiscountSetBack)
	if this.AccountId != 0 {
		sess.Where("agency_id=?", this.AccountId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	ok, err := sess.Table(set.TableName()).
		Select("site_index_id,agency_id,offer,add_mosaic,is_ip").
		Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, ok, err
	}
	return data, ok, err
}

//修改、增加会员注册设定(代理)
func (*AgencyMemberRegisterDiscountSetBean) UpdataThirdSet(this *input.ThirdDiscountUpdata) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	set := new(schema.AgencyMemberRegisterDiscountSet)
	var count int64
	var err error
	has, err := sess.Table(set.TableName()).Where("delete_time=?", 0).
		Where("agency_id=?", this.AgencyId).
		Where("site_index_id=?", this.SiteIndexId).
		Get(set)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	set.SiteId = this.SiteId
	set.SiteIndexId = this.SiteIndexId
	set.AgencyId = this.AgencyId
	set.AddMosaic = this.AddMosaic
	set.IsIp = this.IsIp
	set.Offer = this.Offer
	if !has {
		count, err = sess.Table(set.TableName()).InsertOne(set)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else {
		count, err = sess.Table(set.TableName()).Where("site_id=?", this.SiteId).
			Where("agency_id=?", this.AgencyId).
			Where("site_index_id=?", this.SiteIndexId).
			Where("delete_time=?", 0).
			Cols("site_id,site_index_id,agency_id,offer,add_mosaic,is_ip").
			Update(set)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	}
	return count, err
}

//股东查询总代对会员注册设定
func (*AgencyMemberRegisterDiscountSetBean) GetOneThirdOtherDiscountSet(this *input.ThirdDiscountSet) (
	*back.ThirdDiscountSetBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	set := new(schema.AgencyMemberRegisterDiscountSet)
	agency := new(schema.AgencyCount)
	data := new(back.ThirdDiscountSetBack)
	if this.AccountId != 0 {
		sess.Where("agency_id=?", this.AccountId)
	}
	has, err := sess.Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	if !has {
		return nil, false, err
	}
	if agency.AgencyId != 0 {
		sess.Where("agency_id=?", this.AccountId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	ok, err := sess.Table(set.TableName()).
		Select("site_index_id,agency_id,offer,add_mosaic,is_ip").
		Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, ok, err
	}
	return data, ok, err
}

//股东修改、增加会员注册设定(代理)
func (*AgencyMemberRegisterDiscountSetBean) UpdataThirdOtherSet(this *input.ThirdDiscountUpdata) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	set := new(schema.AgencyMemberRegisterDiscountSet)
	var count int64
	var err error
	has, err := sess.Table(set.TableName()).
		Where("delete_time=?", 0).
		Where("agency_id=?", this.AgencyId).
		Where("site_index_id=?", this.SiteIndexId).Get(set)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	set.SiteId = this.SiteId
	set.SiteIndexId = this.SiteIndexId
	set.AgencyId = this.AgencyId
	set.AddMosaic = this.AddMosaic
	set.IsIp = this.IsIp
	set.Offer = this.Offer
	if !has {
		count, err = sess.Table(set.TableName()).InsertOne(set)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else {
		count, err = sess.Table(set.TableName()).
			Where("site_id=?", this.SiteId).
			Where("agency_id=?", this.AgencyId).
			Where("site_index_id=?", this.SiteIndexId).
			Where("delete_time=?", 0).
			Cols("site_id,site_index_id,agency_id,offer,add_mosaic,is_ip").
			Update(set)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	}
	return count, err
}

//查询站点对会员注册优惠
func (*AgencyMemberRegisterDiscountSetBean) GetSiteMemberReg(this *input.SiteMemberReg) (data *back.HolderRegBack, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	amrds := new(schema.AgencyMemberRegisterDiscountSet)
	data = new(back.HolderRegBack)
	has, err = sess.Table(amrds.TableName()).Where("delete_time=?", 0).Where("agency_id=?", 0).Where("site_id=?", this.Site).Where("site_index_id=?", this.SiteIndex).Get(data)
	return
}

//修改站点对会员注册优惠（表里无数据就添加）
func (*AgencyMemberRegisterDiscountSetBean) UpdataSiteMemberReg(this *input.UpdataSiteMemberReg) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	amrds := new(schema.AgencyMemberRegisterDiscountSet)
	sess.Begin()
	data := new(back.HolderRegBack)
	has, err := sess.Table(amrds.TableName()).Where("delete_time=?", 0).Where("agency_id=?", 0).Where("site_id=?", this.Site).Where("site_index_id=?", this.SiteIndex).Get(data)
	if err != nil {
		sess.Rollback()
		return
	}
	//添加
	if !has {
		amrds.AddMosaic = this.AddMosaic
		amrds.Offer = this.Offer
		amrds.AgencyId = 0
		amrds.IsIp = this.IsIp
		amrds.SiteId = this.Site
		amrds.SiteIndexId = this.SiteIndex
		count, err = sess.Table(amrds.TableName()).InsertOne(amrds)
		if err != nil {
			sess.Rollback()
			return
		}
	} else {
		//更新
		amrds.AddMosaic = this.AddMosaic
		amrds.Offer = this.Offer
		amrds.IsIp = this.IsIp
		count, err = sess.Table(amrds.TableName()).Where("site_id=?", this.Site).Where("site_index_id=?", this.SiteIndex).Where("agency_id=?", 0).Cols("offer,add_mosaic,is_ip").Update(amrds)
		if err != nil {
			sess.Rollback()
			return
		}
	}
	if this.IsClear == 1 {
		var reg_info []schema.AgencyMemberRegisterDiscountSet
		err = sess.Table(amrds.TableName()).Where("delete_time=?", 0).Where("site_id=?", this.Site).Where("agency_id!=?", 0).Find(&reg_info)
		if err != nil {
			sess.Rollback()
			return
		}
		var ids []int64
		for _, v := range reg_info {
			ids = append(ids, v.AgencyId)
		}
		amrd := new(schema.AgencyMemberRegisterDiscountSet)
		amrd.Offer = 0
		amrd.AddMosaic = 0
		amrd.IsIp = 0
		count, err = sess.Table(amrds.TableName()).In("agency_id", ids).Where("delete_time=?", 0).Cols("offer,add_mosaic,is_ip").Update(amrd)
		if err != nil {
			sess.Rollback()
			return
		}

	}
	sess.Commit()
	return
}
