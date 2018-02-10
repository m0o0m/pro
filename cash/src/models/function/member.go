package function

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golyu/sql-build"
	"global"
	"math"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
	"time"
)

type MemberBean struct{}

//会员列表
func (mb *MemberBean) List(l *input.MemberIndex, listParams *global.ListParams) ([]back.MemberIndex, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	data := make([]back.MemberIndex, 0)
	member := new(schema.Member)
	memberInfo := new(schema.MemberInfo)
	agency := new(schema.Agency)

	sess.Where(member.TableName()+".site_id=?", l.SiteId)
	sess.Where(member.TableName() + ".delete_time=0")
	sess.Where(agency.TableName() + ".delete_time=0")
	if l.SiteIndexId != "" {
		sess.Where(member.TableName()+".site_index_id=?", l.SiteIndexId)
	}
	if l.AgencyId != 0 {
		sess.Where(member.TableName()+".third_agency_id=?", l.AgencyId)
	}
	if l.SecondId != 0 {
		sess.Where(member.TableName()+".second_agency_id=?", l.SecondId)
	}
	if l.FirstId != 0 {
		sess.Where(member.TableName()+".first_agency_id = ?", l.FirstId)
	}
	if l.Status != 0 {
		sess.Where(member.TableName()+".status=?", l.Status)
	}
	if l.IsHide != 0 {
		sess.Where(member.TableName()+".id_hide=?", l.IsHide)
	}
	if l.Online != 0 {
		switch l.Online {
		case 1:
			sess.Where(member.TableName() + ".pc_status=1")
		case 2:
			sess.Where(member.TableName() + ".wap_status=1")
		case 3:
			sess.Where(member.TableName() + ".android_status=1")
		case 4:
			sess.Where(member.TableName() + ".ios_status=1")
		case 5:
			sess.Where(member.TableName() + ".pc_status=2")
		case 6:
			sess.Where(member.TableName() + ".wap_status=2")
		case 7:
			sess.Where(member.TableName() + ".android_status=2")
		case 8:
			sess.Where(member.TableName() + ".ios_status=2")
		case 9:
			sess.Where(member.TableName() + ".pc_status=1").And(member.TableName() + ".wap_status=1").And(
				member.TableName() + ".ios_status=1").And(member.TableName() + ".android_status=1")
		case 10:
			sess.Where(member.TableName() + ".pc_status=2").And(member.TableName() + ".wap_status=2").And(
				member.TableName() + ".ios_status=2").And(member.TableName() + ".android_status=2")
		}
	}
	if l.Source != 0 {
		sess.Where(member.TableName()+".register_client_type=?", l.Source)
	}
	loc, _ := time.LoadLocation("Local")
	if l.StartTime != "" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", l.StartTime, loc)
		if err == nil {
			sess.Where(member.TableName()+".create_time>=?", t.Unix())
		}
	}
	if l.EndTime != "" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", l.EndTime, loc)
		if err == nil {
			sess.Where(member.TableName()+".create_time<=?", t.Unix())
		}
	}
	col := ""
	switch l.Type {
	case 1:
		col = member.TableName() + ".account"
	case 2:
		col = member.TableName() + ".realname"
	case 3:
		col = member.TableName() + ".register_ip"
	case 4:
		col = member.TableName() + ".login_ip"
	case 5:
		col = memberInfo.TableName() + ".mobile"
	case 6:
		col = memberInfo.TableName() + ".card"
	case 7:
		col = memberInfo.TableName() + ".email"
	case 8:
		col = memberInfo.TableName() + ".qq"
	case 9:
		col = memberInfo.TableName() + ".wechat"
	}
	if col != "" && l.TypeValue != "" {
		if l.IsVague == 1 {
			sess.Where(col+" like ?", "%"+l.TypeValue+"%")
		} else {
			sess.Where(col+"=?", l.TypeValue)
		}
	}
	switch l.SortBy {
	case "1":
		listParams.OrderBy = member.TableName() + ".create_time"
	case "2":
		listParams.OrderBy = member.TableName() + ".account"
	case "3":
		listParams.OrderBy = member.TableName() + ".login_time"
	case "4":
		listParams.OrderBy = member.TableName() + ".balance"
	}
	switch l.Sort {
	case 1:
		listParams.Desc = true
	default:
		listParams.Desc = false
	}
	if l.PageSize == 50 || l.PageSize == 100 || l.PageSize == 200 {
		listParams.PageSize = l.PageSize
	}
	conds := sess.Conds()
	listParams.Make(sess)

	where1 := fmt.Sprintf("%s.member_id = %s.id", memberInfo.TableName(), member.TableName())
	where2 := fmt.Sprintf("%s.third_agency_id = %s.id", member.TableName(), agency.TableName())
	err := sess.Table(member.TableName()).Join("LEFT", memberInfo.TableName(), where1).
		Join("LEFT", agency.TableName(), where2).Find(&data)
	if err != nil {
		return data, 0, err
	}
	//获取视讯余额
	mvbs, err := mb.MemberVideoBalance(l.SiteId)
	if err != nil {
		return data, 0, err
	}

	var mvb back.MemberVideoBalance
	//平台id和名称（转换使用）
	type s struct {
		Id       int64
		Platform string
	}
	var q s
	var ss []s
	//把会员列表数据和会员视讯融合
	if len(data) > 0 && len(mvbs) > 0 {
		for i := range data {
			for j := range mvbs {
				if len(ss) > 0 {
					var f int64
					for _, s := range ss {
						if s.Id == mvbs[j].PlatformId {
							f += 1
						}
					}
					if f < 1 {
						q.Id = mvbs[j].PlatformId
						q.Platform = mvbs[j].Platform
						ss = append(ss, q)
					}
				} else {
					q.Id = mvbs[j].PlatformId
					q.Platform = mvbs[j].Platform
					ss = append(ss, q)
				}
				if data[i].Id == mvbs[j].MemberId {
					mvb.Platform = mvbs[j].Platform
					mvb.Balance = mvbs[j].Balance
					mvb.MemberId = mvbs[j].MemberId
					mvb.PlatformId = mvbs[j].PlatformId
					data[i].MemberVideoBalance = append(data[i].MemberVideoBalance, mvb)
				}
			}

		}
		//补充会员没有的视讯余额为0
		for u, k := range data {
			for m := range ss {
				var l int64
				for _, v := range k.MemberVideoBalance {
					if ss[m].Platform == v.Platform {
						l += 1
					}
				}
				if l < 1 {
					mvb.Platform = ss[m].Platform
					mvb.Balance = 0
					mvb.MemberId = k.Id
					mvb.PlatformId = ss[m].Id
					data[u].MemberVideoBalance = append(data[u].MemberVideoBalance, mvb)
				}
			}
		}
	}
	for _, b := range data {
		if len(b.MemberVideoBalance) > 0 {
			//排序
			for i := 0; i < len(b.MemberVideoBalance)-1; i++ {
				for j := i + 1; j < len(b.MemberVideoBalance); j++ {
					if b.MemberVideoBalance[i].PlatformId > b.MemberVideoBalance[j].PlatformId {
						b.MemberVideoBalance[i], b.MemberVideoBalance[j] = b.MemberVideoBalance[j], b.MemberVideoBalance[i]
					}
				}
			}
		}
	}
	//给视讯余额排序使用
	type memberBalance struct {
		MemberId int64
		Balance  float64
	}
	var mbb memberBalance
	var mbbs []memberBalance
	var dd back.MemberIndex
	var da []back.MemberIndex
	//给视讯余额排序
	if l.SortBy != "" && l.SortBy != "1" && l.SortBy != "2" && l.SortBy != "3" && l.SortBy != "4" {
		for _, w := range data {
			for _, q := range w.MemberVideoBalance {
				if q.Platform == l.SortBy {
					mbb.Balance = q.Balance
					mbb.MemberId = q.MemberId
					mbbs = append(mbbs, mbb)
				}
			}

		}
		//排序
		for i := 0; i < len(mbbs)-1; i++ {
			for j := i + 1; j < len(mbbs); j++ {
				if l.Sort == 1 {
					//倒序（从大到小）
					if mbbs[i].Balance < mbbs[j].Balance {
						mbbs[i], mbbs[j] = mbbs[j], mbbs[i]
					}
				} else {
					//正序（从小到大）
					if mbbs[i].Balance > mbbs[j].Balance {
						mbbs[i], mbbs[j] = mbbs[j], mbbs[i]
					}
				}
			}
		}
		for _, i := range mbbs {
			for _, t := range data {
				if i.MemberId == t.Id {
					dd = t
					da = append(da, dd)
				}
			}
		}
		data = da
	}
	count, err := sess.Table(member.TableName()).Where(conds).Join("LEFT", memberInfo.TableName(),
		where1).Join("LEFT", agency.TableName(), where2).Count()
	return data, count, err
}

