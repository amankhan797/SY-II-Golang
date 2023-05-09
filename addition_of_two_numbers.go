// Addition of two given numbers
package main

import "fmt"

var a, b int

func add(a, b int) int {
	return a + b
}
func main() {
	var a, b int
	fmt.Print("Enter first number : ")
	fmt.Scan(&a)
	fmt.Print("Enter Second number : ")
	fmt.Scan(&b)
	fmt.Print("Addition is : ", add(a, b))
}
