package global

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//float类型保留2位小数(四舍五入) ,总感觉这种方式不好,有更好的办法请替换该方法
func FloatReserve2(src float64) float64 {
	temp, _ := strconv.ParseFloat(fmt.Sprintf("%0.2f", src), 64)
	return temp
}

//生成随机字符串
func GetRandomString(length int, rs ...*rand.Rand) string {
	var r *rand.Rand
	switch len(rs) {
	case 1:
		r = rs[0]
	case 0:
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	default:
		panic("not found rand.Rand")
	}
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成随机伪账号
func GetRandomAccount() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	start := GetRandomString(3, r)
	end := GetRandomString(3, r)
	return start + "****" + end
}

//判断小数点后面是不是两位
func KeepTwoFloat(money float64) int64 {
	//将float64转为string,再做分割，再判断
	ss := strings.Split(strconv.FormatFloat(money, 'f', -1, 64), ".")
	if len(ss[1]) > 2 {
		return 30251
	}
	return 0
}

//银行卡号展示后四位
func BankStr(card string) (cards string) {
	var sss string
	sss = "**** **** **** " + card[len(card)-4:]
	return sss
}
