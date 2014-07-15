package main

import (
    "os"
    //"io"
    "fmt"
    //"time"
    //"net/http"
    "Utils"
)

//usage = func(program_name string) {
func usage(program_name string) {
    fmt.Fprintf(os.Stdout, "Usage:%s\n", program_name)
}

func main() {
    args := os.Args
    fmt.Fprintf(os.Stdout, "%s start to run.\n", args[0])
    ini_obj, result := Utils.IniFileInit("test.ini")
    if result == true {
        fmt.Printf("The result of Parsing %s is success\n", "test.ini")
    } else {
        fmt.Printf("The result of Parsing %s is failed\n", "test.ini")
    }
    defer ini_obj.Close()
}
