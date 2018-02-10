//[控制器] [平台]维护管理  游戏种类，后台的维护
package site

import (
	"controllers"
	"encoding/json"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
)

//公告管理
type MaintenanceController struct {
	controllers.BaseController
}

//维护页面查询 sales_site_module
func (c *MaintenanceController) GetMaintenanceInfo(ctx echo.Context) error {
	site_a := []string{"a"}

	site_list := []string{"b", "c", "d", "e", "f", "g", "h", "i", "j"}

	site_a_info, err := maintenanceBean.GetList(site_a)
	if err != nil {

	}
	site_list_info, err := maintenanceBean.GetList(site_list)
	if err != nil {
	}
	site_list_data := make(map[int]back.DomainList)
	var pm_key, wap_key, pc_key, fc_key, sp_key, vd_key, dz_key string
	for k, v := range site_a_info {
		site_one_data := back.DomainList{}
		site_one_data.SiteId = v.SiteId
		site_one_data.SiteIndexId = v.SiteIndexId
		arr := []string{} //要改地方太多,没有时间修改这个地方逻辑,谁写的自己修改一下吧
		//arr := strings.Split(v.BackstageDomain, ".")
		//fmt.Println(arr)
		pm_key = "maintain_pm_one_site_ids_" + arr[0]
		key, _ := global.GetRedis().Get(pm_key).Result()
		if key != "" {
			site_one_data.IsPmOne = 1
		} else {
			site_one_data.IsPmOne = 0
		}
		wap_key = "maintain_wap_one_site_ids_" + v.SiteId + v.SiteIndexId
		wapkey, _ := global.GetRedis().Get(wap_key).Result()
		if wapkey != "" {
			site_one_data.IsWapOne = 1
		} else {
			site_one_data.IsWapOne = 0
		}
		pc_key = "maintain_pc_one_site_ids_" + v.SiteId + v.SiteIndexId
		pckey, _ := global.GetRedis().Get(pc_key).Result()
		if pckey != "" {
			site_one_data.IsPcOne = 1
		} else {
			site_one_data.IsPcOne = 0
		}
		fc_key = "maintain_fc_one_site_ids_" + v.SiteId + v.SiteIndexId
		fckey, _ := global.GetRedis().Get(fc_key).Result()
		if fckey != "" {
			site_one_data.IsPcOne = 1
		} else {
			site_one_data.IsPcOne = 0
		}
		sp_key = "maintain_sp_one_site_ids_" + v.SiteId + v.SiteIndexId
		spkey, _ := global.GetRedis().Get(sp_key).Result()
		if spkey != "" {
			site_one_data.IsSpOne = 1
		} else {
			site_one_data.IsSpOne = 0
		}
		vd_key = "maintain_vd_one_site_ids_" + v.SiteId + v.SiteIndexId
		vdkey, _ := global.GetRedis().Get(vd_key).Result()
		if vdkey != "" {
			site_one_data.IsVdOne = 1
		} else {
			site_one_data.IsVdOne = 0
		}
		dz_key = "maintain_dz_one_site_ids_" + v.SiteId + v.SiteIndexId
		dzkey, _ := global.GetRedis().Get(dz_key).Result()
		if dzkey != "" {
			site_one_data.IsDzOne = 1
		} else {
			site_one_data.IsDzOne = 0
		}

		for _, value := range site_list_info {
			//fmt.Println(k1)
			if v.SiteId == value.SiteId {
				listdata := back.DomainList{}

				sub_key_fc := "maintain_fc_one_site_ids_" + value.SiteId + value.SiteIndexId
				sub_key_sp := "maintain_sp_one_site_ids_" + value.SiteId + value.SiteIndexId
				sub_key_vd := "maintain_vd_one_site_ids_" + value.SiteId + value.SiteIndexId
				sub_key_dz := "maintain_dz_one_site_ids_" + value.SiteId + value.SiteIndexId
				sub_key_wap := "maintain_wap_one_site_ids_" + value.SiteId + value.SiteIndexId
				sub_key_pc := "maintain_pc_one_site_ids_" + value.SiteId + value.SiteIndexId
				sub_key_pm := "maintain_gm_one_site_ids_" + arr[0]

				subkeyfc, _ := global.GetRedis().Get(sub_key_fc).Result()
				if subkeyfc != "" {
					listdata.IsFcOne = 1
				} else {
					listdata.IsFcOne = 0
				}
				subkeysp, _ := global.GetRedis().Get(sub_key_sp).Result()
				if subkeysp != "" {
					listdata.IsSpOne = 1
				} else {
					listdata.IsSpOne = 0
				}
				subkeyvd, _ := global.GetRedis().Get(sub_key_vd).Result()
				if subkeyvd != "" {
					listdata.IsVdOne = 1
				} else {
					listdata.IsVdOne = 0
				}
				subkeydz, _ := global.GetRedis().Get(sub_key_dz).Result()
				if subkeydz != "" {
					listdata.IsDzOne = 1
				} else {
					listdata.IsDzOne = 0
				}
				subkeywap, _ := global.GetRedis().Get(sub_key_wap).Result()
				if subkeywap != "" {
					listdata.IsWapOne = 1
				} else {
					listdata.IsWapOne = 0
				}
				subkeypc, _ := global.GetRedis().Get(sub_key_pc).Result()
				if subkeypc != "" {
					listdata.IsPcOne = 1
				} else {
					listdata.IsPcOne = 0
				}
				subkeypm, _ := global.GetRedis().Get(sub_key_pm).Result()
				if subkeypm != "" {
					listdata.IsPmOne = 1
				} else {
					listdata.IsPmOne = 0
				}
				listdata.Id = value.Id
				listdata.SiteId = value.SiteId
				listdata.SiteIndexId = value.SiteIndexId
				//listdata.BackstageDomain = value.BackstageDomain
				site_one_data.List = append(site_one_data.List, listdata)
			}
		}
		site_list_data[k] = site_one_data
	}
	return ctx.JSON(200, global.ReplyItem(site_list_data))
}

