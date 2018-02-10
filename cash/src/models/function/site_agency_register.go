package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strconv"
	"time"
)

type DistributionApplyBeen struct{}

//删除代理申请（修改删除时间）
func (*DistributionApplyBeen) Delete(this *input.AgentRegState) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	dba := new(schema.SiteAgencyRegister)
	dba.Id = this.RegisterId
	dba.DeleteTime = time.Now().Unix()
	count, err := sess.Table(dba.TableName()).Where("id = ?", dba.Id).
		Cols("delete_time").Update(dba)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//审核代理申请
func (*DistributionApplyBeen) Update(this *input.AgentRegEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAgencyReg := new(schema.SiteAgencyRegister)
	acy := new(schema.Agency)
	agency := new(schema.Agency)
	sct := new(schema.SiteCount)
	siteCount := new(schema.SiteCount)
	agencyCount := new(schema.AgencyCount) //添加代理使用
	ac := new(schema.AgencyCount)          //修改总代操作使用
	ae := new(schema.AgencyCount)          //获取first_id使用
	ag := new(schema.AgencyCount)          //总代获取使用
	age := new(schema.AgencyCount)         //股东获取使用
	a := new(schema.AgencyCount)           //修改股东操作使用
	//给代理申请表赋值
	siteAgencyReg.Id = this.RegisterId
	siteAgencyReg.Status = 1
	siteAgencyReg.UpdateTime = time.Now().Unix()
	//给账号表赋值(审核通过就新增账号)
	agency.SiteId = this.SiteId
	agency.SiteIndexId = this.SiteIndexId
	agency.Account = this.Account
	agency.Password = this.Password
	agency.ParentId = this.ParentId
	agency.Username = this.Username
	agency.IsLogin = 2
	agency.Status = 1
	agency.RoleId = 4
	agency.Level = 4
	agency.IsSub = 2
	//给agency_count表赋值
	agencyCount.SiteId = this.SiteId
	agencyCount.SiteIndexId = this.SiteIndexId
	agencyCount.SecondId = this.ParentId
	sess.Begin()
	//查询站点下是否存在股东账号，不存在则新增的为默认股东
	has, err := sess.Table(acy.TableName()).Where("site_index_id = ?", this.SiteIndexId).Where("role_id = ?", 4).Get(acy)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	if !has {
		agency.IsDefault = 1
	} else {
		agency.IsDefault = 2
	}
	count, err = sess.Table(agency.TableName()).InsertOne(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	//获取agencyCount.FirstId
	sess.Table(ae.TableName()).Where("agency_id = ?", this.ParentId).Select("first_id").Get(ae)
	agencyCount.AgencyId = agency.Id
	agencyCount.FirstId = ae.FirstId
	count, err = sess.Table(agencyCount.TableName()).InsertOne(agencyCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	ac.AgencyId = this.ParentId
	//获取总代的代理人数
	_, err = sess.Table(ag.TableName()).Where("agency_id = ?", ac.AgencyId).Select("third_count").Get(ag)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	ac.ThirdCount = ag.ThirdCount + 1
	//修改总代
	count, err = sess.Table(ac.TableName()).Where("agency_id = ?", ac.AgencyId).Cols("third_count").Update(ac)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	if count == 0 {
		sess.Rollback()
		return
	}
	age.AgencyId = this.ParentId
	//获取股东id
	_, err = sess.Table(age.TableName()).Where("agency_id = ?", age.AgencyId).Select("first_id,third_count").Get(age)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	a.AgencyId = age.FirstId
	a.ThirdCount = ac.ThirdCount
	//修改股东
	count, err = sess.Table(a.TableName()).Where("agency_id = ?", a.AgencyId).Cols("third_count").Update(a)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	if count == 0 {
		sess.Rollback()
		return
	}
	//获取site_count中的总代数量
	_, err = sess.Table(sct.TableName()).Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Select("third_agency_count").Get(sct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	siteCount.ThirdAgencyCount = sct.ThirdAgencyCount + 1
	count, err = sess.Table(siteCount.TableName()).
		Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Cols("third_agency_count").Update(siteCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	if count == 0 {
		sess.Rollback()
		return
	}
	siteAgencyReg.AgencyId = agency.Id
	count, err = sess.Table(siteAgencyReg.TableName()).
		Where("id = ?", siteAgencyReg.Id).
		Cols("status,agency_id,update_time").Update(siteAgencyReg)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	if count == 0 {
		sess.Rollback()
		return
	}
	sess.Commit()
	return
}

//查看代理申请
func (*DistributionApplyBeen) Get(this *input.AgencyIndex, listParams *global.ListParams) (
	[]back.AgentIndex, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAgencyReg := new(schema.SiteAgencyRegister)
	var data []back.AgentIndex
	//判断并组合where条件
	sess.Where("site_id = ?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Status == 2 || this.Status == 1 {
		sess.Where("status = ?", this.Status)
	}
	if this.Key == "email" {
		sess.Where("email = ?", this.Value)
	} else if this.Key == "qq" {
		sess.Where("qq = ?", this.Value)
	} else if this.Key == "skype" {
		sess.Where("skype = ?", this.Value)
	} else if this.Key == "wechat" {
		sess.Where("wechat = ?", this.Value)
	}
	sess.Where("delete_time = ?", 0)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	//获得分页记录
	listParams.Make(sess)
	//重新传入表名和where条件查询记录
	err := sess.Table(siteAgencyReg.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	//获得符合条件的记录数
	count, err := sess.Table(siteAgencyReg.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//站点是否存在(代理申请表)
func (*DistributionApplyBeen) SiteIdExists(siteId, siteIndexId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	dtba := new(schema.SiteAgencyRegister)
	has, err := sess.Table(dtba.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Get(dtba)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看agency_id
func (*DistributionApplyBeen) IsAgencyId(id int64) (*schema.SiteAgencyRegister, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sar := new(schema.SiteAgencyRegister)
	_, err := sess.Table(sar.TableName()).
		Where("id = ?", id).Select("agency_id").Get(sar)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return sar, err
	}
	return sar, err
}

//添加代理申请
func (*DistributionApplyBeen) Add(this *input.AgencyRegister) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAgencyReg := new(schema.SiteAgencyRegister)
	//给代理申请表赋值
	pass, _ := global.MD5ByStr(this.Password, global.EncryptSalt)
	BankId, _ := strconv.ParseInt(this.BankId, 10, 64)
	Province, _ := strconv.ParseInt(this.Province, 10, 64)
	Zone, _ := strconv.ParseInt(this.Zone, 10, 64)
	siteAgencyReg.SiteId = this.SiteId
	siteAgencyReg.SiteIndexId = this.SiteIndexId
	siteAgencyReg.Phone = this.Phone
	siteAgencyReg.Qq = this.Qq
	siteAgencyReg.Email = this.Email
	siteAgencyReg.Card = this.Card
	siteAgencyReg.UsName = this.EnglishNickname
	siteAgencyReg.ZhName = this.ChineseNickname
	siteAgencyReg.Account = this.AgencyAccount
	siteAgencyReg.Password = pass
	siteAgencyReg.Status = 2
	siteAgencyReg.UserName = this.UserName
	siteAgencyReg.BackAccount = this.BackAccount
	siteAgencyReg.BankId = BankId
	siteAgencyReg.OtherMethod = this.OtherMethod
	siteAgencyReg.PromoteWebsite = this.PromoteWebsite
	siteAgencyReg.Province = Province
	siteAgencyReg.Zone = Zone
	count, err = sess.InsertOne(siteAgencyReg)
	return
}

//代理账号是否已存在
func (*DistributionApplyBeen) IsExistAccount(account string) (has bool, have bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	siteAgencyRegister := new(schema.SiteAgencyRegister)
	has, err = sess.Where("account = ?", account).Get(agency)
	if err != nil {
		return
	}
	have, err = sess.Where("account = ?", account).Get(siteAgencyRegister)
	return
}
