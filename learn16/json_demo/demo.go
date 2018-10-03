package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func toJson(str string) {
	var person *Person
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Println("json:", person)
}

func toString() (str string) {
	person := &Person{
		Name: "zhangsan",
		Age:  18,
		Sex:  "man",
	}
	strtmp, err := json.Marshal(person)
	if err != nil {
		return
	}
	str = string(strtmp)
	fmt.Println("str:", str)
	return
}

func main() {
	toJson(toString())
}
