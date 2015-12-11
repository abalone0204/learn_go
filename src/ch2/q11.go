package main

import "fmt"

func Map(f func(int) int, xs []int) []int {
    rs := make([]int, len(xs))
    for idx, v := range xs {
        rs[idx] = f(v)
    }
    return rs
}


func main() {
    f := func(x int) int {
        return x*x
    }
    var xs = []int{ 1, 2, 3, 4, 5, 6}
    rs := Map(f, xs)   
    fmt.Println("%v", xs)
    fmt.Println("%v", rs)
}