//会员总数以及今日注册人数
func (*MemberBean) MemberNumberBySite(siteId, siteIndexId string, sTime, eTime int64) (*back.MemberNumberBySite, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mb := new(schema.Member)
	//会员总数
	sess.Where("site_id=?", siteId)
	if siteIndexId != "" {
		sess.Where("site_index_id=?", siteIndexId)
	}
	//今日注册会员
	sess.Where("create_time>=?", sTime)
	sess.Where("create_time<=?", eTime)
	count2, err := sess.Table(mb.TableName()).Count()
	data := new(back.MemberNumberBySite)
	data.RegNum = count2
	return data, err
}

//修改状态
func (*MemberBean) Status(l *input.MemberStatus) (ok bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", l.Id)
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	member := new(schema.Member)
	ok, err = sess.Cols("status").Get(member)
	if err != nil || !ok {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if member.Status == 1 {
		member.Status = 2
	} else {
		member.Status = 1
	}
	//禁用会员后,login_key和各个设备登陆状态应该也需要还原.
	if member.Status == 2 {
		sess.Table(member.TableName())
		sess.Cols("pc_login_key", "wap_login_key", "ios_login_key", "android_login_key")
		sess.Cols("pc_status", "wap_status", "ios_status", "android_status")
		sess.Cols("status")
		member.PcStatus = 2
		member.WapStatus = 2
		member.IosStatus = 2
		member.AndroidStatus = 2
	}
	sess.Where(conds)
	row, err := sess.Update(member)
	if err != nil || row != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return false, err
	}
	return true, nil
}

//获取会员资料
func (*MemberBean) Info(l *input.MemberStatus) (bool, error, *back.MemberInfo) {
	member := new(schema.Member)
	sess := global.GetXorm().NewSession().Table(member.TableName())
	defer sess.Close()
	sess.Where("site_id=?", l.SiteId)
	if l.SiteIndexId != "" {
		sess.Where("site_index_id=?", l.SiteIndexId)
	}
	sess.Where("delete_time=0")
	b := new(back.MemberInfo)
	ok, err := sess.ID(l.Id).Cols("id,account,is_edit_password,realname").Get(b)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err, b
	}
	return ok, err, b
}

//修改会员基本资料
func (*MemberBean) UpdateInfo(l *input.MemberBaseInfo) (bool, error) {
	member := new(schema.Member)
	sess := global.GetXorm().NewSession().Table(member.TableName())
	defer sess.Close()
	sess.Where("site_id=?", l.SiteId)
	if l.SiteIndexId != "" {
		sess.Where("site_index_id=?", l.SiteIndexId)
	}
	sess.Where("delete_time=0")
	ok, err := sess.ID(l.Id).Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err
	}
	member.IsEditPassword = int8(l.IsEditPassword)
	if l.NewPassword != "" {
		member.Password = l.NewPassword
	}
	member.Realname = l.Realname
	count, err := sess.ID(member.Id).
		Cols("password,is_edit_password,realname").
		Update(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err
	}
	if count == 0 {
		ok = false
	}
	return ok, err
}

