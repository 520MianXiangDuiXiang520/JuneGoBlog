package util

import (
	"log"
	"testing"
)

func TestGetHashWithTimeUUID(t *testing.T) {
	str := "Junebao"
	for i := 0; i < 10; i++ {
		log.Println(GetHashWithTimeUUID(str))
	}
}

func TestSha256(t *testing.T) {
	log.Println(Sha256("test"))
}
