package main

import "fmt"

func bubble(xs []int) {
    for i := 0; i < len(xs)-1 ; i++ {
        for j := i+1; j < len(xs); j++ {    
            if xs[j] < xs[i] {
                xs[i], xs[j] = xs[j], xs[i]
            }
        }
    }
}

func main() {
    xs := []int{ 4, 7, 2, 3, 7, 8, 1 }
    fmt.Println("%v", xs)
    bubble(xs)
    fmt.Println("%v", xs)
}