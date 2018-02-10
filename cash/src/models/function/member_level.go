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

type MemberLevelBean struct{}

//添加会员层级
func (*MemberLevelBean) Add(this *input.MemberLevel) (int64, error) {
	schemaMemLev := &schema.MemberLevel{
		SiteId:       this.SiteId,
		SiteIndexId:  this.SiteIndexId,
		DepositCount: this.DepositCount,
		DepositNum:   this.DepositNum,
		Description:  this.Description,
		Remark:       this.Remark,
		LevelId:      this.LevelId,
		IsSelfRebate: 2,
		PaySetId:     this.PaySetId}
	sess := global.GetXorm().NewSession().Table(schemaMemLev.TableName())
	defer sess.Close()
	exist := new(schema.MemberLevel)
	have, err := sess.Where("site_id=?", this.SiteId).And("site_index_id=?", this.SiteIndexId).And("is_default=?", 1).Get(exist)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if have {
		schemaMemLev.IsDefault = 2
	} else {
		schemaMemLev.IsDefault = 1
	}

	schemaMemLev.StartTime = marshalTime(this.StartTime)

	schemaMemLev.EndTime = marshalTime(this.EndTime)
	count, err := sess.InsertOne(schemaMemLev)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//解析2017-10-01这种类型的时间字符串成时间戳(int64)
func marshalTime(timeStr string) int64 {
	loc, _ := time.LoadLocation("UTC")
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return time.Unix()
}

//获取层级信息(返回的信息只包含返回给前端页面显示的内容)
func (*MemberLevelBean) LevelGet(levelId, siteId, siteIndexId string) (*back.MemberLevel, bool, error) {
	schemaMemLev := new(schema.MemberLevel)
	sess := global.GetXorm().NewSession().Table(schemaMemLev.TableName())
	defer sess.Close()
	backMemLev := new(back.MemberLevel)
	sess.Where("level_id=?", levelId)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	have, err := sess.Get(backMemLev)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backMemLev, have, err
	}
	return backMemLev, have, err
}

//判断层级是否为默认层级
func (*MemberLevelBean) LevelDefault(levelId, siteId, siteIndexId string) (bool, error) {
	schemaMemLev := new(schema.MemberLevel)
	sess := global.GetXorm().Table(schemaMemLev.TableName())
	defer sess.Close()
	backMemLev := new(back.MemberLevel)
	sess.Where("level_id=?", levelId)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("is_default=?", 1)
	have, err := sess.Get(backMemLev)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return have, err
	}
	return have, err
}

