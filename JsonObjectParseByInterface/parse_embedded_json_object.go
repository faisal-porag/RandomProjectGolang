package JsonObjectParseByInterface

import (
	"encoding/json"
	"fmt"
)

// Parsing Embedded object in JSON

func ParseEmbeddedJsonObjectExample() {
	//Simple Employee JSON object which we will parse
	empJson := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager",
		"address": {
			"city": "Mumbai",
			"state": "Maharashtra",
			"country": "India"
		}
	}`

	// Declared an empty interface 
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empJson), &result)

	address := result["address"].(map[string]interface{})

	//Reading each value by its key
	fmt.Println("Id :", result["id"],
		"\nName :", result["name"],
		"\nDepartment :", result["department"],
		"\nDesignation :", result["designation"],
		"\nAddress :", address["city"], address["state"], address["country"])
}

