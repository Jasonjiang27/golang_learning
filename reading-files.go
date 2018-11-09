package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程
func check(e error) {
    if e != nil  {
        panic(e)
    }
}

func main() {
    //也许大部分基本文件读取任务是将文件内容读取到内寻中。
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    f, err := os.Open("/tmp/dat")
    check(err)

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)

    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    //你也可以seek到一个文件中已知的位置并从这个位置开始进行读取
    o2, err := f.Seek(6,0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    o3, err := f.Seek(6,0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d : %s\n", n3, o3, string(b3))
    
    //没有内置的回转支持，使用seek(0,0)实现
    _, err = f.Seek(0,0)
    check(err)


    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)

    fmt.Printf("5 bytes: %s\n", string(b4))

    //任务结束后要关闭这个文件，通常可以在open之后立即使用，defer f.close()
    f.Close()

}
