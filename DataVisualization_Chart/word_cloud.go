package DataVisualization_Chart

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"net/http"
	"os"
)

var wordCloudData = map[string]interface{}{
	"java":         10000,
	"HTML":         8000,
	"FAISAL":       5000,
	"NATS":         4000,
	"Redis":        3000,
	"Cassandra":    2500,
	"Python":       2000,
	"Django":       1500,
	"Golang":       1000,
	"Postgres":     800,
	"Sandbox":      500,
	"Microservice": 200,
}

// Generate random data for word cloud
func generateWordCloudData(data map[string]interface{}) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func CreateWordCloud(w http.ResponseWriter, _ *http.Request) {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Popular Cryptocurrencies",
			Subtitle: "Spot your favourite coins",
		}))
	wc.AddSeries("wordcloud", generateWordCloudData(wordCloudData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{40, 80},
					Shape:     "cardioid",
				}),
		)
	//GENERATE HTML FILE
	f, _ := os.Create("./DataVisualization_Chart/word_cloud.html")
	_ = wc.Render(f)

	//OPEN WITH BROWSER
	err := wc.Render(w)
	if err != nil {
		log.Println(err)
		return
	}
}
