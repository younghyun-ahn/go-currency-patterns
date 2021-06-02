package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(n int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		select {
		case results <- fmt.Sprintf("Worker: %d, Job: %d", n, job):
		}
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan string)

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			worker(n, jobs, results)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			jobs <- i
			time.Sleep(10 * time.Millisecond)
		}
		close(jobs)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
