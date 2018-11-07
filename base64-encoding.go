package main

import b64"encoding/base64"
import "fmt"

func main() {
    data := "abc123!?$*&()'-=@~"
    //go同时支持标准的和URL兼容的base64格式。编码需要使用[]byte类型参数，所以要将字符串转成此类型。
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)
    
    //解码可能会返回错误，如果不确定输入信息格式是否正确，那么需要进行错误检查
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    //使用URL兼容的base64格式进行编解码
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))

}
