package service

import . "supermarket/types"

type ProduceGetter interface {
	GetProduce() ([]*ProduceItem, error)
}

type ProduceCreator interface {
	CreateProduce() (*ProduceItem, error)
}

type ProduceDeleter interface {
	DeleteProduce() (bool, error)
}
