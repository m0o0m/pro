package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type GameWhitelistBean struct{}

//视讯白名单列表
func (*GameWhitelistBean) GetGameWhiteList(this *input.GameWhiteList, listparam *global.ListParams) (data []back.GameWhiteList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Ip != "" {
		sess.Where("ip = ?", this.Ip)
	}
	game := new(schema.GameWhiteList)
	//获得分页记录
	listparam.Make(sess)
	err = sess.Table(game.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err = sess.Table(game.TableName()).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//添加视讯白名单IP
func (*GameWhitelistBean) GameWhiteAdd(this schema.GameWhiteList) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.GameWhiteList)
	count, err = sess.Table(admin.TableName()).InsertOne(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改视讯白名单ip
func (*GameWhitelistBean) UpdateInfo(this *input.GameWhiteEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.GameWhiteList)
	admin.Ip = this.Ip
	admin.Remarks = this.Remarks
	count, err = sess.Where("id = ?", this.Id).Cols("ip,remarks").Update(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除视讯白名单ip
func (*GameWhitelistBean) DelGameWhite(Id int) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	admin := new(schema.GameWhiteList)
	count, err = sess.Where("id = ?", Id).Delete(admin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
