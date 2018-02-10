//[控制器] [平台]所有模块管理
package note

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//模块管理
type GameModuleController struct {
	controllers.BaseController
}

//所有游戏种类查询
func (c *GameModuleController) GetGameTypeList(ctx echo.Context) error {

	productType := new(input.ProductTypeList)
	code := global.ValidRequestAdmin(productType, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	data, err := noteModuleBean.ProductTypeList(productType)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10001, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//添加游戏种类
func (c *GameModuleController) PostGameTypeAdd(ctx echo.Context) error {
	productType := new(input.ProductType)
	code := global.ValidRequestAdmin(productType, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断该商品分类是否已经存在该商品名称
	_, have, err := noteModuleBean.GetInfo(productType.Title)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10082, ctx))
	}
	count, err := noteModuleBean.Add(productType)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10083, ctx))
	}
	return ctx.NoContent(204)
}

//游戏种类修改
func (c *GameModuleController) PutGameTypeUpdate(ctx echo.Context) error {
	productType := new(input.ProductTypeEdit)
	code := global.ValidRequestAdmin(productType, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类Id是否存在
	_, have, err := noteModuleBean.GetInfoById(productType.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10085, ctx))
	}
	//判断修改后的商品分类名称是否存在
	data, have, err := noteModuleBean.GetInfo(productType.Title)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		if data.Id != productType.Id {
			return ctx.JSON(200, global.ReplyError(10082, ctx))
		}
	}
	//禁用商品分类需要进行判断
	if productType.Status == 2 {
		/*
			判断该商品分类下是否有商品，如果有提示必须一级一级的往下
			进行删除操作
		*/
		have, err = noteModuleBean.ExistProduct(productType.Id)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(10001, ctx))
		}
		if have {
			return ctx.JSON(200, global.ReplyError(10095, ctx))
		}
	}
	count, err := noteModuleBean.Update(productType)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//游戏种类删除
func (c *GameModuleController) DelGameType(ctx echo.Context) error {
	productTypeInfo := new(input.ProductTypeInfo)
	code := global.ValidRequestAdmin(productTypeInfo, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类Id是否存在
	_, have, err := noteModuleBean.GetInfoById(productTypeInfo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10085, ctx))
	}
	/*
		判断该商品分类下是否有商品，如果有提示必须一级一级的往下
		进行删除操作
	*/
	have, err = noteModuleBean.ExistProduct(productTypeInfo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10095, ctx))
	}
	count, err := noteModuleBean.Delete(productTypeInfo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//游戏平台查询
func (c *GameModuleController) GetGamePlatformList(ctx echo.Context) error {
	GamePlatform := new(input.GamePlatform)
	code := global.ValidRequestAdmin(GamePlatform, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := noteModuleBean.GetGamePlatformList(GamePlatform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//添加游戏平台
func (c *GameModuleController) PostGamePlatformAdd(ctx echo.Context) error {
	GamePlatform := new(input.PlatformAdd)
	code := global.ValidRequestAdmin(GamePlatform, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询添加的平台是否存在
	has, err := noteModuleBean.GetPlatformInfo(GamePlatform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(91001, ctx))
	}
	data, err := noteModuleBean.AddPlatform(GamePlatform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(91002, ctx))
	}
	return ctx.NoContent(204)
}

//游戏平台修改
func (c *GameModuleController) PutGamePlatformUpdate(ctx echo.Context) error {
	GamePlatform := new(input.PlatformUpdate)
	code := global.ValidRequestAdmin(GamePlatform, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := noteModuleBean.UpdatePlatform(GamePlatform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(91003, ctx))
	}
	return ctx.NoContent(204)
}

//游戏平台删除
func (c *GameModuleController) DelGamePlatform(ctx echo.Context) error {
	GamePlatform := new(input.PlatformDelete)
	code := global.ValidRequestAdmin(GamePlatform, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询平台下是否有正常运行的游戏
	has, err := noteModuleBean.GetPlatformGame(GamePlatform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(91004, ctx))
	}
	data, err := noteModuleBean.DeletePlatformGame(GamePlatform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(91005, ctx))
	}
	return ctx.NoContent(204)
}

//具体游戏查询
func (c *GameModuleController) GetGameProductList(ctx echo.Context) error {
	product := new(input.ProductList)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)
	data, count, err := noteModuleBean.ProductList(product, listParams)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(10001, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//具体游戏添加
func (c *GameModuleController) PostGameProductAdd(ctx echo.Context) error {
	product := new(input.Product)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类是否存在
	_, have, err := noteModuleBean.GetInfoById(product.TypeId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10085, ctx))
	}
	//判断该商品分类下的商品名是否已经存在
	_, have, err = noteModuleBean.Exist(product.TypeId, product.ProductName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10089, ctx))
	}
	//添加商品
	count, err := noteModuleBean.AddProduct(product)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10090, ctx))
	}
	return ctx.NoContent(204)
}

//游戏修改
func (c *GameModuleController) PutGameProductUpdate(ctx echo.Context) error {
	product := new(input.ProductEdit)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类是否存在
	_, have, err := noteModuleBean.GetInfoById(product.TypeId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10085, ctx))
	}
	//判断商品是否存在
	_, have, err = noteModuleBean.GetProductInfoById(product.ProductId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10092, ctx))
	}
	//判断修改后的商品名称是否重复
	reply, have, err := noteModuleBean.Exist(product.TypeId, product.ProductName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		if reply.Id != product.ProductId {
			return ctx.JSON(200, global.ReplyError(10089, ctx))
		}
	}
	//判断商品是否在被套餐使用
	if product.Status == 2 {
		have, err := noteModuleBean.GetProductId(product.ProductId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(10001, ctx))
		}
		if have {
			return ctx.JSON(200, global.ReplyError(10094, ctx))
		}
	}
	count, err := noteModuleBean.UpdateProduct(product)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//游戏删除
func (c *GameModuleController) DelGameProduct(ctx echo.Context) error {
	product := new(input.ProductInfo)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品是否存在
	data, have, err := noteModuleBean.GetProductInfoById(product.ProductId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10092, ctx))
	}

	//判断商品是否在被套餐使用
	if data.Status == 1 {
		have, err := noteModuleBean.GetProductId(product.ProductId)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err)
			return ctx.JSON(500, global.ReplyError(10001, ctx))
		}
		if have {
			return ctx.JSON(200, global.ReplyError(10094, ctx))
		}
	}
	//删除商品
	count, err := noteModuleBean.DeleteProduct(product.ProductId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10093, ctx))
	}
	return ctx.NoContent(204)
}
