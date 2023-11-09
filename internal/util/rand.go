package util

import (
	"bytes"
	"math/rand"
	"time"
)

const (
	StrAll = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetRandomString(length int) string {
	result := bytes.NewBufferString("")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result.WriteString(string(StrAll[r.Intn(62)]))
	}
	return result.String()
}
