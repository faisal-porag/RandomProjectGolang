package utils

import "strings"

func YearDashMonthFormat(mmyy string) string {
	if mmyy != "" { //INPUT EXAMPLE: 1021
		strMonthYear := mmyy
		strMonth := strMonthYear[0:2]
		strYear := strMonthYear[len(strMonthYear)-2:]
		strMonthDashYearFormat := "20" + strYear + "-" + strMonth //OUTPUT EXAMPLE : 2021-10
		return strMonthDashYearFormat
	}
	return "N/A"
}


func MonthSlashYearFormat(mmyy string) string {
	if mmyy != "" { //INPUT EXAMPLE: 1021
		strMonthYear := mmyy
		strMonth := strMonthYear[0:2]
		strYear := strMonthYear[len(strMonthYear)-2:]
		strMonthSlashYearFormat := strMonth + "/" + strYear //OUTPUT EXAMPLE: 10/21
		return strMonthSlashYearFormat
	}
	return "N/A"
}



func YYYYDashMMToMMSlashYYFormat(dateString string) string {
	if dateString != "" { //INPUT EXAMPLE: 2021-07
		splitDate := strings.Split(dateString, "-")
		if len(splitDate) > 1 {
			strMonth := splitDate[1]
			strYear := splitDate[0]
			strYear = strYear[len(strYear)-2:]    // GET LAST 2 VALUES
			strResult := strMonth + "/" + strYear //OUTPUT EXAMPLE: 07/21
			return strResult
		}
		return dateString
	}
	return "N/A"
}
