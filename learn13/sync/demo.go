package demo

import (
	"fmt"
	"sync"
)

var num int
var wq sync.WaitGroup // 使用 waitGroup 确保线程运行完
var lock sync.Mutex   // 使用互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		num++
		lock.Unlock()
	}
	wq.Done()
}

func Test() {
	wq.Add(2)
	go add()
	go add()

	wq.Wait()
	fmt.Println("num=", num)
}
