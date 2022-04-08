package JsonFileParse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Employees struct {
	Employees []Employee `json:"employee"`
}

type Employee struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func ReadJsonFile() {
	jsonFile, err := os.Open("./JsonFileParse/employees.json")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Successfully Opened json file")
	defer jsonFile.Close()

	byteEmpValue, _ := ioutil.ReadAll(jsonFile)

	var employees Employees

	unmarshalErr := json.Unmarshal(byteEmpValue, &employees)
	if unmarshalErr != nil {
		log.Println("unmarshal error:",unmarshalErr)
		return
	}

	for i := 0; i < len(employees.Employees); i++ {
		fmt.Println("Employee Name: " + employees.Employees[i].Name)
		fmt.Println("Employee Gender: " + employees.Employees[i].Gender)
		fmt.Println("Employee Age: " + strconv.Itoa(employees.Employees[i].Age))
	}

}
