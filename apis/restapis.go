
package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"regexp"
	"stock-server/config"
)

type StockResp struct {
	TotalReturned int `json:"total_returned"`
	TotalResults  int `json:"total_results"`
	TotalPages    int `json:"total_pages"`
	Limit         int `json:"limit"`
	Page          int `json:"page"`
	Data          []struct {
		Symbol             string `json:"symbol"`
		Name               string `json:"name"`
		Currency           string `json:"currency"`
		Price              string `json:"price"`
		StockExchangeLong  string `json:"stock_exchange_long"`
		StockExchangeShort string `json:"stock_exchange_short"`
	} `json:"data"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetStockPricesByStockSymbol(w http.ResponseWriter, r *http.Request) {