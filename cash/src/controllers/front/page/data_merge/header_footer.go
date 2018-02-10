package data_merge

import (
	"framework/render"
	"global"
	"html/template"
	"models/back"
	"models/input"
	"strings"
)

type HeaderFooter struct {
	Header Header
	Footer Footer //页尾
}

type Header struct {
	LogoUrl  string             //logo地址
	Notice   string             //公告内容
	SiteName string             //站点名称
	PageType int                //页面类型1为首页,2视讯,3电子,4体育,5彩票,6优惠,7会员注册,8文案页面,9线路检测
	Content  []back.OrderModule //页头
	UrlLink  string             //客服连接地址
	CdnUrl   string             //cdn地址
	Phone    []string           //电话
	Email    []string           //邮箱
}

type Footer struct {
	CdnUrl      string            //cdn地址
	SiteName    string            //站点名称
	Content     []back.IndexWord  //页尾内容
	LeftFloat   []back.FloatList  //左浮动框
	RightFloat  []back.FloatList  //右浮动框
	PopAdv      []back.WebPopAdv  //弹窗广告
	PopAdvC     back.WebAdvColor  //弹窗广告
	LeftAdvList []back.WebAdvList //左边广告
	Phone       []string          //电话
	Email       []string          //邮箱
	QQ          []string          //QQ
	UrlLink     string            //客服连接地址
}

//初始化头部尾部数据
func (m *HeaderFooter) initHeaderFooter(siteId, siteIndexId string) (err error) {
	err = m.initHeaderData(siteId, siteIndexId)
	if err != nil {
		return
	}
	err = m.initFooterData(siteId, siteIndexId)
	return
}

//页头数据
func (m *HeaderFooter) initHeaderData(siteId, siteIndexId string) (err error) {
	m.Header.CdnUrl = render.CdnUrl
	var orderModules []back.OrderModule
	orderNumbers, err := webInfoBean.GetOrderNumberBySite(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	//logo
	m.Header.LogoUrl, _ = siteOperateBean.Logo(siteId, siteIndexId, 1)
	//if err != nil {
	//	return err
	//}
	//查询公告
	notices, _ := noticeBean.GetNoticeBySiteId(siteId)
	//if err != nil {
	//	return err
	//}
	if len(notices) > 0 {
		for k, _ := range notices {
			m.Header.Notice += ";" + notices[k].NoticeContent
		}
		m.Header.Notice = m.Header.Notice[1:]
	}

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
	modules := strings.Split(orderNumbers.Module, ",")
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
	var data back.GetSiteInfoPcAndWap
	var has bool
	data, has, err = webInfoBean.GetSiteInfo(&input.OrderModuleList{siteId, siteIndexId})
	var Phone, Email string
	if has {
		m.Header.UrlLink = data.Pc.UrlLink
		Phone = data.Pc.Phone
		Email = data.Pc.Email
	} else {
		m.Header.UrlLink = "未找到客服链接"
	}

	m.Header.Phone = strings.Split(Phone, ",")
	m.Header.Email = strings.Split(Email, ",")

	//查询站点
	m.Header.SiteName, err = siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return
	}
	m.Header.Content = orderModules
	return
}

//页脚数据
func (m *HeaderFooter) initFooterData(siteId, siteIndexId string) (err error) {
	m.Footer.CdnUrl = render.CdnUrl
	//页尾的关于我们那一排菜单
	m.Footer.Content, err = siteIWordBean.IWordList(&input.SiteIWodList{
		SiteId:      siteId,
		SiteIndexId: siteIndexId,
		State:       1, //状态正常
		Itype:       1, //类型:站点文案
	})
	if err != nil {
		return
	}
	//左或者右的浮动框数据
	floatList, err := webFloatBean.GetListBySite(siteId, siteIndexId)
	if err != nil {
		return
	}
	for k, _ := range floatList {
		if floatList[k].Ftype == 1 {
			m.Footer.LeftFloat = append(m.Footer.LeftFloat, floatList[k])
		} else {
			m.Footer.RightFloat = append(m.Footer.RightFloat, floatList[k])
		}
	}

	if len(m.Footer.LeftFloat) > 0 {
		num := len(m.Footer.LeftFloat)
		for k, _ := range m.Footer.LeftFloat {
			if k == num-1 {
				m.Footer.LeftFloat[k].UrlInter = "FloatClose(this);"
				m.Footer.LeftFloat[k].Url = "###"
			}
		}
	}

	if len(m.Footer.RightFloat) > 0 {
		num := len(m.Footer.RightFloat)
		for k, _ := range m.Footer.RightFloat {
			if k == num-1 {
				m.Footer.RightFloat[k].UrlInter = "FloatClose(this);"
				m.Footer.RightFloat[k].Url = "###"
			}
		}
	}

	//查询站点信息
	siteInfo, _, _ := siteOperateBean.GetSingleSite(siteIndexId, siteId)
	m.Footer.PopAdvC.PopoverBgColor = siteInfo.PopoverBgColor
	m.Footer.PopAdvC.PopoverTitleColor = siteInfo.PopoverTitleColor
	m.Footer.PopAdvC.PopoverBarColor = siteInfo.PopoverBarColor
	m.Footer.PopAdvC.AdWay = 2
	//弹窗和左下角广告
	advList, err := webAdv.GetList(siteId, siteIndexId)
	if err != nil {
		return
	}
	for k, _ := range advList {
		if advList[k].Type == 1 {
			info := back.WebPopAdv{}
			info.Id = advList[k].Id
			info.Title = advList[k].Title
			info.Content = template.HTML(advList[k].Content)
			info.State = advList[k].State
			m.Footer.PopAdv = append(m.Footer.PopAdv, info)
		} else {
			m.Footer.LeftAdvList = append(m.Footer.LeftAdvList, advList[k])
		}
	}

	if m.Header.SiteName != "" {
		m.Footer.SiteName = m.Header.SiteName
	} else {
		//查询站点
		m.Footer.SiteName, err = siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	}

	var data back.GetSiteInfoPcAndWap
	var has bool
	data, has, err = webInfoBean.GetSiteInfo(&input.OrderModuleList{siteId, siteIndexId})
	var Phone, Email, QQ string
	if has {
		m.Footer.UrlLink = data.Pc.UrlLink
		Phone = data.Pc.Phone
		Email = data.Pc.Email
		QQ = data.Pc.Qq
	} else {
		m.Footer.UrlLink = "未找到客服链接"
	}

	m.Footer.Phone = strings.Split(Phone, ",")
	m.Footer.Email = strings.Split(Email, ",")
	m.Footer.QQ = strings.Split(QQ, ",")

	return
}

//页脚数据
func GetFooterData(siteId, siteIndexId string) ([]back.IndexWord, error) {
	return siteIWordBean.IWordList(&input.SiteIWodList{
		SiteId:      siteId,
		SiteIndexId: siteIndexId,
		State:       1, //状态正常
		Itype:       1, //类型:站点文案
	})
}

//页头数据
func GetHeaderData(siteId, siteIndexId string) (orderModules []back.OrderModule, err error) {
	orderNumbers, err := webInfoBean.GetOrderNumberBySite(siteId, siteIndexId)
	if err != nil {
		return
	}
	products, err := productBean.ProductList()
	if err != nil {
		return
	}
	productMap := make(map[string]*back.ProductMappingList)
	for k, _ := range products {
		productMap[products[k].VType] = &back.ProductMappingList{
			Name:  products[k].ProductName,
			VType: products[k].VType,
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
	modules := strings.Split(orderNumbers.Module, ",")
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
	return
}
