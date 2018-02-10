package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type AgencyCountBean struct{}

//查询股东
func (*AgencyCountBean) GetSearchFirstAgency(this *input.FirstAgency, listparam *global.ListParams, user *global.RedisStruct) ([]back.FirstAgencyBack, int64, error) {
	//todo index_id 问题  待修复
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	data := make([]back.FirstAgencyBack, 0)
	var agencyData []schema.Agency
	var AgencyCountData []schema.AgencyCount
	var count int64
	var err error
	var agencyIds []int64 //代理id
	AgencyCount := new(schema.AgencyCount)
	Agency := new(schema.Agency)

	//检索条件
	if this.AccountName == "" {
		this.Isvague = 2
	}
	if this.IsOnline != 0 {
		sess.Where("is_login=?", this.IsOnline)
	}
	if this.FormValue != "" {
		sess.Where("site_index_id=?", this.FormValue)
	}
	if this.Status != 0 {
		sess.Where("status=?", this.Status)
	}
	if this.FirstId != 0 {
		sess.Where("agency_id=?", this.FirstId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}

	//精确查询
	if this.Isvague == 0 {
		if this.AccountName != "" {
			sess.Where("account=?", this.AccountName)
		}
	} else {
		if this.AccountName != "" {
			sess.Where("account like ?", "%"+this.AccountName+"%")
		}
	}
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	listparam.Make(sess)
	err = sess.Table(Agency.TableName()).
		Where("site_id=?", user.SiteId).
		Where("delete_time=?", 0).
		Where("level=?", 2).
		OrderBy("create_time DESC").
		Find(&agencyData)

	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}

	for _, v := range agencyData {
		agencyIds = append(agencyIds, v.Id)
	}

	err = sess.Table(AgencyCount.TableName()).
		Where("site_id=?", user.SiteId).
		In("agency_id", agencyIds).
		Find(&AgencyCountData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	for _, v := range agencyData {
		for _, vv := range AgencyCountData {
			if v.Id == vv.AgencyId {
				list := new(back.FirstAgencyBack)
				list.AgencyId = v.Id
				list.Username = v.Username
				list.IsLogin = v.IsLogin
				list.Status = v.Status
				list.SiteIndexId = v.SiteIndexId
				list.SecondCount = vv.SecondCount
				list.ThirdCount = vv.ThirdCount
				list.MemberCount = vv.MemberCount
				list.Account = v.Account
				list.CreateTime = v.CreateTime
				data = append(data, *list)
				break
			}
		}
	}

	count, err = sess.Table(Agency.TableName()).
		Where("site_id=?", user.SiteId).
		Where("delete_time=?", 0).
		Where("level=?", 2).
		Where(conds).
		Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//查询总代
func (*AgencyCountBean) GetSearchSecondAgency(this *input.SecondAgency, listparam *global.ListParams) (
	[]back.SecondAgencyBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	second := new(schema.AgencyCount)
	var data []back.SecondAgencyBack
	var count int64
	if this.AccountName == "" {
		this.Isvague = 2
	}
	if this.FormValue != 0 {
		sess.Where("sales_agency_count.first_id=?", this.FormValue)
	}
	if this.Status != 0 {
		sess.Where("sales_agency.status=?", this.Status)
	}
	if this.SiteIndexId != "" {
		sess.Where("sales_agency_count.site_index_id=?", this.SiteIndexId)
	}
	if this.IsOnline != 0 {
		sess.Where("sales_agency.is_login=?", this.IsOnline)
	}
	if this.SiteId != "" {
		sess.Where("sales_agency_count.site_id=?", this.SiteId)
	}
	if this.FirstId != 0 {
		sess.Where("sales_agency_count.first_id=?", this.FirstId)
	}
	sess.Where("sales_agency.level=?", 3).Where("sales_agency.delete_time=?", 0)
	if this.Isvague == 0 {
		if this.AccountName != "" {
			sess.Where("sales_agency.account=?", this.AccountName)
		}
		var fag back.SecondAgencyBack
		has, err := sess.Table(second.TableName()).
			Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
			Get(&fag)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, 0, err
		}
		if !has {
			return data, 0, err
		}
		var ag schema.Agency
		_, err = sess.Table(agency.TableName()).
			Where("id=?", fag.FirstId).
			Get(&ag)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, 0, err
		}
		fag.FirstAccount = ag.Account
		data = append(data, fag)
		count = 1
	} else {
		//获得分页记录
		listparam.Make(sess)
		sess.Where("sales_agency_count.first_id!=?", 0).
			Where("sales_agency_count.second_id=?", 0)
		if this.AccountName != "" {
			sess.Where("sales_agency.account like ?", "%"+this.AccountName+"%")
		}
		conds := sess.Conds()
		var data1 []schema.Agency
		err := sess.Table(second.TableName()).
			Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
			Find(&data)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, 0, err
		}
		err = sess.Table(second.TableName()).
			Join("LEFT", agency.TableName(), "sales_agency_count.first_id=sales_agency.id").
			Where("sales_agency.id!=?", 0).Where("sales_agency.delete_time=?", 0).
			Find(&data1)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, 0, err
		}
		for k, v := range data {
			for _, vs := range data1 {
				if v.FirstId == vs.Id {
					data[k].FirstAccount = vs.Account
				}
			}
		}
		count, err = sess.Table(second.TableName()).
			Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
			Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, 0, err
		}
	}
	return data, count, nil
}

