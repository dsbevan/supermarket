package main

import (
	"net/http"

	"supermarket/handler"
	. "supermarket/middleware"
)

func main() {
	http.HandleFunc("/supermarket/produce", Authenticate(handler.NewProduceHandler().HandleProduce))
	http.ListenAndServe(":8080", nil)
}
