package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	values []interface{}
}

func (stack Stack) Length() int {
	return len(stack.values)
}

func (stack Stack) Empty() bool {
	return stack.Length() == 0
}

func (stack *Stack) Up(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Unstack() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("Stack is empty")
	}

	value := stack.values[stack.Length()-1]
	stack.values = stack.values[:stack.Length()-1]

	return value, nil
}

func main() {
	stack := Stack{}

	fmt.Println("Stack created with length ", stack.Length())
	fmt.Println("Empty? ", stack.Empty())

	stack.Up("Go")
	stack.Up(2021)
	stack.Up(3.14)
	stack.Up("End")

	fmt.Println("Length after stack up: ", stack.Length())
	fmt.Println("Empty? ", stack.Empty())

	for !stack.Empty() {
		v, _ := stack.Unstack()
		fmt.Println("Unstack: ", v)
		fmt.Println("Length: ", stack.Length())
		fmt.Println("Empty? ", stack.Empty())
	}

	_, err := stack.Unstack()

	if err != nil {
		fmt.Println(err)
	}

}
