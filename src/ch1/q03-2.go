package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    str := "asSASA ddd dsjkdsjs dk"    
    runes := []byte(str)
    fmt.Println(len(runes))
    fmt.Println(utf8.RuneCount(runes))

}