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
