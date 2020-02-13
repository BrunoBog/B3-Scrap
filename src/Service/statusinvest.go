package Service

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	baseUrl = "https://statusinvest.com.br"
)

func ScrapIndice(paperName string) (Paper model.Indice) {
	resp, err := http.Get(baseUrl + "/" + paperName)
	if err != nil {
		log.Printf("[ScrapIndice] - Can´t get response from %v ", baseUrl)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("[ScrapIndice] Can´t read body from baseUrl")
	}

	doc.Find("div .paper .mt-4").Each(func(i int, s *goquery.Selection) {
		Paper.paperName = paperName
		Paper.Value = s.Find("div .value").Text()
	})
}
