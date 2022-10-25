
# stock-server

> Api to fetch stock prices using stock symbol and stock exchange

#### Prerequisite

1. [Install Golang](https://golang.org/doc/install)
2. Setup GOPATH [Link1](https://golang.org/doc/code.html#GOPATH) and [Link2](https://github.com/golang/go/wiki/GOPATH)
3. [Install Dep](https://github.com/golang/dep)

#### Getting Started

1. Clone the repo under `$GOPATH/src`. Run `git clone https://github.com/Gohelraj/stock-server.git stock-server`
2. To install dependencies Run `dep ensure`
5. Add `stockserver` API key in `config.toml` file. If you don't have API key than first [Generate key](https://www.worldtradingdata.com) and than add.
6. Run `go run cmd/main.go`
7. Import `stock-server-postman-collection.json` file in [Postman](https://www.getpostman.com/download?platform=win64) and read the description of APIs to get info about APIs