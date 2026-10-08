package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; //identify enviroment it is running on so that it could adapt its behaviour accordingly
	// reflects target enviroment of compiler not necessarily the host machine
	// like here it shows windows but in website it shows linux coz it is running on linux server
	os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	// switch doesnt need break here
	// evaluates from top to bottom ,stopping at succession
	// doent just do integer or character but can also do string and boolean

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away.")
	}
	t := time.Now() //switch with no condition specified is same as switch true
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
