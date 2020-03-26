package model

//Indice Represents a "Ação"
type Indice struct {
	Name              string `json:"Nome"`
	Value             string `json:"Value"`
	Variation         string `json:"Variation"`
	MinInMonth        string `json:"MinInMonth"`
	MaxInMonth        string `json:"MaxInMonth"`
	LastPayment       string `json:"LastPayment"`
	MonthAppreciation string `json:"MonthAppreciation"`
}
