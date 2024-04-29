package config

import "github.com/fatih/color"

func Green(message string) {
	_, _ = color.New(color.FgGreen).Add(color.Bold).Println(message)
}

func Red(message string) {
	_, _ = color.New(color.FgRed).Add(color.Bold).Add(color.Underline).Println(message)
}

func Yellow(message string) {
	_, _ = color.New(color.FgYellow).Add(color.Bold).Println(message)
}
