package types

import (
	"regexp"
	"strconv"
	"strings"
)

// Type validation
func ValidItem(item ProduceItem) bool {
	if !ValidCode(item.Code) || !ValidName(item.Name) || !ValidPrice(item.Price) {
		return false
	}
	return true
}

func AllValidItems(items []ProduceItem) bool {
	for _, item := range items {
		if !ValidItem(item) {
			return false
		}
	}
	return true
}

// CONDITIONS
// Alphanumeric, does not start with a space, and len() > 0
var nameExp = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9 ]*$`)

func ValidName(name string) bool {
	return nameExp.MatchString(name)
}

// Alphanumeric ####-####-####-####
var codeExp = regexp.MustCompile(`^([a-zA-Z0-9]{4}-){3}[a-zA-Z0-9]{4}$`)

func ValidCode(code string) bool {
	return codeExp.MatchString(code)
}

// Check that number of decimal places <= 2
func ValidPrice(price float64) bool {
	//Format as string representing exact val of 32 bit float
	stringVal := strconv.FormatFloat(price, 'f', -1, 32)
	numDecimals := len(stringVal) - strings.IndexRune(stringVal, '.') - 1
	if numDecimals > 2 {
		return false
	}
	return true
}
