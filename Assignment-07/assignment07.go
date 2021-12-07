package main

import "fmt"

func main() {

	var just_a_number int64 = 78

	if just_a_number < 67 {
		fmt.Println(just_a_number, "less than 67")
	} else if just_a_number%13 == 0 {
		fmt.Println(just_a_number, "divisible by 13")
	} else if just_a_number%78 == 0 {
		fmt.Println(just_a_number, "is 78 itself")
	} else {
		fmt.Println("This not divisible by 13 and probably less than 67")
	}
}
