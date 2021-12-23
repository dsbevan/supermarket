package handler

import . "supermarket/types"

type PostProduceRequestBody struct {
	Produce []ProduceItem `json:"produce"`
}

type DeleteProduceRequestBody struct {
	Code string `json:"code"`
}

type DeleteProduceResponse struct {
	Success bool `json:"success"`
}
