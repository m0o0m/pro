package mongo

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"controllers/admin/site"
	"errors"
	"fmt"
	"global"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	SESSION_IS_NIL    = errors.New("mgo.Session is nil")
	DATABASE_IS_NIL   = errors.New("mgo.Database is nil")
	COLLATION_IS_NIL  = errors.New("mgo.Collation is nil")
	TABLE_NAME_IS_NIL = errors.New("mgo.Collation table is nil")
	PATH_IS_ERROR     = errors.New("path is error")
)

type MongodbFile struct {
	sess       *mgo.Session    `json:"-"`
	database   *mgo.Database   `json:"-"`
	collection *mgo.Collection `json:"-"`
}

//mongo数据存储结构
type FileObject struct {
	Name     string `json:"name" bson:"name"`
	FullName string `json:"-" bson:"full_name"`
	FullPath string `json:"-" bson:"full_path"`
	Rights   string `json:"rights" bson:"rights"`
	Size     string `json:"size" bson:"size"`
	Ext      string `json:"-" bson:"ext"`
	Date     string `json:"date" bson:"date"`
	FileType string `json:"type" bson:"file_type"`
	Content  []byte `json:"-" bson:"content"`
	Status   int8   `json:"status" bson:"status"`
}

func GetDatabase() []FileObject {
	fo := make([]FileObject, 0)
	f1 := FileObject{}
	f1.Name = "conf"
	f1.Rights = "drwxr-xr-x"
	f1.Size = "0"
	f1.Date = "2017-03-03 15:31:40"
	f1.FileType = "dir"
	fo = append(fo, f1)

	f2 := FileObject{}
	f2.Name = "templates"
	f2.Rights = "drwxr-xr-x"
	f2.Size = "0"
	f2.Date = "2017-03-03 15:31:40"
	f2.FileType = "dir"
	fo = append(fo, f2)

	f3 := FileObject{}
	f3.Name = "cache"
	f3.Rights = "drwxr-xr-x"
	f3.Size = "0"
	f3.Date = "2017-03-03 15:31:40"
	f3.FileType = "dir"
	fo = append(fo, f3)
	return fo
}

type BasePath struct {
	Path     string
	Database string
	Table    string
}

//文件系统方法
type NetFile interface {
	Create(bp *BasePath) error
	Read(bp *BasePath) (string, error)
	Write(bp *BasePath, content string) error
	Delete(bp *BasePath) error
	Rename(bp *BasePath, newFullName string) error
	Move(bp *BasePath, moves []string) error
	List(bp *BasePath) (*[]FileObject, error)
	Upload(bp *BasePath, form *multipart.Form) error
}

