package service

import (
	"sync"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
)

type taskFunc func(*entity.BoardDTO, string) bool

func worker(taskChan chan taskFunc, resultChan chan bool, doneChan chan struct{}, wg *sync.WaitGroup, board *entity.BoardDTO, player string) {
	defer wg.Done()
	for task := range taskChan {
		select {
		case <-doneChan:
			return
		default:
			result := task(board, player)
			resultChan <- result
		}
	}
}

func checkDispatcher(tasks []taskFunc, board *entity.BoardDTO, player string) bool {
	var wg sync.WaitGroup

	taskChan := make(chan taskFunc, len(tasks))
	resultChan := make(chan bool, len(tasks))
	doneChan := make(chan struct{})

	// Launch worker
	wg.Add(1)
	go worker(taskChan, resultChan, doneChan, &wg, board, player)

	// Define tasks and distribute them to worker
	go func() {
		for _, task := range tasks {
			select {
			case <-doneChan:
				return
			default:
				taskChan <- task
			}
		}
		close(taskChan)
	}()

	// Collect results
	for i := 0; i < len(tasks); i++ {
		if <-resultChan {
			close(doneChan)
			return true
		}
	}

	// Wait for all workers to finish
	wg.Wait()

	return false
}
