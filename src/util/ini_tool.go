package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func ReadLines(filePath string) []string {
	result := make([]string, 0)
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil && err != io.EOF {
			log.Println(err)
			panic("Read Error!!")
		}
		result = append(result, line)
		if err == io.EOF {
			break
		}
	}
	return result
}

func Load(iniPath, block string, s interface{}) {
	// 1. 读取ini文件
	_, currentfile, _, _ := runtime.Caller(1) // 忽略错误
	filename := path.Join(path.Dir(currentfile), iniPath)
	lines := ReadLines(filename)
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	offset := 0
	for i, line := range lines {
		pat := "(\\[" + block + "\\]).*"
		patt, _ := regexp.Compile(pat)
		ok := patt.MatchString(line)
		if !ok {
			continue
		}
		offset = i
	}
	lines = lines[offset:]
	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}
		pat := "(\\[" + block + "\\]).*"
		patt, _ := regexp.Compile(pat)
		ok := patt.MatchString(line)
		if string(line[0]) == "[" && ok {
			continue
		}
		if string(line[0]) == "[" && !ok {
			break
		}
		line = strings.Trim(line, "\n")
		line = strings.Trim(line, "\r")
		lr := strings.Split(line, "=")
		if len(lr) != 2 {
			panic("ini 格式错误")
		}
		for i := 0; i < t.Elem().NumField(); i++ {
			field := t.Elem().Field(i)
			s := field.Tag.Get("ini")
			if s == lr[0] {
				switch field.Type.Kind() {
				case reflect.String:
					v.Elem().Field(i).SetString(lr[1])
				case reflect.Int:
					a, b := strconv.Atoi(lr[1])
					if b != nil {
						panic("ini 格式错误")
					}
					v.Elem().Field(i).SetInt(int64(a))
				}
			}
		}
	}
}
