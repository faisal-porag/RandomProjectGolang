
package web_page_scraping

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)
/*
func WebPageScrapingVersion4() {
	// the above code

	response, err := http.Get("http://journaldev.com")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		errLog := Body.Close()
		if errLog != nil {
			log.Fatal(errLog)
		}
	}(response.Body)

	//n, err := io.Copy(os.Stdout, response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("Number of bytes copied to STDOUT:", n)

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".articles.article--list.gy--lg h4.article__title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Printf("next")
		txt := s.Text()
		fmt.Printf("Article %d: %s\n\n", i, txt)

		//txt1 := doc.Find(".articles.article--list.gy--lg p").Text()
		//fmt.Printf("%s\n", txt1)

	})

}
*/
