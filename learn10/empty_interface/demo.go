package demo

import (
	"fmt"
)

func emptyDemo(a interface{}) {
	fmt.Printf("type: %T, value: %v \n", a, a)
}

func emptyTest() {
	a := 100
	emptyDemo(a)

	b := "测试啊..."
	emptyDemo(b)

	c := map[string]int{
		"zhangsan": 100,
		"lisi":     200,
	}
	emptyDemo(c)
}

func typeDemo(a interface{}) {
	aa, ok := a.(int)
	if ok {
		fmt.Printf("判断类型是 int, value: %v \n", aa)
		return
	}

	bb, ok := a.(string)
	if ok {
		fmt.Printf("判断类型是 string, value: %v \n", bb)
		return
	}

	fmt.Printf("无法判断类型, value: %v \n", a)
}

func typeTest() {
	a := 100
	typeDemo(a)

	b := "测试啊..."
	typeDemo(b)

	c := map[string]int{
		"zhangsan": 100,
		"lisi":     200,
	}
	typeDemo(c)
}

func typeDemo2(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Printf("判断类型是 int, value: %v \n", v)
	case string:
		fmt.Printf("判断类型是 string, value: %v \n", v)
	case map[string]int:
		fmt.Printf("判断类型是 map, value: %v \n", v)
	default:
		fmt.Printf("无法判断类型, value: %v \n", v)
	}
}

func typeTest2() {
	a := 100
	typeDemo2(a)

	b := "测试啊..."
	typeDemo2(b)

	c := map[string]int{
		"zhangsan": 100,
		"lisi":     200,
	}
	typeDemo2(c)
}

func Test() {
	// emptyTest()
	// typeTest()
	typeTest2()
}
