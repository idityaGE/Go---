package main

import (
	"fmt"
	"sync"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func worker(wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	fmt.Println("Waiting")
	for num := range jobs {
		if isPrime(num) {
			results <- num
		}
	}
}

func collectResult(results <-chan int) int {
	sum := 0
	for prime := range results {
		sum += prime
	}
	return sum
}

func main() {
	var wg sync.WaitGroup
	const numofWorker = 8
	const maxNum = 10000

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for range numofWorker {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}

	go func() {
		for num := range maxNum {
			jobs <- num
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	totalSum := collectResult(results)

	fmt.Println(totalSum)
}
