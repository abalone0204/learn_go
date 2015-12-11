package main

import "fmt"

type Stack struct {
    i int
    data [10]int
}

func (s *Stack) pop()(int) {
    s.i--
    return s.data[i]
}
func (s *Stack) push(x int) {
    s.data[s.i] = x   
    s.i++
}

func main() {
    var s Stack
    s.push(2)
    s.push(3)
    fmt.Printf("%v", s)
}