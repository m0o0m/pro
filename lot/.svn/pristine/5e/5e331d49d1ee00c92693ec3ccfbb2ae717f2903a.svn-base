package sqlBuild

import (
	"testing"
	"github.com/golyu/sql-build"
)

func TestTable(t *testing.T) {
	var tableName string
	for i := rune('A'); i < rune('z')+2; i++ {
		sql, err := sqlBuild.Select(tableName).String()
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(sql)
		tableName = string(i)
	}
}

func TestColumn(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		Column("aaa","zzz").
		Column("bbb as xx").
		Column("").
		String()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(sql)
}

func TestWhere(t *testing.T) {
	//
	sql, err := sqlBuild.Select("myTab").
		Where("nameValue", "name").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").
		Where(0, "phone ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").
		Where("", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestWhere_(t *testing.T) {
	//
	sql, err := sqlBuild.Select("myTab").
		Where_("nameValue", "name").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").
		Where_(0, "phone ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").
		Where_("", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestGroupBy(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
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

func TestOrderBy(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
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

func TestLimit(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		Limit(10).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Limit(10).
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestIn(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		In([]string{"小明", "小华"}, "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		In([]string{"小明", "小华"}, "name").
		In([]int{12, 22, 33, 16}, "age").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		In("小花", "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestNotin(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		NotIn([]string{"小明", "小华"}, "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		NotIn([]string{"小明", "小华"}, "name").
		NotIn([]int{12, 22, 33, 16}, "age").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		NotIn("小花", "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
func TestLike(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		Like("我的家", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Like("%我的家", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Like("我的家%", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestAll(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
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
	sql, err = sqlBuild.Select("myTab").
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

//test bug tag:#2
func TestBug2(t *testing.T)  {
	sql, err := sqlBuild.Select("myTable").
		Where(float64(1503909330000), "time > ").
		String()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(sql)
}
//test bug tag:#3
func TestBug3(t *testing.T) {
	sql, err := sqlBuild.Select("permission ").
		Column("service", "operation").
		WhereFunc("toTimeUUID('2017-10-27 01:00+0000')", "tid > ").
		String()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(sql)
}