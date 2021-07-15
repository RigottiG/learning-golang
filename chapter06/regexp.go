package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "Guilherme is 21 years old"
	expression := regexp.MustCompile("\\d")

	fmt.Println(expression.ReplaceAllString(text, "3"))

	text = "guilherme rigotti"
	expression = regexp.MustCompile("\\b\\w")

	uppercasedName := expression.ReplaceAllStringFunc(
		text,
		func(s string) string {
			return strings.ToUpper(s)
		},
	)

	fmt.Println(uppercasedName)
}
