### sql-build
sql-build是一个支持条件控制的go语言sql拼接库.共分为4个部分,select,insert,update和delete四个部分,生成的结果为完整sql语句,需要和beego,xorm以及其它支持原生sql语句的orm配置使用,sql-build只做拼接工作.

### download
```go
go get github.com/golyu/sql-build
```
### upgrade
```go
go get -u github.com/golyu/sql-build
```
    
主要功能有
- 无序的链式拼接
- 支持string类型的值注入检测
- 通过条件控制,决定数据有效性 (可以自定条件,用户可以自定义自己的规则,来组合`sql-build`,这个可以参考[sql-build-example](https://github.com/golyu/sql-build-example)项目中`models/build`目录中的写法,进行组合设置)
- 不支持sql预编译(这个既是优点也是缺点,优点是,可以在不支持预编译语句的sql中间件上使用,缺点就是预编译的优点)
    
需要注意的是:
- 目前只支持单表操作,不支持联表
- 参数的值在前,列名在后(这个用起来有点别扭,主要是为了使用idea的template)
- 可能缺少一些sql关键字,如果有需要,可以pull request
- 不支持sql预编译
### 解决的痛点
如果的models中,我们组合sql的时候,还需要判断一些参数是否存在,我们使用拼接库是这样的:
```go
func UpdateUser(id int, username, sex, name string, state int, money float64, qq, email string) error {
	if id > 0 {
		set := ""
		if len(username) > 0 {
			set += " name = '" + username + "'"
		}
		if len(sex) > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " sex = '" + sex + "'"
		}
		if len(name) > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " name = '" + name + "'"
		}
		if state > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " state = " + strconv.Itoa(state)
		}
		if money > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " money = " + fmt.Sprintf("%g",money)
		}
		if len(qq) > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " qq = '" + qq + "'"
		}
		if len(email) > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " email = '" + email + "'"
		}
		if len(set) > 0 {
			set = " SET" + set
		} else {
			return errors.New("not need update column")
		}
		sql := fmt.Sprintf("UPDATE user "+set+" WHERE id = %d",
			id)
		o:=orm.NewOrm()
		_, err := o.Raw(sql).Exec()
		return err
	}
	return errors.New("Id can not <=0")
}
```
每一步一个判断,还要考虑数据的类型,一不小心就会出错,下面使用`sql-build`来重构上述代码
```go
func UpdateUser(id int, username, sex, name string, state int, money float64, qq, email string) error {
	sql, err := sqlBuild.Update("user").
		Where_(id, "id").
		Set(username, "username").
		Set(sex, "sex").
		Set(name, "name").
		Set(state, "state").
		Set(money, "money").
		Set(qq, "qq").
		Set(email, "email").
		String()
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.Raw(sql).Exec()
	return err
}
```
一下子,就简洁了很多,还不用类型转换,不担心写错,如果配合idea的template功能一起使用,你会感觉更简单,晒一张图
![image](https://github.com/golyu/sql-build/blob/master/img/temp.gif)

### 使用说明
````go
import (
	"github.com/golyu/sql-build/debug"
)
sqlBuild.Debug() //开启debug模式,可以看到错误和警告的打印
````

### *select*
select功能可以支持以下函数,除了Select和String函数放在语句的头尾处,其余的都可以无序设置
- Select(table string) SelectInf
- Column(column string) SelectInf
- Where(value interface{}, key string, rules ... Rule) SelectInf
- Where_(value interface{}, key string, rules ... Rule) SelectInf
- Like(value string, key string) SelectInf
- In(values interface{}, key string) SelectInf
- NotIn(values interface{}, key string) SelectInf
- OrderBy(orderBy string) SelectInf
- Limit(limit int) SelectInf
- Offset(offset int) SelectInf
- GroupBy(groupBy string) SelectInf
- String() (string, error)

##### String 生成sql的函数
```go
import (
	"testing"
	"github.com/golyu/sql-build"
)

func TestTable(t *testing.T) {
    myTab:="myTab"
	sql, err := sqlBuild.Select(myTab).String()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(sql)
	tableName = string(i)
}
```
输出sql会是:
```go
SELECT * FROM myTab
```
如果把myTab置为空,err则打印
```go
The tabName can not be empty
```
以下示例为了简洁,不显示import

##### Column 查询结果显示列

```go
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
```
sql打印
```go
SELECT aaa,bbb as xx FROM myTab
```
空的数据无法通过数据有效性检测,不予组装(这个数据有效性条件是可以控制的)

##### Where条件控制语句
```go
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
```
sql输入结果为:
```go
	SELECT * FROM myTab WHERE name = 'nameValue'
	SELECT * FROM myTab WHERE name = 'nameValue' and age > 12
	SELECT * FROM myTab WHERE name = 'nameValue' and age > 12
	SELECT * FROM myTab WHERE name = 'nameValue' and age > 12
```
对于key中没有带>=<符号的,sql-build会自定补上=.
第三,四,明明有些条件存在,为什么生成的sql后丢失了呢,这就是sql-build的条件控制功能了,默认的数值类型<=0的条件,string类型值为""的都会被过滤掉,也就是数据无效,如果你说,我们的业务里也有值为0的啊,怎么办,sql-build可以自定义数据有效性过滤条件,在后面会详细说明.

##### Where_ 条件控制语句
```go
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
```
sql输出结果为:
```go
SELECT * FROM myTab WHERE name = 'nameValue'
SELECT * FROM myTab WHERE name = 'nameValue' and age > 12
```
err输出结果为:
```go
Fail to meet the condition
Fail to meet the condition
```
可以看出,`Where_`功能上和`Where`基本一致,只是把过滤替换成了抛出异常,这是对于有些业务需要指定的条件,而该条件可能为空的情况,做出的判断

##### WhereFunc
WhereFunc条件能自动过滤WhereFunc("toTimeUUID('2017-10-27 01:00+0000')", "tid > ").这种函数的调用,在结果两边不会自动加上''
 ```go
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
```
sql输出
```go
SELECT service,operation FROM permission  WHERE tid > toTimeUUID('2017-10-27 01:00+0000')
```

##### GroupBy
```go
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
}
```
sql输出
```go
SELECT * FROM myTab GROUP BY aa,bb
```
##### OrderBy
```go
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
```
sql输出为:
```go
SELECT * FROM myTab ORDER BY aa1,-bb1
```
正序直接写入列名,倒序在前加上-号

##### Limit Offset
```go
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
```
sql输出为:
```go
SELECT * FROM myTab LIMIT 10
SELECT * FROM myTab LIMIT 10 OFFSET 2
```
err输出为:
```go
Need 'Offset' and 'Limit' are used together
```
Limit可以单独使用,Offset需要配合Offset使用,可以先Offset再Limit,我们中间的功能语句无序,只需要出现过就行
##### In
```go
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
```
sql输出:
```go
SELECT * FROM myTab WHERE name IN ('小明','小华')
SELECT * FROM myTab WHERE name IN ('小明','小华') and age IN (12,22,33,16)
```
err输出:
```go
The Value not have need type
```
in语句的值不允许为基本类型切片外的其它任何类型,包括基础类型,上述语句可以改写为
```go
sql, err = sqlBuild.Select("myTab").
	In([]string{"小花"}, "name").
	String()
```
##### NotIn
```go
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
```
##### Like
```go
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
```
sql输出为:
```go
SELECT * FROM myTab WHERE address like '%我的家%'
SELECT * FROM myTab WHERE address like '%我的家'
SELECT * FROM myTab WHERE address like '我的家%'
```
如果like的值没有%,自动在前后补上%

以上就是sql-build中select功能的大致使用,功能块的函数是可以混合使用的,上面的举例只是部分写法,那我们合起来写一次
```go
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
```
sql输出:
```go
SELECT * FROM myTab WHERE class = '一班' and sex = '男' and hobby IN ('语文','数学') and xx NOT IN (6,7) GROUP BY xxx,xxxx ORDER BY -id LIMIT 10 OFFSET 2
```
err输出:
```go
Injection err
```
检查出string类型的值里面包含注入的关键字

### *insert*

insert功能支持以下函数

- Insert(table string) InsertInf
- Option(options ...string) InsertInf
- NoOption(noOptions...string)InsertInf
- Value(value interface{}, rules ... Rule) InsertInf
- Values(value interface{}, rules ... Rule) InsertInf
- OrUpdate()InsertInf
- String() (string, error)

> 注意1:insert方法因为需要在insert的方法上面加上insert的tag,value和values方法不可以同时使用
> 注意2:如果数据库中有自增主键,请在tags中用auto标出,使用`;`隔开,一个struct中只可以标记一个为auto为自增主键,如果使用了mycat中间件来设置自动增长,需要在标注了auto的情况下再标注mycat标签auto;mycat:next value for MYCATSEQ_AGENT,如下面例子
> 注意3:如果使用了option或者noOption方法,请按调用顺序放在value和values前面
##### value

```go
type Tab struct {
	Id   int `insert:"id;auto"`
	Name string `insert:"name"`
	Age  int `insert:"age"`
}

func TestValue(t *testing.T) {
	var tab = Tab{Id: 0, Name: "yiersan", Age: 18}
	sql, err := sqlBuild.Insert("xx").
		Value(&tab).String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
```
首先,给`filed`加上`insert`的tag,然后调用value,传入指针,同`select`一样,value方法也可以自定义规则传入使用.
sql打印:
```go
INSERT INTO xx(id,name,age) VALUES (DEFAULT,'yiersan',18)

```
```go
type Tab struct {
	Id   int    `insert:"id;auto;mycat:next value for MYCATSEQ_AGENT"`
	Name string `insert:"name"`
	Age  int    `insert:"age"`
}
func TestMycat(t *testing.T) {
	debug.Debug = true
	var tab = Tab{Id: 0, Name: "yiersan", Age: 18}
	sql, err := sqlBuild.Insert("xx").
		Value(&tab).
		String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
```
sql打印:
```go
INSERT INTO xx(id,name,age) VALUES (next value for MYCATSEQ_AGENT,'yiersan',18)
```

##### values 批量插入

```go
func TestValues(t *testing.T) {
	var tab1 = Tab{Id: 0, Name: "yiersan", Age: 18}
	var tab2 = Tab{Id: 0, Name: "xx", Age: 16}
	var tab3 = Tab{Id: 0, Name: "pp", Age: 18}
	var tabs = []Tab{tab1, tab2, tab3}
	sql, err := sqlBuild.Insert("xx").
		Values(tabs).String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
```
sql打印:
```go
INSERT INTO xx(id,name,age) VALUES (DEFAULT,'pp',18),(DEFAULT,'xx',16),(DEFAULT,'yiersan',18)
```
这里的批量插入在源码中的反射异步了,所以插入多条和插入单条,理论时间大致是一样的

##### OrUpdate 不存在就插入,存在就更新
```go
func TestOrUpdate(t *testing.T) {
	var tab1 = Tab{Id: 0, Name: "yiersan", Age: 18}
	var tab2 = Tab{Id: 0, Name: "xx", Age: 16}
	var tab3 = Tab{Id: 0, Name: "pp", Age: 18}
	var tabs = []Tab{tab1, tab2, tab3}
	sql, err := sqlBuild.Insert("xx").
		Values(tabs).
		OrUpdate().
		String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
```
sql打印:
```go
INSERT INTO xx(id,name,age) VALUES (DEFAULT,'xx',16),(DEFAULT,'yiersan',18),(DEFAULT,'pp',18)  ON DUPLICATE KEY UPDATE id = values(id),name = values(name),age = values(age)
```
##### Option 只insert指定选项
```go
func TestOption(t *testing.T) {
	debug.Debug = true
	var tab1 = Tab{Id: 0, Name: "yiersan", Age: 18}
	sql, err := sqlBuild.Insert("xx").
		Option("name").
		Value(&tab1).String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
```
sql打印
```go
INSERT INTO xx(name) VALUES ('yiersan') 
```
##### NoOption 不insert指定选项
```go
func TestNoOption(t *testing.T) {
	debug.Debug = true
	var tab1 = Tab{Id: 0, Name: "yiersan", Age: 18}
	var tab2 = Tab{Id: 0, Name: "xx", Age: 16}
	var tab3 = Tab{Id: 0, Name: "pp", Age: 18}
	var tabs = []Tab{tab1, tab2, tab3}
	sql, err := sqlBuild.Insert("xx").
		NoOption("name").
		Values(tabs).OrUpdate().String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
```
sql打印
```go
INSERT INTO xx(id,age) VALUES (NULL,18),(NULL,18),(NULL,16)  ON DUPLICATE KEY UPDATE id = values(id),age = values(age)
```

#### Delete
delete功能支持以下函数

- Delete(table string) DeleteInf
- Where(value interface{}, key string, rules ... Rule) DeleteInf
- Where_(value interface{}, key string, rules ... Rule) DeleteInf
- Like(value string, key string) DeleteInf
- In(values interface{}, key string) DeleteInf
- NotIn(values interface{}, key string) DeleteInf
- OrderBy(orderBy string) DeleteInf
- Limit(limit int) DeleteInf
- Offset(offset int) DeleteInf
- GroupBy(groupBy string) DeleteInf
- String() (string, error)

`delete`功能基本支持`select`功能的所有函数和用法,`column`除外
我们还是写一个例子
```go
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
}	
```
sql打印:
```go
DELETE FROM myTab WHERE class = '一班' and sex = '男' and hobby IN ('语文','数学') and xx NOT IN (6,7) GROUP BY xxx,xxxx ORDER BY -id LIMIT 10 OFFSET 2
```
可能delete主要用到的也只有`Where`,`Where_`,`In`,`NotIn`功能的,其它的有些鸡肋,不过既然sql支持的,我们也就尽量支持,至少提供一个可选项

#### Update

`update`功能也是在`select`功能基础上改造,去掉了`column`,添加了`set`和`set_`函数

- Update(table string) UpdateInf
- Set(value interface{}, key string, rules ... Rule) UpdateInf
- Set_(value interface{}, key string, rules ... Rule) UpdateInf
- Where(value interface{}, key string, rules ... Rule) UpdateInf
- Where_(value interface{}, key string, rules ... Rule) UpdateInf
- Like(value string, key string) UpdateInf
- In(values interface{}, key string) UpdateInf
- NotIn(values interface{}, key string) UpdateInf
- OrderBy(orderBy string) UpdateInf
- Limit(limit int) UpdateInf
- Offset(offset int) UpdateInf
- GroupBy(groupBy string) UpdateInf
- String() (string, error)

同样,我们也写个示例
```go
func TestUpdateWhere(t *testing.T) {
    sql, err := sqlBuild.Update("myTab").
		Set("二中", "school").
		Where("nameValue", "name").String()
	if err != nil {
    		t.Error(err.Error())
    		err = nil
    	}
    	t.Log(sql)
    sql, err = sqlBuild.Update("myTab").
    		Set_("", "school").
    		Where("nameValue", "name").
    		Where(12, "age > ").String()
    	if err != nil {
    		t.Error(err.Error())
    		err = nil
    	}
    	t.Log(sql)
}
```
sql打印:
```go
UPDATE myTab SET school = '二中' WHERE name = 'nameValue'
```
错误打印:
```go
Fail to meet the set
```
有个朋友问了个问题,可能是我的说明没写清楚,问题是这样的,确实需要把一个值`update`成空字符串,怎么办呢,
如果只是少量语句需要这样,可以使用以下的方法,局部条件过滤,只要设置的值不是`*****`,就不会报`Fail to meet the set`错误
```go
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

```
sql打印
```go
UPDATE myTable SET name = '' WHERE id = '123'
```

`set_`,`where_`,这种后面带下划线的都表强制,如果不符合过滤条件,就会返回错误,而不带下划线的`set`,`where`都是跳过不合条件的数据

> 以上的详细示例都可以在`test`文件夹下的测试文件中找到

##### 常见错误对照表
|错误|错误原因|相关函数|
|---|---|---|
|`The tabName can not be empty  meiy`|没有给表名|Select(),Update(),Insert(),Delete()|
|`The Value not have need type`|传的值不是预期的类型|Value(),Values(),Set(),Set_(),Where_(),Where()|
|`Injection err`|可能会有注入危险|all|
|`Not Found Update Data`|没有需要更新的数据|Set(),Set_()|
|`Fail to meet the condition`|不符合过滤条件|Where_()|
|`Fail to meet the set`|不符合过滤条件|Set_()|
|`Need 'Offset' and 'Limit' are used together`|offset方法必须配合limit方法使用|Offset()|
|`Not found Insert Column`|没有找到insert的列名|tag insert:"xxx"|
|`Not found Insert Data`|没有找到要insert的数据|没有要insert的数据|