//获取会员详细信息
func (*MemberBean) Detail(l *input.MemberStatus) (bool, error, *back.MemberDetail) {
	member := new(schema.Member)
	member_info := new(schema.MemberInfo)
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("site_id=?", l.SiteId)
	if l.SiteIndexId != "" {
		sess.Where("site_index_id=?", l.SiteIndexId)
	}
	sess.Where("delete_time=0")
	b := new(back.MemberDetail)
	where := fmt.Sprintf("%s.member_id = %s.id", member_info.TableName(), member.TableName())
	ok, err := sess.Table(member.TableName()).ID(l.Id).
		Join("LEFT", member_info.TableName(), where).
		Get(b)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err, b
	}
	//会员银行卡信息
	var data []back.MemberBank
	memberBank := new(schema.MemberBank)
	sess.Select("id,card,card_address,bank_id")
	err = sess.Table(memberBank.TableName()).Where("member_id=?", l.Id).Where("delete_time=0").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err, b
	}
	b.MemberBank = data
	return ok, err, b
}

//获取会员出款银行卡集合
func (*MemberBean) Bank(l *input.MemberStatus) ([]back.MemberBank, int64, error) {
	member := new(schema.MemberBank)
	sess := global.GetXorm().NewSession().Table(member.TableName())
	defer sess.Close()
	var data []back.MemberBank
	err := sess.Where("member_id=?", l.Id).Where("delete_time=0").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(member.TableName()).
		Where("member_id=?", l.Id).
		Where("delete_time=0").Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取会员出款银行卡详情
func (*MemberBean) BankInfo(l *input.MemberBankInfo) (*back.MemberBank, error) {
	member := new(schema.MemberBank)
	sess := global.GetXorm().NewSession().Table(member.TableName())
	defer sess.Close()
	data := new(back.MemberBank)
	_, err := sess.Where("card=?", l.Card).Where("delete_time=0").Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//修改会员详细资料
func (*MemberBean) UpdateDetail(l *input.MemberDetail, agency_id int64) (bool, error) {
	member := new(schema.Member)
	member_info := new(schema.MemberInfo)
	sess := global.GetXorm().NewSession().Table(member.TableName())
	defer sess.Close()
	sess.Where("site_id=?", l.SiteId)
	if l.SiteIndexId != "" {
		sess.Where("site_index_id=?", l.SiteIndexId)
	}
	sess.Where("delete_time=0")
	//获取会员是否存在、是否属于修改人的站点
	ok, err := sess.ID(l.Id).Get(member)
	if err != nil || !ok || member.SiteId != l.SiteId || (l.SiteIndexId != "" && l.SiteIndexId != member.SiteIndexId) {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err
	}

	//如果是当前登录是代理,那么只能操作当前代理下的会员
	if l.SiteIndexId != "" {
		if member.ThirdAgencyId != agency_id {
			return false, nil
		}
	}

	sess.Begin()
	//更新会员基本资料表
	if l.Realname != "" || l.DrawPassword != "" {
		if l.DrawPassword != "" {
			member.DrawPassword = l.DrawPassword
			sess.Cols("draw_password")
		}
		if l.Realname != "" {
			member.Realname = l.Realname
			sess.Cols("realname")
		}
		_, err := sess.ID(member.Id).Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return false, err
		}
	}

	//更新会员详细资料表
	member_info.Remark = l.Remark
	sd := "remark"
	if l.Birthday != "" {
		loc, _ := time.LoadLocation("Local")
		t, _ := time.ParseInLocation("2006-01-02", l.Birthday, loc)
		member_info.Birthday = t.Unix()
		sd = sd + ",birthday"
	}
	if l.Card != "" {
		member_info.Card = l.Card
		sd = sd + ",card"
	}
	if l.Mobile != "" {
		member_info.Mobile = l.Mobile
		sd = sd + ",mobile"
	}
	if l.Email != "" {
		member_info.Email = l.Email
		sd = sd + ",email"
	}
	if l.QQ != "" {
		member_info.Qq = l.QQ
		sd = sd + ",qq"
	}
	if l.Wechat != "" {
		member_info.Wechat = l.Wechat
		sd = sd + ",wechat"
	}
	_, err = sess.Table(member_info.TableName()).ID(member.Id).Cols(sd).Update(member_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return false, err
	}
	if len(l.Ids) > 0 {
		mb := new(schema.MemberBank)
		sql := "UPDATE " + mb.TableName() + " SET" + " bank_id=CASE id"
		ids := strings.Split(l.Ids, ",")
		bid := strings.Split(l.BankIds, ",")
		for k, v := range ids {
			sql = sql + " WHEN " + v + " THEN " + bid[k]
		}

		sql = sql + " END," + mb.TableName() + ".card_address = CASE id"

		for q, w := range ids {
			sql = sql + " WHEN " + w + " THEN " + "'" + l.CardAddress[q] + "'"
		}

		sql = sql + " END," + mb.TableName() + ".card = CASE id"

		for t, y := range ids {
			sql = sql + " WHEN " + y + " THEN " + "'" + l.BankAccount[t] + "'"
		}
		sql = sql + " END"
		sql = sql + " WHERE id IN (" + l.Ids + ")"
		_, err = sess.Query(sql)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return false, err
		}
	}
	sess.Commit()
	return true, err
}

//修改会员银行信息
func (*MemberBean) UpdateMemberBank(l *input.MemberBankEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.MemberBank)
	member.Id = l.Id
	member.BankId = l.BankId
	member.Card = l.Card
	member.CardName = l.CardName
	member.CardAddress = l.CardAddress
	count, err = sess.Table(member.TableName()).Where("id = ?", member.Id).Cols("bank_id,card,card_name,card_address").Update(member)
	return
}

//删除会员银行
func (*MemberBean) DelMemberBank(l *input.MemberStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.MemberBank)
	member.Id = int64(l.Id)
	member.DeleteTime = time.Now().Unix()
	count, err := sess.Table(member.TableName()).Where("id = ?", member.Id).
		Cols("delete_time").Update(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看卡号是否存在
func (*MemberBean) GetCard(card string, id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.MemberBank)
	ok, err := sess.Table(member.TableName()).
		Where("card = ?", card).
		Where("id != ?", id).Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err
	}
	return ok, err
}

