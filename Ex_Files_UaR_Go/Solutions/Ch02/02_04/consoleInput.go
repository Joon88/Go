package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    
    // var s string
    // fmt.Scanln(&s)  // scans words from console delimited by space
    // fmt.Println(s)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    str, _ := reader.ReadString('\n') // reads a string until \n is encountered
    fmt.Println(str)
    
    fmt.Print("Enter a number: ")
    str, _ = reader.ReadString('\n')
    f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Value of number:", f)
    }
}