package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type GameSetBean struct{}

//游戏类型列表
func (*GameSetBean) GameTypeList() (data []back.VdGameTypeList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	module := new(schema.MgGameType)
	err = sess.Table(module.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//添加游戏类型
func (*GameSetBean) GameTypeAdd(this *input.VdGameType) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	module := new(schema.MgGameType)
	module.Type = this.Type
	count, err = sess.Table(module.TableName()).InsertOne(module)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//更新游戏类型
func (*GameSetBean) GameTypeUpdate(this *input.VdGameTypeUpdate) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	module := new(schema.MgGameType)
	module.Type = this.NewType
	count, err = sess.Table(module.TableName()).Where("type=?", this.OldType).Cols("type").Update(module)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//游戏类型删除
func (*GameSetBean) DelGameType(this *input.VdGameType) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	module := new(schema.MgGameType)
	module.Type = this.Type
	count, err = sess.Table(module.TableName()).Where("type=?", this.Type).Delete(module)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
