package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type PressMoneyBean struct{}

//催款账单列表
func (*PressMoneyBean) GetPressMoneyList(this *input.PressMoneyList) (data []back.SiteMoneyPress, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	press := new(schema.SiteMoneyPress)
	if this.Status != 0 {
		sess.Where("state = ?", this.Status)
	}
	if len(this.SiteId) > 0 {
		sess.Where("site_id like?", "%"+this.SiteId+"%")
	}
	if len(this.StartTime) != 0 {
		StartTime, _ := global.FormatTime2Timestamp2(this.StartTime)
		sess.Where("add_date >= ?", StartTime)
	}
	if len(this.EndTime) != 0 {
		EndTime, _ := global.FormatTime2Timestamp2(this.EndTime)
		sess.Where("add_date <= ?", EndTime)
	}
	//获得符合条件的记录数
	err = sess.Table(press.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//添加催款账单
func (*PressMoneyBean) Add(this []schema.SiteMoneyPress) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteMoneyPress)
	count, err = sess.Table(admin.TableName()).InsertMulti(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//账单详情
func (*PressMoneyBean) GetInfo(this *input.PressMoneyOne) (data back.SiteMoneyPress, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	press := new(schema.SiteMoneyPress)
	has, err = sess.Table(press.TableName()).Where("id = ?", this.Id).Get(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//修改催款账单
func (*PressMoneyBean) UpdateInfo(this *input.PressMoneyEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteMoneyPress)
	admin.SiteId = this.SiteId
	admin.PayName = this.PayName
	admin.PayAddress = this.PayAddress
	admin.PayCard = this.PayCard
	admin.Bank = this.Bank
	admin.Remark = this.Remark
	count, err = sess.Where("id = ?", this.Id).Cols("site_id,pay_name,pay_address,pay_card,bank,remark").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改账单状态
func (*PressMoneyBean) UpdateStatus(this *input.PressMoneyStatus) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.SiteMoneyPress)
	admin.Status = 2
	admin.UpdateDate = time.Now().Unix()
	count, err = sess.Where("id = ?", this.Id).Cols("status").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//site
func (*PressMoneyBean) PressMoneySiteByDrop() ([]back.SiteMoneyPressSiteBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	var data []back.SiteMoneyPressSiteBack
	err := sess.Table(site.TableName()).Where("is_default=?", 1).
		Where("status=?", 1).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//站点获取缴款账单
func (*PressMoneyBean) PressMoneySiteInfo(siteId string) (back.SiteMoneyPress, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data back.SiteMoneyPress
	site := new(schema.SiteMoneyPress)
	sess.Desc("id")
	sess.Limit(1)
	has, err := sess.Table(site.TableName()).Where("site_id = ?", siteId).Get(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//站点根据id和siteId获取缴款账单是否存在
func (*PressMoneyBean) PressIdInfo(siteId string, Id int64) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.SiteMoneyPress)
	count, err = sess.Table(site.TableName()).Where("site_id = ? AND id = ?", siteId, Id).Count()
	return
}

//预缴款
func (*PressMoneyBean) AddPrePayment(this *input.PrePayment) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sitePayRecord := new(schema.SitePayRecord)

	sitePayRecord.SiteId = this.SiteId
	sitePayRecord.SiteIndexId = this.SiteIndexId
	sitePayRecord.State = 1
	sitePayRecord.OrderNum = this.OrderNum
	sitePayRecord.AdminUser = this.AdminUser
	sitePayRecord.Money = float64(this.Money)
	sitePayRecord.Type = 2
	sitePayRecord.Bank = this.Bank
	sitePayRecord.DoTime = global.GetCurrentTime()
	sitePayRecord.Remark = fmt.Sprintf("%s站点预存款%d，转入账号%s", this.SiteId, this.Money, this.PayCard)
	count, err = sess.Table(sitePayRecord.TableName()).Insert(sitePayRecord)
	return
}
