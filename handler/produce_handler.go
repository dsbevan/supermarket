package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"supermarket/service"
)

type ProduceHandler struct {
	produceGetter  service.ProduceGetter
	produceCreator service.ProduceCreator
	produceDeleter service.ProduceDeleter
}

func NewProduceHandler() *ProduceHandler {
	svc := service.NewProduceService()
	return &ProduceHandler{
		produceGetter:  svc,
		produceCreator: svc,
		produceDeleter: svc,
	}
}

// Handle /produce requests
func (ph *ProduceHandler) HandleProduce(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in produce")
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "in GET")

		if res, err := ph.produceGetter.GetProduce(); err != nil {
			// Handle or send error message
		} else if jsn, err := json.Marshal(res); err != nil {
			// Handle or send error message
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsn)
		}

	case "POST":
		fmt.Fprintln(w, "in POST")

		if res, err := ph.produceCreator.CreateProduce(); err != nil {
			// Handle CreateProduce error
		} else if res == nil {
			// Item already exists
		} else if jsn, err := json.Marshal(res); err != nil {
			// Handle json error
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsn)
		}

	case "DELETE":
		fmt.Fprintln(w, "in DELETE")

		if deleted, err := ph.produceDeleter.DeleteProduce(); err != nil {
			// Handle DeleteProduce error
		} else if deleted == false {
			// Item did not exist
		} else if jsn, err := json.Marshal(deleted); err != nil {
			// Handle json error
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsn)
		}

	default:
		w.Write([]byte("in produce"))

	}
}
