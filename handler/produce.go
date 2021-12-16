package handler

import (
	"fmt"
	"net/http"
)

type ProduceHandler struct{}

func HandleProduce(w http.ResponseWriter, r *http.Request) {
	//produceService := service.NewProduceService()
	fmt.Println("in produce")
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "in GET")
	case "POST":
		fmt.Fprintln(w, "in POST")
	case "DELETE":
		fmt.Fprintln(w, "in DELETE")
	default:
		w.Write([]byte("in produce"))
	}
}

func NewProduceHandler() *ProduceHandler {
	return &ProduceHandler{}
}
