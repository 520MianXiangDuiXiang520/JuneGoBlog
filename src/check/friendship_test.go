package check

import (
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestFriendAddCheck(t *testing.T) {
	c := new(gin.Context)

	if _, err := FriendAddCheck(c, &message.FriendAddReq{}); err == nil {
		t.Error("Check FILE")
	}

	if _, err := FriendAddCheck(c, &message.FriendAddReq{
		SiteName: "xxx",
	}); err == nil {
		t.Error("Check FILE")
	}

	if _, err := FriendAddCheck(c, &message.FriendAddReq{
		SiteLink: "xxx",
	}); err == nil {
		t.Error("Check FILE")
	}

	if _, err := FriendAddCheck(c, &message.FriendAddReq{
		SiteLink: "xxx",
		SiteName: "xxx",
	}); err != nil {
		t.Error("Check FILE")
	}


}