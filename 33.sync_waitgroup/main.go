package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)

	// some task is happing
	fmt.Printf("worker %d end\n", i)

}

// worker 1
// worker 2
// worker 3
func main() {
	// fmt.Println("Learning SYNC WAITGROUP")

	var wg sync.WaitGroup

	// Start 3 worker goroutin
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}
	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("worker task completed")
}
