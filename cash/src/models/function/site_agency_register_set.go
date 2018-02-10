package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SecondDistributionRegisterSetupBeen struct{}

//添加代理注册设定
func (*SecondDistributionRegisterSetupBeen) Add(this *input.AgentSetDo) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ssdrs := new(schema.SiteAgencyRegisterSet)
	ssdrs.SiteId = this.SiteId
	ssdrs.SiteIndexId = this.SiteIndexId
	ssdrs.RegisterProxy = this.RegisterProxy
	ssdrs.ChineseNickname = this.ChineseNickname
	ssdrs.EnglishNickname = this.EnglishNickname
	ssdrs.PromoteWebsite = this.PromoteWebsite
	ssdrs.OtherMethod = this.OtherMethod
	ssdrs.IsMustChineseNickname = this.IsMustChineseNickname
	ssdrs.IsMustEmail = this.IsMustEmail
	ssdrs.IsMustEnglishNickname = this.IsMustEnglishNickname
	ssdrs.IsMustIdentity = this.IsMustIdentity
	ssdrs.IsMustMethod = this.IsMustMethod
	ssdrs.IsMustPhone = this.IsMustPhone
	ssdrs.IsMustPromoteWebsite = this.IsMustPromoteWebsite
	ssdrs.IsMustQq = this.IsMustQq
	ssdrs.NeedCard = this.NeedCard
	ssdrs.NeedEmail = this.NeedEmail
	ssdrs.NeedPhone = this.NeedPhone
	ssdrs.NeedQq = this.NeedQq
	count, err := sess.Table(ssdrs.TableName()).InsertOne(ssdrs)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改代理注册设定
func (*SecondDistributionRegisterSetupBeen) Update(this *input.AgentSetDo) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ssdrs := new(schema.SiteAgencyRegisterSet)
	ssdrs.SiteId = this.SiteId
	ssdrs.SiteIndexId = this.SiteIndexId
	ssdrs.RegisterProxy = this.RegisterProxy
	ssdrs.ChineseNickname = this.ChineseNickname
	ssdrs.EnglishNickname = this.EnglishNickname
	ssdrs.PromoteWebsite = this.PromoteWebsite
	ssdrs.OtherMethod = this.OtherMethod
	ssdrs.IsMustChineseNickname = this.IsMustChineseNickname
	ssdrs.IsMustEmail = this.IsMustEmail
	ssdrs.IsMustEnglishNickname = this.IsMustEnglishNickname
	ssdrs.IsMustIdentity = this.IsMustIdentity
	ssdrs.IsMustMethod = this.IsMustMethod
	ssdrs.IsMustPhone = this.IsMustPhone
	ssdrs.IsMustPromoteWebsite = this.IsMustPromoteWebsite
	ssdrs.IsMustQq = this.IsMustQq
	ssdrs.NeedCard = this.NeedCard
	ssdrs.NeedEmail = this.NeedEmail
	ssdrs.NeedPhone = this.NeedPhone
	ssdrs.NeedQq = this.NeedQq
	count, err := sess.Table(ssdrs.TableName()).Where("site_id = ?", ssdrs.SiteId).Where("site_index_id = ?", ssdrs.SiteIndexId).
		Cols("register_proxy", "chinese_nickname", "english_nickname", "need_cardnumber",
			"need_email", "need_qq", "need_phone", "promote_website", "other_method",
			"is_must_chinese_nickname", "is_must_english_nickname", "is_must_email", "is_must_identity",
			"is_must_qq", "is_must_phone", "is_must_promote_website", "is_must_method").Update(ssdrs)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看代理注册设定
func (*SecondDistributionRegisterSetupBeen) SiteIdExist(siteId, siteIndexId string) (
	*back.SiteAgencyRegisterSet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAgencyRegisterSet := new(schema.SiteAgencyRegisterSet)
	siteAgencyRegisterSett := new(back.SiteAgencyRegisterSet)
	has, err := sess.Table(siteAgencyRegisterSet.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Get(siteAgencyRegisterSett)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return siteAgencyRegisterSett, has, err
	}
	return siteAgencyRegisterSett, has, err
}

//查询一条代理申请
func (*SecondDistributionRegisterSetupBeen) GetOneAgencyReg(this *input.OneAgencyReg) (*back.AgentIndex, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	sar := new(schema.SiteAgencyRegister)
	data := new(back.AgentIndex)
	has, err := sess.Table(sar.TableName()).Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}
