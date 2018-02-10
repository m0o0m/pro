package function

import (
	"global"
	"models/schema"
)

type CsGames struct{}

//查询cs彩票游戏类型
func (*CsGames) CsGames() (CsGames []schema.CsGames, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.CsGames)
	err = sess.Table(infoLogoSchema.TableName()).
		Where("cs_state = ?", 1).
		OrderBy("cs_sort desc,cs_lx_type desc").
		Find(&CsGames)
	if err != nil {
		return
	}
	return
}
