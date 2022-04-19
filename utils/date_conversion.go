package utils

import "strings"

func YearDashMonthFormat(mmyy string) string {
	if mmyy != "" { //INPUT EXAMPLE: 1021
		strMonthYear := mmyy
		strMonth := strMonthYear[0:2]
		strYear := strMonthYear[len(strMonthYear)-2:]
		strMonthDashYearFormat := "20" + strYear + "-" + strMonth //EXAMPLE : 2021-10
		return strMonthDashYearFormat
	}
	return ""
}

func MonthSlashYearFormat(mmyy string) string {
	if mmyy != "" { //INPUT EXAMPLE: 1021
		strMonthYear := mmyy
		strMonth := strMonthYear[0:2]
		strYear := strMonthYear[len(strMonthYear)-2:]
		strMonthSlashYearFormat := strMonth + "/" + strYear //EXAMPLE: 10/21
		return strMonthSlashYearFormat
	}
	return ""
}

func YYYYDashMMToMMSlashYYFormat(dateString string) string {
	if dateString != "" { //INPUT EXAMPLE: 2021-07
		splitDate := strings.Split(dateString, "-")
		if len(splitDate) > 0 {
			strMonth := splitDate[1]
			strYear := splitDate[0]
			strYear = strYear[len(strYear)-2:]    // GET LAST 2 VALUE
			strResult := strMonth + "/" + strYear //EXAMPLE: 10/21
			return strResult
		}
	}
	return ""
}
