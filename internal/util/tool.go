package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

// RemoveTitle 用来删除 text 中的标题行
// 标题行指以 1 到 6 个 # 开头，后面紧跟一个空格的行
func RemoveTitle(text string) string {
	res := ""
	for _, v := range strings.Split(text, "\n") {
		if !regexp.MustCompile("^[#]{1,6} .*").MatchString(v) {
			res = fmt.Sprintf("%s%s", res, v)
		}
	}
	return res
}
