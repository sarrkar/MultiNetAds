package helper

import "math/rand"

func RandStr(length int) string {
	b := make([]byte, length)
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
