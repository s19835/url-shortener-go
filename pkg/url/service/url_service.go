package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/s19835/url-shortener-go/internal/util"
)

// create a unique short code from long url
func GenerateShortCode(longURL string, retry int) (string, error) {
	if retry > 3 {
		return "", errors.New("max retries limit reached for generating short code")
	}

	// Step 1: Hash the URL (MD5/SHA-1)
	hash := md5.Sum([]byte(longURL))
	hexHash := hex.EncodeToString(hash[:])

	// Step 2: Truncate to first 8 chars (adjust for collision probability)
	truncatedHash := hexHash[:8]

	// Step 3: Convert hex to uint64 (for Base62 encoding)
	decimalValue, err := hexToUint64(truncatedHash)
	if err != nil {
		return "", err
	}

	// Step 4: Encode to Base62
	shortCode := util.Base62Encode(decimalValue)

	// Step 5: Check for collision (e.g., in DB/Redis)
	// if exist, _ := checkIfShortCodeExist(); exist {}

	return shortCode, nil

}

func hexToUint64(hexStr string) (uint64, error) {
	var num uint64
	_, err := fmt.Sscanf(hexStr, "%x", &num)
	return num, err
}
