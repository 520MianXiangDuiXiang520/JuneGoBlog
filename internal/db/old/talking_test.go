package old

import (
	"log"
	"testing"
)

func TestHasTalk(t *testing.T) {
	unExist := 10
	exist := 4
	if HasTalk(unExist) {
		t.Error("Fail Test")
	}
	if !HasTalk(exist) {
		t.Error("Fail Test")
	}
}

func TestQueryTalksByArticleIDLimit(t *testing.T) {
	talks, e := QueryTalksByArticleIDLimit(65, 2, 10)
	if e != nil {
		t.Error(e)
	}
	if len(talks) <= 0 {
		t.Error("result is empty")
	}
	log.Println(talks)
}
