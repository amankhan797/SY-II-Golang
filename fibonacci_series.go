// Program to print fibonacci series
package main

import "fmt"

func main() {
	var num int
	a := 0
	b := 1
	c := b
	fmt.Print("Series is:", a, b)
	for i := 0; i < 20; i++ {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
		}
		a = c
		fmt.Print(b)
	}
}
