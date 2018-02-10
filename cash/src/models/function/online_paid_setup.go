package function

import (
	"time"

	"global"
	"math/rand"
	"models/back"
	"models/input"
	"models/schema"
)

type OnlinePaidSetupBean struct{}

//新增加线上第三方支付设定
func (*OnlinePaidSetupBean) AddNew(newPaid *input.NewThirdPayOnline, newData *input.AddSetupParse) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	paidSet := new(schema.OnlinePaidSetup)
	paidSet.MerchatId = newPaid.MerchatId                 //商户id
	paidSet.SiteId = newPaid.SiteId                       //代理线
	paidSet.SiteIndexId = newPaid.SiteIndexId             //子代理
	paidSet.FitforLevel = newPaid.LevelId                 //适用层级
	paidSet.PaidLimit = newPaid.PaidLimit                 //支付限额
	paidSet.BackAddress = newPaid.BackAddress             //返回地址
	paidSet.PrivateKey = newPaid.PrviateKey               //私钥
	paidSet.PublicKey = newPaid.PublicKey                 //公钥
	paidSet.PaidPlatform = newPaid.PaidPlatform           //支付平台
	paidSet.PaidType = newPaid.PaidType                   //支付类型
	paidSet.IsJumpApp = newPaid.IsApp                     //是否跳转app
	paidSet.MerUrl = newPaid.MerUrl                       //自填写支付网关
	paidSet.PaidCode = newPaid.PaidCode                   //支付code
	paidSet.Status = newPaid.Staus                        //状态
	paidSet.PaidDomain = newPaid.PaidDomain               //支付域名
	paidSet.SuitableEquipment = newPaid.SuitableEquipment //适用设备
	paidSet.CreateTime = time.Now().Unix()                //创建时间
	paidSet.DelTime = 0                                   //删除时间
	count, err := sess.Insert(paidSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//线上第三方支付设定列表
func (*OnlinePaidSetupBean) PaidList(selectBy *input.ThirdList, listParam *global.ListParams) (
	[]back.OnlinePaidSetupList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var infolist []back.OnlinePaidSetupList
	paidlist := new(schema.OnlinePaidSetup)
	if selectBy.Status == 1 || selectBy.Status == 2 { //状态
		sess.Where("status=?", selectBy.Status)
	}
	if selectBy.ThirdId != 0 { //支付平台
		sess.Where("paid_platform=?", selectBy.ThirdId)
	}
	if selectBy.SiteIndexId != "" {
		sess.Where("site_index_id=?", selectBy.SiteIndexId)
	}
	if selectBy.SiteId != "" {
		sess.Where("site_id=?", selectBy.SiteId)
	}
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	//获得分页记录
	listParam.Make(sess)
	err := sess.Table(paidlist.TableName()).
		Select("sales_online_paid_setup.id,sales_online_paid_setup.status,sales_online_paid_setup.fitfor_level,"+
			"sales_online_paid_setup.paid_type,sales_online_paid_setup.remark,"+
			"sales_online_paid_setup.paid_limit,sales_online_paid_setup.merchat_id,"+
			"sales_online_income_third.title,sales_paid_type.paid_type_name as type_name").
		Join("left", "sales_paid_type", "sales_paid_type.id=sales_online_paid_setup.paid_type").
		Join("left", "sales_online_income_third", "sales_online_income_third.id=sales_online_paid_setup.paid_platform").
		Find(&infolist)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, 0, err
	}
	//获得符合条件的记录数
	count, err := sess.Table(paidlist.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, count, err
	}
	return infolist, count, err
}

//随机选择一条符合条件的第三方
func (*OnlinePaidSetupBean) GetOnePaid(selectBy *input.GetOnePaid) (oneThird back.GetOneThird, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	paidList := new(schema.OnlinePaidSetup)
	paidLists := make([]back.GetOneThird, 0)
	sess.Where("status=?", 1) //状态(1.启用2.停用)
	sess.Where("site_index_id=?", selectBy.SiteIndexId)
	sess.Where("site_id=?", selectBy.SiteId)
	sess.Where("paid_type=?", selectBy.PaidType)
	sess.Where("delete_time=?", 0)
	sess.Where("concat(',',fitfor_level,',') like ?", "%,"+selectBy.LevelId+",%")
	conds := sess.Conds()
	err = sess.Table(paidList.TableName()).
		Select("sales_online_paid_setup.id,sales_online_paid_setup.paid_platform,sales_online_paid_setup.status,sales_online_paid_setup.fitfor_level,"+
			"sales_online_paid_setup.paid_type,sales_online_paid_setup.remark,"+
			"sales_online_paid_setup.paid_limit,sales_online_paid_setup.merchat_id,"+
			"sales_online_income_third.title,sales_paid_type.paid_type_name as type_name,sales_online_paid_setup.paid_code").
		Join("left", "sales_paid_type", "sales_paid_type.id=sales_online_paid_setup.paid_type").
		Join("left", "sales_online_income_third", "sales_online_income_third.id=sales_online_paid_setup.paid_platform").Find(&paidLists)
	//获得符合条件的记录数
	count, err = sess.Table(paidList.TableName()).Where(conds).Count()
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //更新随机种子
	if len(paidLists) > 0 {
		oneThird = paidLists[r.Intn(len(paidLists))] //随机获取一个第三方
	}
	return
}

