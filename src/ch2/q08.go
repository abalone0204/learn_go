package main

import (
    "fmt"
    "strconv"
)

type stack struct {
    i int
    data [10]int
}

func (s *stack) pop()(int) {
    s.i--
    return s.data[s.i]
}
func (s *stack) push(x int) {
    s.data[s.i] = x   
    s.i++
}
func (s stack) String() (string){
    var str string
    for i := 0; i <= s.i; i++ {
        str += "["+strconv.Itoa(i)+
               ":"+strconv.Itoa(s.data[i])+
               "]"    
    }
    return str
}

func main() {
    var s stack
    s.push(2)
    s.push(3)
    a := s.pop()
    fmt.Printf("%v\n", s)
    fmt.Println(a)
}