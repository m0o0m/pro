package controllers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"global"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type BaseController struct {
}

/**
使用方式
type AccountController struct {
	controllers.BaseController
}

//代理登录
func (ac *AccountController) Login(ctx echo.Context) error {
	filters := make([]*controllers.ParamFilter, 0)
	filters = append(filters, controllers.NewParamFilter("username", "string", "", true, true, 0, 0, 1000000))

	params, code := ac.VerifyParam(ctx, filters)
	if code > 0 {
		return ctx.JSON(500, global.GetErrCode(code, ctx.Get(controllers.TranslateLanguageHeaderKey)))
	}
	//类型转换
	username := params[params]

	global.GlobalLogger.Error("111 %s", username)
	return ctx.JSON(200, "123123")
}



*/

//例子
//"notnull"  :true,
var params = map[string]map[string]interface{}{
	"mail": {
		"isrouteparam": false, //是否路由参数
		"mustexist":    true,  //参数是否必须
		"xss":          true,
		"type":         "int",
		"preg_match":   "正则表达式", //格式验证
		"lengthmin":    1,       //格式验证
		"lengthmax":    46,      //格式验证
		"error_code":   100010,  //格式验证
	},
}

type ParamFilter struct {
	Name         string
	IsRouteParam bool
	MustExist    bool
	Xss          bool
	Type         string
	PregMatch    string
	LengthMin    int
	LengthMax    int
	ErrorCode    int64
}

func NewParamFilter(Name, Type, PregMatch string, IsRouteParam, MustExist, Xss bool, LengthMin, LengthMax int, ErrorCode int64) *ParamFilter {
	np := new(ParamFilter)

	np.Name = Name
	np.IsRouteParam = IsRouteParam
	np.Type = Type
	np.PregMatch = PregMatch
	np.MustExist = MustExist
	np.Xss = Xss
	np.LengthMin = LengthMin
	np.LengthMax = LengthMax
	np.ErrorCode = ErrorCode

	return np
}

//参数过滤
//param := make(map[string]interface{})
//获取所有参数
//r.ParseForm()  get
//r.PostFormValue("id")
/*
	r.ParseMultipartForm(32 << 20)
	if r.MultipartForm != nil {
	    values := r.MultipartForm.Value["id"]
	    if len(values) > 0 {
		fmt.Fprintln(w, values[0])
	    }
	}
*/
func (*BaseController) VerifyParam(ctx echo.Context, paramsFilter []*ParamFilter) (param map[string]string, code int64) {
	param = make(map[string]string)
	method := strings.ToUpper(ctx.Request().Method)
	json_param := make(map[string]string)
	if method == "POST" {
		//统一处理
		if strings.Contains(ctx.Request().Header.Get("Content-Type"), "application/json") {
			buf, err := ioutil.ReadAll(ctx.Request().Body)
			if err != nil {
				code = 100000 //TODO
				return
			}
			err = json.Unmarshal(buf, &json_param)
			if err != nil {
				code = 100000
				return
			}
		}
	}
	//是否有 TranslateLanguageContextKey
	lang := ctx.Request().Header.Get(global.TranslateLanguageHeaderKey)
	//语言版本
	ctx.Set(global.TranslateLanguageHeaderKey, lang)
	for _, v := range paramsFilter {
		k := v.Name
		value := ""
		//值的处理
		switch method {
		case "PUT":
			value = strings.TrimSpace(ctx.FormValue(k))
		case "GET":
			value = strings.TrimSpace(ctx.QueryParam(k))
		case "POST":
			if strings.Contains(ctx.Request().Header.Get("Content-Type"), "application/json") {
				value = json_param[k]
			} else {
				value = strings.TrimSpace(ctx.FormValue(k))
			}
		case "DELETE":
			value = strings.TrimSpace(ctx.FormValue(k))
		}

		//检查参数是否是路由参数
		if v.IsRouteParam {
			value = strings.TrimSpace(ctx.Param(k))
		}

		//检查是否必须存在以及是否为空
		if v.MustExist && len(value) == 0 {
			code = v.ErrorCode
			return
		}
		if !v.MustExist && len(value) == 0 {
			continue
		}
		if v.Xss {
			value = xss(value)
		}
		//校验规则
		if !typesOf(v.Type, value) {
			code = v.ErrorCode
			return
		}
		//if !checkPrge(v.PregMatch, value) {
		//	code = v.ErrorCode
		//	return
		//}
		if b, _ := regexp.MatchString(v.PregMatch, value); !b {
			code = v.ErrorCode
			return
		}
		if v.LengthMin > len(value) {
			code = v.ErrorCode
			return
		}

		if v.LengthMax < len(value) {
			code = v.ErrorCode
			return
		}
		//赋值
		param[k] = value
	}
	return param, 0
}

func xss(p string) string {
	return p
}

//
//func checkPrge(z interface{}, value string) bool {
//	switch z.(string) {
//	case "email":
//		flag := global.CheckEmail(value)
//		if !flag {
//			return false
//		}
//		return true
//	case "identity":
//		flag := global.CheckIdentity(value)
//		if !flag {
//			return false
//		}
//		return true
//	case "cardnumber":
//		flag := global.CheckCardNumber(value)
//		if !flag {
//			return false
//		}
//		return true
//	case "qq":
//		flag := global.Checkqq(value)
//		if !flag {
//			return false
//		}
//		return true
//	case "phone":
//		flag := global.CheckPhoneNumber(value)
//		if !flag {
//			return false
//		}
//		return true
//	case "":
//		return true
//	}
//	return false
//}

func typesOf(t interface{}, value string) bool {
	switch t.(type) {
	case int:
		_, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return false
		}
		return true
	case float64:
		_, err := strconv.ParseFloat(value, 10)
		if err != nil {
			return false
		}
		return true
	case string:
		return true
	default:
		return false
	}
}

//获取listparam的值(分页使用)
func (*BaseController) GetParam(listparam *global.ListParams, ctx echo.Context) error {
	param := new(global.ListParams)
	//获取结构体的json数据
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(200, global.ReplyError(30111, ctx))
	}
	listparam.GetCount = param.GetCount //是否统计总数
	listparam.Desc = param.Desc
	listparam.OrderBy = param.OrderBy
	page := param.Page //页码
	if page < 1 {
		page = 1
	}
	listparam.Page = page
	pageSize := param.PageSize //每页数量
	if pageSize < 1 {
		pageSize = 15
	} else {
		if pageSize >= 50 {
			pageSize = 50
		}
	}
	listparam.PageSize = pageSize
	offset := (page - 1) * pageSize
	var limit = []int{pageSize, offset}
	listparam.Limit = limit
	return nil
}
