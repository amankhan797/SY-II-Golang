// Program to Swap two numbers without any temp variable.
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Println("")
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println("Numbers After Swap")
	fmt.Println("first number is : ", num1)
	fmt.Println("second number is : ", num2)
}
