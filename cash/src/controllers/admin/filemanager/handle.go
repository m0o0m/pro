package filemanager

import (
	"controllers"
	"encoding/json"
	"framework/mongo"
	"github.com/labstack/echo"
	"global"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type FileManager struct {
	controllers.BaseController
}

type result struct {
	Result interface{} `json:"result"`
}

type resultStatus struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func (f *FileManager) Handle(ctx echo.Context) error {
	json_param := make(map[string]interface{})
	//解析json
	if strings.Contains(ctx.Request().Header.Get("Content-Type"), "application/json") {
		buf, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.JSONBlob(http.StatusOK, []byte("error1"))
		}
		err = json.Unmarshal(buf, &json_param)
		if err != nil {
			return ctx.JSONBlob(http.StatusOK, []byte("error2"))
		}
	} else if strings.Contains(ctx.Request().Header.Get("content-Type"), "multipart/form-data") {
		form, err := ctx.MultipartForm()
		if err != nil {
			return ctx.JSONBlob(http.StatusOK, []byte("error3"))
		}
		if path, ok := form.Value["destination"]; ok {
			data, _ := upload(path[0], form)
			return ctx.JSON(http.StatusOK, data)
		}
	}
	var data interface{}
	var err error
	switch mongo.ToStr(json_param["action"]) {
	case "list":
		data, err = list(mongo.ToStr(json_param["path"]))
	case "createFolder":
		data, err = createFolder(mongo.ToStr(json_param["newPath"]))
	case "getContent":
		data, err = getContent(mongo.ToStr(json_param["item"]))
	case "edit":
		data, err = edit(mongo.ToStr(json_param["item"]), mongo.ToStr(json_param["content"]))
	case "rename":
		data, err = rename(mongo.ToStr(json_param["item"]), mongo.ToStr(json_param["newItemPath"]))
	case "remove":
		data, err = remove(mongo.ToStr(json_param["items"]))
	case "move":
		data, err = move(mongo.ToStr(json_param["items"]), mongo.ToStr(json_param["newPath"]))
	default:
		global.GlobalLogger.Error("error:%s", "no action")
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return ctx.JSON(http.StatusOK, data)
}

func list(path string) (interface{}, error) {
	rs := new(result)
	if path == "/" {
		rs.Result = mongo.GetDatabase()
	} else {
		bp, err := mongo.ParsePath(path)
		if err != nil {
			return rs, err
		}
		sess, err := mongo.GetMongoSess()
		if err != nil {
			return rs, err
		}
		defer sess.Close()
		mf, err := mongo.NewMongodbFile(sess, bp.Database, bp.Table)
		if err != nil {
			return rs, err
		}
		rs.Result, _ = mf.List(bp)
	}
	return rs, nil
}

func createFolder(path string) (interface{}, error) {
	rs := new(result)
	rsts := new(resultStatus)
	bp, err := mongo.ParsePath(path)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, nil
	}
	if bp.Table == "" {
		rsts.Error = "cannot create database"
		rsts.Success = false
		rs.Result = rsts
		return rs, nil
	} else {
		if bp.Database != "conf" && bp.Database != "templates" && bp.Database != "cache" {
			rsts.Error = "unknown database"
			rsts.Success = false
			rs.Result = rsts
			return rs, nil
		}
		sess, err := mongo.GetMongoSess()
		if err != nil {
			rsts.Error = err.Error()
			rsts.Success = false
			rs.Result = rsts
			return rs, err
		}
		defer sess.Close()
		mf, err := mongo.NewMongodbFile(sess, bp.Database, "")
		if err != nil {
			rsts.Error = err.Error()
			rsts.Success = false
			rs.Result = rsts
			return rs, err
		}
		err = mf.Create(bp)
		if err != nil {
			rsts.Error = err.Error()
			rsts.Success = false
			rs.Result = rsts
			return rs, err
		} else {
			rsts.Success = true
			rs.Result = rsts
			return rs, nil
		}
	}
}

func getContent(path string) (interface{}, error) {
	rs := new(result)
	bp, err := mongo.ParsePath(path)
	if err != nil {
		return rs, err
	}
	sess, err := mongo.GetMongoSess()
	if err != nil {
		return rs, err
	}
	defer sess.Close()
	mf, err := mongo.NewMongodbFile(sess, bp.Database, bp.Table)
	if err != nil {
		return rs, err
	}
	rs.Result, err = mf.Read(bp)
	return rs, err
}

