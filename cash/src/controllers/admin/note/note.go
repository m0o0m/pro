//[控制器] [平台] 注单模块控制管理（合并彩票体育电子视讯）
package note

import (
	"models/function"
)

var noteModuleBean = new(function.NoteModuleBean)       //所有模块管理
var gameSetBean = new(function.GameSetBean)             //电子游戏管理
var gameWhitelistBean = new(function.GameWhitelistBean) //视讯ip白名单
var betRecordInfoBean = new(function.BetRecordInfoBean) //注单管理
var noteGameBean = new(function.NoteGameBean)           //电子游戏管理
