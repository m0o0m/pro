package render

import (
	"bytes"
	"config"
	"errors"
	"fmt"
	"github.com/bluele/gcache"
	"github.com/labstack/echo"
	"global"
	"html/template"
	"strings"
	"sync"
)

var (
	MongoCacheSwitch bool                 //mongo缓存html文件,默认不开启
	sourcePath       = ""                 //模板根目录,在配置文件中配置
	ViewPath         = "template/"        //模板文件全在该目录下
	ModuleViewPath   = "template/public/" //
	DzViewPath       = "games/"           //电子
	FcViewPath       = "lottery/"         //彩票
	SpViewPath       = "sports/"          //体育
	VdViewPath       = "video/"           //视讯
	YhViewPath       = "youhui/"          //优惠

	templates    *template.Template //template.Must(template.ParseGlob("../bin/template/public/error.html"))
	pageCache    gcache.Cache       //缓存页面的容器
	errPageCache sync.Map           //错误因为数量少,就使用map缓存
	CdnUrl       string             //缓存地址
	CacheSwitch  string             //缓存开关
)

func InitRootPath(config *config.TemplateConfig) {

	if !strings.HasSuffix(config.SourcePath, "/") && config.SourcePath != "" {
		sourcePath = config.SourcePath + "/"
	} else {
		sourcePath = config.SourcePath
	}
	ViewPath = sourcePath + ViewPath             //模板文件全在该目录下
	ModuleViewPath = sourcePath + ModuleViewPath //
	DzViewPath = ModuleViewPath + DzViewPath     //电子
	FcViewPath = ModuleViewPath + FcViewPath     //彩票
	SpViewPath = ModuleViewPath + SpViewPath     //体育
	VdViewPath = ModuleViewPath + VdViewPath     //视讯
	YhViewPath = ModuleViewPath + YhViewPath     //优惠
	CdnUrl = config.CdnUrl
	fmt.Println("template:", ViewPath)
	//templatePath := global.GetTemplatePath()
	templates = template.Must(template.ParseGlob(ViewPath + "public/error.html"))
	CacheSwitch = config.CacheSwitch
	InitPageCache(config.CacheSize)
	if config.MongoCache == "on" {
		MongoCacheSwitch = true
	}
}

//初始化缓存大小
func InitPageCache(cacheSize int) {
	if pageCache == nil {
		if cacheSize == 0 {
			cacheSize = 100 //默认100个的缓存数量
		}
		pageCache = gcache.New(cacheSize).ARC().Build()
	}
}

//查询是否存在缓存,有就直接返回
func GetCache(siteId, siteIndexId string, pageData CachePageData) ([]byte, bool) {
	key := pageData.GetPageCacheKey(siteId, siteIndexId, pageData)

	byteBuff, err := pageCache.Get(key)
	if err != nil {
		return nil, false
	}
	bytes, _ := byteBuff.([]byte)
	return bytes, true
}

//清除单站点页面缓存
func DelCache(siteId, siteIndexId string, pageData interface{}) error {
	var templateKey string
	switch pageData.(type) {
	case PcPageData:
		pcPageData := pageData.(PcPageData)
		templateKey = pcPageData.GetPageCacheKey(siteId, siteIndexId, pageData)
	case WapPageData:
		wapPageData := pageData.(WapPageData)
		templateKey = wapPageData.GetPageCacheKey(siteId, siteIndexId, pageData)
	default:
		return errors.New("interface not in (render,PcPageData or render.WapPageData)")
	}
	pageCache.Remove(templateKey)
	return nil
}

