package demo

import (
	"fmt"
	"reflect"
)

func getTypeByReflect(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("a interface is type : %v \n", t)

	k := t.Kind()
	switch k {
	case reflect.Int32:
		fmt.Printf("type is int32 \n")
	default:
		fmt.Printf("not know type \n")
	}
}

func getValueByReflect(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type() // 与 t := reflect.TypeOf(a) 一样
	fmt.Printf("a interface is type : %v \n", t)

	k := t.Kind()
	switch k {
	case reflect.Int32:
		fmt.Printf("type is int32 \n")
	case reflect.Int16:
		fmt.Printf("type is int16 and value = %d \n", v.Int())
	default:
		fmt.Printf("not know type \n")
	}
}

func setValueByReflect(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()

	k := t.Kind()
	switch k {
	case reflect.Int32:
		v.SetInt(11111) // 设置值类型的副本，报错
	case reflect.Int16:
		v.SetInt(22222)
	case reflect.Ptr: // 说明传的是指针
		v.Elem().SetInt(3333) // v.Elem() 获取地址指向的内存
	default:
		fmt.Printf("not know type \n")
	}
}

func Test() {
	var a int16 = 23132
	// getTypeByReflect(a)
	// getValueByReflect(a)
	setValueByReflect(&a)
	fmt.Println("x = ", a)
}
