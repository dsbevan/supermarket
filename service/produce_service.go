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

func (s *ProduceService) CreateProduce() *types.ProduceItem {
	panic("not implemented")
}

func (s *ProduceService) DeleteProduce() bool {
	panic("not implemented")
}

func NewProduceService() *ProduceService {
	return &ProduceService{}
}
