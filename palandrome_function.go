// Check if the given  number is palandrome or not using user define function
package main

import "fmt"

var num int
var rem int
var rev int = 0

func pal(num int) {
	var temp int = num
	for num > 0 {
		rem = num % 10
		rev = rev*10 + rem
		num = num / 10
	}
	if rev == temp {
		fmt.Print("Number is Palandrome :D ")
	} else {
		fmt.Print("number is not Palandrome :|")
	}
}
func main() {
	fmt.Print("Enter number : ")
	fmt.Scan(&num)
	pal(num)
}
