package switch_case_alter

import (
	"fmt"
	"log"
)

type SParams struct {
	s string
	i int
}

var ServiceFuncMap = map[string]func(SParams) (string, int){
	"CASE_ONE_PIN": SendMoneyService,
	"CASE_TWO_PIN":     TopUpService,
}

func SendMoneyService(p SParams) (string, int) {
	status := 200
	if p.s == "" {
		p.s = "blank"
	}
	return p.s + " CASE ONE", status
}

func TopUpService(p SParams) (string, int) {
	status := 200
	if p.s == "" {
		p.s = "blank"
	}
	return p.s + " CASE TWO", status
}

func IsServiceKeyExist(key string) bool {
	data, _ := ServiceFuncMap[key]
	if data == nil {
		return false
	}
	return true
}

func BackEndServiceCall(data SParams, serviceName string) (string, int) {
	test, code := ServiceFuncMap[serviceName](data)
	log.Println(test)

	return test, code
}

func MapFunctionExampleV3() {
	fmt.Println("Test")
	obj := SParams{
		s: "Hello",
		i: 200,
	}

	var response string
	var statusCode int

	isKeyFound := IsServiceKeyExist("CASE_ONE_PIN")
	if isKeyFound {
		response, statusCode = BackEndServiceCall(obj, "CASE_ONE_PIN")
	} else {
		response = "Upcoming feature..."
		statusCode = 200
	}

	fmt.Println(response)
	fmt.Println(statusCode)
}
