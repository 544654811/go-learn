package demo

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func test1() {
	var a map[string]int
	fmt.Printf("a: %v \n", a)
	// a["zhangsan"] = 64  运行时报错
	if a == nil {
		a = make(map[string]int, 16)
		a["zhangsan"] = 64
		fmt.Printf("a: %v \n", a)
	}
}

func test2() {
	a := map[string]int{
		"zhangsan": 1,
		"lisi":     2,
		"wangwu":   3,
		"zhaoliu":  4,
	}
	fmt.Printf("a: %v \n", a)
	fmt.Printf("a: %#v \n", a)

	a["zhangsan"] = 111 // map 修改
	a["qiqi"] = 5       // map 添加
	fmt.Printf("map 修改，a: %v \n", a)
	fmt.Printf("map 添加，a: %v \n", a)
	fmt.Printf("map 查看，zhangsan: %d\n", a["zhangsan"])
	fmt.Printf("a: %#v \n", a)
}

func test3() {
	a := make(map[string]int)
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 11; i++ {
		key := fmt.Sprintf("key%d", i)
		a[key] = rand.Intn(1000)
	}

	for index, item := range a {
		fmt.Printf("a: key=%s, value=%d \n", index, item)
	}
}

func test4() {
	a := map[string]int{
		"zhangsan": 1,
		"lisi":     2,
		"wangwu":   3,
	}

	key := "zhaoliu"
	value := a["zhaoliu"]
	fmt.Printf("a: key=%s, value=%d \n", key, value)

	result, ok := a["zhaoliu"]
	fmt.Printf("a: key=%s, value=%d, ok=%t \n", key, result, ok)

	result, ok = a["wangwu"]
	fmt.Printf("a: key=%s, value=%d, ok=%t \n", key, result, ok)
}

func test5() {
	a := map[string]int{
		"zhangsan": 1,
		"lisi":     2,
		"wangwu":   3,
		"zhaoliu":  4,
	}
	fmt.Printf("a: %v \n", a)
	delete(a, "zhaoliu")
	fmt.Printf("删除后 a: %v \n", a)
}

func test6() {
	a := map[string]int{
		"dzhangsan": 1,
		"flisi":     2,
		"awangwu":   3,
		"czhaoliu":  4,
	}
	fmt.Printf("a: %v \n", a)

	splice := make([]string, 0, len(a))
	for index := range a {
		splice = append(splice, index)
	}
	sort.Strings(splice)

	for _, item := range splice {
		fmt.Printf("排序后 a: key=%s, value=%d \n", item, a[item])
	}
}

func test7() {
	var a map[string][]int
	a = make(map[string][]int, 16)
	fmt.Printf("a: %v \n", a)

	key := "zhangsan"
	value, ok := a[key]
	if !ok {
		value = make([]int, 0, 3)
		value = a[key]
	}

	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	a[key] = value
	fmt.Printf("a: %v \n", a)
}

func test8() {
	var a []map[string]int
	a = make([]map[string]int, 5, 10)
	fmt.Printf("a: %v \n", a)

	var b map[string]int
	b = make(map[string]int, 3)
	b["zhangsan"] = 1
	b["lisi"] = 2
	b["wangwu"] = 3
	a[0] = b
	fmt.Printf("a: %v \n", a)
}

func exercise1() {
	s := "how do you do"
	ss := strings.Split(s, " ")
	fmt.Printf("a: %v \n", ss)

	var result map[string]int
	result = make(map[string]int, len(ss))
	for _, item1 := range ss {
		count := 0
		var key string
		for _, item2 := range ss {
			if item1 == item2 {
				count++
			} else {
				key = item1
			}
		}

		result[key] = count
	}
	fmt.Printf("result: %v \n", result)
}

// 自己写的
func exercise2() {
	rand.Seed(time.Now().UnixNano())
	var a map[int]map[string]string
	a = make(map[int]map[string]string, 1024)
	for i := 1; i < 11; i++ {
		temp, ok := a[i]
		if !ok {
			temp = make(map[string]string, 3)
			// temp = a[i]
		}
		// 存入数据
		temp["name"] = fmt.Sprintf("zhangsan%d", i)
		temp["age"] = fmt.Sprintf("%d", rand.Intn(20))
		temp["score"] = fmt.Sprintf("%d", rand.Intn(120))
		a[i] = temp
	}
	// 输出
	for index, item := range a {
		fmt.Printf("result: id=%d, item=%v \n", index, item)
	}
}

func exercise3() {
	var a map[int]map[string]interface{}
	a = make(map[int]map[string]interface{}, 1024)
	for i := 1; i < 11; i++ {
		temp, ok := a[i]
		if !ok {
			temp = make(map[string]interface{}, 3)
		}
		temp["name"] = fmt.Sprintf("zhangsan%d", i)
		temp["age"] = rand.Intn(20)
		temp["score"] = rand.Float32() * 100.0
		a[i] = temp
	}
	// 输出
	for index, item := range a {
		fmt.Printf("result: id=%d, item=%v \n", index, item)
	}
}

func exercise4() {
	// var result map[string]int
	result := make(map[string]int, 10)
	a := []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	sum := 50
	count := 0
	for _, item1 := range a {
		temp := 0
		for _, item := range []byte(item1) {
			switch item {
			case 'a', 'A', 'e', 'E':
				temp++
			case 'i', 'I':
				temp += 2
			case 'o', 'O':
				temp += 3
			case 'u', 'U':
				temp += 5
			}
		}
		count += temp
		result[item1] = temp
	}
	sy := sum - count
	fmt.Printf("剩余金币: sy=%d \n", sy)
	fmt.Printf("result: %v \n", result)
}

func Test() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// exercise1()
	// exercise2()
	// exercise3()
	exercise4()
}
