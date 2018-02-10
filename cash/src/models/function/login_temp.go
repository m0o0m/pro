package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type LoginTempBean struct{}

//获取注册模板列表
func (*LoginTempBean) GetLoginTempList() (backData []back.LoginTemp, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	videoStyle := new(schema.InfoRegStyle)
	sess.Where("id>?", 0)
	sess.Select("`id`, `title`, `type`, `content`, `color`, `state`")
	err = sess.Table(videoStyle.TableName()).Find(&backData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, err
	}
	return backData, err
}

//获取注册模板详情
func (*LoginTempBean) GetOneLoginRegOneDetail(this *input.LoginRegListDetailIn) (*back.LoginRegDetailBack, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	reg := new(schema.InfoRegStyle)
	info := new(back.LoginRegDetailBack)
	has, err := sess.Table(reg.TableName()).
		Where("id=?", this.Id).Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err
	}
	return info, has, err
}

//新增注册模板
func (*LoginTempBean) AddLogin(this *input.LoginAdd) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	videoStyle := new(schema.InfoRegStyle)
	videoStyle.Title = this.Title
	videoStyle.Content = this.Content
	videoStyle.Color = this.Color
	videoStyle.Type = this.Type
	videoStyle.State = 1
	data, err = sess.Table(videoStyle.TableName()).Insert(videoStyle)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//修改注册模板
func (*LoginTempBean) UpdateLogin(this *input.LoginUpdate) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	videoStyle := new(schema.InfoRegStyle)
	videoStyle.Title = this.Title
	videoStyle.Content = this.Content
	videoStyle.Color = this.Color
	videoStyle.Type = this.Type
	data, err = sess.Table(videoStyle.TableName()).Where("id=?", this.Id).Update(videoStyle)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//修改注册模板状态
func (*LoginTempBean) UpdateStatus(this *input.LoginStatus) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	videoStyle := new(schema.InfoRegStyle)
	switch this.Status {
	case 1:
		videoStyle.State = 2
	case 2:
		videoStyle.State = 1
	}
	data, err = sess.Table(videoStyle.TableName()).Where("id=?", this.Id).Update(videoStyle)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
