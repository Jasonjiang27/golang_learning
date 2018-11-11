package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {
    //我们将从一个简单的命令开始，没有参数或者输入，仅打印一些信息到标准输出流。    //exec.Command 函数帮助我们创建一个表示这个外部进程的对象。
    dateCmd := exec.Command("date")
    // .output是另一个帮助我们处理运行一个命令的常见情况
    // 的函数。它等待命令运行完成，并收集命令的输出。如果没有出错。dateout将获取    //到日期信息的字节
   
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    //下面我们将看看一个稍微复杂的例子。我们将从外部进程的 stdin输入数据并从stdout收集结果.

    grepCmd := exec.Command("grep", "hello")

    //这里我们明确的获取输入/输出管道，运行这个进程，写入一些信息，读取输出结果，最后等待程序运行结束。
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
