package main

import "fmt"

func main() {
    // var 声明1 个或者多个变量
    var a string = "initial"
    fmt.Println(a)

    //你可以声明一次性声明多个变量
    var b, c int = 1, 2
    fmt.Println(b, c)

    //Go 将自动推断已经初始化的变量类型
    var d = true
    fmt.Println(d)

    //声明变量且没有给出对应的初始值时，变量将会初始化为零值，例如一个int 的零值是0
    var e int
    fmt.Println(e)

    //:=语句是申明并初始化变量的简写，例如
    //这个例子中的 var f string = "short".
    f := "short"
    fmt.Println(f)
}

