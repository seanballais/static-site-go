package main

import (
    "fmt"
    "net/http"
    "time"
    "log"
)

func main() {
    http.Handle("/", handler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello. The time is : " + time.Now().Format(time.RFC850))
}
