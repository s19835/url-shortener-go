package util

import (
	"strings"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// converts and integer to base62 string
func Base62Encode(num uint64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var encoded strings.Builder

	for num > 0 {
		remainder := num % 62
		encoded.WriteByte(base62Chars[remainder])
		num /= 62
	}

	return reverseString(encoded.String())
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j+1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
