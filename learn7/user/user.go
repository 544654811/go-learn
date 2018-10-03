package user

import (
	"fmt"
)

type User struct {
	Username string
	Sex      string
	Age      int
	int
	Address
}

type Address struct {
	Province string
	City     string
}

func NewUser(username, sex string, age int) *User {
	user := &User{
		Username: username,
		Sex:      sex,
		Age:      age,
	}

	return user
}

func Test() {
	var user User
	user.Username = "张三"
	user.Age = 18
	user.int = 100
	// 方式一
	user.Address = Address{
		Province: "广东",
		City:     "深圳",
	}
	fmt.Printf("user: %#v \n", user)

	// 方式二
	var user2 User
	user2.Username = "张三"
	user2.Age = 18
	user2.int = 100
	user2.Province = "湖北"
	user2.City = "荆州"
	fmt.Printf("user2: %#v \n", user2)
}
