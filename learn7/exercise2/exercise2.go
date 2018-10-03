package exercise2

import (
	"fmt"
	"learn7/student"
	"os"
)

func showMenu() {
	fmt.Println("请选择以下功能")
	fmt.Println("输入 1 ：添加学生")
	fmt.Println("输入 2 ：修改学生")
	fmt.Println("输入 3 ：查询所有学生")
	fmt.Println("输入 4 ：退出 \n\n")
}

func add() {
	var (
		username string
		sex      int
		grade    int
		score    float32
	)
	fmt.Println("请输入用户名")
	fmt.Scanf("%s\n", &username)
	fmt.Println("请输入性别。【0|1】")
	fmt.Scanf("%d\n", &sex)
	fmt.Println("请输入年级。【0~6】")
	fmt.Scanf("%d\n", &grade)
	fmt.Println("请输入分数。【0~100】")
	fmt.Scanf("%f\n", &score)

	var stu student.Student
	stu.Add(username, sex, grade, score)
}

func update() {
	var (
		username string
		sex      int
		grade    int
		score    float32
	)
	fmt.Println("请输入用户名")
	fmt.Scanf("%s\n", &username)
	fmt.Println("请输入性别。【0|1】")
	fmt.Scanf("%d\n", &sex)
	fmt.Println("请输入年级。【0~6】")
	fmt.Scanf("%d\n", &grade)
	fmt.Println("请输入分数。【0~100】")
	fmt.Scanf("%f\n", &score)

	var stu student.Student
	stu.Update(username, sex, grade, score)
}

func all() {
	var stu student.Student
	stu.SelectAll()
}

func Test() {
	for {
		showMenu()
		var cel int
		fmt.Scanf("%d\n", &cel)
		switch cel {
		case 1:
			add()
		case 2:
			update()
		case 3:
			all()
		case 4:
			os.Exit(0)
		}
	}
}
