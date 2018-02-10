package data_merge

import (
	"encoding/json"
	"fmt"
	"framework/render"
	"global"
	"html/template"
)

type Youhui struct {
	HeaderFooter
	Key
	Notice  string           //公告内容
	LogoUrl string           //logo地址
	YhTitle []YouhuiDataInfo //优惠分类信息
	YhData  []SiteActivity
	YhWidth int64 // 优惠宽度
}

//平台游戏数据
type YouhuiDataInfo struct {
	Id    int64  `json:"id"`    //游戏总数
	Title string `json:"title"` //游戏总数
}

//站点优惠活动
type SiteActivity struct {
	Id      int64         `xorm:"'id' PK autoincr"` //id
	TopId   int64         `xorm:"top_id"`           //上级栏目
	Title   string        `xorm:"title"`            //标题
	Content template.HTML `xorm:"content"`          //内容
	Img     string        `xorm:"img"`              //标题图片路径
}

//得到电子数据
func (m *Youhui) GetData(siteId, siteIndexId string) (interface{}, error) {
	//初始化头部尾部数据
	err := m.initHeaderFooter(siteId, siteIndexId)
	m.Header.PageType = 6 //优惠
	if err != nil {
		return nil, err
	}

	//查询公告
	notices, err := noticeBean.GetNoticeBySiteId(siteId)
	if err != nil {
		return nil, err
	}

	//logo
	logUrl, err := siteOperateBean.Logo(siteId, siteIndexId, 1)
	if err != nil {
		return nil, err
	}

	//优惠活动查询
	youhuiList, err := siteIWordBean.IndexActivityList(siteId, siteIndexId, 1)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return nil, err
	}

	yhTitleData := []YouhuiDataInfo{}
	yhData := []SiteActivity{}
	//获取分类活动
	for _, v := range youhuiList {
		if v.TopId == 0 && len(v.TypeName) != 0 {
			titleInfo := YouhuiDataInfo{}
			titleInfo.Id = v.Id
			titleInfo.Title = v.Title
			yhTitleData = append(yhTitleData, titleInfo)
		} else {
			info := SiteActivity{}
			info.Id = v.Id
			info.Title = v.Title
			info.TopId = v.TopId
			info.Img = v.Img
			info.Content = template.HTML(v.Content)
			yhData = append(yhData, info)
		}
	}

	m.YhData = yhData
	m.YhTitle = yhTitleData

	//获取优惠宽度
	has, activityPromotionSet, _ := noteGameBean.GameTheme(siteId, siteIndexId)
	if has {
		if activityPromotionSet.MaxWidth == 0 {
			m.YhWidth = 960
		} else {
			m.YhWidth = activityPromotionSet.MaxWidth
		}
	}

	if len(notices) > 0 {
		for k, _ := range notices {
			m.Notice += ";" + notices[k].NoticeContent
		}
		m.Notice = m.Notice[1:]
	}

	m.LogoUrl = logUrl

	js, _ := json.MarshalIndent(m, " ", "	")
	fmt.Println("js", js)

	return m, err
}

//得到页面
func (m *Youhui) GetPage() []string {
	return []string{YH, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

//得到<视讯电子彩票体育的页面>
func (m *Youhui) GetSubPage() map[string]string {
	return map[string]string{
		render.YhViewPath: YH,
	}
}
