package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// конвертация текст или азбука морзе
func Convert(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("empty string")
	}
	res := StringOrMorse(str)
	return res, nil
}

// проверка текст или морзе
func StringOrMorse(sm string) string {
	if strings.ContainsFunc(sm, func(r rune) bool {
		return r != '-' && r != '.' && !strings.ContainsRune(" \t\n", r)
	}) {
		return morse.ToMorse(sm)
	}
	return morse.ToText(sm)

}
