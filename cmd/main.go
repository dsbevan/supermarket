package main

import (
	"fmt"
	"net/http"

	"supermarket/config"
	"supermarket/handler"
	"supermarket/middleware"
	"supermarket/util"
)

const CONFIG_FILENAME = "config.json"

func main() {

	// Load data from config file
	util.LoadConfigFatal(CONFIG_FILENAME)

	http.HandleFunc("/supermarket/produce", middleware.Authenticate(handler.NewProduceHandler().HandleProduce))

	fmt.Printf("Server listening on port %d...\n", config.Config.ListenPort)
	http.ListenAndServe(":"+fmt.Sprint(config.Config.ListenPort), nil)
}
