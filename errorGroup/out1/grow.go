package out1

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Producer1 struct {
	Ch chan int
}

func (p Producer1) Run(ctx context.Context, wg *sync.WaitGroup, chErr chan error) error {
	startupError := rand.Int()%10 == 1
	if startupError {
		return errors.New("producer startup failed")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for i := 0; i < 1000; i++ {
			time.Sleep(2 * time.Millisecond)

			select {
			case <-ctx.Done():
				break loop
			case p.Ch <- i:
				if rand.Int()%100 > 80 {
					chErr <- fmt.Errorf("producer 1: failed to send %v (artificial failure)", i)
					return
				}
			}
		}

		fmt.Println("producer 1 shut down")
	}()

	return nil
}
