//[控制器] [平台]套餐管理
package report

import (
	"controllers"
	"fmt"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
)

//套餐管理
type PackageController struct {
	controllers.BaseController
}

//套餐配置列表查询
func (c *PackageController) GetNoticeList(ctx echo.Context) error {
	list := new(input.GetList)
	code := global.ValidRequestAdmin(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)
	data, err := sitePackageBean.GetList(list)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	comboList := make(map[int64]back.ReturnPackList)
	item := make(map[int64]bool)
	for _, v := range data {
		if item[v.ComboId] == false {
			comboListval := back.ReturnPackList{}
			comboListval.ComboId = v.ComboId
			comboListval.ComboName = v.ComboName
			for _, val := range data {
				if v.ComboId == val.ComboId && v.ProductId != 0 {
					comboListval.List = append(comboListval.List, val)
				}
			}
			item[v.ComboId] = true
			comboList[v.ComboId] = comboListval
		}
	}
	return ctx.JSON(200, global.ReplyItem(comboList))
}

//套餐添加
func (c *PackageController) PostPackageAdd(ctx echo.Context) error {
	sess := global.GetXorm().NewSession()
	list := new(input.PackAdd)
	code := global.ValidRequestAdmin(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	sess.Begin()
	add, data, err := sitePackageBean.PostPackageAdd(list)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	fmt.Println(add)
	for k := range list.List {
		list.List[k].ComboId = add.Id
	}
	adddata, err := sitePackageBean.AddProduct(list)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))

	}
	if adddata == 0 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	sess.Commit()
	return ctx.NoContent(204)
}

//套餐修改
func (c *PackageController) PutPackageUpdate(ctx echo.Context) error {
	updatedata := new(input.PackUpdate)
	code := global.ValidRequestAdmin(updatedata, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := sitePackageBean.PutPackageUpdate(updatedata)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(500, global.ReplyError(30102, ctx))
	}
	return ctx.NoContent(204)
}

//获取套餐详情
func (c *PackageController) GetPackage(ctx echo.Context) error {
	get_package := new(input.GetPackage)
	code := global.ValidRequestAdmin(get_package, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := sitePackageBean.GetPackageInfo(get_package)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//套餐启用/禁用(PUT)
func (c *PackageController) Status(ctx echo.Context) error {
	combo := new(input.ComboId)
	code := global.ValidRequestAdmin(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看站点中是否使用此套餐
	has, err := sitePackageBean.GetComboId(combo.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30108, ctx))
	}
	count, err := sitePackageBean.UpdateStatus(combo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30102, ctx))
	}
	return ctx.NoContent(204)
}

//套餐下拉框
func (c *PackageController) ComboDrop(ctx echo.Context) error {
	data, err := sitePackageBean.ComboDropAll()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
