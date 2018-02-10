//[控制器] [平台] 站点管理
package site

import (
	"controllers"
	"fmt"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
)

//站点管理
type SiteController struct {
	controllers.BaseController
}

//站点列表查询(GET 分页 并且整个站点情况统计返回)
func (c *SiteController) GetSiteList(ctx echo.Context) error {
	site_manage := new(input.SiteManageList)
	code := global.ValidRequestAdmin(site_manage, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if site_manage.MoreCondition != 0 && site_manage.MoreContent != "" {
		switch site_manage.MoreCondition {
		case 1:
			site_manage.SiteId = site_manage.MoreContent
		case 2:
			site_manage.SiteName = site_manage.MoreContent
		case 3:
			site_manage.SiteDomain = site_manage.MoreContent
		}
	}
	//获取listparam的数据
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := siteControllBean.SiteManageList(site_manage, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//查询所有没删除的站点
	dataA, err := siteControllBean.SiteManageListDataA()
	if len(data) > 0 && len(dataA) > 0 {
		for k, v := range data {
			var i int64
			for _, n := range dataA {
				if v.Id == n.Id && n.IsDefault == 2 {
					i = i + 1
				}
			}
			if i < 1 {
				data[k].MoreSite = 2
			} else {
				data[k].MoreSite = 1
			}
		}
	}
	//查询站点数目
	nuM := new(back.SiteListNumber)
	nuM.SiteNum = count
	if len(data) > 0 {
		for _, v := range data {
			if v.Status == 1 {
				nuM.OpenSite = nuM.OpenSite + 1
			}
		}
	}
	var list = make(map[string]interface{})
	list["data"] = data
	list["num"] = nuM
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(data)), count, ctx))
}

//站点添加
func (c *SiteController) PostSiteAdd(ctx echo.Context) error {
	addSite := new(input.AddSite) //增加站点的请求你参数struct
	code := global.ValidRequestAdmin(addSite, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看站点是否存在
	info, flag, err := siteOperateBean.GetInfoBySiteIndexId(addSite.Site, addSite.SiteIndex)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag && info.DeleteTime == 0 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	//站点名称是否已经使用的校验
	flag, err = siteOperateBean.GetSingleSiteByName(addSite.SiteName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag {
		return ctx.JSON(200, global.ReplyError(60002, ctx))
	}
	//检查套餐是否存在
	combo := new(input.ComboId)
	combo.Id = addSite.ComboId
	_, flag, err = comboBeen.GetInfo(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60111, ctx))
	}
	//检验pc域名是否合法
	ok := global.DomainCheck(addSite.Domain)
	if !ok {
		return ctx.JSON(200, global.ReplyError(60013, ctx))
	}

	//检验后台域名是否合法
	ok = global.DomainCheck(addSite.BackstageDomain)
	if !ok {
		return ctx.JSON(200, global.ReplyError(60098, ctx))
	}
	//检验三个域名是否被使用
	num, err := comboBeen.CheckThirdDoMain(addSite.Domain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != 0 {
		return ctx.JSON(200, global.ReplyError(num, ctx))
	}
	//增加
	count, err := siteOperateBean.Add(addSite)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60060, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//站点信息修改
func (c *SiteController) PutSiteUpdate(ctx echo.Context) error {
	editSite := new(input.EditSite) //修改站点的请求参数struct
	code := global.ValidRequest(editSite, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断站点是否存在   站点名是否可用
	num, err := siteOperateBean.GetSiteInfomation(editSite.Site,
		editSite.SiteIndex, editSite.SiteName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != 0 {
		return ctx.JSON(200, global.ReplyError(num, ctx))
	}
	combo := new(input.ComboId)
	combo.Id = editSite.ComboId
	_, flag, err := comboBeen.GetInfo(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60111, ctx))
	}
	//检验pc域名是否合法
	ok := global.DomainCheck(editSite.Domain)
	if !ok {
		return ctx.JSON(200, global.ReplyError(60013, ctx))
	}
	//检验三个域名是否被使用
	num, err = comboBeen.CheckThirdDoMainChange(
		editSite.Site, editSite.SiteIndex,
		editSite.Domain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != 0 {
		return ctx.JSON(200, global.ReplyError(num, ctx))
	}
	//更新操作
	count, err := siteOperateBean.Edit(editSite)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30284, ctx))
	}
	return ctx.NoContent(204)
}

