// Create a program that calculates the average weight
// of 5 people.

package main

import "fmt"

func main() {
	weights := [5]int{45, 60, 50, 45, 55}
	var sum int
	//var a int
	//len := weights.length()
	sum = 0
	for _, w := range weights {
		sum += w
	}
	a := sum / 5
	fmt.Println("Average of 5 people weight is :", a)

}
