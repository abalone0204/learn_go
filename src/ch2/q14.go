package main

import "fmt"

func plusTwo() (func(v int) int){
    return func(v int) int {
        return v+2
    }
}
func plusX(x int)(func(v int) int) {
    return func(v int) int {
        return v+x
    }
}
func main() {
    p := plusTwo()
    px := plusX(10)
    fmt.Println(p(2))
    fmt.Println(px(2))
}