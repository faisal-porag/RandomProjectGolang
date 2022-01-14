package Dictionary_DataStructure

type inputValidatorType struct {
	Regex         string
	FailedMessage string
}

var inputValidatorList = map[string]*inputValidatorType{
	"Test":  {"^[1-5]$", "Input invalid"},
	"Test1": {"^[1-5]$", "Input invalid"},
	"Test2": {"^[1-5]$", "Input invalid"},
	"Test3": {"^[1-5]$", "Input invalid"},
}

func DataValidator(key string) *inputValidatorType {
	return inputValidatorList[key]
}
