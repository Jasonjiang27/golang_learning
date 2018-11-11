package main

import "syscall"
import "os"
import "os/exec"

func main() {
    //在我们例子中，我们将执行ls 命令。go 需要提供哦你 我们需要执行的可执行文件的绝对路径，所以我们将使用 exec.LookPath来得到它 大概是/bin/ls

    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"ls", "-a", "-l", "-h"}
    //Exec同样需要使用环境变量，这里仅提供当前的环境变量
    env := os.Environ()

    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
