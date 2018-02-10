package global

import (
	"math/rand"
	"regexp"
	"time"
)

//随机数生成要用到的
const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

//domain check
func DomainCheck(domain string) bool {
	var match bool
	//支持以http://或者https://开头并且域名中间有/的情况
	IsLine := "^((http://)|(https://))?([a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,6}(/)"
	//支持以http://或者https://开头并且域名中间没有/的情况
	NotLine := "^((http://)|(https://))?([a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,6}"
	match, _ = regexp.MatchString(IsLine, domain)
	if !match {
		match, _ = regexp.MatchString(NotLine, domain)
	}
	return match
}

//随机生成
func RandStringBytesMaskImprSrc(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

//邮箱校验
func CheckEmail(email string) (flag bool) {
	var emailJS = "[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"
	flag, _ = regexp.MatchString(emailJS, email)
	if !flag {
		flag = false
		return
	}
	return
}

//qq校验
func Checkqq(qq string) (flag bool) {
	qqJS := "[1-9][0-9]{4,}"
	flag, _ = regexp.MatchString(qqJS, qq)
	if !flag {
		flag = false
		return
	}
	return
}

//手机号校验，或者座机号校验
func CheckPhoneNumber(phone string) (flag bool) {
	phoneJS := "^1([0-9][0-9]|14[57]|5[^4])\\d{8}$"
	landLineJS := "^([0-9]{3}-[0-9]{4,})"
	flag, _ = regexp.MatchString(phoneJS, phone)
	if !flag {
		flag, _ = regexp.MatchString(landLineJS, phone)
		if !flag {
			flag = false
			return
		}
	}
	return
}

//身份证校验
func CheckIdentity(identity string) (flag bool) {
	identityJS := "^([0-9]){17}([0-9]|X)$"
	flag, _ = regexp.MatchString(identityJS, identity)
	if !flag {
		flag = false
		return
	}
	return
}

//银行卡校验
func CheckCardNumber(card string) (flag bool) {
	cardJS := "^([0-9]){16,19}$"
	flag, _ = regexp.MatchString(cardJS, card)
	if !flag {
		flag = false
		return
	}
	return
}

//IP地址校验
func CheckIp(ip string) (flag bool) {
	ipPattern := regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
	flag = ipPattern.MatchString(ip)
	if !flag {
		flag = false
		return
	}
	return
}
