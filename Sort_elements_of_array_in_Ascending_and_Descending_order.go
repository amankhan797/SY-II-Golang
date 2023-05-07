// WAP  in Go to sort elements of an array in Ascending and Descending order.
package main

import "fmt"

func main() {
	var myarray = [...]int{1, 3, 5, 7, 9, 2, 4, 6, 8}
	var temp int
	fmt.Println("Original Array is :", myarray)
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if myarray[i] > myarray[j] {
				temp = myarray[i]
				myarray[i] = myarray[j]
				myarray[j] = temp
			}
		}
	}
	fmt.Println("Sorted Array is :", myarray)
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if myarray[i] < myarray[j] {
				temp = myarray[i]
				myarray[i] = myarray[j]
				myarray[j] = temp
			}
		}
	}
	fmt.Print("Sorted Array in decending order is :", myarray)
}
