package main

import (
    "fmt"
    "os"
)

func main(){

    target := "World"
    if len(os.Args) > 1 { /* os.Args*/
        target = os.Args[1]
    }
    fmt.Println("hello", target)
}
