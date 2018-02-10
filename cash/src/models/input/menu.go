package input

//添加菜单
type AddMenu struct {
	MenuName    string `json:"menuName" valid:"Required;MaxSize(15);ErrorCode(50052)"` //菜单名称
	Route       string `json:"route" valid:"Required;MaxSize(50);ErrorCode(50053)"`    //菜单路由
	Level       int8   `json:"level" valid:"Min(0);ErrorCode(50138)"`                  //菜单等级
	ParentId    int64  `json:"parentId" valid:"Min(0);ErrorCode(50055)"`               //父级id
	LanguageKey string `json:"languageKey" valid:"MaxSize(30);ErrorCode(50139)"`       // 前端国际化标识
	Type        string `json:"type" valid:"Required;MaxSize(6);ErrorCode(50114)"`      // 菜单类型(agency代理,admin平台)
	Sort        int64  `json:"sort" valid:"Min(1);ErrorCode(70009)"`                   //菜单排序
	Icon        string `json:"icon" valid:"MaxSize(100);ErrorCode(50056)"`             //菜单icon
}

//菜单详情
type MenuInfo struct {
	Id       int64  `query:"id" valid:"Min(0);ErrorCode(50057)"`                //菜单id
	Type     string `query:"type" valid:"Required;MaxSize(6);ErrorCode(50114)"` // 菜单类型(agency代理,admin平台)
	MenuName string `query:"menuName" valid:"MaxSize(20);ErrorCode(50052)"`     //菜单名称
	Level    int8   `query:"level" valid:"Range(0,3);ErrorCode(50138)"`         //菜单等级
}

//修改菜单
type UpdataMenu struct {
	Id          int64  `json:"id" valid:"Required;ErrorCode(50057)"`
	MenuName    string `json:"menuName" valid:"Required;MaxSize(15);ErrorCode(50052)"`    //菜单名称
	ParentId    int64  `json:"parentId" valid:"Min(0);ErrorCode(50055)"`                  //父级id
	Route       string `json:"route" valid:"Required;MaxSize(50);ErrorCode(50053)"`       //菜单路由
	LanguageKey string `json:"languageKey" valid:"Required;MaxSize(30);ErrorCode(50139)"` // 前端国际化标识
	Sort        int64  `json:"sort" valid:"Min(1);ErrorCode(70009)"`                      //菜单排序
	Icon        string `json:"icon" valid:"Required;MaxSize(100);ErrorCode(50056)"`
}

//菜单禁用
type MenuOpenClose struct {
	Id     int64 `json:"id" valid:"Min(1);ErrorCode(50057)"`         //菜单id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(30050)"` //转递过来的是现在所处的状态
}

//菜单删除
type MenuDel struct {
	Id int64 `json:"id" valid:"Min(1);ErrorCode(50057)"` //菜单id
}