//根据会员账号查看会员信息(资金管理)
func (*MemberBean) GetMemberInfo(this *input.MemberInfo) (back.MemberInfos, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memb := new(schema.Member)
	var member back.MemberInfos
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	has, err := sess.Table(memb.TableName()).
		Where("site_id = ?", this.SiteId).
		Where("account = ?",
			this.Account).Get(&member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return member, has, err
	}
	return member, has, err
}

//根据会员Id修改会员余额(1.入款2.出款)
func (*MemberBean) ChangeBalance(id int64, balance float64, types int) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//首先获取原先的余额
	info := new(schema.Member)
	flag, err := sess.Where("id=?", id).Get(info)
	if err != nil {
		return 0, err
	}
	if !flag {
		return 0, err
	}
	inf := new(schema.Member)
	if types == 1 {
		inf.Balance = info.Balance + balance //计算出新的总额度
	}
	if types == 2 {
		inf.Balance = info.Balance - balance
	}
	count, err = sess.Where("id=?", id).Cols("balance").Update(inf)
	return
}

//查询成功推广过下线的会员id和推广人数
func (m *MemberBean) GetSpreadNum(spreadInfo *input.SpreadInfo, listParam *global.ListParams) (ids []int64, spreadNum map[int64]int64, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//获得分页记录
	listParam.Make(sess)
	memberSchema := schema.Member{}
	type Result struct {
		SpreadId     int64 `xorm:"spread_id"`
		SpreadNumber int64 `xorm:"spread_number"`
	}
	var result []Result
	sess.Table(memberSchema.TableName()).
		Select("spread_id,count(*) as spread_number").
		Where("spread_id>?", 0).
		GroupBy("spread_id")
	if spreadInfo.SiteId != "" {
		sess.Where("site_id=?", spreadInfo.SiteId)
	}
	if spreadInfo.SiteIndexId != "" {
		sess.Where("site_index_id=?", spreadInfo.SiteIndexId)
	}
	if spreadInfo.Account != "" {
		sess.Where("account=?", spreadInfo.Account)
	}
	if spreadInfo.RegisterIp != "" {
		sess.Where("register_ip=?", spreadInfo.RegisterIp)
	}
	if spreadInfo.SpreadId != "" {
		sess.Where("spread_id=?", spreadInfo.SpreadId)
	}

	conds := sess.Conds()

	err = sess.Find(&result)
	if err != nil {
		return
	}
	spreadNum = make(map[int64]int64)
	for _, v := range result {
		ids = append(ids, v.SpreadId)
		spreadNum[v.SpreadId] = v.SpreadNumber
	}
	count, err = sess.Table(memberSchema.TableName()).Where(conds).Count()
	return
}

//查询成功推广过的会员信息
func (m *MemberBean) GetSpreadInfo(ids []int64) ([]back.SpreadInfo, error) {
	memberSchema := &schema.Member{}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var members []back.SpreadInfo
	err := sess.Table(memberSchema.TableName()).
		In("id", ids).
		Find(&members)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return members, err
	}
	return members, err
}

//根据会员id获取会员详情
func (*MemberBean) GetInfoById(id int64) (*schema.Member, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", id)
	sess.Where("delete_time=?", 0)
	info := new(schema.Member)
	flag, err := sess.Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, err
}

//根据会员账号获取会员id
func GetMemberIdByAccount(account string) (memberId int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	_, err = sess.Where("account = ?", account).Select("id").Get(member)
	memberId = member.Id
	return
}

//根据账号取出账号详细信息
func (*MemberBean) GetInfoBySite(site, siteIndex, account string) (schema.Member, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.Member
	sess.Where("status=?", 1)
	sess.Where("site_id=?", site)
	sess.Where("site_index_id=?", siteIndex)
	sess.Where("delete_time=?", 0)
	sess.Where("account=?", account)
	flag, err := sess.Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, err
}

