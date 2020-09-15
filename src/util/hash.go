package util

import (
	"crypto/sha256"
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func GetHashWithTimeUUID(str string) string {
	timeUnix := time.Now().UnixNano()
	rand := uuid.NewV4().String()
	sum := sha256.Sum256([]byte(fmt.Sprintf("%v%v%v", rand, str, timeUnix)))
	return fmt.Sprintf("%x", sum)
}
