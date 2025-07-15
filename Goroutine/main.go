package main

import (
	"fmt"
	// "sync"
	"time"
)

func main() {
	jobs := make(chan int, 1000)
	results := make(chan int, 1000)

	for i := range 5 {
		go worker(i+1, jobs, results)
	}

	for i := range 1000 {
		jobs <- i + 1
	}
	close(jobs)

	for i := 0; i < 1000; i++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d: finished job %d\n", id, job)
		results <- job * 2
	}
}
