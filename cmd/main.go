package main

import (
	"net/http"

	"supermarket/handler"
	"supermarket/middleware"
	"supermarket/util"
)

const CONFIG_FILENAME = "config.json"

func main() {

	// Load data from config file
	util.LoadConfigFatal(CONFIG_FILENAME)

	http.HandleFunc("/supermarket/produce", middleware.Authenticate(handler.NewProduceHandler().HandleProduce))
	http.ListenAndServe(":8080", nil)
}
