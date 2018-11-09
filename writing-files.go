package main 

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"

)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    check(err)
    
    //对于更细粒度的写入，先打开一个文件。
    f, err := os.Create("/tmp/dat2")
    check(err)
    defer f.Close()

    //你可以写入你想写入的字节切片
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)
    
     //writeString 也是可以的
    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

    //调用sync来将缓冲区的信息写入磁盘
    f.Sync()
    
    //带缓冲的写入器
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    //使用Flush来确保所有缓存的操作一写入底层写入器。
    w.Flush()

}