//维护信息修改
func (c *MaintenanceController) PutMaintenNoticeUpdate(ctx echo.Context) error {
	main_data := new(input.MaintenanceList)
	code := global.ValidRequestAdmin(main_data, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断是否有下级数据
	if main_data.Type == "vd" || main_data.Type == "dz" {
		if main_data.Module == nil {
			return ctx.JSON(200, global.ReplyError(91101, ctx)) //未选择维护模块
		}
	}
	var maintainkey string
	condition := back.ConditionList{}
	if main_data.Wtype == "one" {
		//单站维护
		if main_data.SiteId == "" {
			return ctx.JSON(200, global.ReplyError(91102, ctx)) //维护site_id不能为空
		}
		if main_data.SiteIndexId == "" {
			maintainkey = "_" + main_data.SiteId
		} else {
			maintainkey = "_" + main_data.SiteId + main_data.SiteIndexId
			condition.SiteId = main_data.SiteId
			condition.SiteIndexId = main_data.SiteIndexId
			condition.DeleteTime = 0
		}
	} else {
		//全网维护
		if main_data.Wtype == "all" {
			condition.DeleteTime = 0
		} else {
			return ctx.JSON(200, global.ReplyError(91103, ctx)) //参数错误
		}
	}
	data, err := maintenanceBean.GetDomainList(condition)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	main_key := "maintain_" + main_data.Type + "_" + main_data.Wtype + "_site_ids" + maintainkey

	result := c.DataHandle(main_data, data)
	result_data, _ := json.Marshal(result)
	//存储新token

	err = global.GetRedis().Set(main_key, result_data, 0).Err()

	if err != nil {
		return ctx.JSON(200, global.ReplyError(91104, ctx))
	}
	//将推进list
	err = global.GetRedis().RPush("maintenance_info", main_key).Err()

	if err != nil {
		return ctx.JSON(200, global.ReplyError(91105, ctx))
	}
	return ctx.NoContent(204)
}

//获取视讯 电子 下级菜单
func (c *MaintenanceController) GetInfoList(ctx echo.Context) error {
	main_data := new(input.InfoList)
	code := global.ValidRequestAdmin(main_data, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var typeId int64
	if main_data.Type == "dz" {
		typeId = 2
	} else if main_data.Type == "vd" {
		typeId = 1
	} else {
		return ctx.JSON(200, global.ReplyError(91106, ctx))
	}
	data, err := maintenanceBean.GetInfoList(main_data, typeId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(91107, ctx))
	}
	info := c.GetOneModule(main_data)
	if info != nil {
		for k, v := range data {
			for _, v1 := range info[0].Module {
				if v.VType == v1.VType {
					data[k].IsState = 1
				}
			}
		}
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//数据处理
func (c *MaintenanceController) DataHandle(mainData *input.MaintenanceList, data []back.DomainInfoList) map[int]back.SaveMainTain {
	datalist := make(map[int]back.SaveMainTain)
	if mainData.Wtype == "one" {
		main_one := back.SaveMainTain{}
		main_one.SiteId = data[0].SiteId
		main_one.Content = mainData.Content
		main_one.QQ = data[0].QQ
		main_one.Wechat = data[0].Wechat
		main_one.Phone = data[0].Phone
		main_one.Email = data[0].Email
		datalist[0] = main_one
	} else {
		for k, v := range data {
			main_one := back.SaveMainTain{}
			main_one.SiteId = v.SiteId
			main_one.SiteIndexId = v.SiteIndexId
			main_one.Module = mainData.Module
			main_one.QQ = data[0].QQ
			main_one.Wechat = data[0].Wechat
			main_one.Phone = data[0].Phone
			main_one.Email = data[0].Email
			main_one.Content = mainData.Content
			datalist[k] = main_one
		}
	}
	return datalist
}

//获取单站维护模块
func (c *MaintenanceController) GetOneModule(data *input.InfoList) map[int]back.SaveMainTain {
	var one_key string
	result := make(map[int]back.SaveMainTain)
	one_key = "maintain_" + data.Type + "_one_site_ids_" + data.SiteId + data.SiteIndexId
	list, _ := global.GetRedis().Get(one_key).Result()
	json.Unmarshal([]byte(list), &result)
	return result
}
