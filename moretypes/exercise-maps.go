package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCounts(s string) map[string]int {
    fields := strings.Fields(s)

    m := make(map[string]int)
    for v := range fields {
        _, ok := m[fields[v]]
        if ok {
            m[fields[v]] = m[fields[v]] + 1
        } else {
            m[fields[v]] = 1
        }
    }
    return m
}

func main() {
    wc.Test(WordCount)
}
