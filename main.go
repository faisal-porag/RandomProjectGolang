package main

import (
	"RandomProjectGolang/CSV_FILES"
	dataChart "RandomProjectGolang/DataVisualization_Chart"
	jsonRead "RandomProjectGolang/JsonFileParse"
	pdfGen "RandomProjectGolang/PDF_Generator"
	qrCode "RandomProjectGolang/QR_Code_Generator"
	comp "RandomProjectGolang/composition"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Random services in GolLang")

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
