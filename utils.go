package main

import (
	"fmt"
	"lucas.blin/v2/config"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ValidNumber(number string) bool {
	if len(number) != 11 {
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

func Contact(argv *ArgT) {
	c := strings.Split(argv.Contact, ",")
	if ValidNumber(c[1]) == true {
		if err := config.SetConfig("contacts", c[0], c[1], argv.ConfigPath); err != nil {
			config.Red("❌  Fail to add the contact to the configuration.")
		}
		config.Green("✅  Contact added to the configuration.")
	} else {
		config.Red("❌  Invalid number, please provide it as : 336XXXXXXXX")
	}
}
