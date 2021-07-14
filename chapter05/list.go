package main

import "fmt"

type GenericList []interface{}

func (list *GenericList) RemoveIndex(index int) interface{} {
	l := *list
	removed := l[index]

	*list = append(l[0:index], l[index+1:]...)
	return removed
}

func (list *GenericList) RemoveStart() interface{} {
	return list.RemoveIndex(0)
}

func (list *GenericList) RemoveEnd() interface{} {
	return list.RemoveIndex(len(*list) - 1)
}

func main() {
	list := GenericList{
		1, "Café", 42, true, 23, "Bola", 3.14, false,
	}

	fmt.Printf("Original list: \n%v\n\n", list)

	fmt.Printf("Remove from start: %v, after remove: \n%v\n", list.RemoveStart(), list)

	fmt.Printf("Remove from end: %v, after remove: \n%v\n", list.RemoveEnd(), list)

	fmt.Printf("Remove from index 3: %v, after remove: \n%v\n", list.RemoveIndex(3), list)

	fmt.Printf("Remove from index 0: %v, after remove: \n%v\n", list.RemoveIndex(0), list)

	fmt.Printf("Remove from last index: %v, after remove: \n%v\n", list.RemoveIndex(len(list)-1), list)
}
