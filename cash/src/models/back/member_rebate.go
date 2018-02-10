package back

//推广返佣比例列表
type WapMemberRebate struct {
	ValidMoney  int64   `json:"valid_money"`  //有效总投注
	ProductName string  `json:"product_name"` //商品名
	Rate        float64 `json:"rate"`         //比例
}

//推广返佣比例列表返回
type WapMemberRebateBack struct {
	ValidMoney  int64         `json:"validMoney"`  //有效总投注
	ProductRate []ProductRate `json:"productRate"` //商品比例
}

//商品对应的比例
type ProductRate struct {
	ProductName string  `json:"product_name"` //商品名
	Rate        float64 `json:"rate"`         //比例
}

//wap 推广人数系数
type WapRebateRanking struct {
	RankingNum float64 `json:"rankingNum"` //排行榜人数系数
}

//wap 排行榜
type WapRanking struct {
	Account         string `json:"account"`         //账号
	PromotionNumber int64  `json:"promotionNumber"` //推广人数
}
