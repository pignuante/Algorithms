package main

import "fmt"

func main() {
    var number int
    fmt.Scanf("%d", &number)

    for i:= 1; i < 10; i++ {
        fmt.Printf("%d * %d = %d\n", number, i, number * i)
    }

}
