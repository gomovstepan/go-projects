package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(stroka string) (string, error) {
	var NewString strings.Builder
	massive := []rune(stroka)
	if len(massive) == 0 {
		return "", nil
	}
	if _, err := strconv.Atoi(string(massive[0])); err == nil {
		return "", ErrInvalidString
	}

	var previos string
	var ecran bool
	var skipDigitsRule bool
	var previosDigit bool

	for i, val := range massive {
		if string(val) == `\` && !ecran {
			if i > 0 {
				NewString.WriteString(previos)
			}
			previos = ""
			ecran = true
			if i+1 == len(massive) {
				return "", ErrInvalidString
			}
			continue
		}
		if ecran {
			if unicode.IsLetter(val) {
				return "", ErrInvalidString
			}
			previos = string(val)
			skipDigitsRule, ecran = true, false
			if i+1 == len(massive) {
				NewString.WriteString(previos)
			}
			continue
		}
		if digit, err := strconv.Atoi(string(val)); err == nil {
			if previosDigit && !skipDigitsRule {
				return "", ErrInvalidString
			} else {
				NewString.WriteString(strings.Repeat(previos, digit))
				skipDigitsRule, ecran, previosDigit = false, false, true
				previos = ""
			}
		} else {
			if ecran {
				return "", ErrInvalidString
			}
			NewString.WriteString(previos)
			previosDigit, skipDigitsRule = false, false
			previos = string(val)
			ecran = false
			if i+1 == len(massive) {
				NewString.WriteString(previos)
			}
			ecran = false
		}
	}
	return NewString.String(), nil
}
