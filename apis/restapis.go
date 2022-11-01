
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