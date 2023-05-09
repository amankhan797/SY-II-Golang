// sum of a given number using recursive function
package main

import "fmt"

var summ int = 0

func sum(num int) {
	var rem int = 0
	if num > 0 {
		rem = num % 10
		summ = summ + rem
		sum(num / 10)
	}
}
func main() {
	var num int
	num = 0
	fmt.Print("Enter number : ")
	fmt.Scan(&num)
	sum(num)
	fmt.Print(summ)
}