//站点状态修改
func (c *SiteController) PutSiteStatusUpdate(ctx echo.Context) error {
	//获取数据
	siteManage := new(input.SiteManageStatus)
	code := global.ValidRequestAdmin(siteManage, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询该站点是否存在
	has, err := siteControllBean.ManageBySiteSiteIndexIdStatus(siteManage)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	if siteManage.Status != 1 {
		//查询该站点在agency表中是否被使用
		has, err = siteControllBean.ManageBySiteSiteIndexID(siteManage)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//如果agency表中存在该站点下级，则不能修改其状态
		if has {
			return ctx.JSON(200, global.ReplyError(50122, ctx))
		}
	}
	count, err := siteControllBean.SiteManageStatus(siteManage)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(60011, ctx))
	}
	return ctx.NoContent(204)
}

//站点删除
func (c *SiteController) SiteManageDelete(ctx echo.Context) error {
	//获取数据
	site_manage := new(input.SiteManageStatus)
	code := global.ValidRequestAdmin(site_manage, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//如果该站点不是在停用的状态，不能删除
	if site_manage.Status != 2 {
		return ctx.JSON(200, global.ReplyError(50123, ctx))
	}
	//查询该站点是否存在
	has, err := siteControllBean.ManageBySiteSiteIndexID(site_manage)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//查询该站点在agency表中是否被使用
	has, err = siteControllBean.ManageBySiteSiteIndexID(site_manage)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//如果agency表中存在该站点下级，则不能删除
	if has {
		return ctx.JSON(200, global.ReplyError(50124, ctx))
	}
	count, err := siteControllBean.SiteManageDelete(site_manage)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//站点模块管理（站点商品查询）
/*func (c *SiteController) GetSiteProductList(ctx echo.Context) error {
	siteProduct := new(input.SiteProductList)
	code := global.ValidRequestAdmin(siteProduct, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看站点是否存在
	has, err := siteProductBean.IsExistSite(siteProduct.SiteId, siteProduct.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//查看站点下的套餐
	comboId, has, err := siteProductBean.GetSiteCombo(siteProduct.SiteId, siteProduct.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30216, ctx))
	}
	//查看站点下剔除的商品id
	productIds, err := siteProductBean.SiteProductDel(siteProduct.SiteId, siteProduct.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//查询该站点下的套餐中所有商品
	list, err := siteProductBean.ProductList(comboId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var listBacks []back.SiteProductBack
	var listBack back.SiteProductBack
	for j := range list {
		listBack.SiteId = siteProduct.SiteId
		listBack.SiteIndexId = siteProduct.SiteIndexId
		listBack.ProductId = list[j].Id
		listBack.ProductName = list[j].ProductName
		listBack.TypeId = list[j].TypeId
		listBack.Title = list[j].Title
		listBacks = append(listBacks, listBack)
	}
	if len(productIds) != 0 {
		for k := range listBacks {
			for i := range productIds {
				if listBacks[k].ProductId == productIds[i] {
					listBacks[k].IsCheck = 2
				}
			}
		}
	}
	return ctx.JSON(200, global.ReplyItem(listBacks))
}*/

//站点模块管理（站点商品剔除）
/*func (c *SiteController) InsertSiteProductDel(ctx echo.Context) error {
	siteProduct := new(input.SiteProductEdit)
	code := global.ValidRequestAdmin(siteProduct, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看站点是否存在
	has, err := siteProductBean.IsExistSite(siteProduct.SiteId, siteProduct.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//查看站点下的套餐
	comboId, has, err := siteProductBean.GetSiteCombo(siteProduct.SiteId, siteProduct.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30216, ctx))
	}
	//查看站点下套餐中商品是否存在
	count, err := siteProductBean.IsExistProduct(siteProduct.ProductId, comboId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != int64(len(siteProduct.ProductId)) {
		return ctx.JSON(200, global.ReplyError(30214, ctx))
	}
	//把前端传过来的商品id以外的商品id查出来
	ids, err := siteProductBean.GetProductId(siteProduct.ProductId, comboId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//如果有剔除的商品，就做修改。
	siteProduct.ProductId = ids
	count, err = siteProductBean.SiteProductUpdate(siteProduct)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30215, ctx))
	}
	return ctx.NoContent(204)
}*/

//多站点-查询列表
func (c *SiteController) GetSiteMoreNews(ctx echo.Context) error {
	siteMore := new(input.SiteMoreList)
	code := global.ValidRequestAdmin(siteMore, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := siteOperateBean.GetSiteMoreList(siteMore.SiteId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	nuM := new(back.SiteListNumberR)
	nuM.ChildNum = int64(len(data))
	if len(data) > 0 {
		for _, v := range data {
			if v.Status == 1 {
				nuM.OpenSite = nuM.OpenSite + 1
			}
		}
	}
	var list = make(map[string]interface{})
	list["data"] = data
	list["num"] = nuM
	return ctx.JSON(200, global.ReplyItem(list))
}

//多站点-添加代理(由于多站点都是都是对应的开户人账号，所以这里就是添加股东)
func (c *SiteController) PostSiteMore(ctx echo.Context) error {
	siteMore := new(input.SiteMoreAdd)
	code := global.ValidRequestAdmin(siteMore, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断站点是否存在
	has, err := memberLevelBean.GetSiteIndexId(siteMore.SiteId, siteMore.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//判断站点下代理账号是否存在
	has, err = agencyBean.IsExistAccountAdd(siteMore.SiteId, siteMore.SiteIndexId, siteMore.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30209, ctx))
	}
	//根据站点获取开户人id
	parentId, has, err := siteOperateBean.GetIdBySite(siteMore.SiteId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30219, ctx))
	}
	count, err := siteOperateBean.PostSiteMore(siteMore, parentId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30051, ctx))
	}
	return ctx.NoContent(204)
}

