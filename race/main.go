// Simple program with race condition
package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup
// Oops!
var Counter int

func main() {

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {

		value := Counter
		value++
		Counter = value
	}

	Wait.Done()
}
