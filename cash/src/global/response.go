package global

import (
	"fmt"
	"framework/validation"
	"math"
	"reflect"
	"strconv"

	"github.com/labstack/echo"
)

type (
	//错误返回
	Err struct {
		Code int64  `json:"code"`
		Msg  string `json:"msg"`
	}
	//有错误,并有结果返回
	ErrItem struct {
		Code int64       `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	//单条数据返回
	Item struct {
		Data interface{} `json:"data"`
	}
	//集合返回
	Collection struct {
		Data interface{}    `json:"data"`
		Meta CollectionMeta `json:"meta"`
	}
	CollectionMeta struct {
		Count int64 `json:"count"` //集合总数
	}
	//集合分页返回
	Pagination struct {
		Data  interface{} `json:"data"`
		Meta  Meta        `json:"meta"`
		Links Links       `json:"links"`
	}
	//集合分页返回2
	Pagination2 struct {
		Data interface{} `json:"data"`
		Meta Meta        `json:"meta"`
	}
	Meta struct {
		Count           int64 `json:"count"`             //数据总数
		PageCount       int64 `json:"page_count"`        //页码总数
		CurrentPage     int64 `json:"current_page"`      //当前页码
		PageSize        int64 `json:"page_size"`         //每页数据数量
		CurrentPageSize int64 `json:"current_page_size"` //当前页数据数量
	}
	Links struct {
		First   string `json:"first"`   //首页链接
		Last    string `json:"last"`    //末页链接
		Current string `json:"current"` //当前链接
		Prev    string `json:"prev"`    //上一页链接
		Next    string `json:"next"`    //下一页链接
	}
)

var (
	//默认错误信息
	DefaultErr = Err{
		Code: 10000,
		Msg:  "Field Illegal",
	}
	//语言参数名称
	TranslateLanguageHeaderKey = "language"
	DefaultLanguage            = "zh"
)

//得到错误信息
func GetError(code int64, ctx echo.Context) string {
	lan, ok := ctx.Get(TranslateLanguageHeaderKey).(string)
	if !ok {
		lan = DefaultLanguage
	}
	if len(lan) == 0 {
		lan = DefaultLanguage
	}
	if v, ok := ErrorCode[lan]; ok {
		if errInfo, ok := v[code]; ok {
			return errInfo
		}
	}
	return ""
}

//返回错误信息
func ReplyError(code int64, ctx echo.Context) *Err {
	err := DefaultErr
	lan, ok := ctx.Get(TranslateLanguageHeaderKey).(string)
	if !ok {
		lan = DefaultLanguage
	}
	if len(lan) == 0 {
		lan = DefaultLanguage
	}
	if v, ok := ErrorCode[lan]; ok {
		if err_info, ok := v[code]; ok {
			err.Code = code
			err.Msg = err_info
		}
	}
	return &err
}

//返回单条数据
func ReplyItem(data interface{}) *Item {
	return &Item{Data: data}
}

//返回错误和结果
func ReplyErrItem(code int64, data interface{}, ctx echo.Context) *ErrItem {
	err := ErrItem{Data: data}
	lan, ok := ctx.Get(TranslateLanguageHeaderKey).(string)
	if !ok {
		lan = DefaultLanguage
	}
	if len(lan) == 0 {
		lan = DefaultLanguage
	}
	if v, ok := ErrorCode[lan]; ok {
		if err_info, ok := v[code]; ok {
			err.Code = code
			err.Msg = err_info
		}
	}
	return &err
}

//返回多个集合
func ReplyCollections(data ...interface{}) *Item {
	datas := make(map[string]interface{})
	if len(data) > 0 && len(data)%2 == 0 {
		for i := 0; i < len(data); i += 2 {
			datas[data[i].(string)] = data[i+1]
		}
	}
	//fmt.Printf(">>> %#v\n",datas)
	return &Item{Data: datas}
}

//返回集合
func ReplyCollection(data interface{}, count interface{}) *Collection {
	var temp int64
	switch count.(type) {
	case int:
		temp = int64(count.(int))
	case int8:
		temp = int64(count.(int8))
	case int16:
		temp = int64(count.(int16))
	case int32:
		temp = int64(count.(int32))
	case int64:
		temp = count.(int64)
	default:
		panic("count to int64 err .ReplyCollection")
	}

	return &Collection{
		Data: data,
		Meta: CollectionMeta{Count: temp},
	}
}

func ReplyPagination(listParams *ListParams, data interface{}, current_page_size int64, count int64, ctx echo.Context) *Pagination {
	var first_num, last_num, current_num, prev_num, next_num int64
	var page_host string = ctx.Scheme() + "://" + ctx.Request().Host
	first_num = 1

	//末页码/页码总数
	last_num = int64(math.Ceil(float64(count) / float64(listParams.PageSize)))
	//当前页码
	current_num = int64(listParams.Page)
	//上一页码
	prev_num = current_num - 1
	if prev_num <= 0 {
		prev_num = 0
	}
	next_num = current_num + 1
	if next_num > last_num {
		next_num = 0
	}

	links := new(Links)
	pre := page_host + ctx.Request().URL.Path + "?"

	first_url_map := ctx.QueryParams()
	first_url_map.Set("page", strconv.FormatInt(first_num, 10))
	links.First = pre + first_url_map.Encode()

	last_url_map := ctx.QueryParams()
	last_url_map.Set("page", strconv.FormatInt(last_num, 10))
	links.Last = pre + last_url_map.Encode()

	links.Current = page_host + ctx.Request().URL.String()

	if prev_num != 0 {
		prev_url_map := ctx.QueryParams()
		prev_url_map.Set("page", strconv.FormatInt(prev_num, 10))
		links.Prev = pre + prev_url_map.Encode()
	} else {
		links.Prev = ""
	}

	if next_num != 0 {
		next_url_map := ctx.QueryParams()
		next_url_map.Set("page", strconv.FormatInt(next_num, 10))
		links.Next = pre + next_url_map.Encode()
	} else {
		links.Next = ""
	}

	return &Pagination{
		Data: data,
		Meta: Meta{
			Count:           count,
			PageCount:       last_num,
			CurrentPage:     current_num,
			PageSize:        int64(listParams.PageSize),
			CurrentPageSize: current_page_size,
		},
		Links: *links,
	}
}
func ReplyPagination2(listParams *ListParams, data interface{}, current_page_size int64, count int64) *Pagination2 {
	var last_num, current_num, prev_num, next_num int64

	//末页码/页码总数
	last_num = int64(math.Ceil(float64(count) / float64(listParams.PageSize)))
	//当前页码
	current_num = int64(listParams.Page)
	//上一页码
	prev_num = current_num - 1
	if prev_num <= 0 {
		prev_num = 0
	}
	next_num = current_num + 1
	if next_num > last_num {
		next_num = 0
	}

	return &Pagination2{
		Data: data,
		Meta: Meta{
			Count:           count,
			PageCount:       last_num,
			CurrentPage:     current_num,
			PageSize:        int64(listParams.PageSize),
			CurrentPageSize: current_page_size,
		},
	}
}
func ValidRequest(column interface{}, ctx echo.Context) int64 {
	//请求参数绑定
	if err := ctx.Bind(column); err != nil {
		return 10000
	}

	//请求参数验证
	valid := validation.Validation{}
	ok, err := valid.Valid(column)
	if err != nil {
		return 60000
	}
	if !ok {
		for _, e := range valid.Errors {
			fmt.Println(e.Error())
			return e.Code()
		}
	}
	//设置site_id、site_index_id参数值
	objV := reflect.ValueOf(column).Elem()
	agency, ok := ctx.Get("user").(*RedisStruct)
	if ok {
		valSiteId := objV.FieldByName("SiteId")

		if valSiteId.IsValid() {
			valSiteId.SetString(agency.SiteId)
		}
		if agency.SiteIndexId != "" {
			valSiteIndexId := objV.FieldByName("SiteIndexId")

			if valSiteIndexId.IsValid() {
				valSiteIndexId.SetString(agency.SiteIndexId)
			}
		}
	}
	return 0
}

//平台管理员
func ValidRequestAdmin(column interface{}, ctx echo.Context) int64 {
	//请求参数绑定
	if err := ctx.Bind(column); err != nil {
		return 10000
	}
	//请求参数验证
	valid := validation.Validation{}
	ok, err := valid.Valid(column)
	if err != nil {
		return 60000
	}
	if !ok {
		for _, e := range valid.Errors {
			fmt.Println(e.Error())
			return e.Code()
		}
	}
	return 0
}

//会员请求参数校验绑定
func ValidRequestMember(column interface{}, ctx echo.Context) int64 {
	//请求参数绑定
	if err := ctx.Bind(column); err != nil {
		return 10000
	}

	//请求参数验证
	valid := validation.Validation{}
	ok, err := valid.Valid(column)
	if err != nil {
		return 60000
	}
	if !ok {
		for _, e := range valid.Errors {
			fmt.Println(e.Error())
			return e.Code()
		}
	}
	return 0
}
