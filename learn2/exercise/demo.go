package exercise

import (
	"fmt"
)

func TestCfb() {
	for a := 1; a < 10; a++ { // 行
		for b := 1; b <= a; b++ { // 列
			fmt.Printf("%d*%d = %d  ", b, a, b*a)
		}
		fmt.Println()
	}
}
