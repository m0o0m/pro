package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

//站点管理-广告管理
type SiteADBean struct{}

//广告列表
func (*SiteADBean) AdminAdvertList(this *input.AdminAdvertList, listParams *global.ListParams) ([]back.AdminAdvert, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	aa := make([]back.AdminAdvert, 0)
	sa := new(schema.SiteAdv)
	if this.Title != "" {
		sess.Where("title = ?", this.Title)
	}
	if this.State != 0 {
		sess.Where("state = ?", this.State)
	}
	//广告管理的type为2
	sess.Where("type = 2").Where("delete_time = 0")
	conds := sess.Conds()
	listParams.Make(sess)
	err := sess.Table(sa.TableName()).Find(&aa)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return aa, 0, err
	}
	//根据sort字段排序
	for i := 0; i < len(aa)-1; i++ {
		for j := i + 1; j < len(aa); j++ {
			if aa[i].Sort > aa[j].Sort {
				aa[j], aa[i] = aa[i], aa[j]
			}
		}
	}
	count, err := sess.Table(sa.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return aa, count, err
	}
	return aa, count, err
}

//添加广告
func (*SiteADBean) AdminAdvertAdd(this *input.AdminAdvertAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sa := new(schema.SiteAdv)
	sa.Sort = this.Sort
	sa.Title = this.Title
	sa.BeforeUrl = this.BeforeUrl
	sa.AfterUrl = this.AfterUrl
	sa.State = this.State
	sa.Content = this.Content
	sa.Remark = this.Remark
	sa.SiteText = this.SiteText
	sa.IsLink = this.IsLink
	sa.AddTime = global.GetCurrentTime()
	sa.Type = 2
	count, err := sess.InsertOne(sa)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//编辑广告
func (*SiteADBean) AdminAdvertPut(this *input.AdminAdvertPut) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sa := new(schema.SiteAdv)
	sa.Sort = this.Sort
	sa.Title = this.Title
	sa.BeforeUrl = this.BeforeUrl
	sa.AfterUrl = this.AfterUrl
	sa.State = this.State
	sa.Content = this.Content
	sa.Remark = this.Remark
	sa.SiteText = this.SiteText
	sa.IsLink = this.IsLink
	count, err := sess.Cols("sort,title,before_url,after_url,state,content,remark,stateText,is_link").
		Where("id = ?", this.Id).Update(sa)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//广告详情
func (*SiteADBean) AdminAdvertInfo(id int64) (*back.AdminAdvertInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sa := new(schema.SiteAdv)
	aai := new(back.AdminAdvertInfo)
	sess.Table(sa.TableName())
	sess.Where("delete_time=?", 0)
	has, err := sess.Where("id = ?", id).Get(aai)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return aai, has, err
	}
	return aai, has, err
}

//修改广告状态
func (*SiteADBean) AdminAdvertState(this *input.AdminAdvertState) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sa := new(schema.SiteAdv)
	sa.State = this.State
	count, err := sess.Cols("state").Where("id = ?", this.Id).Update(sa)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改广告排序
func (*SiteADBean) AdminAdvertSort(this *input.AdminAdvertSort) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sa := new(schema.SiteAdv)
	sa.Sort = this.Sort
	count, err := sess.Cols("sort").Where("id = ?", this.Id).Update(sa)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
