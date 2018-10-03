package splice

import (
	"fmt"
)

func test1() {
	a := []int{1, 2, 3, 5}
	fmt.Println("初始化方式1：", a)

	b := [5]int{11, 22, 33, 44, 55}
	b1 := b[1:4]
	fmt.Println("初始化方式2：", b1)
	b2 := b[:4]
	fmt.Println("初始化方式2-1：", b2)
	b3 := b[2:]
	fmt.Println("初始化方式2-2：", b3)
	b4 := b[:]
	fmt.Println("初始化方式2-3：", b4)
}

func test2() {
	var a []int
	if a == nil {
		fmt.Println("切片a is nil")
	} else {
		fmt.Println("切片a is not nil")
	}

	b := []int{}
	if b == nil {
		fmt.Println("切片b is nil")
	} else {
		fmt.Println("切片b is not nil, b = ", b)
	}
}

func test3() {
	a := [...]int{2, 23, 11, 34, 65, 213, 12}
	a1 := a[:]
	a2 := a[:]
	fmt.Println("数组a：", a)
	fmt.Println("数组a的切片a1：", a1)
	fmt.Println("数组a的切片a2：", a2)

	// 修改切片元素的值
	a1[1] = 1111
	fmt.Println("修改a1[1]后，数组a：", a)
	fmt.Println("修改a1[1]后，数组a的切片a1：", a1)
	fmt.Println("修改a1[1]后，数组a的切片a2：", a2)
	a2[3] = 2222
	fmt.Println("修改a2[3]后，数组a：", a)
	fmt.Println("修改a2[3]后，数组a的切片a1：", a1)
	fmt.Println("修改a2[3]后，数组a的切片a2：", a2)
}

func test4() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := a[1:6]
	fmt.Printf("数组 a：%v, len: %d, cap: %d \n", a, len(a), cap(a))
	fmt.Printf("切片 b：%v, len: %d, cap: %d \n", b, len(b), cap(b))
	c := b[:cap(b)]
	fmt.Printf("切片 c：%v, len: %d, cap: %d \n", c, len(c), cap(c))
}

func test5() {
	var a []int
	if a == nil {
		fmt.Println("切片a is nil")
	}

	a = append(a, 100)
	fmt.Printf("切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	a = append(a, 200)
	fmt.Printf("切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	a = append(a, 300)
	fmt.Printf("切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	a = append(a, 400)
	fmt.Printf("切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	a = append(a, 500)
	fmt.Printf("切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	a = append(a, 600)
	fmt.Printf("切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
}

func test6() {
	a := []int{1, 2, 3}
	b := []int{11, 22, 33}
	fmt.Printf("切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	fmt.Printf("切片 b：%v, add: %p, len: %d, cap: %d \n", b, b, len(b), cap(b))
	a = append(a, b...)
	fmt.Printf("追加后切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	fmt.Printf("追加后切片 b：%v, add: %p, len: %d, cap: %d \n", b, b, len(b), cap(b))
}

func sumArray(arrs []int) int {
	sum := 0
	for _, item := range arrs {
		sum += item
	}
	return sum
}

func testSumArray() {
	a := [6]int{1, 2, 234, 4234, 432}
	sum := sumArray(a[:])
	fmt.Println("sum: ", sum)
}

func spliceType(arr []int) {
	arr[1] = 1111
}

func testSpliceType() {
	a := [6]int{1, 2, 234, 4234, 432}
	fmt.Println("数组a：", a)
	spliceType(a[:])
	fmt.Println("spliceType()方法后，数组a：", a)
}

func testSpliceCopy() {
	a := []int{1, 2, 3}
	b := []int{11, 22, 33, 44}
	c := []int{111, 222}
	fmt.Printf("切片 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	copy(a, b)
	fmt.Printf("copy b切片后 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
	copy(a, c)
	fmt.Printf("copy c切片后 a：%v, add: %p, len: %d, cap: %d \n", a, a, len(a), cap(a))
}

func Test() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// testSumArray()
	// testSpliceType()
	testSpliceCopy()
}
