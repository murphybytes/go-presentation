package main

// worker or fan-out send to single channel, multiple workers read from
// channel and perform work.
import (
	"fmt"
	"sync"
)

func worker(ch <-chan int, wait *sync.WaitGroup) {
	defer wait.Done()
	for {

		i, ok := <-ch
		// Ok is false if channel is closed AND all values have been processed
		if !ok {
			return
		}

		fmt.Printf("Did task %d\n", i)

	}

}

func pool(wg *sync.WaitGroup, workers int, tasks int) {
	ch := make(chan int)
	// kick off worker goroutines
	for i := 0; i < workers; i++ {
		go worker(ch, wg)
	}
	// send workers work
	for i := 0; i < tasks; i++ {
		ch <- i
	}

	// close channel to signal workers we're done
	close(ch)

}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	pool(&wg, 5, 25)
	wg.Wait()

}
