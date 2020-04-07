package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	var (
		errCount int

		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(n)
	tasksCh := make(chan Task, len(tasks))

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			isDone := false
			for task := range tasksCh {
				err := task()

				mu.Lock()
				if errCount >= m {
					isDone = true
				}
				if err != nil {
					errCount++
				}
				mu.Unlock()

				if isDone {
					return
				}
			}
		}()
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	wg.Wait()
	if errCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
