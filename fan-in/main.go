package main

import (
	"crypto/rand"
	"fmt"
	"sync"
)

// Fan - in split up a bunch of work and multiplex it into a single channel
// Here are task is summing 10 million numbers. We break the task of summing
// these numbers into 1000 chunks, collect the results, then sum them.
func main() {

	collector := make(chan uint64)
	count := 1000
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		nums := make([]byte, 10000)
		// create a bunch of random numbers
		rand.Read(nums)
		wg.Add(1)

		// Sum each subseries in a seperate gorouting
		go func(nums []byte, ch chan<- uint64, wg *sync.WaitGroup) {
			defer wg.Done()
			var sum uint64

			for _, v := range nums {
				sum += uint64(v)
			}
			ch <- sum
		}(nums, collector, &wg)
	}

	// Collect all the results, and print them when we've recieved them
	// all.
	go func(ch <-chan uint64) {
		var sum uint64

		for subsum := range ch {
			sum += subsum
		}

		// print sum when we have all our results
		fmt.Printf("Sum = %d\n", sum)
	}(collector)

	wg.Wait()
	close(collector)

}
