package main

import "fmt"

func main() {

	var a int = 2
	var b float64 = 2.5
	var c int = 3
	var d float64 = 3.5

	fmt.Printf("The result of addition between 2 + 3 is %d!\n", a+c)
	fmt.Printf("The result of substraction between 2 - 3 is %d\n", a-c)
	fmt.Printf("The result of multiplication between 2 * 3 is %d\n", a*c)
	fmt.Printf("The result of division between 2 / 3 is %f\n", float64(a)/float64(c)) // Don't forget to cast the type Kings & Queens.
	fmt.Printf("The result of addition between 2.5 + 3.5 is %f!\n", b+d)
	fmt.Printf("The result of substraction between 2.5 - 3.5 is %d!\n", int64(b)-int64(d))
	fmt.Printf("The result of multiplication between 2.5 * 3.5 is %f!\n", b*d)
	fmt.Printf("The result of division between 2.5 / 3.5 is %f!\n", b/d)

}
