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

func BackEndServiceCall(data, serviceName string) (string, int) {
	test, code := ServiceFuncMap[serviceName](data)

	return test, code
}

func MapFunction_Example() {
	response, statusCode := BackEndServiceCall("hello", "CASE_TWO")
	fmt.Println(response)
	fmt.Println(statusCode)
}
