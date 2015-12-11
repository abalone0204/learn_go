package main

import (
    "fmt"
)

func main() {
    sum := 0.0
    var xs = [...]float64{ 10, 20, 30} 
    switch len(xs) {
    case 0:
        fmt.Println(0)
    default:
        for _, v := range xs {
            sum += v
        }
        fmt.Println(sum/float64(len(xs)))
    }
}