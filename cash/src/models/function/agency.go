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
	"strconv"
	"time"
)

type AgencyBean struct{}

//取site_index_id下的所有股东(名称)
func (*AgencyBean) GetAllFirstIdName(this *input.FirstIdNameBySite) ([]back.FirstIdNameBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.FirstIdNameBack
	agency := new(schema.Agency)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	sess.Where("status=1")
	sess.Where("level=2")
	sess.Where("role_id=2")
	err := sess.Table(agency.TableName()).
		Where("delete_time=?", 0).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//取site_index_id下的所有股东(账号和名称)
func (*AgencyBean) GetAllFirstIdNameAccount(this *input.FirstIdNameBySite) ([]back.SecondIdNameBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.SecondIdNameBack
	agency := new(schema.Agency)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	sess.Where("status=1")
	sess.Where("level=2")
	sess.Where("role_id=2")
	err := sess.Table(agency.TableName()).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//根据股东id取所有的总代id
func (*AgencyBean) GetAllSecondIdName(this *input.SecondIdNameBySite) ([]back.SecondIdNameBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	var data []back.SecondIdNameBack
	if this.FirstId != 0 {
		sess.Where("parent_id=?", this.FirstId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("site_id=?", this.SiteId)
	err := sess.Table(agency.TableName()).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//根据前台站点获取所有的代理id和名称
func (*AgencyBean) GetAllThirdIdName(this *input.SecondIdNameBySite) ([]back.SecondIdNameBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	var data []back.SecondIdNameBack
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	err := sess.Table(agency.TableName()).
		Where("delete_time=?", 0).
		Where("level=?", 4).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//根据agencyid更改在线状态
func ChangeLoginStatus(agencyId int64) (num int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.IsLogin = 2
	agency.LoginKey = ""
	num, err = sess.ID(agencyId).Cols("login_key", "is_login").Update(agency)
	return
}

//根据loginkey判断是否登录
func (*AgencyBean) GetInfoByLoginKey(token string) (info *schema.Agency, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info = new(schema.Agency)
	flag, err = sess.Table(info.TableName()).Where("is_login=?", 1).Where("login_key=?", token).Get(info)
	return
}

//开户人列表
func (*AgencyBean) GetAllAccountHolder(this *input.AccountHolderList, listparam *global.ListParams) ([]back.HoldersBacks, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	agency := new(schema.Agency)
	site := new(schema.Site)
	siteCount := new(schema.SiteCount)
	siteDomain := new(schema.SiteDomain)
	combo := new(schema.Combo)
	data := make([]back.HoldersBacks, 0)
	var (
		count int64
		err   error
	)

	// a == sales_agency, b == sales_site, c == sales_site_count, d == sales_site_domain, e == sales_combo
	if this.Account != "" {
		sess.Where("a.account=?", this.Account)
	}
	if this.IsLogin != 0 {
		sess.Where("a.is_login=?", this.IsLogin)
	}
	if this.Status != 0 {
		sess.Where("a.status=?", this.Status)
	}
	if this.SiteId != "" {
		sess.Where("a.site_id=?", this.SiteId)
	}
	if this.ShunXu == true {
		switch this.PaiXu {
		case "create_time":
			sess.Desc("create_time")
		case "first_agency_count":
			sess.Desc("A")
		case "second_agency_count":
			sess.Desc("B")
		case "third_agency_count":
			sess.Desc("C")
		case "member_count":
			sess.Desc("D")
		case "site_index_id":
			sess.Desc("E")
		}
	} else if this.ShunXu == false {
		switch this.PaiXu {
		case "create_time":
			sess.Asc(agency.TableName() + "create_time")
		case "first_agency_count":
			sess.Asc("A")
		case "second_agency_count":
			sess.Asc("B")
		case "third_agency_count":
			sess.Asc("C")
		case "member_count":
			sess.Asc("D")
		case "site_index_id":
			sess.Asc("E")
		}
	}

	s := "a.id, a.is_login, a.username, a.account, a.create_time, a.`status`, b.combo_id, e.combo_name, d.domain, " +
		"SUM(c.first_agency_count) AS A, SUM(c.second_agency_count) AS B, " +
		"SUM(c.third_agency_count) AS C, SUM(c.member_count) AS D, COUNT(c.site_index_id) AS E"
	sess.Select(s).Join("LEFT", "`"+site.TableName()+"` AS b", "b.agency_id = a.id").
		Join("LEFT", "`"+siteCount.TableName()+"` AS c", "c.site_id = a.site_id").
		Join("LEFT", "`"+siteDomain.TableName()+"` AS d",
			"d.site_id = a.site_id AND d.site_index_id = a.site_index_id AND d.type = "+strconv.Itoa(global.DOMAIN_MANAGE)).
		Join("LEFT", "`"+combo.TableName()+"` AS e", "e.id = b.combo_id AND e.delete_time=0")
	sess.Table(agency.TableName()).Alias("a").
		Where("a.delete_time=?", 0).
		Where("a.is_sub=?", 2).
		Where("a.parent_id=?", 0)
	listparam.Make(sess)
	err = sess.Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(agency.TableName()).Alias("a").
		Join("LEFT", "`"+site.TableName()+"` AS b", "b.agency_id = a.id").
		Join("LEFT", "`"+siteCount.TableName()+"` AS c", "c.site_id = a.site_id").
		Join("LEFT", "`"+siteDomain.TableName()+"` AS d",
			"d.site_id = a.site_id AND d.site_index_id = a.site_index_id AND d.type = "+strconv.Itoa(global.DOMAIN_MANAGE)).
		Join("LEFT", "`"+combo.TableName()+"` AS e", "e.id = b.combo_id AND e.delete_time=0").
		Where("a.delete_time=?", 0).
		Where("a.is_sub=?", 2).
		Where("a.parent_id=?", 0).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return data, count, err
}

//添加开户人
func (*AgencyBean) AddAccountHolder(this *input.AddAccountHolder) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	sess.Begin()
	agency.Account = this.Account
	agency.Password = this.Password
	agency.OperatePassword = this.OperatePassword
	agency.Username = this.Username
	agency.Remark = this.Remark
	agency.Level = 1
	agency.Status = this.Status
	agency.IsSub = 2
	agency.SiteId = this.Site
	agency.RoleId = 1
	agency.IsDefault = 2
	agency.IsLogin = 2
	//站点
	site := new(schema.Site)
	site.Id = this.Site
	site.SiteName = this.SiteName
	site.IndexId = this.SiteIndex
	site.ComboId = this.ComboId
	site.DomainUp = this.DomainUp
	site.UpCose = this.UpCose
	site.Status = 1
	site.IsDefault = 1
	//站点统计表
	siteCount := new(schema.SiteCount)
	siteCount.SiteId = this.Site
	siteCount.SiteIndexId = this.SiteIndex
	//站点域名
	siteDomains := make([]schema.SiteDomain, 2)
	//开户人后台域名
	siteDomains[0].SiteId = this.Site
	siteDomains[0].SiteIndexId = ""
	siteDomains[0].Domain = this.ManageDomain
	siteDomains[0].IsDefault = 1
	siteDomains[0].CreateTime = time.Now().Unix()
	siteDomains[0].DeleteTime = 0
	siteDomains[0].Type = global.DOMAIN_MANAGE
	//子站点代理后台域名
	siteDomains[1].SiteId = this.Site
	siteDomains[1].SiteIndexId = this.SiteIndex
	siteDomains[1].Domain = this.AgencyDomain
	siteDomains[1].IsDefault = 1
	siteDomains[1].CreateTime = time.Now().Unix()
	siteDomains[1].DeleteTime = 0
	siteDomains[1].Type = global.DOMAIN_AGENCY
	//添加开户人
	count, err := sess.Table(agency.TableName()).InsertOne(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//站点
	site.AgencyId = agency.Id
	count, err = sess.Table(site.TableName()).InsertOne(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//站点客服信息
	siteInfo := new(schema.SiteInfo)
	siteInfo.SiteId = site.Id
	siteInfo.SiteIndexId = site.IndexId
	count, err = sess.InsertOne(siteInfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//site_count添加数据
	count, err = sess.Table(siteCount.TableName()).InsertOne(siteCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//site_domain添加数据
	sd := new(schema.SiteDomain)
	count, err = sess.Table(sd.TableName()).Insert(&siteDomains)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//查询开户人详情
func (*AgencyBean) GetAccountHelderInfo(this *input.AccountNameId) (data *back.AccountHolderInfo, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if this.Id != 0 {
		sess.Where("id = ?", this.Id)
	}
	data = new(back.AccountHolderInfo)
	has, err = sess.Table(agency.TableName()).Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//查询一条开户人信息
func (*AgencyBean) GetOneAccountHelder(this *input.AccountNameId) (data *back.AccountHolderInfo, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if this.Id != 0 {
		sess.Where("id != ?", this.Id)
	}
	if this.Account != "" {
		sess.Where("account=?", this.Account)
	}
	data = new(back.AccountHolderInfo)
	has, err = sess.Table(agency.TableName()).Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//查询一条开户人信息
func (*AgencyBean) AccountHelderGet(this *input.UpdataAccountHolder) (data *schema.Agency, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	data = new(schema.Agency)
	has, err = sess.Table(agency.TableName()).Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//修改开户人信息
func (*AgencyBean) UpdataHolder(this *input.UpdataAccountHolder) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	var agency_info schema.Agency
	agency_info.Username = this.Username
	agency_info.Password = this.Password
	agency_info.OperatePassword = this.OperatePassword
	agency_info.Status = this.Status
	count, err = sess.Table(agency.TableName()).
		Where("id=?", this.Id).
		Cols("password,username,status,operate_password").
		Update(agency_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//开户人禁用
func (*AgencyBean) AccountHolderDisable(this *input.HolderNameId) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if this.Status == 1 {
		agency.Status = 2
	}
	if this.Status == 2 {
		agency.Status = 1
	}
	count, err = sess.Table(agency.TableName()).
		Where("id=?", this.Id).
		Where("delete_time=?", 0).
		Cols("status").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//开户人删除
func (*AgencyBean) AccountHolderDelete(this *input.DelAccountHolder) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.DeleteTime = time.Now().Unix()
	count, err = sess.Table(agency.TableName()).
		Where("id=?", this.Id).
		Cols("delete_time").
		Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询一条
func (*AgencyBean) GetOneAgencyByid(this *schema.Agency) (*schema.Agency, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	data := new(schema.Agency)
	has, err := sess.Table(data.TableName()).Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//根据id获取详情
func (*AgencyBean) GetAgency(id int64) (schema.Agency, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data schema.Agency
	has, err := sess.Where("id=?", id).
		Where("status=?", 1).
		Where("delete_time=?", 0).
		Get(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//根据站点取得默认的三级代理详情
func (*AgencyBean) GetDefault(site, siteIndex string) (info schema.Agency, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Table(info.TableName()).
		Where("is_default=?", 1).
		Where("level=?", 4).
		Where("delete_time=?", 0).
		Where("status=?", 1).
		Get(&info)
	return
}

//查询代理列表
func (*AgencyBean) AgencyManageList(this *input.AgencyManageList, listparams *global.ListParams) (data []back.AgencyManageListBack, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ag := new(schema.Agency)
	ac := new(schema.AgencyCount)
	at := new(schema.AgencyThirdInfo)
	if this.Name != "" {
		if this.SpreadId == 1 { //账号
			sess.Where("", this.Name)
		} else if this.SpreadId == 2 { //代理账号
			sess.Where(ag.TableName()+".account = ?", this.Name)
		} else { //推广id
			sess.Where(at.TableName()+".spread_id=?", this.Name)
		}
	}
	if this.SiteId != "" {
		sess.Where(ag.TableName()+".site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where(ag.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	if this.Type != 0 {
		sess.Where(ag.TableName()+".level=?", this.Type)
	}
	if this.Status != 0 {
		sess.Where(ag.TableName()+".status=?", this.Status)
	}
	sess.Where(ag.TableName()+".delete_time=?", 0)
	sess.Where(ag.TableName()+".is_sub=?", 2)
	conds := sess.Conds()
	listparams.Make(sess)
	err = sess.Table(ag.TableName()).
		Join("LEFT", at.TableName(), ag.TableName()+".id="+at.TableName()+".agency_id").
		Join("LEFT", ac.TableName(), ag.TableName()+".id="+ac.TableName()+".agency_id").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(ag.TableName()).
		Join("LEFT", at.TableName(), ag.TableName()+".id="+at.TableName()+".agency_id").
		Join("LEFT", ac.TableName(), ag.TableName()+".id="+ac.TableName()+".agency_id").
		Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//代理占成比
func (*AgencyBean) OccupationRatio(this *input.AgencyOccupationRatio) (data []back.AgencyOccupationRatioBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	apc := new(schema.AgencyProductCommission)
	pt := new(schema.ProductType)
	err = sess.Table(apc.TableName()).
		Join("LEFT", pt.TableName(), apc.TableName()+".product_id="+pt.TableName()+".id").
		Where(apc.TableName()+".agency_id=?", this.Id).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//站点管理-代理列表查询
func (*AgencyBean) AgencyList(this *input.SiteAgencyList) (agencyList []back.SiteAgencyList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	role := new(schema.Role)
	sql := fmt.Sprintf("%s.role_id=%s.id", agency.TableName(), role.TableName())
	sess.Select(agency.TableName() + ".id," + agency.TableName() + ".role_id," + agency.TableName() +
		".site_id," + agency.TableName() + ".site_index_id," + agency.TableName() + ".username," +
		agency.TableName() + ".create_time," + agency.TableName() + ".account," + role.TableName() + ".role_name," + agency.TableName() + ".remark")
	err = sess.Table(agency.TableName()).Join("LEFT", role.TableName(), sql).Where("level!=1").
		Where("is_sub=2").Where("site_id=?", this.SiteId).Where(agency.TableName() + ".delete_time=0").
		Find(&agencyList)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return agencyList, err
	}
	return agencyList, err
}

//站点管理-代理查询
func (*AgencyBean) AgencyInfo(this *input.SiteAgencyInfo) (agencyInfoBack back.SiteAgencyInfoBack, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	role := new(schema.Role)
	sql1 := fmt.Sprintf("%s.role_id=%s.id", agency.TableName(), role.TableName())
	sess.Select(agency.TableName() + ".id," + agency.TableName() + ".parent_id," + agency.TableName() + ".site_id," + agency.TableName() +
		".site_index_id," + agency.TableName() + ".username," + agency.TableName() + ".account," + role.TableName() +
		".role_name," + agency.TableName() + ".remark")
	//查看代理
	agencyInfo := new(back.SiteAgencyInfo)
	has, err = sess.Table(agency.TableName()).Join("LEFT", role.TableName(), sql1).
		Where("level!=1").Where("is_sub=2").Where(agency.TableName()+".delete_time=0").
		Where(agency.TableName()+".id=?", this.Id).Get(agencyInfo)
	if err != nil || !has {
		global.GlobalLogger.Error("error:%s", err.Error())
		return agencyInfoBack, has, err
	}
	//查看代理父级账号
	a := new(schema.Agency)
	has, err = sess.Where("id=?", agencyInfo.ParentId).Select("account").Get(a)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return agencyInfoBack, has, err
	}
	agencyInfoBack.Id = agencyInfo.Id
	agencyInfoBack.Account = agencyInfo.Account
	agencyInfoBack.SiteId = agencyInfo.SiteId
	agencyInfoBack.SiteIndexId = agencyInfo.SiteIndexId
	agencyInfoBack.Username = agencyInfo.Username
	agencyInfoBack.Remark = agencyInfo.Remark
	agencyInfoBack.RoleName = agencyInfo.RoleName
	agencyInfoBack.ParentAccount = a.Account
	return agencyInfoBack, has, err
}

//站点管理-代理修改
func (*AgencyBean) AgencyEdit(this *input.SiteAgencyEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.Remark = this.Remark
	agency.Username = this.Username
	count, err = sess.Where("id=?", this.Id).Cols("username,remark").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点管理-代理添加
func (a *AgencyBean) AgencyAdd(this *input.SiteAgencyAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	//根据父级账号查询父级id,父级角色
	ab, err := a.GetParentId(this.ParentAccount)
	if err != nil {
		return
	}
	//获取站点下是否有代理或股东
	has, err := a.IsAgency(this.SiteId, this.SiteIndexId, ab.RoleId)
	if err != nil {
		return
	}
	if has {
		agency.IsDefault = 2
	} else {
		agency.IsDefault = 1
	}
	agency.ParentId = ab.Id
	agency.RoleId = ab.RoleId + 1
	agency.Remark = this.Remark
	agency.Username = this.Username
	agency.SiteId = this.SiteId
	agency.SiteIndexId = this.SiteIndexId
	agency.Account = this.Account
	agency.CreateTime = time.Now().Unix()
	agency.IsSub = 2
	agency.Level = int8(agency.RoleId)
	agency.Status = 1
	count, err = sess.InsertOne(agency)
	return
}

//站点管理-代理删除
func (*AgencyBean) AgencyDel(this *input.SiteAgencyInfo) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.DeleteTime = time.Now().Unix()
	count, err = sess.Where("id=?", this.Id).Cols("delete_time").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//代理账号是否存在(修改)
func (*AgencyBean) IsExistAccountEdit(id int64, siteId, siteIndexId, account string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	if id != 0 {
		sess.Where("id=?", id)
	}
	if siteId != "" {
		sess.Where("site_id=?", siteId)
	}
	if siteIndexId != "" {
		sess.Where("site_index_id=?", siteIndexId)
	}
	if account != "" {
		sess.Where("account=?", account)
	}
	has, err = sess.Where("delete_time=?", 0).Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//代理账号是否存在（添加）
func (*AgencyBean) IsExistAccountAdd(siteId, siteIndexId, account string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err = sess.Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Where("account=?", account).
		Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//根据父级账号查询父级id,父级角色
func (*AgencyBean) GetParentId(account string) (agencyBack *back.SiteAgency, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agencyBack = new(back.SiteAgency)
	_, err = sess.Table(agency.TableName()).Where("account=?", account).Get(agencyBack)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return agencyBack, err
	}
	return agencyBack, err
}

//更新代理保证金
func (*AgencyBean) UpdateVideoMoneyById(id int64, videoMoney float64, sessArgs ...*xorm.Session) error {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	num, err := sess.Table(new(schema.Agency).TableName()).
		Where("id = ?", id).
		Update(map[string]interface{}{"video_balance": videoMoney})
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("update 0 row")
	}
	return nil
}

//代理视讯余额加款
func (m *AgencyBean) AddVideoMoneyById(siteId string, id int64, money float64, sessArgs ...*xorm.Session) (float64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	sql, err := sqlBuild.Update(new(schema.Agency).TableName()).
		Where(id, "id").
		Where(siteId, "site_id").
		Set_(money, "video_balance = video_balance+", sqlBuild.Rule{Float64Value: -math.MaxFloat64}).
		String()
	if err != nil {
		return 0, err
	}
	// TODO 更新余额
	err = m.updateVideoMoney(sql, sess)
	if err != nil {
		return 0, err
	}
	// TODO 查询余额
	return m.getVideoMoney(siteId, id, sess)
}

//代理视讯余额扣款
func (m *AgencyBean) DelVideoMoneyById(siteId string, id int64, money float64, sessArgs ...*xorm.Session) (float64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	sql, err := sqlBuild.Update(new(schema.Agency).TableName()).
		Where(id, "id").
		Where(siteId, "site_id").
		Set_(money, "video_balance = video_balance-", sqlBuild.Rule{Float64Value: -math.MaxFloat64}).
		String()
	if err != nil {
		return 0, err
	}
	// TODO 更新余额
	err = m.updateVideoMoney(sql, sess)
	if err != nil {
		return 0, err
	}
	// TODO 查询余额
	return m.getVideoMoney(siteId, id, sess)
}

//更新代理
func (m *AgencyBean) updateVideoMoney(sql string, sess *xorm.Session) error {
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

//查询代理视讯余额
func (m *AgencyBean) getVideoMoney(siteId string, id int64, sess *xorm.Session) (money float64, err error) {
	if sess == nil {
		panic("<sessArgs> incorrect parameter passed")
	}
	b, err := sess.Table(new(schema.Agency).TableName()).
		Where("site_id = ?", siteId).
		Where("id = ?", id).
		Select("video_balance").
		Get(&money)
	if err != nil {
		return
	}
	if !(b) {
		err = errors.New("not found member balance")
	}
	return
}

//根据站点id和前台id修改视讯余额
func (*AgencyBean) UpdateSiteVideoBalance(upData *input.SiteVideoBalance, userInfo global.AdminRedisStruct) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err := sess.Begin()
	agency := new(schema.Agency)
	if upData.Operate == 1 {
		agency.VideoBalance = upData.NowBalance + upData.OperateBalance
	} else if upData.Operate == 2 {
		agency.VideoBalance = upData.NowBalance - upData.OperateBalance
	} else if upData.Operate == 3 {
		agency.VideoBalance = upData.NowBalance + upData.OperateBalance
	}
	sess.Where("site_id=?", upData.SiteId)
	sess.Where("role_id=?", 1)
	count, err := sess.Cols("video_balance").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//视讯额度修改记录表
	scr := new(schema.SiteCashRecord)
	scr.SiteId = upData.SiteId
	scr.SiteIndexId = upData.SiteIndexId
	scr.Money = upData.OperateBalance
	scr.State = 2
	scr.AdminName = userInfo.Account
	scr.Remark = upData.Remark
	scr.Balance = agency.VideoBalance
	scr.CreateTime = global.GetCurrentTime()
	if upData.Operate == 1 {
		scr.DoType = 1
		scr.CashType = 2
	} else if upData.Operate == 2 {
		scr.DoType = 2
		scr.CashType = 3
	} else if upData.Operate == 3 {
		scr.DoType = 1
		scr.CashType = 4
	}
	count, err = sess.Insert(scr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	err = sess.Commit()
	return count, err
}

//查询账号是否存在
func (*AgencyBean) IsExistAccount(account string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.Account = account
	flag, err := sess.Exist(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return flag, err
	}
	return flag, err
}

//获取某个站点下面是否存在默认代理线
func (*AgencyBean) IsExistDefaultAgency(siteId, siteIndexId string) ([]schema.Agency, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []schema.Agency
	agency := new(schema.Agency)
	sess.Where("is_default=?", 1)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("is_sub=?", 2)
	sess.Where("level!=?", 1)
	err := sess.Table(agency.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//根据站点id获取开户人信息
func (*AgencyBean) GetOpenInfo(site string) (schema.Agency, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var backData schema.Agency
	sess.Where("site_id=?", site)
	sess.Where("level=?", 1)
	sess.Where("is_sub=?", 2)
	sess.Where("role_id=?", 1)
	flag, err := sess.Get(&backData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, flag, err
	}
	return backData, flag, err
}

//一键增加三级代理
func (*AgencyBean) AddThirdAgency(inputData *input.GenerationAgency, openId int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	//首先增加股东
	shaleHolder := new(schema.Agency)
	shaleHolder.SiteId = inputData.SiteId
	shaleHolder.SiteIndexId = inputData.SiteIndexId
	//根据站点取得开户人信息
	shaleHolder.ParentId = openId
	shaleHolder.RoleId = 2
	shaleHolder.Account = inputData.DefaultShareholdersAccount
	shaleHolder.Remark = inputData.DefaultShareholdersRemark
	shaleHolder.Level = 2
	shaleHolder.IsDefault = 1
	shaleHolder.Status = 1
	shaleHolder.IsSub = 2
	shaleHolder.CreateTime = time.Now().Unix()
	shaleHolder.DeleteTime = 0
	shaleHolder.VideoBalance = 0
	shaleHolder.Username = inputData.DefaultShareholdersName
	password, err := global.MD5ByStr("123456", global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	shaleHolder.Password = password
	shaleHolder.OperatePassword = password
	count, err := sess.Insert(shaleHolder)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//增加总代理
	totalAgency := new(schema.Agency)
	totalAgency.ParentId = shaleHolder.Id
	totalAgency.SiteId = inputData.SiteId
	totalAgency.SiteIndexId = inputData.SiteIndexId
	totalAgency.RoleId = 3
	totalAgency.Account = inputData.DefaultTotalAgencyAccount
	totalAgency.Password = password
	totalAgency.OperatePassword = password
	totalAgency.Username = inputData.DefaultTotalAgencyName
	totalAgency.Remark = inputData.DefaultTotalAgencyRemark
	totalAgency.Level = 3
	totalAgency.IsSub = 2
	totalAgency.IsDefault = 1
	totalAgency.VideoBalance = 0
	totalAgency.Status = 1
	totalAgency.CreateTime = time.Now().Unix()
	totalAgency.DeleteTime = 0
	count, err = sess.Insert(totalAgency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//增加代理
	agency := new(schema.Agency)
	agency.SiteIndexId = inputData.SiteIndexId
	agency.SiteId = inputData.SiteId
	agency.ParentId = totalAgency.Id
	agency.RoleId = 4
	agency.Account = inputData.DefaultAgencyAccount
	agency.Password = password
	agency.OperatePassword = password
	agency.Username = inputData.DefaultAgencyName
	agency.Level = 4
	agency.IsSub = 2
	agency.IsDefault = 1
	agency.VideoBalance = 0
	agency.Status = 1
	agency.CreateTime = time.Now().Unix()
	agency.DeleteTime = 0
	agency.Remark = inputData.DefaultAgencyRemark
	count, err = sess.Insert(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//给代理人数统计表添加数据(股东)
	shaleHolderCount := new(schema.AgencyCount)
	shaleHolderCount.SiteId = inputData.SiteId
	shaleHolderCount.SiteIndexId = inputData.SiteIndexId
	shaleHolderCount.AgencyId = shaleHolder.Id
	shaleHolderCount.FirstId = 0
	shaleHolderCount.SecondId = 0
	shaleHolderCount.SecondCount = 1
	shaleHolderCount.ThirdCount = 1
	count, err = sess.Insert(shaleHolderCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//给代理人数统计表添加数据(总代	)
	totalAgencyCount := new(schema.AgencyCount)
	totalAgencyCount.SiteId = inputData.SiteId
	totalAgencyCount.SiteIndexId = inputData.SiteIndexId
	totalAgencyCount.AgencyId = totalAgency.Id
	totalAgencyCount.FirstId = shaleHolder.Id
	totalAgencyCount.SecondId = 0
	totalAgencyCount.SecondCount = 0
	totalAgencyCount.ThirdCount = 1
	count, err = sess.Insert(totalAgencyCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//给代理人数统计表添加数据(代理	)
	agencyCount := new(schema.AgencyCount)
	agencyCount.SiteId = inputData.SiteId
	agencyCount.SiteIndexId = inputData.SiteIndexId
	agencyCount.AgencyId = agency.Id
	agencyCount.FirstId = shaleHolder.Id
	agencyCount.SecondId = totalAgency.Id
	agencyCount.SecondCount = 0
	agencyCount.ThirdCount = 0
	count, err = sess.Insert(agencyCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//站点运营统计表
	siteCount := new(schema.SiteCount)
	//获取站点运营统计表数据
	_, err = sess.Where("site_id=?", inputData.SiteId).Where("site_index_id=?", inputData.SiteIndexId).Get(siteCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	siteCount.FirstAgencyCount = siteCount.FirstAgencyCount + 1
	siteCount.SecondAgencyCount = siteCount.SecondAgencyCount + 1
	siteCount.ThirdAgencyCount = siteCount.ThirdAgencyCount + 1
	//修改操作
	count, err = sess.Where("site_id=?", inputData.SiteId).
		Where("site_index_id=?", inputData.SiteIndexId).Update(siteCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//站点下是否有总代或代理
func (*AgencyBean) IsAgency(siteId, siteIndexId string, roleId int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err = sess.Select("id").Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).
		Where("role_id=?", roleId+1).Get(agency)
	return
}

//添加开户人
func (*AgencyBean) AddAccount(this *input.AddAccount) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	site := new(schema.Site)
	sess.Begin()
	agency.Account = this.Account
	agency.Password = this.Password
	agency.Username = this.Username
	agency.Remark = this.Remark
	agency.Level = 1
	agency.Status = 1
	agency.IsSub = 2
	agency.SiteId = this.Site
	agency.SiteIndexId = this.SiteIndex
	agency.RoleId = 1
	agency.IsDefault = 2
	agency.IsLogin = 2
	//添加开户人
	count, err = sess.InsertOne(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	//站点客服信息
	siteInfo := new(schema.SiteInfo)
	siteInfo.SiteId = site.Id
	siteInfo.SiteIndexId = site.IndexId
	count, err = sess.InsertOne(siteInfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	sess.Commit()
	return
}

//获取站点下代理
func (*AgencyBean) GetAgencyId(siteId, siteIndexId string) (*schema.Agency, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err := sess.Select("id").Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).Where("level=?", 4).Where("is_default=?", 1).Where("is_sub=?", 2).Where("delete_time=?", 0).Where("role_id=?", 4).Get(agency)
	return agency, has, err

}

//通过代理账号获取代理id
func (m *AgencyBean) GetAgencyIdByAccount(siteId, account string) (id int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agencySchema := new(schema.Agency)
	sess.Table(agencySchema.TableName()).
		Select("id").
		Where("account = ?", account).
		Where("status = 1").
		Where("delete_time = 0")
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	b, err := sess.Get(&id)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	if !(b) {
		err = errors.New("not found id with account")
	}
	return
}
