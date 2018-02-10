package input

//添加游戏种类、删除游戏种类
type VdGameType struct {
	Type string `query:"type" json:"type"` //类型
}

//修改游戏种类
type VdGameTypeUpdate struct {
	OldType string `json:"old_type"` //原类型
	NewType string `json:"new_type"` //新类型
}
