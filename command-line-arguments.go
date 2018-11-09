// go run hello.go 程序go使用 了 run 和 hello.go 两个参数

package main

import "os"
import "fmt"

func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    //也可以使用标准索引位置方式取得单个参数的值
    arg := os.Args[3]
    fmt.Println(argsWithProg)

    fmt.Println(argsWithoutProg)
    fmt.Println(arg)

}
