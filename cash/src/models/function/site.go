package function

import (
	"errors"
	"fmt"
	"framework/logger"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

//网站运营
type SiteOperateBean struct{}

//增加
func (*SiteOperateBean) Add(siteRequest *input.AddSite) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//赋值
	site := new(schema.Site)
	//查询是否有默认站点
	has, err := sess.Where("id=?", siteRequest.Site).
		Where("is_default=?", 1).
		Where("delete_time=?", 0).
		Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	sess.Begin()
	if !has {
		site.IsDefault = 1
	} else {
		site.IsDefault = 2
	}
	site.Id = siteRequest.Site
	site.IndexId = siteRequest.SiteIndex
	site.DeleteTime = 0
	site.CreateTime = global.GetCurrentTime()
	site.SiteName = siteRequest.SiteName
	site.ComboId = siteRequest.ComboId
	site.DomainUp = siteRequest.DomUp
	site.UpCose = siteRequest.UpCharge
	site.Status = 1
	site.IsDownApp = siteRequest.IsDown
	//插入
	count, err = sess.InsertOne(site)
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
	//站点运营统计
	siteCount := new(schema.SiteCount)
	siteCount.SiteId = site.Id
	siteCount.SiteIndexId = site.IndexId
	count, err = sess.InsertOne(siteCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	//site_domain
	sD := new(schema.SiteDomain)
	sD.IsDefault = 1
	sD.SiteIndexId = siteRequest.SiteIndex
	sD.SiteId = siteRequest.Site
	sD.Domain = siteRequest.Domain
	sD.CreateTime = global.GetCurrentTime()
	sD.DeleteTime = 0
	count, err = sess.InsertOne(sD)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	sess.Commit()
	return
}

//增加(多站点)
func (*SiteOperateBean) MoreSiteAdd(siteRequest *input.AddSiteIn) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//赋值
	site := new(schema.Site)
	//查询是否有默认站点
	has, err := sess.Where("id=?", siteRequest.Site).
		Where("is_default=?", 1).
		Where("delete_time=?", 0).
		Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	sess.Begin()
	if !has {
		site.IsDefault = 1
	} else {
		site.IsDefault = 2
	}
	site.Id = siteRequest.Site
	site.IndexId = siteRequest.SiteIndex
	site.DeleteTime = 0
	site.CreateTime = global.GetCurrentTime()
	site.SiteName = siteRequest.SiteName
	site.DomainUp = siteRequest.DomUp
	site.UpCose = siteRequest.UpCharge
	site.Status = 1
	site.IsDownApp = siteRequest.IsDown
	//插入
	count, err = sess.InsertOne(site)
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
	//站点运营统计
	siteCount := new(schema.SiteCount)
	siteCount.SiteId = site.Id
	siteCount.SiteIndexId = site.IndexId
	count, err = sess.InsertOne(siteCount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	//site_domain
	sD := new(schema.SiteDomain)
	sD.IsDefault = 1
	sD.SiteIndexId = siteRequest.SiteIndex
	sD.SiteId = siteRequest.SiteIndex
	sD.Domain = siteRequest.Domain
	sD.CreateTime = global.GetCurrentTime()
	sD.DeleteTime = 0
	count, err = sess.InsertOne(sD)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	sess.Commit()
	return
}

//删除
func (*SiteOperateBean) Delete(delSite *input.DelSite) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	site.DeleteTime = time.Now().UTC().Unix()
	sess.Where("id=?", delSite.Site)
	count, err = sess.Where("index_id=?", delSite.SiteIndex).Cols("delete_time").Update(site)
	return
}

//修改
func (*SiteOperateBean) Edit(this *input.EditSite) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//赋值
	err := sess.Begin()
	site := new(schema.Site)
	site.UpCose = this.UpCharge
	site.SiteName = this.SiteName
	site.ComboId = this.ComboId
	site.IsDownApp = this.IsDown
	//条件
	count, err := sess.Where("id=?", this.Site).
		Where("index_id=?", this.SiteIndex).
		Where("delete_time=?", 0).
		Cols("site_name,combo_id,domain_up,up_cose,is_down_app").Update(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//查询是site_domain是否有域名
	sD := new(schema.SiteDomain)
	has, err := sess.Where("site_id=?", this.Site).
		Where("site_index_id=?", this.SiteIndex).
		Where("delete_time=?", 0).
		Where("is_default=?", 1).
		Get(sD)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//域名
	sD.Domain = this.Domain
	if !has {
		sD.SiteId = this.Site
		sD.SiteIndexId = this.SiteIndex
		sD.CreateTime = global.GetCurrentTime()
		sD.DeleteTime = 0
		sD.IsDefault = 1
		count, err = sess.InsertOne(sD)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	} else {
		count, err = sess.Where("site_id=?", this.Site).
			Where("site_index_id=?", this.SiteIndex).
			Where("is_default=?", 1).
			Update(sD)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	}
	err = sess.Commit()
	return count, err
}

