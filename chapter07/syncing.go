package chapter07

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func execute(control *sync.WaitGroup) {
	defer control.Done()

	duration := time.Duration(1+rand.Intn(5)) * time.Second
	fmt.Printf("Sleeping for %s...\n", duration)
	time.Sleep(duration)
}

func main() {
	initial := time.Now()
	rand.Seed(initial.UnixNano())

	var control sync.WaitGroup

	for i := 0; i < 5; i++ {
		control.Add(1)
		go execute(&control)
	}

	control.Wait()

	fmt.Printf("Finishing at %s \n", time.Since(initial))
}
