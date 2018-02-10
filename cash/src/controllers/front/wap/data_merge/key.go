package data_merge

import (
	"framework/render"
	"strings"
)

type Key struct {
}

//data_merge:得到页面名称组合,作为缓存的key
func (m *Key) GetPageCacheKey(siteId, siteIndexId string, pageData interface{}) string {
	wapPageData, _ := pageData.(render.WapPageData)
	var keys = []string{siteId, siteIndexId}
	pages := wapPageData.GetPage()
	keys = append(keys, pages...)
	return strings.Join(keys, "$")
}
