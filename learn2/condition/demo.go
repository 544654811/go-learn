package condition

import (
	"fmt"
)

func TestIf() {
	num := 10
	if num%2 == 0 {
		fmt.Println("num 为偶数...")
	}
}

func TestIfelse() {
	num := 11
	if num%2 == 0 {
		fmt.Println("num 为偶数...")
	} else {
		fmt.Println("num 为奇数...")
	}
}

func Test() {
	if num := getNum(); num%2 == 0 {
		fmt.Println("num 为偶数...")
	} else {
		fmt.Println("num 为奇数...")
	}
}

func getNum() int{
	return 11
}
