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
	switch r.Method {
	case "GET":
		produce := h.produceGetter.GetProduce()
		res := GetProduceResponse{produce}
		if jsn, err := json.Marshal(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(jsn)
		}

	case "POST":
		body := PostProduceRequest{}
		getBody(w, r, &body)

		//TODO check format of produce items in body

		// Fulfill request
		produce := h.producePoster.PostProduce(body.Produce)
		res := PostProduceResponse{produce}
		if jsn, err := json.Marshal(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(jsn)
		}

	case "DELETE":
		fmt.Println("in DELETE")

		body := DeleteProduceRequest{}
		getBody(w, r, &body)

		//TODO check format of produce code in body

		// Fulfill request
		response := DeleteProduceResponse{}
		response.Success = h.produceDeleter.DeleteProduce(body.Code)
		if jsn, err := json.Marshal(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(jsn)
		}

	default:
		w.Write([]byte("in produce"))

	}
}

func getBody(w http.ResponseWriter, r *http.Request, bodyObjectPointer interface{}) {
	// Get body
	if r.Body == nil {
		// No body when there should be
		w.WriteHeader(http.StatusBadRequest)
	}
	b := make([]byte, 2048, 2048)
	if bytesRead, err := r.Body.Read(b); err != nil {
		// Error reading body
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		b = b[0:bytesRead]
	}

	// Parse body
	if err := json.Unmarshal(b, bodyObjectPointer); err != nil {
		// Incorrectly formatted body
		w.WriteHeader(http.StatusBadRequest)
	}
}
