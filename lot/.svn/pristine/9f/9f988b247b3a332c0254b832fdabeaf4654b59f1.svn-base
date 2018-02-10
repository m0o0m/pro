package sqlBuild

import (
	"github.com/golyu/sql-build"
	"testing"
)

func TestDeleteTable(t *testing.T) {
	var tableName string
	for i := rune('A'); i < rune('z')+2; i++ {
		sql, err := sqlBuild.Delete(tableName).String()
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(sql)
		tableName = string(i)
	}
}

func TestDeleteWhere(t *testing.T) {
	//
	sql, err := sqlBuild.Delete("myTab").
		Where("nameValue", "name").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").
		Where(0, "phone ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").
		Where("", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestDeleteWhere_(t *testing.T) {
	//
	sql, err := sqlBuild.Delete("myTab").
		Where_("nameValue", "name").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").
		Where_(0, "phone ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").
		Where_("", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestDeleteGroupBy(t *testing.T) {
	sql, err := sqlBuild.Delete("myTab").
		GroupBy("aa").
		GroupBy("bb").
		GroupBy("").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//PS:这里需要提醒一下,如果GroupBy多项的时候,出现xx,这是数据库配置的问题,改一下
}

func TestDeleteOrderBy(t *testing.T) {
	sql, err := sqlBuild.Delete("myTab").
		OrderBy("aa1").
		OrderBy("-bb1").
		OrderBy("").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestDeleteLimit(t *testing.T) {
	sql, err := sqlBuild.Delete("myTab").
		Limit(10).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Limit(10).
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestDeleteIn(t *testing.T) {
	sql, err := sqlBuild.Delete("myTab").
		In([]string{"小明", "小华"}, "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		In([]string{"小明", "小华"}, "name").
		In([]int{12, 22, 33, 16}, "age").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		In("小花", "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestDeleteNotin(t *testing.T) {
	sql, err := sqlBuild.Delete("myTab").
		NotIn([]string{"小明", "小华"}, "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		NotIn([]string{"小明", "小华"}, "name").
		NotIn([]int{12, 22, 33, 16}, "age").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		NotIn("小花", "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
func TestDeleteLike(t *testing.T) {
	sql, err := sqlBuild.Delete("myTab").
		Like("我的家", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Like("%我的家", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Delete("myTab").
		Like("我的家%", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestDeleteAll(t *testing.T) {
	sql, err := sqlBuild.Delete("myTab").
		Where("一班", "class").
		Where(0, "age>").
		Where("c", "").
		Where_("男", "sex").
		In([]string{"语文", "数学"}, "hobby").
		NotIn([]int{6, 7}, "xx").
		GroupBy("xxx").
		GroupBy("xxxx").
		OrderBy("-id").
		Limit(10).
		Offset(2).String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//injection
	sql, err = sqlBuild.Delete("myTab").
		Where("一班", "class").
		Where(0, "age>").
		Where("c", "").
		Where_("男' and 0<>(select count(*) from myTab) and ''=''", "sex").
		In([]string{"语文", "数学"}, "hobby").
		NotIn([]int{6, 7}, "xx").
		GroupBy("xxx").
		GroupBy("xxxx").
		OrderBy("-id").
		Limit(10).
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}