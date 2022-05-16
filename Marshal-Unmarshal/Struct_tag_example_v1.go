package Marshal-Unmarshal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

/* NOTE
Removing Empty JSON Fields
Most commonly, we want to suppress outputting fields that are unset in JSON. 
Since all types in Go have a “zero value,” some default value that they are set to, 
the encoding/json package needs additional information to be able to tell that some 
field should be considered unset when it assumes this zero value. Within the value 
part of any json struct tag, you can suffix the desired name of your field with ,omitempty to tell 
the JSON encoder to suppress the output of this field when the field is set to the zero value. 
The following example fixes the previous examples to no longer output empty fields:
*/

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

func StructTagExampleV1() {
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
