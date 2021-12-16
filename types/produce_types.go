package types

import "strings"

type ProduceItem struct {
	name  string  `json:"name"`
	code  []rune  `json:"code"`
	price float32 `json:"price"`
}

func (p *ProduceItem) GetName() string {
	return p.name
}

func (p *ProduceItem) GetCode() string {
	return string(p.code[:])
}

func (p *ProduceItem) GetPrice() float32 {
	return p.price
}

func NewProduceItem(name *string, code *string, price float32) *ProduceItem {
	// Name, 19-char code (4 4-char sequences separated by '-'), and nonnegaticve price are required
	if name == nil || code == nil || len(*code) != 19 || price < 0 {
		return nil
	}

	item := ProduceItem{
		name:  *name,
		code:  []rune(strings.ToLower(*code)),
		price: price,
	}
	return &item
}
