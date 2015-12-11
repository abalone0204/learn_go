package main

import "fmt"

func max(xs []int) (maxVal int) {
    maxVal = xs[0]
    for _, v := range xs {
        if v> maxVal {
            maxVal = v
        }
    }
    return 
}

func min(xs []int) (minVal int) {
    minVal = xs[0]
    for _, v := range xs {
        if v < minVal {
            minVal = v
        }
    }
    return 
}
func main() {
    xs := []int {1,2,300,4,56,6}
    fmt.Println(max(xs))
    fmt.Println(min(xs))
}