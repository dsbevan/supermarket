package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"supermarket/service"
	"supermarket/types"
)

type ProduceHandler struct {
	produceGetter  service.ProduceGetter
	producePoster  service.ProducePoster
	produceDeleter service.ProduceDeleter
}

// Returns a new ProduceHandler that uses service.ProduceService
// as the default service implementation.
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
			return
		} else {
			w.Write(jsn)
		}

	case "POST":
		body := PostProduceRequest{}
		if ok := getBody(w, r, &body); !ok {
			return
		}

		// Validate and format request values
		for _, item := range body.Produce {
			if !types.ValidItem(item) {
				badRequest(w, "Invalid code, name, or price format")
				return
			}
		}

		// Fulfill request
		produce := h.producePoster.PostProduce(body.Produce)
		res := PostProduceResponse{produce}
		if jsn, err := json.Marshal(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.Write(jsn)
		}

	case "DELETE":
		// Check for produce code parameter
		values := r.URL.Query()["Produce Code"]
		if len(values) < 1 {
			badRequest(w, "Missing 'Produce Code' url parameter")
			return
		}
		code := values[0]

		// Validate code format
		if !types.ValidCode(code) {
			badRequest(w, "Invalid produce code format")
			return
		}

		// Fulfill request
		response := DeleteProduceResponse{}
		response.Success = h.produceDeleter.DeleteProduce(code)
		if jsn, err := json.Marshal(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.Write(jsn)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Unmarshals the request body into the given request object, sending a 400 response
// if the request body format does not match.
func getBody(w http.ResponseWriter, r *http.Request, bodyObjectPointer interface{}) bool {
	// Get body
	if r.Body == nil {
		// No body when there should be
		badRequest(w, "Missing body")
		return false
	}
	b := make([]byte, 2048, 2048)
	if bytesRead, err := r.Body.Read(b); err != nil && bytesRead == 0 {
		// Error reading body
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)
		return false
	} else {
		b = b[0:bytesRead]
	}

	// Parse body
	if err := json.Unmarshal(b, bodyObjectPointer); err != nil {
		// Incorrectly formatted body
		badRequest(w, "Invalid body format")
		return false
	}
	return true
}

func badRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}
