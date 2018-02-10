package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type MemberRegisterSettingBean struct{}

//添加会员注册配置
func (*MemberRegisterSettingBean) Add(this *input.MemberRegisterSetting) (int64, error) {
	schemaMemRegSet := &schema.SiteMemberRegisterSet{
		SiteId:      this.SiteId,
		Quota:       this.Quota,
		SiteIndexId: this.SiteIndexId,
		IsReg:       this.IsReg,
		Email:       this.Email,
		Passport:    this.Passport,
		Qq:          this.Qq,
		Wechat:      this.Wechat,
		Mobile:      this.Mobile,
		Birthday:    this.Birthday,
		IsName:      this.IsName,
		IsShowName:  this.IsShowName,
		IsWechat:    this.IsWechat,
		IsCardReply: this.IsCardReply,
		IsTel:       this.IsTel,
		IsWapSingle: this.IsWapSingle,
		IsEmail:     this.IsEmail,
		IsQq:        this.IsQq,
		IsCode:      this.IsCode,
		TryPlay:     this.TryPlay,
		Offer:       this.Offer,
		AddMosaic:   this.AddMosaic,
		IsIp:        this.IsIp}

	sess := global.GetXorm().NewSession().Table(schemaMemRegSet.TableName())
	defer sess.Close()
	count, err := sess.Insert(schemaMemRegSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看会员注册配置
func (*MemberRegisterSettingBean) Get(this *input.MemberRegisterSettingGet) (*back.MemberRegisterSetting, bool, error) {
	schemaMemRegSet := new(schema.SiteMemberRegisterSet)
	sess := global.GetXorm().NewSession().Table(schemaMemRegSet.TableName())
	defer sess.Close()
	backMemRegSet := new(back.MemberRegisterSetting)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	have, err := sess.Get(backMemRegSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backMemRegSet, have, err
	}
	return backMemRegSet, have, err
}

//修改会员注册配置
func (*MemberRegisterSettingBean) Update(this *input.MemberRegisterSetting) (int64, error) {
	schemaMemRegSet := &schema.SiteMemberRegisterSet{
		SiteId:      this.SiteId,
		Quota:       this.Quota,
		SiteIndexId: this.SiteIndexId,
		IsReg:       this.IsReg,
		Email:       this.Email,
		Passport:    this.Passport,
		Qq:          this.Qq,
		Wechat:      this.Wechat,
		Mobile:      this.Mobile,
		Birthday:    this.Birthday,
		IsName:      this.IsName,
		IsShowName:  this.IsShowName,
		IsWechat:    this.IsWechat,
		IsCardReply: this.IsCardReply,
		IsTel:       this.IsTel,
		IsWapSingle: this.IsWapSingle,
		IsEmail:     this.IsEmail,
		IsQq:        this.IsQq,
		IsCode:      this.IsCode,
		TryPlay:     this.TryPlay,
		Offer:       this.Offer,
		AddMosaic:   this.AddMosaic,
		IsIp:        this.IsIp}
	sess := global.GetXorm().NewSession().Table(schemaMemRegSet.TableName())
	defer sess.Close()
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	count, err := sess.Omit("site_id,site_index_id").Update(schemaMemRegSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取会员注册设置
func (*MemberRegisterSettingBean) GetOneSet(site, siteIndex string) (schema.SiteMemberRegisterSet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.SiteMemberRegisterSet
	sess.Where("site_id=?", site)
	sess.Where("site_index_id=?", siteIndex)
	flag, err := sess.Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return info, flag, err
}
