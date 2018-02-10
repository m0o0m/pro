package global

import (
	"github.com/go-xorm/xorm"
	"io"
	"mime/multipart"
)

//列表参数
type ListParams struct {
	GetCount bool   `query:"getCount"` //是否统计总数
	Page     int    `query:"page"`     //页码数
	PageSize int    `query:"pageSize"` //页面数量
	Limit    []int  //分页
	OrderBy  string `query:"orderBy"` //排序字段
	Desc     bool   `query:"desc"`    //排序顺序 true 正 false 反
}

func (this *ListParams) Make(model *xorm.Session) {
	if this.Limit != nil {
		model.Limit(this.Limit[0], this.Limit[1])
	}
	if this.OrderBy != "" {
		if this.Desc == true {
			model.Desc(this.OrderBy)
		} else {
			model.Asc(this.OrderBy)
		}
	}
}

type Times struct {
	StartTime int64 `query:"start_time"` //开始时间
	EndTime   int64 `query:"end_time"`   //结束时间
}

//根据时间查询
func (this *Times) Make(timeParam string, model *xorm.Session) {
	if this.StartTime > 0 && this.EndTime > 0 {
		model.Where(timeParam+">=?", this.StartTime).And(timeParam+"<=?", this.EndTime)
	} else if this.StartTime > 0 && this.EndTime == 0 {
		model.Where(timeParam+">=?", this.StartTime)
	} else if this.StartTime == 0 && this.EndTime > 0 {
		model.Where(timeParam+"<=?", this.EndTime)
	}
}

//文件二进制流读取
func ReadByte(src *multipart.FileHeader) ([]byte, error) {
	fi, err := src.Open()
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)

	}
	return chunks, err
}