//根据id获取单个站点
func (*SiteOperateBean) GetSingleSite(indexId, id string) (schema.Site, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.Site
	sess.Where("id=?", id)
	sess.Where("index_id=?", indexId)
	sess.Where("delete_time=?", 0)
	flag, err := sess.Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, err
}

//根据名称获取站点信息
func (*SiteOperateBean) GetSingleSiteByName(siteName string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	sess.Where("site_name=?", siteName)
	sess.Where("delete_time=?", 0)
	flag, err := sess.Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return flag, err
	}
	return flag, err
}

//获取某个开户人下面的所有站点
func (*SiteOperateBean) GetOpuerSite(allSite *input.GetAllSite) ([]back.OpenUserAllSite, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var infolist []back.OpenUserAllSite
	site := new(schema.Site)
	if allSite.OperUser > 0 {
		sess.Where("sales_site.agency_id=?", allSite.OperUser)
	}
	if allSite.ComboId > 0 {
		sess.Where("sales_site.combo_id=?", allSite.ComboId)
	}
	if allSite.Status > 0 {
		sess.Where("sales_site.status=?", allSite.Status)
	}
	if allSite.SiteName != "" {
		sess.Where("sales_site.site_name=?", allSite.SiteName)
	}
	err := sess.Table(site.TableName()).Select("*,sales_combo.combo_name,count(sales_site_domain.id) as num").
		Join("left", "sales_site_info", "sales_site.id=sales_site_info.site_id"+
			" and sales_site.index_id=sales_site_info.site_index_id").
		Join("left", "sales_combo", "sales_site.combo_id=sales_combo.id").
		And("sales_site.delete_time=?", 0).
		Join("left", "sales_site_domain",
			"sales_site.id = sales_site_domain.site_id and sales_site.index_id = sales_site_domain.site_index_id "+
				"and sales_site_domain.delete_time=0").
		GroupBy("sales_site.id,sales_site.index_id").Find(&infolist)
	lLen := len(infolist)
	for i := 0; i < lLen; i++ {
		//额外费用计算
		s := int(infolist[i].ExistDomain) - infolist[i].DomUp
		if s < 0 {
			infolist[i].ExtraCharge = 0
		} else {
			infolist[i].ExtraCharge = float64(s) * infolist[i].UpCharge
		}
	}
	return infolist, err
}

//站点的启用和禁用
func (*SiteOperateBean) UpSiteStatus(siteStatus schema.Site) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	site.Status = siteStatus.Status
	if siteStatus.Id != "" {
		sess.Where("id=?", siteStatus.Id)
	}
	if siteStatus.IndexId != "" {
		sess.Where("index_id=?", siteStatus.IndexId)
	}
	if siteStatus.Status == 1 {
		site.Status = 2
	} else {
		site.Status = 1
	}
	count, err := sess.Where("delete_time=?", 0).Cols("status").Update(site)
	return count, err
}

//根据Index_id获取site详细信息
func (*SiteOperateBean) GetSiteInfomation(siteId, siteIndex, siteName string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var coun int64
	site := new(schema.Site)
	has, err := sess.Where("delete_time=?", 0).
		Where("id=?", siteId).
		Where("index_id=?", siteIndex).
		Get(site)
	if !has && err == nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		coun = 60010
		return coun, err
	}
	has, err = sess.Where("delete_time=?", 0).
		Where("site_name=?", siteName).
		Where("id!=?", siteId).
		Where("index_id!=?", siteIndex).
		Get(site)
	if has && err == nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		coun = 50077
		return coun, err
	}
	return coun, err
}

//获取某个开户人下面所有的前台id
func (*SiteOperateBean) GetSiteIndexId(site *schema.Site) ([]back.IndexBackStruct, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info []back.IndexBackStruct
	sess.Where("id=?", site.Id)
	sess.Where("status=?", 1)
	sess.Where("delete_time=?", 0)
	err := sess.Table(site.TableName()).
		Select("id,index_id,site_name,is_default").
		Find(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, err
	}
	return info, err
}

