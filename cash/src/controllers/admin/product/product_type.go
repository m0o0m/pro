package product

import (
	"controllers"
	"global"
	"models/input"

	"github.com/labstack/echo"
)

type ProductController struct {
	controllers.BaseController
}

//添加商品(POST)
func (pc *ProductController) ProductAdd(ctx echo.Context) error {
	product := new(input.Product)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类是否存在
	_, have, err := productTypeBean.GetInfoById(product.TypeId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10085, ctx))
	}
	//判断该商品分类下的商品名是否已经存在
	_, have, err = productBean.Exist(product.TypeId, product.ProductName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10089, ctx))
	}
	//添加商品
	count, err := productBean.Add(product)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10090, ctx))
	}
	return ctx.NoContent(204)
}

//获取商品信息(GET)
func (pc *ProductController) EditProduct(ctx echo.Context) error {
	product := new(input.ProductInfo)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, have, err := productBean.GetInfo(product)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10092, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改商品(PUT)
func (pc *ProductController) ProductEdit(ctx echo.Context) error {
	product := new(input.ProductEdit)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类是否存在
	_, have, err := productTypeBean.GetInfoById(product.TypeId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10085, ctx))
	}
	//判断商品是否存在
	_, have, err = productBean.GetInfoById(product.ProductId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10092, ctx))
	}
	//判断修改后的商品名称是否重复
	reply, have, err := productBean.Exist(product.TypeId, product.ProductName)
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
		have, err := comboBeen.GetProductId(product.ProductId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(10001, ctx))
		}
		if have {
			return ctx.JSON(200, global.ReplyError(10094, ctx))
		}
	}
	count, err := productBean.Update(product)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//商品启用/禁用(PUT)
func (pc *ProductController) Status(ctx echo.Context) error {
	product := new(input.ProductStatus)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品是否存在
	data, have, err := productBean.GetInfoById(product.ProductId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10092, ctx))
	}
	//判断商品是否在被套餐使用
	if data.Status == 1 {
		have, err := comboBeen.GetProductId(product.ProductId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(10001, ctx))
		}
		if have {
			return ctx.JSON(200, global.ReplyError(10094, ctx))
		}
	}
	count, err := productBean.Status(product)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//删除商品(DELETE)
func (pc *ProductController) DeleteProduct(ctx echo.Context) error {
	product := new(input.ProductInfo)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品是否存在
	data, have, err := productBean.GetInfoById(product.ProductId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10092, ctx))
	}

	//判断商品是否在被套餐使用
	if data.Status == 1 {
		have, err := comboBeen.GetProductId(product.ProductId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(10001, ctx))
		}
		if have {
			return ctx.JSON(200, global.ReplyError(10094, ctx))
		}
	}
	//删除商品
	count, err := productBean.Delete(product.ProductId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10093, ctx))
	}
	return ctx.NoContent(204)
}

//商品列表(GET)
func (pc *ProductController) Index(ctx echo.Context) error {
	product := new(input.ProductList)
	code := global.ValidRequestAdmin(product, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	pc.GetParam(listparam, ctx)
	list, count, err := productBean.List(product, listparam)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10001, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//添加商品分类(POST)
func (pc *ProductController) ProductTypeAdd(ctx echo.Context) error {
	productType := new(input.ProductType)
	code := global.ValidRequestAdmin(productType, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断该商品分类是否已经存在该商品名称
	_, have, err := productTypeBean.GetInfo(productType.Title)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10082, ctx))
	}
	count, err := productTypeBean.Add(productType)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10083, ctx))
	}
	return ctx.NoContent(204)
}

//获取商品分类信息(GET)
func (pc *ProductController) EditProductType(ctx echo.Context) error {
	productTypeInfo := new(input.ProductTypeInfo)
	code := global.ValidRequestAdmin(productTypeInfo, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, have, err := productTypeBean.GetInfoById(productTypeInfo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10085, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改商品分类(PUT)
func (pc *ProductController) ProductTypeEdit(ctx echo.Context) error {
	productType := new(input.ProductTypeEdit)
	code := global.ValidRequestAdmin(productType, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类Id是否存在
	_, have, err := productTypeBean.GetInfoById(productType.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10085, ctx))
	}
	//判断修改后的商品分类名称是否存在
	data, have, err := productTypeBean.GetInfo(productType.Title)
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
		have, err = productTypeBean.ExistProduct(productType.Id)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(10001, ctx))
		}
		if have {
			return ctx.JSON(200, global.ReplyError(10095, ctx))
		}
	}
	count, err := productTypeBean.Update(productType)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//商品分类启用/禁用(PUT)
func (pc *ProductController) ProductTypeStatus(ctx echo.Context) error {
	productType := new(input.ProductTypeStatus)
	code := global.ValidRequestAdmin(productType, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类Id是否存在
	_, have, err := productTypeBean.GetInfoById(productType.Id)
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
	have, err = productTypeBean.ExistProduct(productType.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10095, ctx))
	}
	//修改商品分类的状态
	count, err := productTypeBean.Status(productType)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//删除商品分类(DELETE)
func (pc *ProductController) ProductTypeDel(ctx echo.Context) error {
	productTypeInfo := new(input.ProductTypeInfo)
	code := global.ValidRequestAdmin(productTypeInfo, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:%d", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断商品分类Id是否存在
	_, have, err := productTypeBean.GetInfoById(productTypeInfo.Id)
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
	have, err = productTypeBean.ExistProduct(productTypeInfo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10095, ctx))
	}
	count, err := productTypeBean.Delete(productTypeInfo.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//商品分类下拉框
func (pc *ProductController) ProductTypeList(ctx echo.Context) error {
	data, err := productTypeBean.List()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))

}

//商品分类列表
func (pc *ProductController) ProductTypeIndex(ctx echo.Context) error {
	productType := new(input.ProductTypeList)
	code := global.ValidRequestAdmin(productType, ctx)
	if code != 0 {
		global.GlobalLogger.Error("errCode:", code)
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	pc.GetParam(listParams, ctx)
	data, count, err := productTypeBean.ProductTypeList(productType, listParams)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10001, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//商品下拉框
func (pc *ProductController) ProductListDrop(ctx echo.Context) error {
	data, err := productTypeBean.ProductListBeanDrop()
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10001, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
