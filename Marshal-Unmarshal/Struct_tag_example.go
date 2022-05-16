package Marshal-Unmarshal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

//Using Struct Tags to Control Encoding

/* NOTE: 
You can modify the previous example to have exported fields that are properly encoded with 
camel-cased field names by annotating each field with a struct tag. The struct tag that encoding/json 
recognizes has a key of json and a value that controls the output. By placing the camel-cased version 
of the field names as the value to the json key, the encoder will use that name instead. This example 
fixes the previous two attempts:
*/

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish"`
	CreatedAt     time.Time `json:"createdAt"`
}

func StructTagExample() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
