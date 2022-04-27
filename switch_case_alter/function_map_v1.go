package switch_case_alter

import "fmt"

var ServiceFuncMap = map[string]func(string) (string, int){
	"CASE_ONE": func(s string) (string, int) {
		status := 200
		if s == "" {
			s = "blank"
		}
		return s + " CASE_ONE", status
	},
	"CASE_TWO": func(s string) (string, int) {
		status := 200
		if s == "" {
			s = "blank"
		}
		return s + " CASE_TWO", status
	},
}

func IsServiceKeyExist(key string) bool {
	data, _ := ServiceFuncMap[key]
	if data == nil {
		return false
	}
	return true
}

func BackEndServiceCall(data, serviceName string) (string, int) {
	test, code := ServiceFuncMap[serviceName](data)

	return test, code
}

func MapFunction_Example() {
	var response string
	var statusCode int
	isKeyFound := IsServiceKeyExist("CASE_TWO") // CASE_TWO1
	if isKeyFound {
	    response, statusCode := BackEndServiceCall("hello", "CASE_TWO")
	} else {
	    response = "Upcoming ..."
	    statusCode = 200
	}
	
	fmt.Println(response)
	fmt.Println(statusCode)
}
