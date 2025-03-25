package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"unicode"
)

const (
	base62Chars     = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	shortCodeLength = 6 //set max code length for shortcode
)

// convert uint64 num to base64 string
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

// generate short code from url
func GenerateShortCode(longURL string) (string, error) {
	// normalize provided long URL
	normalizedURL, err := normalizeURL(longURL)
	if err != nil {
		return "", err
	}

	// step 1: Hash the URL (MD5/SHA-1)
	hash := md5.Sum([]byte(normalizedURL))
	hexHash := hex.EncodeToString(hash[:])

	// step 2: Truncate to first 8 characters to reduce collision prob
	truncatedHash := hexHash[:8]

	// step 3: convert hex to uint64 for base62 encoding
	var num uint64
	_, err = fmt.Sscanf(truncatedHash, "%x", &num)
	if err != nil {
		return "", err
	}

	fmt.Println(num)

	return Base62Encode(num), nil
}

// normalize URL for consistant hashing
func normalizeURL(rawURL string) (string, error) {
	// add schem if missing
	if !strings.Contains(rawURL, "://") {
		rawURL = "http://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// standardized scheme and host
	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)

	// Remove default ports
	u.Host = strings.TrimSuffix(u.Host, ":80")
	u.Host = strings.TrimSuffix(u.Host, ":443")

	// Sort Query parameters?

	return u.String(), nil
}

// validate short code
func ValidateShortCode(shortCode string) error {
	if len(shortCode) < 3 || len(shortCode) > 8 {
		return errors.New("Code must be 3-8 character long.")
	}

	for _, r := range shortCode {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return errors.New("Code must contain letters and numbers.")
		}
	}

	return nil
}

// helper function for base62 encoding
func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
