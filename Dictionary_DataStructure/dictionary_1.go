package Dictionary_DataStructure

var operatorDictionary = map[string]string{
	"1": "Airtel",
	"2": "Banglalink",
	"3": "Grameenphone",
	"4": "Robi",
	"5": "Teletalk",
}

func GetOperatorName(key string) string {
	return operatorDictionary[key]
}
