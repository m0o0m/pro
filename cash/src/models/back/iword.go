package back

import "html/template"

type IwordList struct {
	Id          int64       `json:"id"`            //
	TopId       int64       `json:"top_id"`        //上级栏目
	SiteId      string      `json:"site_id"`       //
	SiteIndexId string      `json:"site_index_id"` //
	Title       string      `json:"title"`         //标题
	TitleColor  string      `json:"title_color"`   //标题颜色
	Content     string      `json:"content"`       //内容
	Url         string      `json:"url"`           //链接地址
	Img         string      `json:"img"`           //图片路径
	State       int8        `json:"state"`         //状态 1-启用 2-关闭
	Sort        int64       `json:"sort"`          //排序
	From        int8        `json:"from"`          //1-PC 2-WAP
	Itype       int64       `json:"itype"`         //类型代码
	TypeName    string      `json:"type_name"`     //类型名称
	AddTime     int64       `json:"add_time"`      //操作时间
	IList       []IwordList `json:"ilist"`         //子级菜单
	ContentHtml template.HTML
}

type AgreeContent struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`   //标题
	Content string `json:"content"` //内容
}
