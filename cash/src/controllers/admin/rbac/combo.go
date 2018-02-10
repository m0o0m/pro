package rbac

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/schema"
)

//套餐管理
type ComboController struct {
	controllers.BaseController
}

//套餐列表(GET)
func (cc *ComboController) Index(ctx echo.Context) error {
	combo := new(input.ComboList)
	code := global.ValidRequestAdmin(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	cc.GetParam(listparam, ctx)
	list, count, err := comboBeen.GetProductList(combo, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(list) == 0 {
		//如果商品占比没配置，那就先查询套餐
		comboList, count, err := comboBeen.GetList(combo, listparam)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyPagination(listparam, comboList, int64(len(comboList)), count, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//添加套餐(POST)
func (cc *ComboController) ComboAdd(ctx echo.Context) error {
	combo := new(input.ComboAdd)
	code := global.ValidRequestAdmin(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看套餐名是否存在
	has, err := comboBeen.GetComboName(combo.ComboName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30100, ctx))
	}
	count, err := comboBeen.Add(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30089, ctx))
	}
	return ctx.NoContent(204)
}

//获取套餐信息(GET)
func (cc *ComboController) EditCombo(ctx echo.Context) error {
	combo := new(input.ComboId)
	code := global.ValidRequestAdmin(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, has, err := comboBeen.GetInfo(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30103, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//套餐启用/禁用(PUT)
func (cc *ComboController) Status(ctx echo.Context) error {
	combo := new(input.ComboStatus)
	code := global.ValidRequestAdmin(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看站点中是否使用此套餐
	has, err := siteOperateBean.GetComboId(combo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30108, ctx))
	}
	count, err := comboBeen.UpdateStatus(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30102, ctx))
	}
	return ctx.NoContent(204)
}

//修改套餐(PUT)
func (cc *ComboController) ComboEdit(ctx echo.Context) error {
	combo := new(input.ComboEdit)
	code := global.ValidRequestAdmin(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看套餐名是否被使用
	has, err := comboBeen.GetComboNames(combo.Id, combo.ComboName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30100, ctx))
	}
	count, err := comboBeen.UpdateInfo(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30089, ctx))
	}
	return ctx.NoContent(204)
}

//删除套餐(DELETE)
func (cc *ComboController) DeleteCombo(ctx echo.Context) error {
	combo := new(input.ComboId)
	code := global.ValidRequestAdmin(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看站点中是否使用此套餐
	has, err := siteOperateBean.GetComboId(combo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30109, ctx))
	}
	count, err := comboBeen.Delete(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30101, ctx))
	}
	return ctx.NoContent(204)
}

//获取套餐商品(GET)
func (cc *ComboController) ComboProduct(ctx echo.Context) error {
	comboProduct := new(input.ComboProductId)
	code := global.ValidRequestAdmin(comboProduct, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := comboBeen.GetProductInfo(comboProduct)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	product := new(schema.Product)
	data, err := productTypeBean.GetAllProductByType(product)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for k := range list {
		for i := range data {
			for j := range data[i].Children {
				if list[k].ProductId == data[i].Children[j].Id {
					data[i].Children[j].IsProduct = 1
					data[i].Children[j].Proportion = list[k].Proportion
				}
			}
		}
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//设置套餐商品(POST)
func (cc *ComboController) ComboProductEdit(ctx echo.Context) error {
	comboProduct := new(input.ComboProducts)
	code := global.ValidRequestAdmin(comboProduct, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//商品占比不能大于百分之百
	for k := range comboProduct.Params {
		if comboProduct.Params[k].Proportion > 100 {
			return ctx.JSON(200, global.ReplyError(30112, ctx))
		}
	}
	count, err := comboBeen.AddProductProportion(comboProduct)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30107, ctx))
	}
	return ctx.NoContent(204)
}

//根据商品类型名称模糊搜索商品类型名
func (cc *ComboController) GetProductType(ctx echo.Context) error {
	comboProduct := new(input.ProductTypeName)
	code := global.ValidRequestAdmin(comboProduct, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := comboBeen.GetProductByType(comboProduct)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data, err := comboBeen.GetProductType(comboProduct)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for k := range list {
		for i := range data {
			for j := range data[i].Children {
				if list[k].ProductId == data[i].Children[j].Id {
					data[i].Children[j].IsProduct = 1
					data[i].Children[j].Proportion = list[k].Proportion
				}
			}
		}
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
