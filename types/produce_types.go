package types

type ProduceItem struct {
	Name  string   `json:"name"`
	Code  [19]rune `json:"code"`
	Price float32  `json:"price"`
}
