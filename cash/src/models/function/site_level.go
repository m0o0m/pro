package function

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
	"time"
)

type SiteLevelBean struct{}

var siteLevelPlatform = new(LevelPlatformBean)

//添加站点层级
func (*SiteLevelBean) AddSiteLevel(this *input.AddSiteLevel) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteLevel := new(schema.SiteLevel)
	siteLevelPla := new(schema.SiteLevelPlatform)
	siteLevelPlas := make([]schema.SiteLevelPlatform, 0)
	siteLevel.Lid = this.Lid
	siteLevel.LevelName = this.LevelName
	siteLevel.Talk = this.Talk
	siteLevel.Remark = this.Remark
	siteLevel.DoTime = time.Now().Unix()
	siteLevel.State = 1
	siteLevel.SiteLevel = ""
	sess.Begin()
	count, err = sess.Table(siteLevel.TableName()).InsertOne(siteLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	for _, p := range this.Platforms {
		siteLevelPla.LevelId = siteLevel.Id
		siteLevelPla.PlatformId = p.PlatformId
		siteLevelPla.Proportion = p.Proportion
		siteLevelPlas = append(siteLevelPlas, *siteLevelPla)
	}
	count, err = sess.Table(siteLevelPla.TableName()).Insert(siteLevelPlas)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	sess.Commit()
	return

}

//修改站点层级
func (*SiteLevelBean) EditSiteLevel(this *input.EditSiteLevel) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteLevel := new(schema.SiteLevel)
	siteLevelPla := new(schema.SiteLevelPlatform)
	siteLevelPlas := make([]schema.SiteLevelPlatform, 0)
	siteLevel.LevelName = this.LevelName
	siteLevel.Talk = this.Talk
	siteLevel.Remark = this.Remark
	sess.Begin()
	count, err = sess.Table(siteLevel.TableName()).Where("id = ?", this.Id).Cols("level_name,talk,remark").Update(siteLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	//删除设置id下的所有数据
	count, err = sess.Table(siteLevelPla.TableName()).Where("level_id = ?", this.Id).Delete(siteLevelPla)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	for _, p := range this.Platforms {
		siteLevelPla.LevelId = this.Id
		siteLevelPla.PlatformId = p.PlatformId
		siteLevelPla.Proportion = p.Proportion
		siteLevelPlas = append(siteLevelPlas, *siteLevelPla)
	}
	count, err = sess.Table(siteLevelPla.TableName()).Insert(siteLevelPlas)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return
	}
	sess.Commit()
	return

}

