// 1.	Write Program in Go to print out the numbers from 0 to 10, waiting 100ms after each one using delay function
package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go numbers()
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("main terminated")
}
