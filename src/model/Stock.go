package model

import "log"

//Stock Represents a "Ação"
type Stock struct {
	Name              string `json:"Nome"`
	Value             string `json:"Value"`
	Variation         string `json:"Variation"`
	MinInMonth        string `json:"MinInMonth"`
	MaxInMonth        string `json:"MaxInMonth"`
	LastPayment       string `json:"LastPayment"`
	MonthAppreciation string `json:"MonthAppreciation"`
	Roe               string `json:"Roe"`
	TotalAssets       string `json:"TotalAssets"`
	Valuation         string `json:"Valuation"`
	TotalStocks       string `json:"TotalStocks"`
}

//PrintPtBr data in PtBr
func (stock *Stock) PrintPtBr() {
	log.Printf(stock.Name)
	log.Printf("Valor atual: %s", stock.Value)
	log.Printf("Variação: %s", stock.Variation)
	log.Println("Ultimo valor pago:", stock.LastPayment)
	log.Println("Maxima no mês:", stock.MaxInMonth)
	log.Println("Minima mensal:", stock.MinInMonth)
	log.Println("Apreciação ( ou depreciação ):", stock.MonthAppreciation)
	log.Println("ROE:", stock.Roe)
	log.Println("Patrimonio Liquido:", stock.TotalAssets)
	log.Println("Valor de Mercado:", stock.Valuation)
	log.Println("Total de Ações:", stock.TotalStocks)

}
