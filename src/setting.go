package src

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var Setting setting

type setting struct {
	AbstractLen int    `json:"abstractLen"` // 默认的摘要长度
	Redis       bool   `json:"redis"`       // 是否使用redis
	MyEmail     string `json:"my_email"`
}

func readJson(filePath string) (result string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')
		result += s
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read is ok")
				break
			} else {
				fmt.Println("ERROR:", err)
				return
			}
		}
	}
	return result
}

func InitSetting(settingPath string) {
	once := sync.Once{}
	once.Do(func() {
		result := readJson(settingPath)
		err := json.Unmarshal([]byte(result), &Setting)
		if err != nil {
			log.Printf("配置文件加载失败！")
			panic(err)
		}
	})
}
