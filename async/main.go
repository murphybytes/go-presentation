package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	taskNum   int
	processed int
	resp      chan task
}

// Simulate async task processing we send a task to another goroutine
// for processing, and while we're waiting for task to be processed
// we do other work
func taskSender(taskNum int, ch chan<- task, wg *sync.WaitGroup) {
	defer wg.Done()

	t := task{
		taskNum: taskNum,
		// Note task contains channel that will be used
		// for processor to send completed task to sender
		resp: make(chan task),
	}

	ch <- t

	for {
		select {
		case processedTask := <-t.resp:
			fmt.Printf("Task: %d Process: %d\n", processedTask.taskNum, processedTask.processed)
			return
		default:
			// simulate doing some work while we wait
			time.Sleep(time.Millisecond * 10)
		}
	}

}

func taskProcessor(ch <-chan task) {
	processed := 0
	for task := range ch {
		task.processed = processed
		processed++
		task.resp <- task
	}

}

func main() {
	ch := make(chan task, 10)
	tasks := 20
	var wg sync.WaitGroup
	wg.Add(tasks)

	go taskProcessor(ch)

	for i := 0; i < tasks; i++ {
		go taskSender(i, ch, &wg)
	}

	wg.Wait()
	close(ch)

}
