package utils

import "time"
import "math/rand"

func GenerateSixDigitOtp() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rng.Intn(900000) + 100000
}