package Dynamic_Json

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

/* TODO NOTE:-> Gabs is a small utility for dealing with dynamic or unknown JSON structures in Go. */

func DynamicJsonRead() {
	data := []byte(`{
		"employees":{
		   "protected":false,
		   "address":{
			  "street":"22 Saint-Lazare",
			  "postalCode":"75003",
			  "city":"Paris",
			  "countryCode":"FRA",
			  "country":"France"
		   },
		   "employee":[
			  {
				 "id":1,
				 "first_name":"Jeanette",
				 "last_name":"Penddreth"
			  },
			  {
				 "id":2,
				 "firstName":"Giavani",
				 "lastName":"Frediani"
			  }
		   ]
		}
	 }`)

	jsonParsed, err := gabs.ParseJSON(data)
	if err != nil {
		panic(err)
	}

	// Search JSON
	fmt.Println("Get value of Protected:\t", jsonParsed.Path("employees.protected").Data())
	fmt.Println("Get value of Country:\t", jsonParsed.Search("employees", "address", "country").Data())
	fmt.Println("ID of first employee:\t", jsonParsed.Path("employees.employee.0.id").String())
	fmt.Println("Check Country Exists:\t", jsonParsed.Exists("employees", "address", "countryCode"))

	// Iterating address objects
	for key, child := range jsonParsed.Search("employees", "address").ChildrenMap() {
		fmt.Printf("Key=>%v, Value=>%v\n", key, child.Data().(string))
	}

	// Iterating employee array
	for _, child := range jsonParsed.Search("employees", "employee").Children() {
		fmt.Println(child.Data())
	}

	// Use index in your search
	for _, child := range jsonParsed.Search("employees", "employee", "0").Children() {
		fmt.Println(child.Data())
	}
}
