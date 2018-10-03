package demo

import (
	"fmt"
)

type Employee interface {
	Calc() float32
}

type Programmer struct {
	name string
	base float32
}

func NewProgrammer(name string, base float32) *Programmer {
	return &Programmer{
		name: name,
		base: base,
	}
}

func (p *Programmer) Calc() float32 {
	return p.base
}

type Sale struct {
	name  string
	base  float32
	extra float32
}

func NewSale(name string, base, extra float32) *Sale {
	return &Sale{
		name:  name,
		base:  base,
		extra: extra,
	}
}

func (s *Sale) Calc() float32 {
	return s.base + s.extra*s.base*0.5
}

func sum(employees []Employee) (sum float32) {
	for _, e := range employees {
		sum += e.Calc()
	}
	return
}

func Test() {
	p1 := NewProgrammer("p1", 7000)
	p2 := NewProgrammer("p2", 7000)
	p3 := NewProgrammer("p3", 7000)

	s1 := NewSale("s1", 700, 1.2)
	s2 := NewSale("s2", 700, 1.2)
	s3 := NewSale("s3", 700, 1.2)

	var employees []Employee
	employees = append(employees, p1)
	employees = append(employees, p2)
	employees = append(employees, p3)
	employees = append(employees, s1)
	employees = append(employees, s2)
	employees = append(employees, s3)

	sum := sum(employees)
	fmt.Println("总工资：", sum)
}
