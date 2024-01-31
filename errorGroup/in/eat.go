package in

import (
	"context"
	"fmt"
)

type Consumer struct {
	Ch chan int
}

func (c Consumer) Run(ctx context.Context, errCh chan error) {
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("consumer cancelled")
				break loop
			case v, ok := <-c.Ch:
				if !ok {
					fmt.Println("consumer channel closed")
					break loop
				}
				fmt.Printf("consuming %v\n", v)
			}
		}

		fmt.Println("consumer shut down")
	}()
}
