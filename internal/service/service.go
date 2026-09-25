package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertAuto(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("empty input")
	}

	isInvalidChar := func(r rune) bool {
		return r != '.' && r != '-' && r != ' ' && r != '\n' && r != '\t'
	}

	if strings.ContainsFunc(input, isInvalidChar) {
		return morse.ToMorse(input), nil
	}
	return morse.ToText(input), nil
}
