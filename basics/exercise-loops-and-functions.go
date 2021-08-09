package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for i := 1; i < 100; i++ {
        z_prev := z;
        z -= (z*z - x) / (2*z)
        if z == z_prev {
            return z
        }
    }
    return z
}

func main() {
    fmt.Println(Sqrt(10))
}
