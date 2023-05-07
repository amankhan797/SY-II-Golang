// WAP in Go to illustrate the concept of returning multiples values from a function.
package main

import "fmt"

var a int
var b int

func Retmulval(a, b int) (int, int, int) {
	return a - b, a * b, a + b
}

func main() {
	fmt.Print("Enter first integer:")
	fmt.Scan(&a)
	fmt.Print("Enter second integer:")
	fmt.Scan(&b)

	var res1, res2, res3 = Retmulval(a, b)

	fmt.Println("Returning multiple values using function-------:")
	fmt.Println("Returning subtraction:", res1)
	fmt.Println("Returning multiplication:", res2)
	fmt.Println("Returning addition:", res3)
}
