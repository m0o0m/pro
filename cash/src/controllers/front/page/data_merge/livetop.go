package data_merge

import (
	"fmt"
	"framework/render"
	"models/back"
	"strings"
)

type LiveTop struct {
	HeaderFooter
	Key
	Notice    string        //公告内容
	LogoUrl   string        //logo地址
	SiteName  string        //站点名称
	LiveOrder []LiveTopName //视讯平台排序
	LiveLen   int           //视讯平台个数
}

type LiveTopName struct {
	ProductName string `json:"product_name"` //商品名
	VType       string `json:"v_type"`       //游戏类型
}

//得视讯页面的数据
func (m *LiveTop) GetData(siteId, siteIndexId string) (interface{}, error) {
	err := m.initHeaderFooter(siteId, siteIndexId)
	m.Header.PageType = 2 //视讯页面
	//查询视讯公告
	notices, _ := noticeBean.GetNoticeBySiteId(siteId)
	//logo
	logUrl, _ := siteOperateBean.Logo(siteId, siteIndexId, 1)
	//查询站点
	siteName, _ := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)

	if len(notices) > 0 {
		for k, _ := range notices {
			m.Notice += ";" + notices[k].NoticeContent
		}
		m.Notice = m.Notice[1:]
	}
	m.LogoUrl = logUrl
	m.SiteName = siteName

	//获取视讯游戏和排序
	live_order, err := siteOrderModule.LiveModule(siteId, siteIndexId)
	//获取商品表数据
	products, err := siteProductBean.GetProductList()
	var PName LiveTopName
	Order := strings.Split(live_order, ",")
	if len(Order) > 1 {
		for _, v := range Order {
			PName.VType = v
			for _, val := range products {
				if v == val.VType {
					PName.ProductName = val.ProductName
				}
			}
			m.LiveOrder = append(m.LiveOrder, PName) //将v_type与名字对应
		}
	} else {
		m.LiveOrder = nil
	}
	m.LiveLen = len(m.LiveOrder)
	fmt.Println(m)
	return m, err
}

func (*LiveTop) GetPage() []string {
	return []string{LIVE_TOP, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

func (*LiveTop) GetSubPage() map[string]string {
	return map[string]string{
		render.VdViewPath: VD,
	}
}

//获取视讯页面的数据
func GetLiveTopAjax(siteId, siteIndexId string) (interface{}, error) {
	result := new(back.LiveTopAjaxData)

	//查询视讯模版ID
	liveid, _ := siteInfoVideoUser.LiveStyle(siteId, siteIndexId)
	if liveid > 0 {
		result.LiveId = liveid
	} else {
		result.LiveId = 0
	}

	//获取视讯游戏和排序
	live_order, err := siteOrderModule.LiveModule(siteId, siteIndexId)
	//获取商品表数据
	products, err := siteProductBean.GetProductList()
	var PName back.LiveTopName
	Order := strings.Split(live_order, ",")
	if len(Order) > 1 {
		for _, v := range Order {
			PName.VType = v
			for _, val := range products {
				if v == val.VType {
					PName.ProductName = val.ProductName
				}
			}
			result.LiveOrder = append(result.LiveOrder, PName) //将v_type与名字对应
		}
	} else {
		result.LiveOrder = nil
	}
	return result, err
}

func (*LiveTopName) Mult(a int, b float64) (c float64) {
	c = float64(a) * b
	return
}
