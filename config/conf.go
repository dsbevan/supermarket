package config

type ConfigData struct {
	ArraySize int
}

// This variable is set in main() using the configuration
// data read from config.json file in the project root.
var Config ConfigData
