package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))


    const goldi int = 345
    fmt.Printf("%+v %T \n", goldi,goldi)

    const(
        a = iota
        b = iota 
        c = iota 
    )
    const (a2 = iota )
    fmt.Println(a,b,c)
    fmt.Println(a2)
}