// Program to check if the number is even or Odd.
package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Print("Number is Even")
	} else {
		fmt.Print("Number is Odd")
	}
}
