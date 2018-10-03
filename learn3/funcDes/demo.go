package funcDes

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func test2() {
	f1 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("函数类型：f1 = %T \n", f1)
	sum := f1(1, 2)
	fmt.Printf("sum = %d \n", sum)
}

func test3() {
	f1 := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Printf("f1 = %d \n", f1)
}

func test4() {
	num := 1
	defer fmt.Printf("defer语句： num = %d \n", num)
	defer func() {
		fmt.Printf("defer语句调用匿名函数： num = %d \n", num)
	}()

	num = 100
	fmt.Printf("num = %d \n", num)
}

func sub(a, b int) int {
	return a - b
}

func test5(a, b int, fc func(int, int) int) int {
	return fc(a, b)
}

func Test() {
	// f1 := add
	// fmt.Printf("函数类型：f1 = %T \n", f1)
	// sum := f1(1, 2)
	// fmt.Printf("sum = %d \n", sum)
	// test2()
	// test3()
	// test4()
	sum := test5(1, 2, add)
	sub := test5(1, 2, sub)
	fmt.Printf("sum = %d \n", sum)
	fmt.Printf("sub = %d \n", sub)
}
