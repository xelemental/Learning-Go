package main

import "fmt"

func main() {

	var number_iseven int64 = 89

	if number_iseven%2 == 0 {

		fmt.Printf("The number %d is even\n", number_iseven)

	} else {

		fmt.Printf("The number %d is not even\n", number_iseven)
	}
}
