// 4.	Write Program in Go to create buffered channel to display country names.
package main

import "fmt"

func main() {
	//buffered channel
	countyName := make(chan string, 3)
	//Send country names to buffered channel.
	countyName <- "India"
	countyName <- "USA"
	countyName <- "UK"
	//Receive country names from buffered channel.
	fmt.Println(<-countyName)
	fmt.Println(<-countyName)
	fmt.Println(<-countyName)
}