//根据会员id刷新会员信息
func (*MemberBean) RefreshMember(info *schema.Member, result, system, loginIp string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	member.LoginCount = info.LoginCount + 1
	member.LastLoginIp = info.LoginIp
	member.LastLoginTime = info.LastLoginTime
	member.LoginTime = global.GetBeijingtime()
	member.LoginIp = loginIp
	if system == "pc" {
		member.PcLoginKey = result
		member.WapLoginKey = "0"
		member.IosLoginKey = "0"
		member.AndroidLoginKey = "0"
		member.PcStatus = 1
		count, err := sess.Where("id=?", info.Id).Cols("last_login_time,last_login_ip,login_time," +
			"login_ip,pc_login_key,wap_login_key,ios_login_key,android_login_key,login_count," +
			"pc_status").Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else if system == "wap" {
		member.PcLoginKey = "0"
		member.WapLoginKey = result
		member.IosLoginKey = "0"
		member.AndroidLoginKey = "0"
		member.WapStatus = 1
		count, err := sess.Where("id=?", info.Id).Cols("last_login_time,last_login_ip,login_time,login_ip,pc_login_key,wap_login_key,ios_login_key,android_login_key,login_count,wap_status").Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else if system == "ios" {
		member.PcLoginKey = "0"
		member.WapLoginKey = "0"
		member.IosLoginKey = result
		member.AndroidLoginKey = "0"
		member.IosStatus = 1
		count, err := sess.Where("id=?", info.Id).Cols("last_login_time,last_login_ip,login_time,login_ip,pc_login_key,wap_login_key,ios_login_key,android_login_key,login_count,ios_status").Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	} else {
		member.PcLoginKey = "0"
		member.WapLoginKey = "0"
		member.IosLoginKey = "0"
		member.AndroidLoginKey = result
		member.AndroidStatus = 1
		count, err := sess.Where("id=?", info.Id).Cols("last_login_time,last_login_ip,login_time,login_ip,pc_login_key,wap_login_key,ios_login_key,android_login_key,login_count,android_status").Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
	}
	return 1, nil
}

//查询某个登录key是否存在
func (*MemberBean) GetLoginKey(loginkey, platform string) (flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	if platform == "pc" {
		sess.Where("pc_login_key=?", loginkey)
		sess.Where("delete_time=?", 0)
		sess.Where("status=?", 1)
		flag, err = sess.Table(member.TableName()).Exist()
	} else if platform == "wap" {
		sess.Where("wap_login_key=?", loginkey)
		sess.Where("delete_time=?", 0)
		sess.Where("status=?", 1)
		flag, err = sess.Table(member.TableName()).Exist()
	} else if platform == "ios" {
		sess.Where("ios_login_key=?", loginkey)
		sess.Where("delete_time=?", 0)
		sess.Where("status=?", 1)
		flag, err = sess.Table(member.TableName()).Exist()
	} else {
		sess.Where("android_login_key=?", loginkey)
		sess.Where("delete_time=?", 0)
		sess.Where("status=?", 1)
		flag, err = sess.Table(member.TableName()).Exist()
	}
	return
}

//判断该账号是否存在
func (*MemberBean) CheckIsExist(account string, siteId string) (flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	flag, err = sess.Where("account=?", account).Where("site_id=?", siteId).Exist(member)
	return
}

//查询会员推广信息(简要)
func (m *MemberBean) GetMemberSpreadById(siteId, indexId string, id int64) (memberRebateInfo back.MemberRebateInfo, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	sess.Table(member.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", indexId).
		Where("delete_time = ?", 0).
		Where("status = ?", 1)
	conds := sess.Conds()
	b, err := sess.Where("id = ?", id).
		Select("account,spread_money").
		Get(&memberRebateInfo)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	if !b {
		err = errors.New("not found member by <" + fmt.Sprintf("%d", id) + ">")
	}
	var num int64
	//查询推广人数
	b, err = sess.Table(member.TableName()).
		Where(conds).
		Select("count(*)").
		Where("spread_id = ?", id).
		Get(&num)
	if err != nil {
		return
	}
	if !b {
		err = errors.New("not found member by <" + fmt.Sprintf("%d", id) + ">")
	}
	memberRebateInfo.SpreadNum = num
	return
}

//查询有效会员(有推广)的详细信息
func (*MemberBean) GetValidMemberList(siteId, siteIndexId string) (
	[]*back.ValidMember, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var members []*back.ValidMember
	memberSchema := new(schema.Member)
	agencySchema := new(schema.Agency)
	var selectValue = []string{
		"t1.id",
		"t2.account as agency",
		"t1.account",
	}

	sql := "SELECT " + strings.Join(selectValue, ",") + " FROM " + memberSchema.TableName() + " t1 LEFT JOIN " + agencySchema.TableName() + " t2 ON t1.third_agency_id = t2.id WHERE t1.id IN (SELECT spread_id FROM sales_member WHERE spread_id > 0 GROUP BY spread_id) and t1.status = 1 and t1.delete_time = 0"
	if siteId != "" {
		sql += " and t1.site_id = '" + siteId + "'"
	}
	if siteIndexId != "" {
		sql += " and t1.site_index_id = '" + siteIndexId + "'"
	}
	//fmt.Println("sql", sql)
	err := sess.SQL(sql).Find(&members)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return members, err
	}
	return members, err
}

//查询出被推广的会员的详细信息
func (*MemberBean) GetBePromotedMemberList(siteId, siteIndexId string) (members []*schema.Member, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberSchema := new(schema.Member)
	sess.Table(memberSchema.TableName())
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	if siteIndexId != "" {
		sess.Where("site_index_id = ?", siteIndexId)
	}
	err = sess.Where("status = ?", 1).
		Where("spread_id > ?", 0).
		Where("delete_time = ?", 0).
		Find(&members)
	return
}

//根据会员账号获取会员资料
func (*MemberBean) AccountInfo(Account, SiteId, SiteIndexId string) (bool, error, *back.MemberInfo) {
	member := new(schema.Member)
	sess := global.GetXorm().NewSession().Table(member.TableName())
	defer sess.Close()
	sess.Where("site_id=?", SiteId)
	if SiteIndexId != "" {
		sess.Where("site_index_id=?", SiteIndexId)
	}
	sess.Where("delete_time=0")
	b := new(back.MemberInfo)
	ok, err := sess.Where("account = ?", Account).
		Cols("id,account,is_edit_password,realname,draw_password").Get(b)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ok, err, b
	}
	return ok, err, b
}

//注册
func (*MemberBean) MemberRegister(newMember *schema.Member, register *input.MemberRegister, setting schema.SiteMemberRegisterSet, isDiscout int) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err = sess.Begin()
	//增加会员
	count, err = sess.Insert(newMember)
	if err != nil {
		sess.Rollback()
		return 0, err
	}
	//判断是否有优惠，有增加现金流水记录
	if isDiscout == 1 {
		memberCashRecord := new(schema.MemberCashRecord)
		memberCashRecord.SiteId = newMember.SiteId
		memberCashRecord.SiteIndexId = newMember.SiteIndexId
		memberCashRecord.MemberId = newMember.Id
		memberCashRecord.UserName = newMember.Account
		memberCashRecord.AgencyId = newMember.ThirdAgencyId
		memberCashRecord.SourceType = 6
		memberCashRecord.TradeNo = ""
		memberCashRecord.Balance = newMember.Balance
		memberCashRecord.Type = 1
		memberCashRecord.Remark = "注册优惠"
		memberCashRecord.AfterBalance = newMember.Balance
		memberCashRecord.ClientType = int64(newMember.RegisterClientType)
		memberCashRecord.CreateTime = global.GetCurrentTime()
		memberCashRecord.DisBalance = newMember.Balance
		agencyBean := new(AgencyBean)
		info, _, err := agencyBean.GetAgency(newMember.ThirdAgencyId)
		if err != nil {
			sess.Rollback()
			return 0, err
		}
		memberCashRecord.AgencyAccount = info.Account
		count, err = sess.Insert(memberCashRecord)
		if err != nil || count != 1 {
			sess.Rollback()
			return 0, err
		}
	}
	//增加稽核记录
	memberAudit := new(schema.MemberAudit)
	memberAudit.SiteId = newMember.SiteId
	memberAudit.SiteIndexId = newMember.SiteIndexId
	memberAudit.MemberId = newMember.Id
	memberAudit.Account = newMember.Account
	memberAudit.Status = 1
	memberAudit.BeginTime = global.GetCurrentTime()
	memberAudit.EndTime = 0
	memberAudit.NormalMoney = 0
	memberAudit.MultipleMoney = newMember.Balance * float64(setting.AddMosaic)
	memberAudit.AdminMoney = 0
	memberAudit.DepositMoney = newMember.Balance
	count, err = sess.Insert(memberAudit)
	if err != nil || count != 1 {
		sess.Rollback()
		return 0, err
	}
	//增加会员详细信息
	memberinfo := new(schema.MemberInfo)
	memberinfo.MemberId = newMember.Id
	memberinfo.Card = register.PassPort
	memberinfo.Email = register.Email
	memberinfo.Mobile = register.Phone
	memberinfo.Qq = register.Qq
	memberinfo.Wechat = register.Wechat
	memberinfo.Remark = ""
	if register.Birthday != "" {
		tineUnix, err := time.Parse("2006-01-02", register.Birthday)
		//time
		if err != nil {
			sess.Rollback()
			return 0, err
		}
		memberinfo.Birthday = tineUnix.Unix()
	}
	memberinfo.LocalCode = register.LocalCode
	count, err = sess.Insert(memberinfo)
	if err != nil || count != 1 {
		sess.Rollback()
		return 0, err
	}
	//sales_agency_count member_count +=1  where agency_id in(股东id，总代id，代理id)
	sql := "update `sales_agency_count` set member_count=member_count+?  where agency_id in(?,?,?)"
	_, err = sess.Exec(sql, 1, newMember.ThirdAgencyId, newMember.SecondAgencyId, newMember.FirstAgencyId)
	if err != nil {
		sess.Rollback()
		return 0, err
	}
	err = sess.Commit()
	return
}

