package main

import "strings"

func cleanInput(text string) []string {
	return strings.Split(strings.Trim(text, " "), " ")
}
