package main

import (
	dataChart "RandomProjectGolang/DataVisualization_Chart"
	jsonRead "RandomProjectGolang/JsonFileParse"
	"log"
	"net/http"
)

func main() {

	//TODO JSON FILE READ
	jsonRead.ReadJsonFile() // Console print

	//TODO WORD CLOUD CHART
	http.HandleFunc("/word-cloud-chart", dataChart.CreateWordCloud)
	//TODO LINE CHART WITH RANDOM DATA
	http.HandleFunc("/line-chart", dataChart.LineChartRandomData)
	//TODO BAR CHART
	http.HandleFunc("/bar-chart", dataChart.CreateBarChart)
	//TODO PIE CHART
	http.HandleFunc("/pie-chart", dataChart.CreatePieChart)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
