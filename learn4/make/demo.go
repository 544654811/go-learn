package make

import (
	"fmt"
)

func test1() {
	a := make([]int, 5, 10)
	fmt.Printf("a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))

	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	}
}

func Test() {
	test1()
}
