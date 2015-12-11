package main

import "fmt"

func main() {
    i:=1
    for i <= 100 {
        switch {
        case (i % 3) == 0:
            fmt.Println("Fizz")
        case (i % 5) == 0:
            fmt.Println("Buzz")
        default: 
            fmt.Println(i)
        }
        i++
    }
}