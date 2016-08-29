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

	collector := make(chan uint64)

	count := 1000

	// send count sets of numbers to goroutines to be summed
	for i := 0; i < count; i++ {
		nums := make([]byte, 10000)
		crand.Read(nums)

		wg.Add(1)
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
		}(nums, collector, &wg)

	}

	// collect results
	go func(expected int, ch <-chan uint64) {
		var sum uint64
		actual := 0

		for v := range ch {
			sum += v
			actual++
		}

		if actual < expected {
			fmt.Printf("Something bad happened. Recieved %d results, expected %d\n", actual, expected)
			return
		}

		fmt.Printf("Sum = %d\n", sum)

	}(count, collector)

	// all workers are done, signal collection function to exit
	wg.Wait()
	close(collector)

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
