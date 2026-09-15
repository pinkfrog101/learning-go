package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func swap2(x, y string) (a, b string) { //named return values
	a, b = y, x
	return
}

func main() {
	a, b := swap("a", "b")
	fmt.Println(a, b)
	a, b = swap2("a", "b")
	fmt.Println(a, b)
}
