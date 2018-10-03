package main

import (
	"fmt"
)

const a = 456

func main() {
	// const a int = 123
	const a1, a2, a3, a4 = 1, 0.1, true, "张三"

	const (
		b = 2
		c
		d
		e
	)

	const (
		A = iota
		B
		C
		D = 8
		E
		F = iota
		G
	)

	fmt.Println(a, b, c, d, e)
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(A, B, C, D, E, F, G)
}
