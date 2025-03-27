package passgen

import (
	"errors"
	"math/rand"
)

var categories = []string{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"0123456789",
	"!@#$%^&*()_+-=",
}

func Generate(length int) (string, error) {
	if length < 0 || length > 32 {
		return "", errors.New("password length must be in [1; 255] range")
	}

	characters := make([]byte, length)

	for i := range len(characters) {
		categoryIndex := 0
		if i < len(categories) {
			// Make sure that every category is present in the password
			categoryIndex = i
		} else {
			categoryIndex = rand.Intn(len(categories))
		}
		category := categories[categoryIndex]

		charIndex := rand.Intn(len(category))

		characters[i] = category[charIndex]
	}

	// Shuffle selected characters
	rand.Shuffle(len(characters), func(i, j int) {
		characters[i], characters[j] = characters[j], characters[i]
	})

	return string(characters), nil
}
