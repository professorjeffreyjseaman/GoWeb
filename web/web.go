package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function for the root URL "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    // Define a handler function for the "/about" URL
    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "About page")
    })

    // Start the web server on port 8080
    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
