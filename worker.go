package main

import (
	"context"
	"fmt"
	"sync"
)

func startWorker(
	ctx context.Context,
	wg *sync.WaitGroup,
	id int,
	tasks <-chan Task,
	results chan<- Result,
) {
	defer wg.Done()

	fmt.Printf("Start worker #%d\n", id)

	for {
		select {
		case <-ctx.Done():
			return

		case task, ok := <-tasks:
			if !ok {
				return
			}

			output, err := process(task)

			select {
			case <-ctx.Done():
				return
			case results <- Result{
				TaskID: task.ID,
				Output: output,
				Err:    err,
			}:
			}
		}
	}
}
