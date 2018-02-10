package data_merge

import (
	"framework/render"
	"models/back"
	"models/schema"
	"strings"
)

type Lottery struct {
	HeaderFooter
	Key
	Notice       string   //公告内容
	LogoUrl      string   //logo地址
	SiteName     string   //站点名称
	LotteryOrder []string //彩票平台排序
	CdnUrl       string
	PkData       map[string]map[string]interface{}
	EgData       map[string]map[string]interface{}
	CsData       map[string]map[string]interface{}
}

func (m *Lottery) GetData(siteId, siteIndexId string) (interface{}, error) {
	err := m.initHeaderFooter(siteId, siteIndexId)
	m.Header.PageType = 5 //彩票页面
	//查询体育公告
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

	//获取彩票平台和排名
	fc_order, err := siteOrderModule.LotteryModule(siteId, siteIndexId)

	Order := strings.Split(fc_order, ",")
	if len(Order) > 1 {
		for _, v := range Order {
			str := strings.Split(v, "_")
			if str[0] == "eg" {
				str[0] = "egtc"
			}
			m.LotteryOrder = append(m.LotteryOrder, str[0])
		}
	} else {
		m.LotteryOrder = nil
	}
	//获取pk彩种
	pkgames, _ := pkGames.PkGames()
	pkdata := make(map[string]map[string]interface{})
	pkdata["ssc"] = make(map[string]interface{})
	pkdata["yb"] = make(map[string]interface{})
	pkdata["xy"] = make(map[string]interface{})
	pkdata["x11"] = make(map[string]interface{})
	pkdata["k3"] = make(map[string]interface{})
	pkdata["klc"] = make(map[string]interface{})
	pkdata["sf"] = make(map[string]interface{})
	pkdata["gpc"] = make(map[string]interface{})
	for _, val := range pkgames {
		if val.LType == "ssc" {
			pkdata["ssc"]["name"] = "时时彩"
			pkdata["ssc"]["img"] = "/shared/lottlys/images/images/lotteryhall/shishicai.png"
			var types []back.PkLottery
			if temp, ok := pkdata["ssc"]["type"]; ok {
				types, _ = temp.([]back.PkLottery)
			}
			types = append(types, val)
			pkdata["ssc"]["type"] = types
		}
		if val.LType == "yb" {
			pkdata["yb"]["name"] = "一般彩票"
			pkdata["yb"]["img"] = "/shared/lottlys/images/images/lotteryhall/yibancaiqiu.png"
			var yb []back.PkLottery
			if temp, ok := pkdata["yb"]["type"]; ok {
				yb, _ = temp.([]back.PkLottery)
			}
			yb = append(yb, val)
			pkdata["yb"]["type"] = yb
		}
		if val.LType == "xy" {
			pkdata["xy"]["name"] = "幸运彩"
			pkdata["xy"]["img"] = "/shared/lottlys/images/images/lotteryhall/xingyunma.png"
			var xy []back.PkLottery
			if temp, ok := pkdata["xy"]["type"]; ok {
				xy, _ = temp.([]back.PkLottery)
			}
			xy = append(xy, val)
			pkdata["xy"]["type"] = xy
		}
		if val.LType == "11" {
			pkdata["x11"]["name"] = "十一选五"
			pkdata["x11"]["img"] = "/shared/lottlys/images/images/lotteryhall/shiyixuanwu.png"
			var x11 []back.PkLottery
			if temp, ok := pkdata["x11"]["type"]; ok {
				x11, _ = temp.([]back.PkLottery)
			}
			x11 = append(x11, val)
			pkdata["x11"]["type"] = x11
		}
		if val.LType == "k3" {
			pkdata["k3"]["name"] = "快三"
			pkdata["k3"]["img"] = "/shared/lottlys/images/images/lotteryhall/kuai3.png"
			var k3 []back.PkLottery
			if temp, ok := pkdata["k3"]["type"]; ok {
				k3, _ = temp.([]back.PkLottery)
			}
			k3 = append(k3, val)
			pkdata["k3"]["type"] = k3
		}
		if val.LType == "klc" {
			pkdata["klc"]["name"] = "快乐彩"
			pkdata["klc"]["img"] = "/shared/lottlys/images/images/lotteryhall/kuaile8.png"
			var klc []back.PkLottery
			if temp, ok := pkdata["klc"]["type"]; ok {
				klc, _ = temp.([]back.PkLottery)
			}
			klc = append(klc, val)
			pkdata["klc"]["type"] = klc
		}
		if val.LType == "sf" {
			pkdata["sf"]["name"] = "快乐十分"
			pkdata["sf"]["img"] = "/shared/lottlys/images/images/lotteryhall/kuaileshifen.png"
			var sf []back.PkLottery
			if temp, ok := pkdata["sf"]["type"]; ok {
				sf, _ = temp.([]back.PkLottery)
			}
			sf = append(sf, val)
			pkdata["sf"]["type"] = sf
		}
		if val.LType == "gpc" {
			pkdata["gpc"]["name"] = "幸运赛车"
			pkdata["gpc"]["img"] = "/shared/lottlys/images/images/lotteryhall/saichejingshu.png"
			var gpc []back.PkLottery
			if temp, ok := pkdata["gpc"]["type"]; ok {
				gpc, _ = temp.([]back.PkLottery)
			}
			gpc = append(gpc, val)
			pkdata["gpc"]["type"] = gpc
		}
	}
	m.PkData = pkdata

	if strings.Contains(fc_order, "eg_fc") {
		//获取eg彩票彩种信息
		eggames, err := egGames.EgGames()
		if err != nil {
			m.EgData = nil
		} else {
			egdata := make(map[string]map[string]interface{})
			egdata["hot"] = make(map[string]interface{})
			egdata["ssc"] = make(map[string]interface{})
			egdata["yb"] = make(map[string]interface{})
			egdata["xyc"] = make(map[string]interface{})
			egdata["x11"] = make(map[string]interface{})
			egdata["k3"] = make(map[string]interface{})
			egdata["klc"] = make(map[string]interface{})
			egdata["sf"] = make(map[string]interface{})
			for _, val := range eggames {
				if val.Hot == 1 {
					egdata["hot"]["name"] = "热门彩票"
					egdata["hot"]["img"] = "/shared/lottlys/images/images/lotteryhall/remen_cp.png"
					var hot []schema.EgGames
					if temp, ok := egdata["hot"]["type"]; ok {
						hot, _ = temp.([]schema.EgGames)
					}
					hot = append(hot, val)
					egdata["hot"]["type"] = hot
				}
				if val.EgLxType == "ssc" {
					egdata["ssc"]["name"] = "时时彩"
					egdata["ssc"]["img"] = "/shared/lottlys/images/images/lotteryhall/shishicai.png"
					var types []schema.EgGames
					if temp, ok := egdata["ssc"]["type"]; ok {
						types, _ = temp.([]schema.EgGames)
					}
					types = append(types, val)
					egdata["ssc"]["type"] = types
				}
				if val.EgLxType == "yb" {
					egdata["yb"]["name"] = "一般彩票"
					egdata["yb"]["img"] = "/shared/lottlys/images/images/lotteryhall/yibancaiqiu.png"
					var yb []schema.EgGames
					if temp, ok := egdata["yb"]["type"]; ok {
						yb, _ = temp.([]schema.EgGames)
					}
					yb = append(yb, val)
					egdata["yb"]["type"] = yb
				}
				if val.EgLxType == "xyc" {
					egdata["xyc"]["name"] = "幸运彩"
					egdata["xyc"]["img"] = "/shared/lottlys/images/images/lotteryhall/xingyunma.png"
					var xyc []schema.EgGames
					if temp, ok := egdata["xyc"]["type"]; ok {
						xyc, _ = temp.([]schema.EgGames)
					}
					xyc = append(xyc, val)
					egdata["xyc"]["type"] = xyc
				}
				if val.EgLxType == "11" {
					egdata["x11"]["name"] = "十一选五"
					egdata["x11"]["img"] = "/shared/lottlys/images/images/lotteryhall/shiyixuanwu.png"
					var x11 []schema.EgGames
					if temp, ok := egdata["x11"]["type"]; ok {
						x11, _ = temp.([]schema.EgGames)
					}
					x11 = append(x11, val)
					egdata["x11"]["type"] = x11
				}
				if val.EgLxType == "k3" {
					egdata["k3"]["name"] = "快三"
					egdata["k3"]["img"] = "/shared/lottlys/images/images/lotteryhall/kuai3.png"
					var k3 []schema.EgGames
					if temp, ok := egdata["k3"]["type"]; ok {
						k3, _ = temp.([]schema.EgGames)
					}
					k3 = append(k3, val)
					egdata["k3"]["type"] = k3
				}
				if val.EgLxType == "klc" {
					egdata["klc"]["name"] = "快乐彩"
					egdata["klc"]["img"] = "/shared/lottlys/images/images/lotteryhall/saichejingshu.png"
					var klc []schema.EgGames
					if temp, ok := egdata["klc"]["type"]; ok {
						klc, _ = temp.([]schema.EgGames)
					}
					klc = append(klc, val)
					egdata["klc"]["type"] = klc
				}
				if val.EgLxType == "sf" {
					egdata["sf"]["name"] = "快乐十分"
					egdata["sf"]["img"] = "/shared/lottlys/images/images/lotteryhall/kuaileshifen.png"
					var sf []schema.EgGames
					if temp, ok := egdata["sf"]["type"]; ok {
						sf, _ = temp.([]schema.EgGames)
					}
					sf = append(sf, val)
					egdata["sf"]["type"] = sf
				}
			}
			m.EgData = egdata
		}
	}

	if strings.Contains(fc_order, "cs_fc") {
		//获取cs彩票彩种信息
		csgames, err := csGames.CsGames()
		if err != nil {
			m.EgData = nil
		} else {
			csdata := make(map[string]map[string]interface{})
			csdata["hot"] = make(map[string]interface{})
			csdata["efc"] = make(map[string]interface{})
			csdata["ssc"] = make(map[string]interface{})
			csdata["lhc"] = make(map[string]interface{})
			csdata["pk10"] = make(map[string]interface{})
			csdata["g11"] = make(map[string]interface{})
			csdata["k3"] = make(map[string]interface{})
			csdata["fc"] = make(map[string]interface{})
			csdata["pl5"] = make(map[string]interface{})
			for _, val := range csgames {
				if val.Hot == 1 {
					csdata["hot"]["name"] = "热门彩票"
					csdata["hot"]["img"] = "/shared/lottlys/images/cs/ahot.png"
					var hot []schema.CsGames
					if temp, ok := csdata["hot"]["type"]; ok {
						hot, _ = temp.([]schema.CsGames)
					}
					hot = append(hot, val)
					csdata["hot"]["type"] = hot
				}
				if val.CsColumn == "efc" {
					csdata["efc"]["name"] = "卡司二分彩"
					csdata["efc"]["img"] = "/shared/lottlys/images/cs/befc1.png"
					var types []schema.CsGames
					if temp, ok := csdata["efc"]["type"]; ok {
						types, _ = temp.([]schema.CsGames)
					}
					types = append(types, val)
					csdata["efc"]["type"] = types
				}
				if val.CsLxType == "ssc" {
					csdata["ssc"]["name"] = "时时彩"
					csdata["ssc"]["img"] = "/shared/lottlys/images/cs/cssc.png"
					var types []schema.CsGames
					if temp, ok := csdata["ssc"]["type"]; ok {
						types, _ = temp.([]schema.CsGames)
					}
					types = append(types, val)
					csdata["ssc"]["type"] = types
				}
				if val.CsLxType == "lhc" {
					csdata["lhc"]["name"] = "六合彩"
					csdata["lhc"]["img"] = "/shared/lottlys/images/cs/elhc.png"
					var lhc []schema.CsGames
					if temp, ok := csdata["yb"]["type"]; ok {
						lhc, _ = temp.([]schema.CsGames)
					}
					lhc = append(lhc, val)
					csdata["lhc"]["type"] = lhc
				}
				if val.CsLxType == "pk10" {
					csdata["pk10"]["name"] = "幸运飞车"
					csdata["pk10"]["img"] = "/shared/lottlys/images/cs/fpk10.png"
					var pk10 []schema.CsGames
					if temp, ok := csdata["xyc"]["type"]; ok {
						pk10, _ = temp.([]schema.CsGames)
					}
					pk10 = append(pk10, val)
					csdata["pk10"]["type"] = pk10
				}
				if val.CsLxType == "11" {
					csdata["g11"]["name"] = "十一选五"
					csdata["g11"]["img"] = "/shared/lottlys/images/cs/g11.png"
					var g11 []schema.CsGames
					if temp, ok := csdata["x11"]["type"]; ok {
						g11, _ = temp.([]schema.CsGames)
					}
					g11 = append(g11, val)
					csdata["g11"]["type"] = g11
				}
				if val.CsLxType == "k3" {
					csdata["k3"]["name"] = "快三"
					csdata["k3"]["img"] = "/shared/lottlys/images/cs/hk3.png"
					var k3 []schema.CsGames
					if temp, ok := csdata["k3"]["type"]; ok {
						k3, _ = temp.([]schema.CsGames)
					}
					k3 = append(k3, val)
					csdata["k3"]["type"] = k3
				}
				if val.CsLxType == "fc" {
					csdata["fc"]["name"] = "福彩3D"
					csdata["fc"]["img"] = "/shared/lottlys/images/cs/ifc.png"
					var fc []schema.CsGames
					if temp, ok := csdata["klc"]["type"]; ok {
						fc, _ = temp.([]schema.CsGames)
					}
					fc = append(fc, val)
					csdata["fc"]["type"] = fc
				}
				if val.CsLxType == "pl5" {
					csdata["pl5"]["name"] = "排列五"
					csdata["pl5"]["img"] = "/shared/lottlys/images/cs/jpl5.png"
					var pl5 []schema.CsGames
					if temp, ok := csdata["sf"]["type"]; ok {
						pl5, _ = temp.([]schema.CsGames)
					}
					pl5 = append(pl5, val)
					csdata["pl5"]["type"] = pl5
				}
			}
			m.CsData = csdata
		}
	}
	return m, err
}

func (*Lottery) GetPage() []string {
	return []string{LOTTERY, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

func (*Lottery) GetSubPage() map[string]string {
	return map[string]string{
		render.FcViewPath: FC,
	}
}
