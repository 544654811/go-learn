package switchDemo

import (
	"fmt"
)

func TestIf() {
	var num int = 3
	if num == 1 {
		fmt.Println("num = 1")
	} else if num == 2 {
		fmt.Println("num = 2")
	} else if num == 3 {
		fmt.Println("num = 3")
	} else if num == 4 {
		fmt.Println("num = 4")
	} else {
		fmt.Println("num = 其他")
	}
}

func TestSwitch1() {
	var num int = 3
	switch num {
	case 1:
		fmt.Println("num = 1")
	case 2:
		fmt.Println("num = 2")
	case 3:
		fmt.Println("num = 3")
	case 4:
		fmt.Println("num = 4")
	default:
		fmt.Println("num = 其他")
	}
}

func TestSwitch2() {
	switch num := 3; num {
	case 1:
		fmt.Println("num = 1")
	case 2:
		fmt.Println("num = 2")
	case 3:
		fmt.Println("num = 3")
	case 4:
		fmt.Println("num = 4")
	default:
		fmt.Println("num = 其他")
	}
}

func TestSwitch3() {
	switch str := "hello"; str {
	case "你好", "hello", "hi":
		fmt.Println("str 为 你好或者hello或者hi")
	default:
		fmt.Println("str 为 其他")
	}
}

func TestSwitch4() {
	num := 60
	switch {
	case num >= 0 && num < 50:
		fmt.Println("num >=0 && num <50")
	case num >= 50 && num < 100:
		fmt.Println("num >=50 && num <100")
	case num >= 100:
		fmt.Println("num >= 100")
	default:
		fmt.Println("num < 0")
	}
}

func TestSwitch5() {
	num := 60
	switch {
	case num >= 0 && num < 50:
		fmt.Println("num >=0 && num <50")
	case num >= 50 && num < 100:
		fmt.Println("num >=50 && num <100")
		fallthrough
	case num >= 100:
		fmt.Println("num >= 100")
		fallthrough
	default:
		fmt.Println("num < 0")
	}
}
