package thirdParty

import (
	"encoding/json"
	"errors"
	"fmt"
	"global"
	"io/ioutil"
	"models/input"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type PayBean struct {
}

/*type ceShiData struct {
	clientUserId int64  //是	客户ID 询问客服	1
	clientName   string //是	客户名称 询问客服	pkClient
	clientSecret string //是	客户授权证书 询问客服	h1qN7RYH9xpvugZhaFu5Inmdk6bJyIopJrsbCAmj
	order        string //是	订单号	170826162633758040
	agentId      string //是	代理线(站点ID)	t
	agentNum     string //是	自代理线(子站ID)	a
	amount       string //是	金额(元)	100
	payWay       int64  //是	支付方式	1为网银,2为微信，3为支付宝，4为QQ钱包，5为财付通,6银联，7京东钱包，8百度钱包
	merchantId   int64  //是	商户表ID	3
	businessNum  string //是	商户号	3112374124
	bank         string //是	银行特殊编码	可为空字符串，网银提交时填写银行编码
	isApp        int64  //否	是否跳转APP	移动端若跳转APP则传1，否则不传或者传0
	cardMoney    string //否	点卡面额	若为点卡支付则必传，否则可不传，若为空不参与签名
	cardNumber   string //否	点卡编号	若为点卡支付则必传，否则可不传，若为空不参与签名
	CardPwd      string //否	点卡密码	若为点卡支付则必传，否则可不传，若为空不参与签名
	sign         string //是	签名	4925754189ed2bad3bbbef290454704e
	tokSign      string //是	订单token	b81da644ae506caa1142abfcd9e2e121

	TransferType string  //IN OUT
	Credit       float64 //额度不能为负数
	TradeNo      string  //订单号
}*/

//传入 金额，订单号，支付方式
func (*PayBean) ThirdPayTest(this *input.ThirdPayData) error {
	//这里添加post的body内容
	data := make(url.Values)
	data["clientUserId"] = []string{strconv.FormatInt(this.ClientUserId, 10)}
	data["clientName"] = []string{this.ClientName}
	data["clientSecret"] = []string{this.ClientSecret}
	data["order"] = []string{this.Order}
	data["agentId"] = []string{this.SiteId}
	data["agentNum"] = []string{this.SiteIndexId}
	fmt.Println(this.Order)
	res, err := http.PostForm("http://olpay.pk1358.com/api/v1/token", data) //获取token
	if err != nil {
		return err
	}
	fmt.Println(res)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	f := struct {
		Status bool   `json:"status"`
		Token  string `json:"token"`
	}{}
	fmt.Println(string(b))
	err = json.Unmarshal(b, &f)
	if err != nil {
		return err
	}
	if !f.Status {
		return errors.New("获取token失败")
	}
	fmt.Println("-------------------------" + f.Token)

	data["amount"] = []string{strconv.FormatFloat(this.Amount, 'f', -1, 64)}
	data["bank"] = []string{this.PaidCode}
	data["businessNum"] = []string{this.BusinessNum}
	data["payWay"] = []string{strconv.Itoa(this.PaidWay)}
	data["merchantId"] = []string{strconv.FormatInt(this.MerchatId, 10)} //strconv.FormatInt(merchatId,10)
	data["tokSign"] = []string{f.Token}
	if this.PaidWay == 10 { //点卡
		data["cardMoney"] = []string{strconv.FormatFloat(this.CardMoney, 'f', -1, 64)}
	} else {
		data["cardMoney"] = []string{""}
	}
	data["cardNumber"] = []string{this.CardNumber}
	data["cardPwd"] = []string{this.CardPwd}
	str := "&&&agentId=#" + data["agentId"][0] + "#&&&&&&agentNum=#" + data["agentNum"][0] + "#&&&&&&amount=#" + data["amount"][0] + "#&&&&&&bank=#" + data["bank"][0] + "#&&&&&&businessNum=#" + data["businessNum"][0] + "#&&&&&&cardMoney=#" + data["cardMoney"][0] + "#&&&&&&cardNumber=#" + data["cardNumber"][0] + "#&&&&&&cardPwd=#" + data["cardPwd"][0] + "#&&&&&&clientSecret=#" + data["clientSecret"][0] + "#&&&&&&merchantId=#" + data["merchantId"][0] + "#&&&&&&order=#" + data["order"][0] + "#&&&&&&payWay=#" + data["payWay"][0] + "#&&&"
	for key, value := range data { //去除空值后签名
		if value[0] == "" {
			strings.Replace(str, "&&&"+key+"=##&&&", "", 1)
		}
	}
	fmt.Println(str)
	data["sign"] = []string{global.Md5(global.Md5(str))}

	//senddata:="agentId="+data["agentId"][0]+"&agentNum="+data["agentNum"][0]+"&amount="+data["amount"][0]+"&bank="+data["bank"][0]+"&businessNum="+data["businessNum"][0]+"&merchantId="+data["merchantId"][0]+"&order="+data["order"][0]+"&payWay="+data["payWay"][0]+"&sign="+global.Md5(global.Md5(str))+"&clientUserId="+data["clientUserId"][0]+"&tokSign="+data["tokSign"][0]
	//delete(data, "clientUserId")
	delete(data, "clientName")
	delete(data, "clientSecret")
	fmt.Println(data)
	res, err = http.PostForm("http://olpay.pk1358.com/api/v1/buy", data)
	if err != nil {
		return err
	}
	c, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	//ctx.Set("token", string(c))
	fmt.Println(this.PayRedisKey)
	err = global.GetRedis().Set(this.PayRedisKey, string(c), 0).Err()
	if err != nil {
		return err
	}
	fmt.Println(string(c))
	/*g := struct {
		Status bool   `json:"status"`
		Code   int64  `json:"code"`
		Data   string `json:"data"`
	}{}
	fmt.Println("888888888")
	err = json.Unmarshal(c, &g)
	if err != nil {
		//说明已经请求到了成功地址，未发生错误
	}
	//把post表单发送给目标服务器
	fmt.Println("9999999999")
	fmt.Println(res, "-------------------------"+g.Data)*/
	fmt.Println("post send success")
	defer res.Body.Close()
	return err
}
