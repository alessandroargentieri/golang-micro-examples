package main

import (
    "log"
    "os"
)

func main() {
    // If the file doesn't exist, create it or append to the file
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
    
    log.Println("Hello world!")
   // log.Warn("This is a Warn")
   // log.Info("This is an Info")
   // log.Debug("This is a Debug")
    
}
