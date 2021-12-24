package data

import . "supermarket/types"

const PRODUCE_CAPACITY int = 20

var Produce [PRODUCE_CAPACITY]ProduceItem
var ProduceSlice []ProduceItem = make([]ProduceItem, 0, PRODUCE_CAPACITY)
