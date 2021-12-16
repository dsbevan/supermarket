package service

import (
	"supermarket/types"
)

type ProduceGetter interface {
	GetProduce()
}

type ProduceCreator interface {
	CreateProduce()
}

type ProduceDeleter interface {
	DeleteProduce()
}

type ProduceService struct{}

func (s *ProduceService) GetProduce() []*types.ProduceItem {
	panic("not implemented")
}

func (s *ProduceService) CreateProduce() {

}

func (s *ProduceService) DeleteProduce() {

}

func NewProduceService() *ProduceService {
	return &ProduceService{}
}
