package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	rand.New(src)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var string_builder strings.Builder
	string_builder.Grow(n)
	for i := 0; i < n; i++ {
		string_builder.WriteByte(characters[RandomInt(0, int64(len(characters)-1))])
	}
	return string_builder.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "INR"}
	n := len(currencies)
	index := RandomInt(0, int64(n-1))
	return currencies[index]
}
