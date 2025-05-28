package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(sm string) (string, error) {
	if len(sm) == 0 {
		return "", errors.New("empty string")
	}
	var res string
	for _, ch := range sm {
		if ch != '.' && ch != '-' && ch != ' ' {
			res = morse.ToMorse(sm)
		} else {
			res = morse.ToText(sm)
		}
	}
	return res, nil
}
