package main

import "fmt"

func split(numbers []int, even, odd chan<- int, done chan<- bool) {
	for _, n := range numbers {
		if n%2 == 0 {
			even <- n
		} else {
			odd <- n
		}
	}

	done <- true
}

func main() {
	even, odd := make(chan int), make(chan int)
	done := make(chan bool)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	go split(numbers, even, odd, done)

	var evens, odds []int
	finish := false

	for !finish {
		select {
		case n := <-odd:
			odds = append(odds, n)
		case n := <-even:
			evens = append(evens, n)
		case finish = <-done:
		}
	}

	fmt.Printf("Odds: %v | Evens: %v", odds, evens)

}
