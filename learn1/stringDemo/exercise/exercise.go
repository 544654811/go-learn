package exercise

import "fmt"

func TheoryByte() {
	str := "hello"
	fmt.Printf("index[0]=%c，长度：%d\n", str[0], len(str))
	for index, item := range str {
		fmt.Printf("index[%d]=%c\n", index, item)
	}

	// str[0] = 'a' 提示错误

	bytes := []byte(str)
	bytes[0] = 'a'
	str = string(bytes)
	fmt.Println("修改之后：", str)
}

func TheoryStr() {
	str1 := "hello"
	str2 := "hello，你好吗"
	fmt.Printf("str1 = %s 长度：%d\n", str1, len(str1))
	fmt.Printf("str2 = %s 长度：%d\n", str2, len(str2))

	fmt.Printf("str2 = %s 长度：%d\n", str2, len([]rune(str2)))
	// for index, item := range str {
	// 	fmt.Printf("index[%d]=%c\n", index, item)
	// }
}

func ReverseByte() {
	str := "hello"
	strByte := []byte(str)
	for i := 0; i < len(str)/2; i++ {
		temp := strByte[len(strByte)-i-1] // 末尾的值
		strByte[len(strByte)-i-1] = strByte[i]
		strByte[i] = temp
	}

	str = string(strByte)
	fmt.Println("倒序后：", str)
}

func ReverseRune() {
	str := "hello，你好吗"
	strRune := []rune(str)
	for i := 0; i < len(strRune)/2; i++ {
		temp := strRune[len(strRune)-i-1]
		strRune[len(strRune)-i-1] = strRune[i]
		strRune[i] = temp
	}
	str = string(strRune)
	fmt.Println("倒序后：", str)
}
