package main

import (
    "fmt"
)

func main() {
    str := "asSASA ddd dsjkdsjs dk"    
    r := []rune(str)
    copy(r[4:4+3], []rune("哈哈哈"))
    fmt.Printf("%s", string(r))

}