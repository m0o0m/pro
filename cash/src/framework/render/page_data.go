package render

type PageData interface {
	//得到页面数据
	GetData(siteId, siteIndexId string) (interface{}, error)
	//得到页面html文件
	GetPage() []string //为了兼容,不改为GetPcPage
}
type CachePageData interface {
	PageData
	//得到缓存的key
	GetPageCacheKey(siteId, siteIndexId string, pageData interface{}) string
}

//lottery接口
type LotteryPageData interface {
	PageData
}

//wap的接口
type WapPageData interface {
	CachePageData
}

//pc的接口
type PcPageData interface {
	CachePageData
	//得到<视讯电子彩票体育的页面>
	GetSubPage() map[string]string
	GetThemeName(siteId, siteIndexId string) (themeDirName string, err error)
	GetSubThemeName(siteId, siteIndexId, moduleType string) (themeDirName string, err error)
}
