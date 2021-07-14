package main

import "fmt"

type ShoppingList []string

func main() {
	list := ShoppingList{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}

	printSlice([]string(list))
	printShoppingList(ShoppingList(slice))
}

func printSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

func printShoppingList(list ShoppingList) {
	fmt.Println("Shopping List: ", list)
}
