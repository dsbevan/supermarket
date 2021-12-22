package service

import (
	"supermarket/dao"
	. "supermarket/types"
)

type ProduceService struct {
	produceGetter   dao.ProduceGetter
	produceInserter dao.ProduceInserter
	produceDeleter  dao.ProduceDeleter
}

func NewProduceService() *ProduceService {
	dao := dao.NewProduceDao()
	return &ProduceService{
		produceGetter:   dao,
		produceInserter: dao,
		produceDeleter:  dao,
	}
}

// Implements service.ProduceGetter
func (s *ProduceService) GetProduce() ([]*ProduceItem, error) {
	panic("not implemented")
}

// Implements service.ProduceInserter
func (s *ProduceService) InsertProduce(item *ProduceItem) (*ProduceItem, error) {
	panic("not implemented")
}

// Implements service.ProduceDeleter
func (s *ProduceService) DeleteProduce(code string) (bool, error) {
	panic("not implemented")
}
