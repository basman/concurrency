package producerRandom

import (
	"fmt"
	"math/rand"
)

type Rand struct {
	seed int64
	loop int64
	src  rand.Source
}

func New(loop int64) *Rand {
	seed := rand.Int63() // our seed is always random
	r := &Rand{seed: seed, loop: loop}
	r.reseed()
	return r
}

func (r *Rand) reseed() {
	r.src = rand.New(rand.NewSource(r.seed))
	fmt.Println("random number generator has been reset")
}

// Run will generate random numbers, repeating themselves after <loop>
func (r *Rand) Run(stop chan bool) chan int64 {
	ch := make(chan int64, 10)

	go func() {
		i := int64(0)
	loop:
		for {
			select {
			case <-stop:
				stop <- true // resend for other consumers
				break loop
			default:
				ch <- r.src.Int63()
				i++
				if i >= r.loop {
					i = 0
					r.reseed()
				}
			}
		}

		fmt.Println("Rand terminated")
	}()

	return ch
}
