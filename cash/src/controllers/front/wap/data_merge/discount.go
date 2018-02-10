package data_merge

import (
	"encoding/json"
	"fmt"
	"framework/render"
)

//优惠活动
type Discount struct {
	Key
	CdnUrl   string
	Notice   string //公告内容
	LogoUrl  string //logo地址
	CurrPage int64  //当前页面
	UrlLink  string
}

//平台游戏数据
type DiscountInfo struct {
	Id    int64  `json:"id"`    //游戏总数
	Title string `json:"title"` //游戏总数
}

//得到电子数据
func (m *Discount) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CurrPage = 2
	m.CdnUrl = render.CdnUrl
	//查询公告
	notices, err := noticeBean.GetNoticeBySiteId(siteId)
	if err != nil {
		return nil, err
	}

	//logo
	logUrl, err := siteOperateBean.Logo(siteId, siteIndexId, 2)
	if err != nil {
		return nil, err
	}

	if len(notices) > 0 {
		for k, _ := range notices {
			m.Notice += ";" + notices[k].NoticeContent
		}
		m.Notice = m.Notice[1:]
	}

	m.LogoUrl = logUrl

	js, _ := json.MarshalIndent(m, " ", "	")
	fmt.Println("js", string(js))
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
	return m, err
}

//得到页面
func (m *Discount) GetPage() []string {
	return []string{WAP_APPLY_PRO, WAP_FOOTER}
}