func edit(path, content string) (interface{}, error) {
	rs := new(result)
	bp, err := mongo.ParsePath(path)
	if err != nil {
		return rs, err
	}
	sess, err := mongo.GetMongoSess()
	if err != nil {
		return rs, err
	}
	defer sess.Close()
	mf, err := mongo.NewMongodbFile(sess, bp.Database, bp.Table)
	if err != nil {
		return rs, err
	}
	err = mf.Write(bp, content)
	rsts := new(resultStatus)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
	} else {
		rsts.Success = true
		rs.Result = rsts
	}
	return rs, nil
}

func rename(fullName, newFullName string) (interface{}, error) {
	rs := new(result)
	rsts := new(resultStatus)
	bp, err := mongo.ParsePath(fullName)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	sess, err := mongo.GetMongoSess()
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	mf, err := mongo.NewMongodbFile(sess, bp.Database, "")
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	err = mf.Rename(bp, newFullName)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	} else {
		rsts.Success = true
		rs.Result = rsts
		return rs, nil
	}
}

func remove(paths string) (interface{}, error) {
	paths = strings.TrimLeft(paths, "[")
	paths = strings.TrimRight(paths, "]")
	items := strings.Split(paths, " ")
	rs := new(result)
	rsts := new(resultStatus)
	sess, err := mongo.GetMongoSess()
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	for _, v := range items {
		bp, err := mongo.ParsePath(v)
		if err != nil {
			rsts.Error = err.Error()
			rsts.Success = false
			rs.Result = rsts
			return rs, err
		}
		mf, err := mongo.NewMongodbFile(sess, bp.Database, "")
		if err != nil {
			rsts.Error = err.Error()
			rsts.Success = false
			rs.Result = rsts
			return rs, err
		}
		err = mf.Delete(bp)
		if err != nil {
			rsts.Error = err.Error()
			rsts.Success = false
			rs.Result = rsts
			return rs, err
		}
	}
	rsts.Success = true
	rs.Result = rsts
	return rs, nil
}

func move(paths, newPath string) (interface{}, error) {
	paths = strings.TrimLeft(paths, "[")
	paths = strings.TrimRight(paths, "]")
	items := strings.Split(paths, " ")
	rs := new(result)
	rsts := new(resultStatus)
	bp, err := mongo.ParsePath(newPath)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	//验证移动后的目录是否为database或者collection
	if bp.Table == "" {
		rsts.Error = "Cannot move to the database"
		rsts.Success = false
		rs.Result = rsts
		return rs, nil
	}
	for _, v := range items {
		vbp, err := mongo.ParsePath(v)
		if err != nil {
			rsts.Error = err.Error()
			rsts.Success = false
			rs.Result = rsts
			return rs, err
		}
		//验证是否同数据库之间移动
		if vbp.Database != bp.Database {
			rsts.Error = "It's not the same database, don't can move this"
			rsts.Success = false
			rs.Result = rsts
			return rs, nil
		}
		//验证要移动的路径是否是database或者collection
		p := strings.Split(strings.TrimLeft(v, "/"), "/")
		if len(p) < 3 {
			rsts.Error = "cannot move database or collection"
			rsts.Success = false
			rs.Result = rsts
			return rs, nil
		}
	}
	sess, err := mongo.GetMongoSess()
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	mf, err := mongo.NewMongodbFile(sess, bp.Database, "")
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	err = mf.Move(bp, items)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	} else {
		rsts.Success = true
		rs.Result = rsts
		return rs, nil
	}
}

func upload(path string, form *multipart.Form) (interface{}, error) {
	rs := new(result)
	rsts := new(resultStatus)
	bp, err := mongo.ParsePath(path)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	if bp.Database == "" || bp.Table == "" {
		rsts.Error = "Cannot upload to the directory"
		rsts.Success = false
		rs.Result = rsts
		return rs, nil
	}
	sess, err := mongo.GetMongoSess()
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	mf, err := mongo.NewMongodbFile(sess, bp.Database, bp.Table)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	err = mf.Upload(bp, form)
	if err != nil {
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	} else {
		rsts.Success = true
		rs.Result = rsts
		return rs, nil
	}
}
