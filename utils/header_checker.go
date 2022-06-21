package utils

import (
    "fmt"
    "net/http"
)

// TODO THIS IS AN EXAMPLE CODE OF HEADER CHECK
// Check if a particular header is present in a HTTP request
func main() {
    handler := http.HandlerFunc(handleRequest)
    http.Handle("/example", handler)
    http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    headers := r.Header
    val, ok := headers["Content-Type"]
    if ok {
        fmt.Printf("Content-Type key header is present with value %s\n", val)
    } else {
        fmt.Println("Content-Type key header is not present")
    }
    val, ok = headers["someKey"]
    if ok {
        fmt.Printf("someKey key header is present with value %s\n", val)
    } else {
        fmt.Println("someKey key header is not present")
    }
}
