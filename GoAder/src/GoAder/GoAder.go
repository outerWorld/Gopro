package main

import (
    "os"
    //"io"
    "fmt"
    //"time"
    //"net/http"
)

//usage = func(program_name string) {
func usage(program_name string) {
    fmt.Fprintf(os.Stdout, "Usage:%s\n", program_name)
}

func main() {
    args := os.Args
    fmt.Fprintf(os.Stdout, "%s start to run.\n", args[0])
}
