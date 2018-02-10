package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

//会员个人资料操作
type MemberSelfInfoBean struct{}

//查询登录会员的个人资料
func (*MemberSelfInfoBean) GetMemberInfoSelf(id int64) (*back.MemberSelfInfoBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mem := new(schema.Member)
	mi := new(schema.MemberInfo)
	data := new(back.MemberSelfInfoBack)
	sess.Where(mem.TableName()+".delete_time=?", 0)
	sess.Where(mem.TableName()+".status=?", 1)
	has, err := sess.Table(mem.TableName()).
		Join("LEFT", mi.TableName(), mem.TableName()+".id="+mi.TableName()+".member_id").
		Where(mem.TableName()+".id=?", id).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//会员邮箱添加或者修改
func (*MemberSelfInfoBean) EmailChangeOrAdd(this *input.EmailAddOrChangeIn) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	mi := new(schema.MemberInfo)
	//先判断member_info中是否有该会员的基本信息
	has, err := sess.Where("member_id=?", this.Id).Get(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	mi.Email = this.Email
	//添加
	if !has {
		mi.MemberId = this.Id
		count, err := sess.InsertOne(mi)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		sess.Commit()
		return count, err
	} else {
		//修改
		count, err := sess.Where("member_id=?", this.Id).Cols("email").Update(mi)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		sess.Commit()
		return count, err
	}
}

//会员出生日期添加或者修改
func (*MemberSelfInfoBean) BirthChangeOrAdd(this *input.BirthAddOrChangeIn) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	mi := new(schema.MemberInfo)
	//先判断member_info中是否有该会员的基本信息
	has, err := sess.Where("member_id=?", this.Id).Get(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	tineUnix, err := time.Parse("2006-01-02", this.Birth)
	mi.Birthday = tineUnix.Unix()
	//添加
	if !has {
		mi.MemberId = this.Id
		count, err := sess.InsertOne(mi)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		sess.Commit()
		return count, err
	}
	//修改
	count, err := sess.Where("member_id=?", this.Id).Cols("birthday").Update(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	sess.Commit()
	return count, err
}

//会员手机号添加或者修改
func (*MemberSelfInfoBean) PhoneChangeOrAdd(this *input.PhoneBind) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	mi := new(schema.MemberInfo)
	//先判断member_info中是否有该会员的基本信息
	has, err := sess.Where("member_id=?", this.Id).Get(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	mi.LocalCode = this.LocalCode
	mi.Mobile = this.Phone
	//添加
	if !has {
		mi.MemberId = this.Id
		count, err := sess.InsertOne(mi)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		sess.Commit()
		return count, err
	}
	//修改
	count, err := sess.Where("member_id=?", this.Id).Cols("local_code,mobile").Update(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	sess.Commit()
	return count, err
}

//会员中心主页
func (*MemberSelfInfoBean) MemberHomePage(id int64) (*back.MemberHomePageBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.Member)
	mp := new(schema.MemberProductClassifyBalance)
	ml := new(schema.MemberLevel)
	//所需要搜索的字段
	sess.Select(mb.TableName() + ".id," + mb.TableName() + ".realname," +
		mb.TableName() + ".balance," +
		mb.TableName() + ".account," +
		"SUM(" + mp.TableName() + ".balance) as a," +
		ml.TableName() + ".is_self_rebate")
	//状态开启，删除时间为0
	sess.Where(mb.TableName()+".delete_time=?", 0)
	sess.Where(mb.TableName()+".status=?", 1)
	//group by
	sess.GroupBy(mb.TableName() + ".id")
	info := new(back.MemberHomePageBack)
	has, err := sess.Table(mb.TableName()).
		Join("LEFT", mp.TableName(), mb.TableName()+".id="+mp.TableName()+".member_id").
		Where(mb.TableName()+".id=?", id).
		Join("LEFT", ml.TableName(), mb.TableName()+".level_id="+ml.TableName()+
			".level_id AND "+mb.TableName()+".site_id="+ml.TableName()+
			".site_id AND "+mb.TableName()+".site_index_id="+ml.TableName()+
			".site_index_id").Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err
	}
	return info, has, err
}

//根据id查询一条会员信息
func (*MemberSelfInfoBean) MemberOneInfo(id int64) (*schema.Member, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info := new(schema.Member)
	has, err := sess.Where("id=?", id).
		Where("delete_time=?", 0).
		Where("status=?", 1).Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err

	}
	return info, has, err
}

//修改密码
func (*MemberSelfInfoBean) MemberSelfPassword(this *input.PasswordMemberChange, id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.Member)
	if this.Type == 1 {
		mb.Password = this.Password
		sess.Cols("password")
	} else if this.Type == 2 {
		mb.DrawPassword = this.Password
		sess.Cols("draw_password")
	}
	count, err := sess.Where("id=?", id).Update(mb)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取会员消息列表
func (*MemberSelfInfoBean) GetMesList(siteId, siteIndexId string, Id int64, times *global.Times, listparams *global.ListParams) (data []back.MemberMessage, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ms := new(schema.MemberMessage)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("member_id=?", Id)
	if times != nil {
		times.Make("create_time", sess)
	}
	sess.Where("delete_time=?", 0)
	sess.OrderBy("id desc")
	conds := sess.Conds()
	listparams.Make(sess)
	err = sess.Table(ms.TableName()).Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(ms.TableName()).Where(conds).Count()
	return
}

//消息状态修改
func (*MemberSelfInfoBean) PutMesStatus(siteId, siteIndexId string, mes *input.WapMesStatus) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ms := new(schema.MemberMessage)
	mes.State = 2
	data, err = sess.Table(ms.TableName()).Where("id=?", mes.Id).Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).Update(mes)
	return
}

//检测支付密码
func (*MemberSelfInfoBean) CheckDrawPass(this *input.CheckDrawPass) (*back.DrawPassData, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mm := new(schema.Member)
	data := new(back.DrawPassData)
	has, err := sess.Table(mm.TableName()).
		Select("draw_password").
		Where("id=?", this.MemberId).
		Get(data)
	return data, has, err
}

//根据id查询一条会员信息
func (*MemberSelfInfoBean) MemberOneMesInfo(id int64) (*schema.MemberMessage, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info := new(schema.MemberMessage)
	has, err := sess.Where("id=?", id).
		Where("delete_time=?", 0).
		Where("state=?", 1).Get(info)
	return info, has, err
}