//查询代理
func (*AgencyCountBean) GetSearchThirdAgency(this *input.ThirdAgency, listparam *global.ListParams) (
	[]back.ThirdAgencyBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.ThirdAgencyBack
	var count int64
	var err error
	agency := new(schema.Agency)
	third := new(schema.AgencyCount)
	if this.AccountName == "" {
		this.Isvague = 2
	}
	if this.FormValue != 0 {
		sess.Where("sales_agency_count.second_id=?", this.FormValue)
	}
	if this.Status != 0 {
		sess.Where("sales_agency.status=?", this.Status)
	}
	if this.SiteIndexId != "" {
		sess.Where("sales_agency_count.site_index_id=?", this.SiteIndexId)
	}
	if this.FirstId != 0 {
		sess.Where("sales_agency_count.first_id=?", this.FirstId)
	}
	if this.SecondId != 0 {
		sess.Where("sales_agency_count.second_id=?", this.SecondId)
	}
	if this.ThirdId != 0 {
		sess.Where("sales_agency_count.agency_id=?", this.ThirdId)
	}
	if this.SiteId != "" {
		sess.Where("sales_agency_count.site_id=?", this.SiteId)
	}
	if this.IsOnline != 0 {
		sess.Where("sales_agency.is_login=?", this.IsOnline)
	}
	sess.Where("sales_agency.level=?", 4).
		Where("sales_agency.delete_time=?", 0)
	if this.Isvague == 0 {
		if this.AccountName != "" {
			sess.Where("sales_agency.account=?", this.AccountName)
		}
		var fag back.ThirdAgencyBack
		var has bool
		has, err := sess.Table(third).
			Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").Get(&fag)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		if !has {
			return nil, count, err
		}
		var th_arr []int64
		th_arr = append(th_arr, fag.FirstId)
		th_arr = append(th_arr, fag.SecondId)
		var ag []schema.Agency
		ses := global.GetXorm().NewSession()
		err = ses.Table(agency.TableName()).In("id", th_arr).Find(&ag)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		for _, v := range ag {
			if fag.FirstId == v.Id {
				fag.FirstAccount = v.Account
			}
			if fag.SecondId == v.Id {
				fag.SecondAccount = v.Account
			}
		}
		data = append(data, fag)
		count = 1
	} else {
		if this.AccountName != "" {
			sess.Where("sales_agency.account like ?", "%"+this.AccountName+"%")
		}
		listparam.Make(sess)
		conds := sess.Conds()
		var data1 []schema.Agency
		var data2 []schema.Agency
		err = sess.Table(third.TableName()).
			Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
			Find(&data)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		err = sess.Table(third.TableName()).
			Join("LEFT", agency.TableName(), "sales_agency_count.first_id=sales_agency.id").
			Where("sales_agency.id!=?", 0).Where("sales_agency.delete_time=?", 0).
			Find(&data1)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		err = sess.Table(third.TableName()).
			Join("LEFT", agency.TableName(), "sales_agency_count.second_id=sales_agency.id").
			Where("sales_agency.id!=?", 0).Where("sales_agency.delete_time=?", 0).
			Find(&data2)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		for k, v := range data {
			for _, vs := range data1 {
				if v.FirstId == vs.Id {
					data[k].FirstAccount = vs.Account
				}
			}
			for _, vd := range data2 {
				if v.SecondId == vd.Id {
					data[k].SecondAccount = vd.Account
				}
			}
		}
		count, err = sess.Table(third.TableName()).
			Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
			Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
	}
	return data, count, err
}

