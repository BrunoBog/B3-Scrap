package main

import (
	"log"
	"os"

	"github.com/brunobog/B3-Scrap/src/service"
)

func main() {
	args := os.Args[1:]
	for _, x := range args {
		log.Println("Procurando numeros para o a ação: ")
		service.ScrapIndice(x)
	}
}
