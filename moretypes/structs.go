package main

import "fmt"

// uppercase means public
type Vertex struct {
    // want these uppercase to make effectively public
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})
}
