package util

import (
	"regexp"
	"strconv"
	"time"
)

func Int64ToTime(u int64) time.Time {
	return time.Unix(u, 0)
}

func IntString2Time(str string) time.Time {
	uTime, _ := strconv.Atoi(str)
	return time.Unix(int64(uTime), 0)
}

func TimeString2Time(str string) time.Time {
	// Mon, 02 Jan 2006 15:04:05 MST"
	t, _ := time.Parse("2006-01-02 15:04:05 +0000 UTC", str)
	return t
}

func IsEmail(email string) bool {
	if len(email) <= 0 {
		return false
	}
	emailRules := `^[A-Za-z0-9_.]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	ok, err := regexp.MatchString(emailRules, email)
	return ok && err == nil
}
