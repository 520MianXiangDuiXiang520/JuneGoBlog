package util

import (
	"log"
	"math/rand"
	"testing"
)

func TestTimeString2Time(t *testing.T) {
	time := TimeString2Time("2020-01-31 00:00:00 +0000 UTC")
	log.Println(time.Unix() + rand.Int63n(3000))
}

func TestIsEmail(t *testing.T) {
	email1 := "17719@st.nuc.edu.cn"
	email2 := "15264897862_ll.n@162.com"
	email3 := "1233__l@111"
	if !IsEmail(email1) {
		t.Error("email1 mis judgment")
	}
	if !IsEmail(email2) {
		t.Error("email2 mis judgment")
	}
	if IsEmail(email3) {
		t.Error("email3 mis judgment")
	}
}

func TestRemoveTitle(t *testing.T) {
	text1 := `
# 测试
this is a test text
## title2
`
	if res := RemoveTitle(text1); res != "this is a test text" {
		t.Error("test1: ", res)
	}

	text2 := `
this is a test text
`
	if res := RemoveTitle(text2); res != "this is a test text" {
		t.Error("test1: ", res)
	}

	text3 := `
this is a test text
## title
`
	if res := RemoveTitle(text3); res != "this is a test text" {
		t.Error("test1: ", res)
	}
}
