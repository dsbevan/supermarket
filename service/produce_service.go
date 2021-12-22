package service

import (
	"supermarket/dao"
	. "supermarket/types"
)

type ProduceService struct {
	produceGetter  dao.ProduceGetter
	producePoster  dao.ProducePoster
	produceDeleter dao.ProduceDeleter
}

func NewProduceService() *ProduceService {
	dao := dao.NewProduceDao()
	return &ProduceService{
		produceGetter:  dao,
		producePoster:  dao,
		produceDeleter: dao,
	}
}

// Implements service.ProduceGetter
func (s *ProduceService) GetProduce() ([]ProduceItem, error) {
	panic("not implemented")
}

// Implements service.ProducePoster
func (s *ProduceService) PostProduce(item *ProduceItem) (*ProduceItem, error) {
	panic("not implemented")
}

// Implements service.ProduceDeleter
func (s *ProduceService) DeleteProduce(code string) (bool, error) {
	panic("not implemented")
}
