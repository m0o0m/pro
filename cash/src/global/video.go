package global

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

//额度转换
type Res struct {
	Status      int8    `json:"status"`       //状态
	ConfirmTime int64   `json:"confirm_time"` //确认时间
	Balance     float64 `json:"balance"`      //视讯余额
}

//模拟调用视讯接口(额度转换)
func Video() (res Res, err error) {
	//resp, err := http.Get("http://127.0.0.1:10086/video")
	//if err != nil {
	//	GlobalLogger.Error("error:", err.Error())
	//	return
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	GlobalLogger.Error("error:", err.Error())
	//	return
	//}
	//err = json.Unmarshal(body, &res)
	res = Res{
		Status:      20,
		ConfirmTime: time.Now().Unix(),
		Balance:     float64(rand.Intn(1000)),
	}
	return
}

//额度转换-单个平台余额刷新
type ResBalance struct {
	Balance float64 `json:"balance"` //余额
}

//模拟调用视讯接口(单个平台余额刷新）
func PlatformBalanceRefresh() (resBalance ResBalance, err error) {
	resp, err := http.Get("http://127.0.0.1:10086/platform/balance/refresh")
	if err != nil {
		GlobalLogger.Error("error:", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		GlobalLogger.Error("error:", err.Error())
		return
	}
	err = json.Unmarshal(body, &resBalance)
	return
}
