package function

import (
	"errors"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type RedPacketStyleBean struct {
}

//添加或则修改红包皮肤
func (*RedPacketStyleBean) AddOrUpdate(add *input.RedPacketStyleAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketStyleSchema := new(schema.RedPacketStyle)
	redPacketStyleSchema.SiteId = add.SiteId
	redPacketStyleSchema.Status = 1
	redPacketStyleSchema.Name = add.Name
	redPacketStyleSchema.BgPic = add.BgPic
	redPacketStyleSchema.ClickPic = add.ClickPic
	redPacketStyleSchema.ClickCss = add.ClickCss
	redPacketStyleSchema.TimeCss = add.TimeCss
	if add.Id > 0 {
		num, err := sess.Where("id=?", add.Id).
			Cols("site_id,status,name,bg_pic,click_pic,click_css,time_css").
			Update(redPacketStyleSchema)
		return num, err
	}
	num, err := sess.InsertOne(redPacketStyleSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return num, err
	}
	return num, err
}

//查询红包皮肤列表
func (*RedPacketStyleBean) List(siteId string) ([]*back.RedPacketStyleList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var redPacketStyles []*back.RedPacketStyleList
	redPacketStyleSchema := new(schema.RedPacketStyle)
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	sess.Where("status = ?", 1)
	err := sess.Table(redPacketStyleSchema.TableName()).Find(&redPacketStyles)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return redPacketStyles, err
	}
	return redPacketStyles, err
}

//查询红包皮肤下拉框
func (*RedPacketStyleBean) RedStyleDrop(siteId string) ([]back.RedStyleDrop, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketStyleSchema := new(schema.RedPacketStyle)
	var data []back.RedStyleDrop
	sess.Where("site_id = ?", siteId)
	sess.Where("status = ?", 1)
	err := sess.Table(redPacketStyleSchema.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询红包皮肤详情
func (*RedPacketStyleBean) Details(id int64) (back.RedPacketStyle, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var redPacketStyle back.RedPacketStyle
	redPacketStyleSchema := new(schema.RedPacketStyle)
	sess.Where("id = ?", id)
	sess.Where("status = ?", 1)
	b, err := sess.Table(redPacketStyleSchema.TableName()).Get(&redPacketStyle)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return redPacketStyle, err
	}
	if !b {
		err = errors.New("read null")
	}
	return redPacketStyle, err
}

//删除红皮肤
func (*RedPacketStyleBean) Del(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketStyleSchema := new(schema.RedPacketStyle)
	redPacketStyleSchema.Status = 2
	sess.Where("id = ?", id)
	sess.Where("status = ?", 1)
	count, err := sess.Table(redPacketStyleSchema.TableName()).Update(redPacketStyleSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询红包皮肤图片
func (*RedPacketStyleBean) RedStyleDropPicture(id int64) (*back.RedStyleDropPicture, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	redPacketStyleSchema := new(schema.RedPacketStyle)
	data := new(back.RedStyleDropPicture)
	has, err := sess.Table(redPacketStyleSchema.TableName()).Where("id=?", id).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}
