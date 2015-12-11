package main

import "fmt"

func print_multi_args(numbers ... int) {
    for _, v := range numbers {
       fmt.Println(v)
    }
}

func main() {
    print_multi_args(2, 5, 6)
}