package function

import (
	"global"
	"models/back"
	"models/schema"
)

type ApiClientsBean struct{}

//查询一条三方对接验证
func (*ApiClientsBean) GetOneApiClients(siteId string) (data *back.GetOneApiClients, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("site_id=?", siteId)
	sess.Where("revoked=?", 1)
	data = new(back.GetOneApiClients)
	has, err = sess.Table(new(schema.ApiClients).TableName()).Get(data)
	return
}
