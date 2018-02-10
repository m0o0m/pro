package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type OnlineIncomeThirdBean struct{}

//判断数据库是否存在数据
func (*OnlineIncomeThirdBean) IsExistData() (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineThird := new(schema.OnlineIncomeThird)
	flag, err := sess.Exist(onlineThird)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return flag, err
	}
	return flag, err
}

//获取所有的数据
func (*OnlineIncomeThirdBean) GetAllData() ([]schema.OnlineIncomeThird, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var infolist []schema.OnlineIncomeThird
	onlineThird := new(schema.OnlineIncomeThird)
	err := sess.Table(onlineThird.TableName()).Find(&infolist)
	return infolist, err
}

//获取站点管理--三方操作返回列表
func (*OnlineIncomeThirdBean) GetThirdNeedData(listparams *global.ListParams) ([]back.OnlineThird, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var backData []back.OnlineThird
	listparams.Make(sess)
	onlineThird := new(schema.OnlineIncomeThird)
	sess.Table(onlineThird.TableName())
	sess.Where("delete_time=?", 0)
	err := sess.Find(&backData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, 0, err
	}
	count, err := sess.Table(onlineThird.TableName()).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, 0, err
	}
	return backData, count, err
}

//删除所有
func (*OnlineIncomeThirdBean) DelAllData(record []schema.OnlineIncomeThird) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err := sess.Begin()
	onlineThird := new(schema.OnlineIncomeThird)
	count, err := sess.Where("id>?", 0).Delete(onlineThird)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	count, err = sess.Table(onlineThird.TableName()).Insert(&record)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	return count, err
}

//批量插入数据
func (*OnlineIncomeThirdBean) BatchInsertData(insertData []schema.OnlineIncomeThird) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineThird := new(schema.OnlineIncomeThird)
	count, err := sess.Table(onlineThird.TableName()).Insert(&insertData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取第三方支付平台id和名称
func (*OnlineIncomeThirdBean) GetThirdNameAndId() ([]back.BackSelectThird, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.BackSelectThird
	newOnlineIncome := new(schema.OnlineIncomeThird)
	err := sess.Table(newOnlineIncome.TableName()).Select("id,title").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//修改本地第三方状态
func (*OnlineIncomeThirdBean) ChangeThirdStatus(inputStatus *input.UpStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineThird := new(schema.OnlineIncomeThird)
	sess.Table(onlineThird.TableName())
	sess.Where("id=?", inputStatus.Id)
	onlineThird.PayStatus = inputStatus.Status
	count, err := sess.Cols("pay_status").Update(onlineThird)
	return count, err
}

//修改网银信息
func (*OnlineIncomeThirdBean) UpdateInfo(this *input.UpdateThird) (int64, error) {

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineThird := new(schema.OnlineIncomeThird)
	sess.Table(onlineThird.TableName())
	sess.Where("delete_time=?", 0)
	sess.Where("id=?", this.Id)
	onlineThird.AliPayUrl = this.PayUrl
	onlineThird.BankUrl = this.BankUrl
	onlineThird.OutState = this.IsOpenOut
	onlineThird.DepositState = this.IsOpenOut
	onlineThird.PayName = this.ThirdName
	onlineThird.IpState = this.IsIpLimit
	onlineThird.PayCode = this.Code
	onlineThird.PayStatus = this.Status
	sess.Cols("pay_status", "pay_code", "ip_state")
	sess.Cols("title", "deposit_state", "out_state")
	sess.Cols("bank_url", "ali_pay_url")
	count, err := sess.Update(onlineThird)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return count, err
}

//根据Id获取网银信息
func (*OnlineIncomeThirdBean) GetInfoById(id int64) (*schema.OnlineIncomeThird, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineThird := new(schema.OnlineIncomeThird)
	sess.Table(onlineThird.TableName())
	sess.Where("delete_time=?", 0)
	sess.Where("id=?", id)
	have, err := sess.Get(onlineThird)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return onlineThird, have, err
}

//删除网银信息
func (*OnlineIncomeThirdBean) Del(this *input.DelThird) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineThird := new(schema.OnlineIncomeThird)
	sess.Table(onlineThird.TableName())
	sess.Where("delete_time=?", 0)
	sess.Where("id=?", this.Id)
	onlineThird.DeleteTime = time.Now().Unix()
	count, err := sess.Cols("delete_time").Update(onlineThird)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return count, err
}

//添加网银信息
func (*OnlineIncomeThirdBean) Add(this *input.AddThird) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineThird := new(schema.OnlineIncomeThird)
	sess.Table(onlineThird.TableName())
	onlineThird.AliPayUrl = this.PayUrl
	onlineThird.BankUrl = this.BankUrl
	onlineThird.DepositState = this.IsOpenIn
	onlineThird.IpState = this.IsIpLimit
	onlineThird.OutState = this.IsOpenOut
	onlineThird.PayCode = this.Code
	onlineThird.PayModels = this.ModelName
	onlineThird.PayName = this.ThirdName
	onlineThird.PayStatus = this.Status
	count, err := sess.Insert(onlineThird)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return count, err
}
