package stack

import (
	"fmt"
)

//  先进后出的栈
type Stack struct {
	data [1024]string
	top  int
}

func (s *Stack) Push(d string) {
	s.data[s.top] = d
	s.top++
}

func (s *Stack) Pop() (data string, err error) {
	if s.top == 0 {
		err = fmt.Errorf("stack is empty")
		return
	}
	s.top--
	data = s.data[s.top]
	return
}

func (s *Stack) Top() (data string, err error) {
	if s.top == 0 {
		err = fmt.Errorf("stack is empty")
		return
	}
	data = s.data[s.top-1]
	return
}

func (s *Stack) Emtyp() bool {
	return s.top == 0
}
