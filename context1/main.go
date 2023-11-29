package main

import (
	"fmt"
	"time"

	"context1/producerRandom"
	"context1/producerTick"
)

type NumberStream interface {
	Run(stop chan bool) chan int64
}

func Run(rnd NumberStream, ticker NumberStream) chan bool {
	stop := make(chan bool, 1)
	ch1 := ticker.Run(stop)
	ch2 := rnd.Run(stop)

	startTime := time.Now()

	go func() {
		for i := 0; i < 10; i++ {
			t := <-ch1
			r := <-ch2
			fmt.Printf("time: %vms, #%v, value: %v\n", time.Now().Sub(startTime).Milliseconds(), t, r)
		}
	}()

	return stop
}

func main() {
	// create producers
	r := producerRandom.New(4)
	t := producerTick.New([]int{5, 25, 70})

	var stop chan bool

	stop = Run(r, t)

	// issue 1
	time.Sleep(2 * time.Second)

	// issue 2
	stop <- true

	fmt.Println("main goroutine terminated")
}
