package main

import (
	"fmt"
	"sync"
)

// Goroutines are lightweight threads managed by the Go runtime, allowing you to perform concurrent operations with minimal overhead.
// Channels provide a way for goroutines to communicate with each other and synchronize their execution.

// Goroutine for processing each data set, and a channel to receive the results from each goroutine or control access to shared resources, preventing race conditions.

func processData(data int, wg *sync.WaitGroup, results chan<-int) {
	defer wg.Done()
	result := data * 2
	results <- result
}

func main() {
	// sync.WaitGroup is used to wait for all goroutines to finish before closing the results channel.
	var wg sync.WaitGroup
	dataSets := []int{1, 2, 3, 4, 5}

	// The results channel is used to receive the results from each goroutine.
	results := make(chan int, len(dataSets))

	// Start a goroutine for each data set.
	for _, data := range dataSets {
		wg.Add(1)
		go processData(data, &wg, results)
	}

	// Wait for all goroutines to finish before closing the results channel.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read the results from the channel.
	for result := range results {
		fmt.Println(result)
	}
}