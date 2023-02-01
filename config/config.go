
package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path"
	"sync"
)

type Service struct {
	Server      Server
	StockServer StockServer
}

type Server struct {
	Port string
}

type StockServer struct {
	ApiKey string
}

var Conf Service
var once sync.Once

func Load() {
	once.Do(func() {
		gp := os.Getenv("GOPATH")
		filePath := path.Join(gp, "/src/stock-server/config.toml")
		_, err := toml.DecodeFile(filePath, &Conf)
		if err != nil {