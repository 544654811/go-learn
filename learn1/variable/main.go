package main

import "fmt"

func main() {
	// var a int
	// var b float32
	// var c bool
	// var d string

	// var (
	// 	a int
	// 	b float32
	// 	c bool
	// 	d string
	// )

	a, b, c, d := 1, 0.123, true, "李四"

	a = 1
	d = "张三"
	fmt.Printf("a=%d b=%f c=%t d=%s", a, b, c, d)
}
