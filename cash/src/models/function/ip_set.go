package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strconv"
	"strings"
)

type IpSetBean struct{}

//ip开关列表
func (*IpSetBean) IpSetList(this *input.IpSetList, listparams *global.ListParams) ([]back.IpSetList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.IpSetList
	if this.IpStart != "" {
		sess.Where("ip_start like ?", "%"+this.IpStart+"%")
	}
	if this.IpEnd != "" {
		sess.Where("ip_end like ?", "%"+this.IpEnd+"%")
	}
	conds := sess.Conds()
	listparams.Make(sess)
	ip := new(schema.BanIp)
	err := sess.Table(ip.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(ip.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//查询ip开关是否已存在
func (*IpSetBean) BeOneIpSet(this *input.IpSetAdd) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.IpStart != "" {
		sess.Where("ip_start=?", this.IpStart)
	}
	if this.IpEnd != "" {
		sess.Where("ip_end=?", this.IpEnd)
	}
	if this.Type != "" {
		sess.Where("type=?", this.Type)
	}
	ban_ip := new(schema.BanIp)
	has, err = sess.Get(ban_ip)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return
}

//ip开关添加
func (*IpSetBean) IpSetAdd(this *input.IpSetAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ban_ip := new(schema.BanIp)
	ban_ip.IpStart = this.IpStart
	ban_ip.IpEnd = this.IpEnd
	ban_ip.Type = this.Type
	ban_ip.Remark = this.Remark
	ban_ip.State = this.State
	count, err = sess.Table(ban_ip.TableName()).InsertOne(ban_ip)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//ip开关修改
func (*IpSetBean) IpSetChange(this *input.IpSetChange) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ban_ip := new(schema.BanIp)
	ban_ip.Id = this.Id
	ban_ip.IpStart = this.IpStart
	ban_ip.IpEnd = this.IpEnd
	ban_ip.Type = this.Type
	ban_ip.Remark = this.Remark
	ban_ip.State = this.State
	count, err := sess.Where("id=?", this.Id).Update(ban_ip)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询该ip开关是否存在
func (*IpSetBean) BeIpSet(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ban_ip := new(schema.BanIp)
	has, err := sess.Where("id=?", id).Get(ban_ip)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//ip白名单查询
func (*IpSetBean) WhiteList(this *input.WhiteList, listparam *global.ListParams) ([]back.WhiteListBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.WhiteListBack
	if this.Ip != "" {
		sess.Where("ip like ?", "%"+this.Ip+"%")
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	conds := sess.Conds()
	listparam.Make(sess)
	white := new(schema.SiteWhitelist)
	err := sess.Table(white.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(white.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//查询一条白名单
func (*IpSetBean) BeWhiteList(this *input.WhiteListAdd) (data *back.WhiteListBack, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	white := new(schema.SiteWhitelist)
	data = new(back.WhiteListBack)
	has, err = sess.Table(white.TableName()).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//添加白名单
func (*IpSetBean) WhiteListAdd(this *input.WhiteListAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	white := new(schema.SiteWhitelist)
	white.SiteId = this.SiteId
	white.Ip = this.Ip
	white.State = this.Status
	white.Remark = this.Remark
	count, err = sess.InsertOne(white)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改白名单
func (*IpSetBean) WhiteListChange(this *input.WhiteListChange) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	white := new(schema.SiteWhitelist)
	white.Ip = this.Ip
	white.State = this.Status
	white.Remark = this.Remark
	count, err = sess.Where("id=?", this.Id).Where("site_id=?", this.SiteId).
		Cols("ip,state,remark").Update(white)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return
}

//根据站点查询出ip白名单
func (*IpSetBean) GetWhiteBySite(siteId string) (bool, schema.SiteWhitelist, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var white schema.SiteWhitelist
	ok, err := sess.Where("state=1").Where("site_id=?", siteId).Get(&white)
	return ok, white, err
}

//根据类型查询出ip开关
func (*IpSetBean) GetListByType(t int) ([]schema.BanIp, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var ips []schema.BanIp
	err := sess.Where("state=1").Where("type like ?", "%"+strconv.Itoa(t)+"%").Find(&ips)
	return ips, err
}

func (isb *IpSetBean) IpCheck(accountType int, siteId, checkIp string) (bool, error) {
	//ip控制
	ips, err := isb.GetListByType(accountType)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return false, err
	}
	//如果当前角色开启了ip段控制
	if len(ips) > 0 {
		ci := strings.Split(checkIp, ".")
		preIp := ci[0] + "." + ci[1] + "." + ci[2]
		isLimit := false
		for _, ip := range ips {
			if strings.HasPrefix(ip.IpStart, preIp) && strings.HasPrefix(ip.IpEnd, preIp) {
				is := strings.Split(ip.IpStart, ".")
				ie := strings.Split(ip.IpEnd, ".")
				cii, _ := strconv.Atoi(ci[3])
				isi, _ := strconv.Atoi(is[3])
				iei, _ := strconv.Atoi(ie[3])
				if cii >= isi && cii <= iei {
					isLimit = true
					break
				}
			}
		}
		if isLimit {
			//获取当前站点ip白名单
			ok, white, err := isb.GetWhiteBySite(siteId)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return false, err
			}
			if ok {
				isWhite := false
				whiteIps := strings.Split(white.Ip, ",")
				for _, ip := range whiteIps {
					if ip == checkIp {
						isWhite = true
						break
					}
				}
				if !isWhite {
					return false, nil
				}
			} else {
				return false, nil
			}
		}
	}
	return true, nil
}
