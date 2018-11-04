//Go在传统的`printf`中对字符串格式化提供了优异的支持

package main

import "fmt"
import "os"

type point struct {
    x, y int
}

func main() {
    //Go为常规 Go值的格式化设计提供了多种打印方式，例如这里打印了'point'结构体的一个实例。
    p := point{1, 2}
    fmt.Printf("%v\n", p)

    //如果值是一个结构体，`％＋`的格式化输出内容将包括结构体的字段名。
    fmt.Printf("%+v\n", p)

    fmt.Printf("%#v\n", p)

    fmt.Printf("%T\n", p)
    
    fmt.Printf("%t\n", true)

    fmt.Printf("%d\n", 123)

    fmt.Printf("%b\n", 14)

    fmt.Printf("%c\n", 33)

    fmt.Printf("%x\n", 456)

    fmt.Printf("%f\n", 78.9)

    fmt.Printf("%e\n", 1234000000.0)
    fmt.Printf("%E\n", 1234000000.0)

    fmt.Printf("%s\n", "\"string\"")

    fmt.Printf("%q\n", "\"string\"")

    fmt.Printf("%x\n", "hex this")

    fmt.Printf("%p\n", &p)

    fmt.Printf("|%6d|%6d|\n", 12, 345)
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    //sprintf格式化并返回一个字符串，不带任何输出
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
    
    
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
    
