package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
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

func HashByMD5(strList []string) (h string) {
	r := strings.Join(strList, "")
	hash := md5.New()
	hash.Write([]byte(r))
	return hex.EncodeToString(hash.Sum(nil))
}
