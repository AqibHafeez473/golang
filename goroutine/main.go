package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Project: Concurrent Task Processor

Goal:
- Simulate multiple workers processing tasks concurrently using goroutines.
- Use channels to send tasks to workers.
- Use WaitGroup to wait for all tasks to finish.

Concepts used:
1. Goroutines (lightweight threads)
2. Channels (communication between goroutines)
3. WaitGroup (wait for all goroutines to finish)
4. Simulated delay to mimic real task processing
*/

// Worker function that processes tasks
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // tell WaitGroup when done

	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		time.Sleep(time.Second) // simulate task processing
		fmt.Printf("Worker %d finished task %d\n", id, task)
	}
}

func main() {

	// STEP 1: Create a channel for tasks
	tasks := make(chan int, 10) // buffered channel with capacity 10

	// STEP 2: WaitGroup to wait for all workers
	var wg sync.WaitGroup

	// STEP 3: Start 3 worker goroutines
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // increment WaitGroup counter for each worker
		go worker(i, tasks, &wg)
	}

	// STEP 4: Send tasks to workers
	numTasks := 5
	for i := 1; i <= numTasks; i++ {
		tasks <- i
		fmt.Println("Task added:", i)
	}

	// STEP 5: Close channel so workers know no more tasks
	close(tasks)

	// STEP 6: Wait for all workers to finish
	wg.Wait()

	fmt.Println("All tasks processed. Program finished!")
}