//生成pc网页模板数据并缓存
func GenPcCache(siteId, siteIndexId string, pageData PcPageData) (bs []byte, code int64) {
	themeDirName, err := pageData.GetThemeName(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("get theme name %s", err.Error())
		code = 71028
		return
	}
	data, err := pageData.GetData(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("get pc data %s", err.Error())
		code = 71029
		return
	}
	//页面
	pages := pageData.GetPage()
	for k, _ := range pages {
		pages[k] = ViewPath + themeDirName + "/" + pages[k] + ".html"
	}
	//子页面<视讯电子彩票体育的>
	subPageMap := pageData.GetSubPage()
	if len(subPageMap) != 0 {
		dz, ok := subPageMap[DzViewPath]
		if ok {
			subThemeDirName, err := pageData.GetSubThemeName(siteId, siteIndexId, DzViewPath)
			if err != nil {
				global.GlobalLogger.Error("get pc sub theme %s", err.Error())
				return nil, 71030
			}
			pages = append(pages, DzViewPath+subThemeDirName+"/"+dz+".html")
		}
		fc, ok := subPageMap[FcViewPath]
		if ok {
			subThemeDirName, err := pageData.GetSubThemeName(siteId, siteIndexId, FcViewPath)
			if err != nil {
				global.GlobalLogger.Error("get pc sub theme %s", err.Error())
				return nil, 71031
			}
			pages = append(pages, FcViewPath+subThemeDirName+"/"+fc+".html")
		}
		sp, ok := subPageMap[SpViewPath]
		if ok {
			subThemeDirName, err := pageData.GetSubThemeName(siteId, siteIndexId, SpViewPath)
			if err != nil {
				global.GlobalLogger.Error("get pc sub theme %s", err.Error())
				return nil, 71032
			}
			pages = append(pages, SpViewPath+subThemeDirName+"/"+sp+".html")
		}
		vd, ok := subPageMap[VdViewPath]
		if ok {
			subThemeDirName, err := pageData.GetSubThemeName(siteId, siteIndexId, VdViewPath)
			if err != nil {
				global.GlobalLogger.Error("get pc sub theme %s", err.Error())
				return nil, 71033
			}
			pages = append(pages, VdViewPath+subThemeDirName+"/"+vd+".html")
		}
		yh, ok := subPageMap[YhViewPath]
		if ok {
			subThemeDirName, err := pageData.GetSubThemeName(siteId, siteIndexId, YhViewPath)
			if err != nil {
				global.GlobalLogger.Error("get pc sub theme %s", err.Error())
				return nil, 71033
			}
			pages = append(pages, YhViewPath+subThemeDirName+"/"+yh+".html")
		}

	}
	templateKey := pageData.GetPageCacheKey(siteId, siteIndexId, pageData)
	templates := template.Must(parseFiles(pages...))
	pageBytes := new(bytes.Buffer)
	err = templates.Execute(pageBytes, data)
	if err != nil {
		global.GlobalLogger.Error("错误:%s", err.Error())
		return nil, 72002
	}
	bs = pageBytes.Bytes()
	err = pageCache.Set(templateKey, bs)
	if err != nil {
		return nil, 71034
	}
	return bs, 0
}

//生成wap网页模板数据并缓存
func GenWapCache(siteId, siteIndexId string, pageData WapPageData) (bs []byte, code int64) {
	data, err := pageData.GetData(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("get data error:%s", err.Error())
		code = 71029
		return
	}
	//页面
	pages := pageData.GetPage()
	for k, _ := range pages {
		pages[k] = ViewPath + pages[k] + ".html"
	}
	templateKey := pageData.GetPageCacheKey(siteId, siteIndexId, pageData)
	fmt.Println(pages)
	templates := template.Must(parseFiles(pages...))
	pageBytes := new(bytes.Buffer)
	err = templates.Execute(pageBytes, data)
	if err != nil {
		global.GlobalLogger.Error("错误:%s", err.Error())
		return nil, 72002
	}
	bs = pageBytes.Bytes()
	err = pageCache.Set(templateKey, bs)
	if err != nil {
		global.GlobalLogger.Error("generate key err:%s", err.Error())
		return nil, 71034
	}
	return bs, 0
}

//生成彩票大厅页面
func GenLottery(siteId, siteIndexId string, pageData LotteryPageData) (bs []byte, code int64) {
	data, err := pageData.GetData(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("get data error:%s", err.Error())
		code = 71029
		return
	}
	//页面
	pages := pageData.GetPage()
	for k, _ := range pages {
		pages[k] = ViewPath + pages[k] + ".html"
	}
	fmt.Println(pages)
	templates := template.Must(template.ParseFiles(pages...))
	pageBytes := new(bytes.Buffer)
	err = templates.Execute(pageBytes, data)
	if err != nil {
		global.GlobalLogger.Error("错误:%s", err.Error())
		return nil, 72002
	}
	return pageBytes.Bytes(), 0
}

//根据错误码返回错误页面
func PageErr(code int64, ctx echo.Context) error {
	lan, ok := ctx.Get(global.TranslateLanguageHeaderKey).(string)
	if !ok {
		lan = global.DefaultLanguage
	}
	if len(lan) == 0 {
		lan = global.DefaultLanguage
	}
	if v, ok := global.ErrorCode[lan]; ok {
		if errInfo, ok := v[code]; ok {
			b, ok := errPageCache.Load(code)
			//缓存存在
			if ok {
				bs, _ := b.([]byte)
				return ctx.HTMLBlob(int(code), bs)
			} else {
				errBuff := new(bytes.Buffer)
				templates.Execute(errBuff, map[string]interface{}{"value": errInfo})
				errPageCache.Store(code, errBuff.Bytes())
				return ctx.HTMLBlob(int(code), errBuff.Bytes())
			}
		} else {
			panic("not found code")
		}
	} else {
		panic("not found language code")
	}
}
