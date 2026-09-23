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

	isPureMorse := true

	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\t' {
			isPureMorse = false
			break
		}
	}

	if isPureMorse {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}
