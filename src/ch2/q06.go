package main

import "fmt"

func order(a,b int) ( int, int){
    if a > b {
        return b, a
    }
    return a, b

}
func main() {
    fmt.Println(order(1,2))
    fmt.Println(order(10,5))
}