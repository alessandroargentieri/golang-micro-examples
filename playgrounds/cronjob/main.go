package main

import (
    "fmt"
    "github.com/robfig/cron/v3"
)

func main() {
    c := cron.New()
    c.AddFunc("@every 1m", func() {
        fmt.Println("This job runs every 1 minute.")
    })
    c.AddFunc("0 0 1 * * *", func() {
        fmt.Println("This job runs at 1am on the first day of every month.")
    })
    c.Start()

    // Wait for the cron job to complete
    select {}
}

