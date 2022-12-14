
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
	params := mux.Vars(r)
	stockSymbol := params["symbol"]
	stockExchange := r.FormValue("stock_exchange")

	regexPattern, err := regexp.Compile(`\b[A-Z][A-Z]+\b`)
	if err != nil {
		ErrorResponseFunction(w, http.StatusInternalServerError, "Something Went Wrong")
		return
	}

	if !regexPattern.MatchString(stockSymbol) {
		ErrorResponseFunction(w, http.StatusBadRequest, "Invalid stock symbol")
		return
	}

	if stockExchange == "" {
		stockExchange = "AMEX"
	} else {
		if !regexPattern.MatchString(stockExchange) {
			ErrorResponseFunction(w, http.StatusBadRequest, "Invalid stock exchange")
			return
		}
	}

	url := fmt.Sprintf("https://www.worldtradingdata.com/api/v1/stock_search?search_term=%s&stock_exchange=%s&search_by=symbol&limit=5&api_token=%s", stockSymbol, stockExchange, config.Conf.StockServer.ApiKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ErrorResponseFunction(w, http.StatusInternalServerError, "Something Went Wrong")
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ErrorResponseFunction(w, http.StatusInternalServerError, "Something Went Wrong")
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ErrorResponseFunction(w, http.StatusInternalServerError, "Something Went Wrong")
		return
	}
