//go 对时间和时间段提供大量的支持，这里是一些例子

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println
    //得到当前的时间
    now := time.Now()
    p(now)

    then := time.Date(
        2009,11,17,20,34,58,651387237,time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    //输出是星期一到日的`weekday`也是支持的。
    p(then.Weekday())

    //这些方法老比较两个时间，分别测试一下是否是之前，之后或者是同一时刻，精确到秒。
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))
   

    //方法sub返回一个duration来表示两个时间点的间隔时间
    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())


    //你可以使用add将时间后移一个时间间隔，或者使用一个－来将时间前移一个间隔
    p(then.Add(diff))
    p(then.Add(-diff))
}
