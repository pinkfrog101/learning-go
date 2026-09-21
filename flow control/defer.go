package main

import "fmt"

func main() {
	defer fmt.Println("def")
	fmt.Println("abc")
}

// defers exectution of the function until surrounding func returns
// deferred call's argument are evaluated immediately but function call isnt executed until surrounding function returns
// deferred functions are executed in LIFO order (into a stack)