//体系查询
func (*AgencyCountBean) SearchSystem(this *input.Search) ([]back.SearchBack, int64, error) {
	agency := new(schema.Agency)
	search := new(schema.AgencyCount)
	member := new(schema.Member)
	var data []back.SearchBack
	var count int64
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Level == 4 {
		var se back.SearchBack
		if this.Account != "" {
			sess.Where("account=?", this.Account)
		}
		has, err := sess.Table(member.TableName()).
			Select("id,site_id,account,site_index_id,first_agency_id,second_agency_id,third_agency_id").
			Get(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, 0, err
		}
		if !has {
			return nil, 0, err
		}
		var age []int64
		age = append(age, member.FirstAgencyId)
		age = append(age, member.SecondAgencyId)
		age = append(age, member.ThirdAgencyId)
		age_e := RemoveDuplicatesAndEmpty(age)
		se.AgencyId = member.Id
		se.MemberAccount = member.Account
		se.Account = member.Account
		//代理
		var ss []schema.Agency
		ses := global.GetXorm().NewSession()
		err = ses.Table(agency.TableName()).In("id", age_e).Find(&ss)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, 0, err
		}
		for _, v := range ss {
			if member.FirstAgencyId == v.Id {
				se.FirstAccount = v.Account
			}
			if member.SecondAgencyId == v.Id {
				se.SecondAccount = v.Account
			}
			if member.ThirdAgencyId == v.Id {
				se.ThirdAccount = v.Account
			}
		}
		count = 1
		data = append(data, se)
	} else {
		if this.Account != "" {
			sess.Where("sales_agency.account=?", this.Account)
		}
		if this.SiteId != "" {
			sess.Where("sales_agency.site_id=?", this.SiteId)
		}
		var se back.SearchBack
		switch this.Level {
		case 1:
			sess.Where("sales_agency.level=?", 2)
			has, err := sess.Table(search.TableName()).
				Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
				Get(&se)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, count, err
			}
			if !has {
				return nil, 0, err
			}
			se.FirstAccount = se.Account
			data = append(data, se)
		case 2:
			if this.SiteId != "" {
				sess.Where("sales_agency.site_id=?", this.SiteId)
			}
			sess.Where("sales_agency.level=?", 3)
			has, err := sess.Table(search.TableName()).
				Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
				Get(&se)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, count, err
			}
			if !has {
				return nil, 0, err
			}
			se.SecondAccount = se.Account
			var ses schema.Agency
			has, err = sess.Table(agency.TableName()).
				Where("id=?", se.FirstId).
				Get(&ses)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, count, err
			}
			se.FirstAccount = ses.Account
			data = append(data, se)
		case 3:
			if this.SiteId != "" {
				sess.Where("sales_agency.site_id=?", this.SiteId)
			}
			sess.Where("sales_agency.level=?", 4)
			has, err := sess.Table(search.TableName()).
				Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
				Get(&se)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, count, err
			}
			if !has {
				return nil, 0, err
			}
			se.ThirdAccount = se.Account
			var age []int64
			age = append(age, se.FirstId)
			age = append(age, se.SecondId)
			//查股东名称
			var sh []schema.Agency
			ses := global.GetXorm().NewSession()
			err = ses.Table(agency.TableName()).
				In("id", age).
				Find(&sh)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, count, err
			}
			for _, v := range sh {
				if se.SecondId == v.Id {
					se.SecondAccount = v.Account
				}
				if se.FirstId == v.Id {
					se.FirstAccount = v.Account
				}
			}
			data = append(data, se)
		}
		count = 1
	}
	return data, count, nil
}

