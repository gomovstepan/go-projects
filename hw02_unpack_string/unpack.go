package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(stroka string) (string, error) {
	massive := []rune(stroka)
	if len(massive) == 0 {
		return "", nil
	}
	if _, err := strconv.Atoi(string(massive[0])); err == nil {
		return "", ErrInvalidString
	}
	ResNewLine, err := Iteration(massive)
	if err != nil {
		return "", err
	}

	return ResNewLine, nil
}

func Iteration(massive []rune) (string, error) {
	var NewString strings.Builder
	var previous string
	var ecran bool
	var skipDigitsRule bool
	var previousDigit bool

	for i, val := range massive {
		if string(val) == `\` && !ecran {
			if i > 0 {
				NewString.WriteString(previous)
			}
			previous = ""
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
			previous = string(val)
			skipDigitsRule, ecran = true, false
			if i+1 == len(massive) {
				NewString.WriteString(previous)
			}
			continue
		}
		if digit, err := strconv.Atoi(string(val)); err == nil {
			if previousDigit && !skipDigitsRule {
				return "", ErrInvalidString
			} else {
				NewString.WriteString(strings.Repeat(previous, digit))
				skipDigitsRule, ecran, previousDigit = false, false, true
				previous = ""
			}
		} else {
			if ecran {
				return "", ErrInvalidString
			}
			NewString.WriteString(previous)
			previousDigit, skipDigitsRule = false, false
			previous = string(val)
			ecran = false
			if i+1 == len(massive) {
				NewString.WriteString(previous)
			}
			ecran = false
		}
	}
	return NewString.String(), nil
}
