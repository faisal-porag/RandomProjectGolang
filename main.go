package main

import (
	dataChart "RandomProjectGolang/DataVisualization_Chart"
	"log"
	"net/http"
)

func main() {
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
