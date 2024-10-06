package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "go-math-api/mvc/handlers"
)

func main() {
    r := mux.NewRouter()

    mathHandler := handlers.NewMathHandler()

    r.HandleFunc("/add", mathHandler.Add).Methods("GET")
    r.HandleFunc("/subtract", mathHandler.Subtract).Methods("GET")
    r.HandleFunc("/multiply", mathHandler.Multiply).Methods("GET")
    r.HandleFunc("/divide", mathHandler.Divide).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", r))
}
