package rbac

import (
	"controllers"
	"fmt"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"

	"time"
)

type ReportController struct {
	controllers.BaseController
}

//数据中心
func (this *ReportController) GetCenter(ctx echo.Context) error {
	betReportAccount := new(input.BetReportAccount)
	code := global.ValidRequest(betReportAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	listParams := new(global.ListParams)
	this.GetParam(listParams, ctx)
	times := new(global.Times)
	if len(betReportAccount.StartTime) != 0 {
		times.StartTime, code = global.FormatTime2Timestamp2(betReportAccount.StartTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.StartTime = global.GetTimeStart(global.GetCurrentTime())
	}
	if len(betReportAccount.EndTime) != 0 {
		times.EndTime, code = global.FormatTime2Timestamp2(betReportAccount.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.EndTime = global.GetTimeStart(global.GetCurrentTime())
	}
	//获取数据
	data, count, err := reportBean.GetCenter(betReportAccount, listParams, times)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("data", data, "count", count))
}

//获取报表查询页面的参数
func (this *ReportController) RepSearch(ctx echo.Context) error {
	sdata := make(map[string]interface{})
	sdata["PorductList"] = []back.ProductlistRep{}
	sdata["SiteIdList"] = []back.InSiteList{}
	sdata["combo"] = []back.GetList{}
	// 获取商品数据
	data, err := productBean.ProductList()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取游戏类型
	sdata["PorductList"] = data
	sdata["SiteIdList"], err = siteOperateBean.SiteList()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取套餐列表
	sdata["combo"], err = comboBeen.GetComboDrop(new(schema.Combo))
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("data", sdata))
}

//查询报表数据
func (this *ReportController) ReportList(ctx echo.Context) error {
	reportExport := new(input.ReportBills)
	code := global.ValidRequest(reportExport, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if len(reportExport.SiteId) == 0 || len(reportExport.VType) == 0 {
		return ctx.JSON(60201, global.ReplyError(code, ctx))
	}
	//账单数据
	sdata := []back.ReportExport{}
	//excel表格数据
	excelData := make(map[string]interface{})
	//查询账单商品名和商品id
	productList, _ := productBean.ProductList()

	fmt.Println(excelData)
	//查询站点套餐信息
	indexList, err := siteOperateBean.SiteCombo(reportExport.SiteId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//将商品名和商品v_type加进套餐信息中
	for k, v := range indexList {
		info := back.ReportExport{}
		info.SiteId = v.SiteId
		info.SiteIndexId = v.SiteIndexId
		info.SiteName = v.SiteName
		info.ComboName = v.ComboName

		for _, v1 := range reportExport.VType {
			vTypeInfo := back.ReportExportList{}
			for _, v2 := range productList {

				if v1 == v2.VType {
					vTypeInfo.VType = v1
					vTypeInfo.ProductId = v2.Id
					vTypeInfo.ProductName = v2.ProductName
					break
				}
			}

			info.List = append(info.List, vTypeInfo)
		}
		for k1, v1 := range v.List {
			for _, v2 := range productList {

				if v1.ProductId == v2.Id {
					indexList[k].List[k1].VType = v2.VType
					indexList[k].List[k1].ProductName = v2.ProductName
					break
				}
			}

			for k2, v2 := range info.List {

				if v2.ProductId == v1.ProductId {
					info.List[k2].Proportion = v1.Proportion
				}
			}
		}
		sdata = append(sdata, info)
	}

	//获取导出数据
	data, err := reportBean.ReportExport(reportExport)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	for _, v := range data {
		for k1, v1 := range indexList {
			for k2, v2 := range v1.List {

				if v.SiteId == v1.SiteId && v.VType == v2.VType {
					//判断是否是彩金
					if strings.Contains(v.VType, "jack") {
						indexList[k1].List[k2].Win = v.Jack
						break
					}
					indexList[k1].List[k2].Win = v.Win
					break

				}
			}
		}
		for k1, v1 := range sdata {
			for k2, v2 := range v1.List {
				if v.SiteId == v1.SiteId && v.VType == v2.VType {
					//判断是否是彩金
					if strings.Contains(v.VType, "jack") {
						sdata[k1].List[k2].Win = v.Jack
						break
					}
					sdata[k1].List[k2].Win = v.Win
					break
				}
			}
		}
		for k1, v1 := range sdata {
			for k2, v2 := range v1.List {
				if v.SiteId == v1.SiteId && v.VType == v2.VType {
					//判断是否是彩金
					if strings.Contains(v.VType, "jack") {
						sdata[k1].List[k2].Win = v.Jack
						break
					}
					sdata[k1].List[k2].Win = v.Win
					break
				}
			}
		}
	}

	//获取上月报表负数
	//negativeList, _ :=  reportBean.NegativeList(reportExport.SiteId, reportExport.StartTime)
	for k, v := range sdata {
		for _, v1 := range indexList {
			if v.SiteId == v1.SiteId {
				sdata[k].Html, _ = ExcelHtml(v1, reportExport.StartTime, reportExport.EndTime)
			}
		}
	}

	//获取上月报表负数
	//negativeList, _ :=  reportBean.NegativeList(reportExport.SiteId, reportExport.StartTime)
	for k, v := range sdata {
		for _, v1 := range indexList {
			if v.SiteId == v1.SiteId {
				sdata[k].Html, _ = ExcelHtml(v1, reportExport.StartTime, reportExport.EndTime)
			}
		}
	}

	//sdata[0].Html = html.EscapeString("<td>asdf</td>")
	fmt.Println("返回的列表数据：", sdata)
	fmt.Println("套餐信息：", indexList)

	return ctx.JSON(200, global.ReplyItem(sdata))
}

func ExcelHtml(data back.ReportExport, startTime, endTime string) (htmlStr string, err error) {

	//将套餐转成大写
	str := strings.ToUpper(data.ComboName)

	//获取excel表格中的月份
	loc, _ := time.LoadLocation("Local") //时区
	t, err := time.ParseInLocation("2006-01-02", startTime, loc)
	startM := MonthNum(t.Month().String())

	endM := MonthNum(t.AddDate(0, 1, 0).Month().String())
	fmt.Println(startM, endM)

	//站点模板
	//types :=make(map[string]string)
	//types["sp"] = "体育"
	//types["fc"] = "彩票"
	//types["sb"] = "SB体育"
	//types["im"] = "IM体育"
	//types["imdx"] = "IM兑现佣金"
	//types["bbin"] = "BBIN全平台"
	//types["mg"] = "MG视讯"
	//types["mgdz"] = "MG电子"
	//types["ag"] = "AG视讯"
	//types["agdz"] = "AG电子"
	//types["agter"] = "AG捕鱼"
	//types["agjack"] = "AG捕鱼彩金"
	//types["gg"] = "GG捕鱼"
	//types["ggjack"] = "GG捕鱼彩金"
	//types["lebo"] = "LEBO视讯"
	//types["og"] = "OG视讯"
	//types["gpi"] = "GPI视讯"
	//types["gpidz"] = "GPI电子"
	//types["gpidzjack"] = "GPI电子彩金"
	//types["lmg"] = "LMG视讯"
	//types["sa"] = "SA视讯"
	//types["dg"] = "DG视讯"
	//types["dg_red"] = "DG(红包)"
	//types["dgxf"] = "DG(小费)"
	//types["gd"] = "GD视讯"
	//types["gddz"] = "GD电子"
	//types["gddzjack"] = "GD电子彩金"
	//types["ab"] = "AB视讯"
	//types["ct"] = "CT视讯"
	//types["pt"] = "PT电子"
	//types["ptjack"] = "PT彩金"
	//types["eg"] = "EG电子"
	//types["egjack"] = "EG电子彩金"
	//types["hb"] = "HABA电子"
	//types["hbjack"] = "HABA彩金"
	//types["egtc"] = "EG彩票"
	//types["cs"] = "官方彩票"

	var oughtCount, negatCount, linecost, moneyCount float64
	for _, v := range data.List {

		moneyCount = moneyCount + v.Win
	}
	//对应金额
	htmlStyle := fmt.Sprintf(`<style type="text/css">
.table-package{border-collapse:collapse;border:none;width: 1000px !important;}
.table-package td{border:solid #000 1px !important;color: #000 !important;line-height: 1 !important;font-family:'宋体' !important;font-size:16px !important;}
.table-package tr{height: 24px !important;line-height: 1 !important;}
.table-package .l{text-align: left !important;}
.table-package .r{text-align: right !important;}
.table-package .c{text-align: center !important;}
</style>`)
	htmlTbody := ""
	htmlStr = ""
	if str == "A套餐" {
		for _, v := range data.List {
			htmlTbody += "<tr>"
			if v == data.List[0] {
				htmlTbody += "<td rowspan='" + fmt.Sprintf("<div>%v</div>", len(data.List)+2) + "' class='c'>" + data.SiteName + "<br><br>(A套餐)</td>"
			}
			htmlTbody += "<td>" + v.ProductName + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", v.Win) + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"</tr>"
		}
		htmlStr = fmt.Sprintf(`%v
<table cellpadding="0" cellspacing="0" class="table-package">
    <tr>
        <td colspan="9" bgcolor="#FFCC99" class='c'>日期：%v ~ %v</td>
    </tr>
    <tr>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>網站</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>項目</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>金額</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>上月负数</td>
        <td width="10%%" bgcolor="#CCFFCC" class='c'>條件</td>
        <td width="10%%" bgcolor="#CCFFCC" class='c'>负数额度</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>應收</td>
        <td width="10%%" bgcolor="#FFFF00" class='c'>应加应扣</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>備註</td>
    </tr>
    <tr>
        <td bgcolor="#CCFFCC" class='c'>比例</td>
        <td bgcolor="#CCFFCC" class='c'>本月应加</td>
        <td bgcolor="#FFFF00" class='c'>视讯额度</td>
    </tr>
	%v
    <tr>
        <td>其他</td>
        <td></td><td></td>
        <td></td><td></td>
        <td class='c'>0.00</td><td></td><td></td>
    </tr>
    <tr style="font-weight: bold">
        <td>总计</td>
        <td bgcolor="#CCFFCC" class='c'>%v</td>
        <td bgcolor="#CCFFCC" class='c'>%v</td><td></td><td></td>
        <td bgcolor="#CCFFCC" class='c'>%v</td><td></td><td></td>
    </tr>
    <tr>
        <td colspan="9" bgcolor="#FFFF00">
            <p>注：1.客户接到财务对账单，三天内如有数据疑问请联系财务专员，财务专员SKYPE:pk.finance<br>&nbsp;&nbsp;2.如无疑问，汇款前请找财务专员索取最新汇款账号，感谢您们的支持，PK娱乐平台团队致敬！<br>&nbsp;&nbsp;3.汇款账号发出后，半小时内没汇好款，要汇款之前请再联系我们确认汇款账号，以免有变动。</p>
        </td>
    </tr>
    <tr>
        <td colspan="4">A套餐（當期負數可累积）</td><td></td><td></td><td></td><td></td><td></td>
    </tr>
    <tr>
        <td colspan="2">PK平臺體育、彩票</td><td colspan="2" class='c'>8%%</td>
        <td colspan="2" bgcolor="#E26B0A">本月抽成</td><td bgcolor="#E26B0A" class='c'> %v </td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>MG視訊</td><td class='c'>12%%</td>
        <td>MG電子</td><td class='c'>15%%</td>
        <td colspan="2" bgcolor="#E26B0A">%v月预收线路费</td><td bgcolor="#E26B0A" class='c'> %v </td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td colspan="3">AG視訊、LEBO視訊、CT視訊、OG視訊</td><td class='c'>12%%</td>
        <td colspan="2" bgcolor="#E26B0A">%v月实收线路费</td><td bgcolor="#E26B0A" class='c'> %v </td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td colspan="3">BBIN（體育、視訊、電子、彩票）</td><td class='c'>12%%</td>
        <td colspan="2" bgcolor="#E26B0A">%v月预收线路费</td><td bgcolor="#E26B0A" class='c'> %v </td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>PT电子</td><td class='c'>17%%</td>
        <td>EG电子</td><td class='c'>10%%</td>
        <td colspan="2" bgcolor="#E26B0A">本月应缴金额</td><td bgcolor="#E26B0A" class='c'> %v </td>
        <td></td>
        <td></td>
    </tr>


</table>`, htmlStyle, startTime, endTime, htmlTbody, moneyCount, negatCount, oughtCount, linecost, startM, linecost, startM, linecost, endM, linecost, linecost)
	} else if str == "B套餐" {
		for _, v := range data.List {
			htmlTbody += "<tr>"
			if v == data.List[0] {
				htmlTbody += "<td rowspan='" + fmt.Sprintf("<div>%v</div>", len(data.List)+2) + "' class='c'>" + data.SiteName + "<br><br>(B套餐)</td>"
			}
			htmlTbody += "<td>" + v.ProductName + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", v.Win) + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"</tr>"
		}
		htmlStr = fmt.Sprintf(`%v
<table cellpadding="0" cellspacing="0" class="table-package">
    <tr>
        <td colspan="7" bgcolor="#FFCC99" class='c'>日期：%v~%v</td>
    </tr>
    <tr>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>網站</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>項目</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>金額</td>
        <td width="10%%" bgcolor="#CCFFCC" class='c'>條件</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>應收</td>
        <td width="10%%" bgcolor="#FFFF00" class='c'>应加应扣</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>備註</td>
    </tr>
    <tr>
        <td bgcolor="#CCFFCC" class='c'>比例</td>
        <td bgcolor="#FFFF00" class='c'>视讯额度</td>
    </tr>
    %v
    <tr>
        <td>其他</td>
        <td></td>
        <td></td>
        <td class='c'>0.00</td><td></td><td></td>
    </tr>
    <tr style="font-weight: bold">
        <td>总计</td>
        <td bgcolor="#CCFFCC" class='c'>%v</td><td></td>
        <td bgcolor="#CCFFCC" class='c'>%v</td><td></td><td></td>
    </tr>
    <tr>
        <td colspan="7" bgcolor="#FFFF00">
            <p>注：1.客户接到财务对账单，三天内如有数据疑问请联系财务专员，财务专员SKYPE:pk.finance<br>&nbsp;&nbsp;2.如无疑问，汇款前请找财务专员索取最新汇款账号，感谢您们的支持，PK娱乐平台团队致敬！<br>&nbsp;&nbsp;3.汇款账号发出后，半小时内没汇好款，要汇款之前请再联系我们确认汇款账号，以免有变动。</p>
        </td>
    </tr>
    <tr>
        <td colspan="4">B套餐（當期負數不可累計，結算日過後清零）</td><td></td><td></td><td></td>
    </tr>
    <tr>
        <td colspan="2">PK平臺體育、彩票</td><td colspan="2" class='c'>0%%</td>
        <td bgcolor="#E26B0A">本月抽成</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
    <tr>
        <td>MG視訊</td><td class='c'>12%%</td>
        <td>MG電子</td><td class='c'>15%%</td>
        <td bgcolor="#E26B0A">%v月预收线路费</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
    <tr>
        <td colspan="3">AG視訊、LEBO視訊、CT視訊、OG視訊</td><td class='c'>12%%</td>
        <td bgcolor="#E26B0A">%v月实收线路费</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
    <tr>
        <td colspan="3">BBIN（體育、視訊、電子、彩票）</td><td class='c'>12%%</td>
        <td bgcolor="#E26B0A">%v月预收线路费</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
    <tr>
        <td>PT电子</td><td class='c'>17%%</td>
        <td>EG电子</td><td class='c'>10%%</td>
        <td bgcolor="#E26B0A">本月应缴金额</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>

</table>`, htmlStyle, startTime, endTime, htmlTbody, moneyCount, oughtCount, linecost, startM, linecost, startM, linecost, endM, linecost, linecost)
	} else if str == "C套餐" {
		for _, v := range data.List {
			htmlTbody += "<tr>"
			if v == data.List[0] {
				htmlTbody += "<td rowspan='" + fmt.Sprintf("<div>%v</div>", len(data.List)+2) + "' class='c'>" + data.SiteName + "<br><br>(C套餐)</td>"
			}
			htmlTbody += "<td>" + v.ProductName + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", v.Win) + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", 0) + "</td>" + //{$negat[$k]}
				"</tr>"
		}
		htmlStr = fmt.Sprintf(`%v
<table cellpadding="0" cellspacing="0" class="table-package">
    <tr>
        <td colspan="7" bgcolor="#FFCC99" class='c'>日期：%v~%v</td>
    </tr>
    <tr>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>網站</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>項目</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>金額</td>
        <td width="10%%" bgcolor="#CCFFCC" class='c'>條件</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>應收</td>
        <td width="10%%" bgcolor="#FFFF00" class='c'>应加应扣</td>
        <td rowspan="2" width="10%%" bgcolor="#CCFFCC" class='c'>備註</td>
    </tr>
    <tr>
        <td bgcolor="#CCFFCC" class='c'>比例</td>
        <td bgcolor="#FFFF00" class='c'>视讯额度</td>
    </tr>
    %v
    <tr>
        <td>其他</td>
        <td></td>
        <td></td>
        <td class='c'>0.00</td><td></td><td></td>
    </tr>
    <tr style="font-weight: bold">
        <td>总计</td>
        <td bgcolor="#CCFFCC" class='c'>%v</td><td></td>
        <td bgcolor="#CCFFCC" class='c'>%v</td><td></td><td></td>
    </tr>
    <tr>
        <td colspan="7" bgcolor="#FFFF00">
            <p>注：1.客户接到财务对账单，三天内如有数据疑问请联系财务专员，财务专员SKYPE:pk.finance<br>&nbsp;&nbsp;2.如无疑问，汇款前请找财务专员索取最新汇款账号，感谢您们的支持，PK娱乐平台团队致敬！<br>&nbsp;&nbsp;3.汇款账号发出后，半小时内没汇好款，要汇款之前请再联系我们确认汇款账号，以免有变动。</p>
        </td>
    </tr>
    <tr>
        <td colspan="4">C套餐（當期負數不可累計，結算日過後清零）</td><td></td><td></td><td></td>
    </tr>
    <tr>
        <td colspan="2">PK平臺體育、彩票</td><td colspan="2" class='c'>6%%</td>
        <td bgcolor="#E26B0A">本月抽成</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
    <tr>
        <td>MG視訊</td><td class='c'>10%%</td>
        <td>MG電子</td><td class='c'>15%%</td>
        <td bgcolor="#E26B0A">%v月预收线路费</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
    <tr>
        <td colspan="3">AG視訊、LEBO視訊、CT視訊、OG視訊</td><td class='c'>10%%</td>
        <td bgcolor="#E26B0A">%v月实收线路费</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
    <tr>
        <td colspan="3">BBIN（體育、視訊、電子、彩票）</td><td class='c'>10%%</td>
        <td bgcolor="#E26B0A">%v月预收线路费</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
    <tr>
        <td>PT电子</td><td class='c'>17%%</td>
        <td>EG电子</td><td class='c'>10%%</td>
        <td bgcolor="#E26B0A">本月应缴金额</td><td bgcolor="#E26B0A" class='c'>%v</td>
        <td></td>
    </tr>
</table>`, htmlStyle, startTime, endTime, htmlTbody, moneyCount, oughtCount, linecost, startM, linecost, startM, linecost, endM, linecost, linecost)
	} else {

		return "暂无模板", err
	}
	return
}

//月份
func MonthNum(Month string) int {
	switch Month {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "Aguest":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	}
	return 0
}
