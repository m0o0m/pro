package global

import (
	"math/rand"
	"time"
	"unsafe"
)

//随机生成10个长度在4-10位之间的字符串
func GetRandStr() []string {
	var accounts []string
	r := rand.New(rand.NewSource(rand.Int63n(time.Now().Unix())))
	tpl := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUvWXYZ0123456789"
	for j := 0; j < 10; j++ {
		rand.Seed(rand.Int63n(time.Now().UnixNano()))
		ra := 4 + rand.Intn(10-4)
		b := make([]byte, ra)
		for i := 0; i < ra; i++ {
			b[i] = tpl[r.Int31n(int32(len(tpl)))]
		}
		account := *(*string)(unsafe.Pointer(&b))
		accounts = append(accounts, account)
	}
	return accounts
}
