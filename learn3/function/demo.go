package function

import (
	"fmt"
)

func testFunc1(a, b int) int {
	return a + b
}

func testFunc2(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func testFunc3(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func testFunc4(a ...int) (sum int) {
	sum = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return
}

func testFunc5(b int, a ...int) (sum int) {
	sum = b
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return
}

func testDefer() {
	i := 0
	for i = 0; i < 5; i++ {
		defer fmt.Printf("defer函数：i = %d \n", i)
	}
	i = 1000
	fmt.Printf("函数运行中：i = %d \n", i)
	fmt.Printf("函数结束... \n")
}

func TestFunc() {
	// fmt.Println(testFunc1(1, 2))

	// sum, sub := testFunc2(10, 20)
	// fmt.Printf("sum = %d, sub = %d \n", sum, sub)

	// sum, sub := testFunc3(10, 20)
	// fmt.Printf("sum = %d, sub = %d \n", sum, sub)

	// _, sub := testFunc3(10, 20)
	// fmt.Printf("sub = %d \n", sub)

	// sum := testFunc4()
	// fmt.Printf("当没有参数时：sum = %d \n", sum)
	// sum1 := testFunc4(10)
	// fmt.Printf("当有1个参数时：sum1 = %d \n", sum1)
	// sum2 := testFunc4(10, 20)
	// fmt.Printf("当有2个参数时：sum2 = %d \n", sum2)
	// sum3 := testFunc4(10, 20, 30, 40, 50)
	// fmt.Printf("当有多个参数时：sum3 = %d \n", sum3)

	// sum := testFunc5(10)
	// fmt.Printf("当没有参数时：sum = %d \n", sum)
	// sum1 := testFunc5(10, 10)
	// fmt.Printf("当有1个参数时：sum1 = %d \n", sum1)
	// sum2 := testFunc5(10, 10, 20)
	// fmt.Printf("当有2个参数时：sum2 = %d \n", sum2)
	// sum3 := testFunc5(10, 10, 20, 30, 40, 50)
	// fmt.Printf("当有多个参数时：sum3 = %d \n", sum3)

	testDefer()
}
