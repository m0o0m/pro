package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type AgencyDomainBeen struct{}

//推广域名列表
func (*AgencyDomainBeen) GetList(this *input.AgencyThirdDomainList, listParams *global.ListParams) ([]back.AgencyThirdDomain, int64, error) {
	sess := global.GetXorm().NewSession()
	agencyDomain := new(schema.AgencyThirdDomain)
	agecy := new(schema.Agency)
	var data []back.AgencyThirdDomain
	defer sess.Close()
	//判断并组合where条件
	if this.AgencyId != 0 {
		sess.Where(agencyDomain.TableName()+".agency_id = ?", this.AgencyId)
	}
	if this.Domain != "" {
		sess.Where(agencyDomain.TableName()+".domain = ?", this.Domain)
	}
	sess.Where(agencyDomain.TableName()+".delete_time = ?", 0)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	//获得分页记录
	listParams.Make(sess)
	//重新传入表名和where条件查询记录
	data = make([]back.AgencyThirdDomain, 0)
	agencyCount := new(schema.AgencyCount)
	sql1 := fmt.Sprintf("%s.agency_id = %s.agency_id", agencyDomain.TableName(), agencyCount.TableName())
	sql2 := fmt.Sprintf("%s.id=%s.agency_id", agecy.TableName(), agencyDomain.TableName())
	err := sess.Table(agencyDomain.TableName()).Join("INNER", agencyCount.TableName(), sql1).Join("INNER", agecy.TableName(), sql2).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	//获得符合条件的记录数
	count, err := sess.Table(agencyDomain.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//添加推广域名
func (*AgencyDomainBeen) Add(this *input.AgencyThirdDomain) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agencyDomain := new(schema.AgencyThirdDomain)
	agencyDomain.AgencyId = this.AgencyId
	agencyDomain.CreateTime = time.Now().Unix()
	agencyDomain.Domain = this.Domain
	count, err = sess.Table(agencyDomain.TableName()).Insert(agencyDomain)
	return
}

//修改推广域名
func (*AgencyDomainBeen) UpdateInfo(this *input.AgencyThirdDomainEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agencyDomain := new(schema.AgencyThirdDomain)
	agencyDomain.Id = this.Id
	agencyDomain.AgencyId = this.AgencyId
	agencyDomain.Domain = this.Domain
	if agencyDomain.AgencyId != 0 {
		sess.Cols("agency_id")
	}
	if agencyDomain.Domain != "" {
		sess.Cols("domain")
	}
	count, err = sess.Table(agencyDomain.TableName()).Where("id = ?", agencyDomain.Id).Update(agencyDomain)
	return
}

//删除推广域名
func (*AgencyDomainBeen) Delete(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agencyDomain := new(schema.AgencyThirdDomain)
	agencyDomain.DeleteTime = time.Now().Unix()
	count, err := sess.Table(agencyDomain.TableName()).
		Where("id = ?", id).Cols("delete_time").
		Update(agencyDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看域名是否存在(添加)
func (AgencyDomainBeen) GetDomain(domain string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agencyDomain := new(schema.AgencyThirdDomain)
	has, err := sess.Table(agencyDomain.TableName()).
		Where("domain = ?", domain).Get(agencyDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查看域名是否存在(修改)
func (AgencyDomainBeen) GetDomains(id int64, domain string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agencyDomain := new(schema.AgencyThirdDomain)
	has, err := sess.Table(agencyDomain.TableName()).
		Where("id != ?", id).Where("domain = ?", domain).
		Get(agencyDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}