//查看某个列的值是否重复
func (*MemberBean) CheckIsExistValue(types int, conditions string) (flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.MemberInfo)
	//1.银行卡号是否可以重复2.电话号码是否可以重复3.邮箱是否可以重复4.qq是否可以重复5.邮箱是否可以重复
	switch types {
	case 1:
		flag, err = sess.Where("card=?", conditions).Exist(member)
	case 2:
		flag, err = sess.Where("mobile=?", conditions).Exist(member)
	case 3:
		flag, err = sess.Where("email=?", conditions).Exist(member)
	case 4:
		flag, err = sess.Where("qq=?", conditions).Exist(member)
	case 5:
		flag, err = sess.Where("wechat=?", conditions).Exist(member)
	}

	return
}

//查看姓名是否重复
func (*MemberBean) CheckRealNameExist(conditions string) (flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	flag, err = sess.Where("realname=?", conditions).Exist(member)
	return
}

//根据站点和注册ip来查询记录
func (*MemberBean) GetRecordByIp(site, siteindex, ip string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	sess.Where("site_id=?", site)
	sess.Where("site_index_id=?", siteindex)
	sess.Where("register_ip=?", ip)
	flag, err := sess.Table(member.TableName()).Exist()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return flag, err
	}
	return flag, err
}

//根据id来修改key
func (*MemberBean) ChangeLoginkey(id int64, system string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	if system == "pc" {
		member.PcLoginKey = ""
		member.PcStatus = 2
		count, err := sess.Where("id=?", id).Cols("pc_login_key,pc_status").Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
		if count != 1 {
			return count, err
		}
	} else if system == "wap" {
		member.WapLoginKey = ""
		member.WapStatus = 2
		count, err := sess.Where("id=?", id).Cols("wap_login_key,wap_status").Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
		if count != 1 {
			return count, err
		}
	} else if system == "ios" {
		member.IosLoginKey = ""
		member.IosStatus = 2
		count, err := sess.Where("id=?", id).Cols("ios_login_key,ios_status").Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
		if count != 1 {
			return count, err
		}
	} else {
		member.AndroidLoginKey = ""
		member.AndroidStatus = 2
		count, err := sess.Where("id=?", id).Cols("android_login_key,android_status").Update(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return count, err
		}
		if count != 1 {
			return count, err
		}
	}
	return 1, nil
}

//修改密码
func (*MemberBean) ChangePassword(id int64, password string, types int) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)

	if types == 1 {
		member.Password = password
		sess.Where("id=?", id)
		sess.Cols("password")
		count, err = sess.Update(member)
	} else if types == 2 {
		//取款密码
		member.DrawPassword = password
		sess.Where("id=?", id)
		sess.Cols("draw_password")
		count, err = sess.Update(member)
	}
	return
}

//根据代理id获取一个会员
func (*MemberBean) GetOneMemberByThird(id int64) (info schema.Member, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("third_agency_id=?", id)
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	flag, err = sess.Limit(1, 0).Get(&info)
	return
}

//添加会员详细信息
func (*MemberBean) InsertMemberInfo(info *schema.MemberInfo) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err = sess.Insert(info)
	return
}

