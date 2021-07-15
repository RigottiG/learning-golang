package main

import (
	"fmt"
	"time"
)

type Operator interface {
	Calculate() int
}

type Sum struct {
	operator1, operator2 int
}

func (s Sum) Calculate() int {
	return s.operator1 + s.operator2
}

func (s Sum) String() string {
	return fmt.Sprintf("%d + %d", s.operator1, s.operator2)
}

type Subtract struct {
	operator1, operator2 int
}

func (s Subtract) Calculate() int {
	return s.operator1 - s.operator2
}

func (s Subtract) String() string {
	return fmt.Sprintf("%d - %d", s.operator1, s.operator2)
}

type Age struct {
	birthYear int
}

func (i Age) Calculate() int {
	return time.Now().Year() - i.birthYear
}

func (i Age) String() string {
	return fmt.Sprintf("Birth year %d", i.birthYear)
}

func main() {
	operators := make([]Operator, 4)

	operators[0] = Sum{10, 20}
	operators[1] = Sum{5, 5}
	operators[2] = Subtract{30, 15}
	operators[3] = Subtract{25, 5}

	fmt.Println("Accumulate values: ", accumulate(operators))

	ages := make([]Operator, 3)
	ages[0] = Age{1999}
	ages[1] = Age{2000}
	ages[2] = Age{2001}

	fmt.Println("Accumulate age values: ", accumulate(ages))
}

func accumulate(operators []Operator) int {
	accumulator := 0

	for _, op := range operators {
		value := op.Calculate()
		fmt.Printf("%v = %d\n", op, value)
		accumulator += value
	}

	return accumulator
}
