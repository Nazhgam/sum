package calculate

import (
	"fmt"
	"sync"

	"github.com/Nazhgam/sum.git/domain"
)

func Calculate(workers int, numbers []domain.Number) {
	taskCh := make(chan domain.Number)
	resultCh := make(chan int, len(numbers))

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(taskCh, resultCh)
		}()
	}

	// Send tasks to workers
	for _, num := range numbers {
		taskCh <- num
	}
	close(taskCh)

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	sum := 0
	for res := range resultCh {
		sum += res
	}

	fmt.Printf("\n***********\n\n Total sum of numbers is: %d\n\n***********\n", sum)
}

func worker(tasks <-chan domain.Number, results chan<- int) {
	for num := range tasks {
		result := num.First + num.Second
		results <- result
	}
}
