package input

//现金记录查询
type MemberCashRecord struct {
	SiteId     string `query:"siteId"`
	Account    string `query:"account"`    //会员账号
	OrderId    string `query:"orderId"`    //注单号
	StartTime  string `query:"startTime"`  //注册时间,开始
	EndTime    string `query:"endTime"`    //注册时间,结束
	SourceType int8   `query:"sourceType"` //数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换9优惠返水10自助返水11会员返佣
	ClientType int8   `query:"clientType"` //客户端类型1pc 2wap 3android 4ios
	PageSize   int    `query:"pageSize"`   //每页显示
	Page       int    `query:"page"`       //页码
}

//wap交易记录查询
type WapMemberCashRecord struct {
	Account       string `query:"account"`         //会员账号
	StartTime     int64  `query:"start_time"`      //注册时间,开始
	EndTime       int64  `query:"end_time"`        //注册时间,结束
	OrderNum      int64  `query:"order_num"`       //订单号
	SourceType    int64  `query:"source_type"`     //数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换9优惠返水10自助返水11会员返佣
	SourceOneType int64  `query:"source_one_type"` //二级分类
	PageSize      int    `query:"pageSize"`        //每页显示
	Page          int    `query:"page"`            //页码
}

//批量取消或者删除 现金报表
type PutMemberCashRecord struct {
	Ids []int64 `json:"ids"`
}
