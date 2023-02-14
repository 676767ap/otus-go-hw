package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	wg := &sync.WaitGroup{}
	wg.Add(n)
	errCount := int32(0)
	quitChan := make(chan struct{})
	taskChan := make(chan Task)

	for i := 0; i < n; i++ {
		go worker(quitChan, taskChan, &errCount, m, wg)
	}

	for _, task := range tasks {
		select {
		case <-quitChan:
			wg.Wait()
			return ErrErrorsLimitExceeded
		case taskChan <- task:
		}
	}

	close(taskChan)
	wg.Wait()

	return nil
}

func worker(quitChan chan struct{}, taskChan chan Task, errCount *int32, m int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-quitChan:
			return
		default:
		}

		select {
		case <-quitChan:
			return
		case task, ok := <-taskChan:
			if !ok {
				return
			}

			if err := task(); err != nil {
				if atomic.AddInt32(errCount, 1) == int32(m) {
					close(quitChan)
				}
			}
		}
	}
}
