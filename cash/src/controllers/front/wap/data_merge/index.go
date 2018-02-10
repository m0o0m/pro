package data_merge

import (
	"framework/render"
	"global"
	"models/back"
	"strings"
)

type Index struct {
	Key
	CdnUrl    string
	UrlLink   string              //客服链接
	SiteName  string              //站点名称
	LogoUrl   string              //log地址
	FlashList []*NIndexFlashImage //wap端轮播图
	Notice    string              //公告内容
	CurrPage  int64               //当前页面
	Content   []back.OrderModule  //视讯电子彩票体育
}
type NIndexFlashImage struct {
	ImgUrl  string //图片路径
	ImgLink string //链接地址
	IsLink  int    //是否新开页面 1是2否
}

func (m *Index) GetData(siteId, siteIndexId string) (data interface{}, err error) {
	m.CurrPage = 1
	m.CdnUrl = render.CdnUrl
	//轮播图
	flashList, err := siteOperateBean.FlashList(siteId, siteIndexId, 1, 2)
	if err != nil {
		return nil, err
	}
	for k, _ := range flashList {
		m.FlashList = append(m.FlashList, &NIndexFlashImage{flashList[k].ImgUrl, flashList[k].ImgLink, flashList[k].IsLink})
	}
	//查询公告
	notices, err := noticeBean.GetNoticeBySiteId(siteId)
	if err != nil {
		return nil, err
	}
	if len(notices) > 0 {
		for k, _ := range notices {
			m.Notice += ";" + notices[k].NoticeContent
		}
		m.Notice = m.Notice[1:]
	}
	//4大模块
	orderNumbers, err := webInfoBean.GetOrderNumberBySite(siteId, siteIndexId)
	if err != nil {
		return
	}
	modules := strings.Split(orderNumbers.Module, ",")
	products, err := productBean.ProductList()
	if err != nil {
		return
	}
	productMap := make(map[string]*back.ProductMappingList)
	for k, _ := range products {
		productMap[products[k].VType] = &back.ProductMappingList{
			Name:       products[k].ProductName,
			VType:      products[k].VType,
			PlatformId: products[k].PlatformId,
		}
	}

	f := func(vTypes []string) (mappings []*back.ProductMappingList) {
		for k, _ := range vTypes {
			mapping, ok := productMap[vTypes[k]]
			if ok {
				mappings = append(mappings, mapping)
			} else {
				global.GlobalLogger.Error("err:%s not exist", vTypes[k])
			}
		}
		return
	}
	var orderModules []back.OrderModule
	for _, v := range modules {
		var orderModule back.OrderModule
		switch v {
		case "video_module":
			orderModule.Module = back.ProductMapping{"视讯直播", "video"}
			orderModule.SubModule = f(strings.Split(orderNumbers.VideoModule, ","))
		case "fc_module":
			orderModule.Module = back.ProductMapping{"彩票游戏", "fc"}
			orderModule.SubModule = f(strings.Split(orderNumbers.FcModule, ","))
		case "dz_module":
			orderModule.Module = back.ProductMapping{"电子游艺", "dz"}
			orderModule.SubModule = f(strings.Split(orderNumbers.DzModule, ","))
		case "sp_module":
			orderModule.Module = back.ProductMapping{"体育赛事", "sp"}
			orderModule.SubModule = f(strings.Split(orderNumbers.SpModule, ","))
		default:
			global.GlobalLogger.Error("%s not in (video_module,fc_module,dz_module,sp_module)", v)
			continue
		}
		orderModules = append(orderModules, orderModule)
	}
	m.Content = orderModules

	//查询站点名称
	m.SiteName, err = siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return
	}
	//logo
	m.LogoUrl, err = siteOperateBean.Logo(siteId, siteIndexId, 2)
	if err != nil {
		return nil, err
	}
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
	return m, nil
}

func (*Index) GetPage() []string {
	return []string{WAP_INDEX, WAP_FOOTER}
}
