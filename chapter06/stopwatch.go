package main

import (
	"fmt"
	"time"
)

func MakeFibonacci(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b

			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}

func Stopwatcher(function func()) {
	initial := time.Now()

	function()

	fmt.Printf("\n Execution time: %s\n\n", time.Since(initial))
}

func main() {
	Stopwatcher(MakeFibonacci(8))
	Stopwatcher(MakeFibonacci(48))
	Stopwatcher(MakeFibonacci(88))
}
