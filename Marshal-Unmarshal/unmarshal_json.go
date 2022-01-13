package Marshal_Unmarshal

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func JsonUnmarshal() {
	empJsonData := `{"Name":"George Smith","Age":30,"Address":"New york, USA"}`
	empBytes := []byte(empJsonData)
	var emp Response
	json.Unmarshal(empBytes, &emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
	fmt.Println(emp.Address)
}
