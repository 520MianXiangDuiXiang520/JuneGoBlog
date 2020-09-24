package util

import (
	"time"
)

func Int64ToTime(u int64) time.Time {
	return time.Unix(u, 0)
}
