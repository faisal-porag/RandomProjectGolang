package api

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Home Page!")
    fmt.Println("Endpoint Hitted: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func SimpleAPIV1Example() {
    //Call handler function
    handleRequests()
}
