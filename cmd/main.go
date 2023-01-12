
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"stock-server/apis"
	"stock-server/config"
)

func main() {
	config.Load()
	router := mux.NewRouter()
	router.HandleFunc("/stock/{symbol}", apis.GetStockPricesByStockSymbol).Methods("GET")
	log.Printf("Server Listing on port %s", config.Conf.Server.Port)
	log.Fatal(http.ListenAndServe(config.Conf.Server.Port, router))
}