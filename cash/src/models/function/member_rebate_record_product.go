package function

import (
	"github.com/go-xorm/xorm"
	"global"
	"models/back"
	"models/schema"
)

type MemberRebateRecordProductBean struct {
}

//根据返佣记录id查询返佣详情对应商品金额
func (m *MemberRebateRecordProductBean) GetRecordProductListByRecordIds(recordIds []int) (
	[]back.RebateRecordProduct, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	rebateRecordProduct := new(schema.MemberRebateRecordProduct)
	var rebateRecordProducts []back.RebateRecordProduct
	product := new(schema.Product)
	err := sess.Table(rebateRecordProduct.TableName()).
		Join("INNER", product.TableName(), rebateRecordProduct.TableName()+".product_id = "+product.TableName()+".id").
		Select("record_id,product_id,product_name,product_bet,money").
		In(rebateRecordProduct.TableName()+".record_id", recordIds).
		Find(&rebateRecordProducts)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return rebateRecordProducts, err
	}
	return rebateRecordProducts, err
}

//根据返佣详情id,删除所有对应的商品记录
func (m *MemberRebateRecordBean) DelRebateRecordProductByRecordIds(recordIds []int64, sessArgs ...*xorm.Session) (int64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	rebateRecordProduct := new(schema.MemberRebateRecordProduct)
	num, err := sess.Table(rebateRecordProduct.TableName()).
		In("record_id", recordIds).
		Delete(rebateRecordProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return num, err
	}
	return num, err
}
