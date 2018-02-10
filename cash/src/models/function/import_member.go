package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type ImportMemberBean struct{}

//查询会员是否存在
func (*ImportMemberBean) GetMemberAccount(siteId, account string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	has, err := sess.Where("delete_time = 0").Where("account=?", account).
		Where("site_id=?", siteId).Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//添加导入的会员
func (*ImportMemberBean) AddMember(im []input.ImportMember) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var count int64
	var err error
	sess.Begin()
	for i := range im {
		//会员表
		member := new(schema.Member)
		//会员银行表
		mb := new(schema.MemberBank)
		//会员资料表
		mi := new(schema.MemberInfo)
		//给会员表赋值
		member.SiteId = im[i].SiteId
		member.SiteIndexId = im[i].SiteIndexId
		member.LevelId = im[i].LevelId
		member.FirstAgencyId = im[i].FirstAgencyId
		member.SecondAgencyId = im[i].SecondAgencyId
		member.ThirdAgencyId = im[i].ThirdAgencyId
		member.Account = im[i].Account
		member.Balance = im[i].Money
		member.Realname = im[i].UserName
		member.PcStatus = 2
		member.WapStatus = 2
		member.IosStatus = 2
		member.AndroidStatus = 2
		if im[i].Password == "" { //如果密码为空，则默认密码为123456
			im[i].Password = "123456"
		}
		//密码加密
		md5Password, err := global.MD5ByStr(im[i].Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return 0, err
		}
		member.Password = md5Password
		member.IsImport = 1
		member.Status = 1
		member.CreateTime = global.GetCurrentTime()
		//给会员表插入数据
		count, err = sess.Table(member.TableName()).InsertOne(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		//给会员银行表赋值
		mb.MemberId = member.Id
		mb.BankId = im[i].PayCard
		mb.Card = im[i].PayNum
		mb.CardName = im[i].UserName
		mb.CreateTime = global.GetCurrentTime()
		//给会员银行表插入数据
		count, err = sess.Table(mb.TableName()).InsertOne(mb)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		//给会员资料表赋值
		mi.MemberId = member.Id
		//给会员资料表插入数据
		count, err = sess.Table(mi.TableName()).InsertOne(mi)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		//给股东下加一个会员
		sql := "UPDATE `sales_agency_count` SET `member_count` = member_count + ? WHERE agency_id=?"
		_, err = sess.Exec(sql, 1, member.FirstAgencyId)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		//给总代下加一个会员
		sql1 := "UPDATE `sales_agency_count` SET `member_count` = member_count + ? WHERE agency_id=?"
		_, err = sess.Exec(sql1, 1, member.SecondAgencyId)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		//给代理下加一个会员
		sql2 := "UPDATE `sales_agency_count` SET `member_count` = member_count + ? WHERE agency_id=?"
		_, err = sess.Exec(sql2, 1, member.ThirdAgencyId)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
	}
	sess.Commit()
	return count, err
}

//层级下拉
func (*ImportMemberBean) LevelDrop(this *input.ImportSite) ([]back.MemberLevelDrops, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.MemberLevelDrops
	ml := new(schema.MemberLevel)
	err := sess.Table(ml.TableName()).Where("delete_time=0").
		Where("site_id=?", this.SiteId).Where("site_index_id=?", this.SiteIndexId).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//股东下拉
func (*ImportMemberBean) FirstAgencyDrop(this *input.ImportSite) ([]back.AgencyDrop, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.AgencyDrop
	agency := new(schema.Agency)
	err := sess.Table(agency.TableName()).Where("status=?", 1).Where("delete_time=0").
		Where("level = 2").Where("is_sub = 2").Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//总代下拉
func (*ImportMemberBean) SecondAgencyDrop(this *input.ImportAgency) ([]back.AgencyDrop, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.AgencyDrop
	agency := new(schema.Agency)
	err := sess.Table(agency.TableName()).Where("status=?", 1).Where("delete_time=0").
		Where("level = 3").Where("is_sub = 2").Where("site_id=?", this.SiteId).
		Where("parent_id=?", this.AgencyId).Where("site_index_id=?", this.SiteIndexId).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//代理下拉
func (*ImportMemberBean) ThirdAgencyDrop(this *input.ImportAgency) ([]back.AgencyDrop, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.AgencyDrop
	agency := new(schema.Agency)
	err := sess.Table(agency.TableName()).Where("status=?", 1).Where("delete_time=0").
		Where("level = 4").Where("is_sub = 2").Where("site_id=?", this.SiteId).
		Where("parent_id=?", this.AgencyId).Where("site_index_id=?", this.SiteIndexId).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查看站点是否存在
func (*ImportMemberBean) IsExistSite(siteId, siteIndexId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	has, err := sess.Where("delete_time = 0").Where("index_id=?", siteIndexId).
		Where("id=?", siteId).Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看层级是否存在
func (*ImportMemberBean) IsExistMemberLevel(siteId, siteIndexId, levelId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	has, err := sess.Where("delete_time = 0").Where("site_index_id=?", siteIndexId).
		Where("site_id=?", siteId).Where("level_id=?", levelId).Get(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看股东是否存在
func (*ImportMemberBean) IsExistFirstAgency(siteId, siteIndexId, agencyId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err := sess.Where("delete_time = 0").Where("site_index_id=?", siteIndexId).
		Where("site_id=?", siteId).Where("id=?", agencyId).Where("level=2").
		Where("is_sub=2").Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看总代是否存在
func (*ImportMemberBean) IsExistSecondAgency(siteId, siteIndexId, firstAgencyId, secondAgencyId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err := sess.Where("delete_time = 0").Where("site_index_id=?", siteIndexId).
		Where("site_id=?", siteId).Where("id=?", secondAgencyId).Where("level=3").
		Where("is_sub=2").Where("parent_id=?", firstAgencyId).Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看代理是否存在
func (*ImportMemberBean) IsExistThirdAgency(siteId, siteIndexId, secondAgencyId, thirdAgencyId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err := sess.Where("delete_time = 0").Where("site_index_id=?", siteIndexId).
		Where("site_id=?", siteId).Where("id=?", thirdAgencyId).Where("level=4").
		Where("is_sub=2").Where("parent_id=?", secondAgencyId).Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}
