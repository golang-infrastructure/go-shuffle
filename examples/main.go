package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-shuffle"
)

func main() {

	slice := []int{1, 2, 3, 4, 5}
	shuffle.FisherYatesKnuthShuffle(slice)
	fmt.Println(slice)
	// Output:
	// [5 1 2 3 4]

}