//体系查询模糊查询
func (*AgencyCountBean) SearchSystemBlur(this *input.Search, listparam *global.ListParams) ([]back.SearchBack, int64, error) {
	agency := new(schema.Agency)
	search := new(schema.AgencyCount)
	member := new(schema.Member)
	var data []back.SearchBack
	var count int64
	sess := global.GetXorm().NewSession().Table(search.TableName())
	defer sess.Close()
	if this.Level == 4 {
		var se back.SearchBack
		var me []schema.Member
		if this.Account != "" {
			sess.Where("account like ?", "%"+this.Account+"%")
		}
		listparam.Make(sess)
		conds := sess.Conds()
		err := sess.Table(member.TableName()).
			Select("id,site_id,account,site_index_id,first_agency_id,second_agency_id,third_agency_id").
			Find(&me)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		count, err = sess.Table(member.TableName()).Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		var firstids []int64
		for _, v := range me {
			firstids = append(firstids, v.FirstAgencyId)
			firstids = append(firstids, v.SecondAgencyId)
			firstids = append(firstids, v.ThirdAgencyId)
		}
		firstids_w := RemoveDuplicatesAndEmpty(firstids)
		//代理
		var ss []schema.Agency
		err = sess.Table(agency.TableName()).In("id", firstids_w).Find(&ss)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		for _, v := range me {
			for _, vs := range ss {
				if v.FirstAgencyId == vs.Id {
					se.FirstAccount = vs.Account
				}
				if v.SecondAgencyId == vs.Id {
					se.SecondAccount = vs.Account
				}
				if v.ThirdAgencyId == vs.Id {
					se.ThirdAccount = vs.Account
				}
			}
			se.AgencyId = v.Id
			se.SiteIndexId = v.SiteIndexId
			se.MemberAccount = v.Account
			se.Account = v.Account
			data = append(data, se)
		}
	}
	if this.Level == 1 {
		if this.Account != "" {
			sess.Where("sales_agency.account like ?", "%"+this.Account+"%").
				Where("sales_agency_count.first_id=?", 0)
		}
		listparam.Make(sess)
		conds := sess.Conds()
		err := sess.Table(search.TableName()).
			Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
			Find(&data)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		for k, v := range data {
			data[k].FirstAccount = v.Account
		}
		count, err = sess.Table(search.TableName()).
			Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").
			Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
	}
	if this.Level == 2 {
		if this.Account != "" {
			sess.Where("sales_agency.account like ?", "%"+this.Account+"%")
		}
		sess.Where("sales_agency_count.first_id!=?", 0).Where("sales_agency_count.second_id=?", 0)
		listparam.Make(sess)
		conds := sess.Conds()
		err := sess.Table(search.TableName()).Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").Find(&data)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		var firstids []int64
		for _, v := range data {
			firstids = append(firstids, v.FirstId)
		}
		//股东
		var sa []schema.Agency
		err = sess.Table(agency.TableName()).In("id", firstids).Find(&sa)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		for k, v := range data {
			for _, vs := range sa {
				if v.FirstId == vs.Id {
					data[k].FirstAccount = vs.Account
				}
			}
			data[k].SecondAccount = v.Account
		}
		count, err = sess.Table(search.TableName()).Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").Where(conds).Count()
	}
	if this.Level == 3 {
		if this.Account != "" {
			sess.Where("sales_agency.account like ?", "%"+this.Account+"%").Where("sales_agency_count.first_id!=?", 0).Where("sales_agency_count.second_id!=?", 0)
		}
		listparam.Make(sess)
		conds := sess.Conds()
		err := sess.Table(search.TableName()).Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").Find(&data)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		var firstids []int64
		for _, v := range data {
			firstids = append(firstids, v.FirstId)
			firstids = append(firstids, v.SecondId)
		}
		firstids_e := RemoveDuplicatesAndEmpty(firstids)
		//股东
		var sa []schema.Agency
		err = sess.Table(agency.TableName()).In("id", firstids_e).Find(&sa)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
		for k, v := range data {
			for _, vd := range sa {
				if v.FirstId == vd.Id {
					data[k].FirstAccount = vd.Account
				}
				if v.SecondId == vd.Id {
					data[k].SecondAccount = vd.Account
				}
			}
			data[k].ThirdAccount = v.Account
		}
		count, err = sess.Table(search.TableName()).Join("LEFT", "sales_agency", "sales_agency_count.agency_id = sales_agency.id").Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, count, err
		}
	}
	return data, count, nil
}

