package main

/*
   Run this code with "go run -race main.go" and fix any race condition.

   Don't touch main.go, apply your fix in package oneOfAKind.
   Keep the function signature used by the caller as it is.
   Also keep the behavior. First ID returned shall be 0.

   Remember: don't communicate by sharing, share by communicating.
*/

import (
	"fmt"
	"singleton/oneOfAKind"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- oneOfAKind.GetId()
	}()
	go func() {
		ch <- oneOfAKind.GetId()
	}()

	fmt.Printf("id(1): %v\n", <-ch)
	fmt.Printf("id(2): %v\n", <-ch)
}
