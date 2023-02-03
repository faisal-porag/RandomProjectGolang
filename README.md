# RandomProjectGolang
>Data visualization with chart, Create CSV file, Read Write CSV file, Read from json file,
make http request, PDF generator, QR code, web scraping sample, basic rest API and many more in single project ......

First install all package by bellow command
```shell
go get
```


## Error Log Print 
```shell
go get github.com/rs/zerolog/log
```

```shell
log.Error().Err(err).Msg("write your error message ....")
```



```shell
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// an arbitrary json string
	jsonString := `{"name":"Android", "version":13, "code":1}`

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)

	fmt.Println(jsonMap)
	// access it like a map
	fmt.Println(jsonMap["name"])
}
```


#### Apache Kafka and Golang 

https://medium.com/swlh/apache-kafka-with-golang-227f9f2eb818




##### GOLANG Deployment Note 
```
To deploy a Go project in a live server, you can follow these steps:

Choose a hosting provider: You can choose a cloud service provider like AWS, Google Cloud, or Heroku or use a virtual private server (VPS) provider like DigitalOcean, Linode, or Vultr.

Prepare your Go project: Make sure that your Go project can be built and runs without any issues on your local machine. You should also make any necessary changes to the configuration files to make the project run properly in the production environment.

Build the project: Build the binary file for your Go project using the go build command.

Copy the binary to the server: Use a secure file transfer protocol (SFTP) client to upload the binary file to your live server.

Start the binary: Use a terminal to log in to your live server and start the binary file using the ./<binary_name> command.

Configure a reverse proxy: You can use a reverse proxy like Nginx or Apache to handle incoming traffic and direct it to your Go binary. This will allow you to handle HTTPS traffic, manage traffic routing, and handle other tasks.

Monitor and maintain the server: Regularly monitor the server to ensure that the Go binary is running smoothly and make any necessary updates or fixes.

These are the general steps to deploy a Go project in a live server. The exact steps may vary depending on your hosting provider and the specifics of your project.
```