//更新会员层级
func (*MemberLevelBean) Update(this *input.MemberLevelUpdate) (int64, error) {
	var err error
	schemaMemberLev := &schema.MemberLevel{
		DepositCount: this.DepositCount,
		Description:  this.Description,
		DepositNum:   this.DepositNum,
		Remark:       this.Remark,
		LevelId:      this.NewLevelId}
	schemaMemberLev.StartTime = marshalTime(this.StartTime)
	schemaMemberLev.EndTime = marshalTime(this.EndTime)
	sess := global.GetXorm().NewSession().Table(schemaMemberLev.TableName())
	defer sess.Close()
	sess.Begin()
	sess.Where("site_id=?", this.SiteId)
	sess.Where("level_id=?", this.OldLevelId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	member := new(schema.Member)
	//更新层级表中的数据
	count, err := sess.Omit("site_id", "site_index_id").Update(schemaMemberLev)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	member.LevelId = this.NewLevelId
	//获取当前站点,站点前台Id,层级Id下的会员的层级Id
	count, err = sess.Table(member.TableName()).Where(conds).Cols("level_id").Update(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	sess.Commit()
	return count, err
}

//获取会员层级列表
func (*MemberLevelBean) List(this *input.LevelIndex, listParams *global.ListParams) ([]back.MemberLevelList, int64, error) {
	memberLev := new(schema.MemberLevel)
	var data []back.MemberLevelList
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	sess.Where("delete_time=?", 0)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.Remark != "" {
		sess.Where("remark=?", this.Remark)
	}
	if this.Description != "" {
		sess.Where("description=?", this.Description)
	}
	if this.LevelId != "" {
		sess.Where("Level_id=?", this.LevelId)
	}
	conds := sess.Conds()
	listParams.Make(sess)
	sess.Asc("is_default")
	err := sess.Table(memberLev.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(memberLev.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取站点会员层级列表
func (*MemberLevelBean) SiteLevelList(this *input.LevelIndex, listParams *global.ListParams) (data []*back.SiteLevelList, count int64, err error) {
	memberLev := new(schema.MemberLevel)
	sess := global.GetXorm().NewSession().Table(memberLev.TableName())
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	sess.Where("delete_time=?", 0)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	conds := sess.Conds()
	listParams.Make(sess)
	err = sess.Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(memberLev.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//开启层级自助返水功能
func (*MemberLevelBean) SelfRebate(this *input.MemberLevelSelfRebate) (int64, error) {
	schemaMemberLev := new(schema.MemberLevel)
	sess := global.GetXorm().NewSession().Table(schemaMemberLev.TableName())
	defer sess.Close()
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("level_id=?", this.LevelId)
	schemaMemberLev.IsSelfRebate = this.IsSelfRebate
	count, err := sess.Cols("is_self_rebate").Update(schemaMemberLev)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//回归会员层级
func (*MemberLevelBean) ComeBackLevel(this *input.LevelInfoGet) (int64, error) {
	member := new(schema.Member)
	var count int64
	members := make([]schema.Member, 0)
	sess := global.GetXorm().NewSession().Table(member.TableName())
	//查询site_id,site_index_id,level_id没有被锁定的会员列表.
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("level_id=?", this.LevelId)
	sess.Where("is_locked_level=?", 2)
	err := sess.Select("id").Find(&members)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if len(members) == 0 {
		return 0, err
	}
	level, err := defaultInfo(this.SiteId, this.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	//开启事务
	sess.Begin()
	//更新会员的层级
	var i int64
	for index := range members {
		members[index].LevelId = level.LevelId
		sess.Where("id=?", members[index].Id)
		count, err = sess.Table(members[index].TableName()).
			Cols("level_id").Update(members[index])
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if count == 0 {
			return count, err
		}
		i = i + 1
	}
	if i > 0 {
		//更新当前操作层级的人数
		memberLevel := new(schema.MemberLevel)
		sess.Where("site_id=?", this.SiteId)
		sess.Where("site_index_id=?", this.SiteIndexId)
		conds := sess.Conds()
		//获取操作层级的人数
		has, err := sess.Table(memberLevel.TableName()).
			Select("count").Where("level_id=?", this.LevelId).
			Get(memberLevel)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if !has {
			return count, err
		}
		memberLevel.Count -= count
		count, err = sess.Table(memberLevel.TableName()).Where(conds).
			Where("level_id=?", this.LevelId).Cols("count").
			Update(memberLevel)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if count == 0 {
			return count, err
		}
		//更新默认层级的人数
		level.Count += count
		count, err = sess.Table(memberLevel.TableName()).Where(conds).
			Where("level_id=?", level.LevelId).Cols("count").
			Update(level)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		if count == 0 {
			return count, err
		}
	}
	sess.Commit()
	return count, err
}

//查找某个site_id,site_index_id下的默认层级信息
func defaultInfo(siteId, siteIndexId string) (*schema.MemberLevel, error) {
	schemaMemLev := new(schema.MemberLevel)
	sess := global.GetXorm().Table(schemaMemLev.TableName())
	defer sess.Close()
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("is_default=?", 1)
	_, err := sess.Get(schemaMemLev)
	return schemaMemLev, err
}

//会员层级下拉框列表
func (*MemberLevelBean) MemberLevelDrop(this *input.LevelIndex) ([]back.MemberLevelDrop, error) {
	memberLevel := new(schema.MemberLevel)
	var data []back.MemberLevelDrop
	sess := global.GetXorm().NewSession().Table(memberLevel.TableName())
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	err := sess.Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//总后台会员层级下拉
func (*MemberLevelBean) Memberdrop(this *input.MemberLevels) ([]back.MemberLevelDrops, error) {
	memberLevel := new(schema.MemberLevel)
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.MemberLevelDrops
	sess.Where("site_id=?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	err := sess.Table(memberLevel.TableName()).GroupBy("level_id,description").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//入款商户下拉
func (*MemberLevelBean) ThirdPaidList(this *input.MemberLevels) ([]back.ThirdPaidList, error) {
	paidlist := new(schema.OnlinePaidSetup)
	sess := global.GetXorm().NewSession()
	var data []back.ThirdPaidList
	sess.Where("site_id=?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sql1 := "sales_paid_type.id=sales_online_paid_setup.paid_type"
	sql2 := "sales_online_income_third.id=sales_online_paid_setup.paid_platform"
	sess.Select(paidlist.TableName() + ".id as id,sales_online_income_third.title as title," +
		"sales_paid_type.paid_type_name as paid_type_name")
	err := sess.Table(paidlist.TableName()).Join("left", "sales_paid_type", sql1).
		Join("left", "sales_online_income_third", sql2).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取会员详情(列表)
func (*MemberLevelBean) MemberListInfo(this *input.LevelMember, listParams *global.ListParams) ([]back.MemberInfoList, int64, error) {
	member := new(schema.Member)
	var data []back.MemberInfoList
	sess := global.GetXorm().NewSession()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.LevelId != "" {
		sess.Where("level_id=?", this.LevelId)
	}
	if this.Account != "" {
		if this.IsFuzzy == 1 { //模糊查询
			sess.Where("account like ?", "%"+this.Account+"%")
		} else {
			sess.Where("account = ?", this.Account)
		}
	}
	if this.IsLockedLevel != 0 {
		sess.Where("is_locked_level = ?", this.IsLockedLevel)
	}
	if this.AccountList != "" {
		//分割账号
		accounts := strings.Split(this.AccountList, ",")
		sess.In("account", accounts)
	}
	if this.CreateTime != "" {
		loc, _ := time.LoadLocation("Local")
		t1, _ := time.ParseInLocation("2006-01-02", this.CreateTime, loc)
		startTime := t1.Unix()
		dd, _ := time.ParseDuration("24h")
		endTime := t1.Add(dd).Unix()
		sess.Where("create_time>=?", startTime).Where("create_time<?", endTime)
	}
	if this.LastLoginIp != "" {
		sess.Where("login_ip = ?", this.LastLoginIp)
	}
	conds := sess.Conds()
	listParams.Make(sess)
	memberInfo := new(schema.MemberInfo)
	memberCashCount := new(schema.MemberCashCount)
	where1 := fmt.Sprintf("%s.member_id = %s.id", memberInfo.TableName(), member.TableName())
	where2 := fmt.Sprintf("%s.id = %s.member_id", member.TableName(), memberCashCount.TableName())
	err := sess.Table(member.TableName()).Join("LEFT", memberInfo.TableName(), where1).
		Join("LEFT", memberCashCount.TableName(), where2).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(member.TableName()).Join("LEFT", memberInfo.TableName(), where1).
		Join("LEFT", memberCashCount.TableName(), where2).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	for index := range data {
		memberLevel := new(schema.MemberLevel)
		sess.Where("site_id=?", this.SiteId)
		if this.SiteIndexId != "" {
			sess.Where("site_index_id=?", this.SiteIndexId)
		}
		sess.Where("level_id=?", data[index].LevelId)
		sess.Get(memberLevel)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		data[index].IsDefault = memberLevel.IsDefault
	}
	return data, count, err
}

//锁定会员层级
func (*MemberLevelBean) LockMember(this *input.LockMember) (int64, error) {
	member := new(schema.Member)
	sess := global.GetXorm().NewSession().Table(member.TableName())
	defer sess.Close()
	sess.Where("id=?", this.MemberId)
	member.IsLockedLevel = this.Lock
	count, err := sess.Cols("is_locked_level").Update(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//移动分层
func (*MemberLevelBean) MoveLevel(this *input.MoveMemberLevel) (int64, error) {
	//找到移入层级的条件
	var count int64
	memberLevel := new(schema.MemberLevel)
	sess := global.GetXorm().NewSession().Table(memberLevel.TableName())
	defer sess.Close()
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	sess.Where("level_id=?", this.MoveIn)
	_, err := sess.Get(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	/*关联会员信息包括(site_id,site_index_id,level_id,
	is_locked_level)
	还包括会员的注册时间是否在移入层级的起始时间之内。
	*/
	member := new(schema.Member)
	members := make([]schema.Member, 0)
	sess.Table(member.TableName())
	sess.Where(member.TableName()+".site_id=?", this.SiteId)
	sess.Where(member.TableName()+".site_index_id=?", this.SiteIndexId)
	sess.Where(member.TableName()+".level_id=?", this.MoveOut)
	sess.Where(member.TableName()+".is_locked_level=?", 2)
	sess.Where(member.TableName()+".create_time>=?", memberLevel.StartTime)
	sess.Where(member.TableName()+".create_time<=?", memberLevel.EndTime)
	memberCashCount := new(schema.MemberCashCount)
	//关联会员现金表
	where1 := fmt.Sprintf("%s.id = %s.member_id", member.TableName(), memberCashCount.TableName())
	/*
	 会员现金表字段中的存款次数是否满足移入层级的
	 次数
	 会员现金表字段中的存款总额是否满足移入层级的
	 存款次数
	*/
	sess.Where(memberCashCount.TableName()+".deposit_num>=?", memberLevel.DepositNum)
	sess.Where(memberCashCount.TableName()+".deposit_count>=?", memberLevel.DepositCount)
	err = sess.Join("LEFT", memberCashCount.TableName(), where1).Select("id").Find(&members)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if len(members) == 0 {
		return 0, err
	}
	//开启事务,对满足条件的会员的层级Id进行修改
	sess.Begin()
	for index := range members {
		sess.Table(member.TableName()).Where("id=?", members[index].Id)
		members[index].LevelId = this.MoveIn
		_, err = sess.Cols("level_id").Update(members[index])
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		count++
	}
	//对移出组中会员数量字段进行修改
	if count != 0 {
		//对移出组中会员数量字段进行修改
		sess.Where("level_id=?", this.MoveOut)
		levelOut := new(schema.MemberLevel)
		_, err = sess.Table(levelOut.TableName()).Where(conds).Select("count").Get(levelOut)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		levelOut.Count -= count
		sess.Where("level_id=?", this.MoveOut)
		sess.Table(memberLevel.TableName())
		count, err = sess.Where(conds).Cols("count").Update(levelOut)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
		//对移入组中会员数量字段进行修改
		sess.Where("level_id=?", this.MoveIn)
		memberLevel.Count += count
		sess.Table(memberLevel.TableName())
		count, err = sess.Where(conds).Cols("count").Update(memberLevel)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	}
	sess.Commit()
	return count, err
}

//获取会员层级支付设定
func (*MemberLevelBean) MemberLevelPatSetOne(this *input.MemberLevelPaySet) (*back.MemberLevelPaySetBack, bool, error) {
	memberLev := new(schema.MemberLevel)
	sess := global.GetXorm().Table(memberLev.TableName())
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.LevelId != "" {
		sess.Where("level_id=?", this.LevelId)
	}
	data := new(back.MemberLevelPaySetBack)
	has, err := sess.Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//修改会员层级支付方式
func (*MemberLevelBean) UpdataMemberLevelPaySet(this *input.MemberLevelPaySetUpdata) (int64, error) {
	memberLev := new(schema.MemberLevel)
	sess := global.GetXorm().Table(memberLev.TableName())
	defer sess.Close()
	memberLev.PaySetId = this.PaySetId
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("level_id=?", this.LevelId)
	sess.Cols("pay_set_id")
	count, err := sess.Update(memberLev)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取站点的默认分层
func (*MemberLevelBean) SiteDefault(siteId, siteIndexId string) (*schema.MemberLevel, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	schemaMemLev := new(schema.MemberLevel)
	sess.Where("delete_time=?", 0)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("is_default=?", 1)
	flag, err := sess.Get(schemaMemLev)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return schemaMemLev, flag, err
	}
	return schemaMemLev, flag, err
}

//站点管理-层级列表
func (*MemberLevelBean) MemberLevelList(this *input.SiteLevelList) ([]back.SiteMemberLevelList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	var siteMemberLevel []back.SiteMemberLevelList
	err := sess.Table(memberLevel.TableName()).
		Where("site_id=?", this.SiteId).
		Where("delete_time=0").
		Find(&siteMemberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return siteMemberLevel, err
	}
	return siteMemberLevel, err
}

//站点管理-层级详情
func (*MemberLevelBean) MemberLevelInfo(this *input.SiteLevelInfo) (siteMemberLevel *back.SiteMemberLevelInfo, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	siteMemberLevel = new(back.SiteMemberLevelInfo)
	_, err = sess.Table(memberLevel.TableName()).Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).Where("level_id=?", this.LevelId).
		Where("delete_time=0").Get(siteMemberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return siteMemberLevel, err
	}
	return siteMemberLevel, err
}

//站点管理-添加层级
func (*MemberLevelBean) AddMemberLevel(this *input.SiteLevelAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	//判断层级是否存在
	has, err := GetLevel(this.SiteId, this.SiteIndexId)
	if err != nil {
		return
	}
	if has {
		memberLevel.IsDefault = 1
	} else {
		memberLevel.IsDefault = 2
	}
	memberLevel.SiteId = this.SiteId
	memberLevel.SiteIndexId = this.SiteIndexId
	memberLevel.LevelId = this.LevelId
	memberLevel.Description = this.Description
	count, err = sess.Table(memberLevel.TableName()).InsertOne(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点管理-修改层级
func (ml *MemberLevelBean) EditMemberLevel(this *input.SiteLevelEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	memberLevel.Description = this.Description
	count, err = sess.Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("level_id=?", this.LevelId).
		Cols("description").
		Update(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点管理-删除层级
func (ml *MemberLevelBean) DelMemberLevel(this *input.SiteLevelInfo) (err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	memberLevel.DeleteTime = time.Now().Unix()
	sess.Begin()
	count, err := sess.Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("level_id=?", this.LevelId).
		Cols("delete_time").
		Update(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	if count != 0 { //删除层级成功,对应修改会员表
		member := new(schema.Member)
		//获取站点下默认层级
		levelId, has, err := ml.GetDefaultLevel(this.SiteId, this.SiteIndexId)

		if err != nil || has {
			sess.Rollback()
		}
		member.LevelId = levelId
		_, err = sess.Where("site_id = ?", this.SiteId).
			Where("site_index_id=?", this.SiteIndexId).
			Where("level_id=?", this.LevelId).
			Cols("level_id").
			Update(member)
		if err != nil {
			sess.Rollback()
		}
	}
	sess.Commit()
	return
}

//查看站点下会员层级是否存在
func GetLevel(siteId, siteIndexId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	has, err = sess.Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).
		Get(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看站点下站点前台id是否存在
func (*MemberLevelBean) GetSiteIndexId(siteId, siteIndexId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	has, err = sess.Where("id=?", siteId).Where("index_id=?", siteIndexId).
		Where("delete_time=0").Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看站点下站点前台id  level_id是否存在
func (*MemberLevelBean) GetSiteIndexIdLevelId(siteId, siteIndexId, level_id string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.MemberLevel)
	has, err = sess.Where("site_id=?", siteId).Where("level_id=?", level_id).
		Where("site_index_id=?", siteIndexId).Where("delete_time=0").Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看层级是否存在
func (*MemberLevelBean) GetLevel(siteId, siteIndexId, levelId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	has, err = sess.Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).
		Where("level_id=?", levelId).Get(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看层级是否为默认层级
func (*MemberLevelBean) GetLevelIsDefault(siteId, siteIndexId, levelId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	has, err = sess.Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).
		Where("level_id=?", levelId).Where("is_default=1").Get(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看站点下默认层级
func (*MemberLevelBean) GetDefaultLevel(siteId, siteIndexId string) (levelId string, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberLevel := new(schema.MemberLevel)
	has, err = sess.Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Where("is_default=1").
		Select("level_id").
		Get(memberLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return levelId, has, err
	}
	levelId = memberLevel.LevelId
	return levelId, has, err
}

//查看层级下是否有会员
func (*MemberLevelBean) GetMember(siteId, siteIndexId, levelId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	has, err = sess.Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Where("level_id=?", levelId).
		Where("delete_time=0").
		Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//根据levleid获取层级详情
func (*MemberLevelBean) GetLevelInfo(levelId string) (data schema.MemberLevel, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Where("level_id=?", levelId).Get(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, flag, err
	}
	return data, flag, err
}
