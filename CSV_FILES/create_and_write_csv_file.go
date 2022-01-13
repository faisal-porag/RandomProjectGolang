package CSV_FILES

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var empDataWrite = [][]string{
	{"Name", "City", "Skills"},
	{"Smith", "Newyork", "Java"},
	{"William", "Paris", "Golang"},
	{"Rose", "London", "PHP"},
}

type empData struct {
	Name   string
	City   string
	Skills string
}

func CreateCSVFileAndWrite() {

	csvFile, err := os.Create("./CSV_FILES/employee.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvWriter := csv.NewWriter(csvFile)

	for _, empRow := range empDataWrite {
		_ = csvWriter.Write(empRow)
	}
	csvWriter.Flush()
	csvErr := csvFile.Close()
	if csvErr != nil {
		log.Println(csvErr)
		return
	} else {
		fmt.Println("CSV FILE CREATION SUCCESS")
	}
}

func ReadCSVFile() {
	csvFile, err := os.Open("./CSV_FILES/employee.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(csvFile)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		emp := empData{
			Name:   line[0],
			City:   line[1],
			Skills: line[2],
		}
		fmt.Println(emp.Name + " " + emp.City + " " + emp.Skills)
	}
}
