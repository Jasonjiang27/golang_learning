//go 中如何通过通道来处理信号

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
    
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    //signal.Notify 注册则个给定的通道用于接收特定信号

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    
    //这个go协程执行一个阻塞信号接收操作。当它得到一个值时，它将打印这个值，然后通知程序可以退出。
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    //程序将在这里进行等待，直到它得到了期望的信号，然后退出
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
