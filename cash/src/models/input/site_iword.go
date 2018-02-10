package input

//文案列表查询条件
type SiteCopyList struct {
	SiteId      string `query:"siteId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`      //站点前台id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Itype       int64  `query:"itype" valid:"Range(1,3);ErrorCode(90702)"`                           //文案类型
}

//文案列表查询条件
type SiteCopyListInfoOne struct {
	Id int64 `query:"id" valid:"Min(1);ErrorCode(30041)"` //文案类型
}

//文案列表查询条件
type SiteIWodList struct {
	SiteId      string
	SiteIndexId string `query:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Itype       int64  `query:"itype" valid:"Max(3);ErrorCode(90702)"`                               //文案类型
	State       int64  `query:"state"`                                                               //文案类型
}

//优惠文案列表查询条件
type SiteActivityCopyList struct {
	SiteIndexId string `query:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	From        int8   `query:"from" valid:"Range(1,2);ErrorCode(90710)"`                            //文案类别1-PC 2-WAP
}

//优惠文案详情查询条件
type SiteActivityCopyInfo struct {
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(90714)"`                         //文案id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	From        int8   `query:"from" valid:"Max(2);ErrorCode(90710)"`                                //文案类别1-PC 2-WAP
	TypeName    string `query:"typeName" valid:"MaxSize(100);ErrorCode(90712)" `                     //类型名称
	TopId       int64  `query:"topId" valid:"Max(11);ErrorCode(90703)"`                              //上级栏目
}

//文案列表单条查询条件
type SiteCopyInfo struct {
	Id    int64 `query:"id" valid:"Required;Min(0);ErrorCode(90714)"` //文案id
	Itype int64 `query:"itype" valid:"Min(1);ErrorCode(90702)"`       //文案类型
}

type SiteCopyAdd struct {
	SiteId      string `json:"siteId" valid:"RangeSize(1,4);ErrorCode(60105)"`      //操作站点id
	SiteIndexId string `json:"siteIndexId" valid:"RangeSize(1,4);ErrorCode(10050)"` //站点前台id
	//TopId       int64  `json:"topId" valid:"Min(0);ErrorCode(90703)"`                               //上级栏目
	Title string `json:"title" valid:"RangeSize(1,200);ErrorCode(90704)"` //标题
	//TitleColor  string `json:"titleColor" valid:"MaxSize(200);ErrorCode(90705)"`                    //标题颜色
	Content string `json:"content" ` //内容
	//Url         string `json:"url" valid:"MaxSize(200);ErrorCode(90706)" `                          //链接地址
	//Img         string `json:"img" valid:"MaxSize(200);ErrorCode(90707)" `                          //图片路径
	//State       int8   `json:"state" valid:"Max(2);ErrorCode(90708)"`                               //状态 1-启用 2-关闭
	Sort  int64 `json:"sort" valid:"Min(0);ErrorCode(90709)"`     //排序
	From  int8  `json:"from" valid:"Range(1,2);ErrorCode(90710)"` //文案类别1-PC 2-WAP
	Itype int64 `json:"itype" valid:"Min(1);ErrorCode(90711)"`    //类型代码
	//TypeName    string `json:"typeName" valid:"MaxSize(100);ErrorCode(90712)" `                     //类型名称
	//AddTime     int64  `json:"addTime"`                                                             //添加时间
}

////前台文案添加
//type SiteIwordAdd struct {
//	SiteId      string `json:"site_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       //操作站点id
//	SiteIndexId string `json:"site_index_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
//	Title       string `json:"title" valid:"MaxSize(200);ErrorCode(90704)"`                           //标题
//	Content     string `json:"content" `                                                              //内容
//	Sort        int64  `json:"sort" valid:"Max(3);ErrorCode(90709)"`                                  //排序
//	Itype       int64  `json:"itype" valid:"Max(3);ErrorCode(90711)"`                                 //类型代码
//	TypeName    string `json:"type_name" valid:"MaxSize(100);ErrorCode(90712)" `                      //类型名称
//}

//优惠文案添加
type SiteIwordActivityAdd struct {
	SiteId      string `json:"site_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       //操作站点id
	SiteIndexId string `json:"site_index_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Title       string `json:"title" valid:"MaxSize(200);ErrorCode(90704)"`                           //标题
	State       int8   `json:"state" valid:"Max(1);ErrorCode(90708)"`                                 //状态 1-启用 2-关闭
	Sort        int64  `json:"sort" valid:"Max(3);ErrorCode(90709)"`                                  //排序
	From        int8   `json:"from" valid:"Max(1);ErrorCode(90710)"`                                  //文案类别1-PC 2-WAP
	Ctype       int8   `json:"ctype" valid:"Max(1);ErrorCode(90710)"`                                 //文案类型 1-分类 2-内容
	TopId       int64  `json:"top_id" valid:"Max(11);ErrorCode(90703)"`                               //上级栏目
}

