package demo

import (
	"fmt"
)

func hello(a chan int) {
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)
}

func Test() {
	flag := make(chan int)
	go hello(flag)
	fmt.Println("main")
	/*
		for {
			data, ok := <-flag
			if !ok {
				fmt.Println("管道关闭...")
				break
			}
			fmt.Println("从管道中取出元素，", data)
		}
	*/
	for data := range flag {
		fmt.Println("从管道中取出元素，", data)
	}
}
