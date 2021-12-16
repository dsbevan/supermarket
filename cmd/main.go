package main

import (
	"net/http"

	"supermarket/handler"
	. "supermarket/middleware"
)

func main() {
	http.HandleFunc("/supermarket/produce", Authenticate(handler.HandleProduce))
	http.ListenAndServe(":8080", nil)
}
