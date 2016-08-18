package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Fan - in what if one of workers fails?
func main() {

	var wg sync.WaitGroup

	out := make(chan uint64)
	done := make(chan int)

	count := 1000
	wg.Add(count)

	// send count sets of numbers to goroutines to be summed
	for i := 0; i < count; i++ {
		nums := make([]byte, 10000)
		crand.Read(nums)

		// worker func
		go func(nums []byte, ch chan<- uint64, wg *sync.WaitGroup) {
			// signal we're done no matter what
			defer wg.Done()

			if fail() {
				return
			}

			var sum uint64
			for _, v := range nums {
				sum += uint64(v)
			}
			ch <- sum
		}(nums, out, &wg)

	}

	// collect results
	go func(expected int, ch chan uint64, done <-chan int) {
		var sum uint64
		actual := 0

		for {
			select {
			case part := <-ch:
				actual++
				sum += part
			case <-done:
				if actual == expected {
					fmt.Printf("Sum = %d\n", sum)
				} else {
					fmt.Printf("Something went horribly wrong. We got %d results when we expected %d\n", actual, expected)
				}
				break
			}
		}

	}(count, out, done)

	// all workers are done, signal collection function to exit
	wg.Wait()
	done <- 1

}

func fail() bool {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		n := r.Intn(10)

		//fail randomly
		if n == 2 {
			return true
		}

		return false
}
