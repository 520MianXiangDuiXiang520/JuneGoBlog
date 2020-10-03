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
	email2 := "15264897862@162.com"
	email3 := "1233@111"
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
