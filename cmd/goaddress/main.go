package main

import (
    "net/http"
    // "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/covenroven/goaddress/internal/router"
)

func main() {
    // Load .env
    err := godotenv.Load()
    if err != nil {
        panic(err.Error())
    }

    // Initialize log
    file, err := os.OpenFile("address.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        panic("Failed to open log")
    }
    defer file.Close()
    log.SetOutput(file)

    // Initialize router
    r, err := router.Init()
    if err != nil {
        log.Fatal("Failed to initialize router", err)
    }

    log.Fatal(http.ListenAndServe(":3100", r))
}
