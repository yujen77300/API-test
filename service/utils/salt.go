package utils

import (
	"math/rand"
	"time"
)

func GenerateSalt() string {
	var randomSalt string
	salt := "!@#$%^&*()"

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		randomSalt += string(salt[rand.Intn(10)])
	}

	return randomSalt
}
