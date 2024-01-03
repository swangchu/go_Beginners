// Calculate your age today given the date of birth

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	var dob int
	fmt.Println("Please enter you date of birth: ")
	fmt.Scan(&dob)
	age := t.Year() - dob
	fmt.Println("Your age is ", age)
}
