package demo

import (
	"fmt"
	"log"
	"reflect"
)

type Student struct {
	Name  string  `json:"name" db:"NAME"`
	Age   int     `json:"age" db:"AGE"`
	Score float32 `json:"score" db:"SCORE"`
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) Print() {
	fmt.Printf("Print 方法：%#v \n", s)
}

func getMethodByReflect(a interface{}) {
	t := reflect.TypeOf(a)
	methods := t.NumMethod()
	fmt.Printf("Student 中方法数量：methods = %d \n", methods)
	for i := 0; i < methods; i++ {
		method := t.Method(i)
		fmt.Printf("Student 第 %d 个方法中，方法名：%s，方法参数：%v \n", i, method.Name, method.Type)
	}
}

func callMethodByReflect(a interface{}) {
	v := reflect.ValueOf(a)

	var args1 []reflect.Value
	m1 := v.MethodByName("Print")
	m1.Call(args1)

	var args2 []reflect.Value
	args2 = append(args2, reflect.ValueOf("lisi"))
	m2 := v.MethodByName("SetName")
	m2.Call(args2)

	m1.Call(args1)
}

func getTypeByReflect(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("a is type, %v \n", t)
}

func getValueByReflect(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	k := t.Kind()

	switch k {
	case reflect.Int32:
		fmt.Printf("type is int32 \n")
	case reflect.String:
		fmt.Printf("type is string \n")
	case reflect.Struct:
		fmt.Printf("type is struct \n")
		num := v.NumField()
		fmt.Printf("struct field num: %d\n", num)
		for i := 0; i < num; i++ {
			vl := v.Field(i)
			fmt.Printf("第 %d 个字段，vl: %v\n", i, vl)
		}
	default:
		fmt.Printf("not know type \n")
	}
}

func setValueByReflect(a interface{}) {
	v := reflect.ValueOf(a)

	v.Elem().Field(0).SetString("李四")
	v.Elem().FieldByName("Age").SetInt(20)
	v.Elem().FieldByName("Score").SetFloat(22.333)
}

func getTagByReflect(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return
	}
	fields := t.NumField()
	for i := 0; i < fields; i++ {
		temp := t.Field(i)
		fmt.Printf("Student 第 %d 个字段中 tag 信息：json=%s，db=%s \n", i, temp.Tag.Get("json"), temp.Tag.Get("db"))
	}
}

func Test() {
	s := Student{
		Name:  "zhangsan",
		Age:   18,
		Score: 90.1,
	}
	// getTypeByReflect(s)
	// getValueByReflect(s)
	// setValueByReflect(&s)
	// fmt.Printf("s : %#v", s)
	// getMethodByReflect(&s)
	// callMethodByReflect(&s)
	getTagByReflect(&s)
}
