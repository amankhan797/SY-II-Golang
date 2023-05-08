// Program to Table of a given number using for loop.
package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num)
	for i := 0; i < 11; i++ {
		println(num, " X ", i, " = ", num*i)
	}
}
