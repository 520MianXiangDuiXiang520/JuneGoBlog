package dao

import "testing"

func TestQueryAllFriendLink(t *testing.T) {
	var fls []FriendShipLink
	if err := QueryAllFriendLink(&fls); err != nil {
		t.Error("Error!")
	}
	if len(fls) <= 0 {
		t.Error("没拿到")
	}
	if fls[0].SiteName == "" {
		t.Error("Nil")
	}
}
