package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting API server for api-app...")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from api-app API!")
    })
    http.ListenAndServe(":8080", nil)
}
