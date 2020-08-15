package check

import (
	"JuneGoBlog/src/message"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestFriendAddCheck(t *testing.T) {
	c := new(gin.Context)

	if _, err := FriendApplicationCheck(c, &message.FriendApplicationReq{}); err == nil {
		t.Error("Check FILE")
	}

	if _, err := FriendApplicationCheck(c, &message.FriendApplicationReq{
		SiteName: "xxx",
	}); err == nil {
		t.Error("Check FILE")
	}

	if _, err := FriendApplicationCheck(c, &message.FriendApplicationReq{
		SiteLink: "xxx",
	}); err == nil {
		t.Error("Check FILE")
	}

	if _, err := FriendApplicationCheck(c, &message.FriendApplicationReq{
		SiteLink: "xxx",
		SiteName: "xxx",
	}); err != nil {
		t.Error("Check FILE")
	}

}
