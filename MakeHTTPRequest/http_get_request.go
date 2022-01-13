package MakeHTTPRequest

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateHttpGetRequest() {
	response, err := http.Get("https://api.github.com/users/mojombo")
	if err != nil {
		print(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		print(err)
	}

	fmt.Print(string(body))
}
