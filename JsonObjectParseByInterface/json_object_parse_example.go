package JsonObjectParseByInterface

import (
	"encoding/json"
	"fmt"
  "log"
)
// TODO Parsing Embedded object in Array of JSON
func JsonObjectParsing() {
	//Simple Employee JSON object which we will parse
	empArray := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead",
			"address": {
				"city": "Pune",
				"state": "Maharashtra",
				"country": "India"
			}
		}
	]`

	// Declared an empty interface of type Array
	var results []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	unmarshalErr := json.Unmarshal([]byte(empArray), &results)
	if unmarshalErr != nil {
		log.Println("json unmarshal error!")
		return
	}

	for key, result := range results {
		address := result["address"].(map[string]interface{})
		fmt.Println("Reading Value for Key :", key)
		//Reading each value by its key
		fmt.Println("Id :", result["id"],
			"- Name :", result["name"],
			"- Department :", result["department"],
			"- Designation :", result["designation"])
		fmt.Println("Address :", address["city"], address["state"], address["country"])
	}
}
