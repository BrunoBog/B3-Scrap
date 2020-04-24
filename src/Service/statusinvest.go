package service

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/brunobog/B3-Scrap/src/model"
)

var (
	baseURL string = "https://statusinvest.com.br"
)

// ScrapIndice get data from a specific paper
func ScrapIndice(paperName string) (Paper model.Stock) {
	resp, err := http.Get(baseURL + "/acoes/" + paperName)
	if err != nil {
		log.Printf("[ScrapIndice] - Can´t get response from %v ", baseURL)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("[ScrapIndice] Can´t read body from baseUrl")
	}

	doc.Find("div .paper").Each(func(i int, s *goquery.Selection) {
		Paper.Name = paperName
		Paper.Value = s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong").Text()
		Paper.Variation = s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div.info.special.w-100.w-md-33.w-lg-20 > div > div.w-lg-100 > span > b").Text()
		Paper.MaxInMonth = s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(3) > div > div.d-flex.justify-between > div > span.sub-value").Text()
		Paper.LastPayment = s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(4) > div > div.d-flex.justify-between > div > span.sub-value").Text()
		Paper.MinInMonth = s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(2) > div > div.d-flex.justify-between > div > span.sub-value").Text()
		Paper.MonthAppreciation = s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(5) > div > div.d-flex.justify-between > div > span.sub-value > b").Text()
		Paper.Variation = s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(5) > div > div:nth-child(1) > strong").Text()
		Paper.Roe = s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(16) > div > div > strong").Text()
		Paper.TotalAssets = s.Find("#company-section > div > div.top-info.info-3.sm.d-flex.justify-between.mb-5 > div:nth-child(1) > div > div > strong").Text()
		Paper.Valuation = "R$ " + s.Find("#company-section > div > div.top-info.info-3.sm.d-flex.justify-between.mb-5 > div:nth-child(7) > div > div > strong").Text()
		Paper.TotalStocks = "R$ " + s.Find("#company-section > div > div.top-info.info-3.sm.d-flex.justify-between.mb-5 > div:nth-child(9) > div > div > strong").Text()
	})

	return
}
