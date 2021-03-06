package util

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// 新建ini，写入信息
	fw, err := os.Create("./test.ini")
	if err == nil {
		_, _ = fw.WriteString("[mysql]\nuser=testName\npassword=testPassword\n[redis]\npassword=redisPassword")

	}
	// 执行Load
	var s = struct {
		User     string `ini:"user"`
		Password string `ini:"password"`
	}{}
	var rs = struct {
		Password string `ini:"password"`
	}{}
	Load("./test.ini", "mysql", &s)
	Load("./test.ini", "redis", &rs)
	// 断言

	if s.User != "testName" || s.Password != "testPassword" {
		t.Error("LOAD ERROR")
	}

	if rs.Password != "redisPassword" {
		t.Error("LOAD ERROR2" + rs.Password)
	}

	// 删除文件

	defer func() {
		fw.Close()
		err = os.Remove("./test.ini")
		if err != nil {
			t.Error("CAN NOT REMOVE TEST FILE!!")
			t.Error(err)
		}
	}()

}
