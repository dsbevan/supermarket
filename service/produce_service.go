package service

import (
	"supermarket/dao"
	. "supermarket/types"
)

type ProduceService struct {
	produceInserter dao.ProduceInserter
	produceGetter   dao.ProduceGetter
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

func (s *ProduceService) GetProduce() ([]*ProduceItem, error) {
	panic("not implemented")
}

func (s *ProduceService) CreateProduce() (*ProduceItem, error) {
	panic("not implemented")
}

func (s *ProduceService) DeleteProduce() (bool, error) {
	panic("not implemented")
}
