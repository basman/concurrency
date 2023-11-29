package producerTick

import (
	"fmt"
	"time"
)

type Ticker struct {
	patternMillis []int
}

func New(patternMillis []int) *Ticker {
	return &Ticker{patternMillis: patternMillis}
}

// Run will provide time ticks, roughly in the given interval patterns
func (t Ticker) Run(stop chan bool) chan int64 {
	ch := make(chan int64, 10)

	go func() {
		i := 0
	loop:
		for {
			select {
			case <-stop:
				stop <- true // resend for other consumers
				break loop
			default:
				ch <- int64(i)
				i = (i + 1) % len(t.patternMillis)
				time.Sleep(time.Duration(t.patternMillis[i]) * time.Millisecond)
			}
		}

		fmt.Println("Ticker terminated")
	}()

	return ch
}
