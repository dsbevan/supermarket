package handler

import (
	"regexp"
	"strconv"
	"strings"
)

// Request validation functions

// Alphanumeric and len() > 0
var nameExp = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func validName(name string) bool {
	return nameExp.MatchString(name)
}

// Alphanumeric ####-####-####-####
var codeExp = regexp.MustCompile(`^([a-zA-Z0-9]{4}-){3}[a-zA-Z0-9]{4}$`)

func validCode(code string) bool {
	return codeExp.MatchString(code)
}

// Check that number of decimal places <= 2
func validPrice(price float64) bool {
	//Format as string representing exact val of 32 bit float
	stringVal := strconv.FormatFloat(price, 'f', -1, 32)
	numDecimals := len(stringVal) - strings.IndexRune(stringVal, '.') - 1
	if numDecimals > 2 {
		return false
	}
	return true
}
