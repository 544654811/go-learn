package forDemo

import (
	"fmt"
)

func TestFor1() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i = %d \n", i)
	}
	// fmt.Printf("循环结束后：i = %d \n", i)
}

func TestFor2() {
	i := 0
	for i <= 10 {
		fmt.Printf("i = %d \n", i)
		i++
	}
	fmt.Printf("循环结束后：i = %d \n", i)
}

func TestFor3() {
	i := 0
	for i <= 10 {
		fmt.Printf("i = %d \n", i)
		i += 2
	}
	fmt.Printf("循环结束后：i = %d \n", i)
}

func TestFor4() {
	for a, b := 1, 10; a <= 10 && b <= 20; a, b = a+1, b+1 {
		fmt.Printf("%d * %d = %d \n", a, b, a*b)
	}
}

func TestFor5() {
	var i int
	for i = 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("i = %d \n", i)
	}
	fmt.Printf("循环结束后：i = %d \n", i)
}

func TestFor6() {
	var i int
	for i = 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("i = %d \n", i)
	}
	fmt.Printf("循环结束后：i = %d \n", i)
}
