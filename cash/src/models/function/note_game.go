package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
)

//电子管理
type NoteGameBean struct{}

//查询电子游戏列表
func (*NoteGameBean) GetVdGameList(this *input.VdGameList, listparam *global.ListParams) (data []back.VdGameList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)

	if len(this.Name) != 0 {
		sess.Where("name link ?", "%"+this.Name+"%")
	}
	sess.Where("type = ?", this.Type)

	conds := sess.Conds()
	listparam.Make(sess)

	err = sess.Table(mgGame.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(mgGame.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//查询电子游戏列表(不分页)
func (*NoteGameBean) GameListRedis() (data []schema.MgGame, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)
	err = sess.Table(mgGame.TableName()).Find(&data)
	return
}

//查询电子游戏列表(前台查询)
func (*NoteGameBean) IndexGameList(siteId, siteIndexId, Type string) (data []back.MgGame, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)
	siteGameDel := new(schema.SiteGameDel)
	ses := global.GetXorm().NewSession()

	var dataDel []schema.SiteGameDel //查询剔除的电子游戏
	err = ses.Table(siteGameDel.TableName()).
		Where("site_id = ? and site_index_id = ?", siteId, siteIndexId).
		Select("id").Find(&dataDel)
	var gameIdDel []int64
	for _, v := range dataDel {
		gameIdDel = append(gameIdDel, v.GameId)
	}
	if len(Type) != 0 {
		sess.Where("type = ?", strings.ToLower(Type))
	}

	sess.Desc("recommend") //排序

	sess.NotIn("id", gameIdDel)
	err = sess.Table(mgGame.TableName()).Find(&data)
	return
}

//查询电子游戏列表(wap查询)
func (*NoteGameBean) WapGameList(siteId, siteIndexId, Type string) (data []back.MgGame, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)
	if len(Type) != 0 {
		sess.Where("type = ?", strings.ToLower(Type)+"h5")
	} else {
		sess.Where("type = ?", "egh5")
	}
	sess.Where("status >= ?", 1)

	sess.NotIn("id", "select game_id id from sales_site_game_del where site_id = "+siteId+" and site_index_id "+siteIndexId)
	err = sess.Table(mgGame.TableName()).Find(&data)
	return
}

//根据电子游戏数据判断哪些平台有数据
func (*NoteGameBean) GameDataType() (data []back.MgGame, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)

	sess.Where("status >= ?", 1)
	sess.GroupBy("type")
	sess.Select("type")
	err = sess.Table(mgGame.TableName()).Find(&data)
	return
}

//查询电子游戏导航栏(前台查询)
func (*NoteGameBean) IndexGameTitle(siteId, siteIndexId string) (gameModel string, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteOrderModule := new(schema.SiteOrderModule)
	sess.Where("site_id = ?", siteId)
	sess.Where("site_index_id = ?", siteIndexId)
	_, err = sess.Table(siteOrderModule.TableName()).Get(siteOrderModule)
	if err != nil {
		return
	}
	gameModel = siteOrderModule.DzModule
	return
}

//添加电子游戏
func (*NoteGameBean) PostVdGameAdd(this *input.VdGameAdd) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)
	//检查是否存在此游戏  存在就不添加
	has, err := sess.Table(mgGame.TableName()).
		Where("type=? and gameid = ?", this.Type, this.Gameid).Get(mgGame)
	if has {
		code = 60300
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}

	mgGame.Type = this.Type
	mgGame.Name = this.Name
	mgGame.Topid = this.Topid
	mgGame.Itemid = this.Itemid
	mgGame.Gameid = this.Gameid
	mgGame.Image = this.Image
	mgGame.Status = 1
	mgGame.IsSw = 2
	mgGame.IsZs = 2
	//添加游戏
	count, err = sess.Table(mgGame.TableName()).InsertOne(mgGame)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//修改电子游戏
func (n *NoteGameBean) PutVdGameUpdate(this *input.VdGameUpdate) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)
	//检查是否存在此游戏  存在就不添加
	mgGame, has, err := n.GameInfo(this.Id)
	if !has {
		code = 60308
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}

	mgGame.Status = this.Status
	//修改电子游戏
	count, err = n.updata(*mgGame, "status")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//电子游戏修改（修改内容）
