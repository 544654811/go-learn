package exercise

import (
	"fmt"
)

func judgePrime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func testPrime() {
	for i := 1; i < 100; i++ {
		if judgePrime(i) {
			fmt.Printf("[%d] 为质数... \n", i)
		}
	}
}

func is_sxhs(num int) bool {
	a := num % 10         // 个位
	b := (num / 10) % 10  // 十位
	c := (num / 100) % 10 // 百位
	sum := a*a*a + b*b*b + c*c*c
	if num == sum {
		return true
	}
	return false
}

func TestSxhs() {
	for i := 100; i < 1000; i++ {
		if is_sxhs(i) {
			fmt.Printf("[%d] 为水仙花数... \n", i)
		}
	}
}

func count2(str string) (letters int, numbers int, spaces int, others int) {
	for _, item := range str {
		switch {
		case item >= 'a' && item <= 'z' || item >= 'A' && item <= 'Z':
			letters++
		case item >= '0' && item <= '9':
			numbers++
		case item == ' ':
			spaces++
		default:
			others++
		}
	}
	return
}

func count(str string) (letters int, numbers int, spaces int, others int) {
	for _, item := range str {
		if item >= 'a' && item <= 'z' || item >= 'A' && item <= 'Z' {
			letters++
			continue
		}

		if item >= '0' && item <= '9' {
			numbers++
			continue
		}

		if item == ' ' {
			spaces++
			continue
		}

		others++
	}
	return
}

func TestCount() {
	// letters, numbers, spaces, others := count("afa23fs  1a31  /./.张s三")
	// fmt.Printf("字母：%d, 数字：%d, 空格：%d, 其他：%d \n", letters, numbers, spaces, others)
	letters, numbers, spaces, others := count2("afa23fs  1a31  /./.张s三")
	fmt.Printf("字母：%d, 数字：%d, 空格：%d, 其他：%d \n", letters, numbers, spaces, others)
}

func TestExercise() {
	// testPrime()
	// TestSxhs()
	TestCount()
}