//新增股东
func (*AgencyCountBean) FirstAgencyAdd(this *input.FirstAgencyAdd) (int64, error) {
	acy := new(schema.Agency)
	agency := new(schema.Agency)
	siteCount := new(schema.SiteCount)
	sct := new(schema.SiteCount)
	agencyCount := new(schema.AgencyCount)
	//给agency表赋值
	agency.SiteId = this.SiteId
	agency.SiteIndexId = this.SiteIndexId
	agency.Account = this.Account
	agency.Username = this.Username
	agency.Password = this.Password
	agency.ParentId = this.ParentId
	agency.Status = this.Status
	agency.IsLogin = 2
	agency.RoleId = 2
	agency.Level = 2
	agency.IsSub = 2
	//给agency_count表赋值
	agencyCount.SiteId = this.SiteId
	agencyCount.SiteIndexId = this.SiteIndexId
	//给site_count表赋值
	siteCount.SiteId = this.SiteId
	siteCount.SiteIndexId = this.SiteIndexId
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	//查询站点下是否存在股东账号，不存在则新增的为默认股东
	has, err := sess.
		Table(acy.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Where("role_id = ?", 2).
		Get(acy)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err)
		sess.Rollback()
		return 0, err
	}
	if !has {
		agency.IsDefault = 1
	} else {
		agency.IsDefault = 2
	}
	count, err := sess.Table(agency.TableName()).InsertOne(agency)
	if err != nil || count <= 0 {
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
		}
		sess.Rollback()
		return 0, err
	}
	agencyCount.AgencyId = agency.Id
	count, err = sess.
		Table(agencyCount.TableName()).
		InsertOne(agencyCount)
	if err != nil || count <= 0 {
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
		}
		sess.Rollback()
		return 0, err
	}
	//获取site_count中的股东数量
	has, err = sess.
		Table(sct.TableName()).
		Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Select("first_agency_count").
		Get(sct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	siteCount.FirstAgencyCount = sct.FirstAgencyCount + 1
	count, err = sess.
		Table(siteCount.TableName()).
		Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Cols("first_agency_count").
		Update(siteCount)
	if err != nil || count != 1 {
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
		}
		sess.Rollback()
		return 0, err
	}
	sess.Commit()
	return 1, err
}

