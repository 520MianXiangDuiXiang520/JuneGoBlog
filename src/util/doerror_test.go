package util

import (
	"strconv"
	"testing"
)

func TestCatchException(t *testing.T) {
	s := []string{
		"1", "2", "3", "4",
	}
	for _, v := range s {
		_, e := strconv.Atoi(v)
		CatchException(e)
	}
}
