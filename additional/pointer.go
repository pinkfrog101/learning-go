package main

import "fmt"

func main() {
	a := 10
	b := 20
	p := &a
	fmt.Println(*p)
	*p = 30
	fmt.Println(a)

	p = &b
	fmt.Println(*p)
	*p = 40
	fmt.Println(b)
}
