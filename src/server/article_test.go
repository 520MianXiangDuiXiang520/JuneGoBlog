package server

import (
	"testing"
)

func TestGetAbstract(t *testing.T) {
	r := getAbstract(`# 震惊，醋和酱油竟然有如此妙用！

酱油和醋混合是怎么回事呢？酱油和醋相信大家都很熟悉，但是酱油和醋混合是怎么回事呢，下面就让小编带大家一起了解吧。

<!-- more -->

`)
	if r != "酱油和醋混合是怎么回事呢？酱油和醋相信大家都很熟悉，但是酱油和醋混合是怎么回事呢，下面就让小编带大家一起了解吧。" {
		t.Error("Fail")
	}
}
