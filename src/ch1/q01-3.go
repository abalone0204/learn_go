package main

import "fmt"

func main() {
    var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
    for _, value := range arr {
        fmt.Println(value)
    }
}