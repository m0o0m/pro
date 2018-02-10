package function

import (
	"errors"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type MemberMessageBean struct {
}

//修改消息  (设为已读)
func (m *MemberMessageBean) UpdateMessage(ids ...int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	messageSchema := new(schema.MemberMessage)
	messageSchema.State = 2
	sess.Cols("state").
		Where("delete_time = ?", 0).
		Where("state = ?", 1)

	if len(ids) > 0 {
		sess.Where("id = ? ", ids[0])
	}
	count, err := sess.Update(messageSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除消息
func (m *MemberMessageBean) DelMessage(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	messageSchema := new(schema.MemberMessage)
	messageSchema.ID = id
	messageSchema.DeleteTime = time.Now().Unix()
	count, err := sess.Update(messageSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//消息详情
func (m *MemberMessageBean) MessageDetails(id int64) (back.MemberMessage, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var message back.MemberMessage
	messageSchema := new(schema.MemberMessage)
	has, err := sess.Table(messageSchema.TableName()).
		Where("id = ?", id).
		Where("delete_time = ?", 0).
		Get(&message)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return message, err
	}
	if !has {
		err = errors.New("get message failure")
	}
	return message, err
}

//消息列表 1,站点,2系统
func (m *MemberMessageBean) MessageList(siteId, siteIndexId string, id int64, listParams *global.ListParams) (
	messageList []back.MemberMessageList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	messageSchema := new(schema.MemberMessage)

	if id != 0 {
		sess.Where("member_id = ?", id)
	}
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	if siteIndexId != "" {
		sess.Where("site_index_id = ?", siteIndexId)
	}
	sess.Where("delete_time = ?", 0).Desc("create_time")
	listParams.Make(sess)
	conds := sess.Conds()
	err = sess.Table(messageSchema.TableName()).
		Find(&messageList)
	if err != nil {
		return
	}
	count, err = sess.Table(messageSchema.TableName()).Where(conds).Count()
	return
}

//个人消息数量
func (m *MemberMessageBean) MemMessageCount(siteId, siteIndexId string, id int64, times *global.Times) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	messageSchema := new(schema.MemberMessage)
	if times != nil {
		times.Make("create_time", sess)
	}
	count, err = sess.Table(messageSchema.TableName()).
		Where("delete_time = ?", 0).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Where("member_id = ?", id).
		Where("mtype = ?", 1).
		Where("state = ?", 1).
		Count()
	return
}

//获取个人信息列表或者游戏公告列表
func (m *MemberMessageBean) WapList(this *input.WapMemberMessageList, listParams *global.ListParams) ([]*back.WapMemberMessageList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	message := new(schema.MemberMessage)
	listParams.Make(sess)
	sess.Where("delete_time = ?", 0).
		Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Where("member_id = ?", this.MemberId).
		Where("mtype=?", 1) //个人消息
	conds := sess.Conds()
	sess.Table(message.TableName())
	data := make([]*back.WapMemberMessageList, 0)
	err := sess.Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(message.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//获取个人消息或游戏公告详情
func (m *MemberMessageBean) WapInfo(this *input.WapMemberMessageInfo) (*back.WapMemberMessageInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	message := new(schema.MemberMessage)
	sess.Table(message.TableName()).
		Where("delete_time = ?", 0).
		Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Where("id = ?", this.Id).
		Where("member_id=?", this.MemberId)
	data := new(back.WapMemberMessageInfo)
	have, err := sess.Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, have, err
	}
	return data, have, err
}

//删除个人消息或者游戏公告
func (m *MemberMessageBean) WapDel(this *input.WapMemberMessageDel) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	message := new(schema.MemberMessage)
	sess.Table(message.TableName()).
		Where("delete_time = ?", 0).
		Where("site_id = ?", this.SiteId).
		Where("site_index_id = ?", this.SiteIndexId).
		Where("id = ?", this.Id).
		Where("member_id=?", this.MemberId)
	message.DeleteTime = time.Now().Unix()
	count, err := sess.Cols("delete_time").Update(message)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除消息
func (m *MemberMessageBean) MemberDelMessage(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	messageSchema := new(schema.MemberMessage)
	messageSchema.ID = id
	messageSchema.DeleteTime = global.GetCurrentTime()
	data, err := sess.Table(messageSchema.TableName()).Where("id=?", id).Cols("delete_time").Update(messageSchema)
	return data, err
}
