package old

import (
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	u, o := GetUser("a", "b")
	log.Println(u, o)
}
