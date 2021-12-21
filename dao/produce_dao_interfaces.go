package dao

import . "supermarket/types"

type ProduceInserter interface {
	InsertProduce(item ProduceItem) bool
}

type ProduceGetter interface {
	GetProduce(code string) []*ProduceItem
}

type ProduceDeleter interface {
	DeleteProduce(code string) bool
}