//删除站点层级
func (*SiteLevelBean) DelSiteLevel(this *input.DelSiteLevel) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteLevel := new(schema.SiteLevel)
	siteLevel.State = 2
	count, err = sess.Table(siteLevel.TableName()).Where("id = ?", this.Id).Cols("state").Update(siteLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点层级列表
func (*SiteLevelBean) ListSiteLevel(this *input.SiteList) (data back.AllLevelBySite, count int64, err error) {
	siteLevel := new(schema.SiteLevel)
	siteLevelPla := new(schema.SiteLevelPlatform)
	platform := new(schema.Platform)
	sess := global.GetXorm().Table(siteLevel.TableName())
	defer sess.Close()
	sess.Where("sales_site_level.state=?", 1)
	sess.Where("sales_site_level.site_level like?", "%"+this.SiteId+"%")
	conds := sess.Conds()
	sess.Where("sales_site_level_platform.proportion!=?", 0)
	data1 := make([]back.ListSiteLevelAll, 0)
	sql := fmt.Sprintf("%s.id = %s.level_id", siteLevel.TableName(), siteLevelPla.TableName())
	sql2 := fmt.Sprintf("%s.platform_id = %s.id", siteLevelPla.TableName(), platform.TableName())
	err = sess.Table(siteLevel.TableName()).
		Select("sales_site_level.id,lid,level_name,site_level,talk,remark,platform_id,platform,proportion").
		Join("LEFT", siteLevelPla.TableName(), sql).
		Join("LEFT", platform.TableName(), sql2).OrderBy("sales_site_level.id").
		Find(&data1)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	var checkid int64 = 0
	siteLevels := make([]back.ListSiteLevel, 0)
	siteLevelPlas := make([]back.SiteLevelProductRate, 0)
	siteLevel2 := new(back.ListSiteLevel)
	siteLevelPla2 := new(back.SiteLevelProductRate)
	for _, d := range data1 {
		if checkid != d.Id { //id不同时，站点层级，组装平台列表，总组装
			if checkid != 0 {
				siteLevel2.Platforms = siteLevelPlas         //平台列表 组装到总组装Platforms参数
				siteLevels = append(siteLevels, *siteLevel2) //总组装
			}
			siteLevelPlas = nil
			checkid = d.Id

			//站点层级
			siteLevel2.Id = d.Id
			siteLevel2.Lid = d.Lid
			siteLevel2.LevelName = d.LevelName
			siteLevel2.SiteLevel = d.SiteLevel
			siteLevel2.Talk = d.Talk
			siteLevel2.Remark = d.Remark
			//组装平台列表
			siteLevelPla2.PlatformId = d.PlatformId
			siteLevelPla2.Platform = d.Platform
			siteLevelPla2.Proportion = d.Proportion
			siteLevelPlas = append(siteLevelPlas, *siteLevelPla2)
		} else { //id相同时，只组装平台列表
			siteLevelPla2.PlatformId = d.PlatformId
			siteLevelPla2.Platform = d.Platform
			siteLevelPla2.Proportion = d.Proportion
			siteLevelPlas = append(siteLevelPlas, *siteLevelPla2)
		}

	}
	siteLevel2.Platforms = siteLevelPlas
	if checkid != 0 {
		siteLevels = append(siteLevels, *siteLevel2)
	}
	data.ListSiteLevel = siteLevels
	count, err = sess.Table(siteLevel.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	var strr []back.SiteLevelProductPlatform
	var strro back.SiteLevelProductPlatform
	for _, v := range data.ListSiteLevel {
		for _, k := range v.Platforms {
			il := 0
			if len(strr) > 0 {
				for _, g := range strr {
					if k.PlatformId == g.PlatformId {
						il = il + 1
					}
				}
				if il < 1 {
					strro.PlatformId = k.PlatformId
					strro.Platform = k.Platform
					strr = append(strr, strro)
				}
			} else {
				strro.PlatformId = k.PlatformId
				strro.Platform = k.Platform
				strr = append(strr, strro)
			}
		}
	}
	data.PingTai = strr
	var nh back.SiteLevelProductRate
	for _, v := range data.PingTai {
		for m, n := range data.ListSiteLevel {
			var q int64
			if len(n.Platforms) > 0 {
				for _, u := range n.Platforms {
					if v.PlatformId == u.PlatformId {
						q = q + 1
					}
				}
			} else {
				nh.PlatformId = v.PlatformId
				nh.Platform = v.Platform
				nh.Proportion = 0
				data.ListSiteLevel[m].Platforms = append(data.ListSiteLevel[m].Platforms, nh)
			}
			if q < 1 {
				nh.PlatformId = v.PlatformId
				nh.Platform = v.Platform
				nh.Proportion = 0
				data.ListSiteLevel[m].Platforms = append(data.ListSiteLevel[m].Platforms, nh)
			}
		}
	}
	for _, b := range data.ListSiteLevel {
		if len(b.Platforms) > 0 {
			//排序
			for i := 0; i < len(b.Platforms)-1; i++ {
				for j := i + 1; j < len(b.Platforms); j++ {
					if b.Platforms[i].PlatformId > b.Platforms[j].PlatformId {
						b.Platforms[i], b.Platforms[j] = b.Platforms[j], b.Platforms[i]
					}
				}
			}
		}
	}
	return data, count, err
}

//获取单个层级详情
func (*SiteLevelBean) DetailSiteLevel(this *input.DelSiteLevel) (data back.DetailSiteLevel, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var has bool
	siteLevel := new(schema.SiteLevel)
	sess.Where("id=?", this.Id)
	sess.Where("state=?", 1)
	//conds := sess.Conds()
	siteLevelPlas := make([]schema.SiteLevelPlatform, 0)
	siteLevelPla := new(schema.SiteLevelPlatform)
	platform := new(schema.Platform)
	platforms := make([]schema.Platform, 0)
	//查询站点层级表
	has, err = sess.Table(siteLevel.TableName()).Get(siteLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	if !has {
		count = 0
		return data, count, err
	}
	data.Id = siteLevel.Id
	data.Remark = siteLevel.Remark
	data.LevelName = siteLevel.LevelName
	data.Lid = siteLevel.Lid
	data.Talk = siteLevel.Talk
	//查询站点对应平台占成比
	sess.Where("level_id=?", this.Id)
	err = sess.Table(siteLevelPla.TableName()).Cols("level_id,platform_id,proportion").Find(&siteLevelPlas)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count = int64(1)
	ids := make([]int64, 0)
	if len(siteLevelPlas) > 0 {
		for _, v := range siteLevelPlas {
			ids = append(ids, v.PlatformId)
			list := new(back.SiteLevelProductRate)
			list.PlatformId = v.PlatformId
			list.Proportion = v.Proportion
			data.Platforms = append(data.Platforms, *list)
		}
	} else {
		return data, count, err
	}

	//查询商品表
	sess.In("id", ids)
	err = sess.Table(platform.TableName()).Find(&platforms)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	for _, v := range platforms {
		for k2, v2 := range data.Platforms {
			if v.Id == v2.PlatformId {
				data.Platforms[k2].Platform = v.Platform
				break
			}
			fmt.Println(v.Platform)
		}
	}

	return data, count, err
}

//初始化站点层级设置，将未分层站点分到默认分层
func (*SiteLevelBean) InitSiteLevel() (code int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteLevel := new(schema.SiteLevel)
	var count int64
	count, err = sess.Table(siteLevel.TableName()).Where("level_name = ?", "默认未分层").Count()
	if count == 0 {
		//err = errors.New("请先手动添加层级名为 默认未分层 的站点层级")
		code = 10133
		return code, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, err
	}
	siteLevels := make([]schema.SiteLevel, 0)
	err = sess.Table(siteLevel.TableName()).
		Select("site_level").
		Where("site_level<>?", "").
		Find(&siteLevels)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, err
	}
	arrSiteId := make([]string, 0)
	for _, v := range siteLevels {
		for _, v2 := range strings.Split(v.SiteLevel, ",") {
			arrSiteId = append(arrSiteId, v2)
		}
	}
	site := new(schema.Site)
	sites := make([]schema.Site, 0)
	err = sess.Table(site.TableName()).
		Select("id").
		NotIn("id", arrSiteId).
		GroupBy("id").
		Find(&sites)
	if len(sites) == 0 {
		//err = errors.New("没有需要添加到默认未分层的站点")
		code = 10134
		return code, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, err
	}
	var siteLevel2 string
	for i, v := range sites {
		if i == len(sites)-1 {
			siteLevel2 = siteLevel2 + v.Id
		} else {
			siteLevel2 = siteLevel2 + v.Id + ","
		}
	}
	siteLevel.SiteLevel = siteLevel2
	count, err = sess.Table(siteLevel.TableName()).
		Where("level_name = ?", "默认未分层").
		Cols("site_level").Update(siteLevel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, err
	}
	return code, err
}

//移动站点层级
func (*SiteLevelBean) MoveSiteLevel(this *input.MoveSiteLevel) (code int64, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteLevel := new(schema.SiteLevel)
	var has bool
	has, err = sess.Table(siteLevel.TableName()).
		Select("site_level").
		Where("id=?", this.OldId).Get(siteLevel)
	if !has {
		//err = errors.New("没找到老站点的站点层级")
		code = 10136
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if strings.Contains(siteLevel.SiteLevel, this.SiteId) == false {
		//err = errors.New("老站点的站点层级不含有传入的站点")
		code = 10137
		return code, count, err
	}
	if siteLevel.SiteLevel == this.SiteId {
		siteLevel.SiteLevel = ""
	} else {
		siteLevel.SiteLevel = strings.Replace(","+siteLevel.SiteLevel+",", ","+this.SiteId+",", ",", -1) //去除老层级中的站点
		siteLevel.SiteLevel = siteLevel.SiteLevel[1 : len(siteLevel.SiteLevel)-1]
		fmt.Println(siteLevel.SiteLevel)
	}
	sess.Begin()
	count, err = sess.Table(siteLevel.TableName()).
		Where("id = ?", this.OldId).
		Cols("site_level").
		Update(siteLevel) //设置老层级站点
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return code, count, err
	}
	siteLevel2 := new(schema.SiteLevel)
	has, err = sess.Table(siteLevel2.TableName()).
		Select("site_level").
		Where("id=?", this.NewId).Get(siteLevel2)
	if !has {
		//err = errors.New("没找到新站点的站点层级")
		code = 10138
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	if strings.Contains(siteLevel2.SiteLevel, this.SiteId) == true {
		//err = errors.New("新站点的站点层级已含有传入的站点")
		code = 10139
		return code, count, err
	}
	if siteLevel2.SiteLevel == "" {
		siteLevel2.SiteLevel = this.SiteId
	} else {
		siteLevel2.SiteLevel = siteLevel2.SiteLevel + "," + this.SiteId
	}
	count, err = sess.Table(siteLevel2.TableName()).
		Where("id = ?", this.NewId).
		Cols("site_level").
		Update(siteLevel2) //设置新层级站点
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return code, count, err
	}
	sess.Commit()
	return code, count, err

}

//站点层级列表(以站点搜索)
func (*SiteLevelBean) ListSite(this *input.SiteList) (data back.AllLevelBySiteIdBack, count int64, err error) {
	siteLevel := new(schema.SiteLevel)
	siteLevelPla := new(schema.SiteLevelPlatform)
	platform := new(schema.Platform)
	sess := global.GetXorm().Table(siteLevel.TableName())
	defer sess.Close()
	sess.Where("sales_site_level.state=?", 1)
	sess.Where("sales_site_level_platform.proportion!=?", 0)
	sess.Where("sales_site_level.site_level like?", "%"+this.SiteId+"%")
	data1 := make([]back.ListSiteLevelAll, 0)
	sql := fmt.Sprintf("%s.id = %s.level_id", siteLevel.TableName(), siteLevelPla.TableName())
	sql2 := fmt.Sprintf("%s.platform_id = %s.id", siteLevelPla.TableName(), platform.TableName())
	err = sess.Table(siteLevel.TableName()).
		Select("sales_site_level.id,lid,level_name,site_level,talk,remark,platform_id,platform,proportion").
		Join("LEFT", siteLevelPla.TableName(), sql).
		Join("LEFT", platform.TableName(), sql2).OrderBy("sales_site_level.id").
		Find(&data1)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	var checkid int64 = 0
	siteLevels := make([]back.ListSiteLevel, 0)
	siteLevelPlas := make([]back.SiteLevelProductRate, 0)
	siteLevel2 := new(back.ListSiteLevel)
	siteLevelPla2 := new(back.SiteLevelProductRate)
	for _, d := range data1 {
		if checkid != d.Id { //id不同时，站点层级，组装平台列表，总组装
			if checkid != 0 {
				siteLevel2.Platforms = siteLevelPlas         //平台列表 组装到总组装Platforms参数
				siteLevels = append(siteLevels, *siteLevel2) //总组装
			}
			siteLevelPlas = nil
			checkid = d.Id

			//站点层级
			siteLevel2.Id = d.Id
			siteLevel2.Lid = d.Lid
			siteLevel2.LevelName = d.LevelName
			siteLevel2.SiteLevel = d.SiteLevel
			siteLevel2.Talk = d.Talk
			siteLevel2.Remark = d.Remark
			//组装平台列表
			siteLevelPla2.PlatformId = d.PlatformId
			siteLevelPla2.Platform = d.Platform
			siteLevelPla2.Proportion = d.Proportion
			siteLevelPlas = append(siteLevelPlas, *siteLevelPla2)
		} else { //id相同时，只组装平台列表
			siteLevelPla2.PlatformId = d.PlatformId
			siteLevelPla2.Platform = d.Platform
			siteLevelPla2.Proportion = d.Proportion
			siteLevelPlas = append(siteLevelPlas, *siteLevelPla2)
		}

	}
	siteLevel2.Platforms = siteLevelPlas
	if checkid != 0 {
		siteLevels = append(siteLevels, *siteLevel2)
	}
	data2 := new(back.ListSite)
	if this.SiteId == "" {
		for _, s := range siteLevels {
			data2.Id = s.Id
			data2.LevelName = s.LevelName
			data2.Lid = s.Lid
			data2.Platforms = s.Platforms
			data2.Remark = s.Remark
			data2.Talk = s.Talk
			for _, ss := range strings.Split(s.SiteLevel, ",") {
				data2.SiteId = ss
				data.ListSite = append(data.ListSite, *data2)
				count = count + 1
			}
		}
	} else {
		for _, s := range siteLevels {
			for _, ss := range strings.Split(s.SiteLevel, ",") {
				if this.SiteId == ss {
					data2.SiteId = ss
					data2.Id = s.Id
					data2.LevelName = s.LevelName
					data2.Lid = s.Lid
					data2.Platforms = s.Platforms
					data2.Remark = s.Remark
					data2.Talk = s.Talk
					data.ListSite = append(data.ListSite, *data2)
					count = 1
					return data, count, err
				}
			}
		}
	}
	var strr []back.SiteLevelProductPlatform
	var strro back.SiteLevelProductPlatform
	for _, v := range data.ListSite {
		for _, k := range v.Platforms {
			var il int64
			if len(strr) > 0 {
				for _, g := range strr {
					if k.PlatformId == g.PlatformId {
						il = il + 1
					}
				}
				if il < 1 {
					strro.PlatformId = k.PlatformId
					strro.Platform = k.Platform
					strr = append(strr, strro)
				}
			} else {
				strro.PlatformId = k.PlatformId
				strro.Platform = k.Platform
				strr = append(strr, strro)
			}
		}
	}
	data.PingTai = strr
	var nh back.SiteLevelProductRate
	for _, v := range data.PingTai {
		for m, n := range data.ListSite {
			var q int64
			if len(n.Platforms) > 0 {
				for _, u := range n.Platforms {
					if v.PlatformId == u.PlatformId {
						q = q + 1
					}
				}
			} else {
				nh.PlatformId = v.PlatformId
				nh.Platform = v.Platform
				nh.Proportion = 0
				data.ListSite[m].Platforms = append(data.ListSite[m].Platforms, nh)
			}
			if q < 1 {
				nh.PlatformId = v.PlatformId
				nh.Platform = v.Platform
				nh.Proportion = 0
				data.ListSite[m].Platforms = append(data.ListSite[m].Platforms, nh)
			}
		}
	}
	for _, b := range data.ListSite {
		if len(b.Platforms) > 0 {
			//排序
			for i := 0; i < len(b.Platforms)-1; i++ {
				for j := i + 1; j < len(b.Platforms); j++ {
					if b.Platforms[i].PlatformId > b.Platforms[j].PlatformId {
						b.Platforms[i], b.Platforms[j] = b.Platforms[j], b.Platforms[i]
					}
				}
			}
		}
	}
	return data, count, err
}

//层级下拉框
func (*SiteLevelBean) LevelSiteListDrop() ([]back.ListSiteDropBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.ListSiteDropBack
	sL := new(schema.SiteLevel)
	err := sess.Table(sL.TableName()).Where("state=?", 1).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//得到站点所属层级id
func (m *SiteLevelBean) GetLevelBySite(siteId string, sessArgs ...*xorm.Session) (levelId int64, err error) {
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
	tabName := new(schema.SiteLevel).TableName()
	b, err := sess.Table(tabName).
		Select("id").
		Where("site_level like ?", "%"+siteId+"%").
		Where("state = 1").
		Get(&levelId)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return levelId, err
	}
	// DESCRIPTION:如果没有对应层级,就查询未分层(id为1)的
	if !(b) {
		b, err := sess.Table(tabName).
			Select("id").
			Where("id = ?", 1).
			Where("state = 1").
			Get(&levelId)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			return levelId, err
		}
		if !(b) {
			err = errors.New("There is no hierarchy")
			return levelId, err
		}
	}
	return levelId, err
}

//得到站点占成比,用来计算视讯转出和转入时,站点需要扣除和加上的视讯余额
func (m *SiteLevelBean) GetProportionBySite(siteId string, platformId int64, sessArgs ...*xorm.Session) (proportion float64, err error) {
	// DESCRIPTION:查询层级
	levelId, err := m.GetLevelBySite(siteId, sessArgs...)
	if err != nil {
		return proportion, err
	}
	// DESCRIPTION:查询占成比
	return siteLevelPlatform.GetProportionByLevelIdPlatformId(levelId, platformId, sessArgs...)
}
