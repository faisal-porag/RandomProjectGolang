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
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	rawJSON := []byte(`{
	   "level": "debug",
	   "encoding": "json",
	   "outputPaths": ["stdout"],
	   "errorOutputPaths": ["stderr"],
	   "encoderConfig": {
		 "messageKey": "message",
		 "levelKey": "level",
		 "levelEncoder": "lowercase"
	   }
	 }`,
	)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		log.Fatal(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer func(logger *zap.Logger) {
		errLogger := logger.Sync()
		if errLogger != nil {
			log.Fatal(errLogger)
		}
	}(logger)
	logger.Info("Random services in GolLang With ZAP Logger!")
	//logger.Warn("Custom logger is warning you!")
	//logger.Error("Let's do error instead.")

	//Location and time zone configured
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

	//TODO GenerateULID
	id := utils.GenerateULIDID()
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

	errListener := http.ListenAndServe(":8081", nil)
	if errListener != nil {
		log.Println(errListener)
		return
	}
}
