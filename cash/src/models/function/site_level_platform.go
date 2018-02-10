package function

import (
	"errors"
	"github.com/go-xorm/xorm"
	"global"
	"models/schema"
)

type LevelPlatformBean struct {
}

//通过层级Id和平台id获取占成比
func (m *LevelPlatformBean) GetProportionByLevelIdPlatformId(levelId int64, platformId int64, sessArgs ...*xorm.Session) (proportion float64, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
		defer sess.Close()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	if levelId == 0 {
		return proportion, errors.New("Not found level_id")
	}
	b, err := sess.Table(new(schema.SiteLevelPlatform).TableName()).
		Select("proportion").
		Where("level_id = ?", levelId).
		Where("platform_id = ?", platformId).
		Get(&proportion)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return proportion, err
	}
	if !(b) {
		return proportion, errors.New("Not found SiteLevelPlatform")
	}
	return proportion, nil
}
