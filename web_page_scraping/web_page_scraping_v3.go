package web_page_scraping

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

type Earthquake struct {
	Timestamp string
	Latitide  string
	Longitude string
	Magnitudo string
	Depth     string
	Territory string
}

var arrayEarthquake []*Earthquake


func PostScrapeDataV3() []*Earthquake {
	doc, err := goquery.NewDocument("http://www.bmkg.go.id/gempabumi/gempabumi-terkini.bmkg")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {

		e := &Earthquake{}

		tr.Find("td").Each(func(ix int, td *goquery.Selection) {
			switch ix {
			case 1:
				e.Timestamp = td.Text()
			case 2:
				e.Latitide = td.Text()
			case 3:
				e.Longitude = td.Text()
			case 4:
				e.Magnitudo = td.Text()
			case 5:
				e.Depth = td.Text()
			case 6:
				e.Territory = td.Text()
			}
		})

		arrayEarthquake = append(arrayEarthquake, e)
	})

	return arrayEarthquake
}


func WebpageScrapingVersion3() {
	data := PostScrapeDataV3()
	if len(data) > 0 {
		for i := 0; i < len(data); i++ {
			fmt.Println(data[i])
		}
	}
}

