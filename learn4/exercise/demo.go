package exercise

import (
	"fmt"
	"sort"
)

func exercise1() {
	var sa = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		sa = append(sa, i)
	}
	fmt.Println(sa)
}

func exercise2(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func testExercise2() {
	a := []int{3, 4, 1, 7, 4, 9, 11, 5, 10}
	fmt.Println("排序前 a：", a)
	a = exercise2(a)
	fmt.Println("排序后 a：", a)
}

func Test() {
	exercise1()
	testExercise2()
}
