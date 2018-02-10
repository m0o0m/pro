package data_merge

import (
	"global"
	"models/back"
	"strings"
)

type NIndex struct {
	HeaderFooter
	Key
	Notice    string              //公告内容
	FlashList []*NIndexFlashImage //轮播图
	Platform  []back.TypeList     //快速额度转换平台
}

type NIndexFlashImage struct {
	ImgUrl  string //图片路径
	ImgLink string //链接地址
	IsLink  int    //是否新开页面 1是2否
}

//得到首页数据
func (m *NIndex) GetData(siteId, siteIndexId string) (interface{}, error) {
	//初始化头部尾部数据
	err := m.initHeaderFooter(siteId, siteIndexId)
	m.Header.PageType = 1 //我是首页
	if err != nil {
		return nil, err
	}
	m.Notice = m.Header.Notice //头部有就直接赋上去
	//查询公告
	//notices, err := noticeBean.GetNoticeBySiteId(siteId)
	//if err != nil {
	//	return nil, err
	//}
	//if len(notices) > 0 {
	//	for k, _ := range notices {
	//		m.Notice += ";" + notices[k].NoticeContent
	//	}
	//	m.Notice = m.Notice[1:]
	//}
	//轮播图
	flashList, err := siteOperateBean.FlashList(siteId, siteIndexId, 1, 1)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil, err
	}
	themeName, err := m.GetThemeName(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil, err
	}
	if !strings.HasPrefix(themeName, "/") {
		themeName = "/" + themeName
	}
	for k, _ := range flashList {
		m.FlashList = append(m.FlashList, &NIndexFlashImage{themeName + flashList[k].ImgUrl, flashList[k].ImgLink, flashList[k].IsLink})
	}

	//查询当前站点启用的平台
	m.Platform, err = productBean.SitePlatform(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil, err
	}

	//js, _ := json.MarshalIndent(m, " ", "	")
	//fmt.Println(string(js))

	return m, err
}

//得到页面
func (m *NIndex) GetPage() []string {
	return []string{N_INDEX, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

//得到<视讯电子彩票体育的页面>
func (m *NIndex) GetSubPage() map[string]string {
	return nil
}
