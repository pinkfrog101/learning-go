package main

import "fmt"

func main() {
	sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	for sum < 10 {
		sum += sum + 1
	} //init ,post are optional
	fmt.Println(sum)
	s := 1
	for s < 1000 {
		s += s
	}
	fmt.Println(s) // while in go is spelled for
}
