package Marshal-Unmarshal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name          string
	Password      string
	PreferredFish []string
	CreatedAt     time.Time
}

func MarshalIndentExample() {
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
