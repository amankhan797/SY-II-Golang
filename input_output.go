// Simple Go program that take's input and gives output
package main
import "fmt"
func main() {
    var name string
    var rollno int
    var division string
    var collage string
    fmt.Print("Enter your Name :")
    fmt.Scan(&name)

    fmt.Print("Enter your rollno :")
    fmt.Scan(&rollno)

    fmt.Print("Enter your division :")
    fmt.Scan(&division)

    fmt.Print("Enter your collage  :")
    fmt.Scan(&collage)

    fmt.Println("Your name is :", name)
    fmt.Println("Your rollno is :", rollno)
    