//一键生成的账号-下发
func (c *SiteController) ProductAgentsAccount(ctx echo.Context) error {
	productAccount := new(input.ProductAccount) //这里传递过来的站点id和站点前台id用于账号生成
	code := global.ValidRequestAdmin(productAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	newAccount := []string{}
	for {
		account := global.RandStringBytesMaskImprSrc(5)
		accountNew := fmt.Sprintf("%s%s_%s", productAccount.SiteId, productAccount.SiteIndexId, account)
		//判断是否重复
		if len(newAccount) == 0 {
			flag, err := agencyBean.IsExistAccount(accountNew)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			if !flag {
				newAccount = append(newAccount, accountNew)
			}
		} else {
			i := 0
			for _, k := range newAccount {
				if accountNew == k {
					i++
				}
			}
			if i == 0 {
				flag, err := agencyBean.IsExistAccount(accountNew)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					return ctx.JSON(500, global.ReplyError(60000, ctx))
				}
				if !flag {
					newAccount = append(newAccount, accountNew)
				}
			}
		}
		if len(newAccount) == 3 {
			break
		}
	}

	return ctx.JSON(200, global.ReplyItem(newAccount))
}

//一键生成站点默认三级代理 股东 总代 代理之后的提交
func (c *SiteController) PostSiteAgentsAdd(ctx echo.Context) error {
	addAgency := new(input.GenerationAgency)
	code := global.ValidRequestAdmin(addAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//首先判断对应的站点的三级默认代理是否都不存在
	infolist, err := agencyBean.IsExistDefaultAgency(addAgency.SiteId, addAgency.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(infolist) != 0 {
		return ctx.JSON(200, global.ReplyError(60081, ctx))
	}
	//判断代理账号是否存在
	flag, err := agencyBean.IsExistAccount(addAgency.DefaultAgencyAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag {
		return ctx.JSON(200, global.ReplyError(60082, ctx))
	}
	//判断股东账号是否存在
	flag, err = agencyBean.IsExistAccount(addAgency.DefaultShareholdersAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag {
		return ctx.JSON(200, global.ReplyError(60084, ctx))
	}
	//判断总代理账号是否存在
	flag, err = agencyBean.IsExistAccount(addAgency.DefaultTotalAgencyAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag {
		return ctx.JSON(200, global.ReplyError(60083, ctx))
	}
	//判断开户人是否存在
	info, flag, err := agencyBean.GetOpenInfo(addAgency.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if info.Id == 0 || !flag {
		return ctx.JSON(200, global.ReplyError(60085, ctx))
	}
	//增加
	count, err := agencyBean.AddThirdAgency(addAgency, info.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30268, ctx))
	}
	return ctx.NoContent(204)
}

//站点代理数据列表查询
func (c *SiteController) GetSiteAgentsList(ctx echo.Context) error {
	agency := new(input.SiteAgencyList)
	code := global.ValidRequestAdmin(agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := agencyBean.AgencyList(agency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//站点代理数据查询
func (c *SiteController) GetSiteAgentsInfo(ctx echo.Context) error {
	agency := new(input.SiteAgencyInfo)
	code := global.ValidRequestAdmin(agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := agencyBean.AgencyInfo(agency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30208, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//站点代理数据添加
func (c *SiteController) PostSiteAgentAdd(ctx echo.Context) error {
	agency := new(input.SiteAgencyAdd)
	code := global.ValidRequestAdmin(agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断站点是否存在
	has, err := memberLevelBean.GetSiteIndexId(agency.SiteId, agency.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//判断站点下代理账号是否存在
	has, err = agencyBean.IsExistAccountAdd(agency.SiteId, agency.SiteIndexId, agency.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30209, ctx))
	}
	//根据账号获取上级账号详情
	ab, err := agencyBean.GetParentId(agency.ParentAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !(ab.RoleId == 2 || ab.RoleId == 3) {
		return ctx.JSON(200, global.ReplyError(30218, ctx))
	}
	//添加
	count, err := agencyBean.AgencyAdd(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30054, ctx))
	}
	return ctx.NoContent(204)
}

//站点代理数据修改
func (c *SiteController) PutSiteAgentUpdate(ctx echo.Context) error {
	agency := new(input.SiteAgencyEdit)
	code := global.ValidRequestAdmin(agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断站点是否存在
	has, err := memberLevelBean.GetSiteIndexId(agency.SiteId, agency.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	////判断站点下代理账号是否存在【有需要修改账号的再打开修改】
	//has, err = agencyBean.IsExistAccountEdit(agency.Id, agency.SiteId, agency.SiteIndexId, agency.Account)
	//if err != nil {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}
	//if !has {
	//	return ctx.JSON(200, global.ReplyError(50007, ctx))
	//}
	count, err := agencyBean.AgencyEdit(agency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30279, ctx))
	}
	return ctx.NoContent(204)
}

//站点代理数据删除
func (c *SiteController) DeleteSiteAgentDel(ctx echo.Context) error {
	agency := new(input.SiteAgencyInfo)
	code := global.ValidRequestAdmin(agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断站点下代理账号是否存在
	has, err := agencyBean.IsExistAccountEdit(agency.Id, "", "", "")
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	count, err := agencyBean.AgencyDel(agency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30210, ctx))
	}
	return ctx.NoContent(204)
}

//会员层级列表查询
func (c *SiteController) GetMemberLevelList(ctx echo.Context) error {
	site := new(input.SiteLevelList)
	code := global.ValidRequestAdmin(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := memberLevelBean.MemberLevelList(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//会员层级查询
func (c *SiteController) GetMemberLevelInfo(ctx echo.Context) error {
	site := new(input.SiteLevelInfo)
	code := global.ValidRequestAdmin(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := memberLevelBean.MemberLevelInfo(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//会员层级添加
func (c *SiteController) PostMemberLevelAdd(ctx echo.Context) error {
	site := new(input.SiteLevelAdd)
	code := global.ValidRequestAdmin(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断站点下站点前台id是否存在
	has, err := memberLevelBean.GetSiteIndexId(site.SiteId, site.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//判断层级是否已存在
	has, err = memberLevelBean.GetLevel(site.SiteId, site.SiteIndexId, site.LevelId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30199, ctx))
	}
	count, err := memberLevelBean.AddMemberLevel(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(30198, ctx))
	}
	return ctx.NoContent(204)
}

//会员层级修改
func (c *SiteController) PutMemberLevelUpdate(ctx echo.Context) error {
	site := new(input.SiteLevelEdit)
	code := global.ValidRequestAdmin(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断站点下站点前台id是否存在
	has, err := memberLevelBean.GetSiteIndexIdLevelId(site.SiteId, site.SiteIndexId, site.LevelId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10017, ctx))
	}
	count, err := memberLevelBean.EditMemberLevel(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30278, ctx))
	}
	return ctx.NoContent(204)
}

//会员层级删除
func (c *SiteController) PutMemberLevelDel(ctx echo.Context) error {
	site := new(input.SiteLevelInfo)
	code := global.ValidRequestAdmin(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断层级下是否有会员
	has, err := memberLevelBean.GetMember(site.SiteId, site.SiteIndexId, site.LevelId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30217, ctx))
	}
	//判断层级是否为默认层级
	has, err = memberLevelBean.GetLevelIsDefault(site.SiteId, site.SiteIndexId, site.LevelId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30207, ctx))
	}
	err = memberLevelBean.DelMemberLevel(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//站点管理员查询
func (c *SiteController) GetSiteAdminInfo(ctx echo.Context) error {
	account := new(input.AccountHolderList)
	code := global.ValidRequestAdmin(account, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取listparam的数据
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := agencyBean.GetAllAccountHolder(account, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var total back.OnlineNumberAndTotal
	if len(data) > 0 {
		var i int64
		for _, v := range data {
			if v.IsLogin == 1 {
				i = i + 1
			}
		}
		total.OnlineNumber = i

	} else {
		total.OnlineNumber = 0
	}
	total.TotalNumber = count
	var list = make(map[string]interface{})
	list["data"] = data
	list["total"] = total
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(data)), count, ctx))
}

//站点管理员添加
func (c *SiteController) PostSiteAdminAdd(ctx echo.Context) error {
	accountInfo := new(input.AddAccountHolder)
	code := global.ValidRequestAdmin(accountInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	accountNameId := new(input.AccountNameId)
	accountNameId.Account = accountInfo.Account
	//查看帐号是否存在
	_, has, err := agencyBean.GetOneAccountHelder(accountNameId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//查看站点是否存在
	siteSchema := new(schema.Site)
	siteSchema.Id = accountInfo.Site
	has, err = siteOperateBean.GetOneSiteId(siteSchema)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(50076, ctx))
	}
	flag, err := siteOperateBean.GetSingleSiteByName(accountInfo.SiteName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag {
		return ctx.JSON(200, global.ReplyError(50077, ctx))
	}
	//密码和重复密码是否相同
	if accountInfo.Password != accountInfo.RePassword {
		return ctx.JSON(200, global.ReplyError(50070, ctx))
	}
	//客户后台域名不能跟代理后台域名一样
	if accountInfo.ManageDomain == accountInfo.AgencyDomain {
		return ctx.JSON(200, global.ReplyError(50199, ctx))
	}
	//客户后台域名是否被使用
	ok, err := siteDomainBean.IsExistDomain(accountInfo.ManageDomain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if ok {
		return ctx.JSON(200, global.ReplyError(50197, ctx))
	}
	//代理后台域名是否被使用
	ok, err = siteDomainBean.IsExistDomain(accountInfo.AgencyDomain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if ok {
		return ctx.JSON(200, global.ReplyError(50198, ctx))
	}
	//密码加密
	if accountInfo.Password != "" {
		md5Password, err := global.MD5ByStr(accountInfo.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		accountInfo.Password = md5Password
	}
	//操作密码加密
	if accountInfo.OperatePassword != "" {
		md5Password, err := global.MD5ByStr(accountInfo.OperatePassword, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		accountInfo.OperatePassword = md5Password
	}
	//添加帐号
	count, err := agencyBean.AddAccountHolder(accountInfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//站点管理员修改
func (c *SiteController) PutSiteAdminUpdate(ctx echo.Context) error {
	accountInfo := new(input.UpdataAccountHolder)
	code := global.ValidRequestAdmin(accountInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看帐号是否存在
	accountNameId := new(input.AccountNameId)
	accountNameId.Id = accountInfo.Id
	accountNameId.Account = accountInfo.Account
	_, has, err := agencyBean.GetOneAccountHelder(accountNameId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//密码和重复密码是否相同
	if accountInfo.Password != accountInfo.RePassword {
		return ctx.JSON(200, global.ReplyError(50070, ctx))
	}
	//取出原密码
	data, _, err := agencyBean.AccountHelderGet(accountInfo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//密码加密
	if accountInfo.Password != "" {
		md5Password, err := global.MD5ByStr(accountInfo.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		accountInfo.Password = md5Password
	} else {
		accountInfo.Password = data.Password
	}
	//操作密码加密
	//if account_info.OperatePassword != "" {
	//	md5Password, err := global.MD5ByStr(account_info.OperatePassword, global.EncryptSalt)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(30044, ctx))
	//	}
	//	account_info.OperatePassword = md5Password
	//} else {
	//	account_info.OperatePassword = data.OperatePassword
	//}
	count, err := agencyBean.UpdataHolder(accountInfo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30280, ctx))
	}
	return ctx.NoContent(204)
}

//站点管理员删除(也是更新修改的一种，看能否和上面合并)
func (c *SiteController) PutSiteAdminDel(ctx echo.Context) error {
	account_info := new(input.DelAccountHolder)
	code := global.ValidRequest(account_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	has, err := siteOperateBean.BeSiteOne(account_info.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))

	}
	if has {
		return ctx.JSON(200, global.ReplyError(50075, ctx))
	}
	count, err := agencyBean.AccountHolderDelete(account_info)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//开户人信息(GET)
func (ahc *SiteController) EditAccountHolderInfo(ctx echo.Context) error {
	a_info := new(input.AccountHolderInfoIn)
	code := global.ValidRequestAdmin(a_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	account_info := new(input.AccountNameId)
	account_info.Id = a_info.Id
	data, _, err := agencyBean.GetAccountHelderInfo(account_info)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//设置开户人状态[开启/禁用](PUT)
func (ahc *SiteController) AccountHolderOpenClose(ctx echo.Context) error {
	accountInfo := new(input.HolderNameId)
	code := global.ValidRequestAdmin(accountInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//状态开启要做禁用功能
	if accountInfo.Status == 1 {
		siteSchema := new(schema.Site)
		siteSchema.AgencyId = accountInfo.Id
		data, err := siteOperateBean.GetSiteIndexId(siteSchema)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))

		}
		if data != nil {
			return ctx.JSON(200, global.ReplyError(50071, ctx))
		}
		count, err := agencyBean.AccountHolderDisable(accountInfo)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(200, global.ReplyError(30281, ctx))
		}
	}
	//状态禁用要做开启功能
	if accountInfo.Status == 2 {
		count, err := agencyBean.AccountHolderDisable(accountInfo)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(200, global.ReplyError(30281, ctx))
		}
	}
	return ctx.NoContent(204)
}

//查询站点视讯账号是否存在
func (c *SiteController) GetSiteVedioAccount(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//一键生成站点视讯账号
func (c *SiteController) PostSiteVedioAccountAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//站点后台管理导航栏目控制查询
func (c *SiteController) GetSiteRole(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//站点后台管理导航栏目控制  添加修改  存在即修改，不存在就添加
func (c *SiteController) PostSiteRoleAddUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//站点代理（代理登录）后台管理导航栏目控制查询
func (c *SiteController) GetSiteAgentRole(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//站点代理（代理登录）后台管理导航栏目控制  添加修改  存在即修改，不存在就添加
func (c *SiteController) PostSiteAgentRoleAddUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//站点上线时间修改
func (c *SiteController) PutSiteOnlinTimeUpdate(ctx echo.Context) error {
	online := new(input.SiteOnline)
	code := global.ValidRequest(online, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//站点是否存在
	info, flag, err := siteOperateBean.GetInfoBySiteIndexId(online.SiteId, online.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag || info.Id == "" {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}

	count, err := siteOperateBean.UpSiteOnlineTime(online)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//站点负数情况查询
/*func (c *SiteController) GetSiteNegative(ctx echo.Context) error {
	negative := new(input.ReportNegativeList)
	code := global.ValidRequest(negative, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := siteOperateBean.NegativeList(negative)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}*/

//站点负数情况添加记录
/*func (c *SiteController) PostSiteNegativeAdd(ctx echo.Context) error {
	negative := new(input.ReportNegativeAdd)
	code := global.ValidRequest(negative, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteOperateBean.NegativeAdd(negative)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(10147, ctx))
	}
	return ctx.NoContent(204)
}*/

//站点负数情况修改
/*func (c *SiteController) PutSiteNegativeUpdate(ctx echo.Context) error {
	negative := new(input.ReportNegativeEdit)
	code := global.ValidRequest(negative, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteOperateBean.NegativeEdit(negative)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(10148, ctx))
	}
	return ctx.NoContent(204)
}*/

//全站商品维护查询
func (c *SiteController) GetSiteMaintenance(ctx echo.Context) error {
	list, err := siteOperateBean.GetSiteMaintenance()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//获取站点是否选择维护
func (sc *SiteController) SiteIsSelect(ctx echo.Context) error {
	maintenance := new(input.SiteIsSelect)
	code := global.ValidRequestAdmin(maintenance, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取站点数据
	siteList, err := siteOperateBean.SiteSiteIndexIdBy()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//根据id获取维护表设置站点
	siteIdS, err := siteOperateBean.GetSiteIdS(maintenance.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var sis back.SiteIsSelect
	list := make([]back.SiteIsSelect, 0)
	//遍历站点数据并用下划线拼接site_id和site_index_id
	if len(siteList) > 0 {
		for _, k := range siteList {
			sis.Id = k.Id
			sis.IndexId = k.IndexId
			sis.SiteName = k.SiteName
			sis.IsSelect = 2
			list = append(list, sis)
		}
		for o, j := range list {
			ids := j.Id + "_" + j.IndexId
			if siteIdS.SiteIdS == "0" {
				list[o].IsSelect = 1
			} else {
				ss := strings.Split(siteIdS.SiteIdS, ",")
				for _, l := range ss {
					if ids == l {
						list[o].IsSelect = 1
					}
				}
			}
		}
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//全站商品维护修改
func (c *SiteController) PostSiteMaintenanceUpdate(ctx echo.Context) error {
	maintenance := new(input.SiteMaintenance)
	code := global.ValidRequestAdmin(maintenance, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteOperateBean.PutSiteMaintenance(maintenance)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30285, ctx))
	}
	return ctx.NoContent(204)
}

//SiteQuota 站点视讯额度操作
func (c *SiteController) SiteQuota(ctx echo.Context) error {
	siteVideo := new(input.SiteVideoBalance)
	code := global.ValidRequest(siteVideo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作人信息
	userinfo := ctx.Get("admin").(*global.AdminRedisStruct)
	//判断该开户人是否存在
	info, flag, err := agencyBean.GetOpenInfo(siteVideo.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if info.Id == 0 || !flag {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//更改额度
	count, err := agencyBean.UpdateSiteVideoBalance(siteVideo, *userinfo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//前台-文案类型
func (c SiteController) SiteCopyType(ctx echo.Context) error {
	data, err := siteOperateBean.SiteCopyType()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//多站点-前台首页文案列表
func (c *SiteController) SiteCopyList(ctx echo.Context) error {
	copylist := new(input.SiteCopyList)
	code := global.ValidRequest(copylist, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := siteOperateBean.SiteCopyList(copylist)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//多站点-前台首页文案详情
func (c *SiteController) SiteCopyListInfoOne(ctx echo.Context) error {
	copylist := new(input.SiteCopyListInfoOne)
	code := global.ValidRequest(copylist, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, has, err := siteOperateBean.SiteCopyListInfo(copylist)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(90716, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//添加文案
func (c *SiteController) SiteCopyAdd(ctx echo.Context) error {
	copylist := new(input.SiteCopyAdd)
	code := global.ValidRequest(copylist, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := siteOperateBean.SiteCopyAdd(copylist)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(500, global.ReplyError(90701, ctx))
	}
	return ctx.NoContent(204)
}

//文案修改
func (c *SiteController) SiteCopyUpdate(ctx echo.Context) error {
	copyUpdate := new(input.CopyUpdate)
	code := global.ValidRequest(copyUpdate, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := siteOperateBean.SiteCopyUpdate(copyUpdate)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(500, global.ReplyError(90713, ctx))
	}
	return ctx.NoContent(204)
}

//轮播查询
func (c *SiteController) FlashList(ctx echo.Context) error {
	flash_list := new(input.FlashList)
	code := global.ValidRequest(flash_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := siteOperateBean.FlashList(flash_list.SiteId, flash_list.SiteIndexId, 0, 0)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(data) == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	datas := make(map[string][]string)
	for i := range data {
		if data[i].Ftype == 1 {
			datas["pc"] = append(datas["pc"], data[i].ImgUrl)
		}
		if data[i].Ftype == 2 {
			datas["wap"] = append(datas["wap"], data[i].ImgUrl)
		}
	}
	return ctx.JSON(200, global.ReplyItem(datas))
}

//轮播添加
func (c *SiteController) FlashAdd(ctx echo.Context) error {
	flash_add := new(input.FlashAdd)
	code := global.ValidRequest(flash_add, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := siteOperateBean.FlashAdd(flash_add)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 5 {
		return ctx.JSON(500, global.ReplyError(50169, ctx))
	}
	return ctx.NoContent(204)
}

//站点logo图片管理
func (c *SiteController) LogoList(ctx echo.Context) error {
	logo_list := new(input.LogoList)
	code := global.ValidRequest(logo_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := siteOperateBean.LogoList(logo_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//站点logo图片添加
func (c *SiteController) LogoAdd(ctx echo.Context) error {
	logoAdd := new(input.LogoAdd)
	code := global.ValidRequest(logoAdd, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询是否已存在
	has, err := siteOperateBean.LogoInfo(logoAdd)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(50170, ctx))
	}
	data, err := siteOperateBean.LogoAdd(logoAdd)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(30295, ctx))
	}
	return ctx.NoContent(204)
}

//站点域名编辑
func (c *SiteController) SiteDomainEdit(ctx echo.Context) error {
	siteDomain := new(input.SiteDomainEdit)
	code := global.ValidRequest(siteDomain, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断站点是否存在   站点名是否可用
	num, err := siteOperateBean.GetSiteInfomation(siteDomain.SiteId,
		siteDomain.SiteIndexId, siteDomain.SiteName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != 0 {
		return ctx.JSON(200, global.ReplyError(num, ctx))
	}
	//检验pc域名是否合法
	ok := global.DomainCheck(siteDomain.Domain)
	if !ok {
		return ctx.JSON(200, global.ReplyError(60013, ctx))
	}
	//检验三个域名是否被使用
	num, err = comboBeen.CheckThirdDoMainChange(
		siteDomain.SiteId, siteDomain.SiteIndexId,
		siteDomain.Domain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != 0 {
		return ctx.JSON(200, global.ReplyError(num, ctx))
	}
	//修改[判断是否存在，不存在就进行添加]
	count, err := siteOperateBean.ChangeSiteDomain(siteDomain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//站点详情
func (*SiteController) SiteInfoByDomainAndInfo(ctx echo.Context) error {
	site := new(input.SiteDomainInfo)
	code := global.ValidRequest(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, has, err := siteOperateBean.SiteInfoBySiteAndSiteIndex(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//站点添加
func (c *SiteController) PostSiteChildAdd(ctx echo.Context) error {
	addSite := new(input.AddSiteIn) //增加站点的请求你参数struct
	code := global.ValidRequestAdmin(addSite, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看站点是否存在
	info, flag, err := siteOperateBean.GetInfoBySiteIndexId(addSite.Site, addSite.SiteIndex)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag && info.DeleteTime == 0 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	//站点名称是否已经使用的校验
	flag, err = siteOperateBean.GetSingleSiteByName(addSite.SiteName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag {
		return ctx.JSON(200, global.ReplyError(60002, ctx))
	}
	//检验域名是否合法
	ok := global.DomainCheck(addSite.Domain)
	if !ok {
		return ctx.JSON(200, global.ReplyError(60013, ctx))
	}
	//检验三个域名是否被使用
	num, err := comboBeen.CheckThirdDoMain(addSite.Domain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != 0 {
		return ctx.JSON(200, global.ReplyError(num, ctx))
	}
	//增加
	count, err := siteOperateBean.MoreSiteAdd(addSite)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60060, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//初始化密码
func (*SiteController) PutInitPassword(ctx echo.Context) error {
	admin := new(input.InitPassword)
	code := global.ValidRequestAdmin(admin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//MD5加密
	md5Password, err := global.MD5ByStr("123456", global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(30044, ctx))
	}
	//查看开户人id是否存在
	has, err := siteOperateBean.GetAccountById(admin.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30044, ctx))
	}
	count, err := siteOperateBean.InitPassword(admin.Id, md5Password)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//站点管理-数据-管理员-添加管理员(初始密码为123456)
func (c *SiteController) PostSiteAdmin(ctx echo.Context) error {
	account := new(input.AddAccount)
	code := global.ValidRequestAdmin(account, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	accountNameId := new(input.AccountNameId)
	accountNameId.Account = account.Account
	//查看帐号是否存在
	_, has, err := agencyBean.GetOneAccountHelder(accountNameId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//查看站点是否存在
	siteSchema := new(schema.Site)
	siteSchema.Id = account.Site
	has, err = siteOperateBean.GetOneSiteId(siteSchema)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//查看站点下是否有开户人
	has, err = siteOperateBean.GetAccount(account.Site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30267, ctx))
	}
	//密码加密
	md5Password, err := global.MD5ByStr("123456", global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(30044, ctx))
	}
	account.Password = md5Password
	//添加帐号
	count, err := agencyBean.AddAccount(account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}