func (n *NoteGameBean) PutVdGameContentUpdate(this *input.VdGameContentUpdate) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)
	//检查是否存在此游戏  存在就不添加
	mgGame, has, err := n.GameInfo(this.Id)
	if !has {
		code = 60308
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}

	mgGame.Type = this.Type
	mgGame.Name = this.Name
	mgGame.Gameid = this.Gameid
	mgGame.Image = this.Image
	mgGame.Itemid = this.Itemid
	mgGame.Topid = this.Topid
	//修改电子游戏
	count, err = n.updata(*mgGame, "type,name,topid,itemid,gameid,image")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//排序推荐度修改
func (n *NoteGameBean) PutVdGameOrderUpdate(this *input.VdGameUpdate) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)
	//检查是否存在此游戏  存在就不添加
	mgGame, has, err := n.GameInfo(this.Id)
	if !has {
		code = 60308
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}

	if this.Recommend != 0 {
		mgGame.Recommend = this.Recommend
	} else {
		mgGame.Recommend = this.Ckr
	}

	//排序推荐度修改
	count, err = n.updata(*mgGame, "recommend")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//状态修改（剔除）
func (n *NoteGameBean) PutVdGameStatusUpdate(this *input.VdGameStatusUpdate) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteGameDel := new(schema.SiteGameDel)
	//检查是否存在此游戏  存在就不添加
	has, err := sess.Table(siteGameDel.TableName()).
		Where("site_id=? AND site_index_id = ? AND game_id = ?", this.SiteId, this.SiteIndexId, this.GameId).
		Get(siteGameDel)
	if has {
		code = 60300
		return code, count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}

	siteGameDel.SiteId = this.SiteId
	siteGameDel.SiteIndexId = this.SiteIndexId
	siteGameDel.GameId = this.GameId
	//剔除
	count, err = sess.Table(siteGameDel.TableName()).InsertOne(siteGameDel)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return code, count, err
	}
	return code, count, err
}

//排序推荐度修改
func (n *NoteGameBean) PutVdGameDemoSwitch(this *input.VdGameUpdate) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame := new(schema.MgGame)
	//检查是否存在此游戏  存在就不添加
	has, err := sess.Table(mgGame.TableName()).
		Where("id=?", this.Id).Get(mgGame)
	if !has {
		code = 60308
		return
	}
	if err != nil {
		return
	}

	var findStr string
	if this.IsZs != 0 {
		mgGame.IsZs = this.IsZs
		findStr = "is_zs"
	} else if this.IsSw != 0 {
		mgGame.IsSw = this.IsSw
		findStr = "is_sw"
	}

	//排序推荐度修改
	count, err = n.updata(*mgGame, findStr)
	return
}

//根据id查询电子游戏数据
func (*NoteGameBean) GameInfo(Id int64) (mgGame *schema.MgGame, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGame = new(schema.MgGame)
	has, err = sess.Table(mgGame.TableName()).
		Where("id=?", Id).Get(mgGame)
	return
}

//修改
func (*NoteGameBean) updata(mgGame schema.MgGame, fieldStr string) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id = ?", mgGame.Id)
	count, err = sess.Table(mgGame.TableName()).Cols(fieldStr).Update(mgGame)
	return
}

//添加游戏视讯类型
func (n *NoteGameBean) PostVdGameTypeAdd(this *input.VdGameTypeAdd) (code, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mgGameType := new(schema.MgGameType)
	//检查是否存在此游戏  存在就不添加
	has, err := sess.Table(mgGameType.TableName()).
		Where("type=?", this.Type).Get(mgGameType)
	if has {
		code = 60300
		return
	}
	if err != nil {
		return
	}

	mgGameType.Type = this.Type
	//添加游戏视讯类型
	count, err = sess.Table(mgGameType.TableName()).InsertOne(mgGameType)
	return
}

//电子内页主题设置
func (*NoteGameBean) GameTheme(siteId, siteIndexId string) (has bool, infoActivityPromotionSet *schema.InfoActivityPromotionSet, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	infoActivityPromotionSet = new(schema.InfoActivityPromotionSet)
	sess.Where("site_id = ?", siteId)
	sess.Where("site_index_id = ?", siteIndexId)
	has, err = sess.Table(infoActivityPromotionSet.TableName()).Get(infoActivityPromotionSet)
	return
}
