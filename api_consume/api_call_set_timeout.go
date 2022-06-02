package api_consume

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

//Set a timeout while making an HTTP request in
func SampleAPITimeOutExample() {
    client := &http.Client{
        Timeout: time.Nanosecond * 1,
    }
    _, err := client.Get("https://google.com")
    if os.IsTimeout(err) {
        fmt.Println("Timeout Happened")
    } else {
        fmt.Println("Timeout Did not Happened")
    }
}
