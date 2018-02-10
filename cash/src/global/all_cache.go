package global

import (
	"github.com/bluele/gcache"
	"strings"
	"sync"
)

var (
	SiteModuleCache  sync.Map //站点维护的缓存
	ThemeCache       sync.Map //所有站点与其对应的皮肤文件夹name数据,比如,aaa$a = theme1
	ModuleThemeCache sync.Map //所有<视讯电子体育彩票>与其对应的皮肤文件的name数据,例如 aaa&a = default

	DepositCache    gcache.Cache //存款推送的载体,存款页面点击存款后,一直读取该map下是否有成功
	WithdrawalCache gcache.Cache //存款推送的载体,存款页面点击存款后,一直读取该map下是否有成功
)

func init() {
	DepositCache = gcache.New(1000).ARC().Build()
	WithdrawalCache = gcache.New(1000).ARC().Build()
}

func GenKey(args ...string) string {
	return strings.Join(args, "$")
}
