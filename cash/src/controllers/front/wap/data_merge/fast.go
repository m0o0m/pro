package data_merge

import (
	"framework/render"
	"models/back"
	"models/input"
)

type Fast struct {
	Key
	CdnUrl          string              //
	PaidType        map[string]int      //支付类型
	OnlineIncomeSet []back.GetPayeeInfo //收款账号列表
	SiteIncomeBank  []back.SiteBank     //入款银行
	UrlLink         string
	PageType        int    //页面类型 1存款 2取款 3额度转换 4客服
	SiteName        string //站点名称
}

func (m *Fast) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.PageType = 1
	m.CdnUrl = render.CdnUrl
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
	//入款银行
	siteBank, err := bankCardBean.GetAllIncomeBank(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	m.SiteIncomeBank = siteBank
	//站点客服链接
	siteinfo, has, err := SiteInfoBean.GetSingleSiteInfo(siteId, siteIndexId)
	if err != nil {
		m.UrlLink = ""
	}
	if has == false {
		m.UrlLink = ""
	} else {
		m.UrlLink = siteinfo.UrlLink
	}
	//查询站点名称
	m.SiteName, err = siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		m.SiteName = ""
	}
	return m, nil
}

func (m *Fast) GetPage() []string {
	return []string{WAP_FAST, WAP_MEM_FOOTER}
}
