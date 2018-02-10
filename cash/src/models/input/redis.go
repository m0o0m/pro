package input

// 查询缓存
type SearchRedis struct {
	Keyword string `json:"keyword"` // 缓存的Key
}
