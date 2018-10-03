package student

import (
	"fmt"
)

var AllStudents []*Student

type Student struct {
	Username string
	Sex      int
	Grade    int
	Score    float32
}

func NewStudent(username string, sex, grade int, score float32) (stu *Student) {
	stu = &Student{
		Username: username,
		Sex:      sex,
		Grade:    grade,
		Score:    score,
	}
	return stu
}

func (s *Student) Add(username string, sex, grade int, score float32) {
	stu := NewStudent(username, sex, grade, score)
	for index, item := range AllStudents {
		if item.Username == username {
			AllStudents[index] = stu
		}
	}
	AllStudents = append(AllStudents, stu)
}

func (s *Student) Update(username string, sex, grade int, score float32) {
	stu := NewStudent(username, sex, grade, score)
	for index, item := range AllStudents {
		if item.Username == username {
			AllStudents[index] = stu
		}
	}
}

func (s *Student) SelectAll() {
	for _, item := range AllStudents {
		fmt.Printf("student: %#v \n", item)
	}
}
