package out2

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

type Producer2 struct {
	Ch chan int
}

func (p Producer2) Run(ctx context.Context, wg *sync.WaitGroup, chErr chan error) error {
	startupError := rand.Int()%100 < 5
	if startupError {
		return errors.New("producer startup failed")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for i := 0; i < 1000; i++ {
			select {
			case <-ctx.Done():
				break loop
			case p.Ch <- i:
				if rand.Int()%100 > 97 {
					chErr <- fmt.Errorf("producer 2: failed to send %v (artificial failure)", i)
					return
				}
			}
		}

		fmt.Println("producer 2 shut down")
	}()

	return nil
}
