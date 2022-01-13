package MakeHTTPRequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateHttpPostRequest() {
	requestBody, error3 := json.Marshal(map[string]string{
		"username": "test",
		"email":    "test@gmail.com",
	})

	if error3 != nil {
		print(error3)
	}

	response, error2 := http.Post("https://httpbin.org/post",
		"application/json", bytes.NewBuffer(requestBody))
	if error2 != nil {
		print(error2)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)
	body, error1 := ioutil.ReadAll(response.Body)
	if error1 != nil {
		print(error1)
	}

	fmt.Println(string(body))
}
