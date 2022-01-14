package All_Json_Structure

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

/* Decode (unmarshal) JSON to struct
The json.Unmarshal function in package encoding/json parses JSON data. */

type FruitBasket1 struct {
	Name    string
	Fruit   []string
	Id      int64 `json:"ref"`
	Created time.Time
}

func JsonToStruct() {
	jsonData := []byte(`
	{
		"Name": "Standard",
		"Fruit": [
			"Apple",
			"Banana",
			"Orange"
		],
		"ref": 999,
		"Created": "2018-04-09T23:00:00Z"
	}`)

	var basket FruitBasket1
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	fmt.Println(basket.Created)
}
