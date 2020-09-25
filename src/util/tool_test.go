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
