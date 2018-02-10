//[控制器] [平台]站点报表管理
package report

import (
	//"bytes"
	"bufio"

	"controllers"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"os"
	"strconv"
	"strings"
	"time"
)

//套餐管理
type ReportController struct {
	controllers.BaseController
}

//报表查询页，显示所需查询条件
func (c *ReportController) GetReportTerm(ctx echo.Context) error {
	sdata := make(map[string]interface{})
	sdata["porductList"] = []back.ProductlistRep{}
	sdata["siteIdList"] = []back.InSiteList{}
	sdata["combo"] = []back.GetList{}
	// 获取商品数据
	data, err := productBean.ProductList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取游戏类型
	sdata["porductList"] = data
	sdata["siteIdList"], err = siteOperateBean.SiteList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取套餐列表
	sdata["combo"], err = comboBeen.GetComboDrop(new(schema.Combo))
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//站点报表查询
func (c *ReportController) GetSiteReportList(ctx echo.Context) error {
	getCenterList := new(input.GetDataCenterList)
	code := global.ValidRequest(getCenterList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)

	times := new(global.Times)
	if len(getCenterList.StartTime) != 0 {
		times.StartTime, code = global.FormatTime2Timestamp2(getCenterList.StartTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.StartTime = global.GetTimeStart(global.GetCurrentTime())
	}
	if len(getCenterList.EndTime) != 0 {
		times.EndTime, code = global.FormatTime2Timestamp2(getCenterList.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.EndTime = global.GetCurrentTime()
	}

	list, count, err := reportFormBean.GetSiteReportList(getCenterList, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//报表导出
func (c *ReportController) GetSiteReportExport(ctx echo.Context) error {
	reportExports := new(input.ReportExports)
	code := global.ValidRequest(reportExports, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	TableBtmlstr, _ := base64.StdEncoding.DecodeString(reportExports.TableBtml)

	//写入
	f, err := os.Create("bills.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	file := os.NewFile(0, "asdf")
	w := bufio.NewWriter(file)
	w.Write(TableBtmlstr)
	w.Flush()

	return ctx.Stream(200, "application/vnd.ms-excel", f)

}

//报表统计数据查询
func (c *ReportController) GetSiteReportAccount(ctx echo.Context) error {
	reportExport := new(input.ReportBills)
	code := global.ValidRequest(reportExport, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if len(reportExport.SiteId) == 0 || len(reportExport.VType) == 0 {
		return ctx.JSON(60201, global.ReplyError(code, ctx))
	}
	//账单数据
	sdata := []back.ReportInfo{}
	//查询账单商品名和商品id
	productList, _ := productBean.ProductList()

	//查询站点套餐信息
	indexList, err := siteOperateBean.SiteCombo(reportExport.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//将商品名和商品v_type加进套餐信息中
	for k, v := range indexList {
		info := back.ReportInfo{}
		info.SiteId = v.SiteId
		info.SiteIndexId = v.SiteIndexId
		info.SiteName = v.SiteName
		info.ComboName = v.ComboName
		for _, v1 := range reportExport.VType {
			vTypeInfo := back.ReportInfoList{}
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
		}
		sdata = append(sdata, info)
	}

	//获取导出数据
	data, err := reportBean.ReportExport(reportExport)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//获取上月报表负数
	//获取月期数
	loc, _ := time.LoadLocation("Local") //时区
	t, err := time.ParseInLocation("2006-01-02", reportExport.StartTime, loc)

	years := t.AddDate(0, -1, 0).Year()
	months := MonthNum(t.AddDate(0, -1, 0).Month().String())
	qishu := ""
	if months < 10 {
		qishu = strconv.Itoa(years) + "-0" + strconv.Itoa(months)
	} else {
		qishu = strconv.Itoa(years) + "-" + strconv.Itoa(months)
	}
	negativeList, _ := siteOperateBean.SiteNegativeList(reportExport.SiteId, qishu)

	if err != nil {
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
	}

	for k, v := range indexList {
		for k1, v1 := range v.List {
			for _, v2 := range negativeList[v.SiteId].Products {

				if v1.ProductId == v2.ProductId {
					indexList[k].List[k1].Negative = v2.ReportWin
				}
			}
		}
	}

	for k, v := range sdata {
		for _, v1 := range indexList {
			if v.SiteId == v1.SiteId {

				sdata[k].Html, _ = ExcelHtml(v1, reportExport.StartTime, reportExport.EndTime)
			}
		}
	}
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//报表账单查询
func (c *ReportController) GetSiteReportBill(ctx echo.Context) error {
	billList := new(input.BillList)
	code := global.ValidRequest(billList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询账单
	data, err := reportBean.GetSiteReportBill(billList)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	return ctx.JSON(200, global.ReplyCollections("data", data))
}

//报表账单批量下发
func (c *ReportController) GetSiteReportBillBatch(ctx echo.Context) error {
	billList := new(input.BillListBatch)
	code := global.ValidRequest(billList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := reportBean.ReportBillBatch(billList)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//报表账单添加
func (c *ReportController) PostSiteBillAdd(ctx echo.Context) error {
	billsAdd := new(input.BillsAdd)
	code := global.ValidRequest(billsAdd, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if len(billsAdd.SiteId) == 0 {
		return ctx.JSON(60101, global.ReplyError(code, ctx))
	}
	//查询账单是否已经存在
	billList := new(input.BillList)
	billList.SiteId = billsAdd.SiteId
	billList.Year = billsAdd.Year
	billList.Qishu = billsAdd.Qishu
	data, err := reportBean.GetSiteReportBill(billList)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(data) != 0 {
		return ctx.JSON(200, global.ReplyError(60210, ctx))
	}

	//获取日期
	startDate, endDate := c.PeriodDate(billsAdd.Year, billsAdd.Qishu)
	//添加
	//dat:=new(back.ReportExport)
	//dat.SiteId = billsAdd.SiteId
	//dat.ComboName
	//billsAdd.ReportData = ExcelHtml()
	count, err := reportBean.BillsAdd(billsAdd, startDate, endDate)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30096, ctx))
	}
	return ctx.NoContent(204)
}

//报表账单修改
func (c *ReportController) PutSiteBillUpdate(ctx echo.Context) error {
	siteBillUpdate := new(input.SiteBillUpdate)
	code := global.ValidRequest(siteBillUpdate, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	if siteBillUpdate.Id == 0 {
		return ctx.JSON(30041, global.ReplyError(code, ctx))
	}
	if len(siteBillUpdate.SiteId) == 0 {
		return ctx.JSON(60207, global.ReplyError(code, ctx))
	}
	if len(siteBillUpdate.Qishu) == 0 {
		return ctx.JSON(60208, global.ReplyError(code, ctx))
	}

	code, count, err := reportBean.PutSiteBillUpdate(siteBillUpdate)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(60209, ctx))
	}
	return ctx.NoContent(204)
}

//报表账单删除
func (c *ReportController) DelSiteBill(ctx echo.Context) error {
	billList := new(input.BillList)
	code := global.ValidRequest(billList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if len(billList.SiteId) == 0 {
		return ctx.JSON(60207, global.ReplyError(code, ctx))
	}

	code, count, err := reportBean.DelSiteBill(billList)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(60211, ctx))
	}
	return ctx.NoContent(204)
}

//根据期数获取日期
func (*ReportController) PeriodDate(Year, QishuId string) (startDate, endDate int64) {
	dataStr := Year + "-" + QishuId + "-01"

	loc, _ := time.LoadLocation("Local") //时区
	t, _ := time.ParseInLocation("2006-01-02", dataStr, loc)

	num, _ := withoutMonday(t.Weekday().String())
	startDate = t.AddDate(0, 0, num).Unix()
	_, eNum := withoutMonday(t.AddDate(0, 1, 0).Weekday().String())
	endDate = t.AddDate(0, 1, eNum).Unix()
	return
}

//距离星期一差几天
func withoutMonday(week string) (num, eNum int) {
	switch week {
	case "Sunday":
		return 1, 6
	case "Monday":
		return 0, 0
	case "Tuesday":
		return 6, 1
	case "Wednesday":
		return 5, 2
	case "Thursday":
		return 4, 3
	case "Friday":
		return 3, 4
	case "Saturday":
		return 2, 5
	}
	return
}

func ExcelHtml(data back.ReportExport, startTime, endTime string) (htmlStr string, err error) {

	//将套餐转成大写
	str := strings.ToUpper(data.ComboName)

	//获取excel表格中的月份
	loc, _ := time.LoadLocation("Local") //时区
	t, err := time.ParseInLocation("2006-01-02", startTime, loc)
	startM := MonthNum(t.Month().String())

	endM := MonthNum(t.AddDate(0, 1, 0).Month().String())

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
	for k, v := range data.List {

		data.List[k].Ought = v.Win + v.Negative
		oughtCount = oughtCount + data.List[k].Ought
		moneyCount = moneyCount + v.Win
		negatCount = negatCount + v.Negative
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
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", v.Negative) + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v%%</div>", v.Proportion) + "</td><td></td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", v.Ought) + "</td><td></td><td></td>" +
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
				"<td class='c'>" + fmt.Sprintf("<div>%v%%</div>", v.Proportion) + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", v.Ought) + "</td><td></td><td></td>" +
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
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>%%", v.Proportion) + "</td>" +
				"<td class='c'>" + fmt.Sprintf("<div>%v</div>", v.Ought) + "</td><td></td><td></td>" +
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
