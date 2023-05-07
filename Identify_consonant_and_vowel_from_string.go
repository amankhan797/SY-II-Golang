// Take one string from user, identify consonant and vowel.
package main

import "fmt"

func main() {
	var mystr string
	fmt.Print("Enter your String :")
	fmt.Scan(&mystr)
	for i := 0; i < len(mystr); i++ {
		if mystr[i] == 'A' || mystr[i] == 'E' || mystr[i] == 'I' || mystr[i] == 'O' || mystr[i] == 'U' {
			fmt.Printf("%c : is a vowel \n", mystr[i])
		} else {
			fmt.Printf("%c : is a Consonant \n", mystr[i])
		}
	}
}
