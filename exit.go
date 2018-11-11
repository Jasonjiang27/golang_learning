//使用os.Exit来立即进行代给定状态的退出

package main

import "fmt"
import "os"

func main() {
    //当使用os.Exit时defer将不会执行，所以这里的fmt.Println将永远不会被调用
    defer fmt.Println("!")
    //退出并且退出状态为3。
    os.Exit(3)
}

//注意，不像例如c语言，Go不使用在main中返回一个整数来指明退出状态。如果你想以非零状态退出，那么你就要使用os.Exit。

