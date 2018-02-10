package function

import (
	"global"
	"models/schema"
)

type EgGames struct{}

//查询EG彩票游戏类型
func (*EgGames) EgGames() (EgGames []schema.EgGames, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.EgGames)
	err = sess.Table(infoLogoSchema.TableName()).
		Where("eg_state = ?", 1).
		OrderBy("eg_sort desc,eg_lx_type desc").
		Find(&EgGames)
	if err != nil {
		return
	}
	return
}
