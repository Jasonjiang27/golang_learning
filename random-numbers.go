//go的math/rand包提供了伪随机数生成器

package main

import "fmt"
import "math/rand"

func main() {
    //例如 rand.Intn返回一个随机的整数n, 0《＝n<=100
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    fmt.Println(rand.Float64())

    fmt.Println((rand.Float64()*5)+5, ",")
    fmt.Println((rand.Float64() * 5) +5)

    fmt.Println()

    s1 := rand.NewSource(42)
    r1 := rand.New(s1)

    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
}
