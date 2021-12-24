package types

type ProduceItem struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}

func (p *ProduceItem) GetName() string {
	return p.Name
}

func (p *ProduceItem) GetCode() string {
	return p.Code
}

func (p *ProduceItem) GetPrice() float64 {
	return p.Price
}
