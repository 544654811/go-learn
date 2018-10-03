package test

import (
	"fmt"
	"time"
)

func TestTime() {
	now := time.Now()
	timestamp := now.Unix()
	nowObj := time.Unix(timestamp, 0)
	fmt.Printf("获取当前时间，返回Time类型：now = %v \n", now)
	fmt.Printf("获取当前时间戳，返回ine64类型：now = %d \n", timestamp)
	fmt.Printf("时间戳转Time类型，返回Time类型：now = %v \n", nowObj)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)
}

func TestConst() {
	nanosecond := time.Nanosecond
	microsecond := time.Microsecond
	millisecond := time.Millisecond
	second := time.Second
	minute := time.Minute
	hour := time.Hour

	fmt.Printf("nanosecond = %d \n", nanosecond)
	fmt.Printf("microsecond = %d \n", microsecond)
	fmt.Printf("millisecond = %d \n", millisecond)
	fmt.Printf("second = %d \n", second)
	fmt.Printf("minute = %d \n", minute)
	fmt.Printf("hour = %d \n", hour)
}

func task() {
	fmt.Println("定时器调用方法")
}

func TestTicker() {
	ticks := time.Tick(time.Second)
	for i := range ticks {
		fmt.Printf("定时器当前时间： %v \n", i)
		task()
	}
}

func TestFormat() {
	nowTime_start := time.Now().UnixNano()
	now := time.Now()
	nowformat := now.Format("2006-01-02 15:04:05")
	fmt.Println("格式化前：", now)
	fmt.Println("now.Format()格式化后：", nowformat)
	nowTime_end := time.Now().UnixNano()
	fmt.Println("now.Format()所用时间（纳秒数）：", nowTime_end-nowTime_start)

	nowTime_start1 := time.Now().UnixNano()
	nowformat2 := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("fmt.Sprintf()格式化后：", nowformat2)
	nowTime_end1 := time.Now().UnixNano()
	fmt.Println("fmt.Sprintf()所用时间（纳秒数）：", nowTime_end1-nowTime_start1)
}

func TestFormat2() {
	nowTime_start := time.Now().UnixNano()
	now := time.Now()
	nowformat2 := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("fmt.Sprintf()格式化后：", nowformat2)
	nowTime_end := time.Now().UnixNano()
	fmt.Println("fmt.Sprintf()所用时间（纳秒数）：", nowTime_end-nowTime_start)
}
