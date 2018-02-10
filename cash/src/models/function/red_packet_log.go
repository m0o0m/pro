package function

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"global"
	"math/rand"
	"models/schema"
	"time"
)

type RedPacketLogBean struct {
}

//将生成的红包添加到数据库存放
func (*RedPacketLogBean) AddRedPacket(redPackets []*schema.RedPacketLog, sessArgs ...*xorm.Session) (num int64, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	return sess.InsertMulti(&redPackets)
}

//分配红包
func (*RedPacketLogBean) GenerateRedPacket(sumMoney float64, num int64, minMoney float64) (amounts []float64, code int64) {
	other, code := compute(int(sumMoney*100), int(num), int(minMoney*100))
	if code > 0 {
		return nil, code
	}
	fmt.Println(other)
	var result = make([]float64, len(other))
	for i := 0; i < len(other)/2+1; i++ {
		result[i], result[int(num)-1-i] = global.FloatReserve2(float64(other[i])/100), global.FloatReserve2(float64(other[int(num)-1-i])/100)
	}
	return result, 0
}

//传入总金额，红包个数，每个红包最低金额，传入单位是分，返回单位是分
func compute(sumMoney int, num int, minMoney int) ([]int, int64) {
	if (minMoney)*num > sumMoney {
		return nil, 71022
	}
	other := make([]int, num)
	//分配
	distribution := sumMoney - (minMoney * num)
	if distribution < 0 {
		//global.GlobalLogger.Error("(num:%d * minMoney:%d) > sumMoney:%d", num, minMoney, sumMoney)
		return other, 71022
	} else if distribution == 0 {
		for index := range other {
			other[index] = minMoney
		}
		return other, 0
	}
	fmt.Println("剩余", float64(distribution)/100, num, float64(minMoney)/100)
	currentTime := time.Now().UnixNano()
	r := rand.New(rand.NewSource(currentTime))
	addReadPack(distribution, &other, num, r)
	for index, _ := range other {
		other[index] += minMoney
	}
	return other, 0
}

func addReadPack(remainAmount int, amounts *[]int, num int, r *rand.Rand) []int {
	if num <= 1 {
		(*amounts)[num-1] = remainAmount
		return (*amounts)
	}
	maxMoney := remainAmount / num * 2
	ranMoney := r.Intn(maxMoney + 1)
	(*amounts)[num-1] = ranMoney
	return addReadPack(remainAmount-ranMoney, amounts, num-1, r)
}
