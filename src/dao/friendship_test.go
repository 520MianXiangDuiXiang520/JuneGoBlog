package dao

import "testing"

func TestAddFriendship(t *testing.T) {
	// test FriendshipLink has ImgLink and Intro
	err := AddFriendship(&FriendShipLink{
		SiteLink: "https://draveness.me/",
		SiteName: "面向信仰编程",
		ImgLink: "https://draveness.me/images/draven-logo.png",
		Intro: "面向信仰编程",
	})
	if err != nil {
		t.Error(err)
	}
	// test FriendshipLink without ImgLink and Intro
}
