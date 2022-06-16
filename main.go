package main

import (
	"RandomProjectGolang/CSV_FILES"
	dataChart "RandomProjectGolang/DataVisualization_Chart"
	jsonRead "RandomProjectGolang/JsonFileParse"
	pdfGen "RandomProjectGolang/PDF_Generator"
	qrCode "RandomProjectGolang/QR_Code_Generator"
	comp "RandomProjectGolang/composition"
	"RandomProjectGolang/config"
	"RandomProjectGolang/utils"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {

	fmt.Println("Random services in GolLang")

	loc, _ := time.LoadLocation(config.Config.TimeZone)
	time.Local = loc

	//log.Println(time.Local)

	//TODO JSON FILE READ
	jsonRead.ReadJsonFile() // Console print
	log.Println("---------------------------------------")

	//TODO CREATE & WRITE CSV FILE
	CSV_FILES.CreateCSVFileAndWrite()
	log.Println("---------------------------------------")

	//TODO READ FROM CSV FILE
	CSV_FILES.ReadCSVFile() // Console print
	log.Println("---------------------------------------")

	//TODO GENERATE PDF FILE
	pdfErr := pdfGen.GeneratePdf("test.pdf")
	if pdfErr != nil {
		log.Println(pdfErr)
	}

	//TODO GET COMPOSITION EXAMPLE RESULT
	comp.CompositionFunc()

	//TODO UUID GENERATOR
	newUUID, err1 := exec.Command("uuidgen").Output()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("Generated UUID:")
	fmt.Printf("%s", newUUID)

	//TODO GenerateUlidID
	id := utils.GenerateUlidID()
	fmt.Println(id)

	isValid := utils.ValidatePinNumber("878")
	log.Println(isValid)

	//TODO WORD CLOUD CHART
	http.HandleFunc("/word-cloud-chart", dataChart.CreateWordCloud)
	//TODO LINE CHART WITH RANDOM DATA
	http.HandleFunc("/line-chart", dataChart.LineChartRandomData)
	//TODO BAR CHART
	http.HandleFunc("/bar-chart", dataChart.CreateBarChart)
	//TODO PIE CHART
	http.HandleFunc("/pie-chart", dataChart.CreatePieChart)

	//TODO QR CODE
	http.HandleFunc("/create-qr-code", qrCode.CreateQRCode)
	http.HandleFunc("/view-qr-code/", qrCode.ViewQRCode)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
