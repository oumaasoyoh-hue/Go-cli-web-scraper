package parser

import (
	"os"
	"regexp"
)

func ParseFIle(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	linkRegex := regexp.
}
