package main

import "fmt"

func main() {
	var i int = 100
	switch i {
	case -5:
		fmt.Print("-5")
	case 10:
		fmt.Print("10")
	case 20:
		fmt.Print("20")
	default:
		fmt.Print("default")
	}
}
