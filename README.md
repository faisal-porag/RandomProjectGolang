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
