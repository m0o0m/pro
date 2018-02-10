package data_merge

import (
	"errors"
	"framework/render"
	"global"
	"models/schema"
	"strings"
)

type Key struct {
}

const THEME_REDIS_KEY = "theme_redis_key"

//pc得到页面名称组合,作为缓存的key
func (*Key) GetPageCacheKey(siteId, siteIndexId string, pageData interface{}) string {
	pcPageData, _ := pageData.(render.PcPageData)
	var keys = []string{siteId, siteIndexId}
	pages := pcPageData.GetPage()
	keys = append(keys, pages...)
	subPageMap := pcPageData.GetSubPage()
	if len(subPageMap) != 0 {
		dz, ok := subPageMap[render.DzViewPath]
		if ok {
			keys = append(keys, dz)
		}
		fc, ok := subPageMap[render.FcViewPath]
		if ok {
			keys = append(keys, fc)
		}
		sp, ok := subPageMap[render.SpViewPath]
		if ok {
			keys = append(keys, sp)
		}
		vd, ok := subPageMap[render.VdViewPath]
		if ok {
			keys = append(keys, vd)
		}
	}
	return strings.Join(keys, "$")
}

//查询站点主题文件夹名称
func (*Key) GetThemeName(siteId, siteIndexId string) (themeDirName string, err error) {
	key := siteId + "$" + siteIndexId
	temp, ok := global.ThemeCache.Load(key)
	if ok {
		themeDirName, _ = temp.(string)
		return
	}
	//如果没有,就从redis查询
	themeDirName, err = getPagePathByRedis(siteId, siteIndexId)
	return
}

//从redis中获取对应主题下的页面路径 并缓存到内存中
func getPagePathByRedis(siteId, siteIndexId string) (themeDirName string, err error) {
	key := siteId + "$" + siteIndexId
	themeDirName, err = global.GetRedis().HGet(THEME_REDIS_KEY, key).Result()
	if err == nil {
		//将redis中查询的主题名放入内存
		global.ThemeCache.Store(key, themeDirName)
		return
	}
	//如果没有,就从mysql查询
	themeDirName, err = getPagePathByMysql(siteId, siteIndexId)
	return
}

//从数据库获取对应主题下的页面路径  并缓存到redis和内存中
func getPagePathByMysql(siteId, siteIndexId string) (themeDirName string, err error) {
	key := siteId + "$" + siteIndexId
	themeDirName, err = siteOperateBean.GetThemeBySiteId(siteId, siteIndexId)
	if err != nil {
		return
	}
	//将mysql查询的主题名放入redis
	global.GetRedis().HSet(THEME_REDIS_KEY, key, themeDirName)
	//将mysql查询的主题名放入内存
	global.ThemeCache.Store(key, themeDirName)
	return
}

//查询站点子主题<视讯电子彩票体育>文件夹名称 ,数据库目前缺少一个主题名称对应表,这里就先不做缓存了,用style列来判定
func (*Key) GetSubThemeName(siteId, siteIndexId, moduleType string) (themeDirName string, err error) {
	//key := siteId + "$" + siteIndexId
	//目前没确定表,先返回一个虚拟的
	if moduleType == render.VdViewPath {
		liveid, _ := liveStyle(siteId, siteIndexId)
		themeDirName = getLiveStyle(liveid)
		return
	} else {
		return "default", nil
	}
}

//获取视讯选择模版ID
func liveStyle(siteId, siteIndexId string) (LiveId int, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoLogoSchema := new(schema.InfoVideoUse)
	b, err := sess.Table(infoLogoSchema.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Where("state = ?", 1).
		Select("style").
		Get(&LiveId)
	if err != nil {
		return
	}
	if !b {
		err = errors.New("not found style")
	}
	return
}

func getLiveStyle(liveid int) (livestr string) {
	switch liveid {
	case 3:
		livestr = "default"
	case 4:
		livestr = "live2"
	case 5:
		livestr = "live3"
	case 7:
		livestr = "live4"
	case 8:
		livestr = "live5"
	case 9:
		livestr = "live6"
	default:
		livestr = "default"
	}
	return
}
