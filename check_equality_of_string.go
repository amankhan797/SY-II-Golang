// check if the given string is equal or not.
package main

import "fmt"

func main() {
	var str1 string = ""
	var str2 string = ""
	fmt.Print("Enter 1st String : ")
	fmt.Scan(&str1)
	fmt.Print("Enter 2nd String : ")
	fmt.Scan(&str2)
	if str1 == str2 {
		fmt.Print("Strings Are Equal :D ")
	} else {
		fmt.Print("Strings Are not Equal :|")
	}
}
