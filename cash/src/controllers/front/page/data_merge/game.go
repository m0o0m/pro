package data_merge

import (
	"framework/render"
	//"global"
	"html/template"
	"models/back"
	"regexp"
	"strings"
)

type Game struct {
	HeaderFooter
	Key
	Notice  string //公告内容
	LogoUrl string //logo地址

	GameTitle []string      //电子导航
	GameData  EgameDataInfo //电子信息
	GameTheme template.HTML //电子内页主题信息

	Type string
}

type EgameDataInfo struct {
	Count int           `json:"count"` //游戏总数
	Data  []back.MgGame `json:"data"`
	Wh    int8          `json:"wh"`   // 1 维护 2  不维护
	Type  string        `json:"type"` // 类型
}

//得到电子数据
func (m *Game) GetData(siteId, siteIndexId string) (interface{}, error) {
	//初始化头部尾部数据
	err := m.initHeaderFooter(siteId, siteIndexId)
	m.Header.PageType = 3 //电子页面
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

	//电子导航查询
	gameModel, err := noteGameBean.IndexGameTitle(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	gameTitleList := strings.Split(gameModel, ",")
	for k, v := range gameTitleList {
		gameTitleList[k] = strings.ToUpper(strings.Split(v, "_")[0])
	}
	m.GameTitle = gameTitleList

	if len(m.Type) == 0 {
		m.Type = m.GameTitle[0]
	}

	////电子信息查询
	//gameList, err := noteGameBean.IndexGameList(siteId, siteIndexId, m.VType)
	//if err != nil {
	//	return nil, err
	//}
	//gameDataList := EgameDataInfo{}
	//gameDataList.Count = len(gameList)
	//gameDataList.Wh = 1
	//gameDataList.Data = gameList
	//gameDataList.VType = gameTitleList[0]

	if len(notices) > 0 {
		for k, _ := range notices {
			m.Notice += ";" + notices[k].NoticeContent
		}
		m.Notice = m.Notice[1:]
	}

	//获取电子内页主题样式
	has, activityPromotionSet, _ := noteGameBean.GameTheme(siteId, siteIndexId)
	if has {
		if len(activityPromotionSet.Bcolor) != 0 {
			colordate := strings.Split(activityPromotionSet.Bcolor, ",")
			m.GameTheme = template.HTML(EgameHtmlColor(colordate))
		}
	}

	m.LogoUrl = logUrl

	//m.GameData = gameDataList

	return m, err
}

//得到页面
func (m *Game) GetPage() []string {
	return []string{DZ, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

//得到<视讯电子彩票体育的页面>
func (m *Game) GetSubPage() map[string]string {
	return map[string]string{
		render.DzViewPath: DZ,
	}
}

//ajax获取电子数据
//func AjaxEgame(siteId, siteIndexId, Type string) (interface{}, error) {
//	result := EgameDataInfo{}
//	//电子信息查询
//	gameList, err := noteGameBean.IndexGameList(siteId, siteIndexId, Type)
//	if err != nil {
//		global.GlobalLogger.Error("err:s%", err.Error())
//		return nil, err
//	}
//	result.Count = len(gameList)
//	result.Wh = 1
//	result.Data = gameList
//	result.Type = Type
//
//	return result, err
//}

func EgameHtmlColor(colordate []string) (colorCss string) {
	colorCss = `<style type="text/css">`

	var hzRegexp = regexp.MustCompile("^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3})$")

	if len(colordate[0]) != 0 && hzRegexp.MatchString(colordate[0]) {
		colorCss = colorCss + `.tab1{background:` + colordate[0] + `}`
	}
	if len(colordate[1]) != 0 && hzRegexp.MatchString(colordate[1]) {
		colorCss = colorCss + `.tab1 .divgmenu .ul_ul li.zhu_gameClass.off .bg_col{background:` + colordate[1] + `}`
		colorCss = colorCss + `.tab1 .divgmenu .ul_ul li.zhu_gameClass.off .act-img{border-top:13px solid ` + colordate[1] + `}`
	}
	colorCss = colorCss + `.tab1 ul.game_category li a, .tab1 .search a.serch_but{`
	if len(colordate[2]) != 0 && hzRegexp.MatchString(colordate[2]) {
		colorCss = colorCss + `background:` + colordate[2] + `;`
	}
	if len(colordate[4]) != 0 && hzRegexp.MatchString(colordate[4]) {
		colorCss = colorCss + `border: 1px solid ` + colordate[4] + `;`
	}
	colorCss = colorCss + `}`
	if len(colordate[3]) != 0 && hzRegexp.MatchString(colordate[3]) {
		colorCss = colorCss + `.tab1 ul.game_category li a.active, .tab1 ul.game_category li a:hover{background:` + colordate[3] + `}`
	}
	if len(colordate[5]) != 0 && hzRegexp.MatchString(colordate[5]) {
		colorCss = colorCss + `.tab1 .menudiv{background:` + colordate[5] + `}`
	}
	colorCss = colorCss + `</style>`
	return colorCss
}
