package main

import (
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("public"))
    http.ListenAndServer(":8080", fs)
}