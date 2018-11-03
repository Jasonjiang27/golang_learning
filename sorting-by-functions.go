

package main

import "sort"
import "fmt"

type Bylength []string

func (s Bylength) Len() int {
    return len(s)
}

func (s Bylength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s Bylength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}


func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(Bylength(fruits))
    fmt.Println(fruits)
}


