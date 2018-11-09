//一个行过滤器 在读取标准输入流的输入，处理该输入，然后将得到一些的结果输出到标///准的程序中将是常见的一个功能。grep 和 sed 是常见的行过滤器

//这里是一个行过滤器示例，它将所有的输入文字转化为大写的版本

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    //对 os.stdin使用一个带缓冲的scanner,让我们可以直接使用方便的 scan方法来直接读取一行，每次调用该方法可以让 scanner读取下一行

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        //text 返回当前的 token 现在是输入的下一行
        ucl := strings.ToUpper(scanner.Text())
        
        fmt.Println(ucl)
    }

    //检查 scan 的错误。文件结束符是可以接受的，并且不会被scan当作一个错误
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
         os.Exit(1)
    }
}