//根据id选择对应的一条第三方
func (*OnlinePaidSetupBean) GetOnePaidById(selectBy *input.GetOnePaidById) (oneThird *back.GetOneThird, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	paidList := new(schema.OnlinePaidSetup)
	paidType := new(schema.PaidType)
	third := new(schema.OnlineIncomeThird)
	oneThird = new(back.GetOneThird)
	sess.Where(paidList.TableName()+".id=?", selectBy.Id)
	sess.Where("status=?", 1) //状态(1.启用2.停用)
	sess.Where("site_index_id=?", selectBy.SiteIndexId)
	sess.Where("site_id=?", selectBy.SiteId)
	sess.Where("paid_type=?", selectBy.PaidType)
	sess.Where("delete_time=?", 0)
	sess.Where("concat(',',fitfor_level,',') like ?", "%,"+selectBy.LevelId+",%")
	has, err = sess.Table(paidList.TableName()).
		Select(paidList.TableName()+".id,"+
			paidList.TableName()+".paid_platform,"+
			paidList.TableName()+".status,"+
			paidList.TableName()+".fitfor_level,"+
			paidList.TableName()+".paid_type,"+
			paidList.TableName()+".remark,"+
			paidList.TableName()+".paid_limit,"+
			paidList.TableName()+".merchat_id,"+
			third.TableName()+".title,"+
			paidType.TableName()+".paid_type_name as type_name,"+
			paidList.TableName()+".paid_code").
		Join("left", paidType.TableName(), paidType.TableName()+".id="+paidList.TableName()+".paid_type").
		Join("left", third.TableName(), third.TableName()+".id="+paidList.TableName()+".paid_platform").Get(oneThird)
	return
}

//线上第三方支付设定修改
func (*OnlinePaidSetupBean) ChangeSet(newOnlinePaidSet *input.ChangeThisPaidSetup) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newPaidSet := new(schema.OnlinePaidSetup)
	newPaidSet.Status = newOnlinePaidSet.Status                       //状态
	newPaidSet.FitforLevel = newOnlinePaidSet.FitforLevel             //适用层级
	newPaidSet.PaidLimit = newOnlinePaidSet.PaidLimit                 //支付限额
	newPaidSet.PaidDomain = newOnlinePaidSet.PaidDomain               //支付域名
	newPaidSet.BackAddress = newOnlinePaidSet.BackAddress             //回调地址
	newPaidSet.PrivateKey = newOnlinePaidSet.PrivateKey               //私钥
	newPaidSet.PublicKey = newOnlinePaidSet.PublicKey                 //公钥
	newPaidSet.SuitableEquipment = newOnlinePaidSet.SuitableEquipment //适用设备
	newPaidSet.Remark = newOnlinePaidSet.Remark                       //备注
	newPaidSet.PaidCode = newOnlinePaidSet.PaidCode                   //支付编码
	count, err := sess.Where("id=?", newOnlinePaidSet.Id).Cols("paid_code,remark,suitable_equipment,public_key,private_key,back_address,status,fitfor_level,paid_limit,paid_domain").Update(newPaidSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//线上第三方支付设定停用
func (*OnlinePaidSetupBean) StopThirdPaid(stopPaid *input.StopThisPaidSetup) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newPaid := new(schema.OnlinePaidSetup)
	newPaid.Status = stopPaid.Status
	count, err := sess.Where("id=?", stopPaid.Id).
		Where("delete_time=?", 0).Cols("status").Update(newPaid)
	if err != nil {
		global.GlobalLogger.Error("error%s", err.Error())
		return count, err
	}
	return count, err
}

//线上第三方支付设定删除
func (*OnlinePaidSetupBean) DelThirdPaid(delPaid *input.DelThisPaidSetup) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newPaid := new(schema.OnlinePaidSetup)
	newPaid.DelTime = time.Now().Unix()
	sess.Where("id=?", delPaid.Id)
	count, err := sess.Cols("delete_time").Update(newPaid)
	if err != nil {
		global.GlobalLogger.Error("error%s", err.Error())
		return count, err
	}
	return count, err
}

//根据id获取某条支付设定记录
func (*OnlinePaidSetupBean) GetOnePaidSet(id int) (schema.OnlinePaidSetup, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.OnlinePaidSetup
	sess.Where("id=?", id)
	sess.Where("delete_time=?", 0)
	flag, err := sess.Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, err
}

//当拉取线上的数据的时候有选择性的刷新本地数据库
func (*OnlinePaidSetupBean) UpdateLocalData(newData *input.OnlinePaidSetParse) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	refreshData := new(schema.OnlinePaidSetup)
	refreshData.PrivateKey = newData.PrivateKey
	refreshData.PublicKey = newData.PublicKey
	refreshData.BackAddress = newData.NotifyUrl
	refreshData.FitforLevel = newData.LevelId
	refreshData.MerUrl = newData.MerUrl
	refreshData.PaidCode = newData.PayCode
	count, err = sess.Where("id=?", newData.MerId).Update(refreshData)
	return
}

//根据支付类型和商户号获取支付设定详情
func (*OnlinePaidSetupBean) GetInfoBy(newInput *input.NewThirdPayOnline) (
	schema.OnlinePaidSetup, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.OnlinePaidSetup
	sess.Where("merchat_id=?", newInput.MerchatId)
	sess.Where("paid_type=?", newInput.PaidType)
	sess.Where("delete_time=?", 0)
	flag, err := sess.Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, err
}
