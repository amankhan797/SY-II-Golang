package main
import "fmt"
func main() {
    var myarray = [10]int{}
    a := 0
    b := 1
    c := b
    fmt.Print("Fibonacci Serirs without array is: ")
    for i := 0; i < 10; i++ {
        c = b
        b = a + b
        a = c
        myarray[i] = b
        fmt.Print(b, " ")
    }
    fmt.Println()
    fmt.Print("Fibonacci Serirs with array is: ", myarray)
}
