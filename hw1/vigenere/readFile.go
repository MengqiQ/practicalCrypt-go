package vigenere

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func ReadFile(filename string) string {
	content, _ := ioutil.ReadFile(filename)

	reg, _ := regexp.Compile("[^a-zA-Z]+")

	processedString := strings.ToUpper(reg.ReplaceAllString(string(content), ""))

	return processedString
}
