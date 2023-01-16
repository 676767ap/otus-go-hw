package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inptStr string) (string, error) {
	first, _ := utf8.DecodeRuneInString(inptStr)
	if unicode.IsDigit(first) {
		return "", ErrInvalidString
	}
	var sliceStr []string
	ra := []rune(inptStr)
	for i, r := range ra {
		if unicode.IsDigit(r) {
			if unicode.IsDigit(ra[i-1]) {
				return "", ErrInvalidString
			}
			var numRunes string
			curNum, _ := strconv.Atoi(string(r))
			if curNum != 0 {
				numRunes = strings.Repeat(string(ra[i-1]), curNum)
				sliceStr = append(sliceStr, numRunes)
			}
			continue
		}
		if i < (len(ra) - 1) {
			if !unicode.IsDigit(ra[i+1]) {
				sliceStr = append(sliceStr, string(r))
			}
		} else {
			sliceStr = append(sliceStr, string(r))
		}
	}
	return strings.Join(sliceStr, ""), nil
}
