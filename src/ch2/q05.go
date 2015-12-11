package main

import "fmt"

func avg(xs []float64) (mean float64){
    sum := 0.0
    for _, value := range xs {
        sum += value
    }
    mean = sum/float64(len(xs))
    return 
}

func main() {
    var xs = []float64{ 1, 3, 5, 6 }
    mean := avg(xs)
    fmt.Println(mean)
}