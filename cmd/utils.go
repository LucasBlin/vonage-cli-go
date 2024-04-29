package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func ValidNumber(number string) bool {
	if len(number) != 10 {
		return false
	}

	for _, char := range number {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	match, _ := regexp.MatchString(`[a-zA-Z]`, number)
	if match {
		return false
	}

	return true
}

func AsEuro(str string) (string, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("€%.2f", val), nil
}
