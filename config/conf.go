package config

import . "supermarket/types"

type ConfigData struct {
	InitialProduce []ProduceItem
	ListenPort     int
}

// This variable is set in main() using the configuration
// data read from config.json file in the project root.
var Config ConfigData
