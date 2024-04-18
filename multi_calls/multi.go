package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		// Simulate some work
		results <- j * 2 // Sending result back
	}
}

func main() {
	numJobs := 100
	numWorkers := 1

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers

	// pipeline
	var wg sync.WaitGroup

	// invoke workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		fmt.Printf("Sending job %d\n", j)
		// time.Sleep(1 * time.Second)
		jobs <- j
	}
	close(jobs)

	// Collect results
	go func() {
		fmt.Print("Collecting results\n")
		wg.Wait()
		close(results)
	}()

	// Print results
	for r := range results {
		fmt.Println("Result:", r)
	}
}
