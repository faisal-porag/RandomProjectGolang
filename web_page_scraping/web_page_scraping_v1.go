
package web_page_scraping

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

//NOTE: GET the goquery command
// go get -u github.com/PuerkitoBio/goquery

func WebPageScrapingVersion1() {

	webPage := "http://webcode.me"
	resp, err := http.Get(webPage)

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		errLog := Body.Close()
		if errLog != nil {
			log.Fatal(errLog)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()
	fmt.Println(title)

	title1 := doc.Find("p").Text()
	fmt.Println(title1)
}