//修改股东
func (*AgencyCountBean) FirstAgencyEdit(this *input.FirstAgencyEdit) (int64, error) {
	agency := new(schema.Agency)
	//给账号表赋值
	agency.Username = this.Username
	agency.Password = this.Password
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if agency.Password != "" {
		sess.Cols("password")
	}
	if agency.OperatePassword != "" {
		sess.Cols("operate_password")
	}
	count, err := sess.Table(agency.TableName()).
		Where("id = ?", this.Id).
		Cols("username").
		Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//新增总代
func (*AgencyCountBean) SecondAgencyAdd(this *input.AgencyAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	acy := new(schema.Agency)
	agency := new(schema.Agency)
	agc := new(schema.Agency)
	sct := new(schema.SiteCount)
	siteCount := new(schema.SiteCount)
	a := new(schema.AgencyCount)
	agencyCount := new(schema.AgencyCount) //添加总代操作使用
	ac := new(schema.AgencyCount)          //修改股东操作使用
	//给agency表赋值
	agency.SiteId = this.SiteId
	agency.SiteIndexId = this.SiteIndexId
	agency.Account = this.Account
	agency.Username = this.Username
	agency.Password = this.Password
	agency.ParentId = this.ParentId
	agency.Status = this.Status
	agency.RoleId = 3
	agency.IsLogin = 2
	agency.Level = 3
	agency.IsSub = 2
	agency.IsDefault = 2
	//给agency_count表赋值
	agencyCount.SiteId = this.SiteId
	agencyCount.SiteIndexId = this.SiteIndexId
	agencyCount.FirstId = this.ParentId
	var count int64
	var err error
	//查询站点下的默认股东id
	_, err = sess.Table(agc.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Where("is_default = ?", 1).Where("role_id = ?", 2).
		Select("id").Get(agc)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//查询站点下默认股东是否存在总代账号，不存在则新增的为默认总代
	has, err := sess.Table(acy.TableName()).Where("site_id=?", this.SiteId).Where("site_index_id = ?", this.SiteIndexId).Where("parent_id = ?", agc.Id).Where("role_id = ?", 3).Get(acy)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if !has {
		agency.IsDefault = 1
	} else {
		agency.IsDefault = 2
	}
	sess.Begin()
	count, err = sess.Table(agency.TableName()).InsertOne(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	agencyCount.AgencyId = agency.Id
	count, err = sess.Table(agencyCount.TableName()).InsertOne(agencyCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	ac.AgencyId = this.ParentId
	//获取股东的总代人数
	_, err = sess.Table(a.TableName()).Where("agency_id = ?", ac.AgencyId).Select("second_count").Get(a)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	ac.SecondCount = a.SecondCount + 1
	count, err = sess.Table(ac.TableName()).Where("agency_id = ?", ac.AgencyId).Cols("second_count").Update(ac)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//获取site_count中的总代数量
	_, err = sess.Table(sct.TableName()).Where("site_id = ?", this.SiteId).Where("site_index_id = ?", this.SiteIndexId).Select("second_agency_count").Get(sct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	siteCount.SecondAgencyCount = sct.SecondAgencyCount + 1
	count, err = sess.Table(siteCount.TableName()).Where("site_id = ?", this.SiteId).Where("site_index_id = ?", this.SiteIndexId).Cols("second_agency_count").Update(siteCount)
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

//新增代理
func (*AgencyCountBean) ThirdAgencyAdd(this *input.AgencyAdd) (int64, error) {
	acy := new(schema.Agency)
	agency := new(schema.Agency)
	agc := new(schema.Agency)
	sct := new(schema.SiteCount)
	siteCount := new(schema.SiteCount)
	agencyCount := new(schema.AgencyCount) //添加代理使用
	ac := new(schema.AgencyCount)          //修改总代操作使用
	ae := new(schema.AgencyCount)          //获取first_id使用
	ag := new(schema.AgencyCount)          //总代获取使用
	age := new(schema.AgencyCount)         //股东获取使用
	a := new(schema.AgencyCount)           //修改股东操作使用
	//给agency表赋值
	agency.SiteId = this.SiteId
	agency.SiteIndexId = this.SiteIndexId
	agency.Account = this.Account
	agency.Username = this.Username
	agency.Password = this.Password
	agency.ParentId = this.ParentId
	agency.Status = this.Status
	agency.IsLogin = 2
	agency.RoleId = 4
	agency.Level = 4
	agency.IsSub = 2
	//给agency_count表赋值
	agencyCount.SiteId = this.SiteId
	agencyCount.SiteIndexId = this.SiteIndexId
	agencyCount.SecondId = this.ParentId
	sess := global.GetXorm().NewSession()
	var count int64
	var err error
	defer sess.Close()
	sess.Begin()
	//查询站点下的默认总代id
	_, err = sess.Table(agc.TableName()).Where("site_id=?", this.SiteId).Where("site_index_id = ?", this.SiteIndexId).Where("is_default = ?", 1).Where("role_id = ?", 3).Select("id").Get(agc)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//查询站点下默认总代是否存在代理账号，不存在则新增的为默认代理
	has, err := sess.Table(acy.TableName()).Where("site_id=?", this.SiteId).Where("site_index_id = ?", this.SiteIndexId).Where("parent_id = ?", agc.Id).Where("role_id = ?", 4).Get(acy)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
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
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//获取agencyCount.FirstId
	sess.Table(ae.TableName()).Where("agency_id = ?", this.ParentId).Select("first_id").Get(ae)
	agencyCount.AgencyId = agency.Id
	agencyCount.FirstId = ae.FirstId
	count, err = sess.Table(agencyCount.TableName()).InsertOne(agencyCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	ac.AgencyId = this.ParentId
	//获取总代的代理人数
	_, err = sess.Table(ag.TableName()).Where("agency_id = ?", ac.AgencyId).Select("third_count").Get(ag)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	ac.ThirdCount = ag.ThirdCount + 1
	//修改总代
	count, err = sess.Table(ac.TableName()).
		Where("agency_id = ?", ac.AgencyId).
		Cols("third_count").Update(ac)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	age.AgencyId = this.ParentId
	//获取股东id
	_, err = sess.Table(age.TableName()).Where("agency_id = ?", age.AgencyId).Select("first_id,third_count").Get(age)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	a.AgencyId = age.FirstId
	a.ThirdCount = ac.ThirdCount
	//修改股东
	count, err = sess.Table(a.TableName()).Where("agency_id = ?", a.AgencyId).Cols("third_count").Update(a)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//获取site_count中的总代数量
	_, err = sess.Table(sct.TableName()).Where("site_id = ?", this.SiteId).Where("site_index_id = ?", this.SiteIndexId).Select("third_agency_count").Get(sct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	siteCount.ThirdAgencyCount = sct.ThirdAgencyCount + 1
	count, err = sess.Table(siteCount.TableName()).Where("site_id = ?", this.SiteId).Where("site_index_id = ?", this.SiteIndexId).Cols("third_agency_count").Update(siteCount)
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

//数组去重
func RemoveDuplicatesAndEmpty(list []int64) (ret []int64) {
	var x []int64 = []int64{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	ret = x
	return
}
