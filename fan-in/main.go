package main

import (
	"crypto/rand"
	"fmt"
)

// Fan - in split up a bunch of work and multiplex it into a single channel
// Here are task is summing 10 million numbers. We break the task of summing
// these numbers into 1000 chunks, collect the results, then sum thm.
func main() {

	out := make(chan uint64)
	count := 1000

	for i := 0; i < count; i++ {
		nums := make([]byte, 10000)
		// create a bunch of random numbers
		rand.Read(nums)

		go func(nums []byte, ch chan uint64) {
			var sum uint64
			for _, v := range nums {
				sum += uint64(v)
			}
			ch <- sum
		}(nums, out)
	}

	var sum uint64
	recieved := 0

	for subsum := range out {
		sum += subsum
		recieved++
		// break when we've recieved all messages or program will hang
		if recieved >= count {
			break
		}
	}

	fmt.Printf("Sum = %d\n", sum)

}
