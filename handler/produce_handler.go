package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"supermarket/service"
)

type ProduceHandler struct {
	produceGetter  service.ProduceGetter
	producePoster  service.ProducePoster
	produceDeleter service.ProduceDeleter
}

func NewProduceHandler() *ProduceHandler {
	svc := service.NewProduceService()
	return &ProduceHandler{
		produceGetter:  svc,
		producePoster:  svc,
		produceDeleter: svc,
	}
}

// Handle /produce requests
func (h *ProduceHandler) HandleProduce(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in produce")
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "in GET")

		res := h.produceGetter.GetProduce()
		if jsn, err := json.Marshal(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsn)
		}

	case "POST":
		fmt.Fprintln(w, "in POST")

		body := PostProduceRequestBody{}
		getBody(w, r, body)

		//TODO check format of produce items in body

		// Fulfill request
		res := h.producePoster.PostProduce(body.Produce)
		if jsn, err := json.Marshal(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsn)
		}

	case "DELETE":
		fmt.Fprintln(w, "in DELETE")

		body := DeleteProduceRequestBody{}
		getBody(w, r, body)

		//TODO check format of produce code in body

		// Fulfill request
		response := DeleteProduceResponse{}
		response.Success = h.produceDeleter.DeleteProduce(body.Code)
		if jsn, err := json.Marshal(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsn)
		}

	default:
		w.Write([]byte("in produce"))

	}
}

func getBody(w http.ResponseWriter, r *http.Request, bodyObject interface{}) {
	// Get body
	if r.Body == nil {
		// No body when there should be
		w.WriteHeader(http.StatusBadRequest)
	}
	b := make([]byte, 0, 512)
	if _, err := r.Body.Read(b); err != nil {
		// Error reading body
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Parse body
	if err := json.Unmarshal(b, &bodyObject); err != nil {
		// Incorrectly formatted body
		w.WriteHeader(http.StatusBadRequest)
	}
}