func (mf *MongodbFile) Create(bp *BasePath) error {
	tables, err := mf.database.CollectionNames()
	if err != nil {
		return err
	}
	hasTable := false
	for _, v := range tables {
		if v == bp.Table {
			hasTable = true
			break
		}
	}
	mf.collection = mf.database.C(bp.Table)
	if !hasTable {
		err = mf.collection.Create(new(mgo.CollectionInfo))
		if err != nil {
			return err
		}
		idx := mgo.Index{}
		idx.Key = []string{"full_name"}
		idx.Unique = true
		err = mf.collection.EnsureIndex(idx)
		if err != nil {
			return err
		}
	}
	fos, err := ParsePathToFileObject(bp.Path)
	if err != nil {
		return err
	}
	for _, fo := range fos {
		_, err := mf.collection.Upsert(bson.M{"full_name": fo.FullName}, &fo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mf *MongodbFile) Read(bp *BasePath) (string, error) {
	fo := FileObject{}
	err := mf.collection.Find(bson.M{"full_name": bp.Path}).One(&fo)
	if err != nil {
		return "", err
	} else {
		return string(fo.Content), nil
	}
}

func (mf *MongodbFile) Write(bp *BasePath, content string) error {
	fo := FileObject{}
	err := mf.collection.Find(bson.M{"full_name": bp.Path}).One(&fo)
	if err != nil {
		return err
	}
	fo.Content = []byte(content)
	rd := bytes.NewReader(fo.Content)
	fo.Size = strconv.FormatInt(rd.Size(), 10)
	fo.Date = time.Now().String()
	err = mf.collection.Update(bson.M{"full_name": bp.Path}, &fo)
	return err
}

func (mf *MongodbFile) Delete(bp *BasePath) error {
	path := strings.Split(filepath.Clean(strings.TrimRight(bp.Path, "/")), string(os.PathSeparator))
	if len(path) < 3 {
		return errors.New("cannot remove database or table")
	}
	mf.collection = mf.database.C(bp.Table)
	fo := FileObject{}
	err := mf.collection.Find(bson.M{"full_name": bp.Path}).One(&fo)
	if err != nil {
		return err
	}
	fo.Status = 2
	fo.Date = time.Now().String()
	err = mf.collection.Update(bson.M{"full_name": bp.Path}, &fo)
	if err != nil {
		return err
	}
	if fo.FileType == global.DIR {
		fos := make([]FileObject, 0)
		//如果删除的是目录,查找出所有full_path以旧目录开始的文档
		err := mf.collection.Find(bson.M{"full_path": bson.M{"$regex": "^" + bp.Path}}).All(&fos)
		if err != nil {
			return err
		}
		if len(fos) > 0 {
			for _, v := range fos {
				v.Status = 2
				v.Date = time.Now().String()
				err = mf.collection.Update(bson.M{"full_name": v.FullName}, &v)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (mf *MongodbFile) Rename(bp *BasePath, newFullName string) error {
	path := strings.Split(strings.TrimLeft(bp.Path, "/"), "/")
	if len(path) < 3 {
		return errors.New("cannot rename database or table")
	}
	mf.collection = mf.database.C(bp.Table)
	fo := FileObject{}
	err := mf.collection.Find(bson.M{"full_name": bp.Path}).One(&fo)
	if err != nil {
		return err
	}
	fo.Name = strings.TrimPrefix(newFullName, fo.FullPath+"/")
	fo.FullName = newFullName
	fo.Date = time.Now().String()
	err = mf.collection.Update(bson.M{"full_name": bp.Path}, &fo)
	if err != nil {
		return err
	}
	if fo.FileType == global.DIR {
		fos := make([]FileObject, 0)
		//如果修改的是目录,查找出所有full_path以旧目录开始的文档
		err := mf.collection.Find(bson.M{"full_name": bson.M{"$regex": "^" + bp.Path + "/"}}).All(&fos)
		if err != nil {
			return err
		}
		if len(fos) > 0 {
			for _, v := range fos {
				ofn := v.FullName
				v.FullName = strings.Replace(v.FullName, bp.Path, newFullName, 1)
				v.FullPath = strings.Replace(v.FullPath, bp.Path, newFullName, 1)
				v.Date = time.Now().String()
				err = mf.collection.Update(bson.M{"full_name": ofn}, &v)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (mf *MongodbFile) Move(bp *BasePath, moves []string) error {
	/*
		moves			bp.path			newFullName
		/cache/a/a1						  /cache/a/abc/a1
		/cache/a/b1  => /cache/a/abc =>   /cache/a/abc/b1
		/cache/a/c1						  /cache/a/abc/c1
	*/
	//查询移动目标目录是否存在
	mf.collection = mf.database.C(bp.Table)
	if mf.collection == nil {
		return TABLE_NAME_IS_NIL
	}
	for _, p := range moves {
		paths := strings.Split(filepath.Clean(strings.TrimLeft(p, "/")), string(os.PathSeparator)) //[cache a a1]
		name := paths[len(paths)-1:]                                                               // [a1]
		newFullName := bp.Path + "/" + name[0]                                                     // /cache/a/abc/a1
		//这里使用one查询,如果没有数据,会直接返回err("not found")
		fo := FileObject{}
		iter := mf.collection.Find(bson.M{"full_name": newFullName}).Iter()
		_ = iter.Next(&fo)
		/*//如果移动后的文档在数据库中不存在或者状态为2
		if !ok || fo.Status == 2 {*/
		//移动后存在同名情况直接覆盖
		mf1, err := NewMongodbFile(mf.sess, paths[0], paths[1])
		if err != nil {
			return err
		}
		ofo := FileObject{}
		err = mf1.collection.Find(bson.M{"full_name": p}).One(&ofo)
		if err != nil {
			return err
		}

		//如果移动的文档是个目录
		if ofo.FileType == global.DIR {
			fos := make([]FileObject, 0)
			err := mf1.collection.Find(bson.M{"full_name": bson.M{"$regex": "^" + p + "/"}}).All(&fos)
			if err != nil {
				return err
			}
			if len(fos) > 0 {
				for _, v := range fos {
					/*
						/cache/a/a1/a    => /cache/a/abc  /cache/a/abc/a1/a
						/cache/a/a1/a/b  => /cache/a/abc  /cache/a/abc/a1/a/b
					*/
					prefix := paths[:len(paths)-1]             //[cache a]
					prePath := "/" + strings.Join(prefix, "/") // /cache/a
					ofn := v.FullName
					/*
						v.FullName = /cache/a/a1/a
						prePath = /cache/a
						bp.path = /cache/a/abc
					*/
					v.FullName = strings.Replace(v.FullName, prePath, bp.Path, 1) //cache/a/abc/a1/a
					v.FullPath = strings.Replace(v.FullPath, prePath, bp.Path, 1)
					v.Date = time.Now().String()
					err = mf1.collection.Update(bson.M{"full_name": ofn}, &v)
					if err != nil {
						return err
					}
				}
			}
		}

		fo.Name = ofo.Name
		fo.FullName = newFullName
		fo.FullPath = bp.Path
		fo.Rights = ofo.Rights
		fo.Size = ofo.Size
		fo.Ext = ofo.Ext
		fo.Date = time.Now().String()
		fo.FileType = ofo.FileType
		fo.Content = ofo.Content
		fo.Status = 1
		_, err = mf.collection.Upsert(bson.M{"full_name": newFullName}, &fo)
		if err != nil {
			return err
		}
		//删除移动之前数据
		err = mf1.collection.Remove(bson.M{"full_name": p})
		if err != nil {
			return err
		}

		/*} else {
			continue
		}*/
	}
	return nil
}

func (mf *MongodbFile) List(bp *BasePath) (*[]FileObject, error) {
	fo := make([]FileObject, 0)
	if bp.Table == "" {
		tables, err := mf.sess.DB(bp.Database).CollectionNames()
		if err != nil {
			return &fo, err
		}
		for _, v := range tables {
			if strings.HasPrefix(v, "system.") {
				continue
			}
			f1 := FileObject{}
			f1.Name = v
			f1.Rights = "drwxr-xr-x"
			f1.Size = "0"
			f1.Date = "2017-03-03 15:31:40"
			f1.FileType = "dir"
			fo = append(fo, f1)
		}
	} else {
		mf.collection = mf.sess.DB(bp.Database).C(bp.Table)
		mf.collection.Find(bson.M{"full_path": bp.Path, "status": 1}).All(&fo)
	}
	return &fo, nil
}

func (mf *MongodbFile) Upload(bp *BasePath, form *multipart.Form) error {
	time1 := time.Now().UnixNano()
	files := form.File
	fos := make([]FileObject, 0)
	for i := range files {
		if v, ok := files[i]; ok {
			file := v[0]
			var err error
			fos, err = parseFile(bp.Path, file)
			if err != nil {
				return err
			}
			for _, fo := range fos {
				_, err := mf.collection.Upsert(bson.M{"full_name": fo.FullName}, fo)
				if err != nil {
					return err
				}
			}
		}
	}
	fmt.Println("耗时", time.Now().UnixNano()-time1)
	return nil
}

/*func (mf *MongodbFile) Upload(bp *BasePath, form *multipart.Form) error {
	time1 := time.Now().UnixNano()
	files := form.File
	fos := make([]FileObject, 0)
	wg := new(sync.WaitGroup)
	for i := range files {
		if v, ok := files[i]; ok {
			file := v[0]
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				var err error
				fos, err = parseFile(bp.Path, file)
				if err != nil {
					return
				}
			}(wg)
		}
	}
	wg.Wait()
	for _, fo := range fos {
		_, err := mf.collection.Upsert(bson.M{"full_name": fo.FullName}, fo)
		if err != nil {
			return err
		}
	}
	return nil
}*/

func parseFile(path string, file *multipart.FileHeader) ([]FileObject, error) {
	fos := make([]FileObject, 0)
	fo := FileObject{}
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	switch file.Header.Get("Content-Type") {
	case "application/x-gzip":
		gr, err := gzip.NewReader(src)
		if err != nil {
			return nil, err
		}
		defer gr.Close()
		tr := tar.NewReader(gr)
		for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
			if err != nil {
				return nil, err
			}
			fo = parseTarReader(path, hdr)
			if hdr.FileInfo().IsDir() == false {
				fo.Content, err = ioutil.ReadAll(tr)
				if err != nil {
					return nil, err
				}
			}
			fos = append(fos, fo)
		}
	case "application/x-tar":
		tr := tar.NewReader(src)
		for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
			if err != nil {
				return nil, err
			}
			fo = parseTarReader(path, hdr)
			if hdr.FileInfo().IsDir() == false {
				fo.Content, err = ioutil.ReadAll(tr)
				if err != nil {
					return nil, err
				}
			}
			fos = append(fos, fo)
		}
	case "application/x-zip-compressed":
		zr, err := zip.NewReader(src, file.Size)
		if err != nil {
			return nil, err
		}
		if zr.File == nil {
			return fos, nil
		}
		for _, v := range zr.File {
			fo = parseZipReader(path, v)
			if v.FileInfo().IsDir() == false {
				r, err := v.Open()
				if err != nil {
					return nil, err
				}
				var b []byte
				b, err = ioutil.ReadAll(r)
				if err != nil {
					return nil, err
				}
				r.Close()
				fo.Content = b
			}
			fos = append(fos, fo)
		}
	default:
		hdrFullName := strings.Split(filepath.Clean(strings.TrimRight(file.Filename, "/")), string(os.PathSeparator))
		hdrName := hdrFullName[len(hdrFullName)-1:][0]
		hdrPath := strings.Join(hdrFullName[:len(hdrFullName)-1], "/")
		fo.Name = hdrName
		fo.FullName = path + "/" + strings.Join(hdrFullName, "/")
		fo.FullPath = strings.TrimRight(path+"/"+hdrPath, "/")
		fo.Rights = "-rwxr-xr-x"
		fo.Ext = filepath.Ext(file.Filename)
		fo.Date = time.Now().String()
		fo.FileType = global.FILE
		fo.Status = 1
		data, err := ioutil.ReadAll(src)
		if err != nil {
			return nil, err
		}
		size := len(data)
		fo.Content = data
		fo.Size = strconv.Itoa(size)
		fos = append(fos, fo)
	}
	return fos, nil
}

func parseTarReader(path string, tr *tar.Header) FileObject {
	fo := FileObject{}
	hdrFullName := strings.Split(filepath.Clean(strings.TrimRight(tr.Name, "/")), string(os.PathSeparator))
	hdrName := hdrFullName[len(hdrFullName)-1:][0]
	hdrPath := strings.Join(hdrFullName[:len(hdrFullName)-1], "/")
	fo.Name = hdrName
	fo.FullName = path + "/" + strings.Join(hdrFullName, "/")
	fo.FullPath = strings.TrimRight(path+"/"+hdrPath, "/")
	fo.Size = strconv.FormatInt(tr.Size, 10)
	fo.Status = 1
	if tr.FileInfo().IsDir() {
		fo.Rights = "drwxr-xr-x"
		fo.Ext = ""
		fo.Date = time.Now().String()
		fo.FileType = global.DIR
	} else {
		fo.Rights = "-rwxr-xr-x"
		fo.Ext = filepath.Ext(tr.Name)
		fo.Date = time.Now().String()
		fo.FileType = global.FILE
	}
	return fo
}

func parseZipReader(path string, zr *zip.File) FileObject {
	fo := FileObject{}
	hdrFullName := strings.Split(filepath.Clean(strings.TrimRight(zr.Name, "/")), string(os.PathSeparator))
	hdrName := hdrFullName[len(hdrFullName)-1:][0]
	hdrPath := strings.Join(hdrFullName[:len(hdrFullName)-1], "/")
	fo.Name = hdrName
	fo.FullName = path + "/" + strings.Join(hdrFullName, "/")
	fo.FullPath = strings.TrimRight(path+"/"+hdrPath, "/")
	fo.Size = strconv.FormatInt(zr.FileInfo().Size(), 10)
	fo.Status = 1
	if zr.FileInfo().IsDir() {
		fo.Rights = "drwxr-xr-x"
		fo.Ext = ""
		fo.Date = time.Now().String()
		fo.FileType = global.DIR
	} else {
		fo.Rights = "-rwxr-xr-x"
		fo.Ext = filepath.Ext(zr.Name)
		fo.Date = time.Now().String()
		fo.FileType = global.FILE
	}
	return fo
}

func ParsePath(path string) (*BasePath, error) {
	if path == "" {
		return nil, PATH_IS_ERROR
	}
	bp := new(BasePath)
	bp.Path = path
	path = filepath.Clean(strings.TrimLeft(path, "/"))
	path_strs := strings.Split(path, string(os.PathSeparator))
	if path_strs[0] != "conf" && path_strs[0] != "templates" && path_strs[0] != "cache" {
		return nil, errors.New("cannot create database")
	}
	if len(path_strs) >= 2 {
		bp.Database = path_strs[0]
		bp.Table = path_strs[1]
	} else {
		bp.Database = path_strs[0]
	}
	return bp, nil
}

func ParsePathToFileObject(path string) ([]FileObject, error) {
	if path == "" {
		return nil, PATH_IS_ERROR
	}
	path_str := strings.Split(filepath.Clean(strings.Trim(path, "/")), string(os.PathSeparator))
	fos := make([]FileObject, 0)
	for i, p := range path_str {
		if i == 0 || i == 1 {
			continue
		}
		full_path := "/"
		for idx := 0; idx < i; idx++ {
			full_path += path_str[idx] + "/"
		}
		fo := FileObject{}
		fo.Name = p
		fo.Size = "0"
		fo.Date = time.Now().String()
		fo.FullName = full_path + p
		fo.FullPath = strings.TrimRight(full_path, "/")
		fo.Status = 1
		fo.Ext = ""
		fo.Rights = "drwxr-xr-x"
		fo.FileType = global.DIR
		fos = append(fos, fo)
	}
	return fos, nil
}

func NewMongodbFile(sess *mgo.Session, db, table string) (*MongodbFile, error) {
	mf := new(MongodbFile)
	mf.sess = sess
	if mf.sess == nil {
		return nil, SESSION_IS_NIL
	}
	if db != "" {
		mf.database = mf.sess.DB(db)
		if mf.database == nil {
			return nil, DATABASE_IS_NIL
		}
	}
	if table != "" {
		mf.collection = mf.sess.DB(db).C(table)
		if mf.collection == nil {
			return nil, COLLATION_IS_NIL
		}
	}
	return mf, nil
}

func GetMongoSess() (*mgo.Session, error) {
	sess, err := mgo.Dial(site.TotalConfig.MongoDb.Host)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
