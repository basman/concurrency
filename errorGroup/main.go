package main

import (
	"context"
	"fmt"
	"sync"

	"errorgrp/in"
	"errorgrp/out1"
	"errorgrp/out2"
)

/*
TODO introduce error wait groups to simplify the application https://pkg.go.dev/golang.org/x/sync/errgroup

Explanation of the program structure:
 - Wait group wg ensures a synchronised shutdown, avoiding uncontrolled abortion of goroutines as soon as the main goroutine terminates.
 - The channel ch transports user data from producers to the consumer.
 - errCh collects any asynchronous errors that happen after startup of worker goroutines.
 - The context allows distributing cancellation signals over all goroutines, should one goroutine encounter an error.

Behavior:
 - Running this program multiple times will yield different results. One of the following may happen:
	1. Producer 1 fails to start up. The program will abort and not start up any other component of the pipeline.
	2. Producer 2 fails to start up. Producer 1 has already been started and needs to shutdown. A cancellation signal is sent.
       The consumer will start nevertheless but should abort soon.
	3. Producer 1 fails after a couple of items. A cancellation signal ripples through the program, allowing controlled shutdown.
	4. Same for Producer 2.

In all scenarios, the running goroutines are not just cancelled. They get the chance to terminate on their own and
are guaranteed to write a shutdown message. When you introduce the error group, make sure this behavior remains the same.
*/

func main() {
	ch := make(chan int, 2)

	p1 := &out1.Producer1{ch}
	p2 := &out2.Producer2{ch}
	c := &in.Consumer{ch}

	ctx, cancelFunc := context.WithCancel(context.TODO())
	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	err := p1.Run(ctx, &wg, errCh) // launch producer 1
	if err != nil {
		// startup error: abort early
		fmt.Printf("startup error from producer1: %v\n", err)
		cancelFunc()
		return // don't start the rest of the pipeline
	}

	err = p2.Run(ctx, &wg, errCh) // launch producer 2
	if err != nil {
		// startup error: abort early
		fmt.Printf("startup error from producer2: %v\n", err)
		cancelFunc()
		// allow proper shutdown: do not simply return, even if it would prevent the consumer from starting up for nothing.
	}

	c.Run(ctx, errCh) // launch the consumer

	go func() {
		wg.Wait()
		close(ch)
		close(errCh)
	}()

	err = <-errCh
	if err != nil {
		fmt.Printf("pipeline processing error: %v\n", err)
	} else {
		fmt.Println("ok")
	}

	cancelFunc()
	// wait for errCh to close
	<-errCh
}
