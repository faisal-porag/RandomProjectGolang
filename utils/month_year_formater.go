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
	return "N/A"
}
