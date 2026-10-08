package main

import "fmt"

func main() {
	//arrays are values and not reference types
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//range operator is used to create a slice from an array
	//slices are reference types and not values like arrays
	var s []int = primes[1:5]
	fmt.Println(s)
}
