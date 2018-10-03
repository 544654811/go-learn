package exercise1

import (
	"fmt"
)

func maopao_sort(numbers [10]int) [10]int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] < numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

func select_sort(numbers [10]int) [10]int {
	for i := 0; i < len(numbers); i++ {
		for j := 1; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				numbers[j], numbers[i] = numbers[i], numbers[j]
			}
		}
	}
	return numbers
}

func insert_sort(numbers [10]int) [10]int {
	for i := 1; i < len(numbers); i++ {
		for j := i; j > 0; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			} else {
				break
			}
		}
	}
	return numbers
}

func test1() {
	sort := [10]int{3, 0, 6, 9, 10, 8, 11, 34, 23, 56}
	sort1 := maopao_sort(sort)
	fmt.Println(sort)
	fmt.Println(sort1)
}

func test2() {
	arr := [...]int{1, 3, 5, 7, 6, 8, 10}
	for index, item := range arr {
		for i := 0; i < len(arr); i++ {
			if item+arr[i] == 8 {
				fmt.Printf("index：(%d, %d)", index, i)
			}
		}
	}
}

func Test() {
	test2()
}
