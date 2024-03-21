package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function for the root URL "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Write HTML content directly to the response writer
        fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <title>My Go Website</title>
            </head>
            <body>
                <h1>Welcome to my Go website!</h1>
                <p>This is a simple example of serving HTML content with Go.</p>
				<img src="https://jeffreyjseaman.com/images/studentsselfie.jpg" height="250" width="200">
            </body>
            </html>
        `)
    })

    // Define a handler function for the "/about" URL
    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        // Write HTML content directly to the response writer
        fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <title>About</title>
            </head>
            <body>
                <h1>About Page</h1>
                <p>This is the about page of my Go website.</p>
            </body>
            </html>
        `)
    })

    // Start the web server on port 8080
    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
