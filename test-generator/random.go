package test_generator

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return randomStringWithCharset(length, charset)
}

func RandomLowercaseString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	return randomStringWithCharset(length, charset)
}

func RandomUppercaseString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return randomStringWithCharset(length, charset)
}

func randomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RandomUUID() string {
	return uuid.New().String()
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
