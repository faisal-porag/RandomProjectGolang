package utils

import (
	"net/http"
	"time"
)

//https://github.com/alimasyhur/is-connect/blob/master/server.go

//IsOnline is a function to check wheter an internet connection
//return bool

func IsOnline() bool {
	timeout := time.Duration(5000 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}
	//default url to check connection is http://google.com
	_, err := client.Get("https://google.com")
        if err != nil {
		return false
	}
        return true
}

