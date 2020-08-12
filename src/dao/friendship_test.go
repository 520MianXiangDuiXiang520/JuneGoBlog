package dao

import "testing"

func TestQueryAllFriendLink(t *testing.T) {
	fl := make([]FriendShipLink, 0)
	if err := QueryAllFriendLink(&fl); err != nil {
		t.Error("QueryAllFriendLink ERROR!!!")
	}

	if len(fl) <= 0 {
		t.Error("NO Result Query!!")
	}
}

func TestAddFriendship(t *testing.T) {

}