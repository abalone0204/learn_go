package main

import "fmt"

func main() {
    s := "A"
    for i := 0; i < 100; i++ {
        fmt.Println(s)
        s += "A"
    }
}