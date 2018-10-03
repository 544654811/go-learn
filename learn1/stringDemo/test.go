package test

import (
	"fmt"
	"strings"
)

var common = "my name is common str"

func TestFan() {
	str := `
		窗前明月光，疑是地上霜
		举头望明月，低头思故乡
	`
	fmt.Println("测试反引号方式： ", str)
}

func TestLen() {
	len := len(common)
	fmt.Println("测试 len() 方法：长度为 ", len)
}

func TestSplit() {
	str := "192.168.0.1;192.168.0.2;192.168.0.3"
	strs := strings.Split(str, ";")
	fmt.Printf("测试 strings.Split() 方法：strs[0]=%v strs[1]=%s str[2]=%v\n", strs[0], strs[1], strs[2])
}

func TestSprintf() {
	str1 := "how are you?"
	str2 := "i am fine."
	str := str1 + str2
	fmt.Println("测试拼接 + 方法： ", str)
	str3 := fmt.Sprintf(str1, str2)
	fmt.Println("测试拼接 fmt.Sprintf() 方法：", str3)
}

func TestContains() {
	if strings.Contains(common, "my") {
		fmt.Println("测试 strings.Contains() 方法，存在元素")
	} else {
		fmt.Println("测试 strings.Contains() 方法，不存在元素")
	}
}

func TestHasPrefix() {
	str := "https://www.baidu.com"
	if strings.HasPrefix(str, "https") {
		fmt.Println("测试 strings.HasPrefix() 方法，存在前缀元素")
	} else {
		fmt.Println("测试 strings.HasPrefix() 方法，不存在前缀元素")
	}
}

func TestHasSuffix() {
	str := "https://www.baidu.com"
	if strings.HasSuffix(str, "com") {
		fmt.Println("测试 strings.HasSuffix() 方法，存在后缀元素")
	} else {
		fmt.Println("测试 strings.HasSuffix() 方法，不存在后缀元素")
	}
}

func TestIndex() {
	str := "https://www.baidu.com"
	index := strings.Index(str, "baidu")
	fmt.Println("测试 strings.Index() 方法，第%s个字符", index)
}

func TestLastIndex() {
	str := "https://www.baidu baidu.com"
	index := strings.LastIndex(str, "baidu")
	fmt.Println("测试 strings.LastIndex() 方法，第%s个字符", index)
}

func TestJoin() {
	strs := []string{"张三", "李四", "王五", "赵六"}
	str := strings.Join(strs, "-")
	fmt.Println("测试 strings.Join() 方法，字符串为：", str)
}
