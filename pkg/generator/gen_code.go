package generator

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenCode(length int, chars string) (string, error) {
	if length <= 0 {
		return "", errMustBeGTZero
	}
	if len(chars) == 0 {
		return "", errCannotEmpty
	}

	rand.NewSource(time.Now().UnixNano())
	result := make([]byte, length)

	for i := range length {
		// #nosec G404 - Using math/rand for non-cryptographic random generation
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result), nil
}

func GenUUID() string {
	return uuid.New().String()
}
