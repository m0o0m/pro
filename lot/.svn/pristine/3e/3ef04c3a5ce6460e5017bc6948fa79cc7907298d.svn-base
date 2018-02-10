package sqlBuild

import (
	"github.com/golyu/sql-build"
	"testing"
)

func TestUpdateTable(t *testing.T) {
	var tableName string
	for i := rune('A'); i < rune('z')+2; i++ {
		sql, err := sqlBuild.Update(tableName).String()
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(sql)
		tableName = string(i)
	}
}

func TestUpdateWhere(t *testing.T) {
	//
	sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
		Where("nameValue", "name").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set_("", "school").
		Where("nameValue", "name").
		Where(12, "age > ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Where("nameValue", "name").
		Where(12, "age > ").
		Where(0, "phone ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Where("nameValue", "name").
		Where(12, "age > ").
		Where("", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestUpdateWhere_(t *testing.T) {
	//
	sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
		Where_("nameValue", "name").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Where_("nameValue", "name").
		Where_(12, "age > ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Where_("nameValue", "name").
		Where_(12, "age > ").
		Where_(0, "phone ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Where_("nameValue", "name").
		Where_(12, "age > ").
		Where_("", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestUpdateGroupBy(t *testing.T) {
	sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
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

func TestUpdateOrderBy(t *testing.T) {
	sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
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

func TestUpdateLimit(t *testing.T) {
	sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
		Limit(10).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Limit(10).
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestUpdateIn(t *testing.T) {
	sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
		In([]string{"小明", "小华"}, "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		In([]string{"小明", "小华"}, "name").
		In([]int{12, 22, 33, 16}, "age").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		In("小花", "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestUpdateNotin(t *testing.T) {
	sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
		NotIn([]string{"小明", "小华"}, "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		NotIn([]string{"小明", "小华"}, "name").
		NotIn([]int{12, 22, 33, 16}, "age").
		Set("二中", "school").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		NotIn("小花", "name").
		Set("二中", "school").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
func TestUpdateLike(t *testing.T) {
	sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
		Like("我的家", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Like("%我的家", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Update("myTab").
		Set("二中", "school").
		Like("我的家%", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}

func TestUpdateAll(t *testing.T) {
	sql, err := sqlBuild.Update("myTab").
		Where("一班", "class").
		Where(0, "age>").
		Set("二中", "school").
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
	sql, err = sqlBuild.Update("myTab").
		Where("一班", "class").
		Where(0, "age>").
		Where("c", "").
		Set("二中", "school").
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

func TestSetEmpty(t *testing.T) {
	sql, err := sqlBuild.Update("myTable").
		Where_("123", "id").
		Set_("", "name", sqlBuild.Rule{StringValue: "*****"}).
		String()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(sql)
}
