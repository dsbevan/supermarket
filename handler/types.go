package handler

import . "supermarket/types"

type GetProduceResponse struct {
	Produce []ProduceItem `json:"produce"`
}

type PostProduceRequest struct {
	Produce []ProduceItem `json:"produce"`
}

type DeleteProduceRequest struct {
	Code string `json:"code"`
}

type DeleteProduceResponse struct {
	Success bool `json:"success"`
}
