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

//func TestUpdateFriendStatusByID(t *testing.T) {
//	if err := UpdateFriendStatusByID(2, 1); err != nil {
//		t.Error("Update Error")
//	}
//	fl, _ := HasFriendLinkByID(2)
//	if fl.Status != 1 {
//		t.Error("Update Error!!")
//	}
//}

func TestHasFriendLinkByID(t *testing.T) {
	fl, ok := HasFriendLinkByID(1)
	if ok {
		if fl.SiteName != "DeepBlue的小站" {
			t.Error("get friendship error")
		}
	} else {
		t.Error("Error!!")
	}

	fl, ok = HasFriendLinkByID(20)
	if ok {
		t.Error("Error!!")
	}
}
