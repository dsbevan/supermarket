package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"supermarket/config"
	"supermarket/data"
	"supermarket/types"
)

// Loads config data from the given file. Exits with code 1 if
// config file structure or data is invalid.
func LoadConfigFatal(filename string) {
	// Load data from config file
	if jsonConfigBytes, err := ioutil.ReadFile(filename); err != nil {
		// Error reading config file
		fmt.Fprintln(os.Stderr, "Error reading config file.")
		os.Exit(1)
	} else {
		json.Unmarshal(jsonConfigBytes, &config.Config)

		// Validate initial produce items
		if !types.AllValidItems(config.Config.InitialProduce) {
			fmt.Fprintln(os.Stderr, "Invalid values in config file.")
			os.Exit(1)
		}
		// Store initial produce items
		for _, item := range config.Config.InitialProduce {
			data.ProduceSlice = append(data.ProduceSlice, item)
		}
	}
}
