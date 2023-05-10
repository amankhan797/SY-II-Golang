// 3.	Write Program in Go to accept number and calculate the squares of each digit are calculated in a separate Go-routine, cubes of each digit in another Go-routine and the final summation happens in the main Go-routine [20] Example if number is 123 then squares = (1 * 1) + (2 * 2) + (3 * 3) cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3). Output Sum of squares= 170 Sum of cubes= 1366 Final sum of squares and cubes = 1536
package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}
func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
