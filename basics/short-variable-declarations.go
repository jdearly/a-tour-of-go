package main

import "fmt"

func main() {
    var i, j int = 1, 2
    k := 3 // can only use this inside a function, otherwise doesn't exist
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}
