// 2.	Write Program in Go to illustrate how to create an anonymous Go-routine.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome!! to Main function")
	// Creating Anonymous Goroutine
	go func() {
		fmt.Println("Welcome!! to TYBCA")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! ")
}
