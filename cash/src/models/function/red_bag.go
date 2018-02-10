package function

import (
	"global"
	"models/input"
	"models/schema"
	"strconv"
)

type RedBagBean struct{}

//红包数据补全（强行修改某会员的打码或存款数据让其能抢红包）
func (*RedBagBean) RedBagSet(this *input.RedBagData) (code int64, key string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	redPacket := new(schema.RedPacketSet)
	var has bool
	var redType string
	if this.Type == 2 {
		redType = "bet" + this.Account
	} else {

		has, err = sess.Where("site_id=? and site_index_id=? and account=?", this.SiteId, this.SiteIndexId,
			this.Account).Get(member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return code, key, err
		}
		if !has {
			code = 30138 //会员账号不存在!
			return
		}
		redType = "sum" + strconv.FormatInt(member.Id, 10)
	}
	has, err = sess.Where("site_id=? and site_index_id=?", this.SiteId, this.SiteIndexId).Get(redPacket)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, key, err
	}
	if !has {
		code = 10142 //当前无红包!
		return
	}
	key = "red_bag" + this.SiteId + "_" + this.SiteIndexId + "_" + strconv.FormatInt(redPacket.Id, 10) + redType
	global.GetRedis().Set(key, this.Value, 172800) //2天
	return
}
