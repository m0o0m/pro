package function

import (
	"global"
	"models/back"
	"models/schema"
)

type PkGames struct{}

//查询PK彩票游戏类型
func (*PkGames) PkGames() (PkGames []back.PkLottery, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.PkGames)
	err = sess.Table(infoLogoSchema.TableName()).
		Where("state = ?", 1).
		Select("name,type,l_type").
		Find(&PkGames)
	if err != nil {
		return
	}
	return
}
