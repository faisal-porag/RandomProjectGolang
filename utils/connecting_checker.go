package utils

// Internet connection checker function
func connected() (ok bool) {
    _, err := http.Get("http://clients3.google.com/generate_204")
    if err != nil {
        return false
    }
    return true
}

/*

if !connected() {
    fmt.Println("Check your internet connection")
}

*/
