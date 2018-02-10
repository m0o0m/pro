package data_merge

import (
	"fmt"
	"framework/render"
	"models/back"
	"models/input"
	"models/schema"
	//"strconv"
	//"strings"
	//"time"
)

type Bank struct {
	Key
	CdnUrl string //
	//SiteName        string
	//LogoUrl         string
	//Datetime        string
	//TradeDate       string
	//OrderNum        string
	UrlLink         string              //客服链接
	PaidType        map[string]int      //支付类型
	OnlineIncomeSet []back.GetPayeeInfo //收款账号列表
	SiteIncomeBank  []back.SiteBank     //入款银行
	PageType        int                 //页面类型 1存款 2取款 3额度转换 4客服
}

func (m *Bank) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//siteName, err := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	//if err != nil {
	//	return nil, err
	//}

	//list, err := webLogo.GetWebLogoList(&input.LogoInfoList{siteId, siteIndexId})
	//if err != nil {
	//	return nil, err
	//}
	//if len(list) > 0 {
	//	m.LogoUrl = list[0].LogoUrl
	//}

	//客服链接
	siteInfo, _, err := siteInfoBean.GetSingleInfo(&schema.SiteInfo{siteId, siteIndexId, "", "", "", "", "", ""})
	if err != nil {
		return nil, err
	}

	//支付类型
	infoList, err := paidTypeBean.GetPaidTypeData()
	if err != nil {
		return nil, err
	}
	m.PaidType = make(map[string]int)
	for k, _ := range infoList {
		if infoList[k].TypeStatus == 1 {
			m.PaidType[infoList[k].PaidTypeName] = infoList[k].Id
		}
	}

	//收款人信息
	bankInfo, err := memberCompanyIncomeBean.GetSetAccountInfo(&input.SiteId{siteId, siteIndexId})
	if err != nil {
		return nil, err
	}
	m.OnlineIncomeSet = bankInfo
	fmt.Println(bankInfo)
	//入款银行
	siteBank, err := bankCardBean.GetAllIncomeBank(siteId, siteIndexId)
	fmt.Println(siteBank)
	if err != nil {
		return nil, err
	}
	m.SiteIncomeBank = siteBank
	m.UrlLink = siteInfo.UrlLink
	m.PageType = 1

	//m.Datetime = time.Now().String()[:19]
	//m.SiteName = siteName
	//m.OrderNum = strings.Replace(time.Now().String()[:10], "-", "", 2) + siteId + siteIndexId + strconv.FormatInt(time.Now().Unix(), 10)
	//m.TradeDate = time.Now().String()[:10]
	return m, nil
}

func (m *Bank) GetPage() []string {
	return []string{WAP_BANK, WAP_MEM_FOOTER}
}
