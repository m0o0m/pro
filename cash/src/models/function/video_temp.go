package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type VideoTempBean struct{}

//获取视讯模板列表
func (*VideoTempBean) GetVideoList() (backData []back.VideoStyle, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	videoStyle := new(schema.SiteInfoVideoStyle)
	sess.Where("id>?", 0)
	sess.Select("`id`, `name`, `pid`, `aid`, `style`, `remark`, `status`")
	err = sess.Table(videoStyle.TableName()).Where("pid=?", 0).Find(&backData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return backData, err
	}
	return backData, err
}

//新增视讯模板
func (*VideoTempBean) AddVideo(this *input.VideoAdd) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	videoStyle := new(schema.SiteInfoVideoStyle)
	videoStyle.Name = this.Name
	videoStyle.Status = this.Status
	videoStyle.Pid = this.Pid
	videoStyle.Aid = this.Style
	data, err = sess.Insert(videoStyle)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//修改视讯模板
func (*VideoTempBean) UpdateVideo(this *input.VideoUpdate) (data int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	videoStyle := new(schema.SiteInfoVideoStyle)
	videoStyle.Name = this.Name
	videoStyle.Aid = this.Style
	videoStyle.Pid = this.Pid
	videoStyle.Status = this.Status
	data, err = sess.Where("id=?", this.Id).Update(videoStyle)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
