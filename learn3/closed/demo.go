package closed

import (
	"fmt"
	"strings"
)

func test1() func(int) int {
	num := 1
	return func(a int) int {
		num += a
		return num
	}
}

func demo1() {
	f := test1()
	sum := f(2)
	fmt.Printf("f(2)：sum=%d\n", sum)
	sum = f(20)
	fmt.Printf("f(20)：sum=%d\n", sum)
	sum = f(200)
	fmt.Printf("f(200)：sum=%d\n", sum)

	f1 := test1()
	sum1 := f1(1)
	fmt.Printf("f1(1)：sum=%d\n", sum1)
	sum1 = f1(10)
	fmt.Printf("f1(10)：sum=%d\n", sum1)
	sum1 = f1(100)
	fmt.Printf("f1(100)：sum=%d\n", sum1)
}

func test2(str string) func(string) string {
	return func(rule string) string {
		if !strings.HasSuffix(str, rule) {
			return str + rule
		}
		return str
	}
}

func test3(rule string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, rule) {
			return name + rule
		}
		return name
	}
}

func demo3() {
	func1 := test3(".png")
	fmt.Printf("func1：str=%s\n", func1("text"))
	fmt.Printf("func1：str=%s\n", func1("text"))
	func2 := test3(".jpg")
	fmt.Printf("func2：str=%s\n", func2("text"))
}

func demo2() {
	f := test2("text")
	fmt.Printf("f：str=%s\n", f(".png"))
	// f = test2("text2")
	fmt.Printf("f：str=%s\n", f(".pngsdf"))

	f1 := test2("text")
	fmt.Printf("f1：str=%s\n", f1(".jpg"))
}

func Test() {
	// demo1()
	// demo2()
	demo3()
}