//根据代理id获取会员列表(账号名称)
func (*MemberBean) GetAllList(this *input.ThirdAgencyInfo) ([]back.MemberInfo, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.MemberInfo
	member := new(schema.Member)
	if len(this.SiteId) != 0 {
		sess.Where("site_id = ?", this.SiteId)
	}
	if len(this.SiteIndexId) != 0 {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	sess.Where("third_agency_id = ?", this.Id)
	sess.Where("status = ?", 1)
	err := sess.Table(member.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取指定id的会员的信息
func (m *MemberBean) GetMemberByIds(memberIds []int64) (members []schema.Member, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err = sess.Table(new(schema.Member).TableName()).
		In("id", memberIds).
		Where("delete_time = 0").
		Find(&members)
	return
}

//根据会员id,得到冲销待扣款的会员信息
func (m *MemberBean) GetAuditMemberByIds(memberIds []int64, sessArgs ...*xorm.Session) ([]back.RebateAuditMember, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	memberSchema := new(schema.Member)
	agencySchema := new(schema.Agency)
	var auditMembers []back.RebateAuditMember
	var err error
	err = sess.Table(memberSchema.TableName()).
		Alias("t1").
		Join("LEFT", agencySchema.TableName(), "t1.third_agency_id = "+agencySchema.TableName()+".id").
		Select("t1.id,t1.account,"+agencySchema.TableName()+".id as agency_id,"+agencySchema.TableName()+".account as agency_account,t1.balance").
		In("t1.id", memberIds).
		Where("t1.delete_time = ?", 0).
		Find(&auditMembers)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return auditMembers, err
	}
	return auditMembers, err
}

//更新会员金额
func (bean *MemberBean) UpdateMoney(member *back.RebateAuditMember, sessArgs ...*xorm.Session) error {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	num, err := sess.Table(new(schema.Member).TableName()).
		Where("id = ?", member.Id).
		Update(map[string]interface{}{"balance": member.Balance})
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("update 0 row")
	}
	return nil
}

//更新会员金额
func (bean *MemberBean) UpdateMoneyById(id int64, money float64, sessArgs ...*xorm.Session) error {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	num, err := sess.Table(new(schema.Member).TableName()).
		Where("id = ?", id).
		Update(map[string]interface{}{"balance": money})
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("update 0 row")
	}
	return nil
}

//会员加款
func (m *MemberBean) AddMoneyById(siteId string, id int64, money float64, sessArgs ...*xorm.Session) (float64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	sql, err := sqlBuild.Update(new(schema.Member).TableName()).
		Where(id, "id").
		Where(siteId, "site_id").
		Set_(money, "balance = balance+", sqlBuild.Rule{Float64Value: -math.MaxFloat64}).
		String()
	if err != nil {
		return 0, err
	}
	// TODO 更新余额
	err = m.updateMoney(sql, sess)
	if err != nil {
		return 0, err
	}
	// TODO 查询余额
	return m.getMoney(siteId, id, sess)
}

//会员扣款
func (m *MemberBean) DelMoneyById(siteId string, id int64, money float64, sessArgs ...*xorm.Session) (float64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	sql, err := sqlBuild.Update(new(schema.Member).TableName()).
		Where(id, "id").
		Where(siteId, "site_id").
		Set_(money, "balance = balance-", sqlBuild.Rule{Float64Value: -math.MaxFloat64}).
		String()
	if err != nil {
		return 0, err
	}
	// TODO 更新余额
	err = m.updateMoney(sql, sess)
	if err != nil {
		return 0, err
	}
	// TODO 查询余额
	return m.getMoney(siteId, id, sess)
}

//更新会员金额
func (m *MemberBean) updateMoney(sql string, sess *xorm.Session) error {
	if sess == nil {
		panic("<sessArgs> incorrect parameter passed")
	}
	result, err := sess.Exec(sql)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("update 0 row")
	}
	return nil
}

//查询会员余额
func (m *MemberBean) getMoney(siteId string, id int64, sess *xorm.Session) (money float64, err error) {
	if sess == nil {
		panic("<sessArgs> incorrect parameter passed")
	}
	b, err := sess.Table(new(schema.Member).TableName()).
		Where("site_id = ?", siteId).
		Where("id = ?", id).
		Select("balance").
		Get(&money)
	if err != nil {
		return
	}
	if !(b) {
		err = errors.New("not found member balance")
	}
	return
}

//根据账号获取会员
func (*MemberBean) GetMemberByAccount(account string) (member schema.Member, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Where("account=?", account).
		Where("delete_time=?", 0).
		Get(&member)
	return
}

//根据账号获取会员
func (*MemberBean) GetMemberBySiteAccount(siteId, account string, sessArgs ...*xorm.Session) (member schema.Member, err error) {
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
	b, err := sess.
		Where("site_id=?", siteId).
		Where("account=?", account).
		Where("delete_time=?", 0).
		Where("status = ?", 1).
		Get(&member)
	if !(b) {
		err = errors.New("not found  member")
		return
	}
	return
}

//wap会员注册
func (*MemberBean) WapMemberRegister(newMember *schema.Member, register *input.WapMemberRegister, setting schema.SiteMemberRegisterSet, isDiscout int) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err := sess.Begin()
	//增加会员
	count, err := sess.Insert(newMember)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//判断是否有优惠，有增加现金流水记录
	if isDiscout == 1 {
		memberCashRecord := new(schema.MemberCashRecord)
		memberCashRecord.SiteId = newMember.SiteId
		memberCashRecord.SiteIndexId = newMember.SiteIndexId
		memberCashRecord.MemberId = newMember.Id
		memberCashRecord.UserName = newMember.Account
		memberCashRecord.AgencyId = newMember.ThirdAgencyId
		memberCashRecord.SourceType = 6
		memberCashRecord.TradeNo = ""
		memberCashRecord.Balance = newMember.Balance
		memberCashRecord.Type = 1
		memberCashRecord.Remark = "注册优惠"
		memberCashRecord.AfterBalance = newMember.Balance
		memberCashRecord.ClientType = int64(newMember.RegisterClientType)
		memberCashRecord.CreateTime = time.Now().Unix()
		memberCashRecord.DisBalance = newMember.Balance
		agencyBean := new(AgencyBean)
		info, _, err := agencyBean.GetAgency(newMember.ThirdAgencyId)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		memberCashRecord.AgencyAccount = info.Account
		count, err = sess.Insert(memberCashRecord)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return 0, err
		}
		if count != 1 {
			return count, err
		}
	}
	//增加稽核记录
	memberAudit := new(schema.MemberAudit)
	memberAudit.SiteId = newMember.SiteId
	memberAudit.SiteIndexId = newMember.SiteIndexId
	memberAudit.MemberId = newMember.Id
	memberAudit.Account = newMember.Account
	memberAudit.Status = 1
	memberAudit.BeginTime = time.Now().Unix()
	memberAudit.EndTime = 0
	memberAudit.NormalMoney = 0
	memberAudit.MultipleMoney = newMember.Balance * float64(setting.AddMosaic)
	memberAudit.AdminMoney = 0
	memberAudit.DepositMoney = newMember.Balance
	count, err = sess.Insert(memberAudit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count != 1 {
		return count, err
	}
	//增加会员详细信息
	memberinfo := new(schema.MemberInfo)
	memberinfo.MemberId = newMember.Id
	memberinfo.Remark = ""
	//if register.Birthday != "" {
	//	tineUnix, err := time.Parse("2006-01-02", register.Birthday)
	//	if err != nil {
	//		sess.Rollback()
	//		return 0, err
	//	}
	//	memberinfo.Birthday = tineUnix.Unix()
	//}
	//memberinfo.LocalCode = register.LocalCode
	count, err = sess.Insert(memberinfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count != 1 {
		return count, err
	}
	err = sess.Commit()
	return count, err
}

//增加登录日志
func (*MemberBean) AddLoginLog(log *schema.LoginLog) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err := sess.Insert(log)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看站点下是否有该会员
func (*MemberBean) GetMemberBySiteId(member *input.MemberVideoBalance) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	m := new(schema.Member)
	sess.Where("id=?", member.Id)
	sess.Where("status=?", 1)
	sess.Where("delete_time=?", 0)
	sess.Where("site_id=?", member.SiteId)
	has, err := sess.Get(m)
	return has, err
}

