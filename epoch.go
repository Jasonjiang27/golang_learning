
package main

import "fmt"
import "time"

func main() {
    //分别使用带unix或者unxiNano的time.now
    //来获取从自协调世界时起到现在的秒数或者纳秒数
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)
  
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    //你也可以将协调世界时起的整数秒或者纳秒转化为相应的时间
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}


