package main

/*
   The function removeIndex() is not as innocent as it looks.
   Run the unit test or main.go and see for yourself. The
   function is not supposed to modify the slice that belongs
   to the caller.

   Find an implementation that passes all unit tests.
*/

import "fmt"

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	s := []int{7, 1, 12, 3, 89, 5, 6, 13, 8, 9}
	cpy := s[:]
	fmt.Printf("s   before removeIndex: %v\n", s)

	s1 := removeIndex(s[:], 2)
	fmt.Printf("s   after removeIndex: %v (should be same as s before)\n", s)
	fmt.Printf("cpy after removeIndex: %v (should be same as s before)\n", cpy)
	fmt.Printf("value 12 (idx=2) removed: %v\n", s1)
}
