package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("Random number %g generated\n", rand.Intn(10)) //intn is a function that returns a non-negative pseudo-random number in [0,n). It panics if n <= 0.
}

//int is a built in primitive data type. size is implementation-specific
