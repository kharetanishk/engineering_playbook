// Worker pool: fixed goroutines jobs channel se kaam uthate hain. Interview ka favourite.
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs { // jobs channel close hone tak kaam lo
		results <- j * j
	}
	_ = id
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	// sirf 3 workers, chahe jobs kitne bhi ho (concurrency limit ka yahi tarika hai)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for i := 1; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)

	// alag goroutine me wait, taaki results padhna block na ho
	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println("sum of squares 1..9 =", sum) // 285
}
