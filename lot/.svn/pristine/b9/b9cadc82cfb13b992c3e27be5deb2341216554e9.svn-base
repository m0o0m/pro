package pay

import (
	"reflect"
	"fmt"
)

//给struct中的属性排序并用&符连接,排序的规则是:参数名(由json的tag决定)a到z的顺序排序，若遇到相同首字母，则看第二个字母，以此类推
// 组成规则如下：
// 参数名1=参数值1&参数名2=参数值2&……&参数名n=参数值n&key=key值
type tree struct {
	name  string
	value string
	left,
	right *tree //包含自己的的指针类型
}
type result struct {
	data   string //最终的字符串
	symbol string //连接符,默认为&
}

//排序方式是按照字母a-z序
//必须接收指针,第二的参数决定用什么tag来排序,第三个参数是决定连接符,默认是&
func Sort(st interface{}, args ... string) (string) {
	var tagKey = "json"
	var result = result{data: "", symbol: "&"}
	if len(args) > 0 {
		tagKey = args[0]
	}
	if len(args) > 1 {
		result.symbol = args[1]
	}

	val := reflect.ValueOf(st)
	ind := reflect.Indirect(val)
	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("The struct paramer must be use ptr"))
	}
	ty := ind.Type()
	var root *tree
	for i := 0; i < ty.NumField(); i++ {
		name := ty.Field(i).Tag.Get(tagKey)
		value := getValue(ind.Field(i))
		if name != "" && value != "" {
			root = add(root, name, value)
		}
	}
	return join(&result, root).data
}

//排序是按照field的先后顺序
//必须接收指针,第二的参数决定用什么tag来排序,第三个参数是决定连接符,默认是&
func NoSort(st interface{}, args ... string) (string) {
	var tagKey = "json"
	var symbol = "&"
	if len(args) > 0 {
		tagKey = args[0]
	}
	if len(args) > 1 {
		symbol = args[1]
	}

	val := reflect.ValueOf(st)
	ind := reflect.Indirect(val)
	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("The struct paramer must be use ptr"))
	}
	ty := ind.Type()
	var result string
	for i := 0; i < ty.NumField(); i++ {
		name := ty.Field(i).Tag.Get(tagKey)
		value := getValue(ind.Field(i))
		if name != "" && value != "" {
			//root = add(root, name, value)
			if result != "" {
				result += symbol
			}
			result += name + "=" + value
		}
	}
	return result
}

//得到数据类型
func getValue(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Int:
		return fmt.Sprintf("%d", value.Int())
	case reflect.String:
		return value.String()
	}
	return ""
}

func join(result *result, t *tree) *result {
	if t != nil {
		result = join(result, t.left)
		if result.data == "" {
			result.data += t.name + "=" + t.value
		} else {
			result.data += result.symbol + t.name + "=" + t.value
		}
		result = join(result, t.right)
	}
	return result
}

func add(t *tree, name, value string) *tree {
	if t == nil {
		//等价于返回&tree{value:value
		t = new(tree)
		t.name = name
		t.value = value
		return t
	}
	if compare(name, t.name) {
		t.left = add(t.left, name, value)
	} else {
		t.right = add(t.right, name, value)
	}
	return t
}

//比较字符串,顺序按照参数名a到z的顺序排序，若遇到相同首字母，则看第二个字母，以此类推,
// true为va1排前面,false为val2排前面,如果两个字符串一致,也返回true
func compare(va1, va2 string) (bool) {
	byteVa1 := []byte(va1)
	byteVa2 := []byte(va2)
	var size int
	//取短的进行循环
	if len(byteVa1) > len(byteVa2) {
		size = len(byteVa2)
	} else {
		size = len(byteVa1)
	}
	for i := 0; i < size; i++ {
		if byteVa1[i] < byteVa2[i] {
			return true
		} else if byteVa1[i] > byteVa2[i] {
			return false
		} else {
			continue
		}
	}
	//比较完之后,还没有分出胜负,谁短谁排前面
	if len(byteVa1) == size {
		return true
	}
	return false
}
