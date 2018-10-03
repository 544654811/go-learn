package demo

import (
	"fmt"
)

func test1() {
	var a int
	fmt.Printf("a的内存地址: %p \n", &a)
}

func test2() {
	var a *int
	fmt.Printf("指针变量 a 的内存地址：%p, a 的值：%v \n", &a, a)

	b := 100
	fmt.Printf("变量 b 的内存地址：%p, b 的值：%v \n", &b, b)
	a = &b
	fmt.Printf("指针变量 a 的内存地址：%p, a 的值：%v \n", &a, a)
}

func test3() {
	b := 100
	fmt.Printf("变量 b 的内存地址：%p, b 的值：%v \n", &b, b)

	a := &b
	fmt.Printf("指针变量 a 的内存地址：%p, a 的值：%v \n", &a, a)

	*a = 200
	fmt.Printf("修改指针变量 a 后, 内存地址：%p, a 的值：%v, *a：%d \n", &a, a, *a)
	fmt.Printf("修改指针变量 a 后, 变量 b 的内存地址：%p, b 的值：%v \n", &b, b)
}

func test4() {
	a := 100
	fmt.Println("变量a的地址：", &a)
}

func pointArg(a *int) {
	*a = 100
}

func testPointArg() {
	a := 1
	fmt.Println("变量a：", a)
	pointArg(&a)
	fmt.Println("传参后变量a：", a)
}

func Test() {
	// test1()
	// test2()
	// test3()
	// test4()
	testPointArg()
}
