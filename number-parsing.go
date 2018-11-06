// ???????????????????????????
// Go ????????

package main

// ??? `strconv` ???????????
import "strconv"
import "fmt"

func main() {

    // ?? `ParseFloat` ????????? `64` ?????
    // ???????
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // ??? `ParseInt` ????????????? `0` ?
    // ??????????????????`64` ?????
    // ????? 64 ?????
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // `ParseInt` ????????????
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // `ParseUint` ??????
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // `Atoi` ?????? 10 ??????????
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // ???????????????????
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
