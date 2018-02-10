package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type SitePromotionConfigBean struct{}

//自助申请优惠列表
func (*SitePromotionConfigBean) GetSitePromotionConfig(this *input.SitePromotionConfig) (
	[]back.SitePromotionConfig, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.SitePromotionConfig
	pre := new(schema.SitePromotionConfig)
	si := new(schema.Site)
	if this.SiteIndexId != "" {
		sess.Where(pre.TableName()+".site_index_id = ?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where(pre.TableName()+".site_id = ?", this.SiteId)
	}
	//获得符合条件的记录数
	err := sess.Table(pre.TableName()).
		Join("LEFT", si.TableName(), pre.TableName()+".site_id="+si.TableName()+".id AND "+
			pre.TableName()+".site_index_id="+si.TableName()+".index_id").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//添加自助申请配置
func (*SitePromotionConfigBean) Add(this schema.SitePromotionConfig) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SitePromotionConfig)
	count, err := sess.Table(admin.TableName()).InsertOne(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改自助优惠申请状态
func (*SitePromotionConfigBean) UpdateStatus(this *input.SitePromotionConfigStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SitePromotionConfig)
	if this.Status == 1 {
		admin.Status = 2
	} else {
		admin.Status = 1
	}
	admin.Updatetime = time.Now().Unix()
	count, err := sess.Table(admin.TableName()).
		Where("site_id=?", this.SiteId).Where("id = ?", this.Id).
		Cols("status,update_time").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询活动列表
func (*SitePromotionConfigBean) GetSiteProConfig(this *input.SitePromotionConfig) (data []back.SitePromotionConfig, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	pre := new(schema.SitePromotionConfig)
	if this.SiteIndexId != "" {
		sess.Where(pre.TableName()+".site_index_id = ?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where(pre.TableName()+".site_id = ?", this.SiteId)
	}
	//获得符合条件的记录数
	err = sess.Table(pre.TableName()).
		Select("id,site_id,site_index_id,pro_title,pro_content,createtime,status").
		Find(&data)
	return
}
