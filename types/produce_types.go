package types

type ProduceItem struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float32 `json:"price"`
}

func (p *ProduceItem) GetName() string {
	return p.Name
}

func (p *ProduceItem) GetCode() string {
	return p.Code
}

func (p *ProduceItem) GetPrice() float32 {
	return p.Price
}

//func NewProduceItem(name *string, code *string, price float32) *ProduceItem {
//	// Name, 19-char code (4 4-char sequences separated by '-'), and nonnegaticve price are required
//	if name == nil || code == nil || len(*code) != 19 || price < 0 {
//		return nil
//	}
//
//	item := ProduceItem{
//		name:  *name,
//		code:  []rune(strings.ToUpper(*code)),
//		price: price,
//	}
//	return &item
//}
