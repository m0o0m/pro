package function

import (
	"fmt"
	"global"
	"models/back"
	"models/schema"
)

type PaidTypeBean struct{}

//增加
func (*PaidTypeBean) AddNew(newData []schema.PaidType) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	paidType := new(schema.PaidType)
	count, err := sess.Table(paidType.TableName()).Insert(&newData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//整张表有不有数据
func (*PaidTypeBean) ExistDataPaid() (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	paidType := new(schema.PaidType)
	flag, err := sess.Table(paidType.TableName()).Exist(paidType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return flag, err
	}
	return flag, err
}

//删除表数据
func (*PaidTypeBean) DelAllData(record []schema.PaidType) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	paidType := new(schema.PaidType)
	count, err := sess.Where("id>?", 0).Delete(paidType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	count, err = sess.Table(paidType.TableName()).Insert(&record)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	return count, err
}

//获取所有的支付类型
func (*PaidTypeBean) GetTbaleData() ([]back.PaidTypeBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.PaidTypeBack
	newTab := new(schema.PaidType)
	err := sess.Table(newTab.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取所有可用的支付类型
func (*PaidTypeBean) GetPaidTypeData() (data []back.PaidTypeAndStatusBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newTab := new(schema.PaidType)
	err = sess.Table(newTab.TableName()).Find(&data)
	return
}

//获取对应站点所有可用的支付类型
func (*PaidTypeBean) GetSitePaidTypeData(siteId, siteIndexId, levelId string) (data []back.PaidTypeAndStatusBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newTab := new(schema.PaidType)
	newTabs := make([]schema.PaidType, 0)
	setup := new(schema.OnlinePaidSetup)
	setups := make([]*schema.OnlinePaidSetup, 0)
	sess.Where("type_status=?", 1)
	err = sess.Table(newTab.TableName()).Cols("id,paid_type_name,type_status").Find(&newTabs)
	if err != nil {
		fmt.Println(err)
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if len(newTabs) == 0 {
		global.GlobalLogger.Error("error:%s", "未找到可用的支付类型")
		return
	}
	ids := make([]int, 0)
	newTabMap := make(map[int]*schema.PaidType)
	for k, v := range newTabs {
		ids = append(ids, v.Id)
		newTabMap[v.Id] = &newTabs[k]
	}
	sess.In("paid_type", ids)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("concat(',',fitfor_level,',') like ?", "%,"+levelId+",%")
	err = sess.Table(setup.TableName()).Cols("id,paid_type,sort").OrderBy("paid_type,sort desc").Find(&setups)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if len(setups) == 0 {
		global.GlobalLogger.Error("error:%s", "未找到会员对应层级的支付设定")
		return
	}
	for _, v := range setups {
		data = append(data, back.PaidTypeAndStatusBack{
			newTabMap[v.PaidType].Id,
			newTabMap[v.PaidType].PaidTypeName,
			int8(newTabMap[v.PaidType].TypeStatus)})
	}
	return
}

//获取线上存款支付类型+支付设定
func (*PaidTypeBean) GetIncomeData(siteId, siteIndexId, levelId string) (data []back.PaidTypeAndPaySetBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newTab := new(schema.PaidType)
	newTabs := make([]schema.PaidType, 0)
	setup := new(schema.OnlinePaidSetup)
	setups := make([]*schema.OnlinePaidSetup, 0)
	sess.Where("type_status=?", 1)
	err = sess.Table(newTab.TableName()).Cols("id,paid_type_name,type_status").Find(&newTabs)
	if err != nil {
		fmt.Println(err)
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if len(newTabs) == 0 {
		global.GlobalLogger.Error("error:%s", "未找到可用的支付类型")
		return
	}
	ids := make([]int, 0)
	newTabMap := make(map[int]*schema.PaidType)
	for k, v := range newTabs {
		ids = append(ids, v.Id)
		newTabMap[v.Id] = &newTabs[k]
	}
	sess.In("paid_type", ids)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("concat(',',fitfor_level,',') like ?", "%,"+levelId+",%")
	err = sess.Table(setup.TableName()).Cols("id,paid_type,sort").OrderBy("paid_type,sort desc").Find(&setups)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if len(setups) == 0 {
		global.GlobalLogger.Error("error:%s", "未找到会员对应层级的支付设定")
		return
	}
	for _, v := range setups {
		data = append(data, back.PaidTypeAndPaySetBack{
			newTabMap[v.PaidType].Id,
			newTabMap[v.PaidType].PaidTypeName,
			int8(newTabMap[v.PaidType].TypeStatus),
			int64(v.Id),
			v.Sort})
	}
	return
}

//快捷支付 获取线上存款支付类型+支付设定
func (*PaidTypeBean) GetFastIncomeData(siteId, siteIndexId, levelId string) (data []back.FastIncome, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	setup := new(schema.OnlinePaidSetup)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("concat(',',fitfor_level,',') like ?", "%,"+levelId+",%")
	err = sess.Table(setup.TableName()).
		Select("paid_type,id as set_id,sort").OrderBy("paid_type").Find(&data)
	return
}

//组装 公司入款支付类型+收款银行设定 与 线上存款支付类型+支付设定
func (*PaidTypeBean) MergeCompanyAndOnlineData(online []back.PaidTypeAndPaySetBack, company []back.GetPayeeInfo) (mergeData map[int]*back.IncomeData, err error) {
	//支付类型的中文名
	zhName := []string{"网银", "微信支付", "支付宝", "QQ钱包", "财付通", "银联扫码", "京东钱包", "百度钱包", "快捷支付", "点卡支付"}
	//封装数据
	mergeData = make(map[int]*back.IncomeData)
	for _, v := range online {
		//onlineMap[v.PaidTypeName]=&online[k]
		if _, ok := mergeData[v.Id]; !ok {
			mergeData[v.Id] = &back.IncomeData{}
		}
		mergeData[v.Id] = &back.IncomeData{
			v.Id,
			zhName[v.Id-1],
			append(mergeData[v.Id].OnlineIncome, back.OnlineIncomeData{
				int64(v.Id),
				v.Sort}),
			nil}
	}
	for _, v := range company {
		if _, ok := mergeData[v.PayTypeId]; !ok {
			mergeData[v.PayTypeId] = &back.IncomeData{}
		}
		mergeData[v.PayTypeId] = &back.IncomeData{
			v.PayTypeId,
			zhName[v.PayTypeId-1],
			mergeData[v.PayTypeId].OnlineIncome,
			append(mergeData[v.PayTypeId].CompanyIncome, back.GetPayeeInfo{
				int64(v.Id),
				v.PaidTypeName,
				v.Title,
				v.Account,
				v.OpenBank,
				v.Payee,
				v.StopBalance,
				v.BankId,
				v.QrCode,
				v.PayTypeId,
				v.Sort})}
	}
	return
}

/*
//获取网银在线的银行
func (*PaidTypeBean) GetOnlineIncomeBank(this []back.OnlineIncomeData) (data []back.GetOnlineIncomeBank, err error) {
	//这里添加post的body内容
	apiClients, has, err := apiClientsBean.GetOneApiClients("zzz") //member.SiteId todo 真实上线时，要对应上站点
	data := make(url.Values)
	data["clientUserId"] = []string{strconv.FormatInt(this.ClientUserId, 10)}
	data["clientName"] = []string{this.ClientName}
	data["clientSecret"] = []string{this.ClientSecret}
	data["order"] = []string{this.Order}
	data["agentId"] = []string{this.SiteId}
	data["agentNum"] = []string{this.SiteIndexId}
	fmt.Println(this.Order)
	res, err := http.PostForm("http://olpay.pk1358.com/api/v1/token", data) //获取token
	if err != nil {
		return err
	}
	fmt.Println(res)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	f := struct {
		Status bool   `json:"status"`
		Token  string `json:"token"`
	}{}
	fmt.Println(string(b))
	err = json.Unmarshal(b, &f)
	if err != nil {
		return err
	}
	if !f.Status {
		return errors.New("获取token失败")
	}
	fmt.Println("-------------------------" + f.Token)

	return
}
*/
