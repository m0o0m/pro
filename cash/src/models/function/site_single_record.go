package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strconv"
	"time"
)

type SiteSingleRecordBean struct{}

//掉单列表
func (*SiteSingleRecordBean) GetList(this *input.SiteSingleRecordList, listParams *global.ListParams, times *global.Times) (
	back.SiteSingleRecordsBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var list back.SiteSingleRecordsBack
	siteSingleRecord := new(schema.SiteSingleRecord)
	//判断并组合where条件
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Username != "" {
		sess.Where("username = ?", this.Username)
	}
	if this.Type != 0 {
		sess.Where("type = ?", this.Type)
	}
	ctype, _ := strconv.Atoi(this.Ctype)
	if this.Ctype != "" {
		sess.Where("ctype = ?", ctype)
	}
	vtype, _ := strconv.Atoi(this.Vtype)
	if this.Vtype != "" {
		sess.Where("vtype = ?", vtype)
	}
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	ssrd := make([]back.SiteSingleRecord, 0)
	err := sess.Table(siteSingleRecord.TableName()).Select("money").Find(&ssrd)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	var totalMoney float64
	for k := range ssrd {
		totalMoney += ssrd[k].Money
	}
	//获得分页记录
	listParams.Make(sess)
	times.Make("do_time", sess)
	ssr := make([]back.SiteSingleRecord, 0)
	platform := new(schema.Platform)
	sql := fmt.Sprintf("%s.id = %s.ctype or %s.id = %s.vtype", platform.TableName(), siteSingleRecord.TableName(), platform.TableName(), siteSingleRecord.TableName())
	sess.Select("sales_site_single_record.id,sales_site_single_record.username,sales_site_single_record.ctype," +
		"sales_site_single_record.vtype,sales_site_single_record.money,sales_site_single_record.do_time," +
		"sales_site_single_record.type,sales_site_single_record.remark,sales_platform.platform")
	//重新传入表名和where条件查询记录
	err = sess.Table(siteSingleRecord.TableName()).Join("LEFT", platform.TableName(), sql).Where(conds).Find(&ssr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	var ssrb back.SiteSingleRecordBack
	data := make([]back.SiteSingleRecordBack, 0)
	var subTotalMoney float64 //小计金额
	for k := range ssr {
		subTotalMoney += ssr[k].Money
		if ssr[k].Vtype == 0 {
			ssrb.Vtype = "系统额度"
		} else {
			ssrb.Vtype = ssr[k].Platform
		}
		if ssr[k].Ctype == 0 {
			ssrb.Ctype = "系统额度"
		} else {
			ssrb.Ctype = ssr[k].Platform
		}
		ssrb.DoTime = ssr[k].DoTime
		ssrb.Username = ssr[k].Username
		ssrb.Id = ssr[k].Id
		ssrb.Type = ssr[k].Type
		ssrb.Money = ssr[k].Money
		ssrb.Remark = ssr[k].Remark
		data = append(data, ssrb)
	}
	list.TotalMoney = totalMoney
	list.TotalCount = len(ssrd)
	list.SiteSingleRecordBack = data
	list.SubtotalMoney = subTotalMoney
	return list, err
}

//添加掉单申请
func (*SiteSingleRecordBean) Add(this *input.SiteSingleRecordAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteSingleRecord := new(schema.SiteSingleRecord)
	//掉单时间转换成时间戳
	loc, _ := time.LoadLocation("Local")
	st, _ := time.ParseInLocation("2006-01-02 15:04:05", this.DoTime, loc)
	doTime := st.Unix()
	//给账号表赋值
	if this.Ctype == 0 {
		siteSingleRecord.DoType = 1
	} else {
		siteSingleRecord.DoType = 2
	}
	siteSingleRecord.SiteId = this.SiteId
	siteSingleRecord.SiteIndexId = this.SiteIndexId
	siteSingleRecord.Remark = this.Remark
	siteSingleRecord.Money = this.Money
	siteSingleRecord.Username = this.Username
	siteSingleRecord.AdminUser = this.AdminUser
	siteSingleRecord.DoTime = doTime
	siteSingleRecord.Ctype = this.Ctype
	siteSingleRecord.Vtype = this.Vtype
	siteSingleRecord.Type = 1
	count, err := sess.Table(siteSingleRecord.TableName()).Insert(siteSingleRecord)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看会员账号是否是该操作人所在站点
func (*SiteSingleRecordBean) IsExistMemberAccount(this *input.SiteSingleRecordAdd) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	has, err := sess.Select("id").Where("account = ?", this.Username).
		Where("site_id = ?", this.SiteId).Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//审核掉单申请
func (*SiteSingleRecordBean) Check(this *input.SiteSingleRecordEdit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteSingleRecord := new(schema.SiteSingleRecord)
	siteSingleRecord.Type = this.Type
	siteSingleRecord.UpdateTime = time.Now().Unix()
	siteSingleRecord.UpdateUsername = this.UpdateUsername
	count, err := sess.Table(siteSingleRecord.TableName()).Where("id = ?", this.Id).
		Cols("type,update_time,update_username").Update(siteSingleRecord)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
