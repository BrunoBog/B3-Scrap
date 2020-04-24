package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brunobog/B3-Scrap/src/service"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		key, ok := r.URL.Query()["stock"]
		if !ok || len(key[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}

		stock := service.ScrapIndice(key[0])

		json, err := json.Marshal(stock)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})
	fmt.Println("Api is up on port 8181")
	http.ListenAndServe(":8181", nil)
}
