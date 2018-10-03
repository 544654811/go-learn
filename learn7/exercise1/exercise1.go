package exercise1

import (
	"fmt"
	"learn7/student"
	"os"
)

var allStudent []*student.Student

func showMenu() {
	fmt.Println("请选择以下功能")
	fmt.Println("输入 1 ：添加学生")
	fmt.Println("输入 2 ：修改学生")
	fmt.Println("输入 3 ：查询所有学生")
	fmt.Println("输入 4 ：退出")
	fmt.Println()
	fmt.Println()
}

func addStudent() {
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

	stu := student.NewStudent(username, sex, grade, score)
	for index, item := range allStudent {
		if item.Username == stu.Username {
			allStudent[index] = stu
			return
		}
	}
	allStudent = append(allStudent, stu)
}

func updateStudent() {
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

	stu := student.NewStudent(username, sex, grade, score)
	for index, item := range allStudent {
		if item.Username == stu.Username {
			allStudent[index] = stu
		}
	}
	fmt.Printf("用户名为%s的学生没有找到\n", username)
}

func selectAll() {
	for _, item := range allStudent {
		fmt.Printf("stu: %#v \n", item)
	}
}

func Test() {
	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			addStudent()
		case 2:
			updateStudent()
		case 3:
			selectAll()
		case 4:
			os.Exit(0)
		}
	}
}
