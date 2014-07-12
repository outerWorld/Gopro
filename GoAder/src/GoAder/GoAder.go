package main

import (
    "fmt"
    "time"
)

func fcgi() {
    fmt.Println("Go fcgi")
}

func main() {
    fmt.Println("GoAder start to run.")
    go fcgi()
    for i := 0; i == 0; i=i {
        fmt.Println("wait")
        time.Sleep(3)
    }
}
