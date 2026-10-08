package main

import "fmt"

func main() {
	a := 10
	b := 20
	p := &a
	fmt.Println(*p) //read a
	*p = 30         // rewrite a
	fmt.Println(a)  //print rewritten a

	p = &b          //reassign to b
	fmt.Println(*p) //read b
	*p = 40         //rewrite b
	fmt.Println(b)  //print rewritten b
}
