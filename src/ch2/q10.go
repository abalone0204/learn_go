package main

import "fmt"

func fib(n int) []int{
     xs := make([]int, n)
     xs[0], xs[1] = 1, 1
     for i := 2; i < n; i++ {
         xs[i] = xs[i-1] + xs[i-2]
     }
     return xs
    
}

func main() {
    a := fib(10)
    for _, v := range a {
        fmt.Println(v)
    }
}