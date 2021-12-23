package service

import (
	"supermarket/dao"
	. "supermarket/types"
	"sync"
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
func (s *ProduceService) GetProduce() []ProduceItem {
	c := make(chan []ProduceItem, 1)
	go s.getProduceRoutine(c)
	res := <-c
	return res
}

func (s *ProduceService) getProduceRoutine(c chan []ProduceItem) {
	c <- s.produceGetter.GetProduce()
}

// Implements service.ProducePoster
func (s *ProduceService) PostProduce(items []ProduceItem) []ProduceItem {
	numItems := len(items)
	wg := &sync.WaitGroup{}
	wg.Add(numItems)

	c := make(chan *ProduceItem, numItems)
	for _, item := range items {
		go s.postProduceRoutine(c, wg, item)
	}

	// Wait for all routines to finish
	wg.Wait()

	// Build a new slice of successfully posted items
	res := make([]ProduceItem, 0, numItems)
	for i := 0; i < numItems; i++ {
		result := <-c
		if result != nil {
			res = append(res, *result)
		}
	}
	return res
}

func (s *ProduceService) postProduceRoutine(c chan *ProduceItem, wg *sync.WaitGroup, item ProduceItem) {
	posted := s.producePoster.PostProduce(item)
	if posted {
		c <- &item
	} else {
		c <- nil
	}
	wg.Done()
}

// Implements service.ProduceDeleter
func (s *ProduceService) DeleteProduce(code string) bool {
	c := make(chan bool)
	go s.deleteProduceRoutine(c, code)
	res := <-c
	return res
}

func (s *ProduceService) deleteProduceRoutine(c chan bool, code string) {
	c <- s.produceDeleter.DeleteProduce(code)
}
