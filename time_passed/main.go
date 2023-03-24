package main

import (
    "fmt"
    "time"
)

func main() {
    createdAt := time.Now()
    elapsed := time.Since(createdAt)

    if elapsed.Minutes() > 10 {
        fmt.Println("More than 10 minutes have passed since createdAt.")
    } else {
        fmt.Println("Less than 10 minutes have passed since createdAt.")
    }
}

