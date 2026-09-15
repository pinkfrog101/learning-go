package main

import "fmt"

func add(x int, y int) int { //x,y int can be written too for they are of same type
	return x + y
}

func main() {
	fmt.Println(add(2, 1))
}