//查看账号是否存在
func (*SiteOperateBean) GetAccountById(id int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err = sess.Where("id= ?", id).Where("delete_time = 0").Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//初始化管理员密码
func (*SiteOperateBean) InitPassword(id int64, password string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.Password = password
	count, err := sess.Where("id = ?", id).Cols("password").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//根据siteindexid获取站点信息
func (*SiteOperateBean) GetInfoBySiteIndexId(siteId string, siteIndexId string) (schema.Site, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.Site
	sess.Where("id=?", siteId)
	sess.Where("index_id=?", siteIndexId)
	flag, err := sess.Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, err
}

//根据siteindexid获取站点名字
func (*SiteOperateBean) GetSiteNameBySiteIndexId(siteId string, siteIndexId string) (siteName string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteSchema := new(schema.Site)
	sess.Table(siteSchema.TableName())
	sess.Where("id = ?", siteId)
	sess.Where("status = ?", 1)
	sess.Where("delete_time = ?", 0)
	sess.Where("index_id = ?", siteIndexId)
	sess.Select("site_name")
	b, err := sess.Get(&siteName)
	if err != nil {
		return
	}
	if !b {
		err = errors.New("find row 0")
	}
	return
}

//获取某个开户人下面是否有站点
func GetOpenUserSite(siteId string) (flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info := new(schema.Site)
	flag, err = sess.Table(info.TableName()).Where("id=?", siteId).Exist()
	return
}

//上面是开户的时候的site信息，这里的是之后自己完善的部分site_info
//获取单个
func (*SiteOperateBean) GetSingleSiteInfo(siteInfo *schema.SiteCount) (schema.SiteCount, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.SiteCount
	if siteInfo.SiteIndexId != "" {
		sess.Where("site_index_id=?", siteInfo.SiteIndexId)
	} //站点
	if siteInfo.SiteId != "" {
		sess.Where("site_id=?", siteInfo.SiteId)
	}
	flag, err := sess.Get(&info)
	return info, flag, err
}

//查看套餐id是否存在
func (*SiteOperateBean) GetComboId(comboId int64) (has bool, err error) {
	site := new(schema.Site)
	sess := global.GetXorm().NewSession().Table(site.TableName())
	defer sess.Close()
	has, err = sess.Where("combo_id = ?", comboId).Where("status = ?", 1).Where("delete_time = ?", 0).Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//根据siteId查询
func (*SiteOperateBean) GetOneSiteId(this *schema.Site) (has bool, err error) {
	sess := global.GetXorm().NewSession().Table(this.TableName())
	defer sess.Close()
	has, err = sess.Where("id=?", this.Id).Get(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//获取站点下拉
func (*SiteOperateBean) GetSiteDrop(s *input.SelectSiteDrop) (info []back.BackSiteDrop, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	if s.SiteName != "" {
		sess.Where("site_name like ?", "%"+s.SiteName+"%")
	}
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	err = sess.Table(site.TableName()).Find(&info)
	return
}

//站点会员查询
func (m *SiteOperateBean) Member(this *input.SiteMemberInfo, listParams *global.ListParams) (members []*back.SiteMemberInfo, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	if this.SiteId != "" {
		sess.Where(member.TableName()+".site_id=?", this.SiteId)
	}
	sess.Where(member.TableName()+".delete_time=?", 0)
	if this.SiteIndexId != "" {
		sess.Where(member.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where(member.TableName()+".account=?", this.Account)
	}
	if this.StartTime != "" {
		loc, _ := time.LoadLocation("Local")
		t, _ := time.ParseInLocation("2006-01-02", this.StartTime, loc)
		sess.Where(member.TableName()+".create_time>=?", t.Unix())
		t, _ = time.ParseInLocation("2006-01-02", this.EndTime, loc)
		sess.Where(member.TableName()+".create_time<=?", t.Unix())
	}
	switch this.Device {
	case 1:
		{
			sess.Where(member.TableName()+".pc_status=?", 1)
		}
	case 2:
		{
			sess.Where(member.TableName()+".wap_status=?", 1)
		}
	case 3:
		{
			sess.Where(member.TableName()+".ios_status=?", 1)
		}
	case 4:
		{
			sess.Where(member.TableName()+".android_status=?", 1)
		}
	}
	if this.Ip != "" {
		sess.Where(member.TableName()+".login_ip=?", this.Ip)
	}
	if this.Status != 0 {
		sess.Where(member.TableName()+".status=?", this.Status)
	}
	sess.Where(member.TableName() + ".delete_time = 0")
	conds := sess.Conds()
	listParams.Make(sess)
	lL := new(schema.LoginLog)
	sess.Select(member.TableName() + ".id," + member.TableName() + ".site_id," + member.TableName() +
		".site_index_id," + member.TableName() + ".account," + member.TableName() + ".realname," +
		member.TableName() + ".balance," + member.TableName() + ".status," + member.TableName() + ".login_ip," +
		member.TableName() + ".create_time," + lL.TableName() + ".device," + member.TableName() +
		".pc_status," + member.TableName() + ".wap_status," + member.TableName() + ".ios_status," +
		member.TableName() + ".android_status")
	err = sess.Table(member.TableName()).
		Join("LEFT", lL.TableName(), member.TableName()+".account="+lL.TableName()+".account").
		GroupBy(member.TableName() + ".account").
		Find(&members)
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return
	}
	var mN []int64
	if len(members) > 0 {
		for _, v := range members {
			mN = append(mN, v.Id)
		}
	}
	//根据会员id去获取视讯余额
	mB := new(schema.MemberProductClassifyBalance)
	pL := new(schema.Platform)
	var vB []back.MemberVideoBalanceBack
	err = sess.Table(mB.TableName()).Join("LEFT", pL.TableName(), mB.TableName()+".platform_id="+
		pL.TableName()+".id").In(mB.TableName()+".member_id", mN).Find(&vB)
	//取出所有的平台
	var vc back.MemberVideoBalanceBack
	var vd []back.MemberVideoBalanceBack
	if len(vB) > 0 {
		for _, n := range vB {
			for l, o := range members {
				if n.MemberId == o.Id {
					vc.MemberId = n.MemberId
					vc.Balance = n.Balance
					vc.Platform = n.Platform
					vc.PlatformId = n.PlatformId
					members[l].MemberVideoBalance = append(members[l].MemberVideoBalance, vc)
				}
			}
			if len(vd) > 0 {
				var i int64
				for _, k := range vd {
					if n.PlatformId == k.PlatformId {
						i = i + 1
					}
				}
				if i < 1 {
					vc.MemberId = n.MemberId
					vc.Balance = n.Balance
					vc.Platform = n.Platform
					vc.PlatformId = n.PlatformId
					vd = append(vd, vc)
				}
			} else {
				vc.MemberId = n.MemberId
				vc.Balance = n.Balance
				vc.Platform = n.Platform
				vc.PlatformId = n.PlatformId
				vd = append(vd, vc)
			}
		}

	}
	//补0
	if len(members) > 0 {
		for c, j := range members {
			for _, r := range vd {
				var h int64
				for _, y := range j.MemberVideoBalance {
					if r.Platform == y.Platform {
						h = h + 1
					}
				}
				if h < 1 {
					vc.MemberId = j.Id
					vc.Balance = 0
					vc.Platform = r.Platform
					vc.PlatformId = r.PlatformId
					members[c].MemberVideoBalance = append(members[c].MemberVideoBalance, vc)
				}
			}
		}
	}
	//排序
	for _, b := range members {
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
	//总会员数
	count, err = sess.Table(member.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return
	}
	return
}

func (*SiteOperateBean) StatusCount() (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	count, err := sess.Table(member.TableName()).Where("status=1").Where("delete_time=0").Count()
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return count, err
	}
	return count, err
}

//会员资料详情
func (*SiteOperateBean) MemberInfo(id int64) (*back.MemberDetail, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	memberInfo := new(schema.MemberInfo)
	md := new(back.MemberDetail)
	sql := fmt.Sprintf("%s.member_id = %s.id", memberInfo.TableName(), member.TableName())
	has, err := sess.Table(member.TableName()).Join("LEFT", memberInfo.TableName(), sql).
		Where("id=?", id).Get(md)
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return md, has, err
	}
	return md, has, err
}

//会员银行卡资料详情
func (*SiteOperateBean) MemberBankInfoById(id int64) ([]back.MemberBank, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.MemberBank)
	banK := new(schema.Bank)
	var mb []back.MemberBank
	err := sess.Table(member.TableName()).
		Join("LEFT", banK.TableName(), member.TableName()+
			".bank_id="+banK.TableName()+".id").
		Where(member.TableName()+".member_id=?", id).
		Where(member.TableName()+".delete_time=?", 0).Find(&mb)
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return mb, err
	}
	return mb, err
}

//获取站点自助优惠离开关列表
func (*SiteOperateBean) GetSiteSelfSwitch(selectDis *input.SelfDiscountSwitch) (data []back.AutoDiscountSwitch, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if selectDis.SiteId != "" {
		sess.Where("id=?", selectDis.SiteId)
	}
	if selectDis.IndexId != "" {
		sess.Where("index_id=?", selectDis.IndexId)
	}
	sess.Where("status=?", 1)
	autoDiscountSwitch := new(schema.Site)
	err = sess.Table(autoDiscountSwitch.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取站点套餐占成比
func (*SiteOperateBean) SiteCombo(SiteId []string) (sdata map[string]back.ReportExport, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	if len(SiteId) != 0 {
		sess.In(site.TableName()+".id", SiteId)
	}

	sess.Where(site.TableName()+".index_id=?", "a")
	sess.Where(site.TableName()+".delete_time=?", 0)
	sess.Where(site.TableName()+".status=?", 1)

	sess.Join("LEFT", "sales_combo_product", site.TableName()+".combo_id = sales_combo_product.combo_id")
	sess.Join("LEFT", "sales_combo", site.TableName()+".combo_id = sales_combo.id")
	data := []back.SiteIndexList{}
	err = sess.Table(site.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return sdata, err
	}
	sdata = make(map[string]back.ReportExport)
	item := make(map[string]bool)

	for _, v := range data {

		if item[v.SiteId] == false {
			sdataInfo := back.ReportExport{}
			sdataInfo.SiteId = v.SiteId
			sdataInfo.SiteIndexId = v.SiteIndexId
			sdataInfo.SiteName = v.SiteName
			sdataInfo.ComboName = v.ComboName

			for _, v1 := range data {

				if v.SiteId == v1.SiteId {
					info := back.ReportExportList{}
					info.ProductId = v1.ProductId
					info.Proportion = v1.Proportion

					sdataInfo.List = append(sdataInfo.List, info)

				}
			}
			sdata[v.SiteId] = sdataInfo
			item[v.SiteId] = true
		}
	}
	return sdata, err
}

//获取site_id列表 站点列表
func (*SiteOperateBean) SiteList() (data []back.InSiteList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)

	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	sess.GroupBy("id")

	err = sess.Table(site.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//根据agency_id获取站点
func (*SiteOperateBean) BeSiteOne(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	has, err := sess.Table(site.TableName()).Where("delete_time=?", 0).
		Where("agency_id=?", id).Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//前台-文案类型
func (*SiteOperateBean) SiteCopyType() ([]schema.IwordCate, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []schema.IwordCate
	err := sess.Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//多站点-前台首页文案列表
func (*SiteOperateBean) SiteCopyList(this *input.SiteCopyList) (data []back.CopyList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	switch this.Itype {
	case 1:
		sess.In("itype", []int{11, 12, 13, 14, 15, 16, 20})
	case 2:
		sess.In("itype", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	case 3:
		sess.In("itype", []int{17, 18, 19})
	}
	copylist := new(schema.Iword)
	sess.Table(copylist.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//多站点-前台首页文案详情
func (*SiteOperateBean) SiteCopyListInfo(this *input.SiteCopyListInfoOne) (*back.CopyListInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	copylist := new(schema.Iword)
	info := new(back.CopyListInfo)
	has, err := sess.Table(copylist.TableName()).
		Where("id=?", this.Id).
		Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err
	}
	return info, has, err
}

//多站点-前台文案添加
func (*SiteOperateBean) SiteCopyAdd(this *input.SiteCopyAdd) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	iwordType := new(schema.IwordCate)
	iwordType.Id = this.Itype
	ok, err := sess.Get(iwordType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if !ok {
		return 0, errors.New("类型不存在")
	}
	copylist := new(schema.Iword)
	copylist.TopId = 0
	copylist.SiteId = this.SiteId
	copylist.SiteIndexId = this.SiteIndexId
	copylist.Title = this.Title
	copylist.TitleColor = ""
	copylist.Content = this.Content
	copylist.Url = ""
	copylist.Img = ""
	copylist.State = 1
	copylist.Sort = this.Sort
	copylist.From = this.From
	copylist.Itype = this.Itype
	copylist.TypeName = iwordType.Name
	copylist.AddTime = time.Now().Unix()
	data, err = sess.Table(copylist.TableName()).Insert(copylist)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//轮播查询 state 状态 1-启用 2-关闭,ftype 类型 1-PC端 2-WAP端
func (*SiteOperateBean) FlashList(siteId, siteIndexId string, state, fType int64) (data []back.FlashList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flash := new(schema.SiteFlash)
	if state > 0 {
		sess.Where("state = ?", state)
	}
	if fType > 0 {
		sess.Where("ftype = ?", fType)
	}
	err = sess.Table(flash.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Asc("sort").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//轮播新增
func (*SiteOperateBean) FlashAdd(this *input.FlashAdd) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//统计该站点已有几张轮播图
	flash := new(schema.SiteFlash)
	count, err := sess.Table(flash.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("ftype=?", this.Ftype).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if count >= 5 {
		data = 5
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	data, err = sess.Table(flash.TableName()).Insert(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return
}

//站点logo图片管理
func (*SiteOperateBean) LogoList(this *input.LogoList) (data []back.LogoList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	logo_list := new(schema.InfoLogo)
	err = sess.Table(logo_list.TableName()).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询站点logo  form >1pc2wap
func (m *SiteOperateBean) Logo(siteId, siteIndexId string, form int64) (logoUrl string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.InfoLogo)
	b, err := sess.Table(infoLogoSchema.TableName()).
		Where("site_id = ?", siteId).
		Where("type = ?", 5).
		Where("state = ?", 1).
		Where("form = ?", 1).
		Where("site_index_id = ?", siteIndexId).
		Select("logo_url").
		Get(&logoUrl)
	if err != nil {
		return
	}
	if !b {
		err = errors.New("not found log_url")
	}
	return
}

//站点logo是否存在
func (*SiteOperateBean) LogoInfo(this *input.LogoAdd) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	logoList := new(schema.InfoLogo)
	has, err = sess.Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("state=?", 1).
		Where("form=?", this.Form).
		Get(logoList)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//站点logo图片添加
func (*SiteOperateBean) LogoAdd(this *input.LogoAdd) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	logoList := new(schema.InfoLogo)
	data, err = sess.Table(logoList.TableName()).Insert(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return
}

//文案修改
func (*SiteOperateBean) SiteCopyUpdate(this *input.CopyUpdate) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	copylist := new(schema.Iword)
	data, err = sess.Table(copylist.TableName()).
		Where("id=?", this.Id).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//修改站点上线时间
func (*SiteOperateBean) UpSiteOnlineTime(siteOnline *input.SiteOnline) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	times, code := global.FormatTime2Timestamp2(siteOnline.OnlineTime)
	if code != 0 {
		return 0, errors.New("时间转换错误")
	}
	sess.Where("id=?", siteOnline.SiteId)
	sess.Where("index_id=?", siteOnline.SiteIndexId)
	site := new(schema.Site)
	site.OnlineTime = times
	count, err := sess.Cols("online_time").Update(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改站点域名配置
func (*SiteOperateBean) ChangeSiteDomain(inputData *input.SiteDomainEdit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err := sess.Begin()
	//首先更改site
	site := new(schema.Site)
	site.SiteName = inputData.SiteName
	site.IsDownApp = inputData.IsDownApp
	site.UpCose = inputData.UpCose
	site.DomainUp = inputData.DomainUp
	count, err := sess.Where("id=?", inputData.SiteId).
		Where("index_id=?", inputData.SiteIndexId).
		Cols("site_name,is_down_app,domain_up,up_cose").
		Update(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	siteDomain := new(schema.SiteDomain)
	siteDomain.Domain = inputData.Domain
	//更改site_domain,首先判断该域名配置是否存在，存在修改，不存在添加
	if inputData.Id == 0 {
		//不存在
		//todo 添加的时候这些值应该如何处理，暂时全部给的确认得值
		siteDomain.SiteId = inputData.SiteId
		siteDomain.SiteIndexId = inputData.SiteIndexId
		siteDomain.IsDefault = 1
		//siteDomain.FileName = nil
		siteDomain.CreateTime = time.Now().Unix()
		siteDomain.DeleteTime = 0
		count, err = sess.Insert(siteDomain)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	} else {
		count, err = sess.Where("id=?", inputData.Id).
			Where("site_id=?", inputData.SiteId).
			Where("site_index_id=?", inputData.SiteIndexId).
			Cols("domain").
			Update(siteDomain)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return count, err
		}
	}
	err = sess.Commit()
	return count, err
}

//多站点-列表
func (*SiteOperateBean) GetSiteMoreList(siteId string) (siteMoreList []back.SiteMoreList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	var siteInfo []schema.Site
	siteDomain := new(schema.SiteDomain)
	err = sess.Table(siteDomain.TableName()).Select("id,domain,is_default as isDefaultD,site_index_id,is_default,site_id").Where("site_id=?", siteId).Find(&siteMoreList)
	if err != nil {
		return
	}
	err = sess.Table(site.TableName()).Where("id=?", siteId).Find(&siteInfo)

	for k, v := range siteInfo {
		for k1, v1 := range siteMoreList {
			if v.Id == v1.SiteId && v.IndexId == v1.SiteIndexId {
				siteMoreList[k1].CreateTime = siteInfo[k].CreateTime
				siteMoreList[k1].IsDefault = siteInfo[k].IsDefault
				siteMoreList[k1].Status = siteInfo[k].Status
				siteMoreList[k1].IsDownApp = siteInfo[k].IsDownApp
				siteMoreList[k1].UpCose = siteInfo[k].UpCose
				siteMoreList[k1].DomainUp = siteInfo[k].DomainUp
			}
		}
	}

	return
}

//多站点-添加代理
func (sob *SiteOperateBean) PostSiteMore(this *input.SiteMoreAdd, parentId int64) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	agency := new(schema.Agency)
	has, err := sob.GetSecondAgency(this.SiteId, this.SiteIndexId)
	if err != nil {
		sess.Rollback()
		return
	}
	if has {
		agency.IsDefault = 2
	} else {
		agency.IsDefault = 1
	}
	agency.ParentId = parentId
	agency.RoleId = 2
	agency.Remark = this.Remark
	//默认密码123456
	password, err := global.MD5ByStr("123456", global.EncryptSalt)
	if err != nil {
		return 0, err
	}
	agency.Password = password
	agency.Username = this.Username
	agency.SiteId = this.SiteId
	agency.SiteIndexId = this.SiteIndexId
	agency.Account = this.Account
	agency.CreateTime = global.GetCurrentTime()
	agency.IsSub = 2
	agency.Level = 2
	agency.Status = 1
	count, err = sess.InsertOne(agency)
	if err != nil {
		sess.Rollback()
		return
	}
	//代理人数统计表
	AgencyCount := new(schema.AgencyCount)
	//添加数据
	AgencyCount.SiteId = this.SiteId
	AgencyCount.SiteIndexId = this.SiteIndexId
	AgencyCount.AgencyId = agency.Id
	count, err = sess.Insert(AgencyCount)
	if err != nil {
		sess.Rollback()
		return count, err
	}
	//站点运营统计表
	siteCount := new(schema.SiteCount)
	//获取站点运营统计表数据
	_, err = sess.Where("site_id=?", this.SiteId).Where("site_index_id=?", this.SiteIndexId).Get(siteCount)
	if err != nil {
		sess.Rollback()
		return
	}
	siteCount.FirstAgencyCount = siteCount.FirstAgencyCount + 1
	//修改操作
	count, err = sess.Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).Update(siteCount)
	if err != nil {
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return
}

//根据站点获取开户人id
func (*SiteOperateBean) GetIdBySite(siteId string) (id int64, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err = sess.Select("id").Where("site_id=?", siteId).Where("role_id=1").
		Where("delete_time=0").Where("status=1").Where("parent_id=0").
		Where("delete_time=0").Get(agency)
	id = agency.Id
	return
}

//站点下是否有股东
func (*SiteOperateBean) GetSecondAgency(siteId, siteIndexId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	has, err = sess.Select("id").Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).
		Where("role_id=2").Get(agency)
	return
}

//报表负数查询
func (*SiteOperateBean) NegativeList(this *input.ReportNegativeList) (data []back.ReportNegativeList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	negative := new(schema.SiteReportNegative)
	negativePro := new(schema.SiteReportNegativeProduct)
	product := new(schema.Product)
	t1 := negative.TableName()
	t2 := negativePro.TableName()
	t3 := product.TableName()
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	data1 := make([]back.ReportNegativeAndProducts, 0)
	sql1 := fmt.Sprintf("%s.id = %s.negative_id", negative.TableName(), negativePro.TableName())
	sql2 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), negativePro.TableName())
	err = sess.Table(negative.TableName()).
		Select(t1+".id,site_id,site_index_id,years,"+t2+".product_id,"+t3+".product_name,report_win").
		Join("LEFT", negativePro.TableName(), sql1).
		Join("LEFT", product.TableName(), sql2).OrderBy(t1 + ".id").
		Find(&data1)
	if err != nil {
		return
	}
	var checkid int64 = 0
	negatives := make([]back.ReportNegativeList, 0)
	negativePros := make([]back.ReportNegativeProduct, 0)
	negative2 := new(back.ReportNegativeList)
	negativePro2 := new(back.ReportNegativeProduct)
	for _, d := range data1 {
		if checkid != d.Id {
			//id不同时，组装，总组装
			if checkid != 0 {
				negative2.Products = negativePros         //商品列表 组装到总组装products参数
				negatives = append(negatives, *negative2) //总组装
			}
			negativePros = nil
			checkid = d.Id

			//组装站点id，站点前台id,年月
			negative2.SiteId = d.SiteId
			negative2.SiteIndexId = d.SiteIndexId
			negative2.Years = d.Years
			//组装商品列表
			negativePro2.ProductId = d.ProductId
			negativePro2.ProductName = d.ProductName
			negativePro2.ReportWin = d.ReportWin
			negativePros = append(negativePros, *negativePro2)
		} else {
			//id相同时，只组装商品列表
			negativePro2.ProductId = d.ProductId
			negativePro2.ProductName = d.ProductName
			negativePro2.ReportWin = d.ReportWin
			negativePros = append(negativePros, *negativePro2)
		}
	}
	negative2.Products = negativePros
	if checkid != 0 {
		negatives = append(negatives, *negative2)
	}
	data = negatives
	return
}

//报表负数新增
func (*SiteOperateBean) NegativeAdd(this *input.ReportNegativeAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	negative := new(schema.SiteReportNegative)
	var negativePro schema.SiteReportNegativeProduct
	var negativePros []schema.SiteReportNegativeProduct
	negative.SiteIndexId = this.SiteIndexId
	negative.SiteId = this.SiteId
	negative.Years = this.Years
	negative.State = 1
	sess.Begin()
	//新增
	count, err = sess.Table(negative.TableName()).InsertOne(negative)
	if err != nil {
		sess.Rollback()
		return
	}
	for k := range this.Products {
		negativePro.NegativeId = negative.Id
		negativePro.ProductId = this.Products[k].ProductId
		negativePro.ReportWin = this.Products[k].ReportWin
		negativePros = append(negativePros, negativePro)
	}
	//新增
	count, err = sess.Table(negativePro.TableName()).Insert(negativePros)
	if err != nil {
		sess.Rollback()
		return
	}
	sess.Commit()
	return

}

//报表负数修改
func (*SiteOperateBean) NegativeEdit(this *input.ReportNegativeEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	negative := new(schema.SiteReportNegative)
	var negativePro schema.SiteReportNegativeProduct
	var negativePros []schema.SiteReportNegativeProduct
	negative.State = this.State
	sess.Begin()
	count, err = sess.Table(negative.TableName()).
		Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Where("years=?", this.Years).
		Cols("state").
		Update(negative)
	if err != nil {
		sess.Rollback()
		return
	}
	//删除设置id下的所有数据,site_id与years联合唯一
	count, err = sess.Table(negativePro.TableName()).
		Where("negative_id = ?", this.Id).
		Delete(negativePro)
	if err != nil {
		sess.Rollback()
		return
	}
	for k := range this.Products {
		negativePro.NegativeId = this.Id
		negativePro.ProductId = this.Products[k].ProductId
		negativePro.ReportWin = this.Products[k].ReportWin
		negativePros = append(negativePros, negativePro)
	}
	//新增
	count, err = sess.Table(negativePro.TableName()).Insert(negativePros)
	if err != nil {
		sess.Rollback()
		return
	}
	sess.Commit()
	return
}

//根据site_id 和期数查询上月报表负数

func (*SiteOperateBean) SiteNegativeList(SiteId []string, Qishu string) (sdata map[string]back.ReportNegativeList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	negative := new(schema.SiteReportNegative)
	negativePro := new(schema.SiteReportNegativeProduct)

	sess.In(negative.TableName()+".site_id", SiteId)
	sess.Where(negative.TableName()+".years = ?", Qishu)
	sess.Where(negative.TableName()+".state = ?", 1)

	sql := negative.TableName() + ".id = " + negativePro.TableName() + ".negative_id"

	sess.Join("LEFT", negativePro.TableName(), sql)

	var data []back.ReportNegativeAndProducts
	err = sess.Table(negative.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return sdata, err
	}
	sdata = make(map[string]back.ReportNegativeList)
	item := make(map[string]bool)
	for _, v := range data {

		if item[v.SiteId] == false && v.SiteIndexId == "a" {
			//因为站点套餐默认用a站 所以过滤掉其他的

			info := back.ReportNegativeList{}
			info.SiteId = v.SiteId
			info.SiteIndexId = v.SiteIndexId
			info.Years = v.Years

			for _, v1 := range data {
				if v1.Years == v.Years && v1.SiteId == v.SiteId && v1.SiteIndexId == "a" {
					productsInfo := back.ReportNegativeProduct{}
					productsInfo.ProductId = v.ProductId
					productsInfo.ReportWin = v.ReportWin
					info.Products = append(info.Products, productsInfo)
				}
			}
			sdata[v.SiteId] = info
			item[v.SiteId] = true
		}
	}
	return sdata, err
}

//查询所有的皮肤
func (m *SiteOperateBean) GetThemeAll() (themes []*back.Theme, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteSchema := new(schema.Site)
	err = sess.Table(siteSchema.TableName()).
		Select("id as site_id,index_id as site_index_id,theme as theme_name").
		Find(&themes)
	return
}

//查询个前台单站点的皮肤 ,因包冲突,迁移到page.siteTheme
func (m *SiteOperateBean) GetThemeBySiteId(siteId string, siteIndexId string) (themeDirName string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteSchema := new(schema.Site)
	b, err := sess.Table(siteSchema.TableName()).
		Select("theme").
		Where("id = ?", siteId).
		Where("index_id=?", siteIndexId).
		Get(&themeDirName)
	if err != nil {
		return
	}
	if !(b) {
		err = errors.New("get 0 row")
	}
	return
}

//站点详情
func (*SiteOperateBean) SiteInfoBySiteAndSiteIndex(this *input.SiteDomainInfo) (*back.SiteInfoBySiteAndSiteIndexBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	sD := new(schema.SiteDomain)
	info := new(back.SiteInfoBySiteAndSiteIndexBack)
	has, err := sess.Table(site.TableName()).
		Join("LEFT",
			sD.TableName(), site.TableName()+".id="+
				sD.TableName()+".site_id AND "+
				site.TableName()+".index_id="+
				sD.TableName()+".site_index_id AND "+
				sD.TableName()+".is_default=1 AND "+
				sD.TableName()+".delete_time=0").
		Where(site.TableName()+".id=?", this.Site).
		Where(site.TableName()+".index_id=?", this.SiteIndex).
		Where(site.TableName()+".delete_time=?", 0).
		Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err
	}
	return info, has, err
}

//获取全站维护数据
func (*SiteOperateBean) GetSiteMaintenance() ([]back.SiteMaintenance, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sm := make([]back.SiteMaintenance, 0)
	siteModule := new(schema.SiteModule)
	err := sess.Table(siteModule.TableName()).Find(&sm)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return sm, err
	}
	return sm, err
}

//设置全站维护数据
func (*SiteOperateBean) PutSiteMaintenance(this *input.SiteMaintenance) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sm := new(schema.SiteModule)
	sm.Content = this.Content
	sm.SiteIds = this.SiteIdS
	count, err := sess.Table(sm.TableName()).Where("id = ?", this.Id).Update(sm)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取站点数据
func (*SiteOperateBean) SiteSiteIndexIdBy() ([]back.SiteSiteIndexBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	var data []back.SiteSiteIndexBack
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	err := sess.Table(site.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取全站维护对应的站点
func (*SiteOperateBean) GetSiteIdS(id int64) (*back.SiteIdS, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sm := new(back.SiteIdS)
	siteModule := new(schema.SiteModule)
	_, err := sess.Table(siteModule.TableName()).Where("id=?", id).Get(sm)
	return sm, err
}

//查看站点下是否有开户人
func (*SiteOperateBean) GetAccount(siteId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	this := new(schema.Agency)
	has, err = sess.Where("site_id=?", siteId).Where("delete_time=0").Where("is_sub=2").
		Where("parent_id=0").Where("level=1").Get(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}
