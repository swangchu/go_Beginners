// Get a number from the console and check if it’s between 1 and 10.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	num, _ := strconv.Atoi(os.Args[1])

	if num > 1 && num < 10 {
		fmt.Printf("The %d is between 1 and 10.", num)
	} else {
		fmt.Printf("The %d is not between 1 and 10.", num)
	}
}