//轮播查询
type FlashList struct {
	SiteId      string `query:"siteId"`                                                              //操作站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//轮播查询详情
type FlashListInfo struct {
	Id int64 `query:"id" valid:"Min(1);ErrorCode(30041)"` //id
}

//轮播添加
type FlashAdd struct {
	SiteId      string `json:"siteId"`                                                              //操作站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	ImgTitle    string `json:"imgTitle" valid:"MaxSize(50);ErrorCode(90801)"`                       //标题
	ImgUrl      string `json:"imgUrl" valid:"MaxSize(200);ErrorCode(90802)"`                        //图片路径
	ImgLink     string `json:"imgLink" valid:"MaxSize(200);ErrorCode(90803)"`                       //链接地址
	State       int8   `json:"state" valid:"Range(1,2);ErrorCode(90804)"`                           //状态 1-启用 2-关闭
	Sort        int64  `json:"sort" valid:"Min(0);ErrorCode(90805)"`                                //排序
	Ftype       int8   `json:"ftype" valid:"Range(1,2);ErrorCode(90806)"`                           //类型 1-PC端 2-WAP端
}

//轮播图状态更改
type FlashStatus struct {
	SiteId      string `json:"siteId"`                                                              //操作站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Ftype       int8   `json:"ftype" valid:"Range(1,2);ErrorCode(90806)"`                           //类型 1-PC端 2-WAP端
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(300041)"`                                 //id
	Status      int8   `json:"status" valid:"Range(1,2);ErrorCode(30050)"`                          //当前状态
}

//站点logo图片管理
type LogoList struct {
	SiteId      string `query:"siteId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`      //操作站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//站点logo图片添加
type LogoAdd struct {
	SiteId      string `json:"siteId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`      //操作站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Title       string `json:"title" valid:"MaxSize(80);ErrorCode(90901)"`                          //logo名称
	LogoUrl     string `json:"logoUrl" valid:"MaxSize(200);ErrorCode(90902)"`                       //logo地址
	Type        int8   `json:"type" valid:"Min(1);ErrorCode(90711)"`                                //文案类型
	State       int8   `json:"state" valid:"Min(1);ErrorCode(90903)"`                               //状态1启用 2停用
	Form        int8   `json:"form" valid:"Min(1);ErrorCode(90903)"`                                //1pc 2wap
}

//文案修改
type CopyUpdate struct {
	Id          int64  `json:"id" valid:"Required;Min(1);Max(11);ErrorCode(90714)"`                 //文案id
	SiteId      string `json:"siteId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`      //操作站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Content     string `json:"content"`                                                             //内容
}

//文案修改
type IwordUpdate struct {
	Id      int64  `json:"id" valid:"Required;Min(1);Max(11);ErrorCode(90714)"` //文案id
	Itype   int64  `json:"itype" valid:"Min(1);ErrorCode(50114)"`
	Content string `json:"content"` //内容
	//Img     	string `json:"img"`                                                 	//图片路径
}

//优惠修改(标题类型修改)
type ActivityEditeTitle struct {
	Id          int64  `json:"id"`                                                         //文案id
	TopId       int64  `json:"topId" valid:"Max(11);ErrorCode(90703)"`                     //上级栏目
	SiteIndexId string `json:"siteIndexId" valid:"MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Title       string `json:"title" valid:"MaxSize(200);ErrorCode(90704)"`                //标题
	State       int8   `json:"state" valid:"Range(1,2);ErrorCode(90708)"`                  //状态 1-启用 2-关闭
	Sort        int64  `json:"sort" valid:"Max(3);ErrorCode(90709)"`                       //排序
	Ctype       int8   `json:"ctype" valid:"Max(1);ErrorCode(90710)"`                      //文案类型 1-分类 2-内容
	From        int8   `json:"from" valid:"Range(1,2);ErrorCode(90710)"`                   //文案类别1-PC 2-WAP
}

//优惠修改(标题类型修改)
type ActivityEditeContent struct {
	Id      int64  `json:"id" valid:"Required;Min(1);Max(11);ErrorCode(90714)"` //文案id
	Content string `json:"content"`                                             //内容
	Img     string `json:"img"`                                                 //图片路径
}

//优惠修改(标题内容)
type ActivityDel struct {
	Id int64 `json:"id" valid:"Required;Min(1);Max(11);ErrorCode(90714)"` //文案id
}

//站点线路检测查询
type SiteDetect struct {
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//站点线路检测修改

type SiteDetectEdit struct {
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(90717)"`                         //文案id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Content     string `json:"content"`                                                             //域名
	Protocol    int8   `json:"protocol" valid:"Range(1,2);ErrorCode(90718)"`                        //前缀 1 http: 2 https:
}

//站点线路检测删除
type SiteDetectDel struct {
	Id          int64  `json:"id" valid:"Required;Min(1);Max(11);ErrorCode(90717)"`                 //文案id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//站点线路检测修改

type SiteDetectAdd struct {
	SiteIndexId string `json:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Content     string `json:"content"`                                                             //域名
	Protocol    int8   `json:"protocol" valid:"Range(1,2);ErrorCode(90718)"`                        //前缀 1 http: 2 https:
}
