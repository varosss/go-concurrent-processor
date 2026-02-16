package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	const (
		numWorkers = 5
		numTasks   = 20
	)

	tasks := make(chan Task)
	results := make(chan Result)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go startWorker(ctx, &wg, i, tasks, results)
	}

	go func() {
		defer close(tasks)

		for i := 1; i <= numTasks; i++ {
			select {
			case <-ctx.Done():
				return
			case tasks <- Task{ID: i}:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Printf("Task %d failed: %v\n", result.TaskID, result.Err)
			continue
		}
		fmt.Println(result.Output)
	}

	fmt.Println("All done.")
}