//踢线会员
func (*MemberBean) OffLine(this *input.MemberStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	//踢出会员账号需要更新把用户当前登陆的所有设备的login_key给更新
	//也需要更新当前设备的状态
	sess.Table(member.TableName())
	sess.Cols("pc_login_key", "wap_login_key", "ios_login_key", "android_login_key")
	sess.Cols("pc_status", "wap_status", "ios_status", "android_status")
	sess.Where("id=?", this.Id)
	sess.Where("delete_time=?", 0)
	member.Id = this.Id
	member.PcStatus = 2
	member.WapStatus = 2
	member.IosStatus = 2
	member.AndroidStatus = 2
	count, err := sess.Update(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//批量启用禁用会员
func (*MemberBean) BatchStatus(this *input.BatchMember) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	sess.Table(member.TableName())
	sess.Where("delete_time=0")
	sess.In("id", this.Ids)
	//禁用会员后,login_key和各个设备登陆状态应该也需要还原.
	if this.Status == 2 {
		sess.Cols("pc_login_key", "wap_login_key", "ios_login_key", "android_login_key")
		sess.Cols("pc_status", "wap_status", "ios_status", "android_status")
		member.PcStatus = 2
		member.WapStatus = 2
		member.IosStatus = 2
		member.AndroidStatus = 2
	}
	member.Status = this.Status
	return sess.Cols("status").Update(member)

}

//批量踢线会员
func (*MemberBean) BatchOffline(this *input.OfflineMember) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	//踢出会员账号需要更新把用户当前登陆的所有设备的login_key给更新
	//也需要更新当前设备的状态
	sess.Table(member.TableName())
	sess.Cols("pc_login_key", "wap_login_key", "ios_login_key", "android_login_key")
	sess.Cols("pc_status", "wap_status", "ios_status", "android_status")
	sess.In("id", this.Ids)
	sess.Where("delete_time=?", 0)
	member.PcStatus = 2
	member.WapStatus = 2
	member.IosStatus = 2
	member.AndroidStatus = 2
	return sess.Update(member)
}

//会员排序下拉
func (*MemberBean) MemberSortDrop(siteId string) ([]back.MemberSortDrop, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	cp := new(schema.ComboProduct)
	platform := new(schema.Platform)
	msd := make([]back.MemberSortDrop, 0)
	sess.Where(site.TableName()+".id=?", siteId).Where(site.TableName() + ".is_default=1")
	sql1 := fmt.Sprintf("%s.combo_id=%s.combo_id", site.TableName(), cp.TableName())
	sql2 := fmt.Sprintf("%s.platform_id=%s.id", cp.TableName(), platform.TableName())
	err := sess.Table(platform.TableName()).Join("LEFT", cp.TableName(), sql2).
		Join("LEFT", site.TableName(), sql1).Find(&msd)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return msd, err
	}
	return msd, err
}

//会员视讯余额
func (*MemberBean) MemberVideoBalance(siteId string) ([]back.MemberVideoBalance, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	platform := new(schema.Platform)
	mpcb := new(schema.MemberProductClassifyBalance)
	site := new(schema.Site)
	cp := new(schema.ComboProduct)
	var mvbs []back.MemberVideoBalance
	sql1 := fmt.Sprintf("%s.combo_id = %s.combo_id AND %s.is_default = 1 AND %s.id = '%s'", site.TableName(),
		cp.TableName(), site.TableName(), site.TableName(), siteId)
	sql2 := fmt.Sprintf("%s.platform_id=%s.id AND %s.delete_time = 0", cp.TableName(), platform.TableName(),
		platform.TableName())
	sql3 := fmt.Sprintf("%s.platform_id=%s.id AND %s.site_id = %s.id", mpcb.TableName(), platform.TableName(),
		mpcb.TableName(), site.TableName())
	//联表查询会员视讯余额
	err := sess.Table(cp.TableName()).Join("INNER", site.TableName(), sql1).Join("INNER",
		platform.TableName(), sql2).Join("LEFT", mpcb.TableName(), sql3).Find(&mvbs)
	return mvbs, err
}

//更新会员真实姓名
func (*MemberBean) UpdateMemberReallname(siteId, siteIndexId, account, reallyName string) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	member.Realname = reallyName
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("account=?", account)
	sess.Cols("realname")
	count, err = sess.Update(member)
	return
}

//得到指定代理下的会员id
func (m *MemberBean) GetMemberIdsByAgencyId(siteId string, agencyId int64, account string) (ids []int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	sess.Table(member.TableName()).
		Select("id")

	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	if account != "" {
		sess.Where("account = ?", account)
	}
	if agencyId > 0 {
		sess.Where("third_agency_id = ?", agencyId)
	}
	err = sess.Find(&ids)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
	}
	return
}
