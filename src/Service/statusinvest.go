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
func ScrapIndice(paperName string) (Paper model.Indice) {
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
		log.Printf(Paper.Name)
		log.Printf("Valor atual: %s", Paper.Value)
		log.Printf("Variação: %s", Paper.Variation)
		log.Println("Ultimo valor pago:", Paper.LastPayment)
		log.Println("Maxima no mês:", Paper.MaxInMonth)
		log.Println("Minima mensal:", Paper.MinInMonth)
		log.Println("Apreciação ( ou depreciação ):", Paper.MonthAppreciation)
		log.Println("Variação:", Paper.Variation)
	})

	return
}
