package data_merge

import (
	"framework/render"
	"models/input"
	"models/schema"
	"strconv"
	"strings"
	"time"
)

type QuickPay struct {
	Key
	CdnUrl    string //cdn
	SiteName  string
	LogoUrl   string
	UrlLink   string
	Datetime  string
	PaidType  map[string]int8
	TradeDate string
	OrderNum  string
}

//得到首页数据
func (m *QuickPay) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//初始化头部尾部数据
	//查询站点
	siteName, err := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}

	list, err := webLogo.GetWebLogoList(&input.LogoInfoList{siteId, siteIndexId})
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		m.LogoUrl = list[0].LogoUrl
	}

	siteInfo, _, err := siteInfoBean.GetSingleInfo(&schema.SiteInfo{siteId, siteIndexId, "", "", "", "", "", ""})
	if err != nil {
		return nil, err
	}

	infoList, err := paidTypeBean.GetPaidTypeData()
	if err != nil {
		return nil, err
	}
	m.PaidType = make(map[string]int8)
	for k, _ := range infoList {
		m.PaidType[infoList[k].PaidTypeName] = infoList[k].TypeStatus
	}
	m.Datetime = time.Now().String()[:19]
	m.UrlLink = siteInfo.UrlLink
	m.SiteName = siteName
	m.OrderNum = strings.Replace(time.Now().String()[:10], "-", "", 2) + siteId + siteIndexId + strconv.FormatInt(time.Now().Unix(), 10)
	m.TradeDate = time.Now().String()[:10]
	return m, nil
}

//得到页面
func (m *QuickPay) GetPage() []string {
	return []string{QUICK_PAY}
}

//得到<视讯电子彩票体育的页面>
func (m *QuickPay) GetSubPage() map[string]string {
	return nil
}
